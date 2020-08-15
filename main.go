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

func getYaml(input string) []byte {
  yml, err := ioutil.ReadFile(input)
  checkError(err)
  return yml
}

func convertYaml(input []byte) []byte {
  json, err := yaml.YAMLToJSON(input)
  checkError(err)
  return json
}

func writeJson(input []byte, output string) {
  checkError(ioutil.WriteFile(output, input, 0644))
}

func writeStdout(input string) {
  fmt.Print(input)
}

func checkError(e error) {
  if e != nil {
    panic(e)
  }
}

func checkVariable(variable string, required bool) int {
  if variable != "" {
    return 0
  } else {
    if required == true {
      getHelp()
    }
    return 1
  }
}

func init() {
  // Set Variables
  flag.StringVar(&input, "input", "", "file to convert to json")
  flag.StringVar(&output, "output", "", "output file for the json rendered template")

  // Parse flags here so that we can check them
  flag.Parse()

  // Check Variables
  checkVariable(input, true)
  checkVariable(output, false)
}

func main() {
  yml  := getYaml(input)
  json := convertYaml(yml)

  if checkVariable(output, false) == 0 {
    writeJson(json, output)
  } else {
    writeStdout(string(json))
  }
}
