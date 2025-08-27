import { z } from "zod";

// Schema for a single CO2 data point in the time series
export const PointSchema = z.object({
  T: z.string().refine((val: string) => !isNaN(Date.parse(val)), {
    message: "Invalid date string",
  }),
  V: z.number(),
});

// Schema for the latest CO2 metric
export const MetricLatestSchema = z.object({
  timestamp: z.string().refine((val: string) => !isNaN(Date.parse(val)), {
    message: "Invalid date string",
  }),
  value: z.number(),
});

// Schema for the CO2 series response (array of points)
export const SeriesSchema = z.array(PointSchema);

// Types inferred from schemas
export type Point = z.infer<typeof PointSchema>;
export type MetricLatest = z.infer<typeof MetricLatestSchema>;
export type Series = z.infer<typeof SeriesSchema>;
