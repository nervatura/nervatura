---
title: Create a new database
type: docs
weight: 10
bookToC: false
---

Configure your API-KEY and database connection in your environment variables: 
```NT_API_KEY``` and ```NT_ALIAS_[ALIASNAME]``` <br />
Connection string form: *adapter://user:password@host/database* or *adapter://user:password@host?database=database*<br />
Supported database adapters: *sqlite3, postgres, mysql, mssql*

For examples:
- *NT_ALIAS_DEMO=sqlite3://data/database/demo.db*
- *NT_ALIAS_PGDEMO=postgres://postgres:password@localhost:5432/nervatura*
- *NT_ALIAS_MYDEMO=mysql://root:password@localhost:3306/nervatura*
- *NT_ALIAS_MSDEMO=mssql://sa:Password1234_1@localhost:1433?database=nervatura*

Create a new database:
```
./nervatura -c DatabaseCreate -k [YOUR_API_KEY] \
  -o "{\"database\":\"[your_lowercase_alias_name]\",\"demo\":false}"
```
You can use the [**ADMIN GUI**](/docs/start/screenshot#service-admin-gui) Database section:

API-KEY: **YOUR_API_KEY**<br />
Alias name: **your_lowercase_alias_name**<br />
Demo database: **false**

The SQLite databases are created automatically. Other types of databases must be created manually before. For testing you can fill in the database with some dummy data (demo=true). If you don't need those later, then you can create a blank database again. **All data in the database will be destroyed!**

Optional: Install Nervatura Report templates to the database
- Login to the database: [**ADMIN GUI**](/docs/start/screenshot#service-admin-gui) <br />
Username: **admin**<br />
Password: **Empty password: Please change after the first login!**<br />
Database: **your_lowercase_alias_name**
- List reports: Returns all installable files from the ```NT_REPORT_DIR``` directory (empty value: all built-in Nervatura Report templates)
- Install a report to the database
