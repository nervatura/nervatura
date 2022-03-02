import { render, fireEvent, queryByAttribute } from '@testing-library/react';
import '@testing-library/jest-dom/extend-expect';

import { Default, Disabled } from './Menu.stories';

const getById = queryByAttribute.bind(null, 'id');

it('renders in the Default state', () => {
  const onMenu = jest.fn()
  const onClose = jest.fn()

  const { container } = render(
    <Default {...Default.args} id="test_settings"
    onMenu={onMenu} onClose={onClose} />
  );
  expect(getById(container, 'test_settings')).toBeDefined();

  const closeIcon = getById(container, 'closeIcon')
  fireEvent.click(closeIcon)
  expect(onClose).toHaveBeenCalledTimes(1);

  const btn_cancel = getById(container, 'btn_cancel')
  fireEvent.click(btn_cancel)
  expect(onClose).toHaveBeenCalledTimes(2);

  const btn_ok = getById(container, 'btn_ok')
  fireEvent.click(btn_ok)
  expect(onMenu).toHaveBeenCalledTimes(1);

  const fieldname = getById(container, 'fieldname')
  fireEvent.change(fieldname, {target: {value: "value"}})
  expect(fieldname.value).toEqual("value");

  const description = getById(container, 'description')
  fireEvent.change(description, {target: {value: "value"}})
  expect(description.value).toEqual("value");

  const fieldtype = getById(container, 'fieldtype')
  fireEvent.change(fieldtype, {target: {value: "37"}})
  expect(fieldtype.value).toEqual("37");

  const orderby = getById(container, 'orderby')
  fireEvent.change(orderby, {target: {value: "12"}})
  expect(orderby.value).toEqual("12");

})

it('renders in the Disabled state', () => {
  const { container } = render(
    <Disabled {...Disabled.args} id="test_settings" />
  );
  expect(getById(container, 'test_settings')).toBeDefined();
})