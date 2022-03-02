import { render, queryByAttribute, fireEvent } from '@testing-library/react';
import '@testing-library/jest-dom/extend-expect';

import { Default, NewItem, NewPayment, NewMovement, NewResource, Document, 
  DocumentDeleted, DocumentCancellation, DocumentClosed, DocumentReadonly, DocumentNoOptions } from './Edit.stories';

const getById = queryByAttribute.bind(null, 'id');

it('renders in the Default state', () => {
  const onEvent = jest.fn()

  const { container } = render(
    <Default {...Default.args} id="test_menu" onEvent={onEvent} />
  );
  expect(getById(container, "test_menu")).toBeDefined();

  const btn_view = getById(container, 'state_edit')
  fireEvent.click(btn_view)
  expect(onEvent).toHaveBeenCalledTimes(1);

  const btn_group = getById(container, 'new_transitem_group')
  fireEvent.click(btn_group)
  expect(onEvent).toHaveBeenCalledTimes(2);

})

it('renders in the NewItem state', () => {

  const { container } = render(
    <NewItem {...NewItem.args} id="test_menu" />
  );
  expect(getById(container, "test_menu")).toBeDefined();

})

it('renders in the NewPayment state', () => {

  const { container } = render(
    <NewPayment {...NewPayment.args} id="test_menu" />
  );
  expect(getById(container, "test_menu")).toBeDefined();

})

it('renders in the NewMovement state', () => {

  const { container } = render(
    <NewMovement {...NewMovement.args} id="test_menu" />
  );
  expect(getById(container, "test_menu")).toBeDefined();

})

it('renders in the NewResource state', () => {

  const { container } = render(
    <NewResource {...NewResource.args} id="test_menu" />
  );
  expect(getById(container, "test_menu")).toBeDefined();

})

it('renders in the Document state', () => {

  const { container } = render(
    <Document {...Document.args} id="test_menu" />
  );
  expect(getById(container, "test_menu")).toBeDefined();

})

it('renders in the DocumentDeleted state', () => {

  const { container } = render(
    <DocumentDeleted {...DocumentDeleted.args} id="test_menu" />
  );
  expect(getById(container, "test_menu")).toBeDefined();

})

it('renders in the DocumentCancellation state', () => {

  const { container } = render(
    <DocumentCancellation {...DocumentCancellation.args} id="test_menu" />
  );
  expect(getById(container, "test_menu")).toBeDefined();

})

it('renders in the DocumentClosed state', () => {

  const { container } = render(
    <DocumentClosed {...DocumentClosed.args} id="test_menu" />
  );
  expect(getById(container, "test_menu")).toBeDefined();

})

it('renders in the DocumentReadonly state', () => {

  const { container } = render(
    <DocumentReadonly {...DocumentReadonly.args} id="test_menu" />
  );
  expect(getById(container, "test_menu")).toBeDefined();

})

it('renders in the DocumentNoOptions state', () => {

  const { container } = render(
    <DocumentNoOptions {...DocumentNoOptions.args} id="test_menu" />
  );
  expect(getById(container, "test_menu")).toBeDefined();

})