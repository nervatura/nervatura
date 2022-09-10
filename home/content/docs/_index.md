---
title: About
type: docs
bookHidden: false
bookSideMenu: true
bookToC: false
---

Nervatura is a business management framework based on **open-data principle**. It can handle any type of business related information, starting from customer details, up to shipping, stock or payment information. Developed as open-source project and can be used freely under the scope of [LGPLv3 License](http://www.gnu.org/licenses/lgpl.html).

<div class="row infoRow">
  <div class="cell mobile">
    The main aspects of its design were:
    <ul>
      <li>simple and transparent structure</li>
      <li>capability of storing different data types of an average company</li>
      <li>effective, easily expandable and secure data storage</li>
      <li>support of several database types</li>
      <li>well documented, easy data management</li>
    </ul>
  </div>
  <div class="cell contactCol mobile">
    <div class="paypal">
      <a href="http://nervatura.com" target="_blank" rel="noopener noreferrer"
          title="Nervatura Homepage" >
          <img src="/images/logo_green.svg" style="width:80px;" alt="logo" class="logo" />
      </a>
    </div>
    <div class="paypal">
      <form action="https://www.paypal.com/donate" method="post" target="_top">
        <input type="hidden" name="hosted_button_id" value="{{< param DonateId >}}" />
        <input type="image" src="https://www.paypalobjects.com/en_US/i/btn/btn_donateCC_LG.gif" border="0" name="submit" title="PayPal - The safer, easier way to pay online!" alt="Donate with PayPal button" />
        <img alt="" border="0" src="https://www.paypal.com/en_US/i/scr/pixel.gif" width="1" height="1" />
      </form>
    </div>
  </div>
</div>

The framework is based on Nervatura Object [**MODEL**](/docs/model) specification. It is a general **open-data model**, which can store all information generated in the operation of a usual corporation.

The Nervatura service is small and fast. A single ~5 MB file/image contains all the necessary dependencies.
The framework includes:
- [**CLI API**](/docs/service/cli#cli-api) (command line interface)
- [**CGO API**](/docs/service/cli#cgo-api) (C shared library)
- standard HTTP [**RESTful API**](/docs/service/api) for client communication
- HTTP/2-based [**gRPC API**](/docs/service/grpc) for server-side communication
- JWT generation, external token validation, SSL/TLS support and other HTTP security [settings](/docs/install#configuration-options)
- built-in database drivers for postgres, mysql, sqlite databases
- a basic report generation library for creating simple PDF documents (eg. order, invoice, etc.) 
or CSV data files
- sample report templates and [**REPORT EDITOR**](/docs/client/program/editor) GUI
- PWA [**CLIENT**](/docs/client) application and a basic ADMIN interface

The client and report interface supports [multilingualism](/docs/start/customization#customize-the-appearance). The framework can be easily extended with additional interfaces and functions in any languages. See the [**Examples**](/docs/start/examples) for more information

