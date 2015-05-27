package main

import (
  "fmt"
  "os"
  "path/filepath"
  "io/ioutil"
  "bytes"
  "encoding/json"
)

type Config struct {
  Title, Theme string
}

func main() {
  fmt.Println("Welcome to GoLang CMS!\n")

  slash := string(filepath.Separator)

  config := getConfig()

  theme := config.Theme

  // Compile website
  header, _ := ioutil.ReadFile("themes" + slash + theme + slash + "header.html");
  index, _ := ioutil.ReadFile("content" + slash + "index.html");
  footer, _ := ioutil.ReadFile("themes" + slash + theme + slash + "footer.html");
  page := [][]byte{header,index,footer}
  publicindex, _ := os.Create("public" + slash + "index.html")
  publicindex.Write(bytes.Join(page, []byte("")))

  css, _ := ioutil.ReadFile("themes" + slash + theme + slash + "css" + slash + "bootstrap.min.css")
  publiccss, _ := os.Create("public" + slash + "css" + slash + "bootstrap.min.css")
  publiccss.Write(css)
}

func getConfig() Config {

  configFile, _ := ioutil.ReadFile("config.json")
  configDecoded := json.NewDecoder(bytes.NewReader(configFile))

  var config Config
  configDecoded.Decode(&config)

  return config
}
