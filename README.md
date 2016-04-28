Nervatura
=========

Open Source Business Management Framework for [Node.js](http://nodejs.org)

## Features

The purpose of Nervatura Project is to develop such data model, tools and knowledge database which allows to build safe and flexible management solutions quickly and at low cost.
The Nervatura Project includes:
  * [Object Model](https://github.com/nervatura/nervatura/wiki/Nervatura-Object-Model-%28NOM%29)
  * [Application Server and Admin GUI](https://github.com/nervatura/nervatura/wiki/Application-Server-and-Admin-GUI-%28NAS%29)
  * [Data Interface](https://github.com/nervatura/nervatura/wiki/Nervatura-Data-Interface-%28NDI%29)
  * [Programming Interface](https://github.com/nervatura/nervatura/wiki/Nervatura-Programming-Interface-%28NPI%29)
  * [Client- and server-side data reporting](https://github.com/nervatura/nervatura/wiki/Nervatura-Report-%28NDR%29)
  * and other documents, resources, sample code and demo programs

The development are open-source project and can be used freely under the scope of [LGPLv3 License](http://www.gnu.org/licenses/lgpl.html).

Homepage: http://www.nervatura.com

## Installation & Quick Start

    $ npm install nervatura --production --save

or

    $ git clone https://github.com/nervatura/nervatura.git
    $ cd nervatura
    $ npm install --production --save

Start the server

* development mode:
```
  $ NODE_ENV=development node server.js
```
and [http://localhost:3000/](http://localhost:3000/)

* production mode:
```
  $ npm start
```
and [https://localhost:3000/](https://localhost:3000/)

* or change the file [.npmrc](.npmrc): production = true/false, and
```
  $ npm start
```    

Please see more the [Admin Guide](https://rawgit.com/nervatura/nervatura/master/views/docs/nas.html):
  * Optional packages
  * Cloud Hosting Services
  * Server config
  * Other recipes

## Docs & Community

[Nervatura Wiki](https://github.com/nervatura/nervatura/wiki)

More info see http://www.nervatura.com.

## Previous (Python) version

The Nervatura Framework v1.* is based on the Python/WEB2PY technology.
Please see https://github.com/nervatura/nerva2py