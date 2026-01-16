import { html } from 'lit';

import '../../StoryContainer/story-container.js';
import './modal-template.js';

import { APP_THEME, TEMPLATE_DATA_TYPE } from '../../../config/enums.js'
import * as locales from '../../../config/locales.js';

const msg = (defaultValue, props) => locales.en[props.id] || defaultValue

export default {
  title: 'Modal/Template',
  component: 'modal-template',
  excludeStories: ['Template'],
  argTypes: {
    theme: {
      control: 'select',
      options: Object.values(APP_THEME),
    },
    type: {
      control: 'select',
      options: Object.values(TEMPLATE_DATA_TYPE),
    },
    onModalEvent: {
      name: "onModalEvent",
      description: "onModalEvent handler",
      table: {
        type: { 
          summary: "func", 
        },
      },
      action: "onModalEvent" 
    },
  }
};

export function Template({ 
  id, type, name, columns, theme, onModalEvent 
}) {
  const component = html`<modal-template
    id="${id}"
    type="${type}"
    name="${name}"
    columns="${columns}"
    .onEvent=${{ onModalEvent }}
    .msg=${msg}
  ></modal-template>`
  return html`<story-container theme="${theme}">${component}</story-container>`;
}

export const Default = Template.bind({});
Default.args = {
  id: "template",
  theme: APP_THEME.LIGHT,
  type: TEMPLATE_DATA_TYPE.TABLE,
  name: "table",
  columns: "col1,col2,col3",
}

export const DarkForm = Template.bind({});
DarkForm.args = {
  ...Default.args,
  theme: APP_THEME.DARK,
  type: TEMPLATE_DATA_TYPE.TEXT,
  name: "",
  columns: "",
}