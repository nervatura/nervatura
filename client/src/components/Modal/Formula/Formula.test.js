import { render, fireEvent, queryByAttribute } from '@testing-library/react';
import '@testing-library/jest-dom/extend-expect';

import { Default, Disabled } from './Formula.stories';

const getById = queryByAttribute.bind(null, 'id');

it('renders in the Default state', () => {
  const onFormula = jest.fn()
  const onClose = jest.fn()

  const { container } = render(
    <Default {...Default.args} id="test_settings"
    onFormula={onFormula} onClose={onClose} />
  );
  expect(getById(container, 'test_settings')).toBeDefined();

  const closeIcon = getById(container, 'closeIcon')
  fireEvent.click(closeIcon)
  expect(onClose).toHaveBeenCalledTimes(1);

  const btn_cancel = getById(container, 'btn_cancel')
  fireEvent.click(btn_cancel)
  expect(onClose).toHaveBeenCalledTimes(2);

  const btn_formula = getById(container, 'btn_formula')
  fireEvent.click(btn_formula)
  expect(onFormula).toHaveBeenCalledTimes(1);

  const formula = getById(container, 'formula')
  fireEvent.change(formula, {target: {value: "19"}})
  expect(formula.value).toEqual("19");

})

it('renders in the Disabled state', () => {
  const { container } = render(
    <Disabled {...Disabled.args} id="test_settings" />
  );
  expect(getById(container, 'test_settings')).toBeDefined();
})