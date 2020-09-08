# User Guide for Cisco Intersight Terraform Provider
The reader of the Cisco Intersight Terraform Provider guide must be familiar with the following:
* Terraform
* HashiCorp Configuration Language
* Cisco Intersight Platform

## Installing Terraform
* Download terraform zip from https://www.terraform.io/downloads.html as per your operating system specifications.
* Extract the zip and move it to a directory of your choice.
* Add the path of the directory to PATH variable of the system. 
* For a detailed video on installing terraform, visit 
[terraform website](https://learn.hashicorp.com/terraform/getting-started/install.html)

## Using the Cisco Intersight Terraform Provider

### Prerequisites
The system must have:
* Terraform
* Cisco Intersight Terraform Provider 
* An active Cisco Intersight Account.

To start using the provider, API Key, Secret Key and Intersight endpoint URL are required. 
The following code must be included in a .tf file in the working directory, to pair the provider with Intersight account.
```hcl-terraform
provider "intersight" {
  apikey    = "<api key>"
  secretkeyfile = "<secret key in a file>"
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
  name = "server1"
  tags {
    key   = "server"
    value = "demo"
  }
  assigned_server {
    moid        = "56789acbcd67890"
    object_type = "compute.RackUnit"
  }
}

resource "intersight_ntp_policy" "ntp1" {
  name    = "ntp2"
  enabled = true
  ntp_servers = [
    "10.10.10.10",
    "10.10.10.11",
    "10.10.10.12",
    "10.10.10.13"
  ]
  profiles {
    moid        = intersight_server_profile.server1.moid
    object_type = "server.Profile"
  }
}
```
The first resource is for creation of `server profile`. It is named `server1`. This name will not be reflected anywhere 
in the Cisco Intersight. It is for reference among the .tf files. The NTP policy is attached to the server profile created 
earlier. This is done by referencing to the `server1` profile in `profiles.moid`. A resource can point or reference to 
another resource in the format `<resource>.<resource_name>.<property_name>` . 

### Referencing Intersight MOs using data source
The data source objects’ name are in the format `intersight_<model_name_in_snake_case>`. E.g., `intersight_sol_policy` 
is the data source object for SOL Policy. Each data source is assigned a name which can be later used for tracking and 
referencing.
The following is an example of accessing a pre-existing server profile and attaching a NTP policy to the server profile.
```hcl-terraform
data "intersight_server_profile" "server1" {
  name = "demo_server_profile"
}

resource "intersight_ntp_policy" "ntp1" {
  name    = "ntp2"
  enabled = true
  ntp_servers = [
    "10.10.10.10",
    "10.10.10.11",
    "10.10.10.12",
    "10.10.10.13"
  ]
  profiles {
    moid        = data.intersight_server_profile.server1.moid
    object_type = "server.Profile"
  }
}
```
A server profile with the given constraints i.e. name as `demo_server_profile` must exist in the Intersight account. Data source 
is used to access this MO. The source is named `server1`. The NTP policy is attached to the server profile with the name 
`server1`. This is done by referencing to the `server1` profile in `profiles.moid`. A resource can have a reference to a 
data source in the format `data.<data_source>.<data_source_name>.<property_name>`. 

###	Viewing the Logs
`TF_LOG` is a terraform variable that is used for viewing different category of logs. By default, it is left empty. To 
view logs for terraform operations, this variable must be set to *DEBUG*.  
#### Mac and Linux Bash
`export TF_LOG=debug`
#### Windows PowerShell
`$env:TF_LOG=“DEBUG”` 
