import { html } from 'lit';

import '../../StoryContainer/story-container.js';
import './sidebar-template.js';

import { SIDE_VISIBILITY, APP_THEME } from '../../../config/enums.js'
import * as locales from '../../../config/locales.js';

export default {
  title: 'Sidebar/Template',
  component: 'sidebar-template',
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
    onSideEvent: {
      name: "onSideEvent",
      description: "onSideEvent handler",
      table: {
        type: { 
          summary: "func", 
        },
      },
      action: "onSideEvent" 
    },
  }
};

export function Template({ id, side, templateKey, dirty, theme, onSideEvent, msg }) {
  const component = html`<sidebar-template
    id="${id}"
    side="${side}"
    templateKey="${templateKey}"
    ?dirty="${dirty}"
    .onEvent=${{ onSideEvent }}
    .msg=${msg}
  ></sidebar-template>`
  return html`<story-container theme="${theme}">${component}</story-container>`;
}

export const Default = Template.bind({});
Default.args = {
  id: "side_bar",
  theme: APP_THEME.LIGHT,
  side: SIDE_VISIBILITY.AUTO,
  templateKey: "template",
  dirty: false,
  msg: (defaultValue, props) => locales.en[props.id] || defaultValue
}

export const Sample = Template.bind({});
Sample.args = {
  ...Default.args,
  side: SIDE_VISIBILITY.SHOW,
  templateKey: "_sample",
}

export const Dirty = Template.bind({});
Dirty.args = {
  ...Default.args,
  dirty: true,
}
