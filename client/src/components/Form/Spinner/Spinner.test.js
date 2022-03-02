import { render, queryByAttribute } from '@testing-library/react';

import '@testing-library/jest-dom/extend-expect';

import { Default } from './Spinner.stories';

const getById = queryByAttribute.bind(null, 'id');

it('renders in the Default state', () => {

  const { container } = render(
    <Default {...Default.args} />
  );
  expect(getById(container, 'app_loading')).toBeDefined();

});