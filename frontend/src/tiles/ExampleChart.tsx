import React from "react";
import { VegaLite } from "react-vega";
import { themeConfig } from "../charts/theme";

const data = {
  table: Array.from({ length: 30 }, (_, i) => ({
    x: i,
    y: Math.sin(i / 5) + i / 50,
  })),
};

const spec = {
  $schema: "https://vega.github.io/schema/vega-lite/v5.json",
  height: 140,
  mark: { type: "line" as const, interpolate: "monotone" },
  data: { name: "table" },
  encoding: {
    x: { field: "x", type: "quantitative", axis: { title: null } },
    y: { field: "y", type: "quantitative", axis: { title: null } },
    color: { value: "var(--chart-1)" },
  },
};

export default function ExampleChart() {
  return (
    <VegaLite
      spec={spec as any}
      data={data}
      config={themeConfig()}
      actions={false}
    />
  );
}
