import { render, fireEvent, queryByAttribute } from '@testing-library/react';
import '@testing-library/jest-dom/extend-expect';

import { Default, Disabled } from './Report.stories';

const getById = queryByAttribute.bind(null, 'id');

it('renders in the Default state', () => {
  const onOutput = jest.fn()
  const onClose = jest.fn()

  const { container } = render(
    <Default {...Default.args} id="test_settings"
    onOutput={onOutput} onClose={onClose} />
  );
  expect(getById(container, 'test_settings')).toBeDefined();

  const closeIcon = getById(container, 'closeIcon')
  fireEvent.click(closeIcon)
  expect(onClose).toHaveBeenCalledTimes(1);

  const btn_print = getById(container, 'btn_print')
  fireEvent.click(btn_print)
  expect(onOutput).toHaveBeenCalledTimes(1);

  const btn_pdf = getById(container, 'btn_pdf')
  fireEvent.click(btn_pdf)
  expect(onOutput).toHaveBeenCalledTimes(2);

  const btn_xml = getById(container, 'btn_xml')
  fireEvent.click(btn_xml)
  expect(onOutput).toHaveBeenCalledTimes(3);

  const btn_printqueue = getById(container, 'btn_printqueue')
  fireEvent.click(btn_printqueue)
  expect(onOutput).toHaveBeenCalledTimes(4);

  const template = getById(container, 'template')
  fireEvent.change(template, {target: {value: "ntr_invoice_fi"}})
  expect(template.value).toEqual("ntr_invoice_fi");

  const orient = getById(container, 'orient')
  fireEvent.change(orient, {target: {value: "landscape"}})
  expect(orient.value).toEqual("landscape");

  const size = getById(container, 'size')
  fireEvent.change(size, {target: {value: "a5"}})
  expect(size.value).toEqual("a5");

  const copy = getById(container, 'copy')
  fireEvent.change(copy, {target: {value: "2"}})
  expect(copy.value).toEqual("2");

})

it('renders in the Disabled state', () => {
  const { container } = render(
    <Disabled {...Disabled.args} id="test_settings" />
  );
  expect(getById(container, 'test_settings')).toBeDefined();
})