import { render, queryByAttribute, fireEvent } from '@testing-library/react';
import '@testing-library/jest-dom/extend-expect';

import { Default, New, Report, PrintQueue, Notes, Item, Meta, View } from './Editor.stories';

const getById = queryByAttribute.bind(null, 'id');

describe('<Editor />', () => {
  
  beforeAll(() => {
    Object.defineProperty(global.document, 'execCommand', { value: jest.fn() });
  });
  
  afterAll(() => {
    jest.clearAllMocks();
  });

  it('renders in the Default state', () => {
    const onEvent = jest.fn()

    const { container } = render(
      <Default {...Default.args} id="test_editor" onEvent={onEvent} />
    );
    expect(getById(container, "test_editor")).toBeDefined();

    const btn_form = getById(container, 'btn_form')
    fireEvent.click(btn_form)
    expect(onEvent).toHaveBeenCalledTimes(1);

    const btn_fieldvalue = getById(container, 'btn_fieldvalue')
    fireEvent.click(btn_fieldvalue)
    expect(onEvent).toHaveBeenCalledTimes(2);

  })

  it('renders in the New state', () => {
    const { container } = render(
      <New {...New.args} id="test_editor" />
    );
    expect(getById(container, "test_editor")).toBeDefined();

  })

  it('renders in the Report state', () => {
    const { container } = render(
      <Report {...Report.args} id="test_editor" />
    );
    expect(getById(container, "test_editor")).toBeDefined();

  })

  it('renders in the PrintQueue state', () => {
    const { container } = render(
      <PrintQueue {...PrintQueue.args} id="test_editor" />
    );
    expect(getById(container, "test_editor")).toBeDefined();

  })

  it('renders in the Notes state', () => {
    const { container } = render(
      <Notes {...Notes.args} id="test_editor" />
    );
    expect(getById(container, "test_editor")).toBeDefined();

  })

  it('renders in the Item state', () => {
    const { container } = render(
      <Item {...Item.args} id="test_editor" />
    );
    expect(getById(container, "test_editor")).toBeDefined();

  })

  it('renders in the Meta state', () => {
    const { container } = render(
      <Meta {...Meta.args} id="test_editor" />
    );
    expect(getById(container, "test_editor")).toBeDefined();

  })

  it('renders in the View state', () => {
    const { container } = render(
      <View {...View.args} id="test_editor" />
    );
    expect(getById(container, "test_editor")).toBeDefined();

  })
})