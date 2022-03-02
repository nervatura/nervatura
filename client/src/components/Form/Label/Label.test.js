import { render, queryByAttribute } from '@testing-library/react';

import '@testing-library/jest-dom/extend-expect';

import { Default, LeftIcon, RightIcon, Centered } from './Label.stories';

const getById = queryByAttribute.bind(null, 'id');

it('renders in the Default state', () => {

  const { container } = render(
    <Default {...Default.args} id="test_label" />
  );
  expect(getById(container, 'test_label')).toBeDefined();

});

it('renders in the LeftIcon state', () => {

  const { container } = render(
    <LeftIcon {...LeftIcon.args} id="test_label" />
  );
  expect(getById(container, 'test_label')).toBeDefined();

});

it('renders in the RightIcon state', () => {

  const { container } = render(
    <RightIcon {...RightIcon.args} id="test_label" />
  );
  expect(getById(container, 'test_label')).toBeDefined();

});

it('renders in the Centered state', () => {

  const { container } = render(
    <Centered {...Centered.args} id="test_label" />
  );
  expect(getById(container, 'test_label')).toBeDefined();

});