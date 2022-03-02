import { render, queryByAttribute, fireEvent } from '@testing-library/react';
import '@testing-library/jest-dom/extend-expect';

import { Default, Empty, ReadOnly } from './Note.stories';

const getById = queryByAttribute.bind(null, 'id');

describe('Edit - Note', () => {
  beforeAll(() => {
    Object.defineProperty(global.document, 'execCommand', { value: jest.fn() });
  });
  
  afterAll(() => {
    jest.clearAllMocks();
  });

  it('renders in the Default state', () => {
    const onEvent = jest.fn()

    const { container } = render(
      <Default {...Default.args} id="test_note" onEvent={onEvent} />
    );
    expect(getById(container, "test_note")).toBeDefined();

    const btn_pattern_default = getById(container, 'btn_pattern_default')
    fireEvent.click(btn_pattern_default)
    expect(onEvent).toHaveBeenCalledTimes(1);

    const btn_pattern_load = getById(container, 'btn_pattern_load')
    fireEvent.click(btn_pattern_load)
    expect(onEvent).toHaveBeenCalledTimes(2);

    const btn_pattern_save = getById(container, 'btn_pattern_save')
    fireEvent.click(btn_pattern_save)
    expect(onEvent).toHaveBeenCalledTimes(3);

    const btn_pattern_new = getById(container, 'btn_pattern_new')
    fireEvent.click(btn_pattern_new)
    expect(onEvent).toHaveBeenCalledTimes(4);

    const btn_pattern_delete = getById(container, 'btn_pattern_delete')
    fireEvent.click(btn_pattern_delete)
    expect(onEvent).toHaveBeenCalledTimes(5);

    const sel_pattern = getById(container, 'sel_pattern')
    fireEvent.change(sel_pattern, {target: {value: "2"}})
    expect(onEvent).toHaveBeenCalledTimes(6);

  })

  it('renders in the Empty state', () => {
    const { container } = render(
      <Empty {...Empty.args} id="test_note" />
    );
    expect(getById(container, "test_note")).toBeDefined();
  })

  it('renders in the ReadOnly state', () => {
    const { container } = render(
      <ReadOnly {...ReadOnly.args} id="test_note" />
    );
    expect(getById(container, "test_note")).toBeDefined();
  })

})