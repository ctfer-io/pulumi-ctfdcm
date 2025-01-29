// Copyright 2016-2024, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package ctfdcm

import (
	"path"

	// Allow embedding bridge-metadata.json in the provider.
	_ "embed"

	ctfdcm "github.com/ctfer-io/terraform-provider-ctfdcm/provider" // Import the upstream provider

	pfbridge "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/pf/tfbridge"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge/info"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge/tokens"

	"github.com/ctfer-io/pulumi-ctfdcm/provider/pkg/version"
)

// all of the token components used below.
const (
	// This variable controls the default name of the package in the package
	// registries for nodejs and python:
	mainPkg = "ctfdcm"
	// modules:
	mainMod = "index" // the ctfdcm module
)

//go:embed cmd/pulumi-resource-ctfdcm/bridge-metadata.json
var metadata []byte

// Provider returns additional overlaid schema and metadata associated with the provider.
func Provider() tfbridge.ProviderInfo {
	// Create a Pulumi provider mapping
	prov := tfbridge.ProviderInfo{
		//nolint:lll
		P: pfbridge.ShimProvider(ctfdcm.New(version.Version)()),

		Name:              "ctfdcm",
		Version:           version.Version,
		DisplayName:       "CTFd-Chall-Manager",
		Publisher:         "CTFer.io",
		LogoURL:           "https://raw.githubusercontent.com/ctfer-io/pulumi-ctfdcm/main/res/logo.png",
		PluginDownloadURL: "https://github.com/ctfer-io/pulumi-ctfdcm/releases/download/v${VERSION}/",
		Description:       "The CTFd-Chall-Manager provider for Pulumi, to manage its resources as code.",
		Keywords:          []string{"pulumi", "ctfd", "chall-manager", "category/utility"},
		License:           "Apache-2.0",
		Homepage:          "https://ctfer.io",
		Repository:        "https://github.com/ctfer-io/pulumi-ctfdcm",
		GitHubOrg:         "ctfer-io",
		MetadataInfo:      tfbridge.NewProviderMetadata(metadata),
		Config:            map[string]*tfbridge.SchemaInfo{},
		Resources: map[string]*info.Resource{
			"ctfdcm_challenge_dynamiciac": {
				Tok: tfbridge.MakeResource(mainPkg, mainMod, "ChallengeDynamicIaC"), // manual mapping for cases
			},
		},
		JavaScript: &tfbridge.JavaScriptInfo{
			// RespectSchemaVersion ensures the SDK is generated linking to the correct version of the provider.
			RespectSchemaVersion: true,
		},
		Python: &tfbridge.PythonInfo{
			// RespectSchemaVersion ensures the SDK is generated linking to the correct version of the provider.
			RespectSchemaVersion: true,
			// Enable modern PyProject support in the generated Python SDK.
			PyProject: struct{ Enabled bool }{true},
		},
		Golang: &tfbridge.GolangInfo{
			// Set where the SDK is going to be published to.
			ImportBasePath: path.Join(
				"github.com/ctfer-io/pulumi-ctfdcm/sdk/",
				tfbridge.GetModuleMajorVersion(version.Version),
				"go",
				mainPkg,
			),
			// Opt in to all available code generation features.
			GenerateResourceContainerTypes: true,
			GenerateExtraInputTypes:        true,
			// RespectSchemaVersion ensures the SDK is generated linking to the correct version of the provider.
			RespectSchemaVersion: true,
		},
		CSharp: &tfbridge.CSharpInfo{
			// RespectSchemaVersion ensures the SDK is generated linking to the correct version of the provider.
			RespectSchemaVersion: true,
			// Use a wildcard import so NuGet will prefer the latest possible version.
			PackageReferences: map[string]string{
				"Pulumi": "3.*",
			},
		},
	}

	// MustComputeTokens maps all resources and datasources from the upstream provider into Pulumi.
	//
	// tokens.SingleModule puts every upstream item into your provider's main module.
	//
	// You shouldn't need to override anything, but if you do, use the [tfbridge.ProviderInfo.Resources]
	// and [tfbridge.ProviderInfo.DataSources].
	prov.MustComputeTokens(tokens.SingleModule("ctfdcm_", mainMod,
		tokens.MakeStandard(mainPkg)))

	prov.MustApplyAutoAliases()
	prov.SetAutonaming(255, "-")

	return prov
}
