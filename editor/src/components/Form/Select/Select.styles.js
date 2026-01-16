import { css } from 'lit';

export const styles = css`
select {
  font-family: var(--font-family);
  font-size: 13px;
  border-radius: 3px;
  padding: 8px;
  display: block;
  color: var(--text-1);
  background-color: rgba(var(--base-4), 1);
  border: 1px solid rgba(var(--neutral-1), 0.2);
  box-sizing: border-box;
}
select:disabled {
  opacity: 0.5;
}
select:focus, select:hover:not(:disabled) {
  color: rgb(var(--functional-green))!important;
}
option {
  font-size: 14px;
  border: none;
}
option:disabled {
  opacity: 0.5;
}
.full{
  width: 100%;
}
`