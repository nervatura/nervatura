import { html } from 'lit';

import '../StoryContainer/story-container.js';
import './client-login.js';

import { APP_THEME } from '../../config/enums.js'
import * as _locales from '../../config/locales.js';

export default {
  title: 'Login',
  component: 'client-login',
  excludeStories: ['Template'],
  argTypes: {
    theme: {
      control: 'select',
      options: Object.values(APP_THEME),
    },
    onPageEvent: {
      name: "onPageEvent",
      description: "onPageEvent handler",
      table: {
        type: { 
          summary: "func", 
        },
      },
      action: "onPageEvent" 
    },
  }
};

export function Template({ id, version, lang, serverURL, data, locales, theme, onPageEvent, msg }) {
  const component = html`<client-login
    id="${id}"
    version="${version}"
    theme="${theme}"
    serverURL="${serverURL}"
    lang="${lang}"
    .locales="${locales}"
    .data="${data}"
    .onEvent=${{ onPageEvent }}
    .msg=${msg}
  ></client-login>`
  return html`<story-container theme="${theme}">${component}</story-container>`;
}

export const Default = Template.bind({});
Default.args = {
  id: "login_page",
  theme: APP_THEME.LIGHT,
  version: "5.1.0",
  lang: "en",
  serverURL: "",
  data: {
    username: "admin",
    database: "demo",
    server: "http://localhost:5000"
  },
  locales: {
    en:{en: "English"}, jp:{"jp": "日本語"} , au:{}
  },
  msg: (defaultValue, props) => _locales.en[props.id] || defaultValue
}

export const DarkServer = Template.bind({});
DarkServer.args = {
  ...Default.args,
  theme: APP_THEME.DARK,
  serverURL: "SERVER",
  data: {
    username: "admin",
    database: "",
    server: "http://localhost:5000"
  },
}