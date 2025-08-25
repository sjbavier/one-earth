export const reutersLight = {
  background: "transparent",
  padding: 8,
  axis: {
    labelColor: "var(--ink-2)",
    labelFontSize: 12,
    titleColor: "var(--ink)",
    titleFontSize: 12,
    grid: true,
    gridColor: "var(--grid)",
    domainColor: "var(--border)",
    tickColor: "var(--border)",
  },
  legend: { labelColor: "var(--ink-2)", titleColor: "var(--ink)" },
  view: { stroke: "var(--border)" },
  line: { strokeWidth: 2 },
  area: { opacity: 0.7 },
  bar: { cornerRadiusTopLeft: 2, cornerRadiusTopRight: 2 },
  range: {
    category: [
      "var(--chart-1)",
      "var(--chart-2)",
      "var(--chart-3)",
      "var(--chart-4)",
      "var(--chart-5)",
      "var(--chart-6)",
    ],
  },
};

export const higDark = {
  background: "transparent",
  padding: 8,
  axis: {
    labelColor: "var(--ink-2)",
    labelFontSize: 12,
    titleColor: "#fff",
    titleFontSize: 12,
    grid: true,
    gridColor: "var(--grid)",
    domainColor: "var(--border)",
    tickColor: "var(--border)",
  },
  legend: { labelColor: "var(--ink-2)", titleColor: "#fff" },
  view: { stroke: "var(--border)" },
  line: { strokeWidth: 2 },
  area: { opacity: 0.6 },
  bar: { cornerRadiusTopLeft: 2, cornerRadiusTopRight: 2 },
  range: {
    category: [
      "var(--chart-1)",
      "var(--chart-2)",
      "var(--chart-3)",
      "var(--chart-4)",
      "var(--chart-5)",
      "var(--chart-6)",
    ],
  },
};

export const themeConfig = () =>
  document.documentElement.dataset.theme === "dark" ? higDark : reutersLight;
