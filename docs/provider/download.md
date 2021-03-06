# API doc

This is API documentation for Provider versions. This is generated by `httpdoc`. Don't edit by hand.

## Table of contents

- [[200] GET /v1/providers/cappyzawa/concourse/0.1.0/download/darwin/amd64](#200-get-v1providerscappyzawaconcourse0.1.0downloaddarwinamd64)
- [[404] GET /v1/providers/foo/bar/0.1.0/download/darwin/amd64](#404-get-v1providersfoobar0.1.0downloaddarwinamd64)
- [[404] GET /v1/providers/cappyzawa/concourse/11.11.0/download/darwin/amd64](#404-get-v1providerscappyzawaconcourse11.11.0downloaddarwinamd64)
- [[404] GET /v1/providers/cappyzawa/concourse/0.1.0/download/windows/amd64](#404-get-v1providerscappyzawaconcourse0.1.0downloadwindowsamd64)


## [200] GET /v1/providers/cappyzawa/concourse/0.1.0/download/darwin/amd64

existing provider: cappyzawa/concourse:0.1.0(darwin/amd64)

### Request









### Response

Headers

| Name  | Value  | Description |
| ----- | :----- | :--------- |
| Content-Type | application/json |  |





Response example

<details>
<summary>Click to expand code.</summary>

```javascript
{"protocols":["5.3"],"os":"darwin","arch":"amd64","filename":"terraform-provider-concourse_0.1.0_darwin_amd64.zip","download_url":"https://github.com/cappyzawa/terraform-provider-concourse/releases/download/v0.1.0/terraform-provider-concourse_0.1.0_darwin_amd64.zip","shasums_url":"https://github.com/cappyzawa/terraform-provider-concourse/releases/download/v0.1.0/terraform-provider-concourse_0.1.0_SHA256SUMS","shasums_signature_url":"https://github.com/cappyzawa/terraform-provider-concourse/releases/download/v0.1.0/terraform-provider-concourse_0.1.0_SHA256SUMS.sig","shasum":"82abade6ec0c90b88f205c38e9e241ec229df6c7150ee59c242f1924070bad82","signing_keys":{"gpg_public_keys":[{"key_id":"XXXXXXXXXXXXXXXXXXXXXXXXX","ascii_armor":"-----BEGIN PGP PUBLIC KEY BLOCK-----\n\nYYYYYY\n-----END PGP PUBLIC KEY BLOCK-----\n","trust_signature":"","source":"","source_url":""}]}}

```

</details>


## [404] GET /v1/providers/foo/bar/0.1.0/download/darwin/amd64

non existing provider: foo/bar:0.1.0(darwin/amd64)

### Request









### Response

Headers

| Name  | Value  | Description |
| ----- | :----- | :--------- |
| Content-Type | application/json |  |






## [404] GET /v1/providers/cappyzawa/concourse/11.11.0/download/darwin/amd64

non existing provider version: cappyzawa/concourse:11.11.0(darwin/amd64)

### Request









### Response

Headers

| Name  | Value  | Description |
| ----- | :----- | :--------- |
| Content-Type | application/json |  |






## [404] GET /v1/providers/cappyzawa/concourse/0.1.0/download/windows/amd64

non existing provider os: cappyzawa/concourse:0.1.0(windows/amd64)

### Request









### Response

Headers

| Name  | Value  | Description |
| ----- | :----- | :--------- |
| Content-Type | application/json |  |







