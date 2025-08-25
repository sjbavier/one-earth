import React from "react";
import { createRoot } from "react-dom/client";
import App from "./App";
import "./styles/theme.css";
import { initMode } from "./lib/theme";

initMode();

const root = createRoot(document.getElementById("root")!);
root.render(<App />);
