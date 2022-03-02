import { render, fireEvent, queryByAttribute } from '@testing-library/react';
import '@testing-library/jest-dom/extend-expect';
import update from 'immutability-helper';

import { Default, StringMapExtend, StringMapLinkID, TextNull, NotesValue, Float, Button,
  Link, ValueList, Selector, Select, Empty, BoolFalse, BoolTrue, BoolTrueDisabled,
  DateDisabled, DateTime, Password, Color, Integer, FloatLinkValue, SelectOptions,
  SelectOptionsLabel, DateExtend, DateLink, Fieldvalue, StringMapLinkValue, DateLinkValue,
  IntegerLinkID, SelectorExtend, SelectorLnkID, SelectorFieldvalue } from './Field.stories';

const getById = queryByAttribute.bind(null, 'id');

it('renders in the Default state', () => {
  const onEdit = jest.fn()

  const { container } = render(
    <Default {...Default.args} id="test_input" onEdit={onEdit} />
  );
  expect(getById(container, 'test_input')).toBeDefined();

  const test_input = getById(container, 'test_input')

  fireEvent.change(test_input, {target: {value: "change"}})
  expect(onEdit).toHaveBeenCalledTimes(1);

});

it('renders in the StringMapExtend state', () => {

  const { container } = render(
    <StringMapExtend {...StringMapExtend.args} id="test_input" />
  );
  expect(getById(container, 'test_input')).toBeDefined();

});

it('renders in the StringMapLinkID state', () => {

  const { container } = render(
    <StringMapLinkID {...StringMapLinkID.args} id="test_input" />
  );
  expect(getById(container, 'test_input')).toBeDefined();

  render(
    <StringMapLinkID {...StringMapLinkID.args} field={
      update(StringMapLinkID.args.field, {$merge: {
        map: {
          source: "groups",
          value: "id",
          text: "groupvalue",
        }
      }})
    } />
  );

  render(
    <StringMapLinkID {...StringMapLinkID.args} 
      values={{
        id: 14,
        transtype: 61,
        direction: 69,
      }} />
  );

});

it('renders in the StringMapLinkValue state', () => {

  const { container } = render(
    <StringMapLinkValue {...StringMapLinkValue.args} id="test_input" />
  );
  expect(getById(container, 'test_input')).toBeDefined();

});

it('renders in the TextNull state', () => {
  const onEdit = jest.fn()

  const { container, rerender } = render(
    <TextNull {...TextNull.args} id="test_input" onEdit={onEdit} />
  );
  expect(getById(container, 'test_input')).toBeDefined();

  const test_input = getById(container, 'test_input')

  fireEvent.change(test_input, {target: {value: "change"}})
  expect(onEdit).toHaveBeenCalledTimes(1);

  rerender(
    <TextNull {...TextNull.args} 
      field={{
        rowtype: "field",
        name: "notes",
        label: "Comment",
        datatype: "text",
        default: "notes"
      }} />
  )

});

it('renders in the NotesValue state', () => {

  const { container } = render(
    <NotesValue {...NotesValue.args} id="test_input" />
  );
  expect(getById(container, 'test_input')).toBeDefined();

});

it('renders in the Float state', () => {
  const onEdit = jest.fn()

  const { container } = render(
    <Float {...Float.args} id="test_input" onEdit={onEdit} />
  );
  expect(getById(container, 'test_input')).toBeDefined();
 
  const test_input = getById(container, 'test_input')

  fireEvent.blur(test_input, {target: {value: "100"}})
  expect(onEdit).toHaveBeenCalledTimes(1);

  fireEvent.change(test_input, {target: {value: "12"}})
  expect(onEdit).toHaveBeenCalledTimes(2);

});

it('renders in the FloatLinkValue state', () => {

  const { container } = render(
    <FloatLinkValue {...FloatLinkValue.args} id="test_input" />
  );
  expect(getById(container, 'test_input')).toBeDefined();

  render(
    <FloatLinkValue {...FloatLinkValue.args} 
      data={{
      dataset: { 
        fieldvalue: [
          { deleted: 0, fieldname: "trans_wsdistance", id: 76, notes: "", ref_id: 8, value: "200.0" },
        ]}, 
      current: {}, 
      audit: "all"}} />
  );

  render(
    <FloatLinkValue {...FloatLinkValue.args} 
      data={{
      dataset: {}, 
      current: {
        fieldvalue: []
      }, 
      audit: "all"}} />
  );

});

it('renders in the Integer state', () => {
  const onEdit = jest.fn()

  const { container, rerender } = render(
    <Integer {...Integer.args} id="test_input" onEdit={onEdit} />
  );
  expect(getById(container, 'test_input')).toBeDefined();
 
  const test_input = getById(container, 'test_input')

  fireEvent.change(test_input, {target: {value: "12"}})
  expect(onEdit).toHaveBeenCalledTimes(1);

  fireEvent.blur(test_input, {target: {value: "100"}})
  expect(onEdit).toHaveBeenCalledTimes(2);

  rerender(
    <Integer {...Integer.args} id="test_input" onEdit={onEdit}
      field={
        update(Integer.args.field, {$merge: {
          datatype: "percent",
          map: {
            source: "payment",
            value: "trans_id",
            extend: true,
          }
        }})}
    />
  )

  fireEvent.change(test_input, {target: {value: "22"}})
  expect(onEdit).toHaveBeenCalledTimes(3);

});

it('renders in the IntegerLinkID state', () => {

  const { container } = render(
    <IntegerLinkID {...IntegerLinkID.args} id="test_input" />
  );
  expect(getById(container, 'test_input')).toBeDefined();

});

it('renders in the Button state', () => {
  const onEdit = jest.fn()

  const { container } = render(
    <Button {...Button.args} id="test_input" onEdit={onEdit} />
  );
  expect(getById(container, 'test_input')).toBeDefined();
 
  const test_input = getById(container, 'test_input')

  fireEvent.click(test_input)
  expect(onEdit).toHaveBeenCalledTimes(1);

  render(
    <Button {...ValueList.args} 
      field={update(Button.args.field, {$merge: {
        disabled: true,
        focus: false,
        title: null,
        icon: null
      }})}/>
  );

});

it('renders in the Link state', () => {
  const onEvent = jest.fn()

  const { container } = render(
    <Link {...Link.args} id="test_input" onEvent={onEvent} />
  );
  expect(getById(container, 'test_input')).toBeDefined();
 
  const test_input = getById(container, 'link_trans_id')

  fireEvent.click(test_input)
  expect(onEvent).toHaveBeenCalledTimes(1);

  render(
    <Link {...Link.args} 
      data={{
        dataset: {
          translink: []
        }, 
        current: {}, 
        audit: "readonly",
      }} />
  )

  render(
    <Link {...Link.args} field={
      update(Link.args.field, {$merge: {
        map: {
          source: "translink",
          value: "ref_id_1",
          text: "ref",
          label_field: undefined,
          lnktype: "trans",
          transtype: "order",
        }
      }})
    } />
  )

  render(
    <Link {...Link.args} field={
      update(Link.args.field, {$merge: {
        map: {
          source: "translink",
          value: "ref_id_1",
          text: "ref",
          label_field: "",
          lnktype: "trans",
          transtype: "order",
        }
      }})
    } />
  )

  render(
    <Link {...Link.args} values={{
      id: null,
      transtype: 55,
      direction: 68,
      transnumber: "DMINV/00001",
      ref_transnumber: "DMORD/00003",
    }} />
  )

});

it('renders in the ValueList state', () => {
  const onEdit = jest.fn()

  const { container } = render(
    <ValueList {...ValueList.args} id="test_input" onEdit={onEdit} />
  );
  expect(getById(container, 'test_input')).toBeDefined();

  const test_input = getById(container, 'test_input')

  fireEvent.change(test_input, {target: {value: "blue"}})
  expect(onEdit).toHaveBeenCalledTimes(1);

  render(
    <ValueList {...ValueList.args} 
      field={update(ValueList.args.field, {$merge: {
          disabled: true
      }})}/>
  );

});

it('renders in the Selector state', () => {
  const onEvent = jest.fn()
  const onEdit = jest.fn()

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

  const { container } = render(
    <Selector {...Selector.args} id="test_input" 
      onEvent={onEvent} onSelector={onSelector} onEdit={onEdit} />
  );
  expect(getById(container, 'test_input')).toBeDefined();
 
  const test_input = getById(container, 'sel_link_customer_id')
  fireEvent.click(test_input)
  expect(onEvent).toHaveBeenCalledTimes(1);

  const sel_delete = getById(container, 'sel_delete_customer_id')
  fireEvent.click(sel_delete)
  expect(onEdit).toHaveBeenCalledTimes(1);

  const sel_show = getById(container, 'sel_show_customer_id')
  fireEvent.click(sel_show)

  render(
    <Selector {...Selector.args} field={
      update(Selector.args.field, {$merge: {
        disabled: true
      }})
    } />
  )

  render(
    <Selector {...Selector.args} values={
      update(Selector.args.values, {$merge: {
        custname: null
      }})
    } />
  )

  render(
    <Selector {...Selector.args} data={
      update(Selector.args.data, {$merge: {
        dataset: {
          trans: []
        }
      }})
    } />
  )

});

it('renders in the SelectorExtend state', () => {
  const onEdit = jest.fn()
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

  const { container, rerender } = render(
    <SelectorExtend {...SelectorExtend.args} id="test_input" 
      onEdit={onEdit}  onSelector={onSelector1} />
  );
  expect(getById(container, 'test_input')).toBeDefined();

  const sel_show = getById(container, 'sel_show_ref_id')
  fireEvent.click(sel_show)

  rerender(
    <SelectorExtend {...SelectorExtend.args} 
      onEdit={onEdit} onSelector={onSelector2}
      field={{
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
      }}
      values={
        update(SelectorExtend.args.values, {$merge: {
          ref_id: 5
        }})
      }
      data={
        update(SelectorExtend.args.data, {$merge: {
          current: {
            extend: {
              seltype: "transitem",
              ref_id: 5,
              refnumber: "DMINV/00001",
              transtype: "invoice",
            }
          }
        }})
      } />
  );
  
  fireEvent.click(sel_show)

  render(
    <SelectorExtend {...SelectorExtend.args} 
      field={{
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
      }}
      values={
        update(SelectorExtend.args.values, {$merge: {
          ref_id: 5
        }})
      }
      data={
        update(SelectorExtend.args.data, {$merge: {
          current: {
            extend: {
              seltype: "transitem",
              ref_id: 5,
              refnumber: "DMINV/00001",
              transtype: "invoice",
            }
          }
        }})
      } />
  );

  render(
    <SelectorExtend {...SelectorExtend.args} 
      field={{
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
      }}
      data={
        update(SelectorExtend.args.data, {$merge: {
          current: {
            extend: {
              seltype: "transitem",
              refnumber: "DMINV/00001",
              transtype: "invoice",
            }
          }
        }})
      } />
  );

  render(
    <SelectorExtend {...SelectorExtend.args} 
      field={{
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
      }}
      data={
        update(SelectorExtend.args.data, {$merge: {
          current: {
            extend: {
              seltype: "transitem",
              transtype: "invoice",
            }
          }
        }})
      } />
  );

});

it('renders in the SelectorLnkID state', () => {

  const { container } = render(
    <SelectorLnkID {...SelectorLnkID.args} id="test_input" />
  );
  expect(getById(container, 'test_input')).toBeDefined();

});

it('renders in the SelectorFieldvalue state', () => {
  const { container } = render(
    <SelectorFieldvalue {...SelectorFieldvalue.args} id="test_input" />
  );
  expect(getById(container, 'test_input')).toBeDefined();

  render(
    <SelectorFieldvalue {...SelectorFieldvalue.args} id="test_input"
      field={
        update(SelectorFieldvalue.args.field, {$merge: {
          value: null,
          description: null
        }})}
      values={
        update(SelectorFieldvalue.args.values, {$merge: {
          value: null,
          description: null
        }})
    } />
  )

});

it('renders in the Select state', () => {
  const onEdit = jest.fn()

  const { container } = render(
    <Select {...Select.args} id="test_input" onEdit={onEdit} />
  );
  expect(getById(container, 'test_input')).toBeDefined();

  const test_input = getById(container, 'test_input')

  fireEvent.change(test_input, {target: {value: "69"}})
  expect(onEdit).toHaveBeenCalledTimes(1);

  fireEvent.change(test_input, {target: {value: "abc"}})
  expect(onEdit).toHaveBeenCalledTimes(2);

  render(
    <Select {...Select.args} field={
      update(Select.args.field, {$merge: {
        map: {
          source: "direction",
          value: "id",
          text: "groupvalue",
          label: undefined,
        }
      }})
    } />
  )

});

it('renders in the SelectOptions state', () => {
  const { container } = render(
    <SelectOptions {...SelectOptions.args} id="test_input" />
  );
  expect(getById(container, 'test_input')).toBeDefined();

});

it('renders in the SelectOptionsLabel state', () => {
  const { container } = render(
    <SelectOptionsLabel {...SelectOptionsLabel.args} id="test_input" />
  );
  expect(getById(container, 'test_input')).toBeDefined();

  render(
    <SelectOptionsLabel {...SelectOptionsLabel.args} 
      data={{
        dataset: {}, 
        current: {
          extend: {
            ref_id: 4,
            refnumber: "demo",
            transtype: "",
          }
        }, 
        audit: "all",
      }} />
  );

});

it('renders in the Empty state', () => {

  const { container } = render(
    <Empty {...Empty.args} id="test_input" />
  );
  expect(getById(container, 'test_input')).toBeDefined();
});

it('renders in the BoolFalse state', () => {
  const onEdit = jest.fn()

  const { container, rerender } = render(
    <BoolFalse {...BoolFalse.args} id="test_input" onEdit={onEdit} />
  );
  expect(getById(container, 'test_input')).toBeDefined();
 
  let test_input = getById(container, 'test_input')

  fireEvent.click(test_input)
  expect(onEdit).toHaveBeenCalledTimes(1);

  rerender(
    <BoolFalse {...BoolFalse.args} id="test_input" onEdit={onEdit} 
      values={update(BoolFalse.args.field, {$merge: {
        value: "true"
      }})} />
  );
  fireEvent.click(test_input)
  expect(onEdit).toHaveBeenCalledTimes(2);

});

it('renders in the BoolTrue state', () => {
  const onEdit = jest.fn()

  const { container, rerender } = render(
    <BoolTrue {...BoolTrue.args} id="test_input" onEdit={onEdit} />
  );
  expect(getById(container, 'test_input')).toBeDefined();
 
  const test_input = getById(container, 'test_input')

  fireEvent.click(test_input)
  expect(onEdit).toHaveBeenCalledTimes(1);

  rerender(
    <BoolTrue {...BoolTrue.args} id="test_input" onEdit={onEdit} 
      values={update(BoolTrue.args.field, {$merge: {
        paid: 0
      }})} />
  );
  fireEvent.click(test_input)
  expect(onEdit).toHaveBeenCalledTimes(2);

});

it('renders in the BoolTrueDisabled state', () => {
  const { container, rerender } = render(
    <BoolTrueDisabled {...BoolTrueDisabled.args} id="test_input" />
  );
  expect(getById(container, 'test_input')).toBeDefined();

  rerender(
    <BoolTrueDisabled {...BoolTrueDisabled.args} id="test_input" 
      values={update(BoolTrueDisabled.args.field, {$merge: {
        paid: 0
      }})} />
  );

});

it('renders in the DateDisabled state', () => {
  const { container } = render(
    <DateDisabled {...DateDisabled.args} id="test_input" />
  );
  expect(getById(container, 'test_input')).toBeDefined();

});

it('renders in the DateTime state', () => {
  const onEdit = jest.fn()

  const { container } = render(
    <DateTime {...DateTime.args} id="test_input" onEdit={onEdit} />
  );
  expect(getById(container, 'test_input')).toBeDefined();

  const input = getById(container, 'test_input')
  fireEvent.change(input, {target: {value: "2021-12-24"}})
  fireEvent.keyDown(input, { key: 'Enter', code: 'Enter', keyCode: 13 })
  expect(onEdit).toHaveBeenCalledTimes(1);

  fireEvent.change(input, {target: {value: ""}})
  fireEvent.keyDown(input, { key: 'Enter', code: 'Enter', keyCode: 13 })
  expect(onEdit).toHaveBeenCalledTimes(2);

});

it('renders in the DateExtend state', () => {
  const { container } = render(
    <DateExtend {...DateExtend.args} id="test_input" />
  );
  expect(getById(container, 'test_input')).toBeDefined();

});

it('renders in the DateLink state', () => {
  const { container } = render(
    <DateLink {...DateLink.args} id="test_input" />
  );
  expect(getById(container, 'test_input')).toBeDefined();
  
  render(
    <DateLink {...DateLink.args} 
      data={{
        dataset: {},
        current: {
          trans: []
        }, 
        audit: "all",
      }} />
  )
});

it('renders in the DateLinkValue state', () => {
  const { container } = render(
    <DateLinkValue {...DateLinkValue.args} id="test_input" />
  );
  expect(getById(container, 'test_input')).toBeDefined();
  
});

it('renders in the Password state', () => {
  const onEdit = jest.fn()

  const { container } = render(
    <Password {...Password.args} id="test_input" onEdit={onEdit} />
  );
  expect(getById(container, 'test_input')).toBeDefined();

  const test_input = getById(container, 'test_input')

  fireEvent.change(test_input, {target: {value: "123"}})
  expect(onEdit).toHaveBeenCalledTimes(1);

  render(
    <Password {...Password.args} 
      values={{
        password_1: null,
      }}
      data={{
        dataset: {}, 
        current: {}, 
        audit: "readonly",
      }} />
  );

});

it('renders in the Color state', () => {
  const onEdit = jest.fn()

  const { container } = render(
    <Color {...Color.args} id="test_input" onEdit={onEdit} />
  );
  expect(getById(container, 'test_input')).toBeDefined();

  const test_input = getById(container, 'test_input')

  fireEvent.change(test_input, {target: {value: "#ffffff"}})
  expect(onEdit).toHaveBeenCalledTimes(1);

});

it('renders in the Fieldvalue state', () => {

  const { container } = render(
    <Fieldvalue {...Fieldvalue.args} id="test_input" />
  );
  expect(getById(container, 'test_input')).toBeDefined();

});