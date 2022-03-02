import { render, queryByAttribute, fireEvent } from '@testing-library/react';
import '@testing-library/jest-dom/extend-expect';

import { Default, DatabaseGroup, UserGroup, FormItemAll, FormItemNew, FormItemRead } from './Setting.stories';
import { SIDE_VISIBILITY } from "./Setting";

const getById = queryByAttribute.bind(null, 'id');

it('renders in the Default state', () => {
  const onEvent = jest.fn()

  const { container } = render(
    <Default {...Default.args} id="test_menu" onEvent={onEvent} />
  );
  expect(getById(container, "test_menu")).toBeDefined();

  const btn_numberdef = getById(container, 'cmd_numberdef')
  fireEvent.click(btn_numberdef)
  expect(onEvent).toHaveBeenCalledTimes(1);

  const btn_database = getById(container, 'group_database')
  fireEvent.click(btn_database)
  expect(onEvent).toHaveBeenCalledTimes(2);

  render(
    <Default {...Default.args} 
      auditFilter={{
        setting: ["all", 1],
        audit: ["disabled", 1]
      }} />
  )

})

it('renders in the DatabaseGroup state', () => {

  const { container } = render(
    <DatabaseGroup {...DatabaseGroup.args} id="test_menu" />
  );
  expect(getById(container, "test_menu")).toBeDefined();

})

it('renders in the UserGroup state', () => {

  const { container } = render(
    <UserGroup {...UserGroup.args} id="test_menu" />
  );
  expect(getById(container, "test_menu")).toBeDefined();

})

it('renders in the FormItemAll state', () => {

  const { container } = render(
    <FormItemAll {...FormItemAll.args} id="test_menu" side={SIDE_VISIBILITY.SHOW} />
  );
  expect(getById(container, "test_menu")).toBeDefined();

})

it('renders in the FormItemNew state', () => {

  const { container } = render(
    <FormItemNew {...FormItemNew.args} id="test_menu" />
  );
  expect(getById(container, "test_menu")).toBeDefined();

})

it('renders in the FormItemRead state', () => {

  const { container } = render(
    <FormItemRead {...FormItemRead.args} id="test_menu" />
  );
  expect(getById(container, "test_menu")).toBeDefined();

})