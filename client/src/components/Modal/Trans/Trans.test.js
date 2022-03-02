import { render, fireEvent, queryByAttribute } from '@testing-library/react';
import '@testing-library/jest-dom/extend-expect';

import { Default, Order, Worksheet } from './Trans.stories';

const getById = queryByAttribute.bind(null, 'id');

it('renders in the Default state', () => {
  const onCreate = jest.fn()
  const onClose = jest.fn()

  const { container } = render(
    <Default {...Default.args} id="test_settings"
      onCreate={onCreate} onClose={onClose} />
  );
  expect(getById(container, 'test_settings')).toBeDefined();

  const closeIcon = getById(container, 'closeIcon')
  fireEvent.click(closeIcon)
  expect(onClose).toHaveBeenCalledTimes(1);

  const btn_cancel = getById(container, 'btn_cancel')
  fireEvent.click(btn_cancel)
  expect(onClose).toHaveBeenCalledTimes(2);

  const btn_create = getById(container, 'btn_create')
  fireEvent.click(btn_create)
  expect(onCreate).toHaveBeenCalledTimes(1);

  const transtype = getById(container, 'transtype')
  fireEvent.change(transtype, {target: {value: "worksheet"}})
  expect(transtype.value).toEqual("worksheet");

  const direction = getById(container, 'direction')
  fireEvent.change(direction, {target: {value: "in"}})
  expect(direction.value).toEqual("in");

})

it('renders in the Order state', () => {
  const onCreate = jest.fn()

  const { container } = render(
    <Order {...Order.args} id="test_settings"
    onCreate={onCreate} />
  );
  expect(getById(container, 'test_settings')).toBeDefined();

  const transtype = getById(container, 'transtype')
  fireEvent.change(transtype, {target: {value: "receipt"}})
  expect(transtype.value).toEqual("receipt");

  const refno = getById(container, 'refno')
  fireEvent.click(refno)

  const netto = getById(container, 'netto')
  fireEvent.click(netto)

  const from = getById(container, 'from')
  fireEvent.click(from)

  const btn_create = getById(container, 'btn_create')
  fireEvent.click(btn_create)
  expect(onCreate).toHaveBeenCalledTimes(1);

})

it('renders in the Worksheet state', () => {
  const onCreate = jest.fn()

  const { container } = render(
    <Worksheet {...Worksheet.args} id="test_settings"
    onCreate={onCreate} />
  );
  expect(getById(container, 'test_settings')).toBeDefined();

  const transtype = getById(container, 'transtype')
  fireEvent.change(transtype, {target: {value: "receipt"}})
  expect(transtype.value).toEqual("receipt");

  const btn_create = getById(container, 'btn_create')
  fireEvent.click(btn_create)
  expect(onCreate).toHaveBeenCalledTimes(1);

})