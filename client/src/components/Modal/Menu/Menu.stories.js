import { Menu } from "./Menu";

import { getText, store } from 'config/app';

export default {
  title: "Modal/Menu",
  component: Menu
}

const Template = (args) => <Menu {...args} />

export const Default = Template.bind({});
Default.args = {
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
  className: "light",
  getText: (key)=>getText({ locales: store.session.locales, lang: "en", key: key }),
  onClose: undefined,
  onMenu: undefined,
}

export const Disabled = Template.bind({});
Disabled.args = {
  ...Default.args,
  className: "dark",
  idKey: 1,
  fieldname: "", 
  description: "",
  fieldtype: 34,
}