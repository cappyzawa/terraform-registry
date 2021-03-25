# terraform-registry

[![BuildStatus](https://github.com/cappyzawa/terraform-registry/workflows/CI/badge.svg)](https://github.com/cappyzawa/terraform-registry/actions?query=workflow%3ACI)
[![Go Report Card](https://goreportcard.com/badge/github.com/cappyzawa/terraform-registry)](https://goreportcard.com/report/github.com/cappyzawa/terraform-registry)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/cappyzawa/terraform-registry)](https://pkg.go.dev/github.com/cappyzawa/terraform-registry)
[![codecov](https://codecov.io/gh/cappyzawa/terraform-registry/branch/main/graph/badge.svg)](https://codecov.io/gh/cappyzawa/terraform-registry)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://github.com/cappyzawa/terraform-registry/blob/main/LICENSE)

WIP: The implementation of [Provider Registry Protocol \- Terraform by HashiCorp](https://www.terraform.io/docs/internals/provider-registry-protocol.html) for in-houce provider.

## How to use

### Configuration file

This registry runs based on configuration file like as bellow.

```yaml
providers:
- namespace: cappyzawa
  type: sample 
  download_url_fmt: https://github.com/{namespace}/terraform-provider-{type}/releases/download/v{version}/terraform-provider-{type}_{version}_{os}_{arch}.zip
  shasums_url_fmt: https://github.com/{namespace}/terraform-provider-{type}/releases/download/v{version}/terraform-provider-{type}_{version}_SHA256SUMS
  shasums_signature_url_fmt: https://github.com/{namespace}/terraform-provider-{type}/releases/download/v{version}/terraform-provider-{type}_{version}_SHA256SUMS.sig
  signing_keys:
    gpg_public_keys:
    - key_id: XXXXXXXXXXXXXXXXXXXXXXXXX
      ascii_armor: |
        -----BEGIN PGP PUBLIC KEY BLOCK-----

        YYYYYY
        -----END PGP PUBLIC KEY BLOCK-----
  versions:
  - name: "0.1.0"
    assets:
    - os: darwin
      arch: amd64
      shasum: bbbbbbbbbbbbbbbbbbbbbbb
    - os: linux
      arch: amd64
      shasum: aaaaaaaaaaaaaaaaaaaaaaa
  - name: "0.0.5"
    assets:
    - os: darwin
      arch: amd64
      shasum: ccccccccccccccccccccccc
    - os: linux
      arch: amd64
      shasum: ddddddddddddddddddddddd
```

#### Provider

For details on how to publish the provider, please refer to [Terraform Registry \- Publishing Providers \- Terraform by HashiCorp](https://www.terraform.io/docs/registry/providers/publishing.html).

* `namespace`: The namespace of terraform provider.
* `type`: The type of terraform provider.
* `download_url_fmt`: Format of download url for terraform provider asset (`https` only). Available variables are `{namespace}`, `{type}`, `{version}`, `{os}`, `{arch}`.
* `shasums_url_fmt`: Format of shasums url for terraform provider assets (`https` only). Available variables are `{namespace}`, `{type}`, `{version}`, `{os}`, `{arch}`.
* `shasums_signature_url_fmt`: Format of shasums signature (`https` only). Available variables are `{namespace}`, `{type}`, `{version}`, `{os}`, `{arch}`.
* `signing_keys`: This keys vaidates shasums signature.

#### Module
TBD.

### Run

After the implementation of the configuration file is complete, start the terraform registry with the following command.

**`terraform` allows only `https` as a registry schema.**

```bash
$ terraform-registry -h
building terraform-registry in local

Usage:
  terraform-registry [command]

Available Commands:
  help        Help about any command
  server      run terraform registry locally
  version     display version

Flags:
  -h, --help   help for terraform-registry

Use "terraform-registry [command] --help" for more information about a command.
```

#### Using binary

```bash
terraform-registry server -c config.yaml

## If you want to register PID
terraform-registry server -c config.yaml --pid-file /var/run/regsitry.pid

## If you want to output log to file
terraform-registry server -c config.yaml --log-file ./log.log
```

#### Using docker (or Kubernetes)

```bash
docker run -itd --name terraform-registry -p 8080:8080 -v /tmp/config.yaml:/tmp/config.yaml ghcr.io/cappyzawa/terraform-registry -c /tmp/config.yaml
```
