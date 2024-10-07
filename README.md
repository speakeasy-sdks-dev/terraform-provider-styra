# Styra DAS Terraform Provider

Styra Declarative Authorization Service (DAS) - The Enterprise OPA Manager providing unified policy management across the cloud native stack.

This Styra DAS Terraform provider enables managing Styra DAS resources with Terraform.

<!-- No SDK Installation -->
<!-- No SDK Example Usage -->
<!-- No SDK Available Operations -->
<!-- Start Summary [summary] -->
## Summary

Styra API: Styra DAS is entirely API-driven.

Access to the APIs requires authentication that should be provided as an Authorization HTTP header including a Styra DAS-issued token:

`Authorization: Bearer <YOURTOKENHERE>`

To request a token you need to have an Styra account, and create a token via the API Tokens menu.

For more information about the API: [Styra DAS Documentation](https://docs.styra.com)
<!-- End Summary [summary] -->

<!-- Start Table of Contents [toc] -->
## Table of Contents

* [Installation](#installation)
* [Available Resources and Data Sources](#available-resources-and-data-sources)
* [Testing the provider locally](#testing-the-provider-locally)
<!-- End Table of Contents [toc] -->

<!-- Start Installation [installation] -->
## Installation

To install this provider, copy and paste this code into your Terraform configuration. Then, run `terraform init`.

```hcl
terraform {
  required_providers {
    styra = {
      source  = "StyraInc/styra"
      version = "0.3.1"
    }
  }
}

provider "styra" {
  # Configuration options
}
```
<!-- End Installation [installation] -->

<!-- Start Available Resources and Data Sources [operations] -->
## Available Resources and Data Sources

### Resources

* [styra_library](docs/resources/library.md)
* [styra_policy](docs/resources/policy.md)
* [styra_role_binding](docs/resources/role_binding.md)
* [styra_secret](docs/resources/secret.md)
* [styra_stack](docs/resources/stack.md)
* [styra_system](docs/resources/system.md)
### Data Sources

* [styra_library](docs/data-sources/library.md)
* [styra_policy](docs/data-sources/policy.md)
* [styra_role_binding](docs/data-sources/role_binding.md)
* [styra_secret](docs/data-sources/secret.md)
* [styra_stack](docs/data-sources/stack.md)
* [styra_system](docs/data-sources/system.md)
<!-- End Available Resources and Data Sources [operations] -->

<!-- Start Testing the provider locally [usage] -->
## Testing the provider locally

#### Local Provider

Should you want to validate a change locally, the `--debug` flag allows you to execute the provider against a terraform instance locally.

This also allows for debuggers (e.g. delve) to be attached to the provider.

```sh
go run main.go --debug
# Copy the TF_REATTACH_PROVIDERS env var
# In a new terminal
cd examples/your-example
TF_REATTACH_PROVIDERS=... terraform init
TF_REATTACH_PROVIDERS=... terraform apply
```

#### Compiled Provider

Terraform allows you to use local provider builds by setting a `dev_overrides` block in a configuration file called `.terraformrc`. This block overrides all other configured installation methods.

1. Execute `go build` to construct a binary called `terraform-provider-styra`
2. Ensure that the `.terraformrc` file is configured with a `dev_overrides` section such that your local copy of terraform can see the provider binary

Terraform searches for the `.terraformrc` file in your home directory and applies any configuration settings you set.

```
provider_installation {

  dev_overrides {
      "registry.terraform.io/StyraInc/styra" = "<PATH>"
  }

  # For all other providers, install them directly from their origin provider
  # registries as normal. If you omit this, Terraform will _only_ use
  # the dev_overrides block, and so no other providers will be available.
  direct {}
}
```
<!-- End Testing the provider locally [usage] -->

<!-- Placeholder for Future Speakeasy SDK Sections -->

## Community

For questions, discussions and announcements related to Styra products, services and open source projects, please join
the Styra community on [Slack](https://communityinviter.com/apps/styracommunity/signup)!
