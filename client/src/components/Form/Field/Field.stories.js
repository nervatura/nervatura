import { Field } from "./Field";
import { getText, store } from 'config/app';

export default {
  title: "Form/Field",
  component: Field
}

const Template = (args) => <Field {...args} />

export const Default = Template.bind({});
Default.args = {
  field: {
    rowtype: "field",
    name: "custname",
    label: "Customer Name",
    datatype: "string",
    length: 30,
  },
  values: {
    id: 2,
    custtype: 116,
    custnumber: "DMCUST/00001",
    custname: null,
  },
  options: {},
  data: {
    dataset: {}, 
    current: {}, 
    audit: "all",
  },
  className: "light",
  getText: (key)=>getText({ locales: store.session.locales, lang: "en", key: key }),
  onEdit: undefined,
  onEvent: undefined,
  onSelector: undefined,
}

export const StringMapExtend = Template.bind({});
StringMapExtend.args = {
  ...Default.args,
  className: "dark",
  field: {
    name: "id",
    label: "Firstname",
    datatype: "string",
    map: {
      source: "contact",
      value: "ref_id",
      text: "firstname",
      extend: true,
    },
  },
  values: {
    id: 5,
    empnumber: "DMEMP/00001",
    username: null,
  },
  data: {
    dataset: {
    }, 
    current: {
      type: "employee",
      extend: {
        id: 9,
        nervatype: 12,
        ref_id: 5,
        firstname: "John",
        surname: "Strong",
      },
    }, 
    audit: "all",
  },
}

export const StringMapLinkID = Template.bind({});
StringMapLinkID.args = {
  ...Default.args,
  field: {
    name: "direction",
    label: "Delivery Type",
    datatype: "string",
    map: {
      source: "groups",
      value: "id",
      text: "groupvalue",
      label: "delivery",
    },
  },
  values: {
    id: 14,
    transtype: 61,
    direction: 68,
  },
  data: {
    dataset: {
      groups: [
        {
          deleted: 0,
          description: null,
          groupname: "direction",
          groupvalue: "out",
          id: 68,
          inactive: 0,
        },
      ]
    }, 
    current: {
    }, 
    audit: "all",
  },
}

export const StringMapLinkValue = Template.bind({});
StringMapLinkValue.args = {
  ...Default.args,
  field: {
    rowtype: "field",
    name: "trans_rentnote",
    label: "Justification",
    datatype: "string",
    map: {
      source: "fieldvalue",
      value: "fieldname",
      text: "value",
    },
  },
  values: {
    id: 9,
    transtype: 60,
    direction: 68,
    transnumber: "DMRNT/00001",
  },
  data: {
    dataset: {
    }, 
    current: {
      fieldvalue:[
        {
          id: null,
          fieldname: "trans_rentnote",
          ref_id: 9,
          value: "value",
          notes: null,
          deleted: 0,
        },
      ]
    }, 
    audit: "all",
  },
}

export const TextNull = Template.bind({});
TextNull.args = {
  ...Default.args,
  field: {
    rowtype: "field",
    name: "notes",
    label: "Comment",
    datatype: "text",
  },
  values: {
    id: 2,
    custtype: 116,
    notes: null,
    inactive: 0,
    deleted: 0,
  },
  data: {
    dataset: {}, 
    current: {}, 
    audit: "all",
  },
}

export const NotesValue = Template.bind({});
NotesValue.args = {
  ...Default.args,
  field: {
    rowtype: "fieldvalue",
    id: 127,
    name: "fieldvalue_value",
    fieldname: "1aba0e61",
    value: "fieldvalue",
    notes: "",
    label: "Customer notes",
    description: null,
    disabled: true,
    fieldtype: "notes",
    datatype: "notes",
    rows: 5
  },
  values: {
    rowtype: "fieldvalue",
    id: 127,
    name: "fieldvalue_value",
    fieldname: "1aba0e61",
    value: "fieldvalue",
    notes: "",
    label: "Customer notes",
    description: null,
    disabled: false,
    fieldtype: "notes",
    datatype: "notes",
  },
  data: {
    dataset: {}, 
    current: {}, 
    audit: "all",
  },
}

export const Float = Template.bind({});
Float.args = {
  ...Default.args,
  field: {
    name: "creditlimit",
    label: "Credit line",
    datatype: "float",
  },
  values: {
    id: 2,
    custtype: 116,
    creditlimit: null,
  },
  data: {
    dataset: {}, 
    current: {}, 
    audit: "all",
  },
}

export const FloatLinkValue = Template.bind({});
FloatLinkValue.args = {
  ...Default.args,
  field: {
    name: "trans_wsdistance",
    label: "Distance (km)",
    datatype: "float",
    map: {
      source: "fieldvalue",
      value: "fieldname",
      text: "value",
    },
    opposite: true
  },
  values: {
    id: 8,
    transtype: 59,
    direction: 68,
    transnumber: "DMWORK/00001",
  },
  data: {
    dataset: {}, 
    current: {
      fieldvalue: [
        {
          deleted: 0,
          fieldname: "trans_wsdistance",
          id: 76,
          notes: "",
          ref_id: 8,
          value: "200.0",
        },
      ]
    }, 
    audit: "readonly",
  },
}

export const Integer = Template.bind({});
Integer.args = {
  ...Default.args,
  field: {
    name: "id",
    label: "Amount",
    datatype: "integer",
    opposite: true,
    map: {
      source: "payment",
      value: "trans_id",
      text: "amount",
      extend: true,
    },
  },
  values: {
    id: 11,
    transtype: 67,
    direction: 68,
    transnumber: "DMPMT/00002",
    ref_transnumber: null,
  },
  options: {
    opposite: true
  },
  data: {
    dataset: {}, 
    current: {
      extend: {
        id: 4,
        trans_id: 11,
        paiddate: "2020-12-18",
        amount: -488,
        notes: null,
        deleted: 0,
        rid: 4,
      }
    }, 
    audit: "all",
  },
}

export const IntegerLinkID = Template.bind({});
IntegerLinkID.args = {
  ...Default.args,
  field: {
    name: "trans_id",
    label: "Test integer",
    datatype: "integer",
    map: {
      source: "trans",
      value: "id",
      text: "testint",
    },
  },
  values: {
    id: 12,
    item_id: 10,
    trans_id: 14,
  },
  data: {
    dataset: {}, 
    current: {
      trans: [
        {
          id: 14,
          transcast: "normal",
          testint: 345,
          transnumber: "DMDEL/00002",
        },
      ]
    }, 
    audit: "all",
  },
}

export const Button = Template.bind({});
Button.args = {
  ...Default.args,
  field: {
    name: "log_search",
    title: "Search",
    label: "",
    focus: true,
    class: "full",
    icon: "Search",
    datatype: "button",
  },
  values: {},
  data: {
    dataset: {}, 
    current: {}, 
    audit: "all",
  },
  className: "dark",
}

export const Link = Template.bind({});
Link.args = {
  ...Default.args,
  field: {
    name: "id",
    label: "Reference No.",
    datatype: "link",
    map: {
      source: "translink",
      value: "ref_id_1",
      text: "ref_id_2",
      label_field: "transnumber",
      lnktype: "trans",
      transtype: "order",
    },
  },
  values: {
    id: 5,
    transtype: 55,
    direction: 68,
    transnumber: "DMINV/00001",
    ref_transnumber: "DMORD/00003",
  },
  data: {
    dataset: {
      translink: [
        {
          deleted: 0,
          id: 2,
          nervatype_1: 31,
          nervatype_2: 31,
          ref_id_1: 5,
          ref_id_2: 3,
          transnumber: "DMORD/00003",
          transtype: "order",
        },
      ]
    }, 
    current: {}, 
    audit: "all",
  },
  className: "dark"
}

export const ValueList = Template.bind({});
ValueList.args = {
  ...Default.args,
  field: {
    rowtype: "fieldvalue",
    id: 29,
    name: "fieldvalue_value",
    fieldname: "sample_customer_valuelist",
    value: "yellow",
    notes: "",
    label: "Sample valuelist",
    description: [
      "blue",
      "yellow",
      "white",
      "brown",
      "red",
    ],
    disabled: false,
    fieldtype: "valuelist",
    datatype: "valuelist",
  },
  values: {
    rowtype: "fieldvalue",
    id: 29,
    name: "fieldvalue_value",
    fieldname: "sample_customer_valuelist",
    value: "yellow",
    notes: "",
    label: "Sample valuelist",
    description: [
      "blue",
      "yellow",
      "white",
      "brown",
      "red",
    ],
    disabled: false,
    fieldtype: "valuelist",
    datatype: "valuelist",
  },
  data: {
    dataset: {}, 
    current: {}, 
    audit: "all",
  },
  className: "dark"
}

export const Selector = Template.bind({});
Selector.args = {
  ...Default.args,
  field: {
    rowtype: "field",
    name: "customer_id",
    label: "Customer Name",
    datatype: "selector",
    empty: true,
    map: {
      seltype: "customer",
      table: "trans",
      fieldname: "customer_id",
      lnktype: "customer",
      transtype: "",
      label_field: "custname",
    },
  },
  values: {
    id: 5,
    transnumber: "DMINV/00001",
    customer_id: 2,
    custname: "First Customer Co.",
  },
  data: {
    dataset: {
      trans: [
        {
          id: 5,
          transnumber: "DMINV/00001",
          customer_id: 2,
          custname: "First Customer Co.",
        }
      ]
    }, 
    current: {}, 
    audit: "all",
  },
}

export const SelectorExtend = Template.bind({});
SelectorExtend.args = {
  ...Default.args,
  field: {
    name: "ref_id",
    label: "Reference",
    datatype: "selector",
    empty: false,
    map: {
      seltype: "employee",
      table: "extend",
      fieldname: "ref_id",
      lnktype: "employee",
      transtype: "",
      label_field: "refnumber",
      extend: true,
    },
  },
  values: {
    id: 12,
    transtype: 63,
    direction: 68,
    transnumber: "DMMOVE/00001",
    ref_transnumber: null,
    crdate: "2021-11-22",
    transdate: "2020-12-05",
    duedate: null,
    customer_id: null,
    employee_id: 4,
    department: null,
  },
  data: {
    dataset: {}, 
    current: {
      extend: {
        seltype: "employee",
        ref_id: 4,
        refnumber: "demo",
        transtype: "",
      }
    }, 
    audit: "all",
  },
}

export const SelectorLnkID = Template.bind({});
SelectorLnkID.args = {
  ...Default.args,
  field: {
    rowtype: "field",
    name: "ref_id_1",
    label: "Payment No.",
    datatype: "selector",
    empty: false,
    map: {
      seltype: "payment",
      table: "invoice_link",
      fieldname: "ref_id_1",
      lnktype: "trans",
      transtype: "",
      lnkid: "trans_id",
      label_field: "transnumber",
    },
  },
  values: {
    curr: "EUR",
    deleted: 0,
    id: 6,
    lslabel: "2020-12-20 | Bank | 849.0",
    lsvalue: "DMPMT/00001 ~ 2 | bank-transfer",
    nervatype_1: 22,
    nervatype_2: 31,
    ref_id_1: 2,
    ref_id_2: 5,
    trans_id: 10,
    transnumber: "DMPMT/00001 ~ 2",
    transtype: "bank",
  },
  data: {
    dataset: {}, 
    current: {
      invoice_link: [
        {
          curr: "EUR",
          deleted: 0,
          id: 6,
          lslabel: "2020-12-20 | Bank | 849.0",
          lsvalue: "DMPMT/00001 ~ 2 | bank-transfer",
          nervatype_1: 22,
          nervatype_2: 31,
          ref_id_1: 2,
          ref_id_2: 5,
          trans_id: 10,
          transnumber: "DMPMT/00001 ~ 2",
          transtype: "bank",
        },
      ]
    }, 
    audit: "all",
  },
}

export const SelectorFieldvalue = Template.bind({});
SelectorFieldvalue.args = {
  ...Default.args,
  field: {
    rowtype: "fieldvalue",
    id: 128,
    name: "fieldvalue_value",
    fieldname: "trans_transitem_link",
    value: "5",
    notes: "",
    label: "Ref.No.",
    description: "DMINV/00001",
    disabled: false,
    fieldtype: "transitem",
    datatype: "selector",
  },
  values: {
    rowtype: "fieldvalue",
    id: 128,
    name: "fieldvalue_value",
    fieldname: "trans_transitem_link",
    value: "5",
    notes: "",
    label: "Ref.No.",
    description: "DMINV/00001",
    disabled: false,
    fieldtype: "transitem",
    datatype: "selector",
  },
  data: {
    dataset: {}, 
    current: {}, 
    audit: "all",
  },
}

export const Select = Template.bind({});
Select.args = {
  ...Default.args,
  field: {
    name: "direction",
    label: "Invoice Type",
    datatype: "select",
    empty: true,
    map: {
      source: "direction",
      value: "id",
      text: "groupvalue",
      label: "invoice",
    },
    disabled: false,
  },
  values: {
    id: 5,
    transtype: 55,
    direction: 68,
    transnumber: "DMINV/00001",
  },
  data: {
    dataset: {
      direction: [
        {
          deleted: 0,
          description: null,
          groupname: "direction",
          groupvalue: "in",
          id: 69,
          inactive: 0,
        },
        {
          deleted: 0,
          description: null,
          groupname: "direction",
          groupvalue: "out",
          id: 68,
          inactive: 0,
        },
      ]
    }, 
    current: {}, 
    audit: "all",
  },
}

export const SelectOptions = Template.bind({});
SelectOptions.args = {
  ...Default.args,
  field: {
    name: "groupname",
    label: "Group Type",
    datatype: "select",
    default: "",
    options: [
      [ "", "" ],
      [ "department", "department" ],
      [ "eventgroup", "eventgroup" ],
      [ "paidtype", "paidtype" ],
      [ "toolgroup", "toolgroup" ],
      [ "rategroup", "rategroup" ],
    ],
    disabled: true,
  },
  values: {
    deleted: 0,
    description: "transfer",
    groupname: "paidtype",
    groupvalue: "transfer",
    id: 123,
    inactive: 0,
  },
  data: {
    dataset: {}, 
    current: {}, 
    audit: "all",
  },
}

export const SelectOptionsLabel = Template.bind({});
SelectOptionsLabel.args = {
  ...Default.args,
  field: {
    name: "seltype",
    label: "Reference type",
    datatype: "select",
    empty: false,
    olabel: "waybill",
    extend: true,
    options: [
      [ "transitem", "transitem" ],
      [ "customer", "customer" ],
      [ "employee", "employee" ],
    ],
  },
  values: {
    id: 12,
    transtype: 63,
    direction: 68,
    transnumber: "DMMOVE/00001",
    ref_transnumber: null,
    crdate: "2021-11-22",
  },
  data: {
    dataset: {}, 
    current: {
      extend: {
        seltype: "employee",
        ref_id: 4,
        refnumber: "demo",
        transtype: "",
      }
    }, 
    audit: "all",
  },
}

export const Empty = Template.bind({});
Empty.args = {
  ...Default.args,
  field: {
    name: "oslabel",
    label: "Orientation / Size",
    datatype: "label",
  },
  values: {},
  data: {
    dataset: {}, 
    current: {}, 
    audit: "all",
  },
}

export const BoolFalse = Template.bind({});
BoolFalse.args = {
  ...Default.args,
  field: {
    rowtype: "fieldvalue",
    id: 129,
    name: "fieldvalue_value",
    fieldname: "6186fca8",
    value: "false",
    notes: "",
    label: "Test description",
    description: "true",
    disabled: false,
    fieldtype: "bool",
    datatype: "bool",
  },
  values: {
    rowtype: "fieldvalue",
    id: 129,
    name: "fieldvalue_value",
    fieldname: "6186fca8",
    value: "false",
    notes: "",
    label: "Test description",
    description: "true",
    disabled: false,
    fieldtype: "bool",
    datatype: "bool",
  },
  data: {
    dataset: {}, 
    current: {}, 
    audit: "all",
  },
}

export const BoolTrue = Template.bind({});
BoolTrue.args = {
  ...Default.args,
  field: {
    name: "paid",
    label: "Paid",
    datatype: "bool",
  },
  values: {
    id: 5,
    transtype: 55,
    direction: 68,
    transnumber: "DMINV/00001",
    paid: 1,
  },
  data: {
    dataset: {}, 
    current: {}, 
    audit: "all",
  },
}

export const BoolTrueDisabled = Template.bind({});
BoolTrueDisabled.args = {
  ...Default.args,
  field: {
    name: "paid",
    label: "Paid",
    datatype: "bool",
  },
  values: {
    id: 5,
    transtype: 55,
    direction: 68,
    transnumber: "DMINV/00001",
    paid: 1,
  },
  data: {
    dataset: {}, 
    current: {}, 
    audit: "readonly",
  },
}

export const DateDisabled = Template.bind({});
DateDisabled.args = {
  ...Default.args,
  field: {
    name: "crdate",
    label: "Creation",
    datatype: "date",
    disabled: true,
  },
  values: {
    id: 5,
    transtype: 55,
    crdate: "2021-11-22",
  },
  data: {
    dataset: {}, 
    current: {}, 
    audit: "all",
  },
}

export const DateTime = Template.bind({});
DateTime.args = {
  ...Default.args,
  field: {
    name: "duedate",
    label: "End Date",
    datatype: "datetime",
    empty: true,
  },
  values: {
    id: 20,
    transtype: 64,
    duedate: "2020-12-02T00:00:00",
  },
  data: {
    dataset: {}, 
    current: {}, 
    audit: "all",
  },
}

export const DateExtend = Template.bind({});
DateExtend.args = {
  ...Default.args,
  field: {
    name: "id",
    label: "Payment Date",
    datatype: "date",
    map: {
      source: "payment",
      value: "trans_id",
      text: "paiddate",
      extend: true,
    },
  },
  values: {
    id: 11,
    transtype: 67,
    direction: 68,
    transnumber: "DMPMT/00002",
  },
  data: {
    dataset: {}, 
    current: {
      extend: {
        id: 4,
        trans_id: 11,
        paiddate: "2020-12-18",
        amount: -488,
        notes: null,
        deleted: 0,
        rid: 4,
      }
    }, 
    audit: "all",
  },
}

export const DateLink = Template.bind({});
DateLink.args = {
  ...Default.args,
  field: {
    name: "trans_id",
    label: "Shipping Date",
    datatype: "date",
    map: {
      source: "trans",
      value: "id",
      text: "transdate",
    },
  },
  values: {
    id: 12,
    item_id: 10,
    trans_id: 14,
  },
  data: {
    dataset: {}, 
    current: {
      trans: [
        {
          id: 14,
          transcast: "normal",
          transdate: "2020-12-10",
          transnumber: "DMDEL/00002",
        },
      ]
    }, 
    audit: "all",
  },
}

export const DateLinkValue = Template.bind({});
DateLinkValue.args = {
  ...Default.args,
  field: {
    name: "trans_testdate",
    label: "Date link",
    datatype: "date",
    map: {
      source: "fieldvalue",
      value: "fieldname",
      text: "value",
    },
    opposite: true
  },
  values: {
    id: 8,
    transtype: 59,
    direction: 68,
    transnumber: "DMWORK/00001",
  },
  data: {
    dataset: {}, 
    current: {
      fieldvalue: [
        {
          deleted: 0,
          fieldname: "trans_testdate",
          id: 76,
          notes: "",
          ref_id: 8,
          value: "2021-11-01",
        },
      ]
    }, 
    audit: "readonly",
  },
}

export const Password = Template.bind({});
Password.args = {
  ...Default.args,
  field: {
    name: "password_1",
    label: "New password",
    datatype: "password",
  },
  values: {
    username: "admin",
    password_1: "password",
    password_2: "",
  },
  data: {
    dataset: {}, 
    current: {}, 
    audit: "all",
  },
}

export const Color = Template.bind({});
Color.args = {
  ...Default.args,
  field: {
    rowtype: "flip",
    name: "color",
    datatype: "color",
    info: "value in hexadecimal (example: #A0522D) or in decimal (example: 10506797)",
  },
  values: {
    title: "BANK STATEMENT",
    color: "#A0522D",
  },
  data: {
    dataset: {}, 
    current: {}, 
    audit: "all",
  },
}

export const Fieldvalue = Template.bind({});
Fieldvalue.args = {
  ...Default.args,
  field: {
    rowtype: "field",
    name: "fieldvalue_value",
    label: "Value",
    datatype: "fieldvalue",
  },
  values: {
    rowtype: "fieldvalue",
    id: 8,
    fieldname: "default_paidtype",
    fieldvalue_value: "transfer",
    fieldvalue_notes: "",
    label: "default paidtype",
    description: "transfer",
    disabled: "false",
    fieldtype: "string",
    datatype: "string",
  },
  data: {
    dataset: {}, 
    current: {}, 
    audit: "all",
  },
}