---
subcategory: "cloud"
layout: "intersight"
page_title: "Intersight: intersight_cloud_aws_vpc"
description: |-
        VPC (Virtual Private Cloud) object in AWS inventory.It is a service that lets you launch AWS resources in a logically isolated
        virtual network.

---

# Data Source: intersight_cloud_aws_vpc
VPC (Virtual Private Cloud) object in AWS inventory.It is a service that lets you launch AWS resources in a logically isolated
virtual network.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_cloud_aws_vpc.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `dns_host_name`:(bool) If true, instances in the vpc get public DNS hostnames. 
* `dns_resolution`:(bool) Indicates whether the Dns resolution is supported. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `identity`:(string) The internally generated identity of this placement. This entity is not manipulated by users. It aids in uniquely identifying the placement object. 
* `is_default`:(bool) If true, indicates that this is default VPC. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the virtual machine placement. It is the name of the VPC (Virtual Private Cloud) in case of AWSvirtual machine, and datacenter name in case of VMware virtual machine. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `state`:(string) The state of the VPC (pending | available). 
* `tenancy`:(string) The allowed tenancy of instances launched into the VPC. 
* `uuid`:(string) The uuid of this placement. The uuid is internally generated and not user specified. 
 
