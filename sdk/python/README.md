# CTFd-CM Resource Provider

The CTFd-CM Resource Provider lets you manage [CTFd plugin for Chall-Manager](http://github.com/ctfer-io/ctfd-chall-manager) resources.

## Installing

This package is available for several languages/platforms:

### Node.js (JavaScript/TypeScript)

To use from JavaScript or TypeScript in Node.js, install using either `npm`:

```bash
npm install @ctfer-io/ctfdcm
```

or `yarn`:

```bash
yarn add @ctfer-io/ctfdcm
```

### Python

To use from Python, install using `pip`:

```bash
pip install ctferio-ctfdcm
```

### Go

To use from Go, use `go get` to grab the latest version of the library:

```bash
go get github.com/ctfer-io/pulumi-ctfdcm/sdk/go/...
```

### .NET

To use from .NET, install using `dotnet add package`:

```bash
dotnet add package CTFerio.Ctfdcm
```

## Configuration

The following configuration points are available for the `ctfdcm` provider:

- `ctfdcm:region` (environment: `CTFDCM_REGION`) - the region in which to deploy resources

## Reference

For detailed reference documentation, please visit [the Pulumi registry](https://www.pulumi.com/registry/packages/ctfdcm/api-docs/).
