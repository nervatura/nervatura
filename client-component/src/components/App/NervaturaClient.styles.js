import { css } from 'lit';

export const styles = css`
html, :host {
  display: block;
  font-family: var(--font-family);
  font-size: var(--font-size);
  color: var(--text-1);
  fill: var(--text-1);
  width: 100%;
  height: 100%;
  margin: 0;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
}
.main {
  background-color: rgb(var(--base-0));
  height: 100%;
  width: 100%;
  position: absolute;
  left: 0;
  top: var(--menu-top-height);
}
.client-menubar {
  z-index: 5;
  position: fixed;
  width: 100%;
  top: 0;
}
*::-webkit-scrollbar {
  width: 10px;
  height: 5px;
  background-color: transparent;
  visibility: hidden;
}
*::-webkit-scrollbar-track {
  background-color: rgba(var(--accent-1), .05);
  border-radius: 8px;
}
*::-webkit-scrollbar-thumb {
  background-color: rgba(var(--accent-1), .60);
  border-radius: 8px;
}
*::-webkit-scrollbar-thumb:active,
*::-webkit-scrollbar-thumb:hover {
  background-color: rgba(var(--accent-1), .20)
}
`