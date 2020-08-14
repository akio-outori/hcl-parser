package main

import (
  "os"
  "fmt"
  "flag"
  "io/ioutil"
  "github.com/ghodss/yaml"
)

var input string
var output string

// It works every time
func getHelp() {
  fmt.Print("Usage:\n\n")
  flag.PrintDefaults()
  os.Exit(1)
}

func getYAML(input string) []byte {
  yml, err := ioutil.ReadFile(input)
  check_error(err)
  return yml
}

func convertYAML(input []byte) []byte {
  json, err := yaml.YAMLToJSON(input)
  check_error(err)
  return json
}

func writeJSON(input []byte, output string) {
  check_error(ioutil.WriteFile(output, input, 0644))
}

func check_error(e error) {
  if e != nil {
    panic(e)
  }
}

func check_variable(variable string) {
  if variable == "" {
    getHelp()
  }
}

func init() {
  // Set Variables
  flag.StringVar(&input, "input", "", "file to convert to json")
  flag.StringVar(&output, "output", "", "output file for the json rendered template")

  // Parse flags here so that we can check them
  flag.Parse()

  // Check Variables
  check_variable(input)
  check_variable(output)
}

func main() {
  yml  := getYAML(input)
  json := convertYAML(yml)
  writeJSON(json, output)
  fmt.Print(string(json))
}
