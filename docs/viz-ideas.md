# One Earth — Visualization Ideas
_Last updated: 2025-08-20_

A menu of visualization patterns aligned to our data sources and refresh cadences. Use these as building blocks for MVP tiles and future stories.

---

## By Metric

### Atmospheric CO₂ (NOAA GML — daily/weekly/monthly)
- **Live tile:** current ppm + 7–30 day sparkline.
- **Seasonal cycle:** monthly mean with shaded seasonal band + anomaly overlay.
- **Climate spiral:** polar plot of monthly anomalies (signature visual).
- **Small multiples:** year-over-year mini charts to show seasonality.

### Global Temperature Anomaly (GISTEMP — monthly)
- **Warming stripes:** vertical bands by month (baseline vs anomaly).
- **Line/area:** anomaly with optional uncertainty band.
- **Ridgelines by decade:** distribution of monthly anomalies.
- **Calendar heatmap:** months × years.

### Global Mean Sea Level (NASA — monthly)
- **Long-run line:** cumulative mm since 1993 + rolling trend (mm/yr).
- **Milestone meter:** headline cumulative change visual.

### Sea-Ice Extent (NSIDC — daily/monthly)
- **Seasonal ribbon:** current year vs climatological min–max + median.
- **Small multiples:** Arctic vs Antarctic; or one chart per month across years.
- **Threshold markers:** dates crossing key extents.

### Air Quality (OpenAQ/AirNow — hourly/near-RT)
- **City tile:** AQI dial or bullet + pollutant chips (PM2.5, O₃).
- **Map:** hexbin/heatmap with station dots & tooltips.
- **Diurnal curve:** typical 24h profile; **calendar heatmap** (30–90 days).
- **Rankings:** bar chart of best/worst locations (filterable).

### Wildfires (NASA FIRMS — NRT/URT)
- **Live map:** clustered points, size/opacity by detection age; time slider.
- **Timeline:** detections per hour/day (stack by confidence/sensor).
- **Comparative bars:** this week vs last week, region/AOI selectable.

### Grid Carbon Intensity (GB CI, EIA, Electricity Maps — 30-min/minutes)
- **Now tile:** gCO₂/kWh + mix donut (coal/gas/renewables) + mini-forecast.
- **Regional map:** choropleth by zone/BA; arrows for net imports (stretch).
- **Sankey:** generation → consumption (where data allows).

### Deforestation Alerts (GFW GLAD/RADD — days–weeks)
- **Alert density map:** grid/hex with time scrubber.
- **Bump chart:** top regions by new alerts this month.
- **AOI dashboard:** alerts over time + selectable polygon/map extent.

### ENSO (CPC ONI & Niño 3.4 — weekly/monthly)
- **Gauge:** current Niño 3.4 category (La Niña…El Niño).
- **Strip/phase plot:** ONI vs thresholds over time.
- **Seasonality swirl:** monthly phase around the year.
- **(Stretch)** correlation minis: ONI vs anomaly impacts (educational).

### Annual Emissions (GCB/OWID — yearly)
- **Stacked area:** by sector or fuel.
- **Slope chart:** country trajectories.
- **Waterfall:** YoY change contributors (coal, gas, land-use).
- **Per-capita map:** snapshot choropleth + (stretch) time slider.

---

## Cross‑Cutting Patterns
- **Tile anatomy:** Value + unit → sparkline/trend arrow → “Last updated” chip → Source & License → Download data.
- **Compare mode:** Region dropdown + small multiples (e.g., Arctic vs Antarctic; city lists).
- **Brushing/zoom:** 7/30/90‑day ranges for high‑cadence series.
- **Uncertainty & baselines:** shaded climatology bands; baseline notes in tooltips.
- **Downloadables:** link the exact CSV/JSON powering each viz.

## Maps — Recurring Motifs
- **Hexbin/contour** for dense points (AQI, fires).
- **Clustered markers** with “age” styling (opacity/size by hours).
- **Choropleth** for zones/regions (grid carbon, per‑capita emissions).
- **Time slider** wherever data is near‑real‑time.

## “Wow” Pieces (use sparingly)
- **Climate spiral** (temp anomalies).
- **Animated seasonal ribbon** (sea‑ice).
- **Story scrollytell:** ENSO → temp anomalies → fires linkage.

## Accessibility & Style
- Color‑blind‑safe palettes; avoid red/green alone for critical cues.
- Keyboard navigable charts; ARIA labels and descriptive tooltips.
- Dark‑mode friendly; consistent legend & axis patterns.

---

## Next Specs to Scaffold (suggested)
- **CO₂ tile + long‑run:** Vega‑Lite line + seasonal band; 30–60‑point sparkline.
- **Sea‑ice seasonal ribbon:** current year vs climatology band.
- **Air quality city tile:** dial/bullet + pollutant chips; 24h diurnal curve.
- **FIRMS live map:** clustered points with detection‑age styling + 24h slider.

> Keep each tile self‑contained: UI, data hook, and (if needed) its Vega‑Lite spec. Validate responses (Zod) and surface “data unavailable” gracefully.
