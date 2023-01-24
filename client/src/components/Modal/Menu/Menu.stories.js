import { html } from 'lit';

import '../../StoryContainer/story-container.js';
import './modal-menu.js';

import { APP_THEME } from '../../../config/enums.js'
import * as locales from '../../../config/locales.js';

const msg = (defaultValue, props) => locales.en[props.id] || defaultValue

export default {
  title: 'Modal/Menu',
  component: 'modal-menu',
  excludeStories: ['Template'],
  argTypes: {
    theme: {
      control: 'select',
      options: Object.values(APP_THEME),
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
  id, idKey, menu_id, fieldname, description, orderby, fieldtype, fieldtypeOptions, theme, onModalEvent 
}) {
  const component = html`<modal-menu
    id="${id}"
    idKey="${idKey}"
    menu_id="${menu_id}"
    fieldname="${fieldname}"
    description="${description}"
    orderby="${orderby}"
    fieldtype="${fieldtype}"
    .fieldtypeOptions="${fieldtypeOptions}"
    .onEvent=${{ onModalEvent }}
    .msg=${msg}
  ></modal-menu>`
  return html`<story-container theme="${theme}">${component}</story-container>`;
}

export const Default = Template.bind({});
Default.args = {
  id: "menu",
  theme: APP_THEME.LIGHT,
  idKey: null,
  menu_id: 1, 
  fieldname: "fieldname", 
  description: "description", 
  fieldtype: 38, 
  fieldtypeOptions: [
    { value: "33", text: "bool" },
    { value: "34", text: "date" },
    { value: "36", text: "float" },
    { value: "37", text: "integer" },
    { value: "38", text: "string" },
  ], 
  orderby: 1,
}

export const Disabled = Template.bind({});
Disabled.args = {
  ...Default.args,
  theme: APP_THEME.DARK,
  idKey: 1,
  fieldname: "", 
  description: "",
  fieldtype: 34,
}
