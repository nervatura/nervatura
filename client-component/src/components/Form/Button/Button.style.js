import { css } from 'lit';

export const styles = css`
button {
  font-family: var(--font-family);
  font-size: var(--font-size);
  color: var(--text-1);
  fill: var(--text-1);
  border-color: rgba(var(--neutral-1), 0.2);
  background-color: rgba(var(--neutral-1), 0.1);
  font-weight: bold;
  display: inline-block;
  padding: 8px 16px;
  vertical-align: middle;
  overflow: hidden;
  text-decoration: none;
  text-align: center;
  cursor:pointer;
  white-space:nowrap;
  -webkit-touch-callout:none;
  -webkit-user-select:none;
  -khtml-user-select:none;
  -moz-user-select:none;
  -ms-user-select:none;
  user-select:none;
  border-radius: 3px;
  border-width: 1px;
  transition: 0.1s all ease-out;
  box-sizing: border-box;
}
button[button-type='primary'] {
  color: rgba(var(--accent-1c), 0.85);
  fill: rgba(var(--accent-1c), 0.85);
  border-color: rgba(var(--neutral-1), 0.2);
  background-color: rgb(var(--accent-1));
}
button[button-type='border'] {
  background: none;
  border-width: 1px;
  border-style: solid;
  border-color: rgba(var(--neutral-1), 0.3);
}
button[disabled] {
  pointer-events: none;
  opacity: 0.3;
}
button:not(:active):hover {
  background-color: rgba(var(--neutral-1), 0.15);
}
button[button-type='primary']:not(:active):hover {
  background-color: rgb(var(--accent-1b));
}
button[button-type='border']:not(:active):hover {
  color: rgb(var(--accent-1c));
  fill: rgb(var(--accent-1c));
  border-color: rgb(var(--accent-1b));
}
button:disabled:hover {
  box-shadow:none;
}
button:focus, button:hover {
  color: rgb(var(--functional-green));
  fill: rgb(var(--functional-green));
}
.selected {
  color: rgb(var(--functional-yellow))!important;
  fill: rgb(var(--functional-yellow))!important;
}
.small {
  font-size: 12px;
  padding: 6px 8px!important;
}
.full{
  width: 100%;
}
@media (max-width:600px){
  .hidelabel slot { 
    display: none;
  }
  .hidelabel {
    padding: 8px 12px;
  }
}
`