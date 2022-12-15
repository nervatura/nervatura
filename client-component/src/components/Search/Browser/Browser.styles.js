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
.filter-panel {
  width: 100%;
  padding: 16px;
  border: 0.5px solid rgba(var(--neutral-1), 0.2);
  border-top: none;
  border-radius: 0px;
}
.panel-container { 
  padding: 16px; 
}
.align-right { 
  text-align: right; 
}
.dropdown-box {
  position: relative;
  display: inline-block;
}
.section-small-top { 
  padding-top: 8px; 
}
@keyframes opac{
  from{opacity:0} to{opacity:1}
}
.dropdown-content{
  background-color: rgba(var(--base-2), 1);
  border: 0.5px solid rgba(var(--neutral-1), 0.2);
  border-top: none;
  animation: opac 0.8s;
  box-shadow:0 2px 3px 0 rgba(0,0,0,0.16),0 2px 10px 0 rgba(0,0,0,0.12);
  position: absolute;
  min-width: 160px;
  margin: 0;
  padding: 4px 8px;
  z-index: 5;
}
.drop-label {
  font-size: 14px;
  padding: 2px 0px;
  font-weight: bold;
  vertical-align: middle;
  white-space: nowrap;
  width: 100%;
  cursor: pointer;
  -webkit-touch-callout:none;
  -webkit-user-select:none;
  -khtml-user-select:none;
  -moz-user-select:none;
  -ms-user-select:none;
  user-select:none;
}
.drop-label form-label:hover {
  color: rgb(var(--functional-yellow));
  fill: rgb(var(--functional-yellow));
}
.active {
  color: rgb(var(--functional-yellow))!important;
  fill: rgb(var(--functional-yellow))!important;
}
.col-box {
  display: table;
  width: 100%;
  border: 0.5px solid rgba(var(--neutral-1), 0.2);
  border-radius: 4px;
  white-space: nowrap;
  padding: 6px;
  margin-top: 2px;
}
.base-col form-label{
  font-size: 10px;
}
.select-col form-label{
  color: rgb(var(--functional-green));
  fill: rgb(var(--functional-green));
  font-weight: bold;
}
.edit-col form-label:hover {
  font-weight: normal;
}
.edit-col form-label:hover {
  color: rgb(var(--functional-green));
  fill: rgb(var(--functional-green));
}
.col-cell {
  padding: 2px 4px;
  float: left;
  cursor: pointer;
}
.border {
  border: 0.5px solid rgba(var(--neutral-1), 0.2);
}
.result-title {
  background-color: rgba(var(--accent-1),0.1);
  font-weight: bold;
  padding: 8px 16px;
}
.result-title-plus {
  background-color: rgba(var(--accent-1),0.1);
  vertical-align: middle;
  padding: 0px 16px;
  width: 40px;
  text-align: right;
}
@media (max-width:600px){
  .filter-panel {
    padding: 8px;
  }
  .panel-container { 
    padding: 16px 8px; 
  }
  .col-cell { 
    padding: 1px 2px; 
  }
  .mobile{ 
    display: block; 
    width: 100%; 
  }
}
`