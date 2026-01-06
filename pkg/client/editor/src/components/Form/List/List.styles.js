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
.row {
  display: table;
}
.full {
  width: 100%; 
}
.cell {
  display: table-cell;
  vertical-align: middle;
}
.responsive {
  display: grid; 
  width: 100%;
}
.list {
  overflow-x: auto;
  width: 100%;
  border: 1px solid rgba(var(--neutral-1), 0.2);
  border-bottom: none;
  list-style-type: none;
  padding: 0;
  margin: 0;
  box-sizing: border-box;
}
.list li:nth-child(odd) {
  background-color: rgba(var(--functional-yellow),0.1);
}
.list-row {
  display: table;
  width: 100%;
}
.border-bottom {
  border-bottom: 1px solid rgba(var(--neutral-1), 0.2);
}
.edit-cell {
  display: table-cell;
  vertical-align: middle;
  text-align: center;
  width: 45px;
  cursor: pointer;
}
.edit-cell:hover {
  fill: rgb(var(--functional-green));
}
.value-cell {
  display: table-cell;
  vertical-align: middle;
}
.cursor {
  cursor: pointer;
}
.label {
  width: 100%;
  font-weight: bold;
  padding: 5px 8px 2px;
}
.value {
  width: 100%;
  padding: 2px 8px 5px;
}
.delete-cell {
  display: table-cell;
  vertical-align: middle;
  text-align: center;
  width: 45px;
  cursor: pointer;
}
.delete-cell:hover {
  fill: rgb(var(--functional-red));
}
`