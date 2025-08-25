import React from "react";
import { getMode, setMode, Mode } from "../lib/theme";

export default function ThemeToggle() {
  const [mode, setLocal] = React.useState<Mode>(getMode());
  const set = (m: Mode) => {
    setMode(m);
    setLocal(m);
  };

  return (
    <div className="chip" role="group" aria-label="Theme">
      <button onClick={() => set("light")} aria-pressed={mode === "light"}>
        Light
      </button>
      <button onClick={() => set("dark")} aria-pressed={mode === "dark"}>
        Dark
      </button>
      <button onClick={() => set("system")} aria-pressed={mode === "system"}>
        System
      </button>
    </div>
  );
}
