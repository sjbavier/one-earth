import { useQuery } from "@tanstack/react-query";
import { MetricLatestSchema, SeriesSchema } from "../schemas/co2";

const fetchCO2Latest = async () => {
  const response = await fetch("/api/metrics/co2");
  if (!response.ok) {
    throw new Error("Network response was not ok");
  }
  const json = await response.json();
  return MetricLatestSchema.parse(json);
};

const fetchCO2Series = async (days: number) => {
  const response = await fetch("/api/series/co2?days=" + days);
  if (!response.ok) {
    throw new Error("Network response was not ok");
  }
  const json = await response.json();
  return SeriesSchema.parse(json);
};

export function useCO2Latest() {
  return useQuery({
    queryKey: ["co2Latest"],
    queryFn: fetchCO2Latest,
    refetchInterval: 60000, // refetch every 60 seconds
    retry: 3,
  });
}

export function useCO2Series(days: number) {
  return useQuery({
    queryKey: ["co2Series", days],
    queryFn: () => fetchCO2Series(days),
    refetchInterval: 60000, // refetch every 60 seconds
    retry: 3,
  });
}
