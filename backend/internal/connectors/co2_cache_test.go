package connectors

import (
	"context"
	"net/http"
	"testing"
	"time"
)

func TestFetchCO2DataCached(t *testing.T) {
	ctx := context.Background()

	// Start local HTTP server to serve sample CSV
	handler := func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "testdata/sample_co2.csv")
	}
	server := http.Server{
		Addr:    "127.0.0.1:8085",
		Handler: http.HandlerFunc(handler),
	}
	go server.ListenAndServe()
	defer server.Close()

	url := "http://127.0.0.1:8085/sample_co2.csv"

	// First call should fetch fresh data and cache it
	latest1, all1, err1 := FetchCO2DataCached(ctx, url)
	if err1 != nil {
		t.Fatalf("FetchCO2DataCached failed: %v", err1)
	}
	if len(all1) == 0 {
		t.Fatal("Expected non-empty data slice")
	}

	// Second call should hit cache
	latest2, all2, err2 := FetchCO2DataCached(ctx, url)
	if err2 != nil {
		t.Fatalf("FetchCO2DataCached failed on second call: %v", err2)
	}

	if latest1.T != latest2.T || latest1.V != latest2.V {
		t.Fatal("Expected latest points to be equal from cache")
	}

	if len(all1) != len(all2) {
		t.Fatal("Expected all data slices to be equal length from cache")
	}

	// Simulate cache expiration by setting TTL to very short and waiting
	memCache.SetTTL(10 * time.Millisecond)
	time.Sleep(20 * time.Millisecond)

	// Third call should fetch fresh data again
	_, _, err3 := FetchCO2DataCached(ctx, url)
	if err3 != nil {
		t.Fatalf("FetchCO2DataCached failed after cache expiration: %v", err3)
	}
}
