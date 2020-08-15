package nomad

import (
  nomad "github.com/hashicorp/nomad/api"
)

func checkError(e error) {
  if e != nil {
    panic(e)
  }
}

func Echo(message string) string {
  return message
}

func SubmitJob(name string) *Job {
  return nomad.NewServiceJob(name)
}
