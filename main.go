package main

import (
  restservice "myapi/conf"
)

func main() {
  restservice.RestfulAPIServiceInit("HTTP")
  //restservice.RestfulAPIServiceInit("HTTPS")
}
