import { render, fireEvent, queryByAttribute } from '@testing-library/react';
import '@testing-library/jest-dom/extend-expect';

import { Default, Placeholder } from './Select.stories';

const getById = queryByAttribute.bind(null, 'id');

it('renders in the Default state', () => {
  const onChange = jest.fn()

  const { container } = render(
    <Default {...Default.args} id="test_select"
    onChange={onChange} />
  );
  expect(getById(container, 'test_select')).toBeDefined();

  const test_select = getById(container, 'test_select')

  fireEvent.change(test_select, {target: {value: "value2"}})
  expect(onChange).toHaveBeenCalledTimes(1);

});

it('renders in the Placeholder state', () => {
  const { container } = render(
    <Placeholder {...Placeholder.args} id="test_select" />
  );
  expect(getById(container, 'test_select')).toBeDefined();

});