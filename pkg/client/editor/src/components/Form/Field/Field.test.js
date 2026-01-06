import { fixture, expect } from '@open-wc/testing';
import sinon from 'sinon'

import './form-field.js';
import { Template, Default, StringMapExtend, StringMapLinkID, TextNull, NotesValue, Float, Button,
  Link, ValueList, Selector, Select, Empty, BoolFalse, BoolTrue, BoolTrueDisabled,
  DateDisabled, DateTime, Password, Color, Integer, FloatLinkValue, SelectOptions,
  SelectOptionsLabel, DateExtend, DateLink, Fieldvalue, StringMapLinkValue, DateLinkValue,
  IntegerLinkID, SelectorExtend, SelectorLnkID, SelectorFieldvalue } from  './Field.stories.js';

describe('Field', () => {
  afterEach(() => {
    // Restore the default sandbox here
    sinon.restore();
  });

  it('renders in the Default state', async () => {
    const onEdit = sinon.spy()
    const element = await fixture(Template({
      ...Default.args, onEdit
    }));
    const testField = element.querySelector('#test_field');
    expect(testField).to.exist;

    const input = testField.shadowRoot.querySelector('#test_field');
    input._onInput({ target: { value: "change" } })
    expect(input.value).to.equal("change");
    sinon.assert.calledOnce(onEdit);
  })

  it('renders in the StringMapExtend state', async () => {
    const element = await fixture(Template({
      ...StringMapExtend.args
    }));
    const testField = element.querySelector('#test_field');
    expect(testField).to.exist;
  })

  it('renders in the StringMapLinkID state', async () => {
    const element = await fixture(Template({
      ...StringMapLinkID.args
    }));
    const testField = element.querySelector('#test_field');
    expect(testField).to.exist;
  })

  it('renders in the TextNull state', async () => {
    const onEdit = sinon.spy()
    let element = await fixture(Template({
      ...TextNull.args, onEdit
    }));
    let testField = element.querySelector('#test_field');
    expect(testField).to.exist;

    testField._onTextInput({ target: { value: "change" } })
    sinon.assert.calledOnce(onEdit);

    element = await fixture(Template({
      ...TextNull.args,
      field: {
        rowtype: "field",
        name: "notes",
        label: "Comment",
        datatype: "text",
        default: "notes"
      }
    }));
    testField = element.querySelector('#test_field');
    expect(testField).to.exist;
    
  })

  it('renders in the NotesValue state', async () => {
    const element = await fixture(Template({
      ...NotesValue.args
    }));
    const testField = element.querySelector('#test_field');
    expect(testField).to.exist;
  })

  it('renders in the Float state', async () => {
    const onEdit = sinon.spy()
    const element = await fixture(Template({
      ...Float.args, onEdit
    }));
    const testField = element.querySelector('#test_field');
    expect(testField).to.exist;

    const input = testField.shadowRoot.querySelector('#test_field');
    input._onBlur()
    expect(input.value).to.equal(0);
    sinon.assert.callCount(onEdit, 1);

    input._onInput({ target: { value: "100", valueAsNumber: 100 } })
    expect(input.value).to.equal(100);
    sinon.assert.callCount(onEdit, 2);

    input._onBlur()
    expect(input.value).to.equal(100);
    sinon.assert.callCount(onEdit, 3);
  })

  it('renders in the Button state', async () => {
    const onEdit = sinon.spy()
    const element = await fixture(Template({
      ...Button.args, onEdit
    }));
    const testField = element.querySelector('#test_field');
    expect(testField).to.exist;

    const button = testField.shadowRoot.querySelector('#test_field');
    button._onClick({
      stopPropagation: sinon.spy()
    })
    sinon.assert.calledOnce(onEdit);

    const elementDisabled = await fixture(Template({
      ...Button.args,
      field: {
        ...Button.args.field,
        disabled: true,
        focus: false,
        title: null,
        icon: null
      }
    }));
    const testFieldDisabled = elementDisabled.querySelector('#test_field');
    expect(testFieldDisabled).to.exist;
  })

  it('renders in the Link state', async () => {
    const onEvent = sinon.spy()
    let element = await fixture(Template({
      ...Link.args, onEvent
    }));
    let testField = element.querySelector('#test_field');
    expect(testField).to.exist;

    const link = testField.shadowRoot.querySelector('#link_trans_id');
    link.click()
    sinon.assert.calledOnce(onEvent);

    element = await fixture(Template({
      ...Link.args,
      values: {
        id: null,
        transtype: 55,
        direction: 68,
        transnumber: "DMINV/00001",
        ref_transnumber: "DMORD/00003",
      }
    }));
    testField = element.querySelector('#test_field');
    expect(testField).to.exist;
    

  })

  it('renders in the ValueList state', async () => {
    const onEdit = sinon.spy()
    const element = await fixture(Template({
      ...ValueList.args, onEdit
    }));
    const testField = element.querySelector('#test_field');
    expect(testField).to.exist;

    const select = testField.shadowRoot.querySelector('#test_field')
    select._onInput({ target: { value: "blue" } })
    sinon.assert.calledOnce(onEdit);
    expect(select.value).to.equal('blue');
  })

  it('renders in the Selector state', async () => {
    const onEvent = sinon.spy()
    const onEdit = sinon.spy()
    const onSelector = (type, filter, setSelector) => {
      setSelector({
        custname: "Second Customer Name",
        deleted: 0,
        id: "trans/invoice/6",
        label: "DMINV/00002",
        transnumber: "DMINV/00002",
        transtype: "invoice-out",
        trans_id: 123
      }, "DMINV/0000")
    }
    const element = await fixture(Template({
      ...Selector.args, onEvent, onEdit, onSelector
    }));
    const testField = element.querySelector('#test_field');
    expect(testField).to.exist;

    const selLink = testField.shadowRoot.querySelector('#sel_link_customer_id');
    selLink.click()
    sinon.assert.calledOnce(onEvent);

    const selDelete = testField.shadowRoot.querySelector('#sel_delete_customer_id');
    selDelete.click()
    sinon.assert.calledOnce(onEdit);

    const selShow = testField.shadowRoot.querySelector('#sel_show_customer_id');
    selShow.click()

    const elementDisabled = await fixture(Template({
      ...Selector.args, onEvent,
      field: {
        ...Selector.args.field,
        disabled: true,
      },
      values: {
        ...Selector.args.values,
        custname: null
      }
    }));
    const testFieldDisabled = elementDisabled.querySelector('#test_field');
    expect(testFieldDisabled).to.exist;

    const selLinkDisabled = testField.shadowRoot.querySelector('#sel_link_customer_id');
    selLinkDisabled.click()
    sinon.assert.calledTwice(onEvent);

    const elementReftable = await fixture(Template({
      ...Selector.args, onEvent,
      data: {
        ...Selector.args.data,
        dataset: {
          trans: []
        }
      },
    }));
    const testFieldReftable = elementReftable.querySelector('#test_field');
    expect(testFieldReftable).to.exist;

  })

  it('renders in the SelectorExtend state', async () => {
    const onSelector1 = (type, filter, setSelector) => {
      setSelector({
        id: "customer/6",
        item: {
          lslabel: ""
        }
      }, "DMINV/0000")
    }
    const onSelector2 = (type, filter, setSelector) => {
      setSelector({
        custname: "Second Customer Name",
        deleted: 0,
        id: "trans/invoice/6",
        label: "DMINV/00002",
        transnumber: "DMINV/00002",
        transtype: "invoice-out",
      }, "DMINV/0000")
    }
    let element = await fixture(Template({
      ...SelectorExtend.args, onSelector: onSelector1
    }));
    let testField = element.querySelector('#test_field');
    expect(testField).to.exist;

    const selShow = testField.shadowRoot.querySelector('#sel_show_ref_id');
    selShow.click()


    testField.onSelector = onSelector2
    testField.field = {
      name: "ref_id",
      label: "Reference",
      datatype: "selector",
      empty: false,
      map: {
        seltype: "transitem",
        table: "extend",
        fieldname: "ref_id",
        lnktype: "trans",
        transtype: "invoice",
        label_field: "refnumber",
        extend: true,
      },
    }
    testField.values = {
      ...testField.values,
      ref_id: 5
    }
    testField.data = {
      ...testField.data,
      current: {
        extend: {
          seltype: "transitem",
          ref_id: 5,
          refnumber: "DMINV/00001",
          transtype: "invoice",
        }
      }
    }
    selShow.click()

    element = await fixture(Template({
      ...SelectorExtend.args,
      field: {
        name: "ref_id",
        label: "Reference",
        datatype: "selector",
        empty: false,
        map: {
          seltype: "transitem",
          table: "extend",
          fieldname: "ref_id",
          lnktype: "trans",
          transtype: "invoice",
          label_field: "refnumber",
          extend: true,
        },
      },
      values: {
        ...SelectorExtend.args.values,
        ref_id: 5
      },
      data: {
        ...SelectorExtend.args.data,
        current: {
          extend: {
            seltype: "transitem",
            ref_id: 5,
            refnumber: "DMINV/00001",
            transtype: "invoice",
          }
        }
      }
    }));
    testField = element.querySelector('#test_field');
    expect(testField).to.exist;

    element = await fixture(Template({
      ...SelectorExtend.args,
      field: {
        name: "ref_id",
        label: "Reference",
        datatype: "selector",
        empty: false,
        map: {
          seltype: "transitem",
          table: "extend",
          fieldname: "ref_id",
          lnktype: "trans",
          transtype: "invoice",
          label_field: "refnumber",
          extend: true,
        },
      },
      data: {
        ...SelectorExtend.args.data,
        current: {
          extend: {
            seltype: "transitem",
            refnumber: "DMINV/00001",
            transtype: "invoice",
          }
        }
      }
    }));
    testField = element.querySelector('#test_field');
    expect(testField).to.exist;

    element = await fixture(Template({
      ...SelectorExtend.args,
      field: {
        name: "ref_id",
        label: "Reference",
        datatype: "selector",
        empty: false,
        map: {
          seltype: "transitem",
          table: "extend",
          fieldname: "ref_id",
          lnktype: "trans",
          transtype: "invoice",
          label_field: "refnumber",
          extend: true,
        },
      },
      data: {
        ...SelectorExtend.args.data,
        current: {
          extend: {
            seltype: "transitem",
            transtype: "invoice",
          }
        }
      }
    }));
    testField = element.querySelector('#test_field');
    expect(testField).to.exist;
    
  })

  it('renders in the SelectorLnkID state', async () => {
    const element = await fixture(Template({
      ...SelectorLnkID.args
    }));
    const testField = element.querySelector('#test_field');
    expect(testField).to.exist;
  })

  it('renders in the SelectorFieldvalue state', async () => {
    let element = await fixture(Template({
      ...SelectorFieldvalue.args
    }));
    let testField = element.querySelector('#test_field');
    expect(testField).to.exist;

    element = await fixture(Template({
      ...SelectorFieldvalue.args,
      field: {
        ...SelectorFieldvalue.args.field,
        value: null,
        description: null
      },
      values: {
        ...SelectorFieldvalue.args.values,
        value: null,
        description: null
      }
    }));
    testField = element.querySelector('#test_field');
    expect(testField).to.exist;
  })

  it('renders in the Select state', async () => {
    const onEdit = sinon.spy()
    const element = await fixture(Template({
      ...Select.args, onEdit
    }));
    const testField = element.querySelector('#test_field');
    expect(testField).to.exist;

    const selectField = testField.shadowRoot.querySelector('#test_field')
    selectField._onInput({ target: { value: "69" } })
    sinon.assert.calledOnce(onEdit);
    expect(selectField.value).to.equal("69");

    selectField._onInput({ target: { value: "abc" } })
    sinon.assert.calledTwice(onEdit);
    expect(selectField.value).to.equal('abc');
  })

  it('renders in the Empty state', async () => {
    const element = await fixture(Template({
      ...Empty.args
    }));
    const testField = element.querySelector('#test_field');
    expect(testField).to.exist;
  })

  it('renders in the BoolFalse state', async () => {
    const onEdit = sinon.spy()
    let element = await fixture(Template({
      ...BoolFalse.args, onEdit
    }));
    let testField = element.querySelector('#test_field');
    expect(testField).to.exist;

    let link = testField.shadowRoot.querySelector('#test_field');
    link.click()
    sinon.assert.calledOnce(onEdit);

    element = await fixture(Template({
      ...BoolFalse.args, onEdit,
      values: {
        ...BoolFalse.args.values,
        value: "true"
      }
    }));
    testField = element.querySelector('#test_field');
    expect(testField).to.exist;

    link = testField.shadowRoot.querySelector('#test_field');
    link.click()
    sinon.assert.calledTwice(onEdit);
  })

  it('renders in the BoolTrue state', async () => {
    const onEdit = sinon.spy()
    let element = await fixture(Template({
      ...BoolTrue.args, onEdit
    }));
    let testField = element.querySelector('#test_field');
    expect(testField).to.exist;

    let link = testField.shadowRoot.querySelector('#test_field');
    link.click()
    sinon.assert.calledOnce(onEdit);

    element = await fixture(Template({
      ...BoolTrue.args, onEdit,
      values: {
        ...BoolTrue.args.values,
        paid: 0
      }
    }));
    testField = element.querySelector('#test_field');
    expect(testField).to.exist;

    link = testField.shadowRoot.querySelector('#test_field');
    link.click()
    sinon.assert.calledTwice(onEdit);
  })

  it('renders in the BoolTrueDisabled state', async () => {
    const onEdit = sinon.spy()
    const element = await fixture(Template({
      ...BoolTrueDisabled.args, onEdit
    }));
    const testField = element.querySelector('#test_field');
    expect(testField).to.exist;

    const link = testField.shadowRoot.querySelector('#test_field');
    link.click()
    sinon.assert.callCount(onEdit, 0);

    testField.values = {
      ...BoolTrueDisabled.args.values,
      paid: 0
    }
  })

  it('renders in the DateDisabled state', async () => {
    const element = await fixture(Template({
      ...DateDisabled.args
    }));
    const testField = element.querySelector('#test_field');
    expect(testField).to.exist;
  })

  it('renders in the DateTime state', async () => {
    const onEdit = sinon.spy()
    const element = await fixture(Template({
      ...DateTime.args, onEdit
    }));
    const testField = element.querySelector('#test_field');
    expect(testField).to.exist;

    const input = testField.shadowRoot.querySelector('#test_field');
    input._onInput({ target: { value: "2021-12-24T11:34" } })
    expect(input.value).to.equal("2021-12-24T11:34:00");
    sinon.assert.callCount(onEdit, 1);
  })

  it('renders in the Password state', async () => {
    const onEdit = sinon.spy()
    let element = await fixture(Template({
      ...Password.args, onEdit
    }));
    let testField = element.querySelector('#test_field');
    expect(testField).to.exist;

    const input = testField.shadowRoot.querySelector('#test_field');
    input._onInput({ target: { value: "123" } })
    expect(input.value).to.equal("123");
    sinon.assert.callCount(onEdit, 1);

    element = await fixture(Template({
      ...Password.args,
      values: {
        password_1: null,
      },
      data: {
        dataset: {}, 
        current: {}, 
        audit: "readonly",
      }
    }));
    testField = element.querySelector('#test_field');
    expect(testField).to.exist;
  })

  it('renders in the Color state', async () => {
    const element = await fixture(Template({
      ...Color.args
    }));
    const testField = element.querySelector('#test_field');
    expect(testField).to.exist;
  })

  it('renders in the Integer state', async () => {
    const element = await fixture(Template({
      ...Integer.args
    }));
    const testField = element.querySelector('#test_field');
    expect(testField).to.exist;

    const input = testField.shadowRoot.querySelector('#test_field');
    input._onInput({ target: { value: "12", valueAsNumber: 12 } })
    expect(input.value).to.equal(12);

    const elementPercent = await fixture(Template({
      ...Integer.args,
      field: {
        ...Integer.args.field,
        datatype: "percent",
        map: {
          source: "payment",
          value: "trans_id",
          extend: true,
        }
      }
    }));
    const testFieldPercent = elementPercent.querySelector('#test_field');
    expect(testFieldPercent).to.exist;
  })

  it('renders in the FloatLinkValue state', async () => {
    let element = await fixture(Template({
      ...FloatLinkValue.args
    }));
    let testField = element.querySelector('#test_field');
    expect(testField).to.exist;

    const input = testField.shadowRoot.querySelector('#test_field');
    input._onInput({ target: { value: "100", valueAsNumber: 100 } })
    expect(input.value).to.equal(100);

    input._onBlur()
    expect(input.value).to.equal(100);

    element = await fixture(Template({
      ...FloatLinkValue.args,
      data: {
        dataset: { 
          fieldvalue: [
            { deleted: 0, fieldname: "trans_wsdistance", id: 76, notes: "", ref_id: 8, value: "200.0" },
          ]}, 
        current: {}, 
        audit: "all"
      }
    }));
    testField = element.querySelector('#test_field');
    expect(testField).to.exist;
  })

  it('renders in the SelectOptions state', async () => {
    const element = await fixture(Template({
      ...SelectOptions.args
    }));
    const testField = element.querySelector('#test_field');
    expect(testField).to.exist;
  })

  it('renders in the SelectOptionsLabel state', async () => {
    let element = await fixture(Template({
      ...SelectOptionsLabel.args
    }));
    let testField = element.querySelector('#test_field');
    expect(testField).to.exist;

    element = await fixture(Template({
      ...SelectOptionsLabel.args,
      data: {
        dataset: {}, 
        current: {
          extend: {
            ref_id: 4,
            refnumber: "demo",
            transtype: "",
          }
        }, 
        audit: "all",
      }
    }));
    testField = element.querySelector('#test_field');
    expect(testField).to.exist;

  })

  it('renders in the DateExtend state', async () => {
    const element = await fixture(Template({
      ...DateExtend.args
    }));
    const testField = element.querySelector('#test_field');
    expect(testField).to.exist;
  })

  it('renders in the DateLink state', async () => {
    const element = await fixture(Template({
      ...DateLink.args
    }));
    const testField = element.querySelector('#test_field');
    expect(testField).to.exist;
  })

  it('renders in the Fieldvalue state', async () => {
    const element = await fixture(Template({
      ...Fieldvalue.args
    }));
    const testField = element.querySelector('#test_field');
    expect(testField).to.exist;
  })

  it('renders in the StringMapLinkValue state', async () => {
    const element = await fixture(Template({
      ...StringMapLinkValue.args
    }));
    const testField = element.querySelector('#test_field');
    expect(testField).to.exist;
  })

  it('renders in the DateLinkValue state', async () => {
    const element = await fixture(Template({
      ...DateLinkValue.args
    }));
    const testField = element.querySelector('#test_field');
    expect(testField).to.exist;
  })

  it('renders in the IntegerLinkID state', async () => {
    const element = await fixture(Template({
      ...IntegerLinkID.args
    }));
    const testField = element.querySelector('#test_field');
    expect(testField).to.exist;
  })

})