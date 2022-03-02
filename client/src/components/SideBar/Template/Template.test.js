import { render, queryByAttribute, fireEvent } from '@testing-library/react';
import '@testing-library/jest-dom/extend-expect';

import { Default, Sample, Dirty } from './Template.stories';

const getById = queryByAttribute.bind(null, 'id');

it('renders in the Default state', () => {
  const onEvent = jest.fn()

  const { container } = render(
    <Default {...Default.args} id="test_menu" onEvent={onEvent} />
  );
  expect(getById(container, "test_menu")).toBeDefined();

  const cmd_back = getById(container, 'cmd_back')
  fireEvent.click(cmd_back)
  expect(onEvent).toHaveBeenCalledTimes(1);

})

it('renders in the Sample state', () => {

  const { container } = render(
    <Sample {...Sample.args} id="test_menu" />
  );
  expect(getById(container, "test_menu")).toBeDefined();

})

it('renders in the Dirty state', () => {

  const { container } = render(
    <Dirty {...Dirty.args} id="test_menu" />
  );
  expect(getById(container, "test_menu")).toBeDefined();

})