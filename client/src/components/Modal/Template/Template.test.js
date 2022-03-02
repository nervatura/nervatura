import { render, fireEvent, queryByAttribute } from '@testing-library/react';
import '@testing-library/jest-dom/extend-expect';

import { Default, Disabled } from './Template.stories';

const getById = queryByAttribute.bind(null, 'id');

it('renders in the Default state', () => {
  const onData = jest.fn()
  const onClose = jest.fn()

  const { container } = render(
    <Default {...Default.args} id="test_settings"
    onData={onData} onClose={onClose} />
  );
  expect(getById(container, 'test_settings')).toBeDefined();

  const closeIcon = getById(container, 'closeIcon')
  fireEvent.click(closeIcon)
  expect(onClose).toHaveBeenCalledTimes(1);

  const btn_cancel = getById(container, 'btn_cancel')
  fireEvent.click(btn_cancel)
  expect(onClose).toHaveBeenCalledTimes(2);

  const name = getById(container, 'name')
  fireEvent.change(name, {target: {value: "name"}})
  expect(name.value).toEqual("name");

  const columns = getById(container, 'columns')
  fireEvent.change(columns, {target: {value: "col3,col4,col5"}})
  expect(columns.value).toEqual("col3,col4,col5");

  const type = getById(container, 'type')
  fireEvent.change(type, {target: {value: "list"}})
  expect(type.value).toEqual("list");

  const btn_ok = getById(container, 'btn_ok')
  fireEvent.click(btn_ok)
  expect(onData).toHaveBeenCalledTimes(1);

})

it('renders in the Disabled state', () => {
  const { container } = render(
    <Disabled {...Disabled.args} id="test_settings" />
  );
  expect(getById(container, 'test_settings')).toBeDefined();
})