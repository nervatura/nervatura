import { render, queryByAttribute, fireEvent } from '@testing-library/react';
import '@testing-library/jest-dom/extend-expect';
import update from 'immutability-helper';

import { Default, Items, Log } from './Form.stories';

const getById = queryByAttribute.bind(null, 'id');

it('renders in the Default state', () => {
  const onEvent = jest.fn()

  const { container } = render(
    <Default {...Default.args} id="test_editor" onEvent={onEvent} />
  );
  expect(getById(container, "test_editor")).toBeDefined();

  const field_value = getById(container, 'field_fieldvalue_value')
  fireEvent.change(field_value, {target: {value: "test data"}})
  expect(onEvent).toHaveBeenCalledTimes(1);

})

it('renders in the Items state', () => {
  const onEvent = jest.fn()

  const { container } = render(
    <Items {...Items.args} id="test_editor" onEvent={onEvent} />
  );
  expect(getById(container, "test_editor")).toBeDefined();
  
  const btn_add = getById(container, 'btn_add')
  fireEvent.click(btn_add)
  expect(onEvent).toHaveBeenCalledTimes(1);

  const btn_edit = getById(container, 'edit_1')
  fireEvent.click(btn_edit)
  expect(onEvent).toHaveBeenCalledTimes(2);

  const btn_delete = getById(container, 'delete_1')
  fireEvent.click(btn_delete)
  expect(onEvent).toHaveBeenCalledTimes(3);

  const delete_data = update(Items.args.data, {current: {template: {view: {items: {actions: {$merge: {
    delete: null
  }}}}}}})
  render(
    <Items {...Items.args} id="test_editor" data={delete_data}  />
  )

  const edit_data = update(Items.args.data, {current: {template: {view: {items: {actions: {$merge: {
    edit: null
  }}}}}}})
  render(
    <Items {...Items.args} id="test_editor" data={edit_data}  />
  )

})

it('renders in the Log state', () => {
  const onEvent = jest.fn()

  const { container } = render(
    <Log {...Log.args} id="test_editor" onEvent={onEvent} />
  );
  expect(getById(container, "test_editor")).toBeDefined();

})