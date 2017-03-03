Another Charset Reader
======================

## Support charset:

- windows-874

- windows-1250

- windows-1251

- windows-1252

- windows-1253

- windows-1254

- windows-1255

- windows-1256

- windows-1257

- windows-1258

- koi8r

- koi8u

- iso-8859-1

- iso-8859-2

- iso-8859-3

- iso-8859-4

- iso-8859-5

- iso-8859-6

- iso-8859-7

- iso-8859-8

- iso-8859-10

- iso-8859-13

- iso-8859-14

- iso-8859-15

## Example for XML

``` go
package main

import (
  "bytes"
  "log"
  
  "github.com/l-vitaly/acharset"
)

func main() {
    var result map[string]interface{}
    
    buf := bytes.NewBufferString("xml content")
    
    dec := xml.NewDecoder(buf)
    dec.CharsetReader = acharset.CharsetReader
    err = dec.Decode(&result)
    if err != nil {
        panic(err)
    }
    
    log.Println(result)
}
```
