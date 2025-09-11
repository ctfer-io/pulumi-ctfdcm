package ctfd

import (
	"fmt"
	"path"

	// Allow embedding bridge-metadata.json in the provider.
	_ "embed"

	"github.com/ctfer-io/pulumi-ctfdcm/provider/pkg/version"
	ctfdcm "github.com/ctfer-io/terraform-provider-ctfdcm/provider"
	pf "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/pf/tfbridge"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge/tokens"
)

const (
	mainPkg = "ctfdcm"
	mainMod = "index" // the ctfd module
)

//go:embed cmd/pulumi-resource-ctfdcm/bridge-metadata.json
var metadata []byte

func Provider() tfbridge.ProviderInfo {
	prov := tfbridge.ProviderInfo{
		P:                 pf.ShimProvider(ctfdcm.New(version.Version)()),
		Version:           version.Version,
		Name:              "ctfdcm",
		DisplayName:       "CTFd-CM",
		Publisher:         "CTFer.io",
		PluginDownloadURL: "github://api.github.com/ctfer-io/",
		Description:       "The CTFd-CM provider for Pulumi, to manage its resources as code.",
		Keywords:          []string{"pulumi", "ctfd", "chall-manager", "category/cloud"},
		LogoURL:           "https://raw.githubusercontent.com/ctfer-io/pulumi-ctfdcm/main/res/logo.png",
		License:           "Apache-2.0",
		Homepage:          "https://ctfer.io",
		Repository:        "https://github.com/ctfer-io/pulumi-ctfdcm",
		GitHubOrg:         "ctfer-io",
		MetadataInfo:      tfbridge.NewProviderMetadata(metadata),
		Config:            map[string]*tfbridge.SchemaInfo{},
		Resources: map[string]*tfbridge.ResourceInfo{
			"ctfdcm_challenge_dynamiciac": {
				Tok: tfbridge.MakeResource(mainPkg, mainMod, "ChallengeDynamicIaC"),
			},
			"ctfdcm_instance": {
				Tok: tfbridge.MakeResource(mainPkg, mainMod, "Instance"),
			},
		},
		DataSources: map[string]*tfbridge.DataSourceInfo{},
		JavaScript: &tfbridge.JavaScriptInfo{
			PackageName: "@ctfer-io/pulumi-ctfdcm",
			// List any npm dependencies and their versions
			Dependencies: map[string]string{
				"@pulumi/pulumi": "^3.0.0",
			},
			DevDependencies: map[string]string{
				"@types/node": "^10.0.0", // so we can access strongly typed node definitions.
				"@types/mime": "^2.0.0",
			},
		},
		Python: &tfbridge.PythonInfo{
			PackageName: fmt.Sprintf("ctfer-io_pulumi-%[1]s", mainPkg),
			// List any Python dependencies and their version ranges
			Requires: map[string]string{
				"pulumi": ">=3.0.0,<4.0.0",
			},
		},
		Golang: &tfbridge.GolangInfo{
			ImportBasePath: path.Join(
				fmt.Sprintf("github.com/ctfer-io/pulumi-%[1]s/sdk/", mainPkg),
				tfbridge.GetModuleMajorVersion(version.Version),
				"go",
				mainPkg,
			),
			GenerateResourceContainerTypes: true,
		},
		CSharp: &tfbridge.CSharpInfo{
			RootNamespace: "CTFerio",
			PackageReferences: map[string]string{
				"Pulumi": "3.*",
			},
		},
	}

	// These are new API's that you may opt to use to automatically compute resource
	// tokens, and apply auto aliasing for full backwards compatibility.  For more
	// information, please reference:
	// https://pkg.go.dev/github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge#ProviderInfo.ComputeTokens
	prov.MustComputeTokens(tokens.SingleModule("ctfdcm_", mainMod,
		tokens.MakeStandard(mainPkg)))
	prov.MustApplyAutoAliases()
	prov.SetAutonaming(255, "-")

	return prov
}
