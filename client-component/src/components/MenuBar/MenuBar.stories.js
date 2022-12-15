import { html } from 'lit';

import '../StoryContainer/story-container.js';
import './client-menubar.js';

import { SIDE_VISIBILITY, APP_MODULE, APP_THEME } from '../../config/enums.js'

export default {
  title: 'MenuBar',
  component: 'client-menubar',
  excludeStories: ['Template'],
  argTypes: {
    theme: {
      control: 'select',
      options: Object.values(APP_THEME),
    },
    side: {
      control: 'select',
      options: Object.values(SIDE_VISIBILITY),
    },
    module: {
      control: 'select',
      options: Object.values(APP_MODULE),
    },
    onMenuEvent: {
      name: "onMenuEvent",
      description: "onMenuEvent handler",
      table: {
        type: { 
          summary: "func", 
        },
      },
      action: "onMenuEvent" 
    },
  }
};

export function Template({ id, side, module, scrollTop, theme, onMenuEvent }) {
  const component = html`<client-menubar
    id="${id}"
    side="${side}"
    module="${module}"
    ?scrollTop="${scrollTop}"
    .onEvent=${{ 
      onMenuEvent 
    }}
    .msg=${(defValue) => defValue}
    .app=${{ store: {} }}
  ></client-menubar>`
  return html`<story-container theme="${theme}">${component}</story-container>`;
}

export const Default = Template.bind({});
Default.args = {
  id: "menu_bar",
  theme: APP_THEME.LIGHT,
  side: SIDE_VISIBILITY.AUTO,
  module: APP_MODULE.SEARCH,
  scrollTop: false
}

export const ScrollTop = Template.bind({});
ScrollTop.args = {
  ...Default.args,
  side: SIDE_VISIBILITY.SHOW,
  scrollTop: true,
  module: APP_MODULE.SETTING,
}