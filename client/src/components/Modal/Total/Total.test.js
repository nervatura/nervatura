import { render, fireEvent, queryByAttribute } from '@testing-library/react';
import '@testing-library/jest-dom/extend-expect';

import { Default, DarkTotal } from './Total.stories';

const getById = queryByAttribute.bind(null, 'id');

it('renders in the Default state', () => {
  const onClose = jest.fn()

  const { container } = render(
    <Default {...Default.args} id="test_settings" onClose={onClose} />
  );
  expect(getById(container, 'test_settings')).toBeDefined();

  const closeIcon = getById(container, 'closeIcon')
  fireEvent.click(closeIcon)
  expect(onClose).toHaveBeenCalledTimes(1);

  const btn_ok = getById(container, 'btn_ok')
  fireEvent.click(btn_ok)
  expect(onClose).toHaveBeenCalledTimes(2);

})

it('renders in the DarkTotal state', () => {
  const { container } = render(
    <DarkTotal {...DarkTotal.args} id="test_settings" />
  );
  expect(getById(container, 'test_settings')).toBeDefined();
})