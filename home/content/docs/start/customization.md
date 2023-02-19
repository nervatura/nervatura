---
title: Customization
type: docs
weight: 20
bookToC: true
---

## Customize the appearance

1. Nervatura Client language translation
- Create a file based on the [client_config.json](https://raw.githubusercontent.com/nervatura/nervatura/master/service/data/client_config.json) file. All subtitles [can be found here](https://raw.githubusercontent.com/nervatura/nervatura/master/client/src/config/locales.js).
- The **recommended and easy** way to create and edit the language file is the ADMIN GUI [**Translation Helper Tool**](/docs/start/screenshot/#service-admin-gui).
- For more help on setting up and using the language file, see the [**Examples**](/docs/start/examples/#nervatura-client-language-translation) section
- Set ```NT_CLIENT_CONFIG``` environment variable value to *YOUR_CLIENT_CONFIG_FILE.JSON*
- Docker container: mount local folder to the container

2. Nervatura Client custom remote functions: [**MENU SHORTCUTS**](/docs/client/settings/uimenu)

3. Custom PDF Report font
- Set ```NT_FONT_FAMILY``` environment variable value to *YOUR_FONT_FAMILY_NAME*
- Set ```NT_FONT_DIR``` environment variable value to *YOUR_FONTS_PATH*
- Valid font type and filename form: FAMILY_NAME-Regular.ttf, FAMILY_NAME-Italic.ttf, FAMILY_NAME-Bold.ttf, FAMILY_NAME-BoldItalic.ttf
- Docker container: mount local folder to the container

4. Modify installed Nervatura report definitions: [**REPORT EDITOR**](/docs/client/program/editor)

5. The basic report definition files [can be found here](https://github.com/nervatura/nervatura/tree/master/service/pkg/utils/static/templates).

## Bearer Authentication

Environment variables: [.env.example](https://raw.githubusercontent.com/nervatura/nervatura/master/service/.env.example)

User authentication is based on the *employee.username* or *employee.registration_key* fields. <br />
To match the username, Nervatura login processing associates the token fields in the 
following order: *login*.***username*** = *token*.***username*** || *token*.***user_id*** || 
*token*.***sub*** || *token*.***email***

Passwords are not stored in the employee or customer tables. They are anonymized and stored in a unique table with [strong encryption](https://github.com/P-H-C/phc-winner-argon2).

Custom token-based, password-free login can also be used. Password-based login can be enabled or disabled with the ```NT_PASSWORD_LOGIN``` option value.

**You can find more information about this in the** [**Examples**](/docs/start/examples).
