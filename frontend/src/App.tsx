import React, { useEffect, useState } from "react";
import ThemeToggle from "./components/ThemeToggle";
import ExampleChart from "./tiles/ExampleChart";
import { CO2Tile } from "./tiles/CO2Tile";

const guessApi = () => {
  const { origin } = window.location;
  if (origin.endsWith(":8080")) return origin.replace(":8080", ":8081");
  return origin;
};

export default function App() {
  const [msg, setMsg] = useState("…loading");
  const API = guessApi();
  const siteUrl =
    (import.meta as any).env?.VITE_PUBLIC_SITE_URL ?? "https://one-earth.info";

  useEffect(() => {
    fetch(`${API}/api/hello`)
      .then((r) => r.json())
      .then((d) => setMsg(d.message))
      .catch(() => setMsg("API unreachable"));
  }, [API]);

  return (
    <div
      style={{
        fontFamily: "system-ui, sans-serif",
        padding: 24,
        lineHeight: 1.6,
      }}
    >
      <h1>One Earth</h1>
      <ThemeToggle />
      <ExampleChart />
      <CO2Tile />

      <p>
        API says: <strong>{msg}</strong>
      </p>
      <p>
        Health check:{" "}
        <a href={`${API}/health`} target="_blank" rel="noreferrer">
          {API}/health
        </a>
      </p>
      <p>
        Public site URL (baked at build): <code>{siteUrl}</code>
      </p>
    </div>
  );
}
