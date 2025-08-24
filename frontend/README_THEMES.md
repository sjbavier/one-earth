
# One Earth Themes

**What you get**
- `styles/theme.css` — CSS variables (light=Reuters-esque, dark=Apple HIG) + small UI primitives
- `charts/theme.ts` — Vega-Lite theme configs that read those CSS vars
- `lib/theme.ts` — helpers to set/get theme

## Use

1) Import CSS in your app entry:
```ts
import './styles/theme.css'
```

2) Toggle theme:
```ts
import { setMode } from './lib/theme'
setMode('light')  // or 'dark' or 'system'
```

3) Vega-Lite config:
```tsx
import { VegaLite } from 'react-vega'
import { themeConfig } from './charts/theme'
<VegaLite spec={spec} config={themeConfig()} actions={false} />
```

4) Basic UI in JSX:
```tsx
<div className="card">
  <div className="chip">Last updated: just now</div>
  <h2 className="stat">419.2<span className="muted"> ppm</span></h2>
</div>
```
