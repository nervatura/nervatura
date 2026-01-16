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
.page {
  transition: margin-left .4s;
  margin-left: var(--menu-side-width);
  padding: 8px;
}
@media (max-width:768px){
  .page{
    margin-left: 0px;
  }
}
@media (max-width:600px){
  .page {
    padding: 4px;
  }
}
`