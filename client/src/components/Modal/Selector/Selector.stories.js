import { Selector } from "./Selector";

import { getText, store } from 'config/app';

export default {
  title: "Modal/Selector",
  component: Selector
}

const Template = (args) => <Selector {...args} />

export const Default = Template.bind({});
Default.args = {
  view: "customer",
  columns: [
    ["custname", "c.custname"],
    ["custnumber", "c.custnumber"],
    ["city", "addr.city"],
    ["street", "addr.street"]
  ],
  result: [
    {
      city: 'City1',
      custname: 'First Customer Co.',
      custnumber: 'DMCUST/00001',
      id: 'customer//2',
      label: 'First Customer Co.',
      street: 'street 1.'
    },
    {
      city: 'City3',
      custname: 'Second Customer Name',
      custnumber: 'DMCUST/00002',
      id: 'customer//3',
      label: 'Second Customer Name',
      street: 'street 3.',
      deleted: 1
    },
    {
      city: 'City4',
      custname: 'Third Customer Foundation',
      custnumber: 'DMCUST/00003',
      id: 'customer//4',
      label: 'Third Customer Foundation',
      street: 'street 4.'
    }
  ],
  filter: "",
  className: "light",
  getText: (key)=>getText({ locales: store.session.locales, lang: "en", key: key }),
  onClose: undefined,
  onSearch: undefined,
  onSelect: undefined
}

export const QuickView = Template.bind({});
QuickView.args = {
  ...Default.args,
  className: "dark",
  onClose: null,
}