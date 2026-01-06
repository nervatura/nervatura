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
.half { 
  float:left; 
  width:100% 
}
.s12 { 
  float:left; 
  width:99.99999%; 
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
.bold { 
  font-weight: bold; 
}
.padding-small { 
  padding: 4px 8px; 
}
.padding-tiny { 
  padding: 2px 4px; 
}
.container-small { 
  padding: 0px 8px; 
}
.container-row {
  display: table;
  width: 100%;
  padding: 8px;
  border-bottom: 1px solid rgba(var(--neutral-1), 0.2);
}
.label-row {
  color: rgb(var(--functional-blue));
}
.report-field {
  cursor: pointer;
}
.toggle-on {
  fill: rgb(var(--functional-green))!important;
  color: rgb(var(--functional-green))!important;
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
  color: rgb(var(--functional-green))!important;
}
.toggle-off {
  opacity: 0.5;
}
.toggle-disabled {
  opacity: 0.5;
}
.info {
  background-color: rgba(var(--functional-blue),0.2)!important;
}
.leftbar{
  border-left: 8px solid rgba(var(--functional-yellow),0.8)!important;
  padding: 8px;
  font-style: italic;
  font-size: 13px;
}
.center { 
  text-align: center; 
}
.align-right { 
  text-align: right; 
}
textarea {
  font-family: var(--font-family);
  font-size: 12px;
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
.field-cell {
  width: 150px;
}
.fieldvalue-delete {
  cursor: pointer;
}

.fieldvalue-delete:hover {
  fill: rgb(var(--functional-red));
}
@media (max-width:600px){
  .container-row {
    padding: 4px;
  }
  .container-small { 
    padding: 0px 4px; 
  }
  .padding-small { 
    padding: 2px 4px; 
  }
  .padding-tiny { 
    padding: 1px 2px; 
  }
  .hide-small { 
    display: none!important; 
  }
}
@media (min-width:601px) {
  .half { 
    width:49.99999% 
  }
  .m3 { 
    float:left; 
    width:24.99999%; 
  }
  .m4 { 
    float:left; 
    width: 33.33333%; 
  }
  .m6 { 
    float:left; 
    width: 49.99999%; 
  }
}
@media (max-width:992px) and (min-width:601px){
  .hide-medium { 
    display: none!important; 
  }
}
@media (min-width:993px) {
  .hide-large { 
    display: none!important; 
  }
  .l3 { 
    float:left; 
    width: 24.99999%; 
  }
  .l4 { 
    float:left; 
    width: 33.33333%; 
  }
  .l6 {
    float:left; 
    width: 49.99999%; 
  }
}
`