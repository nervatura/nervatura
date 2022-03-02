import { render, fireEvent, queryByAttribute } from '@testing-library/react';
import '@testing-library/jest-dom/extend-expect';

import { Default, InputValue, DefaultCancel } from './InputBox.stories';

const getById = queryByAttribute.bind(null, 'id');

it('renders in the Default state', () => {
  const onOK = jest.fn()
  const onCancel = jest.fn()

  const { container } = render(
    <Default {...Default.args} id="test_input"
      onOK={onOK} onCancel={onCancel} />
  );
  expect(getById(container, 'test_input')).toBeDefined();

  const btn_ok = getById(container, 'btn_ok')
  fireEvent.click(btn_ok)
  expect(onOK).toHaveBeenCalledTimes(1);

  const btn_cancel = getById(container, 'btn_cancel')
  fireEvent.click(btn_cancel)
  expect(onCancel).toHaveBeenCalledTimes(1);

})

it('renders in the InputValue state', () => {
  const onOK = jest.fn()

  const { container } = render(
    <InputValue {...InputValue.args} id="test_input" onOK={onOK} />
  );
  expect(getById(container, 'test_input')).toBeDefined();

  const input_value = getById(container, 'input_value')
  fireEvent.change(input_value, {target: {value: "value"}})
  expect(input_value.value).toEqual("value");

  fireEvent.keyDown(input_value, { key: 'Enter', code: 'Enter', keyCode: 13 })
  expect(onOK).toHaveBeenCalledTimes(1)

})

it('renders in the DefaultCancel state', () => {
  const { container } = render(
    <DefaultCancel {...DefaultCancel.args} id="test_input" />
  );
  expect(getById(container, 'test_input')).toBeDefined();

})