import { css } from 'lit';

export const styles = css`
button {
  font-family: var(--font-family);
  font-size: var(--font-size);
  color: var(--text-1);
  fill: var(--text-1);
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
  color: rgb(255, 255, 255);
  fill: rgb(255, 255, 255);
  border-color: rgba(var(--neutral-1), 0.2);
  background-color: rgb(var(--accent-1));
}
button[button-type='secondary'] {
  border-color: rgba(var(--neutral-1), 0.2);
  background-color: rgba(var(--neutral-1), 0.1);
}
button[button-type='border'] {
  background: none;
  border-width: 1px;
  border-style: solid;
  border-color: rgba(var(--neutral-1), 0.3);
  border-radius: 0;
}
button[disabled] {
  pointer-events: none;
  opacity: 0.3;
}
button[button-type='primary']:not(:active):hover {
  background-color: rgb(var(--accent-1b));
}
button[button-type='secondary']:not(:active):hover {
  background-color: rgba(var(--neutral-1), 0.15);
}
button[button-type='border']:not(:active):hover {
  color: rgb(var(--accent-1b));
  border-color: rgb(var(--accent-1b));
}
button:disabled:hover {
  box-shadow:none;
}
.small {
  font-size: 12px;
  padding: 6px 8px;
}
.full{
  width: 100%;
}
`