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
  padding: 8px;
  border-bottom: 1px solid rgba(var(--neutral-1), 0.2);
}
.paginator-cell {
  display: table-cell;
  vertical-align: middle;
  padding-left: 8px;
  float: right;
}
.padding-small { 
  padding: 4px 8px; 
}
@media (max-width:600px){
  .mobile, .paginator-cell { 
    display: block; 
    width: 100%; 
  }
  .padding-small { 
    padding: 2px 4px; 
  }
}
`