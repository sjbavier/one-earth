package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"one-earth-api/internal/connectors"
)

const noaaCO2URL = "https://gml.noaa.gov/webdata/ccgg/trends/co2/co2_mm_mlo.csv"

type CO2Response struct {
	Latest connectors.Point   `json:"latest"`
	Series []connectors.Point `json:"series"`
}

// MetricLatest handles GET /api/metrics/co2

func MetricLatest(w http.ResponseWriter, r *http.Request) {
	latest, _, err := connectors.FetchCO2DataCached(context.Background(), noaaCO2URL)
	if err != nil {
		// Return JSON error with 503
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusServiceUnavailable)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to fetch CO2 data and no cached data available"})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(latest)
}

// Series handles GET /api/series/co2?days=30
func Series(w http.ResponseWriter, r *http.Request) {
	daysStr := r.URL.Query().Get("days")
	days := 30
	if daysStr != "" {
		if d, err := strconv.Atoi(daysStr); err == nil && d > 0 {
			days = d
		}
	}

	_, all, err := connectors.FetchCO2DataCached(context.Background(), noaaCO2URL)
	if err != nil {
		// Return JSON error with 503
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusServiceUnavailable)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to fetch CO2 data and no cached data available"})
		return
	}

	cutoff := time.Now().AddDate(0, 0, -days)
	filtered := []connectors.Point{}
	for _, p := range all {
		if p.T.After(cutoff) {
			filtered = append(filtered, p)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(filtered)
}
