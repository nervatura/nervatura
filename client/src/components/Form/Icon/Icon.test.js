import { render, queryByAttribute } from '@testing-library/react';

import '@testing-library/jest-dom/extend-expect';

import { Default, ColorPointer } from './Icon.stories';

const getById = queryByAttribute.bind(null, 'id');

it('renders in the Default state', () => {

  const { container } = render(
    <Default {...Default.args} id="default" />
  );
  expect(getById(container, 'default')).toBeDefined();

});

it('renders in the ColorPointer state', () => {

  const { container } = render(
    <ColorPointer {...ColorPointer.args} id="color" />
  );
  expect(getById(container, 'color')).toBeDefined();

});