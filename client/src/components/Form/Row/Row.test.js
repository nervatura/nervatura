import { render, fireEvent, queryByAttribute } from '@testing-library/react';
import '@testing-library/jest-dom/extend-expect';

import { Default, FlipStringOn, FlipStringOff, FlipTextOn, FlipTextOff, FlipImageOn, FlipImageOff,
  FlipChecklistOn, FlipChecklistOff, Field, Reportfield, ReportfieldEmpty, Fieldvalue,
  Col2, Col3, Col4, Missing, Label } from './Row.stories';

const getById = queryByAttribute.bind(null, 'id');

it('renders in the Default state', () => {
  const { container } = render(
    <Default {...Default.args} id="test_input" />
  );
  expect(getById(container, 'test_input')).toBeDefined();

});

it('renders in the Label state', () => {
  const { container } = render(
    <Label {...Label.args} id="test_input" />
  );
  expect(getById(container, 'test_input')).toBeDefined();

});

it('renders in the FlipStringOn state', () => {
  const onEdit = jest.fn()

  const { container } = render(
    <FlipStringOn {...FlipStringOn.args} id="test_input" onEdit={onEdit} />
  );
  expect(getById(container, 'test_input')).toBeDefined();

  const checkbox = getById(container, 'checkbox_title')
  fireEvent.click(checkbox)
  expect(onEdit).toHaveBeenCalledTimes(1);

});

it('renders in the FlipStringOff state', () => {
  const { container } = render(
    <FlipStringOff {...FlipStringOff.args} id="test_input" />
  );
  expect(getById(container, 'test_input')).toBeDefined();

});

it('renders in the FlipTextOn state', () => {
  const { container } = render(
    <FlipTextOn {...FlipTextOn.args} id="test_input" />
  );
  expect(getById(container, 'test_input')).toBeDefined();

});

it('renders in the FlipTextOff state', () => {
  const { container } = render(
    <FlipTextOff {...FlipTextOff.args} id="test_input" />
  );
  expect(getById(container, 'test_input')).toBeDefined();

});

it('renders in the FlipImageOn state', () => {
  const onEdit = jest.fn()

  const { container } = render(
    <FlipImageOn {...FlipImageOn.args} id="test_input" onEdit={onEdit} />
  );
  expect(getById(container, 'test_input')).toBeDefined();

  const input_src = getById(container, 'input_src')
  fireEvent.change(input_src, {target: {value: "data:image"}})
  expect(onEdit).toHaveBeenCalledTimes(1);

  const file_src = getById(container, 'file_src')
  fireEvent.change(file_src, {target: {value: ""}})
  expect(onEdit).toHaveBeenCalledTimes(2);

  render(
    <FlipImageOn {...FlipImageOn.args} 
      values={{ src: "" }} />
  )

  render(
    <FlipImageOn {...FlipImageOn.args} 
      values={{ src: "data:image" }} />
  )

  render(
    <FlipImageOn {...FlipImageOn.args} 
      values={{ src: "data" }} />
  )

});

it('renders in the FlipImageOff state', () => {
  const { container } = render(
    <FlipImageOff {...FlipImageOff.args} id="test_input" />
  );
  expect(getById(container, 'test_input')).toBeDefined();

});

it('renders in the FlipChecklistOn state', () => {
  const onEdit = jest.fn()

  const { container } = render(
    <FlipChecklistOn {...FlipChecklistOn.args} id="test_input" onEdit={onEdit} />
  );
  expect(getById(container, 'test_input')).toBeDefined();

  const checkbox = getById(container, 'checklist_border_0')
  fireEvent.click(checkbox)
  expect(onEdit).toHaveBeenCalledTimes(1);

});

it('renders in the FlipChecklistOff state', () => {
  const { container } = render(
    <FlipChecklistOff {...FlipChecklistOff.args} id="test_input" />
  );
  expect(getById(container, 'test_input')).toBeDefined();

});

it('renders in the Field state', () => {
  const { container } = render(
    <Field {...Field.args} id="test_input" />
  );
  expect(getById(container, 'test_input')).toBeDefined();

});

it('renders in the Reportfield state', () => {
  const onEdit = jest.fn()

  const { container } = render(
    <Reportfield {...Reportfield.args} id="test_input" onEdit={onEdit} />
  );
  expect(getById(container, 'test_input')).toBeDefined();

  const cb_posdate = getById(container, 'cb_posdate')
  fireEvent.click(cb_posdate)
  expect(onEdit).toHaveBeenCalledTimes(0);

});

it('renders in the ReportfieldEmpty state', () => {
  const onEdit = jest.fn()

  const { container } = render(
    <ReportfieldEmpty {...ReportfieldEmpty.args} id="test_input" onEdit={onEdit} />
  );
  expect(getById(container, 'test_input')).toBeDefined();

  const cb_curr = getById(container, 'cb_curr')
  fireEvent.click(cb_curr)
  expect(onEdit).toHaveBeenCalledTimes(1);

});

it('renders in the Fieldvalue state', () => {
  const onEdit = jest.fn()

  const { container } = render(
    <Fieldvalue {...Fieldvalue.args} id="test_input" onEdit={onEdit} />
  );
  expect(getById(container, 'test_input')).toBeDefined();

  const checkbox = getById(container, 'delete_sample_customer_float')
  fireEvent.click(checkbox)
  expect(onEdit).toHaveBeenCalledTimes(1);

  const notes = getById(container, 'notes_sample_customer_float')
  fireEvent.change(notes, {target: {value: "value"}})
  expect(onEdit).toHaveBeenCalledTimes(2);

  render(
    <Fieldvalue {...Fieldvalue.args} id="test_input" 
      data={{
        dataset: {}, 
        current: {}, 
        audit: "readonly"}} />
  );

});

it('renders in the Col2 state', () => {
  const { container } = render(
    <Col2 {...Col2.args} id="test_input" />
  );
  expect(getById(container, 'test_input')).toBeDefined();

});

it('renders in the Col3 state', () => {
  const { container } = render(
    <Col3 {...Col3.args} id="test_input" />
  );
  expect(getById(container, 'test_input')).toBeDefined();

});

it('renders in the Col4 state', () => {
  const { container } = render(
    <Col4 {...Col4.args} id="test_input" />
  );
  expect(getById(container, 'test_input')).toBeDefined();

});

it('renders in the Missing state', () => {
  const { container } = render(
    <Missing {...Missing.args} id="test_input" />
  );
  expect(getById(container, 'test_input')).toBeDefined();

});