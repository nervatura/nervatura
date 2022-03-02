import { Shipping } from "./Shipping";

import { getText, store } from 'config/app';

export default {
  title: "Modal/Shipping",
  component: Shipping
}

const Template = (args) => <Shipping {...args} />

export const Default = Template.bind({});
Default.args = {
  partnumber: "DMPROD/00008",
  description: "DMPROD/00008 | Pallet",
  unit: "piece",
  batch_no: "",
  qty: 6,
  className: "light",
  getText: (key)=>getText({ locales: store.session.locales, lang: "en", key: key }),
  onClose: undefined,
  onShipping: undefined,
}

export const DarkShipping = Template.bind({});
DarkShipping.args = {
  ...Default.args,
  className: "dark",
}