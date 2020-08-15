package yaml

import (
  "fmt"
  "io/ioutil"
  "github.com/ghodss/yaml"
)

func checkError(e error) {
  if e != nil {
    panic(e)
  }
}

func getYaml(input string) []byte {
  yml, err := ioutil.ReadFile(input)
  checkError(err)
  return yml
}

func Convert(input string) []byte {
  yml       := getYaml(input)
  json, err := yaml.YAMLToJSON(yml)
  checkError(err)
  return json
}

func writeJson(input []byte, output string) {
  checkError(ioutil.WriteFile(output, input, 0644))
}

func writeStdout(input string) {
  fmt.Print(input)
}

func Write(input []byte, output string) {
  if output != "" {
    writeJson(input, output)
  } else {
    writeStdout(string(input))
  }
}
