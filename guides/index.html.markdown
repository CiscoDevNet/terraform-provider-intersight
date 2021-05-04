---
layout: "intersight"
page_title: "Provider: Intersight"
description: |-
The Cisco Intersight provider is used to interact with the many resources supported by Intersight.
---

# Cisco Intersight Terraform Provider
The Cisco Intersight Terraform Provider is used to interact with the
many resources supported by Cisco Intersight. You must have a valid Intersight account in order to use the provider.

Use the navigation to the left to read about the available resources.

## Installing Terraform
* Follow the steps [here](https://www.terraform.io/downloads.html) for installing terraform.

## Using the Cisco Intersight Terraform Provider

### Prerequisites
The system must have:
* Terraform
* Cisco Intersight Terraform Provider
* An active Cisco Intersight Account.


A `apikey`, `secretkey` and Intersight `endpoint` URL are required to start using the provider.
The following code must be included in a `.tf` file in the working directory, to establish connection between
the provider and your Intersight account.
```hcl-terraform
provider "intersight" {
  apikey    = "<api key>"
  secretkey = "<secret key as a string or in a file>"
  endpoint = "https://intersight.com"
}
```

###	Writing a managed source / resource
The resource objects’ name are in the format `intersight_<model_name_in_snake_case>`. E.g., `intersight_sol_policy`
is the resource object for **SOL Policy**. Each resource is assigned a name which can be later used for
tracking and referencing.  
The following is an example of a resource for server profile and NTP policy attached to the server profile:
```hcl-terraform
resource "intersight_server_profile" "server1" {
  name = "demo_server_profile"
  organization {
    moid   = "changeMe"
    object_type = "organization.Organization"
  }
  policy_bucket {
    object_type = "ntp.Policy"
    moid = intersight_ntp_policy.ntp1.moid
  }
}

resource "intersight_ntp_policy" "ntp1" {
  name    = "demo_ntp"
  enabled = true
  ntp_servers = [
    "10.10.10.10",
    "10.10.10.11",
    "10.10.10.12",
    "10.10.10.13"
  ]
  organization {
    moid   = "changeMe"
    object_type = "organization.Organization"
  }
}
```
The first resource is for creation of `server profile`. It is named `server1`. This name will not be reflected anywhere
in the Cisco Intersight. It is for reference among the .tf files. The server profile is attached to the NTP policy created
later. This is done by referencing to the `ntp1` policy in `policy_bucket`. A resource can point or reference to
another resource in the format `<resource>.<resource_name>.<property_name>` .

### Referencing Intersight MOs using data source
The data source objects’ name are in the format `intersight_<model_name_in_snake_case>`. E.g., `intersight_sol_policy`
is the data source object for SOL Policy. Each data source is assigned a name which can be later used for tracking and
referencing.
The following is an example of accessing a pre-existing server profile and attaching a NTP policy to the server profile.
```hcl-terraform
resource "intersight_server_profile" "server1" {
  name = "demo_server_profile"
  organization {
    moid   = "changeMe"
    object_type = "organization.Organization"
  }
  policy_bucket {
    object_type = "ntp.Policy"
    moid = data.intersight_ntp_policy.ntp1.results.0.moid
  }
}

data "intersight_ntp_policy" "ntp1" {
  name    = "demo_ntp"
}
```
An NTP policy with the given constraints i.e. name as `demo_ntp` must exist in the Intersight account. Data source
is used to access this object. The data source is named `ntp1`. The NTP policy is attached to the server profile with the name
`server1`. Data sources are referenced in the format `data.<data_source>.<data_source_name>.results.<index>.<property_name>`.

###	Viewing the Logs
`TF_LOG` is a terraform-specific environment variable that is used for viewing different category of logs.
By default, it is left empty. To view logs for terraform operations, this variable must be set to *DEBUG*.
You can view the request sent by the provider and the response received by the endpoint for debugging purposes.
#### Mac and Linux Bash
`export TF_LOG=debug`
#### Windows PowerShell
`$env:TF_LOG=“DEBUG”` 
