import { html } from 'lit';

import '../../StoryContainer/story-container.js';
import './form-icon.js';
import { IconData } from './IconData.js'

import { APP_THEME } from '../../../config/enums.js'

export default {
  title: 'Form/Icon',
  component: 'form-icon',
  excludeStories: ['Template'],
  argTypes: {
    theme: {
      control: 'select',
      options: Object.values(APP_THEME),
    },
    onClick: {
      name: "onClick",
      description: "onClick click handler",
      table: {
        type: { 
          summary: "func", 
        },
      },
      action: "onClick" 
    },
    iconKey: {
      control: 'select',
      options: Object.keys(IconData).sort(),
    }
  }
};

export function Template({ id, iconKey, width, height, color, style, theme, onClick }) {
  const component = html`<form-icon
    id="${id}"
    iconKey="${iconKey}"
    width="${width}"
    height="${height}"
    .color=${color}
    .style=${style}
    @click=${onClick}
  ></form-icon>`
  return html`<story-container theme="${theme}">${component}</story-container>`;
}

export const Default = Template.bind({});
Default.args = {
  id: "test_icon",
  iconKey: "ExclamationTriangle",
  theme: APP_THEME.LIGHT,
  style: {}
};

export const DarkIcon = Template.bind({});
DarkIcon.args = {
  ...Default.args,
  theme: APP_THEME.DARK,
  style: "opacity: 1;"
};

export const ColorPointer = Template.bind({});
ColorPointer.args = {
  ...Default.args,
  iconKey: "Copy",
  width: 42,
  height: 48,
  color: "red",
  style: {cursor: "pointer"}
}

export const InvalidIcon = Template.bind({});
InvalidIcon.args = {
  ...Default.args,
  iconKey: "Copy123"
}