package main

import (
  "fmt"
  "flag"
  "io/ioutil"
  "github.com/ghodss/yaml"
)

func getYAML(input string) []byte {
  yml, err := ioutil.ReadFile(input)
  check(err)
  return yml
}

func convertYAML(input []byte) []byte {
  json, err := yaml.YAMLToJSON(input)
  check(err)
  return json
}

func writeJSON(input []byte, output string) {
  check(ioutil.WriteFile(output, input, 0644))
}

func check(e error) {
  if e != nil {
    panic(e)
  }
}

func main() {

  var input string
  var output string

  flag.StringVar(&input, "input", "", "file to convert to json")
  flag.StringVar(&output, "output", "", "output file for the json rendered template")
  flag.Parse()

  yml  := getYAML(input)
  json := convertYAML(yml)

  writeJSON(json, output)
  fmt.Print(string(json))

}
