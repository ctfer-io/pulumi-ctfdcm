---
title: CTFdCM Provider
meta_desc: Overview of the setup for the CTFd-Chall-Manager Pulumi provider (aka CTFdCM).
layout: package
---

[CTFd](https://ctfd.io) is an open-source CTFd platform famous for its simplicity and extensibility.
Using the [Chall-Manager plugin](https://github.com/ctfer-io/ctfd-chall-manager) you can deploy any challenge using Infra as Code (Kubernetes, GCP, AWS, ...) using the [Pulumi Auto API](https://www.pulumi.com/automation/).

It should be used in completion of the [CTFd Provider](https://www.pulumi.com/registry/packages/ctfd/).

## Use Cases

- Deploy and manage `dynamic_iac` challenges on CTFd when using the Chall-Manager plugin.

## Provider

To use the CTFdCM Provider you will have to configure it and use it to create resources: it can't infer the CTFd url and credentials.

{{< chooser language "javascript,typescript,go,python,csharp" >}}

{{% choosable language javascript %}}

```javascript
import * as ctfd from '@ctfer-io/pulumi-ctfd';
import * as ctfdcm from '@ctfer-io/pulumi-ctfdcm';

// Stack configuration, other resources, etc.
// ...

// Create providers
const pvArgs = {
    url: 'https://my-ctfd.lan',
    apiKey: 'ctfd_xxx', // please do not hardcode your credentials/api keys
};
let pvCtfd = new ctfd.Provider('ctfd-fine-grained', pvArgs);
let pvCtfdCM = new ctfdcm.Provider('ctfdcm-fine-grained', pvArgs);

// Create the scenario for Chall-Manager
let scenario = new ctfd.File('scenario', {
    // ... see CTFd provider documentation
}, { provider: pvCtfd });

// Create resources with the custom provider
let ch = new ctfdcm.ChallengeDynamicIaC('some-challenge', {
    name: 'My Challenge',
    category: 'misc',
    description: '...',
    value: 500,
    scenarioId: scenario.id,
    timeout: 600,
}, { provider: pvCtfdcm });

// Other resources, export, etc.
// ...
```

{{% /choosable %}} {{% choosable language typescript %}}

```typescript
import * as ctfd from '@ctfer-io/pulumi-ctfd';
import * as ctfdcm from '@ctfer-io/pulumi-ctfdcm';

// Stack configuration, other resources, etc.
// ...

// Create providers
const pvArgs = {
    url: 'https://my-ctfd.lan',
    apiKey: 'ctfd_xxx', // please do not hardcode your credentials/api keys
};
let pvCtfd = new ctfd.Provider('ctfd-fine-grained', pvArgs);
let pvCtfdCM = new ctfdcm.Provider('ctfdcm-fine-grained', pvArgs);

// Create the scenario for Chall-Manager
let scenario = new ctfd.File('scenario', {
    // ... see CTFd provider documentation
}, { provider: pvCtfd });

// Create resources with the custom provider
let ch = new ctfdcm.ChallengeDynamicIaC('some-challenge', {
    name: 'My Challenge',
    category: 'misc',
    description: '...',
    value: 500,
    scenarioId: scenario.id,
    timeout: 600,
}, { provider: pvCtfdcm });

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
        pvArgs := &ctfd.ProviderArgs{
            Url:    pulumi.String("https://my-ctf.lan"),
            ApiKey: pulumi.String("ctfd_xxx"), // please do not hardcode your credentials/api keys
        }
        pvCtfd, err := ctfd.NewProvider(ctx, "ctfd-fine-grained", pvArgs)
        if err != nil {
            return err
        }
        pvCtfdcm, err := ctfdcm.NewProvider(ctx, "ctfdcm-fine-grained", (*ctfdcm.ProviderArgs)(pvArgs))
        if err != nil {
            return err
        }

        // Build resource options to use it
        opts := []pulumi.ResourceOption{
            pulumi.Provider(pvCtfd),
            pulumi.Provider(pvCtfdcm),
        }

        // Create the scenario for Chall-Manager
        scenario, err := ctfd.NewFile(ctx, "scenario", &ctfd.FileArgs{
            // ... see CTFd provider documentation
        })
        if err != nil {
            return err
        }

        // Create resources with the custom provider
        _, err = ctfdcm.NewChallengeDynamicIaC(ctx, "some-challenge", &ctfdcm.ChallengeDynamicIaCArgs{
            Name:        pulumi.String("My Challenge"),
            Category:    pulumi.String("misc"),
            Description: pulumi.String("..."),
            Value:       pulumi.Int(500),
            ScenarioId:  scenario.ID(),
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

# Create providers
pvCtfd = ctfd.Provider("ctfd-fine-grained", url="https://my-ctf.lan", api_key="ctfd_xxx") # please do not hardcode your credentials/api keys
pvCtfdcm = ctfdcm.Provider("ctfdcm-fine-grained", url="https://my-ctf.lan", api_key="ctfd_xxx") # please do not hardcode your credentials/api keys

# Create the scenario for Chall-Manager
file = ctfd.File("scenario", opts=pulumi.ResourceOptions(provider=pvCtfd)) # ... see CTFd provider documentation

# Create resources with the custom provider
ch = ctfdcm.ChallengeDynamicIaC("some-challenge", name="My Challenge", category="misc", description="...", value=500, scenarioId=file.id, opts=pulumi.ResourceOptions(provider=pvCtfdcm))

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

            // Create providers
            var pvCtfd = new Ctfd.Provider("ctfd-fine-grained", new Ctfd.ProviderArgs{
                Url = "https://my-ctf.lan",
                ApiKey = "ctfd_xxx" // please do not hardcode your credentials/api keys
            });
            var pvCtfdcm = new Ctfdcm.Provider("ctfdcm-fine-grained", new Ctfdcm.ProviderArgs{
                Url = "https://my-ctf.lan",
                ApiKey = "ctfd_xxx" // please do not hardcode your credentials/api keys
            });

            // Create the scenario for Chall-Manager
            var file = new Ctfd.File("scenario", new Ctfd.FileArgs{
                // ... see CTFd provider documentation
            });

            // Create resources with the custom provider
            var ch = new Ctfdcm.ChallengeDynamicIaC("my-challenge", new Ctfdcm.ChallengeDynamicIaCArgs{
                Name = "My Challenge",
                Category = "misc",
                Description = "...",
                Value = 500,
                ScenarioId = scenario.Id
            }, new Pulumi.CustomResourceOptions { Provider = pvCtfdcm });

            // Other resources, export, etc.
            // ...
        });
}
```

{{% /choosable %}}

{{< /chooser >}}
