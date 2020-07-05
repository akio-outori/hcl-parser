package main

import (
  "fmt"
  "log"
  "gopkg.in/yaml.v2"
)

type kv struct {
  String string `yaml:"string,omitempty"`
}

func create_doc(value []byte) string {
  return fmt.Sprintf("---\n\n%v\n\n", string(value))
}

func main() {
  output, err := yaml.Marshal(&kv{String: "hello world"})
  if err != nil { log.Fatalf("error: %v", err) }

  fmt.Printf(create_doc(output))
}
