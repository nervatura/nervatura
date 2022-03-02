import { render, queryByAttribute, fireEvent, screen } from '@testing-library/react';
import '@testing-library/jest-dom/extend-expect';

import { Default, Table, ReadOnlyList, ReadOnlyTable } from './View.stories';

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

  const btn_edit = getById(container, 'row_edit_0')
  fireEvent.click(btn_edit)
  expect(onEvent).toHaveBeenCalledTimes(2);

  const btn_delete = getById(container, 'row_delete_0')
  fireEvent.click(btn_delete)
  expect(onEvent).toHaveBeenCalledTimes(3);

  const btn_next = getById(container, 'btn_next')
  fireEvent.click(btn_next)
  expect(onEvent).toHaveBeenCalledTimes(4);

})

it('renders in the Table state', () => {
  const onEvent = jest.fn()

  const { container } = render(
    <Table {...Table.args} id="test_editor" onEvent={onEvent} />
  );
  expect(getById(container, "test_editor")).toBeDefined();

  const btn_add = getById(container, 'btn_add')
  fireEvent.click(btn_add)
  expect(onEvent).toHaveBeenCalledTimes(1);

  const btn_edit = getById(container, 'edit_1')
  fireEvent.click(btn_edit)
  expect(onEvent).toHaveBeenCalledTimes(2);

  const row_item = screen.getAllByRole('row')[2]
  fireEvent.click(row_item)
  expect(onEvent).toHaveBeenCalledTimes(3);

})

it('renders in the ReadOnlyList state', () => {

  const { container } = render(
    <ReadOnlyList {...ReadOnlyList.args} id="test_editor" />
  );
  expect(getById(container, "test_editor")).toBeDefined();

})

it('renders in the ReadOnlyTable state', () => {

  const { container } = render(
    <ReadOnlyTable {...ReadOnlyTable.args} id="test_editor" />
  );
  expect(getById(container, "test_editor")).toBeDefined();

})