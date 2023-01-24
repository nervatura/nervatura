import { css } from 'lit';

export const styles = css`
:host {
  font-family: var(--font-family);
  font-size: var(--font-size);
}
div {
  box-sizing: border-box;
}
.shadow {
  box-shadow: 0 2px 5px 0 rgba(0,0,0,0.16),0 2px 10px 0 rgba(0,0,0,0.12);
}
.menubar {
  display: table;
  width: 100%;
  height: var(--menu-top-height);
  padding: 0px 8px;
  background-color: rgb(var(--accent-1));
  font-size: 14px;
  overflow:hidden;
  -webkit-touch-callout: none; -webkit-user-select: none; -khtml-user-select: none; 
  -moz-user-select: none; -ms-user-select: none; user-select: none;
}
.cell {
  display: table-cell;
  vertical-align: middle;
}
.menuitem {
  font-size: 17px;
  text-align: center;
  padding: 8px 6px;
  font-weight: bold;
  vertical-align: middle;
  width: auto;
  display: table-cell;
  cursor: pointer;
  -webkit-touch-callout:none;
  -webkit-user-select:none;
  -khtml-user-select:none;
  -moz-user-select:none;
  -ms-user-select:none;
  user-select:none;
}
.exit:hover {
  color: rgb(var(--functional-red))!important;
  fill: rgb(var(--functional-red))!important;
}
.selected {
  color: rgb(var(--functional-yellow))!important;
  fill: rgb(var(--functional-yellow))!important;
}
.menu-label {
  color: rgba(var(--accent-1c), 0.85);
  fill: rgba(var(--accent-1c), 0.85);
}
.menu-label:hover {
  color: rgb(var(--functional-yellow));
  fill: rgb(var(--functional-yellow));
}
.right { 
  float: right; 
}
.container { 
  padding: 0px 4px; 
}
@media (min-width:769px){
  .sidebar{
    display: none;
  }
}
@media (max-width:600px){
  .hide-small { 
    display: none!important; 
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
}
`