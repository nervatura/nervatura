import { render, queryByAttribute, fireEvent, screen } from '@testing-library/react';
import '@testing-library/jest-dom/extend-expect';

import { Default, Details, Row, Cell, Data, StringData, ListData, ListDataItem, TableData, TableDataItem, Meta, VGap } from './TemplateEditor.stories';

const getById = queryByAttribute.bind(null, 'id');

it('renders in the Default state', () => {
  const onEvent = jest.fn()

  const { container } = render(
    <Default {...Default.args} id="test_template" onEvent={onEvent} />
  );
  expect(getById(container, "test_template")).toBeDefined();

  //navButton
  const btn_next = getById(container, 'btn_next')
  fireEvent.click(btn_next)
  expect(onEvent).toHaveBeenCalledTimes(2);

  //mapButton
  const btn_tmp_report = getById(container, 'btn_tmp_report')
  fireEvent.click(btn_tmp_report)
  expect(onEvent).toHaveBeenCalledTimes(3);

  //Row editItem
  const field_title = getById(container, 'field_title')
  fireEvent.change(field_title, {target: {value: "test value"}})
  expect(onEvent).toHaveBeenCalledTimes(4);

  //tabButton
  const btn_data = getById(container, 'btn_data')
  fireEvent.click(btn_data)
  expect(onEvent).toHaveBeenCalledTimes(5);

})

it('renders in the Details state', () => {
  const onEvent = jest.fn()

  const { container } = render(
    <Details {...Details.args} id="test_template" onEvent={onEvent} />
  );
  expect(getById(container, "test_template")).toBeDefined();

  const btn_add_item = getById(container, 'btn_add_item')
  fireEvent.click(btn_add_item)
  expect(onEvent).toHaveBeenCalledTimes(2);

})

it('renders in the Row state', () => {
  const onEvent = jest.fn()

  const { container } = render(
    <Row {...Row.args} id="test_template" onEvent={onEvent} />
  );
  expect(getById(container, "test_template")).toBeDefined();

  const btn_add_item = getById(container, 'btn_add_item')
  fireEvent.click(btn_add_item)
  expect(onEvent).toHaveBeenCalledTimes(2);

  const sel_add_item = getById(container, 'sel_add_item')
  fireEvent.change(sel_add_item, {target: {value: "image"}})
  expect(onEvent).toHaveBeenCalledTimes(3);

})

it('renders in the Cell state', () => {
  const onEvent = jest.fn()

  const { container } = render(
    <Cell {...Cell.args} id="test_template" onEvent={onEvent} />
  );
  expect(getById(container, "test_template")).toBeDefined();

})

it('renders in the Data state', () => {
  const onEvent = jest.fn()

  const { container } = render(
    <Data {...Data.args} id="test_template" onEvent={onEvent} />
  );
  expect(getById(container, "test_template")).toBeDefined();

  const btn_add = getById(container, 'btn_add')
  fireEvent.click(btn_add)
  expect(onEvent).toHaveBeenCalledTimes(1);

  const row_edit = getById(container, 'row_edit_0')
  fireEvent.click(row_edit)
  expect(onEvent).toHaveBeenCalledTimes(2);

  const row_delete = getById(container, 'row_delete_0')
  fireEvent.click(row_delete)
  expect(onEvent).toHaveBeenCalledTimes(3);

})

it('renders in the StringData state', () => {
  const onEvent = jest.fn()

  const { container } = render(
    <StringData {...StringData.args} id="test_template" onEvent={onEvent} />
  );
  expect(getById(container, "test_template")).toBeDefined();

  const html_text_value = getById(container, 'html_text_value')
  fireEvent.change(html_text_value, {target: {value: "test value"}})
  expect(onEvent).toHaveBeenCalledTimes(1);

})

it('renders in the ListData state', () => {
  const onEvent = jest.fn()

  const { container } = render(
    <ListData {...ListData.args} id="test_template" onEvent={onEvent} />
  );
  expect(getById(container, "test_template")).toBeDefined();

  const btn_add = getById(container, 'btn_add')
  fireEvent.click(btn_add)
  expect(onEvent).toHaveBeenCalledTimes(1);

  const row_edit = getById(container, 'row_edit_0')
  fireEvent.click(row_edit)
  expect(onEvent).toHaveBeenCalledTimes(2);

  const row_delete = getById(container, 'row_delete_0')
  fireEvent.click(row_delete)
  expect(onEvent).toHaveBeenCalledTimes(3);

})

it('renders in the ListDataItem state', () => {
  const onEvent = jest.fn()

  const { container } = render(
    <ListDataItem {...ListDataItem.args} id="test_template" onEvent={onEvent} />
  );
  expect(getById(container, "test_template")).toBeDefined();

  const title_value = getById(container, 'title_value')
  fireEvent.change(title_value, {target: {value: "test value"}})
  expect(onEvent).toHaveBeenCalledTimes(1);

})

it('renders in the TableData state', () => {
  const onEvent = jest.fn()

  const { container } = render(
    <TableData {...TableData.args} id="test_template" onEvent={onEvent} />
  );
  expect(getById(container, "test_template")).toBeDefined();
  
  const btn_add = getById(container, 'btn_add')
  fireEvent.click(btn_add)
  expect(onEvent).toHaveBeenCalledTimes(1);

  const delete_row = getById(container, 'delete_0')
  fireEvent.click(delete_row)
  expect(onEvent).toHaveBeenCalledTimes(2);

  const row_selected = screen.getAllByRole('row')[2]
  fireEvent.click(row_selected)
  expect(onEvent).toHaveBeenCalledTimes(3);

})

it('renders in the TableDataItem state', () => {
  const onEvent = jest.fn()

  const { container } = render(
    <TableDataItem {...TableDataItem.args} id="test_template" onEvent={onEvent} />
  );
  expect(getById(container, "test_template")).toBeDefined();

  const text_value = getById(container, 'text_value')
  fireEvent.change(text_value, {target: {value: "test value"}})
  expect(onEvent).toHaveBeenCalledTimes(1);

})

it('renders in the Meta state', () => {
  const onEvent = jest.fn()

  const { container } = render(
    <Meta {...Meta.args} id="test_template" onEvent={onEvent} />
  );
  expect(getById(container, "test_template")).toBeDefined();

})

it('renders in the VGap state', () => {
  const onEvent = jest.fn()

  const { container } = render(
    <VGap {...VGap.args} id="test_template" onEvent={onEvent} />
  );
  expect(getById(container, "test_template")).toBeDefined();

})