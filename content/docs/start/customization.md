---
title: Customization
type: docs
weight: 20
bookToC: true
---

## Customize the appearance

1. Nervatura Client language translation
- Create a file based on the [client_config.json](https://github.com/nervatura/nervatura/tree/master/dist) file. All subtitles [can be found here](https://raw.githubusercontent.com/nervatura/nervatura-client/master/src/config/locales.js).
- Set ```NT_CLIENT_CONFIG``` environment variable value to *YOUR_CLIENT_CONFIG_FILE_PATH*
- Docker container: mount local folder to the container

2. Nervatura Client custom remote functions: [**MENU SHORTCUTS**](/docs/client/settings/uimenu)

3. Custom PDF Report font
- Set ```NT_FONT_FAMILY``` environment variable value to *YOUR_FONT_FAMILY_NAME*
- Set ```NT_FONT_DIR``` environment variable value to *YOUR_FONTS_PATH*
- Valid font type and filename form: FAMILY_NAME-Regular.ttf, FAMILY_NAME-Italic.ttf, FAMILY_NAME-Bold.ttf, FAMILY_NAME-BoldItalic.ttf
- Docker container: mount local folder to the container

4. Modify installed Nervatura report definitions: [**REPORT EDITOR**](/docs/client/program/editor)

## Bearer Authentication

Environment variables: [.env.example](https://github.com/nervatura/nervatura-service/blob/master/.env.example)<br />
User authentication is based on the *employee.username* or *customer.custnumber* fields. The identifier can be the following types: username (employee), email, phone number (customer).<br />
Passwords are not stored in the employee or customer tables. They are anonymized and stored in a unique table with [strong encryption](https://github.com/P-H-C/phc-winner-argon2).

External authorization: ```NT_TOKEN_PUBLIC_KEY_TYPE```, ```NT_TOKEN_PUBLIC_KEY_URL```

## Other Recipes

- [CHANGELOG](https://raw.githubusercontent.com/nervatura/nervatura-service/master/CHANGELOG
)
- [![GoDoc](https://godoc.org/github.com/nervatura/nervatura-service?status.svg)](https://godoc.org/github.com/nervatura/nervatura-service)
- [gRPC API proto file](https://github.com/nervatura/nervatura/tree/master/dist)
- [Python gRPC packages](https://pypi.org/project/nervatura/)
- [Report templates files](https://github.com/nervatura/nervatura-service/tree/master/pkg/utils/static/templates)
- [Node.js sample application](https://github.com/nervatura/nervatura-express)
- [Python sample application](https://github.com/nervatura/nervatura-fastapi)
