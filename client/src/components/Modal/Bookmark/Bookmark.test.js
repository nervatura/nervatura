import { render, fireEvent, queryByAttribute } from '@testing-library/react';
import '@testing-library/jest-dom/extend-expect';

import { Default, BookmarkData, DarkTheme } from './Bookmark.stories';

const getById = queryByAttribute.bind(null, 'id');

it('renders in the Default state', () => {

  const { container } = render(
    <Default {...Default.args} id="test_bookmark" />
  );
  expect(getById(container, 'test_bookmark')).toBeDefined();
})

it('renders in the BookmarkData state', () => {
  const onSelect = jest.fn()
  const onDelete = jest.fn()
  const onClose = jest.fn()

  const { container } = render(
    <BookmarkData {...BookmarkData.args} id="test_bookmark" 
      onSelect={onSelect} onDelete={onDelete} onClose={onClose} />
  );
  expect(getById(container, 'test_bookmark')).toBeDefined();

  const row_item = getById(container, 'row_item_1')
  fireEvent.click(row_item)
  expect(onSelect).toHaveBeenCalledTimes(1);

  const row_delete = getById(container, 'row_delete_1')
  fireEvent.click(row_delete)
  expect(onDelete).toHaveBeenCalledTimes(1);

  const closeIcon = getById(container, 'closeIcon')
  fireEvent.click(closeIcon)
  expect(onClose).toHaveBeenCalledTimes(1);

  const btn_bookmark = getById(container, 'btn_bookmark')
  fireEvent.click(btn_bookmark)

  const btn_history = getById(container, 'btn_history')
  fireEvent.click(btn_history)
  
})

it('renders in the DarkTheme state', () => {

  const { container } = render(
    <DarkTheme {...DarkTheme.args} id="test_bookmark" />
  );
  expect(getById(container, 'test_bookmark')).toBeDefined();
})