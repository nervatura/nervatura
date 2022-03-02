import { Meta } from "./Meta";

import { getText, store } from 'config/app';
import { Default as EditorDefault } from 'components/Editor/Editor/Editor.stories'

export default {
  title: "Editor/Meta",
  component: Meta,
}

const Template = (args) => <Meta {...args} />

export const Default = Template.bind({});
Default.args = {
  current: {
    ...EditorDefault.args.current,
    view: "fieldvalue"
  }, 
  template: EditorDefault.args.template, 
  dataset: EditorDefault.args.dataset, 
  audit: "all",
  className: "light",
  onEvent: undefined,
  getText: (key)=>getText({ locales: store.session.locales, lang: "en", key: key })
}

export const NewField = Template.bind({});
NewField.args = {
  ...Default.args,
  current: {
    ...Default.args.current,
    deffield: "trans_transitem_link",
    page: 0,
  },
  dataset: {
    ...Default.args.dataset,
    deffield_prop: []
  },
  paginationPage: 2
}

export const ReadOnly = Template.bind({});
ReadOnly.args = {
  ...Default.args,
  audit: "readonly"
}

export const Customer = Template.bind({});
Customer.args = {
  ...Default.args,
  current: {
    type: "customer",
    transtype: "",
    item: { id: 2, custtype: 116, custnumber: "DMCUST/00001", custname: "First Customer Co.",
      taxnumber: "12345678-1-12", account: null, notax: 0, terms: 8, creditlimit: 1000000,
      discount: 2, notes: null, inactive: 0, deleted: 0 },
    state: "normal",
    page: 0,
    fieldvalue: [
      { deleted: 0, fieldname: "sample_customer_date", id: 27, notes: "", 
        ref_id: 2, value: "2022-08-12" },
      {
        deleted: 0, fieldname: "sample_customer_float", id: 28, notes: "",
        ref_id: 2, value: "123.4" }
    ],
    view: "form"
  }
}