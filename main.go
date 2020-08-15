package main

import (
  "os"
  "fmt"
  "flag"
  "github.com/akio-outori/hcl-parser/yaml"
  "github.com/akio-outori/hcl-parser/nomad"
)

var action string
var input string
var output string

// It works every time
func getHelp() {
  fmt.Print("Usage:\n\n")
  flag.PrintDefaults()
  os.Exit(1)
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
  flag.stringVar(&action, "", "", "action to take")
  flag.StringVar(&input, "input", "", "file to convert to json")
  flag.StringVar(&output, "output", "", "output file for the json rendered template")

  // Parse flags here so that we can check them
  flag.Parse()

  // Check Variables
  checkVariable(action, true)
  checkVariable(input, true)
  checkVariable(output, false)
}

func main() {

  switch {

    case action == "convert":
      json := yaml.Convert(input)
      yaml.Write(json, output)

    case action == "nomad":
      nomad.Echo("test")
      nomad.SubmitJob("test")

    default:
      getHelp()

  }
}
