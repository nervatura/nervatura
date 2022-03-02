import { render, fireEvent, queryByAttribute } from '@testing-library/react';
import '@testing-library/jest-dom/extend-expect';

import { Default, DarkServer } from './Server.stories';

const getById = queryByAttribute.bind(null, 'id');

it('renders in the Default state', () => {
  const onOK = jest.fn()
  const onClose = jest.fn()

  const { container } = render(
    <Default {...Default.args} id="test_shipping"
      onOK={onOK} onClose={onClose} />
  );
  expect(getById(container, 'test_shipping')).toBeDefined();

  const closeIcon = getById(container, 'closeIcon')
  fireEvent.click(closeIcon)
  expect(onClose).toHaveBeenCalledTimes(1);

  const btn_cancel = getById(container, 'btn_cancel')
  fireEvent.click(btn_cancel)
  expect(onClose).toHaveBeenCalledTimes(2);

  const btn_ok = getById(container, 'btn_ok')
  fireEvent.click(btn_ok)
  expect(onOK).toHaveBeenCalledTimes(1);

  const field_step = getById(container, 'field_step')
  fireEvent.click(field_step)

})

it('renders in the DarkServer state', () => {
  const { container } = render(
    <DarkServer {...DarkServer.args} id="test_shipping" />
  );
  expect(getById(container, 'test_shipping')).toBeDefined();

})