import { Report } from "./Report";

import { getText, store } from 'config/app';

export default {
  title: "Modal/Report",
  component: Report
}

const Template = (args) => <Report {...args} />

export const Default = Template.bind({});
Default.args = {
  title: "DMINV/00001",
  template: "ntr_invoice_en",
  templates: [
    { text: "Invoice EN", value: "ntr_invoice_en" },
    { text: "Lasku FI", value: "ntr_invoice_fi" }
  ],
  orient: "portrait",
  size: "a4", 
  copy: 1,
  className: "light",
  getText: (key)=>getText({ locales: store.session.locales, lang: "en", key: key }),
  onClose: undefined,
  onOutput: undefined,
}

export const Disabled = Template.bind({});
Disabled.args = {
  ...Default.args,
  className: "dark",
  template: "",
  templates: []
}