import { Server } from "./Server";

import { getText, store } from 'config/app';

export default {
  title: "Modal/Server",
  component: Server
}

const Template = (args) => <Server {...args} />

export const Default = Template.bind({});
Default.args = {
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
  className: "light",
  getText: (key)=>getText({ locales: store.session.locales, lang: "en", key: key }),
  onClose: undefined,
  onOK: undefined,
}

export const DarkServer = Template.bind({});
DarkServer.args = {
  ...Default.args,
  className: "dark",
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