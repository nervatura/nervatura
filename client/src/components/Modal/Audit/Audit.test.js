import { render, fireEvent, queryByAttribute } from '@testing-library/react';
import '@testing-library/jest-dom/extend-expect';

import { Default, Existing, Disabled } from './Audit.stories';

const getById = queryByAttribute.bind(null, 'id');

it('renders in the Default state', () => {
  const onAudit = jest.fn()
  const onClose = jest.fn()

  const { container } = render(
    <Default {...Default.args} id="test_settings"
      onAudit={onAudit} onClose={onClose} />
  );
  expect(getById(container, 'test_settings')).toBeDefined();

  const closeIcon = getById(container, 'closeIcon')
  fireEvent.click(closeIcon)
  expect(onClose).toHaveBeenCalledTimes(1);

  const btn_cancel = getById(container, 'btn_cancel')
  fireEvent.click(btn_cancel)
  expect(onClose).toHaveBeenCalledTimes(2);

  const nervatype = getById(container, 'nervatype')
  fireEvent.change(nervatype, {target: {value: "18"}})
  expect(nervatype.value).toEqual("18");

  fireEvent.change(nervatype, {target: {value: "28"}})
  expect(nervatype.value).toEqual("28");

  fireEvent.change(nervatype, {target: {value: "10"}})
  expect(nervatype.value).toEqual("10");

  const inputfilter = getById(container, 'inputfilter')
  fireEvent.change(inputfilter, {target: {value: "106"}})
  expect(inputfilter.value).toEqual("106");

  const supervisor = getById(container, 'supervisor')
  fireEvent.click(supervisor)
  expect(getById(container, 'check_0')).toBeDefined();
  fireEvent.click(supervisor)
  expect(getById(container, 'check_1')).toBeDefined();

  const btn_ok = getById(container, 'btn_ok')
  fireEvent.click(btn_ok)
  expect(onAudit).toHaveBeenCalledTimes(1);

})

it('renders in the Existing state', () => {
  const onAudit = jest.fn()
  const { container } = render(
    <Existing {...Existing.args} id="test_settings" onAudit={onAudit} />
  );
  expect(getById(container, 'test_settings')).toBeDefined();

  const subtype = getById(container, 'subtype')
  fireEvent.change(subtype, {target: {value: "58"}})
  expect(subtype.value).toEqual("58");

  const btn_ok = getById(container, 'btn_ok')
  fireEvent.click(btn_ok)
  expect(onAudit).toHaveBeenCalledTimes(1);
})

it('renders in the Disabled state', () => {
  const { container } = render(
    <Disabled {...Disabled.args} id="test_settings" />
  );
  expect(getById(container, 'test_settings')).toBeDefined();

})