import { html } from 'lit';

import '../StoryContainer/story-container.js';
import './login-page.js';

export default {
  title: 'LoginPage',
  component: 'login-page',
  excludeStories: ['Template'],
  argTypes: {
    theme: {
      control: 'select',
      options: ["light", "dark"],
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

export function Template({ id, version, lang, serverURL, data, locales, theme, onPageEvent }) {
  const component = html`<login-page
    id="${id}"
    version="${version}"
    theme="${theme}"
    serverURL="${serverURL}"
    lang="${lang}"
    .locales="${locales}"
    .data="${data}"
    .onEvent=${{ onPageEvent }}
  ></login-page>`
  return html`<story-container theme="${theme}">${component}</story-container>`;
}

export const Default = Template.bind({});
Default.args = {
  id: "login_page",
  theme: "light",
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
  }
}

export const DarkServer = Template.bind({});
DarkServer.args = {
  ...Default.args,
  theme: "dark",
  serverURL: "SERVER",
  data: {
    username: "admin",
    database: "",
    server: "http://localhost:5000"
  },
}