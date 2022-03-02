import { render, fireEvent, queryByAttribute } from '@testing-library/react';
import '@testing-library/jest-dom/extend-expect';

import { Default, NumberInput, IntegerInput, CommaSeparator, Disabled } from './Input.stories';

const getById = queryByAttribute.bind(null, 'id');

it('renders in the Default state', () => {
  const onChange = jest.fn()
  const onBlur = jest.fn()

  const { container } = render(
    <Default {...Default.args} id="test_input"
    onChange={onChange} onBlur={onBlur} />
  );
  expect(getById(container, 'test_input')).toBeDefined();

  const test_input = getById(container, 'test_input')

  fireEvent.change(test_input, {target: {value: "change"}})
  expect(onChange).toHaveBeenCalledTimes(1);

  fireEvent.blur(test_input, {target: {value: "blur"}})
  expect(onBlur).toHaveBeenCalledTimes(1);

});

it('renders in the NumberInput state', () => {
  const onChange = jest.fn()
  const onBlur = jest.fn()
  const onEnter = jest.fn()

  const { container } = render(
    <NumberInput {...NumberInput.args} id="test_input"
    onChange={onChange} onBlur={onBlur} onEnter={onEnter} />
  );
  expect(getById(container, 'test_input')).toBeDefined();

  const test_input = getById(container, 'test_input')

  fireEvent.change(test_input, {target: {value: "123.12"}})
  expect(onChange).toHaveBeenCalledTimes(1);

  fireEvent.change(test_input, {target: {value: "123."}})
  expect(onChange).toHaveBeenCalledTimes(2);

  fireEvent.blur(test_input, {target: {value: "123.45"}})
  expect(onBlur).toHaveBeenCalledTimes(1);
  
  fireEvent.keyDown(test_input, { key: 'Enter', code: 'Enter', keyCode: 13 })
  expect(onEnter).toHaveBeenCalledTimes(1)

  fireEvent.keyDown(test_input, { key: 'Enter', code: 'Enter', keyCode: 66 })
  expect(onEnter).toHaveBeenCalledTimes(1)

});

it('renders in the IntegerInput state', () => {
  const onChange = jest.fn()
  const onBlur = jest.fn()

  const { container } = render(
    <IntegerInput {...IntegerInput.args} id="test_input" minValue={10} maxValue={100}
    onChange={onChange} onBlur={onBlur} />
  );
  expect(getById(container, 'test_input')).toBeDefined();

  const test_input = getById(container, 'test_input')

  fireEvent.change(test_input, {target: {value: "123"}})
  expect(onChange).toHaveBeenCalledTimes(1);

  fireEvent.change(test_input, {target: {value: "123.55"}})
  expect(onChange).toHaveBeenCalledTimes(2);

  fireEvent.change(test_input, {target: {value: ""}})
  expect(onChange).toHaveBeenCalledTimes(3);

  fireEvent.blur(test_input, {target: {value: "125"}})
  expect(onBlur).toHaveBeenCalledTimes(1);

  fireEvent.blur(test_input, {target: {value: "5"}})
  expect(onBlur).toHaveBeenCalledTimes(2);

});

it('renders in the CommaSeparator state', () => {
  const onChange = jest.fn()

  const { container } = render(
    <CommaSeparator {...CommaSeparator.args} id="test_input"
    onChange={onChange} />
  );
  expect(getById(container, 'test_input')).toBeDefined();

  const test_input = getById(container, 'test_input')
  fireEvent.blur(test_input, {target: {value: "5a"}})
  expect(onChange).toHaveBeenCalledTimes(1);

});

it('renders in the Disabled state', () => {

  const { container } = render(
    <Disabled {...Disabled.args} id="test_input" />
  );
  expect(getById(container, 'test_input')).toBeDefined();

});
