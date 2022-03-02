import { Login } from "./Login";

import { getText, store } from 'config/app';

export default {
  title: "Login",
  component: Login
}

const Template = (args) => <Login {...args} />

export const Default = Template.bind({});
Default.args = {
  theme: "light",
  version: store.session.version,
  locales: { en:{en: "English"}, jp:{"jp": "日本語"} , au:{} },
  lang: "en",
  configServer: false,
  getText: (key)=>getText({ locales: store.session.locales, lang: "en", key: key }),
  //changeData: (key, value) => {},
  //setLocale: (value) => {}
}

export const DarkLogin = Template.bind({});
DarkLogin.args = {
  ...Default.args,
  theme: "dark",
  username: "admin",
  database: "demo",
  configServer: true
}