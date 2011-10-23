package main

import (
"exec"
//"io/ioutil"
)

func Crawl(link string) {
 // contents,err :=
  exec.Command(
  "wget", 
  //"-qO-", //write the webpages to the stdout
  "--no-parent", 
  "--wait=2", 
  "--timeout=100", 
  "--recursive", 
  "--accept=html,htm,php,net,aspx,asp,",
  link).Output()
}

//testing wget in go
func main() {
  seed := "http://www.cnn.com"
  Crawl(seed);
}


