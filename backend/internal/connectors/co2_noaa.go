package connectors

import (
	"context"
	"encoding/csv"
	"errors"
	"io"
	"math"
	"net/http"
	"strconv"
	"time"
)

// Point represents a data point with time and value.
type Point struct {
	T time.Time
	V float64
}

// FetchCO2Data fetches the NOAA daily CO2 CSV data from Mauna Loa, parses it into []Point,
// skipping missing/flagged rows and dropping NaNs. It returns the latest point and the full slice.
func FetchCO2Data(ctx context.Context, url string) (latest Point, all []Point, err error) {
	const maxRetries = 5
	var resp *http.Response

	client := &http.Client{}

	// Retry with exponential backoff
	for i := 0; i < maxRetries; i++ {
		req, reqErr := http.NewRequestWithContext(ctx, "GET", url, nil)
		if reqErr != nil {
			err = reqErr
			return
		}

		resp, err = client.Do(req)
		if err == nil && resp.StatusCode == http.StatusOK {
			break
		}

		if resp != nil {
			resp.Body.Close()
		}

		backoff := time.Duration(math.Pow(2, float64(i))) * time.Second
		select {
		case <-time.After(backoff):
		case <-ctx.Done():
			err = ctx.Err()
			return
		}
	}

	if err != nil {
		return
	}

	defer resp.Body.Close()

	r := csv.NewReader(resp.Body)

	// Read header
	_, err = r.Read()
	if err != nil {
		return
	}

	var points []Point

	for {
		record, errRead := r.Read()
		if errRead == io.EOF {
			break
		}
		if errRead != nil {
			err = errRead
			return
		}

		// Expected CSV columns (example from NOAA):
		// Year,Month,Decimal Date,Average,Interpolated,Trend,Number of Days
		// We will parse Decimal Date (index 2) and Average (index 3)
		// Skip rows with missing or flagged data (Average == -99.99 or empty)

		if len(record) < 4 {
			continue
		}

		avgStr := record[3]
		if avgStr == "" || avgStr == "-99.99" {
			continue
		}

		avg, errParse := strconv.ParseFloat(avgStr, 64)
		if errParse != nil || math.IsNaN(avg) {
			continue
		}

		decimalDateStr := record[2]
		decimalDate, errParse := strconv.ParseFloat(decimalDateStr, 64)
		if errParse != nil {
			continue
		}

		// Convert decimal date to time.Time
		t, errTime := decimalDateToTime(decimalDate)
		if errTime != nil {
			continue
		}

		points = append(points, Point{T: t, V: avg})
	}

	if len(points) == 0 {
		err = errors.New("no valid data points found")
		return
	}

	latest = points[len(points)-1]
	all = points
	return
}

// decimalDateToTime converts a decimal year (e.g. 2025.65) to time.Time approximating the date.
func decimalDateToTime(decimalDate float64) (time.Time, error) {
	year := int(decimalDate)
	frac := decimalDate - float64(year)

	// Calculate the number of days in the year (accounting for leap years)
	daysInYear := 365
	if isLeapYear(year) {
		daysInYear = 366
	}

	// Calculate the day of year
	dayOfYear := int(frac * float64(daysInYear))

	// Construct time.Time from year and day of year
	t := time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC)
	t = t.AddDate(0, 0, dayOfYear)

	return t, nil
}

func isLeapYear(year int) bool {
	if year%400 == 0 {
		return true
	}
	if year%100 == 0 {
		return false
	}
	return year%4 == 0
}
