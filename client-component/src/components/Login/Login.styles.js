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
  .modal {
    position: fixed;
    left: 0;
    right: 0;
    top: 0;
    height: 100vh;
    overflow: auto;
    background-color: rgba(25, 25, 25, 0.7);
    padding: 10px 5px;
  }
  .middle {
    margin: 20px 10px 10px;
  }
  .dialog {
    border-radius: 0px;
    border: 0.5px solid rgba(var(--neutral-1), 0.2);
    box-shadow:0 4px 10px 0 rgba(0,0,0, 0.2),0 4px 20px 0 rgba(0,0,0, 0.19);
    background-color: rgba(var(--base-2), 1);
    margin: 0 auto;
    animation: animatezoom 0.6s;
    width: 100%;
    max-width: 400px;
    min-width: 280px;
    top: 0;
    left: 0;
  }
  .title {
    width: 100%;
    border-bottom: 0.5px solid rgba(var(--neutral-1), 0.2);
    background-color: rgb(var(--accent-1));
    color: rgba(var(--accent-1c), 0.85);
    fill: rgba(var(--accent-1c), 0.85);
  }
  .title-cell {
    padding: 8px 16px;
    font-weight: bold;
  }
  .version-cell {
    padding: 8px 16px;
    text-align: right;
    font-size: 13px;
    font-weight: normal;
  }
  .label-cell {
    width: 35%;
  }
  .buttons {
    background-color: rgb(var(--base-1));
    border-top: 0.5px solid rgba(var(--neutral-1), 0.2);
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
  .section { 
    padding-top: 16px!important; 
    padding-bottom: 16px!important; 
  }
  .section-small { 
    padding-top: 8px!important; 
    padding-bottom: 8px!important; 
  }
  .section-small-bottom { 
    padding-bottom: 8px!important; 
  }
  .container { 
    padding: 0.01em 16px; 
  } 
  .container::after, .container::before { 
    content: ""; display: table; clear: both; 
  }
  .align-right { 
    text-align: right; 
  }
  .padding-normal { 
    padding: 8px 16px; 
  }
  @media (max-width:600px){
    .mobile{ 
      display: block; 
      width: 100%; 
    }
    .container { 
      padding: 0px 8px; 
    }
    .padding-normal { 
      padding: 4px 8px; 
    }
  }
  @media only screen and (min-width: 601px){
    .middle {
      margin: 0px;
      position:absolute;
      top:50%;
      left:50%;
      transform:translate(-50%,-50%);
      -ms-transform:translate(-50%,-50%)
    }
    .dialog {
      min-width: 400px;
    }
  }
`