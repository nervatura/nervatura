import { Formula } from "./Formula";

import { getText, store } from 'config/app';

export default {
  title: "Modal/Formula",
  component: Formula
}

const Template = (args) => <Formula {...args} />

export const Default = Template.bind({});
Default.args = {
  formula: "18",
  formulaValues: [
    {value: "18", text: 'DMFRM/00001'},
    {value: "19", text: 'DMFRM/00002'}
  ],
  partnumber: "DMPROD/00004",
  description: "Car",
  className: "light",
  getText: (key)=>getText({ locales: store.session.locales, lang: "en", key: key }),
  onClose: undefined,
  onFormula: undefined,
}

export const Disabled = Template.bind({});
Disabled.args = {
  ...Default.args,
  className: "dark",
  formula: "",
  formulaValues: []
}