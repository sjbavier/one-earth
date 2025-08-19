# Climate & Environment Data Sources (MVP) 

Each source below is reputable, offers programmatic access, and has a predictable refresh cadence. For every tile on the site, display: **source**, **method**, **update cadence**, **license**, and a visible **“Last updated”** timestamp.  
_Last updated: 2025-08-17_

---

## 1) Atmospheric CO₂ (ppm) — Mauna Loa

**Primary**: NOAA Global Monitoring Laboratory (GML) — daily, weekly, monthly.  
- **Programmatic**: CSV/TXT endpoints for daily/weekly/monthly means.  
- **Cadence**: Daily (preliminary), weekly & monthly official.  
- **Docs/Data**:  
  - https://gml.noaa.gov/ccgg/trends/  
  - https://gml.noaa.gov/ccgg/trends/data.html

**Alternative (backup/compare)**: Scripps “Keeling Curve” (daily)  
- https://keelingcurve.ucsd.edu/  
- https://scrippsco2.ucsd.edu/data/atmospheric_co2/

---

## 2) Global Temperature Anomaly (°C vs baseline)

**Primary**: NASA GISTEMP v4  
- **Programmatic**: Tables/files from GISTEMP site; mirrored via NOAA PSL.  
- **Cadence**: ~Monthly (mid-month) once prerequisite datasets update.  
- **Docs/Data**:  
  - https://data.giss.nasa.gov/gistemp/  
  - https://data.giss.nasa.gov/gistemp/updates_v4/  
  - https://psl.noaa.gov/data/gridded/data.gistemp.html

---

## 3) Global Mean Sea Level (mm since 1993)

**Primary**: NASA Sea Level Change (altimetry time series); corroboration from AVISO/CNES.  
- **Programmatic**: Time-series downloads via NASA portal; AVISO HTTPS/FTP for GMSL.  
- **Cadence**: Monthly (satellite altimetry).  
- **Docs/Data**:  
  - https://sealevel.nasa.gov/  
  - https://sealevel.nasa.gov/understanding-sea-level/key-indicators/global-mean-sea-level/  
  - https://www.aviso.altimetry.fr/en/data/products/ocean-indicators-products/mean-sea-level/data-acces.html

---

## 4) Arctic & Antarctic Sea-Ice Extent (million km²)

**Primary**: NSIDC Sea Ice Index (daily + monthly)  
- **Programmatic**: ASCII/CSV downloads via “Get Source Data”; map services available.  
- **Cadence**: Daily & monthly.  
- **Docs/Data**:  
  - https://nsidc.org/data/seaice_index

---

## 5) Air Quality Now (PM₂.₅ / AQI by city)

**Global**: OpenAQ  
- **Programmatic**: REST API (OpenAPI), bulk via S3.  
- **Cadence**: Near-real-time (varies by network).  
- **Docs**: https://docs.openaq.org/ • https://api.openaq.org/

**U.S.**: EPA AirNow  
- **Programmatic**: REST API (key required).  
- **Cadence**: Hourly/near-real-time; daily forecasts.  
- **Docs**: https://www.airnowapi.org/ • https://docs.airnowapi.org/webservices

---

## 6) Active Wildfires (near-real-time points)

**Primary**: NASA LANCE/FIRMS (MODIS & VIIRS detections)  
- **Programmatic**: CSV/KML/GeoJSON via API; WMS/WFS web services; “Ultra Real-Time” (URT) for parts of N. America.  
- **Cadence**: NRT (minutes–hours), URT (sub-hour) where available.  
- **Docs/API**: https://firms.modaps.eosdis.nasa.gov/api/ • https://firms.modaps.eosdis.nasa.gov/web-services/

---

## 7) Electric Grid Carbon Intensity (optional but compelling)

**Global (commercial)**: Electricity Maps (real-time + forecast)  
- **Programmatic**: Paid API (current/forecast intensity & mix).  
- **Docs**: https://portal.electricitymaps.com/docs

**Great Britain (free)**: Carbon Intensity API (National Energy System Operator)  
- **Programmatic**: Public REST, regional & national series.  
- **Docs**: https://carbonintensity.org.uk/ • https://api.carbonintensity.org.uk/

**United States (free/open)**: U.S. Energy Information Administration (EIA) API v2  
- **Programmatic**: REST with API key; RTO/BA load, generation, interchanges; combine with emissions factors.  
- **Docs**: https://www.eia.gov/opendata/ • https://www.eia.gov/opendata/documentation.php

---

## 8) Deforestation Alerts (near-real-time)

**Primary**: Global Forest Watch (GLAD-L, GLAD-S2, RADD) + Integrated Deforestation Alerts  
- **Programmatic**: GFW Data API (auth); query by AOI/admin; returns JSON/GeoJSON.  
- **Cadence**: Days–weeks depending on product/cloud cover.  
- **Docs/Datasets**:  
  - https://data-api.globalforestwatch.org/  
  - GLAD alerts overview: https://www.globalforestwatch.org/blog/data-and-tools/glad-deforestation-alerts/  
  - RADD: https://data.globalforestwatch.org/datasets/gfw::deforestation-alerts-radd/about  
  - Integrated alerts explainer: https://www.globalforestwatch.org/blog/data-and-tools/integrated-deforestation-alerts/

---

## 9) ENSO (Niño 3.4 / ONI)

**Primary**: NOAA Climate Prediction Center (CPC); Niño regions via NOAA PSL  
- **Programmatic**: Plain-text tables for weekly SSTs; downloadable monthly series; ONI.  
- **Cadence**: Weekly (SST regions), monthly (ONI/diagnostics).  
- **Docs/Data**:  
  - ONI v5: https://origin.cpc.ncep.noaa.gov/products/analysis_monitoring/ensostuff/ONI_v5.php  
  - ENSO discussion: https://www.cpc.ncep.noaa.gov/products/analysis_monitoring/enso_advisory/ensodisc.shtml  
  - Weekly SSTs file: https://www.cpc.ncep.noaa.gov/data/indices/wksst9120.for  
  - Niño 3.4 (PSL): https://psl.noaa.gov/data/timeseries/month/Nino34/

> Note: The **Keeling Curve** belongs with CO₂, not ENSO.

---

## 10) Annual Emissions (context; not “live”)

**Primary**: Global Carbon Budget (GCB) + Our World in Data (harmonized series)  
- **Programmatic**: GCB spreadsheets (annual); OWID CSV/JSON repo with codebook.  
- **Cadence**: Yearly (typically Nov/Dec).  
- **Docs/Data**:  
  - GCB hub: https://globalcarbonbudget.org/gcb-2024/  
  - Latest data portal: https://globalcarbonbudgetdata.org/latest-data.html  
  - OWID GitHub: https://github.com/owid/co2-data  
  - OWID explorer: https://ourworldindata.org/explorers/co2

---

## Data Governance & Licensing (safe to publish)

- **NOAA & NASA**: U.S. Federal works → public domain; attribution recommended.  
  - NOAA open data/licensing: https://www.noaa.gov/disclaimer  
  - NASA Earthdata policy: https://earthdata.nasa.gov/earthdata/citing-our-data

- **Copernicus (ECMWF CDS/C3S/CAMS)**: Copernicus products are under **CC BY 4.0** (attribution required).  
  - Licence info: https://cds.climate.copernicus.eu/licence  
  - ECMWF/Forum announcements & guidance: https://www.ecmwf.int/en/about/licensing

> Include an explicit license line on each tile (e.g., “NOAA public domain; attribution appreciated” or “Copernicus CC BY 4.0”).

---

## Implementation Notes (for “live” feel)

- **Timestamps**: Persist `fetched_at` and `source_last_modified` (from HTTP headers) and render “Updated <date>”.  
- **Polling** (respect rate limits & ETags):  
  - CO₂ daily; Sea Ice daily; FIRMS 5–15 min (URT where available); OpenAQ hourly; GISTEMP/Sea Level monthly; ENSO weekly/monthly; GCB yearly.  
- **Transparency**: Link to the raw CSV/JSON powering each viz; include a short method blurb (baseline, anomaly definition, instruments).  
- **Redundancy**: Keep a read-through cache and raw snapshots in object storage; use append-only storage for observations.

---

## Nice-to-have (backups & general catalogs)

- NOAA NCEI Climate Data Online (CDO): https://www.ncei.noaa.gov/cdo-web/  
- Copernicus Climate Data Store (CDS): https://cds.climate.copernicus.eu/  
