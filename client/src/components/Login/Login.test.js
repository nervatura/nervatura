import { render, fireEvent, queryByAttribute } from '@testing-library/react';
import '@testing-library/jest-dom/extend-expect';

import { Default, DarkLogin } from './Login.stories';

const getById = queryByAttribute.bind(null, 'id');

it('renders in the Default state', () => {
  const setLocale = jest.fn()
  const changeData = jest.fn()
  const setTheme = jest.fn()

  const { container } = render(
    <Default {...Default.args} id="test_login"
    changeData={changeData}  setLocale={setLocale} setTheme={setTheme} />
  );
  expect(getById(container, 'test_login')).toBeDefined();

  const username = getById(container, 'username')
  fireEvent.change(username, {target: {value: "username"}})
  expect(changeData).toHaveBeenCalledTimes(1);

  const password = getById(container, 'password')
  fireEvent.change(password, {target: {value: "password"}})
  expect(changeData).toHaveBeenCalledTimes(2);

  const database = getById(container, 'database')
  fireEvent.change(database, {target: {value: "database"}})
  expect(changeData).toHaveBeenCalledTimes(3);

  const server = getById(container, 'server')
  fireEvent.change(server, {target: {value: "server"}})
  expect(changeData).toHaveBeenCalledTimes(4);

  const sb_lang = getById(container, 'lang')
  fireEvent.change(sb_lang, {target: {value: "jp"}})
  expect(setLocale).toHaveBeenCalledTimes(1);

  const cmd_theme = getById(container, 'theme')
  fireEvent.click(cmd_theme)
  expect(setTheme).toHaveBeenCalledTimes(1);

});

it('renders in the DarkLogin state', () => {
  const onLogin = jest.fn()

  const { container } = render(
    <DarkLogin {...DarkLogin.args} id="test_login"
    onLogin={onLogin} />
  );
  expect(getById(container, 'test_login')).toBeDefined();

  const cmd_login = getById(container, 'login')
  fireEvent.click(cmd_login)
  expect(onLogin).toHaveBeenCalledTimes(1);

});