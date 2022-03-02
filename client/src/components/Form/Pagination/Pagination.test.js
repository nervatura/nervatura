import { render, fireEvent, queryByAttribute } from '@testing-library/react';

import '@testing-library/jest-dom/extend-expect';

import { Default, Items } from './Pagination.stories';

const getById = queryByAttribute.bind(null, 'id');

it('renders in the Default state', () => {

  const { container } = render(
    <Default {...Default.args} id="test_paginator" />
  );
  expect(getById(container, 'test_paginator')).toBeDefined();

});

it('renders in the Items state', () => {
  const onEvent = jest.fn()
  const { container } = render(
    <Items {...Items.args} id="test_paginator" onEvent={onEvent} />
  );
  expect(getById(container, 'test_paginator')).toBeDefined();
  
  const btn_first = getById(container, 'btn_first')
  fireEvent.click(btn_first)
  expect(onEvent).toHaveBeenCalledTimes(1);

  const btn_previous = getById(container, 'btn_previous')
  fireEvent.click(btn_previous)
  expect(onEvent).toHaveBeenCalledTimes(2);

  const btn_next = getById(container, 'btn_next')
  fireEvent.click(btn_next)
  expect(onEvent).toHaveBeenCalledTimes(3);

  const btn_last = getById(container, 'btn_last')
  fireEvent.click(btn_last)
  expect(onEvent).toHaveBeenCalledTimes(4);

  const input_goto = getById(container, 'input_goto')
  fireEvent.change(input_goto, {target: {value: "1"}})
  expect(onEvent).toHaveBeenCalledTimes(5);

  const sel_page_size = getById(container, 'sel_page_size')
  fireEvent.change(sel_page_size, {target: {value: "50"}})
  expect(onEvent).toHaveBeenCalledTimes(6);

});