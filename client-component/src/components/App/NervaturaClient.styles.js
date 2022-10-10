import { css } from 'lit';

export const styles = css`
  html, :host {
    display: block;
    font-family: var(--font-family);
    font-size: var(--font-size);
    width: 100%;
    height: 100%;
    margin: 0;
    -webkit-font-smoothing: antialiased;
    -moz-osx-font-smoothing: grayscale;
  }
  main {
    height: 100%;
    width: 100%;
    position: absolute;
    left: 0;
    top: var(--menu-top-height);
  }
  /* scrollbar */
  *::-webkit-scrollbar {
    width: 10px;
    height: 5px;
    background-color: transparent;
    visibility: hidden;
  }
  *::-webkit-scrollbar-track {
    background-color: rgba(var(--neutral-1), .05);
    border-radius: 8px;
  }
  *::-webkit-scrollbar-thumb {
    background-color: rgba(var(--neutral-1), .10);
    border-radius: 8px;
  }
  *::-webkit-scrollbar-thumb:active,
  *::-webkit-scrollbar-thumb:hover {
    background-color: rgba(var(--neutral-1), .20)
  }
`