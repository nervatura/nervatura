import { render, queryByAttribute, fireEvent } from '@testing-library/react';
import '@testing-library/jest-dom/extend-expect';

import { Default, List, ReadOnly, Empty, DeleteOnly, ReadOnlyList } from './View.stories';

const getById = queryByAttribute.bind(null, 'id');

it('renders in the Default state', () => {
  const onEvent = jest.fn()

  const { container } = render(
    <Default {...Default.args} id="test_editor" onEvent={onEvent} />
  );
  expect(getById(container, "test_editor")).toBeDefined();

  const btn_add = getById(container, 'btn_add')
  fireEvent.click(btn_add)
  expect(onEvent).toHaveBeenCalledTimes(1);

  const btn_edit = getById(container, 'edit_18')
  fireEvent.click(btn_edit)
  expect(onEvent).toHaveBeenCalledTimes(2);

  const btn_delete = getById(container, 'delete_18')
  fireEvent.click(btn_delete)
  expect(onEvent).toHaveBeenCalledTimes(3);

})

it('renders in the List state', () => {
  const onEvent = jest.fn()

  const { container } = render(
    <List {...List.args} id="test_editor" onEvent={onEvent} />
  );
  expect(getById(container, "test_editor")).toBeDefined();

  const row_edit = getById(container, 'row_edit_0')
  fireEvent.click(row_edit)
  expect(onEvent).toHaveBeenCalledTimes(1);

  const row_delete = getById(container, 'row_delete_0')
  fireEvent.click(row_delete)
  expect(onEvent).toHaveBeenCalledTimes(2);

  const btn_add = getById(container, 'btn_add')
  fireEvent.click(btn_add)
  expect(onEvent).toHaveBeenCalledTimes(3);

})

it('renders in the ReadOnly state', () => {
  const onEvent = jest.fn()

  const { container } = render(
    <ReadOnly {...ReadOnly.args} id="test_editor" onEvent={onEvent} />
  );
  expect(getById(container, "test_editor")).toBeDefined();

})

it('renders in the Empty state', () => {
  const { container } = render(
    <Empty {...Empty.args} id="test_editor" />
  );
  expect(getById(container, "test_editor")).toBeDefined();

})

it('renders in the DeleteOnly state', () => {
  const { container } = render(
    <DeleteOnly {...DeleteOnly.args} id="test_editor" />
  );
  expect(getById(container, "test_editor")).toBeDefined();

})

it('renders in the ReadOnlyList state', () => {
  const { container } = render(
    <ReadOnlyList {...ReadOnlyList.args} id="test_editor" />
  );
  expect(getById(container, "test_editor")).toBeDefined();

})