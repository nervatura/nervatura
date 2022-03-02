import { render, fireEvent, queryByAttribute, screen } from '@testing-library/react';
import '@testing-library/jest-dom/extend-expect';

import { Default, QuickView } from './Selector.stories';

const getById = queryByAttribute.bind(null, 'id');

it('renders in the Default state', () => {
  const onSelect = jest.fn()
  const onSearch = jest.fn()
  const onClose = jest.fn()

  const { container } = render(
    <Default {...Default.args} id="test_selector"
      onSelect={onSelect} onSearch={onSearch} onClose={onClose} />
  );
  expect(getById(container, 'test_selector')).toBeDefined();

  const row_item = screen.getAllByRole('row')[2]

  fireEvent.click(row_item)
  expect(onSelect).toHaveBeenCalledTimes(1);

  const closeIcon = getById(container, 'closeIcon')
  fireEvent.click(closeIcon)
  expect(onClose).toHaveBeenCalledTimes(1);

  const btn_search = getById(container, 'btn_search')
  fireEvent.click(btn_search)
  expect(onSearch).toHaveBeenCalledTimes(1);

  const filter = getById(container, 'filter')
  fireEvent.change(filter, {target: {value: "filter"}})
  expect(filter.value).toEqual("filter");
  fireEvent.keyDown(filter, { key: 'Enter', code: 'Enter', keyCode: 13 })
  expect(onSearch).toHaveBeenCalledTimes(2)

})

it('renders in the QuickView state', () => {
  const { container } = render(
    <QuickView {...QuickView.args} id="test_selector" />
  );
  expect(getById(container, 'test_selector')).toBeDefined();
})