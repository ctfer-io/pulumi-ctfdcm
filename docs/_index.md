---
title: CTFd Provider-CM
meta_desc: Provides an overview on how to setup the CTFd-CM Provider for Pulumi.
layout: package
---

The CTFd-CM Resource Provider lets you manage [CTFd plugin for Chall-Manager](https://github.com/ctfer-io/ctfd-chall-manager) resources.

## Use Cases

Users of the CTFd-CM provider can:

- Manage `dynamic_iac` challenges

## Provider

To use the CTFd-CM Provider you will have to configure it and use it to create resources: it can't infer the CTFd url and credentials.
It should be used in conjunction of the [CTFd Provider](https://github.com/ctfer-io/pulumi-ctfd).

{{< chooser language "javascript,typescript,go,python,csharp" >}}

{{% choosable language javascript %}}

```javascript
import * as ctfd from '@ctfer-io/pulumi-ctfd';
import * as ctfdcm from '@ctfer-io/pulumi-ctfdcm';

// Stack configuration, other resources, etc.
// ...

// Create provider
let pv1 = new ctfd.Provider('ctfd', {
    url: 'https://my-ctf.lan',
    apiKey: 'ctfd_xxx', // please do not hardcode your credentials/api keys
});
let pv2 = new ctfdcm.Provider('ctfdcm', {
    url: 'https://my-ctf.lan',
    apiKey: 'ctfd_xxx', // please do not hardcode your credentials/api keys
});

// Create scenario file
let f = new ctfd.File('scenario', { /***/ }, { provider: pv1 });

// Create resources with the custom provider
let ch = new ctfdcm.ChallengeDynamicIaC('some-challenge', {
    name: 'My Challenge',
    category: 'misc',
    description: '...',
    value: 500,
    scenarioId: f.id,
}, { provider: pv2 });

// Other resources, export, etc.
// ...
```

{{% /choosable %}} {{% choosable language typescript %}}

```typescript
import * as ctfd from '@ctfer-io/pulumi-ctfd';
import * as ctfdcm from '@ctfer-io/pulumi-ctfdcm';

// Stack configuration, other resources, etc.
// ...

// Create provider
let pv1 = new ctfd.Provider('ctfd', {
    url: 'https://my-ctf.lan',
    apiKey: 'ctfd_xxx', // please do not hardcode your credentials/api keys
});
let pv2 = new ctfdcm.Provider('ctfdcm', {
    url: 'https://my-ctf.lan',
    apiKey: 'ctfd_xxx', // please do not hardcode your credentials/api keys
});

// Create scenario file
let f = new ctfd.File('scenario', { /***/ }, { provider: pv1 });

// Create resources with the custom provider
let ch = new ctfdcm.ChallengeDynamicIaC('some-challenge', {
    name: 'My Challenge',
    category: 'misc',
    description: '...',
    value: 500,
    scenarioId: f.id,
}, { provider: pv2 });

// Other resources, export, etc.
// ...
```

{{% /choosable %}} {{% choosable language go %}}

```go
package main

import (
    "github.com/ctfer-io/pulumi-ctfd/sdk/v2/go/ctfd"
    "github.com/ctfer-io/pulumi-ctfdcm/sdk/go/ctfdcm"
    "github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
    pulumi.Run(func(ctx *pulumi.Context) error {
        // Stack configuration, other resources, etc.
        // ...

        // Create provider
        pv1, err := ctfd.NewProvider(ctx, "ctfd", &ctfd.ProviderArgs{
            Url:    pulumi.String("https://my-ctf.lan"),
            ApiKey: pulumi.String("ctfd_xxx"), // please do not hardcode your credentials/api keys
        })
        if err != nil {
            return err
        }
        pv2, err := ctfdcm.NewProvider(ctx, "ctfdcm", &ctfdcm.ProviderArgs{
            Url:    pulumi.String("https://my-ctf.lan"),
            ApiKey: pulumi.String("ctfd_xxx"), // please do not hardcode your credentials/api keys
        })
        if err != nil {
            return err
        }

        // Build resource options to use it
        opts := []pulumi.ResourceOption{
            pulumi.Provider(pv1),
            pulumi.Provider(pv2),
        }

        // Create scenario file
        f, err := ctfd.NewFile(ctx, "scenario", &ctfd.File{ /***/ }, opts...)
        if err != nil {
            return err
        }

        // Create resources with the custom provider
        _, err = ctfdcm.NewChallengeDynamicIaC(ctx, "some-challenge", &ctfdcm.ChallengeDynamicIaCArgs{
            Name:        pulumi.String("My Challenge"),
            Category:    pulumi.String("misc"),
            Description: pulumi.String("..."),
            Value:       pulumi.Int(500),
            ScenarioId:  f.ID(),
        }, opts...)
        if err != nil {
            return err
        }

        // Other resources, export, etc.
        // ...

        return nil
    })
}
```

{{% /choosable %}} {{% choosable language python %}}

```python
import ctfer-io_pulumi_ctfd as ctfd
import ctfer-io_pulumi_ctfdcm as ctfdcm
import pulumi

# Stack configuration, other resources, etc.
# ...

# Create provider
pv1 = ctfd.Provider("ctfd", url="https://my-ctf.lan", api_key="ctfd_xxx") # please do not hardcode your credentials/api keys
pv2 = ctfdcm.Provider("ctfdcm", url="https://my-ctf.lan", api_key="ctfd_xxx") # please do not hardcode your credentials/api keys

# Create scenario file
f = ctfd.File('scenario', opts=pulumi.ResourceOptions(provider=pv1))

# Create resources with the custom provider
ch = ctfdcm.ChallengeDynamicIaC("some-challenge", name="My Challenge", category="misc", description="...", value=500, scenarioId=f.id opts=pulumi.ResourceOptions(provider=pv2))

# Other resources, export, etc.
# ...
```

{{% /choosable %}} {{% choosable language csharp %}}

```csharp
using System.Collections.Generic;
using System.Threading.Tasks;
using Pulumi;
using CTFerio.Ctfd;
using CTFerio.Ctfdcm;

class Program
{
    static Task Main() =>
        Deployment.Run(() => {
            // Stack configuration, other resources, etc.
            // ...

            // Create provider
            var pv1 = new Ctfd.Provider("ctfd", new Ctfd.ProviderArgs{
                Url = "https://my-ctf.lan",
                ApiKey = "ctfd_xxx" // please do not hardcode your credentials/api keys
            });
            var pv2 = new Ctfdcm.Provider("ctfdcm", new Ctfdcm.ProviderArgs{
                Url = "https://my-ctf.lan",
                ApiKey = "ctfd_xxx" // please do not hardcode your credentials/api keys
            });

            // Create scenario file
            var f = new Ctfd.File("scenario", new Ctfd.File{
                // ...
            }, new Pulumi.CustomResourceOptions { Provider = pv1 });

            // Create resources with the custom provider
            var ch = new Ctfdcm.ChallengeDynamicIaC("my-challenge", new Ctfdcm.ChallengeDynamicIaCArgs{
                Name = "My Challenge",
                Category = "misc",
                Description = "...",
                Value = 500,
                ScenarioId = f.ID,
            }, new Pulumi.CustomResourceOptions { Provider = pv });

            // Other resources, export, etc.
            // ...
        });
}
```

{{% /choosable %}}

{{< /chooser >}}
