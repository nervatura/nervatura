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
.table-wrap {
  display: block;
  max-width: 100%;
  overflow-x: auto;
  overflow-y: hidden;
}
.sort {
  cursor: pointer;
  cursor: hand;
}
.sort:after {
  padding-left: 1em;
  font-size: 0.5em;
}
.sort-none:after {
  padding-left: 0.5em;
  font-size: 1em;
}
.sort-asc:after {
  content: '▲';
  font-size: 14px;
  color: rgb(var(--functional-yellow));
}
.sort-desc:after {
  content: '▼';
  font-size: 14px;
  color: rgb(var(--functional-yellow));
}
.sort-order {
  margin-left: 0.5em;
}
.number-cell {
  text-align: right;
}
.link-cell {
  color: rgb(var(--functional-blue));
  cursor: pointer;
}
.link-cell:hover {
  text-decoration: underline;
}
.ui-table {
  overflow-x: auto;
  border-collapse: collapse;
  border-spacing: 0;
  width: 100%;
  font-size: 14px;
  border: 1px solid rgba(var(--neutral-1), 0.2);
  box-sizing: border-box;
}
.ui-table thead th {
  border-bottom: 1px solid rgba(var(--neutral-1),0.2);
}
.ui-table tbody tr:nth-child(odd) {
  background-color: rgba(var(--functional-yellow),0.1);
}
.ui-table tbody th, .ui-table tbody td {
  vertical-align: top;
}
.ui-table tbody tr:hover {
  color: rgb(var(--functional-yellow));
}
.ui-table th:hover {
  color: rgb(var(--functional-yellow));
}
.cell-label { 
	display: none;
  font-weight: bold;
  font-size: 12px;
}
/* Mobile first styles: Begin with the stacked presentation at narrow widths */ 
@media only all {
	/* Hide the table headers */ 
	.ui-table thead td, 
	.ui-table thead th {
		display: none;
    font-size: 13px;
    background-color: rgb(var(--accent-1));
    color: rgb(255, 255, 255);
    fill: rgb(255, 255, 255);
    white-space: nowrap;
	}
	/* Show the table cells as a block level element */ 
	.ui-table td,
	.ui-table th { 
		text-align: left;
		display: block;
    padding: 8px;
	}
	/* Add a fair amount of top margin to visually separate each row when stacked */  
	.ui-table tbody th {
		margin-top: 3em;
	}
	/* Make the label elements a percentage width */ 
	.cell-label { 
		padding: 4px 6px;
		min-width: 30%; 
		display: inline-block;
		margin: -.4em 1em -.1em -.4em;
	}
}
/* Breakpoint to show as a standard table at 560px (35em x 16px) or wider */ 
@media ( min-width: 35em ) {
	/* Show the table header rows */ 
	.ui-table td,
	.ui-table th,
	.ui-table tbody th,
	.ui-table tbody td,
	.ui-table thead td,
	.ui-table thead th {
		display: table-cell;
    margin: 0;
	}
	/* Hide the labels in each cell */ 
  .cell-label { 
		display: none;
	}
}
/* Hack to make IE9 and WP7.5 treat cells like block level elements, scoped to ui-responsive class */ 
/* Applied in a max-width media query up to the table layout breakpoint so we don't need to negate this*/ 
@media ( max-width: 35em ) {
	.ui-table td,
	.ui-table th {
		width: 100%;
		-webkit-box-sizing: border-box;
		-moz-box-sizing: border-box;
		box-sizing: border-box;
		float: left;
		clear: left;
    font-size: 13px;
    padding: 4px 6px;
    line-height: 11.5px;
	}
  .number-cell {
    text-align: left;
  }
}
.cursor-pointer {
  cursor: pointer;
}
.cursor-disabled {
  cursor: not-allowed;
}
.middle {
  vertical-align: middle;
}
`