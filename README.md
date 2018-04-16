Nervatura
=========

Open Source Business Management Framework

## Features

Nervatura is a business management framework based on **open-data principle**. It can handle any type of business related information, starting from customer details, up to shipping, stock or payment information.

>The open-data principle regulates the access to our business data. The point is the logic of the data storage. It means that the data are defined for storage so, that those are compliant with an open data-model which could be accessed and interpreted by anyone. It doesn't concern the physical storage of the data, that can be implemented according to one's needs. However it should ensure that data can be managed safely according to published description. Retrieving, new data creation, possibility to export the entire data structure should be provided.

>What are the main advantages of open-data applications?
* **safety:** provided one's have proper usage rights and physical access to a database, then will be able to interpret and process the data correctly without any help or permission from a third party. Information in the data becomes independent of any management system, its treatment is not tied to specific programs or technologies.
* **efficiency and cost reduction:** the business management system can be developed and diversified in accordance with one's needs. There is no need and pressure to be tied to a solution of any vendor, the most appropriate tools and programs can be selected for all tasks. The only criteria is that selected applications should be able to communicate and exchange data with each other or with a central database according to the open-data description. The elements of the system can at any time be flexibly developed or new ones added by choosing the best offers available on the market.

The framework is based on Nervatura [Object Model](https://github.com/nervatura/nervatura/wiki/Nervatura-Object-Model-%28NOM%29) specification. The main aspects of its design were:

* simple and transparent structure
* capability of storing different data types of an average company
* effective, easily expandable and secure data storage
* support of several database types
* well documented, easy data management

The Nervatura Framework is a set of open source applications and documentations. It enables to easily create a basic open-data business management system. Its key elements are:
  * [Nervatura Admin Tools](https://github.com/nervatura/nervatura/wiki/Nervatura-Admin-Tools-(NAS)) to handle Nervatura databases: creation, data export, access rights system
  * [Data Interface](https://github.com/nervatura/nervatura/wiki/Nervatura-Data-Interface-%28NDI%29) as a non-graphical, command-based user interface
  * [Programming Interface](https://github.com/nervatura/nervatura/wiki/Nervatura-Programming-Interface-%28NPI%29) for applications that would use the data
  * [Client- and server-side data reporting](https://github.com/nervatura/nervatura/wiki/Nervatura-Report-%28NDR%29)
  * and other documents, resources, sample code and demo programs

Developed as open-source project and can be used freely under the scope of [LGPLv3 License](http://www.gnu.org/licenses/lgpl.html).

Homepage: http://www.nervatura.com

## Nervatura packages

[Express Framework](https://github.com/nervatura/nervatura-express/), [Azure Functions](https://github.com/nervatura/nervatura-azure/)

Experimental: [Google Cloud Functions](https://github.com/nervatura/nervatura-gcloud/), [AWS Lambda](https://github.com/nervatura/nervatura-lambda/), [IBM Cloud Functions](https://github.com/nervatura/nervatura-openwhisk/)

## Quick Start

Nervatura Core Package:

    $ npm install nervatura --production --save

Nervatura with Express Framework:

    $ git clone https://github.com/nervatura/nervatura-express.git nervatura
    $ cd nervatura
    $ npm install --production

For more packages and deployment options, please see the [Admin Guide](https://rawgit.com/nervatura/nervatura/master/views/docs/nas.html):
  * Nervatura packages
  * Cloud Hosting Services
  * Server config
  * Other recipes

## More Resources and Sample Applications

Sample Applications: [Nervatura Booking](https://github.com/nervatura/nervatura-booking/) (Express+React), [Nervatura Report](https://github.com/nervatura/nervatura-react/) (React), [wxPython Client](https://rawgit.com/nervatura/nervatura/master/public/download/nwx.zip), [Javascript RPC functions](https://rawgit.com/nervatura/nervatura/master/public/download/npijson.zip), [LibreOffice Calc](https://rawgit.com/nervatura/nervatura/master/public/download/calc_demo.ods)

More Resources: [Nervatura Touch](http://nervatura.com#touch), Data Interface Wizard ([nervatura-express](https://github.com/nervatura/nervatura-express/)), Nervatura Report ([nervatura-express](https://github.com/nervatura/nervatura-express/))

## Docs & Community

[Nervatura Wiki](https://github.com/nervatura/nervatura/wiki)

More info see http://www.nervatura.com.