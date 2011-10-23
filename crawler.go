package main

import (
"exec"
"fmt"
"io/ioutil"
)


//testing wget in go
func main() {
  website := "http://www.google.com"
  val := exec.Command("wget", website)
  val.Output()
  infile := "index.html"
  contents,_ := ioutil.ReadFile(infile)
  fmt.Println(string(contents))
}


