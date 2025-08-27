import type { TopLevelSpec } from "vega-lite";

export const co2Spec: TopLevelSpec = {
  width: 300,
  height: 100,
  padding: 5,
  data: { name: "values" },
  mark: {
    type: "line",
    interpolate: "monotone",
    color: "#4B5563", // Tailwind gray-600
  },
  encoding: {
    x: {
      field: "T",
      type: "temporal",
      axis: null,
    },
    y: {
      field: "V",
      type: "quantitative",
      axis: null,
    },
  },
  config: {
    axis: {
      grid: false,
      domain: false,
      ticks: false,
      labels: false,
    },
  },
};
