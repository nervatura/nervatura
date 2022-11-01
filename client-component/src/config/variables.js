import { css } from 'lit';

export default css`
html, :host {
  --font-family: "Noto Sans";
  --font-size: 14px;
  --menu-top-height: 43.5px;
  --menu-side-width: 250px;
}
html, :host, *[theme="light"] {
  --neutral-1: 0, 0, 0;
  --neutral-2: 255, 255, 255;
  --accent-1: 0, 28, 50;
  --accent-1b: 0, 71, 93;
  --accent-1c: 255, 255, 255;
  --base-0: 255, 255, 255;
  --base-1: 235, 235, 235;
  --base-2: 245, 245, 245;
  --base-3: 255, 255, 255;
  --base-4: 255, 255, 255;
  --functional-blue: 20, 120, 220;
  --functional-red: 210, 105, 125;
  --functional-yellow: 220, 168, 40;
  --functional-green: 50, 168, 40;
  --text-1: rgba(0, 0, 0, .90);
  --text-2: rgba(0, 0, 0, .60);
  --text-3: rgba(0, 0, 0, .20);
  --shadow-1: 0 2px 8px rgba(0,0,0,.1), 0 1px 4px rgba(0,0,0,.05);
}
*[theme="dark"] {
  --neutral-1: 255, 255, 255;
  --neutral-2: 0, 0, 0;
  --accent-1: 0, 28, 50;
  --accent-1b: 0, 71, 93;
  --accent-1c: 255, 255, 255;
  --base-0: 0, 0, 2;
  --base-1: 15, 15, 15;
  --base-2: 25, 25, 25;
  --base-3: 35, 35, 35;
  --base-4: 45, 45, 45;
  --functional-blue: 20, 120, 220;
  --functional-red: 210, 105, 125;
  --functional-yellow: 220, 160, 40;
  --functional-green: 40, 160, 40;
  --text-1: rgba(255, 255, 255, .90);
  --text-2: rgba(255, 255, 255, .60);
  --text-3: rgba(255, 255, 255, .20);
  --shadow-1: 0 2px 8px rgba(0,0,0,.2), 0 1px 4px rgba(0,0,0,.15);
}
`