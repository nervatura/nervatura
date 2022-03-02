import { render, queryByAttribute, fireEvent, screen } from '@testing-library/react';
import '@testing-library/jest-dom/extend-expect';

import { Default, NewField, Customer, ReadOnly } from './Meta.stories';

const getById = queryByAttribute.bind(null, 'id');

it('renders in the Default state', () => {
  const onEvent = jest.fn()

  const { container } = render(
    <Default {...Default.args} id="test_editor" onEvent={onEvent} />
  );
  expect(getById(container, "test_editor")).toBeDefined();

  //onSelector
  const sel_show_fieldvalue_value = getById(container, 'sel_show_fieldvalue_value')
  fireEvent.click(sel_show_fieldvalue_value)
  expect(onEvent).toHaveBeenCalledTimes(1);

  //editItem
  const field_vlist = getById(container, 'field_4e451b7f-72d1-b19c-7cbe-2c80495b5a8e')
  fireEvent.change(field_vlist, {target: {value: "red"}})
  expect(onEvent).toHaveBeenCalledTimes(2);

  const sel_deffield = getById(container, 'sel_deffield')
  fireEvent.change(sel_deffield, {target: {value: "trans_transitem_link"}})
  expect(onEvent).toHaveBeenCalledTimes(3);

})

it('renders in the NewField state', () => {
  const onEvent = jest.fn()

  const { container } = render(
    <NewField {...NewField.args} id="test_editor" onEvent={onEvent} />
  );
  expect(getById(container, "test_editor")).toBeDefined();

  const btn_new = getById(container, 'btn_new')
  fireEvent.click(btn_new)
  expect(onEvent).toHaveBeenCalledTimes(1);

  const page_2 = screen.getByText("2")
  fireEvent.click(page_2)
  expect(onEvent).toHaveBeenCalledTimes(2);

  const sel_page_size = getById(container, 'sel_page_size')
  fireEvent.change(sel_page_size, {target: {value: "50"}})

})

it('renders in the Customer state', () => {
  const { container } = render(
    <Customer {...Customer.args} id="test_editor" />
  );
  expect(getById(container, "test_editor")).toBeDefined();

})

it('renders in the ReadOnly state', () => {
  const { container } = render(
    <ReadOnly {...ReadOnly.args} id="test_editor" />
  );
  expect(getById(container, "test_editor")).toBeDefined();

  render(<ReadOnly {...ReadOnly.args} id="test_editor"
    current={{
      ...Default.args.current,
    }}
    paginationPage={2}
  />)

})