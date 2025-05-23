---
title: CTFd-CM Installation & Configuration
meta_desc: Provides an overview on how to setup the CTFd-CM Provider for Pulumi.
layout: package
---

## Installation

The CTFd-CM provider is available as a package in all Pulumi languages:

- JavaScript/TypeScript: [`@ctfer-io/pulumi-ctfdcm`](https://www.npmjs.com/package/@ctfer-io/pulumi-ctfdcm)
- Python: [`ctfer-io_pulumi-ctfdcm`](https://pypi.org/project/ctfer-io_pulumi-ctfdcm/)
- Golang: [`github.com/ctfer-io/pulumi-ctfdcm/sdk/go/ctfdcm`](https://github.com/ctfer-io/pulumi-ctfdcm)
- .NET: [`CTFerio.Ctfdcm`](https://www.nuget.org/packages/CTFerio.Ctfdcm)

## Provider Binary

The CTFd-CM provider binary is a third party binary. It can be installed using the pulumi plugin command.

```bash
pulumi plugin install resource ctfdcm <version> --server github://api.github.com/ctfer-io
```

Replace the version string with your desired version.

## Configuration

Pulumi relies on the CTFd REST JSON API to authenticate requests from the host machine to CTFd. Your credentials are never sent to pulumi.com.
The Pulumi CTFd-CM Provider needs to be configured with CTFd credentials before it can be used to manage resources.

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

We recommend using an API key rather than a nonce/session combo for security purposes: the API key is natively supported thus enable better logging, authentication & authorization.
Moreover, it is possible to rotate the keys and revoke them on the fly using the API, while a session/nonce is not build in this way.

### Provider Configuration

You can also configure the provider using the stack configuration.
For instance, you can set the url using `pulumi config set url https://my-ctf.lan` then use it with, for instance, the following Go code.

```go
package main

import (
    "github.com/ctfer-io/pulumi-ctfdcm/sdk/go/ctfdcm"
    "github.com/pulumi/pulumi/sdk/v3/go/pulumi"
    "github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

func main() {
    pulumi.Run(func(ctx *pulumi.Context) error {
        conf := config.New(ctx, "prod")

        // Create provider
        _, err := ctfdcm.NewProvider(ctx, "ctfd-fine-grained", &ctfdcm.ProviderArgs{
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
