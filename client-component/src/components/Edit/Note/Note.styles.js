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
.panel {
  width: 100%;
  background-color: rgba(var(--base-0), 1);
  border: 0.5px solid rgba(var(--neutral-1), 0.2);
  border-top: none;
  border-bottom: none;
}
.actionbar {
  width: 100%;
  background-color: rgba(var(--base-2), 1);
}
.cell {
  display: table-cell;
  vertical-align: middle;
}
.padding-tiny { 
  padding: 4px 12px 4px 4px; 
}
.padding-small { 
  padding: 4px 8px; 
}
.rtf-editor {
  width: 100%;
  min-height: 65px;
  background-color: rgba(var(--base-0), 1);
  border: 0.5px solid rgba(var(--neutral-1), 0.2);
  border-left: none;
  border-right: none;
}
.editor-content {
  height: 300px;
  outline: 0;
  overflow-y: auto;
  padding: 10px; 
}
@media (max-width:600px){
  .padding-tiny { 
    padding: 4px 12px 4px 2px; 
  }
  .padding-small { 
    padding: 4px 4px; 
  }
  .mobile{ 
    display: block; 
    width: 100%; 
  }
}
`