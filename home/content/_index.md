---
title: About
type: docs
bookHidden: false
bookSideMenu: true
bookToC: false
---
<div class="row" style="paddin-top:32px;padding-bottom:16px;">
  <a href="{{< param VideoS1 >}}" target="_blank" rel="noopener noreferrer">
    <img alt="Move on to Open-data by Nervatura!" src="/images/video_s1.jpg"
      style="display:block;margin-left:auto;margin-right:auto;" />
  </a>
</div>

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
- built-in database drivers for postgres, mysql, mssql, sqlite databases
- a basic report generation library for creating simple PDF documents (eg. order, invoice, etc.) 
or CSV data files
- sample report templates and [**REPORT EDITOR**](/docs/client/program/editor) GUI
- [**CLIENT**](/docs/client) Web Component application and a basic ADMIN interface

The Nervatura [**Service**](/docs/service) is a simple interface layer that provides multiple, well-documented data access protocols for handling data. With their help, we can use the best data access for every development language and environment. Using the functions of the interfaces, we can be sure that the data is always read or written from the databases correctly and simply. The data can be stored in several types of databases, but they can be handled in the same format, and the database types can be easily changed.

The Nervatura Service has a modular structure, where most modules are optional. The default service includes all modules, but you [can build a customized service from them](/docs/install/#other-platforms-and-custom-build).

The Nervatura [**Client**](/docs/client) is a standard HTML5/ES6 [Web Component](https://developer.mozilla.org/en-US/docs/Web/Web_Components) application that contains no other external dependencies apart from the [lit helper functions](https://lit.dev). The standard HTML5 Web Components can be easily integrated or called from other javascript frameworks. It was created so that all the business data of the framework can be managed immediately after installation through a graphical interface. The client and report interface supports [multilingualism](/docs/start/customization#customize-the-appearance).

The Nervatura Framework **can be used independently**, but it is basically designed to provide a stable and secure foundation for self-developed, customized enterprise business systems. The framework **can be easily extended** with additional user interfaces or data management functions in **any programming language** or technology. Using the data from the framework, you can easily create your own web stores, user input interfaces or data interfaces for other systems. 

Nervatura Client supports the business processes that most companies may need. During your own developments, you can only focus on those that really require unique solutions, and you can use the technology that best suits the purpose. This type of development means greater **technological independence and security**, since your self-developed applications are only connected to other systems through well-documented interfaces, so unnecessary external technological dependencies cannot develop.

You can find more information about the use of different programming languages and development environments in the [**Examples**](/docs/start/examples) section.

