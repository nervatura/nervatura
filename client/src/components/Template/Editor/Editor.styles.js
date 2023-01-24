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
.row::before {
  content: "";
  display: table;
  clear: both;
}
.row::after {
  content: "";
  display: table;
  clear: both;
}
.cell {
  display: table-cell;
  vertical-align: middle;
}
.full { 
  width: 100%; 
}
.panel {
  background-color: rgba(var(--base-2), 1);
  border: 0.5px solid rgba(var(--neutral-1), 0.2);
  box-shadow:0 2px 5px 0 rgba(0,0,0,0.16),0 2px 10px 0 rgba(0,0,0,0.12);
  border-radius: 4px;
  max-width: 800px;
}
.panel-title {
  display: table;
  border-radius: 4px;
  font-weight: bold;
  padding: 8px 16px;
  width: 100%;
  border-bottom: 0.5px solid rgba(var(--neutral-1), 0.2);
  background-color: rgb(var(--accent-1));
}
.title-cell {
  color: rgba(var(--accent-1c), 0.85);
  fill: rgba(var(--accent-1c), 0.85);
}
.section-container {
  padding: 16px;
  display: table;
  width: 100%;
}
.section-container-small {
  display: table;
  width: 100%;
  padding: 8px;
}
.border {
  border: 0.5px solid rgba(var(--neutral-1), 0.2);
}
.padding-small { 
  padding: 4px 8px; 
}
.third{ 
  width: 33.33333%;
  vertical-align: top;
}
.mapBox {
  width: 100%;
  text-align: center;
  padding: 10px 5px;
  box-sizing: border-box;
  margin-top: 5px;
	margin-bottom: 5px;
  border: 0.5px solid rgba(var(--neutral-1), 0.2);
  background-color: rgba(var(--neutral-1), 0.1);
}
.reportMap {
	border: 1px solid #666; 
	background-color: white;
	width: 110px; 
	height: 165px;
}
.separator {
  height: 10px;
  width: 100%;
  border: none;
  margin: 0;
}
.report-title {
  width: 100%;
  border: 1.5px solid rgb(var(--functional-yellow));
  margin-top: -1px;
}
.template-row {
  display: table;
  width: 100%;
  border-left: 0.5px solid rgba(var(--neutral-1), 0.2);
  border-right: 0.5px solid rgba(var(--neutral-1), 0.2);
}
.report-title-label {
  color: rgb(var(--functional-yellow));
  fill: rgb(var(--functional-yellow));
}
.align-right { 
  text-align: right; 
}
.close-icon {
  fill: rgba(var(--accent-1c), 0.85);
  cursor: pointer;
}
.close-icon form-icon:hover {
  fill: rgb(var(--functional-red));
}
textarea {
  font-family: var(--font-family);
  font-size: 12px;
  width: 100%;
  border-radius: 3px;
  overflow: auto;
  padding: 8px;
  margin-top: 8px;
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
.meta-title-row {
  display: table;
  border-top: 0.5px solid rgba(var(--neutral-1), 0.2);
  border-bottom: 0.5px solid rgba(var(--neutral-1), 0.2);
  margin: auto;
}
.meta-title-cell {
  display: table-cell;
  vertical-align: top;
  padding: 8px;
  font-size: 12px;
}
.bold {
  font-weight: bold;
}
.meta-title-sources {
  display: table;
  width: 100%;
  padding: 16px 0 0px;
}
.meta-sources {
  display: table;
  width: 100%;
  padding: 8px 0;
}
.meta-sources-name {
  border-top: 0.5px solid rgba(var(--neutral-1), 0.2);
  border-bottom: 0.5px solid rgba(var(--neutral-1), 0.2);
  font-weight: bold;
  font-style: italic;
}
.meta-sources-cell {
  display: table-cell;
  vertical-align: top;
  font-size: 10px;
}
@media (max-width:600px){
  .section-container { 
    padding: 8px; 
  }
  .section-container-small{
    padding: 4px; 
  }
  .padding-small { 
    padding: 2px 4px; 
  }
  .third{ 
    width: 100%;
    display: block;
  }
  .meta-title-row{ 
    display:block; 
    width:100%;
  }
  .meta-title-cell {
    display:block; 
    width:100%;
    padding: 4px;
  }
}
`