import { css } from 'lit';

export const styles = css`
input {
  font-family: var(--font-family);
  font-size: var(--font-size);
  border-radius: 3px;
  border-width: 1px;
  padding: 8px;
  display: block;
  color: var(--text-1);
  background-color: rgba(var(--base-4), 1);
  border: 1px solid rgba(var(--neutral-1), 0.2);
  box-sizing: border-box;
}
input[type=color]{
  height: 34px;
  border: none;
  padding: 0px;
  cursor: pointer;
}
input::placeholder, input::-ms-input-placeholder {
  opacity: 0.5;
}
input:disabled {
  opacity: 0.5;
}
input:focus, input:hover:not(:disabled) {
  color: rgb(var(--functional-green))!important;
}
.full{
  width: 100%;
}
`