package main

import (
  "fmt"
  "io/ioutil"
  "github.com/ghodss/yaml"
)

func check(e error) {
  if e != nil {
    panic(e)
  }
}

func main() {

  yml, err := ioutil.ReadFile("jobspec.yaml")
  check(err)

  j2, err := yaml.YAMLToJSON(yml)
  check(err)

  write_err := ioutil.WriteFile("test.json", j2, 0644)
  check(write_err)

  fmt.Print(string(j2))

}
