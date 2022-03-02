import { render, queryByAttribute, fireEvent } from '@testing-library/react';
import '@testing-library/jest-dom/extend-expect';

import { Default, Report } from './Main.stories';

const getById = queryByAttribute.bind(null, 'id');

it('renders in the Default state', () => {
  const onEvent = jest.fn()

  const { container } = render(
    <Default {...Default.args} id="test_editor" onEvent={onEvent} />
  );
  expect(getById(container, "test_editor")).toBeDefined();

  //onSelector
  const sel_show_customer_id = getById(container, 'sel_show_customer_id')
  fireEvent.click(sel_show_customer_id)
  expect(onEvent).toHaveBeenCalledTimes(1);

  //editItem
  const field_curr = getById(container, 'field_curr')
  fireEvent.change(field_curr, {target: {value: "USD"}})
  expect(onEvent).toHaveBeenCalledTimes(2);

})

it('renders in the Report state', () => {
  const { container } = render(
    <Report {...Report.args} id="test_editor" />
  );
  expect(getById(container, "test_editor")).toBeDefined();

})