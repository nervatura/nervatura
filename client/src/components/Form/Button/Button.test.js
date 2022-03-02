import { render, fireEvent, queryByAttribute } from '@testing-library/react';

import '@testing-library/jest-dom/extend-expect';

import { Default, PrimaryIconLabel, IconBorderButton, Disabled, SmallButton } from './Button.stories';

const getById = queryByAttribute.bind(null, 'id');

it('renders in the Default state', () => {
  const onClick = jest.fn()

  const { container } = render(
    <Default {...Default.args} id="test_button" onClick={onClick} />
  );
  expect(getById(container, 'test_button')).toBeDefined();

  const test_button = getById(container, 'test_button')

  fireEvent.click(test_button)
  expect(onClick).toHaveBeenCalledTimes(1);

});

it('renders in the PrimaryIconLabel state', () => {

  const { container } = render(
    <PrimaryIconLabel {...PrimaryIconLabel.args} id="test_button" />
  );
  expect(getById(container, 'test_button')).toBeDefined();

});

it('renders in the IconBorderButton state', () => {

  const { container } = render(
    <IconBorderButton {...IconBorderButton.args} id="test_button" />
  );
  expect(getById(container, 'test_button')).toBeDefined();

});

it('renders in the Disabled state', () => {

  const { container } = render(
    <IconBorderButton {...Disabled.args} id="test_button" />
  );
  expect(getById(container, 'test_button')).toBeDefined();

});

it('renders in the SmallButton state', () => {

  const { container } = render(
    <SmallButton {...SmallButton.args} id="test_button" />
  );
  expect(getById(container, 'test_button')).toBeDefined();

});