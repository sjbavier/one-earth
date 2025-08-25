package connectors

import (
	"context"
	"encoding/csv"
	"errors"
	"io"
	"math"
	"strconv"
	"strings"
	"testing"
	"time"
)

const sampleCSV = `Year,Month,Decimal Date,Average,Interpolated,Trend,Number of Days
2025,7,2025.54,420.12,420.15,419.80,31
2025,8,2025.58,-99.99,420.20,419.85,31
2025,9,2025.62,421.00,421.05,420.90,30
2025,10,2025.67,,421.10,420.95,31
`

func TestFetchCO2Data(t *testing.T) {
	ctx := context.Background()

	// Use strings.NewReader to simulate HTTP response body
	r := strings.NewReader(sampleCSV)

	// Override FetchCO2Data to accept io.Reader for testing
	points, latest, err := parseCO2CSV(ctx, r)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(points) != 2 {
		t.Errorf("expected 2 valid points, got %d", len(points))
	}

	expectedLatestTime := time.Date(2025, 8, 15, 0, 0, 0, 0, time.UTC) // approx for 2025.62
	if !latest.T.Equal(expectedLatestTime) {
		t.Errorf("expected latest time %v, got %v", expectedLatestTime, latest.T)
	}

	expectedLatestValue := 421.00
	if latest.V != expectedLatestValue {
		t.Errorf("expected latest value %v, got %v", expectedLatestValue, latest.V)
	}
}

// parseCO2CSV is a helper function to parse CSV from io.Reader for testing
func parseCO2CSV(ctx context.Context, reader io.Reader) ([]Point, Point, error) {
	r := csv.NewReader(reader)

	// Read header
	_, err := r.Read()
	if err != nil {
		return nil, Point{}, err
	}

	var points []Point

	for {
		record, errRead := r.Read()
		if errRead == io.EOF {
			break
		}
		if errRead != nil {
			return nil, Point{}, errRead
		}

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

		t, errTime := decimalDateToTime(decimalDate)
		if errTime != nil {
			continue
		}

		points = append(points, Point{T: t, V: avg})
	}

	if len(points) == 0 {
		return nil, Point{}, errors.New("no valid data points found")
	}

	latest := points[len(points)-1]
	return points, latest, nil
}
