import React from "react";
import { VegaLite } from "react-vega";
import { useCO2Latest, useCO2Series } from "../hooks/useCO2";
import { format } from "date-fns";
import "../styles/theme.css";
import { co2Spec } from "../charts/co2.spec";

export function CO2Tile() {
  // Fetch latest CO2 metric
  const {
    data: latest,
    isLoading: loadingLatest,
    isError: errorLatest,
  } = useCO2Latest();

  // Fetch last 30 days CO2 series
  const {
    data: series,
    isLoading: loadingSeries,
    isError: errorSeries,
  } = useCO2Series(30);

  if (loadingLatest || loadingSeries) {
    return <div>Loading...</div>;
  }

  if (errorLatest || errorSeries || !latest || !series) {
    return <div>Data unavailable — retrying…</div>;
  }

  // Prepare data for VegaLite sparkline
  const vegaData = {
    values: series.map((point) => ({
      T: point.T,
      V: point.V,
    })),
  };

  // Format last updated timestamp in UTC
  const lastUpdatedUTC = format(
    new Date(latest.timestamp),
    "yyyy-MM-dd HH:mm 'UTC'"
  );

  return (
    <div className="p-4 bg-white dark:bg-gray-900 rounded-lg shadow-md max-w-sm">
      <div className="flex justify-between items-center mb-2">
        <span className="text-sm text-gray-500 dark:text-gray-400">
          Last updated: {lastUpdatedUTC}
        </span>
      </div>
      <div className="text-4xl font-semibold text-gray-900 dark:text-white mb-2">
        {latest.value.toFixed(1)} ppm
      </div>
      <div>
        <VegaLite spec={co2Spec} data={vegaData} />
      </div>
      <div className="mt-2 text-xs text-gray-400">
        Source: NOAA GML — Public Domain
      </div>
    </div>
  );
}
