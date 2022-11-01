import { html } from 'lit';

import '../../StoryContainer/story-container.js';
import './modal-server.js';

import { APP_THEME } from '../../../config/enums.js'
import * as locales from '../../../config/locales.js';

export default {
  title: 'Modal/Server',
  component: 'modal-server',
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

export function Template({ id, cmd, fields, values, theme, onModalEvent, msg }) {
  const component = html`<modal-server
    id="${id}"
    .cmd=${cmd}
    .fields=${fields}
    .values=${values}
    .onEvent=${{ onModalEvent }}
    .msg=${msg}
  ></modal-server>`
  return html`<story-container theme="${theme}" >${component}</story-container>`;
}

export const Default = Template.bind({});
Default.args = {
  id: "server",
  theme: APP_THEME.LIGHT,
  cmd: {
    address: null,
    description: "Server function example",
    funcname: "nextNumber",
    icon: null,
    id: 1,
    menukey: "nextNumber",
    method: 130,
    methodName: "post",
    modul: null
  }, 
  fields: [
    {
      description: "Code (e.g. custnumber)",
      fieldname: "numberkey",
      fieldtype: 38,
      fieldtypeName: "string",
      id: 1,
      menu_id: 1,
      orderby: 0,
    },
    {
      description: "Stepping",
      fieldname: "step",
      fieldtype: 33,
      fieldtypeName: "bool",
      id: 2,
      menu_id: 1,
      orderby: 1,
    },
  ], 
  values: {
    numberkey: "",
    step: false,
  },
  msg: (defaultValue, props) => locales.en[props.id] || defaultValue
}

export const DarkServer = Template.bind({});
DarkServer.args = {
  ...Default.args,
  theme: APP_THEME.DARK,
  cmd: {
    address: "https://www.google.com",
    description: "Internet URL example",
    funcname: "search",
    icon: null,
    id: 2,
    menukey: "google",
    method: 129,
    methodName: "get",
    modul: null,
  },
  fields: [
    {
      description: "google search",
      fieldname: "q",
      fieldtype: 38,
      fieldtypeName: "string",
      id: 3,
      menu_id: 2,
      orderby: 0,
    },
  ],
  values: {
    q: "",
  }
}


