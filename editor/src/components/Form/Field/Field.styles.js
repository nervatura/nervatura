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
textarea {
  font-family: var(--font-family);
  font-size: var(--font-size);
  border-radius: 3px;
  overflow: auto;
  padding: 8px;
  color: var(--text-1);
  background-color: rgba(var(--base-4), 1);
  border: 1px solid rgba(var(--neutral-1), 0.2);
  box-sizing: border-box;
}
textarea:focus, textarea:hover:not(:disabled) {
  color: rgb(var(--functional-green))!important;
}
textarea:disabled {
  opacity: 0.5;
}
textarea::placeholder, textarea::-ms-input-placeholder {
  opacity: 0.5;
}
.link {
  font-family: var(--font-family);
  font-size: var(--font-size);
  width: 100%;
  border-radius: 3px;
  padding: 6px 8px 5px;
  vertical-align: middle;
  min-height: 35px;
  color: var(--text-1);
  background-color: rgba(var(--base-4), 1);
  border: 1px solid rgba(var(--neutral-1), 0.2);
}
.link-text {
  font-size: 14px;
  cursor: pointer;
  color: rgb(var(--functional-blue));
}
.link-text:hover {
  text-decoration: underline;
}
.search-col {
  width: 38px;
}
.times-col {
  width: 34px;
}
.toggle-on {
  fill: rgb(var(--functional-green))!important;
}
.toggle {
  font-family: var(--font-family);
  font-size: var(--font-size);
  text-align: center;
  cursor: pointer;
  width: 100%;
  border-radius: 3px;
  fill: var(--text-1);
  background-color: rgba(var(--base-4), 1);
  border: 1px solid rgba(var(--neutral-1), 0.2);
}
.toggle:hover:not(:disabled) {
  fill: rgb(var(--functional-green))!important;
}
.toggle-off {
  opacity: 0.5;
}
.toggle-disabled {
  opacity: 0.5;
}
`