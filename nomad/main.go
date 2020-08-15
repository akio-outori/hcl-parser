package nomad

import (
  "github.com/hashicorp/nomad/api"
  "github.com/hashicorp/terraform-provider-nomad/nomad/core/jobspec"
)

func Echo(message string) string {
  return message
}
