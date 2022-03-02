import { Trans } from "./Trans";

import { getText, store } from 'config/app';

export default {
  title: "Modal/Trans",
  component: Trans
}

const Template = (args) => <Trans {...args} />

export const Default = Template.bind({});
Default.args = {
  baseTranstype: "offer",
  transtype: "order",
  direction: "out",
  doctypes: ["offer","order","worksheet","rent"],
  directions: ["in","out"],
  refno: true, 
  nettoDiv: false, 
  netto: true, 
  fromDiv: false, 
  from: false,
  elementCount: 10,
  className: "light",
  getText: (key)=>getText({ locales: store.session.locales, lang: "en", key: key }),
  onClose: undefined,
  onCreate: undefined,
}

export const Order = Template.bind({});
Order.args = {
  ...Default.args,
  baseTranstype: "order",
  transtype: "invoice",
  direction: "in",
  doctypes: ["offer","order","worksheet","rent","invoice","receipt"],
  refno: true, 
  nettoDiv: true, 
  netto: true, 
  fromDiv: true, 
  from: false,
  elementCount: 0,
  className: "dark",
}

export const Worksheet = Template.bind({});
Worksheet.args = {
  ...Default.args,
  baseTranstype: "worksheet",
  transtype: "invoice",
  direction: "in",
  doctypes: ["offer","order","worksheet","rent","invoice","receipt"],
  refno: true, 
  nettoDiv: false, 
  netto: true, 
  fromDiv: true, 
  from: false,
  elementCount: 10,
  className: "dark",
}