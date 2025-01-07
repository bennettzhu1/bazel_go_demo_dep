//go:build tools
// +build tools

package tools

import (
    // Blank import to include dbmate as a development tool
    _ "github.com/amacneil/dbmate/v2"
    _ "github.com/terraform-providers/terraform-provider-aws"
)
