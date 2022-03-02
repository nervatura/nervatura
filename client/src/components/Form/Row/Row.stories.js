import { Row } from "./Row";
import { getText, store } from 'config/app';

export default {
  title: "Form/Row",
  component: Row
}

const Template = (args) => <Row {...args} />

export const Default = Template.bind({});
Default.args = {
  row: {
    rowtype: "label",
    name: "description",
  },
  values: {
    id: 7,
    description: "Description text",
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

export const Label = Template.bind({});
Label.args = {
  ...Default.args,
  row: {
    rowtype: "label",
    label: "Label text",
  },
  values: {
  }
}

export const FlipStringOn = Template.bind({});
FlipStringOn.args = {
  ...Default.args,
  row: {
    rowtype: "flip",
    name: "title",
    datatype: "string",
    default: "Nervatura Report",
    info: "Info message"
  },
  values: {
    title: "BANK STATEMENT",
  }
}

export const FlipStringOff = Template.bind({});
FlipStringOff.args = {
  ...Default.args,
  row: {
    rowtype: "flip",
    name: "title",
    datatype: "string",
    default: "Nervatura Report",
  },
  values: {}
}

export const FlipTextOn = Template.bind({});
FlipTextOn.args = {
  ...Default.args,
  className: "dark",
  row: {
    rowtype: "flip",
    name: "html",
    datatype: "text",
    default: "",
    info: "The text specified here can include Basic HTML format elements (bold, italic, etc.) Experimental!",
  },
  values: {
    fieldname: "notes",
    html: "<b>Bold text</b><br />",
  }
}

export const FlipTextOff = Template.bind({});
FlipTextOff.args = {
  ...Default.args,
  row: {
    rowtype: "flip",
    name: "html",
    datatype: "text",
    default: "",
  },
  values: {
    fieldname: "notes",
  }
}

export const FlipImageOn = Template.bind({});
FlipImageOn.args = {
  ...Default.args,
  row: {
    rowtype: "flip",
    name: "src",
    datatype: "image",
    info: "an inlined image (jpg/png), encoded in base64 (or a valid databind key)",
  },
  values: {
    src: "logo",
  },
  data: {
    dataset: {
      logo: "data:image/jpg;base64,/9j/4AAQSkZJRgABAQIA7ADsAAD/2wBDAAoHBwgHBgoICAgLCgoLDhgQDg0NDh0VFhEYIx8lJCIfIiEmKzcvJik0KSEiMEExNDk7Pj4+JS5ESUM8SDc9Pjv/2wBDAQoLCw4NDhwQEBw7KCIoOzs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozv/wgARCABAAEADAREAAhEBAxEB/8QAGwABAAIDAQEAAAAAAAAAAAAAAAQGAQMFAgf/xAAYAQEBAQEBAAAAAAAAAAAAAAAAAQIDBP/aAAwDAQACEAMQAAABuYAAAAAAPCQZno3eCBMc+YjyRJnmzndenq6F3VMefg55Abl2W37p68Hznl4sHolXcWYs2+9h12wUPn5MGCLM9C7u/T1ejmzEaZr2eOiTr66WfXfbaAAAAP/EACMQAAICAgEDBQEAAAAAAAAAAAIDAQQABRATITAREhQVIjH/2gAIAQEAAQUC8xEIRXd8hvDLtdWFtl5O2LC2VksMzZNJfSq5sLkkfMh7QrK61jJ7D/eIiSn8oyZkp1aPQeG1yW/prHCd2ypVKyyIgYy1Fks+tCVMpvVMKZOV9YZYtYqDxf/EACERAAEDAwUBAQAAAAAAAAAAAAEAAhEQE0EDEiAwUSEy/9oACAEDAQE/Ae8Gal4CuhXVcciZTBApqPwOEQmiTy/NNMZqW/YUDK3eUY3dV27CtiEWEKCm6fqAjr//xAAbEQACAgMBAAAAAAAAAAAAAAABEQAQEiAwQP/aAAgBAgEBPwHxOZTKOhROo4C1oBZiipdP/8QAKRAAAQMCBAUEAwAAAAAAAAAAAQACEQMhEDFBYRIiMHGRICNRoTJCYv/aAAgBAQAGPwLrS4gDdPe38G2G+MF8n4C5abj3VqQ8qxDewUvcXd0wam5wNFhhoz39EuzOQTGb3wJxgCStHVPpqkmSjWOthi6l4k6L3Ks7MuuGm3gb9nD+BmUALAYNFB0A5ohziah/cqDTJ3F1am7wprco+NVwMEAdP//EACUQAQACAAUEAQUAAAAAAAAAAAEAERAhMUFRMGFxkdEgobHw8f/aAAgBAQABPyHrOTrdVBvNQ/Ld/GLlPyMF90BNn/N4bX69vLQTlXL0KD3YJyybN30VyV5muWcEtvDfB9gIqlW1wMuTQJWt2evyMcuTNWMdz9XDWU0Ms0APK2bd7O3vSUMzZNN+RwoGZ/Chl0KA2wrmyj3JrMv3JVUcNItSXgcXG7fV8QIHbHT/AP/aAAwDAQACAAMAAAAQAAAAAAANEty9DknkCjioltiArqvAAAAA/8QAIBEBAAICAQQDAAAAAAAAAAAAAQARECExMEFRYSCRsf/aAAgBAwEBPxDrKG2X1OM8kx7BL+ItE5MoTCL8Bqt5laZW94BWia9n8iq2yst3y4xK+X1HShRhV6gVowFRNiF35iXE9cV3ACjp/wD/xAAcEQACAgIDAAAAAAAAAAAAAAABEQAQIEEwMVH/2gAIAQIBAT8Q5wXZARKPCSYCFahgoDOXVDuyhUQHcflMttRVCQiMHuALj//EACUQAQABAwMDBQEBAAAAAAAAAAERACExQWGBEHGhMFGRscEg0f/aAAgBAQABPxD1svViA+auiElpLweHV2B5GZspY5aAZ7T8WaeMjpN9AqIC2f8AVbp1cO04rVUIIZuv2IOOi6wWIdY7GN/4ljjsNngDoZc4iXhJN2N/AjmgAgsU4GUT4pyyFVyvR9TwOVe1QZRsCHb6bWDekkDISrUKcbppbvLBx0QCJI2aOgLdIIkBj9ofKRqE5wOFqEJMI/nTYg26G0Kjt+zd8Zo+oglgYOjrVq4PZZdM4vTGblwoNInHvN6ZaDOXkxzFaDikH6rUecg/h5O1BLC32Lq7+n//2Q=="
    }, 
    current: {}, 
    audit: "all",
  },
}

export const FlipImageOff = Template.bind({});
FlipImageOff.args = {
  ...Default.args,
  row: {
    rowtype: "flip",
    name: "src",
    datatype: "image",
  },
  values: {
  }
}

export const FlipChecklistOn = Template.bind({});
FlipChecklistOn.args = {
  ...Default.args,
  row: {
    rowtype: "flip",
    name: "border",
    datatype: "checklist",
    values: [
      "1|All",
      "L|Left",
      "T|Top",
      "R|Right",
      "B|Bottom",
    ],
    info: "Info message"
  },
  values: {
    name: "label",
    width: 40,
    "font-style": "bold",
    value: "labels.lb_statement_no",
    border: "LBT",
    "border-color": 100,
  }
}

export const FlipChecklistOff = Template.bind({});
FlipChecklistOff.args = {
  ...Default.args,
  row: {
    rowtype: "flip",
    name: "border",
    datatype: "checklist",
    values: [
      "1|All",
      "L|Left",
      "T|Top",
      "R|Right",
      "B|Bottom",
    ],
  },
  values: {
    name: "label",
    width: 40,
    "font-style": "bold",
    value: "labels.lb_statement_no",
    "border-color": 100,
  }
}

export const Field = Template.bind({});
Field.args = {
  ...Default.args,
  row: {
    rowtype: "field",
    name: "customer_id",
    label: "Customer Name",
    datatype: "selector",
    empty: false,
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
    customer_id: 2,
    custname: "First Customer Co.",
  },
  data: {
    dataset: {
      trans: [
        {
          custname: "First Customer Co.",
          customer_id: 2,
          id: 5,
        },
      ] 
    }, 
    current: {}, 
    audit: "all",
  },
}

export const Reportfield = Template.bind({});
Reportfield.args = {
  ...Default.args,
  row: {
    id: 1,
    rowtype: "reportfield",
    datatype: "date",
    name: "posdate",
    label: "Date",
    selected: true,
    empty: "false",
    value: "2021-12-14",
  },
  values: {
    id: 1,
    rowtype: "reportfield",
    datatype: "date",
    name: "posdate",
    label: "Date",
    selected: true,
    empty: "false",
    value: "2021-12-14",
  },
}

export const ReportfieldEmpty = Template.bind({});
ReportfieldEmpty.args = {
  ...Default.args,
  row: {
    id: 2,
    rowtype: "reportfield",
    datatype: "string",
    name: "curr",
    label: "Currency",
    selected: false,
    empty: "true",
    value: "",
  },
  values: {
    id: 2,
    rowtype: "reportfield",
    datatype: "string",
    name: "curr",
    label: "Currency",
    selected: false,
    empty: "true",
    value: "",
  },
}

export const Fieldvalue = Template.bind({});
Fieldvalue.args = {
  ...Default.args,
  row: {
    rowtype: "fieldvalue",
    id: 27,
    name: "fieldvalue_value",
    fieldname: "sample_customer_float",
    value: "123.4",
    notes: "",
    label: "Sample float",
    description: "123.4",
    disabled: false,
    fieldtype: "float",
    datatype: "float",
  },
  values: {
    rowtype: "fieldvalue",
    id: 27,
    name: "fieldvalue_value",
    fieldname: "sample_customer_float",
    value: "123.4",
    notes: "",
    label: "Sample float",
    description: "123.4",
    disabled: false,
    fieldtype: "float",
    datatype: "float",
  },
}

export const Col2 = Template.bind({});
Col2.args = {
  ...Default.args,
  row: {
    rowtype: "col2",
    columns: [
      {
        name: "paidtype",
        label: "Payment",
        datatype: "select",
        empty: false,
        map: {
          source: "paidtype",
          value: "id",
          text: "groupvalue",
          label: "paidtype",
        },
      },
      {
        name: "department",
        label: "Department",
        datatype: "select",
        empty: true,
        map: {
          source: "department",
          value: "id",
          text: "groupvalue",
        },
      },
    ],
  },
  values: {
    id: 5,
    department: 138,
    project_id: null,
    place_id: null,
    paidtype: 123,
    curr: "EUR",
  },
  data: {
    dataset: {
      department: [
        {
          deleted: 0,
          description: "Sample logistics department",
          groupname: "department",
          groupvalue: "logistics",
          id: 139,
          inactive: 0,
        },
        {
          deleted: 0,
          description: "Sample production department",
          groupname: "department",
          groupvalue: "production",
          id: 140,
          inactive: 0,
        },
        {
          deleted: 0,
          description: "Sample sales department",
          groupname: "department",
          groupvalue: "sales",
          id: 138,
          inactive: 0,
        },
      ],
      paidtype: [
        {
          deleted: 0,
          description: "cash",
          groupname: "paidtype",
          groupvalue: "cash",
          id: 122,
          inactive: 0,
        },
        {
          deleted: 0,
          description: "credit card",
          groupname: "paidtype",
          groupvalue: "credit_card",
          id: 124,
          inactive: 0,
        },
        {
          deleted: 0,
          description: "transfer",
          groupname: "paidtype",
          groupvalue: "transfer",
          id: 123,
          inactive: 0,
        },
      ]
    }, 
    current: {}, 
    audit: "all",
  },
}

export const Col3 = Template.bind({});
Col3.args = {
  ...Default.args,
  row: {
    rowtype: "col3",
    columns: [
      {
        name: "crdate",
        label: "Creation",
        datatype: "date",
        disabled: true,
      },
      {
        name: "transdate",
        label: "Invoice Date",
        datatype: "date",
      },
      {
        name: "duedate",
        label: "Due Date",
        datatype: "date",
      },
    ],
  },
  values: {
    id: 5,
    crdate: "2021-11-22",
    transdate: "2020-12-10",
    duedate: "2020-12-20T00:00:00",
  },
  data: {
    dataset: {}, 
    current: {}, 
    audit: "all",
  },
}

export const Col4 = Template.bind({});
Col4.args = {
  ...Default.args,
  row: {
    rowtype: "col4",
    columns: [
      {
        name: "curr",
        label: "Currency",
        datatype: "select",
        empty: true,
        map: {
          source: "currency",
          value: "curr",
          text: "curr",
        },
      },
      {
        name: "acrate",
        label: "Acc.Rate",
        datatype: "float",
        default: 0,
      },
      {
        name: "paid",
        label: "Paid",
        datatype: "flip",
      },
      {
        name: "closed",
        label: "Closed",
        datatype: "flip",
      },
    ],
  },
  values: {
    id: 5,
    curr: "EUR",
    notax: 0,
    paid: 0,
    acrate: 0,
    closed: 0,
    deleted: 0,
  },
  data: {
    dataset: {
      currency: [
        {
          cround: 0,
          curr: "EUR",
          defrate: 0,
          description: "euro",
          digit: 2,
          id: 1,
        },
        {
          cround: 0,
          curr: "USD",
          defrate: 0,
          description: "dollar",
          digit: 2,
          id: 2,
        },
      ]
    }, 
    current: {}, 
    audit: "all",
  },
}

export const Missing = Template.bind({});
Missing.args = {
  ...Default.args,
  row: {
    rowtype: "missing",
  },
}