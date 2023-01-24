import { css } from 'lit';

export const styles = css`
@keyframes animatezoom{from{transform:scale(0)} to{transform:scale(1)}}
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
.modal {
  z-index: 10;
  position: fixed;
  left: 0;
  right: 0;
  top: 0;
  height: 100vh;
  overflow: auto;
  background-color: rgba(25, 25, 25, 0.7);
  padding: 30px 5px;
}
.dialog {
  border-radius: 4px;
  border: 0.5px solid rgba(var(--neutral-1), 0.2);
  box-shadow:0 4px 10px 0 rgba(0,0,0, 0.2),0 4px 20px 0 rgba(0,0,0, 0.19);
  background-color: rgba(var(--base-2), 1);
  margin: 0 auto;
  animation: animatezoom 0.6s;
  width: 100%;
  max-width: 400px;
  min-width: 280px;
}
.panel {
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
.section {
  padding: 16px 0; 
}
.section-row {
  display: table;
  width: 100%;
  padding: 0 16px;
}
.padding-small { 
  padding: 4px 8px; 
}
.padding-tiny { 
  padding: 2px 4px; 
}
.buttons {
  background-color: rgb(var(--base-1));
  border-top: 0.5px solid rgba(var(--neutral-1), 0.2);
}
.table-row {
  display: table-row;
}
@media (max-width:600px){
  .section-row { 
    padding: 0 8px 8px; 
  }
  .padding-small { 
    padding: 2px 4px; 
  }
  .padding-tiny { 
    padding: 1px 2px; 
  }
  .mobile{ 
    display: block; 
    width: 100%; 
  }
}
@media only screen and (min-width: 601px){
  .dialog {
    min-width: 400px;
  }
}
`