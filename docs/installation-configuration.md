---
title: CTFd Installation & Configuration
meta_desc: Provides an overview on how to setup the CTFd Provider for Pulumi.
layout: package
---

## Installation

The CTFd provider is available as a package in all Pulumi languages:

- JavaScript/TypeScript: [`@ctfer-io/pulumi-ctfdcm`](https://www.npmjs.com/package/@ctfer-io/pulumi-ctfdcm)
- Python: [`ctfer-io_pulumi-ctfdcm`](https://pypi.org/project/ctfer-io_pulumi-ctfdcm/)
- Golang: [`github.com/ctfer-io/pulumi-ctfd/sdk/v2/go/ctfdcm`](https://github.com/ctfer-io/pulumi-ctfdcm)
- .NET: [`CTFerio.Ctfdcm`](https://www.nuget.org/packages/CTFerio.Ctfdcm)

## Provider Binary

The CTFd provider binary is a third party binary. It can be installed using the pulumi plugin command.

```bash
pulumi plugin install resource ctfdcm <version> --server github://api.github.com/ctfer-io
```

Replace the version string with your desired version.

## Configuration

Pulumi relies on the CTFd REST JSON API to authenticate requests from the host machine to CTFd. Your credentials are never sent to pulumi.com.
The Pulumi CTFdCM Provider needs to be configured with CTFd credentials before it can be used to manage resources.

Your can either configure it using:

1. [environment variables](#environment-variables)
2. [provider configuration](#provider-configuration)

### Environment variables

Using the following you can configure the CTFd. For description of each variable, refer to the [provider configuration](#provider-configuration).

```bash
export CTFD_URL="https://my-ctf.lan"
export CTFD_API_KEY="ctfd_xxx"
# additional configuration
```

You can either use:
- authentication by API key (officially supported, most cases must implement a key rotation strategy)
- `username`/`password` combo (not officially supported but might be convenient in many use cases)

### Provider Configuration

You can also configure the provider using the stack configuration.
For instance, you can set the url using `pulumi config set url https://my-ctf.lan` then use it with, for instance, the following Go code.

```go
package main

import (
    "github.com/ctfer-io/pulumi-ctfdcm/sdk/v2/go/ctfdcm"
    "github.com/pulumi/pulumi/sdk/v3/go/pulumi"
    "github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

func main() {
    pulumi.Run(func(ctx *pulumi.Context) error {
        conf := config.New(ctx, "prod")

        // Create provider
        _, err := ctfdcm.NewProvider(ctx, "ctfdcm-fine-grained", &ctfdcm.ProviderArgs{
            Url:    pulumi.String(conf.Get("url")),
            ApiKey: conf.GetSecret("api_key"),
        })
        if err != nil {
            return err
        }

        // ...

        return nil
    })
}
```
