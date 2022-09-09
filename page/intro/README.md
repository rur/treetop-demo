# \<template> Request Demo

Experiments using a template protocol for modern web UX.

### What is a template request?

A web request that results in content fragments <sup>[1]</sup> being fetched and applied directly to the loaded web page. It is triggered from a web link, form, button or other JS component embedded in the requesting page.

### Motivation

The goal is to find a new way for HTML web servers to support basic interactivity with an extension of conventional navigation. Modern controls can be achieved with dynamically loaded markup bound to components. The hope is that this will reduce the need for custom APIs and enable more light-weight implementations for common dynamic features.

The protocol used here
An Accept <sup>[2]</sup> header indicates to the web server that the client is expecting a template as a response. If implemented, a HTML template fragment will be sent in the response body. That template will be applied to the requesting document using a find and replace strategy.

```
> GET /content_a HTTP/1.1
  Accept: application/x.treetop-html-template+xml

< HTTP/1.1 200 OK
  Content-Type: application/x.treetop-html-template+xml
  X-Page-URL: /content_a
  Vary: Accept

  <template>
    <div id="content">...</div>
    <div id="nav">...</div>
  </template>
```

That might seem too simple to be useful, so the example apps were created to demonstrate how various controls can be implemented using the protocol.

### What is 'treetop'?

This protocol was developed along with the Treetop library , a Go package for building web pages from nested templates.

#### The Network Inspector

This site is laid out to accommodate a network inspector <sup>[3]</sup> expanded to the right or left of the web browser window. The examples will make more sense if you are monitoring the network log.

### Markup and Styles

As you might have guessed, this page and all demo apps use the Bootstrap Component Library . Styling and markup is intended to be as familiar and conventional as possible <sup>[4]</sup> .

---

#### Footnotes

1. The Content Template element (MDN) was included in the HTML5 spec.
1. HTTP content negotiation (MDN) using a custom MIME type and custom response headers.
1. Network inspector docs: Chrome , Safari , Firefox , Edge
1. Suggestions for improvements to markup and CSS are very welcome, visit GitHub Treetop
