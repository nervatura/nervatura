import { css } from 'lit';

export const styles = css`
:host {
  font-family: var(--font-family);
  font-size: var(--font-size);
  color: var(--text-1);
  fill: var(--text-1);
}
div {
  box-sizing: border-box;
}
.panel {
  width: 100%;
  background-color: rgba(var(--base-2), 1);
  border: 0.5px solid rgba(var(--neutral-1), 0.2);
  border-top: none;
  border-bottom: none;
}
.row {
  display: table;
}
.full { 
  width: 100%; 
}
.title-cell {
  display: table-cell;
  font-weight: bold;
  padding: 8px 16px;
  border: 0.5px solid rgba(var(--neutral-1), 0.2);
  background-color: rgba(var(--neutral-1), 0.1);
}
.title-pre {
  padding-right: 6px;
  opacity: 0.5;
}
@media (max-width:600px){
  .title-cell {
    padding: 4px 8px;
  }
}
`