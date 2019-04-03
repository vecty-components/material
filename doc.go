/*
Quickstart Guide

1. In your project include the all-in-one distribution of the MDC javascript
library and set it to the global variable "mdc". This can be done a number of
ways (HTML script element, webpack, filename "mdc.inc.js" for gopherjs to pick
up, etc).

2. Import a Material component from this project in your Go progrem.

  import "agamigo.io/material/checkbox"

3. Make the HTML suitable for that MDC component available to your GopherJS
program. See: https://material.io/components/web/catalog/

  <html>
    <body>
      <div class="mdc-checkbox">
        <input class="mdc-checkbox__native-control" type="checkbox">
      </div>
    </body>
  </html>

4. Put that HTMLElement into a GopherJS object.

  cbElem := js.Global().Get("document").Get("body").Get("firstElementChild")

5. Create a new instance of the component and start it.

  cb := checkbox.CB{}
  cb.Start(cbElem)
*/
package material // import "agamigo.io/material"
