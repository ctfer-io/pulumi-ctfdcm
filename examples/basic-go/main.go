package main

import (
	"archive/zip"
	"bytes"
	"encoding/base64"
	"io"
	"os"
	"os/exec"

	"github.com/ctfer-io/pulumi-ctfd/sdk/v2/go/ctfd"
	"github.com/ctfer-io/pulumi-ctfdcm/sdk/go/ctfdcm"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"gopkg.in/yaml.v3"
)

var (
	url = os.Getenv("CTFD_URL")
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Build CTFd and CTFd-Chall-Manager providers
		pargs := &ctfd.ProviderArgs{
			Url:      pulumi.String(url),
			Username: pulumi.String("ctfer"), // expected to be configured as so
			Password: pulumi.String("ctfer"), // expected to be configured as so
		}

		p1, err := ctfd.NewProvider(ctx, "ctfd", pargs)
		if err != nil {
			return err
		}

		p2, err := ctfdcm.NewProvider(ctx, "ctfdcm", (*ctfdcm.ProviderArgs)(pargs))
		if err != nil {
			return err
		}

		opts := []pulumi.ResourceOption{
			pulumi.Provider(p1),
			pulumi.Provider(p2),
		}

		// Deploy a DynamicIaC challenge
		// => Create the scenario
		scn := scenario()
		f, err := ctfd.NewFile(ctx, "scenario", &ctfd.FileArgs{
			Name:       pulumi.String("scenario.zip"),
			Contentb64: pulumi.String(scn),
		}, opts...)
		if err != nil {
			return err
		}

		// => Create the challenge that depends on the previous file
		ch, err := ctfdcm.NewChallengeDynamicIaC(ctx, "dynamic", &ctfdcm.ChallengeDynamicIaCArgs{
			Name:     pulumi.String("My Challenge"),
			Category: pulumi.String("Some category"),
			Description: pulumi.String(`My superb description !

And it's multiline :o`),
			State:   pulumi.String("visible"),
			Value:   pulumi.Int(500),
			Decay:   pulumi.Int(17),
			Minimum: pulumi.Int(50),

			Shared:        pulumi.Bool(true),
			DestroyOnFlag: pulumi.Bool(true),
			ManaCost:      pulumi.Int(1),
			ScenarioId:    f.ID(),
			Timeout:       pulumi.Int(600),
		}, opts...)
		if err != nil {
			return err
		}

		ctx.Export("id", ch.ID())
		return nil
	})
}

func scenario() string {
	buf := bytes.NewBuffer([]byte{})
	archive := zip.NewWriter(buf)

	// Add Pulumi.yaml file
	mp := map[string]any{
		"name": "scenario",
		"runtime": map[string]any{
			"name": "go",
			"options": map[string]any{
				"binary": "./main",
			},
		},
		"description": "An example scenario.",
	}
	b, err := yaml.Marshal(mp)
	if err != nil {
		panic(err)
	}
	w, err := archive.Create("Pulumi.yaml")
	if err != nil {
		panic(err)
	}
	if _, err := io.Copy(w, bytes.NewBuffer(b)); err != nil {
		panic(err)
	}

	// Add binary file
	fs, err := compile()
	if err != nil {
		panic(err)
	}
	defer fs.Close()

	fst, err := fs.Stat()
	if err != nil {
		panic(err)
	}
	header, err := zip.FileInfoHeader(fst)
	if err != nil {
		panic(err)
	}
	header.Name = "main"

	// Create archive
	f, err := archive.CreateHeader(header)
	if err != nil {
		panic(err)
	}

	// Copy the file's contents into the archive.
	_, err = io.Copy(f, fs)
	if err != nil {
		panic(err)
	}

	// Complete zip creation
	if err := archive.Close(); err != nil {
		panic(err)
	}

	return base64.StdEncoding.EncodeToString(buf.Bytes())
}

func compile() (*os.File, error) {
	cmd := exec.Command("go", "build", "-o", "main", "scenario/main.go")
	cmd.Env = append(os.Environ(), "CGO_ENABLED=0")
	if err := cmd.Run(); err != nil {
		return nil, err
	}
	defer func() {
		cmd := exec.Command("rm", "main")
		_ = cmd.Run()
	}()
	return os.Open("main")
}
