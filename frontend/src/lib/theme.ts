
export type Mode = 'light' | 'dark' | 'system';

export function getMode(): Mode {
  const t = document.documentElement.getAttribute('data-theme');
  return (t === 'light' || t === 'dark') ? t : 'system';
}

export function isDark(): boolean {
  const t = getMode();
  if (t === 'light') return false;
  if (t === 'dark') return true;
  return window.matchMedia && window.matchMedia('(prefers-color-scheme: dark)').matches;
}

export function setMode(mode: Mode) {
  if (mode === 'system') {
    document.documentElement.removeAttribute('data-theme');
  } else {
    document.documentElement.setAttribute('data-theme', mode);
  }
  document.documentElement.style.colorScheme = isDark() ? 'dark' : 'light';
}

export function watchSystemTheme(callback: (dark: boolean) => void) {
  const mq = window.matchMedia('(prefers-color-scheme: dark)');
  const handler = () => callback(mq.matches);
  mq.addEventListener('change', handler);
  return () => mq.removeEventListener('change', handler);
}
