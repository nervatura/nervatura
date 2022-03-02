import { render, queryByAttribute, fireEvent } from '@testing-library/react';
import '@testing-library/jest-dom/extend-expect';

import { Default, NewItem } from './Item.stories';

const getById = queryByAttribute.bind(null, 'id');

it('renders in the Default state', () => {
  const onEvent = jest.fn()

  const { container } = render(
    <Default {...Default.args} id="test_editor" onEvent={onEvent} />
  );
  expect(getById(container, "test_editor")).toBeDefined();

  //onSelector
  const sel_show_product_id = getById(container, 'sel_show_product_id')
  fireEvent.click(sel_show_product_id)
  expect(onEvent).toHaveBeenCalledTimes(1);

  //editItem
  const field_qty = getById(container, 'field_qty')
  fireEvent.change(field_qty, {target: {value: "2"}})
  expect(onEvent).toHaveBeenCalledTimes(2);

})

it('renders in the NewItem state', () => {
  const { container } = render(
    <NewItem {...NewItem.args} id="test_editor" />
  );
  expect(getById(container, "test_editor")).toBeDefined();

})