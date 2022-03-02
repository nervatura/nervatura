import { render, queryByAttribute, fireEvent, screen } from '@testing-library/react';
import '@testing-library/jest-dom/extend-expect';

import { Default, HideHeader, Columns, Filters, FormActions } from './Browser.stories';

const getById = queryByAttribute.bind(null, 'id');

it('renders in the Default state', () => {
  const onEvent = jest.fn()

  const { container } = render(
    <Default {...Default.args} id="test_browser" onEvent={onEvent} />
  );
  expect(getById(container, "test_browser")).toBeDefined();

  const btn_search = getById(container, 'btn_search')
  fireEvent.click(btn_search)
  expect(onEvent).toHaveBeenCalledTimes(1);

  const btn_bookmark = getById(container, 'btn_bookmark')
  fireEvent.click(btn_bookmark)
  expect(onEvent).toHaveBeenCalledTimes(2);

  const btn_export = getById(container, 'btn_export')
  fireEvent.click(btn_export)
  expect(onEvent).toHaveBeenCalledTimes(3);

  const btn_help = getById(container, 'btn_help')
  fireEvent.click(btn_help)
  expect(onEvent).toHaveBeenCalledTimes(4);

  const btn_views = getById(container, 'btn_views')
  fireEvent.click(btn_views)
  expect(onEvent).toHaveBeenCalledTimes(4);

  const btn_columns = getById(container, 'btn_columns')
  fireEvent.click(btn_columns)
  expect(onEvent).toHaveBeenCalledTimes(5);

  const btn_filter = getById(container, 'btn_filter')
  fireEvent.click(btn_filter)
  expect(onEvent).toHaveBeenCalledTimes(6);

  const btn_total = getById(container, 'btn_total')
  fireEvent.click(btn_total)
  expect(onEvent).toHaveBeenCalledTimes(7);

  const row_item = getById(container, 'edit_customer//2')
  fireEvent.click(row_item)
  expect(onEvent).toHaveBeenCalledTimes(8);

  const page_2 = screen.getByText("2")
  fireEvent.click(page_2)
  expect(onEvent).toHaveBeenCalledTimes(9);

})

it('renders in the HideHeader state', () => {
  const onEvent = jest.fn()

  const { container } = render(
    <HideHeader {...HideHeader.args} id="test_browser" onEvent={onEvent} />
  );
  expect(getById(container, "test_browser")).toBeDefined();

  const btn_header = getById(container, 'btn_header')
  fireEvent.click(btn_header)
  expect(onEvent).toHaveBeenCalledTimes(1);

})

it('renders in the Columns state', () => {
  const onEvent = jest.fn()

  const { container } = render(
    <Columns {...Columns.args} id="test_browser" onEvent={onEvent} />
  );
  expect(getById(container, "test_browser")).toBeDefined();

  const view_item = getById(container, 'view_CustomerAddressView')
  fireEvent.click(view_item)
  expect(onEvent).toHaveBeenCalledTimes(1);

  const col_surname = getById(container, 'col_surname')
  fireEvent.click(col_surname)
  expect(onEvent).toHaveBeenCalledTimes(2);

  const col_status = getById(container, 'col_status')
  fireEvent.click(col_status)
  expect(onEvent).toHaveBeenCalledTimes(3);

})

it('renders in the Filters state', () => {
  const onEvent = jest.fn()

  const { container } = render(
    <Filters {...Filters.args} id="test_browser" onEvent={onEvent} />
  );
  expect(getById(container, "test_browser")).toBeDefined();

  const btn_delete = getById(container, 'btn_delete_filter_0')
  fireEvent.click(btn_delete)
  expect(onEvent).toHaveBeenCalledTimes(1);

  const link_cell = getById(container, 'link_31')
  fireEvent.click(link_cell)
  expect(onEvent).toHaveBeenCalledTimes(2);

  const filter_name = getById(container, 'filter_name_0')
  fireEvent.change(filter_name, {target: {value: "sample_customer_date"}})
  expect(onEvent).toHaveBeenCalledTimes(3);

  const filter_type = getById(container, 'filter_type_0')
  fireEvent.change(filter_type, {target: {value: "==N"}})
  expect(onEvent).toHaveBeenCalledTimes(4);

  const filter_value_number = getById(container, 'filter_value_0')
  fireEvent.change(filter_value_number, {target: {value: "111"}})
  expect(onEvent).toHaveBeenCalledTimes(5);

  const filter_value_date = getById(container, 'filter_value_1')
  filter_value_date.value = "2022-01-02"
  fireEvent.keyDown(filter_value_date, { key: 'Enter', code: 'Enter', keyCode: 13 })
  expect(onEvent).toHaveBeenCalledTimes(6);

  const filter_value_bool = getById(container, 'filter_value_2')
  fireEvent.change(filter_value_bool, {target: {value: "0"}})
  expect(onEvent).toHaveBeenCalledTimes(7);

  const filter_value_string = getById(container, 'filter_value_3')
  fireEvent.change(filter_value_string, {target: {value: "red"}})
  expect(onEvent).toHaveBeenCalledTimes(8);

})

it('renders in the FormActions state', () => {
  const onEvent = jest.fn()

  const { container } = render(
    <FormActions {...FormActions.args} id="test_browser" onEvent={onEvent} />
  );
  expect(getById(container, "test_browser")).toBeDefined();

  const btn_actions_new = getById(container, 'btn_actions_new')
  fireEvent.click(btn_actions_new)
  expect(onEvent).toHaveBeenCalledTimes(1);

})