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
.cell {
  display: table-cell;
  vertical-align: middle;
}
.full { 
  width: 100%; 
}
.container-row {
  display: table;
  width: 100%;
  border-bottom: 1px solid rgba(var(--neutral-1), 0.2);
}
.total-cell {
  display: table-cell;
  vertical-align: middle;
  padding: 8px 16px;
  text-align: right;
}
.total-label {
  font-size: 13px;
  white-space: nowrap;
  font-weight: bold;
}
.total-value {
  font-size: 13px;
  white-space: nowrap;
  font-weight: bold;
  padding: 4px 6px;
  background-color: rgba(var(--base-3), 1);
  border: 1px solid rgba(var(--neutral-1), 0.2);
  border-radius: 3px;
  margin-left: 8px;
}
@media (max-width:600px){
  .mobile { 
    display: block; 
    width: 100%; 
  }
  .total-cell { 
    padding: 4px 8px;; 
  }
}
`