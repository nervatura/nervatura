import { render, fireEvent, queryByAttribute } from '@testing-library/react';
import '@testing-library/jest-dom/extend-expect';

import { Default, DateInput, TimeInput } from './DateTime.stories';

const getById = queryByAttribute.bind(null, 'id');

it('renders in the Default state', () => {
  const onChange = jest.fn()

  const { container } = render(
    <Default {...Default.args} id="test_date"
      onChange={onChange} />
  );
  const test_input = getById(container, 'test_date')
  expect(test_input).toBeDefined();

  fireEvent.keyDown(test_input, { key: 'Enter', code: 'Enter', keyCode: 13 })
  expect(onChange).toHaveBeenCalledTimes(1);
  
  test_input.value = ""
  fireEvent.keyDown(test_input, { key: 'Enter', code: 'Enter', keyCode: 13 })
  expect(onChange).toHaveBeenCalledTimes(2);

});

it('renders in the DateInput state', () => {
  const onChange = jest.fn()

  const { container } = render(
    <DateInput {...DateInput.args} id="test_date" 
      onChange={onChange} />
  );
  const test_input = getById(container, 'test_date')
  expect(test_input).toBeDefined();

  fireEvent.change(test_input, {target: {value: test_input.value}})
  fireEvent.keyDown(test_input, { key: 'Enter', code: 'Enter', keyCode: 13 })
  expect(onChange).toHaveBeenCalledTimes(0);

  fireEvent.change(test_input, {target: {value: "2021-12-24"}})
  fireEvent.keyDown(test_input, { key: 'Enter', code: 'Enter', keyCode: 13 })
  expect(onChange).toHaveBeenCalledTimes(1);
  
  fireEvent.change(test_input, {target: {value: ""}})
  fireEvent.keyDown(test_input, { key: 'Enter', code: 'Enter', keyCode: 13 })
  expect(onChange).toHaveBeenCalledTimes(2);

});

it('renders in the TimeInput state', () => {
  const onChange = jest.fn()

  const { container } = render(
    <TimeInput {...TimeInput.args} id="test_date" 
      onChange={onChange}/>
  );
  const test_input = getById(container, 'test_date')
  expect(test_input).toBeDefined();

  fireEvent.keyDown(test_input, { key: 'Enter', code: 'Enter', keyCode: 13 })
  expect(onChange).toHaveBeenCalledTimes(1);

});

it('onChange', () => {
  const { container } = render(
    <TimeInput {...TimeInput.args} id="test_date" value="" />
  );
  const test_input = getById(container, 'test_date')
  expect(test_input).toBeDefined();

  fireEvent.keyDown(test_input, { key: 'Enter', code: 'Enter', keyCode: 13 })

});
