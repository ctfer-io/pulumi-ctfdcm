// Copyright 2024, Pulumi Corporation.  All rights reserved.
//go:build go || all
// +build go all

package examples

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/pulumi/pulumi/pkg/v3/testing/integration"
)

func Test_I_Go(t *testing.T) {
	opts := getGoBaseOptions(t).With(integration.ProgramTestOptions{
		Dir: filepath.Join(getCwd(t), "basic-go"),
		// The ctfd.File resource is created on the fly thus cannot be guaranteed
		// to keep the same hash... It may be improved in the future to remove this
		// technical debt.
		// TODO improve ctfd_file stability in terraform-provider-ctfd
		AllowEmptyPreviewChanges: true,
		AllowEmptyUpdateChanges:  true,
		Env: []string{
			fmt.Sprintf("CTFD_URL=%s", os.Getenv("CTFD_URL")),
		},
	})
	integration.ProgramTest(t, &opts)
}
