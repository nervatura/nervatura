import { Total } from "./Total";

import { getText, store } from 'config/app';

export default {
  title: "Modal/Total",
  component: Total
}

const Template = (args) => <Total {...args} />

export const Default = Template.bind({});
Default.args = {
  total: {
    totalFields: {
      netamount: 11221,
      vatamount: 2097.2,
      amount: 13318.2,
      acrate: 0,
    },
    totalLabels: {
      netamount: "Net Amount",
      vatamount: "VAT",
      amount: "Amount",
      acrate: "Acc.Rate",
    },
    count: 4,
  },
  className: "light",
  getText: (key)=>getText({ locales: store.session.locales, lang: "en", key: key }),
  onClose: undefined,
}

export const DarkTotal = Template.bind({});
DarkTotal.args = {
  ...Default.args,
  className: "dark",
}