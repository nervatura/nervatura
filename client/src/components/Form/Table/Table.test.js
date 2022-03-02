import { render, fireEvent, screen, queryByAttribute } from '@testing-library/react';

import '@testing-library/jest-dom/extend-expect';

import { Default, TopPagination, BottomPagination, Filtered } from './Table.stories';

const getById = queryByAttribute.bind(null, 'id');

beforeEach(() => {
  Object.defineProperty(global.window, 'scrollTo', { value: jest.fn() });
});

afterEach(() => {
  jest.clearAllMocks();
});

it('renders in the Default state', () => {

  const { container } = render(
    <Default {...Default.args} id="test_table" />
  );
  expect(getById(container, 'test_table')).toBeDefined();

  const row_selected = screen.getAllByRole('row')[2]
  fireEvent.click(row_selected)

});

it('renders in the TopPagination state', () => {
  const onRowSelected = jest.fn()
  const onCurrentPage = jest.fn()
  const onEditCell = jest.fn()

  const { container } = render(
    <TopPagination {...TopPagination.args} id="test_table"
      onRowSelected={onRowSelected} onCurrentPage={onCurrentPage} onEditCell={onEditCell} />
  );
  expect(getById(container, 'test_table')).toBeDefined();

  const row_selected = screen.getAllByRole('row')[2]
  fireEvent.click(row_selected)
  expect(onRowSelected).toHaveBeenCalledTimes(1);

  const link_cell = screen.getByText('Name link')
  fireEvent.click(link_cell)
  expect(onEditCell).toHaveBeenCalledTimes(1);

  const page_2 = screen.getByText("2")
  fireEvent.click(page_2)
  expect(onCurrentPage).toHaveBeenCalledTimes(1);

  const sort_header = screen.getAllByText('Stamp')[0]
  fireEvent.click(sort_header)
  fireEvent.click(sort_header)

});

it('renders in the BottomPagination state', () => {

  const { container } = render(
    <BottomPagination {...BottomPagination.args} id="test_table" />
  );
  expect(getById(container, 'test_table')).toBeDefined();

  const page_1 = screen.getByText("1")
  fireEvent.click(page_1)

});

it('renders in the Filtered state', () => {
  const onAddItem = jest.fn()

  const { container } = render(
    <Filtered {...Filtered.args} id="test_table" onAddItem={onAddItem} />
  );
  expect(getById(container, 'test_table')).toBeDefined();

  const btn_add = getById(container, 'btn_add')
  fireEvent.click(btn_add)
  expect(onAddItem).toHaveBeenCalledTimes(1);

  const filter = getById(container, 'filter')
  fireEvent.change(filter, {target: {value: "filter"}})
  expect(filter.value).toEqual("filter");

});
