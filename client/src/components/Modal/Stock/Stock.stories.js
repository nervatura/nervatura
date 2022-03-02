import { Stock } from "./Stock";

import { getText, store } from 'config/app';

export default {
  title: "Modal/Stock",
  component: Stock
}

const Template = (args) => <Stock {...args} />

export const Default = Template.bind({});
Default.args = {
  partnumber: "DMPROD/00001",
  partname: "Big product",
  rows: [
    {
      batch_no: "demo",
      description: "Big product",
      id: 4,
      partnumber: "DMPROD/00001",
      shipping: "2020-11-08T00:00:00",
      sqty: 5,
      unit: "piece",
      warehouse: "material | Raw material",
    },
    {
      batch_no: "demo",
      description: "Big product",
      id: 3,
      partnumber: "DMPROD/00001",
      shipping: "2020-12-10T00:00:00",
      sqty: 2,
      unit: "piece",
      warehouse: "warehouse | Warehouse"
    }
  ],
  className: "light",
  getText: (key)=>getText({ locales: store.session.locales, lang: "en", key: key }),
  onClose: undefined,
}

export const DarkShipping = Template.bind({});
DarkShipping.args = {
  ...Default.args,
  className: "dark",
}