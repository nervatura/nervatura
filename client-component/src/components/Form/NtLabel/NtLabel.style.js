import { css } from 'lit';

export const styles = css`
  :host {
    font-family: var(--font-family);
    font-size: var(--font-size);
    color: var(--text-1);
    fill: var(--text-1);
  }
  .row {
    display: table;
  }
  .cell {
    display: table-cell;
    vertical-align: middle;
    line-height: 1;
  }
  .label_info_left {
    padding: 0 0 0 3px;
  }
  .label_info_right {
    text-align: right;
    padding: 0 3px 0 0;
  }
  .label_icon_left {
    text-align: center;
    width: auto;
  }
  .label_icon_right {
    text-align: right;
    width: auto;
  }
  .centered {
    margin: auto;
  }
  .bold {
    font-weight: bold;
  }
`