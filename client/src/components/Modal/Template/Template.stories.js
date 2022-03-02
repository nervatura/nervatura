import { Template as TemplateData, DATA_TYPE } from "./Template";

import { getText, store } from 'config/app';

export default {
  title: "Modal/Template",
  component: TemplateData
}

const Template = (args) => <TemplateData {...args} />

export const Default = Template.bind({});
Default.args = {
  type: DATA_TYPE.TABLE,
  name: "table",
  columns: "col1,col2,col3",
  className: "light",
  getText: (key)=>getText({ locales: store.session.locales, lang: "en", key: key }),
  onClose: undefined,
  onData: undefined,
}

export const Disabled = Template.bind({});
Disabled.args = {
  ...Default.args,
  className: "dark",
  type: DATA_TYPE.TEXT,
  name: "",
  columns: "",
}