import { css } from 'lit';

export const styles = css`
:host {
	font-family: var(--font-family);
	font-size: var(--font-size);
	color: var(--text-1);
	fill: var(--text-1);
}
.row {
  display: table;
}
.cell {
  display: table-cell;
  vertical-align: middle;
}
.padding-small { 
  padding: 4px 8px; 
}
@media (max-width:600px){
  .padding-small { 
    padding: 2px 4px; 
  }
  .hide-small { 
    display: none!important; 
  }
}
`