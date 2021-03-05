---
subcategory: "hyperflex"
layout: "intersight"
page_title: "Intersight: intersight_hyperflex_node_profile"
description: |-
  A configuration profile per node in the HyperFlex cluster.
It defines node settings such as IP address configuration for hypervisor management network, storage data network, HyperFlex management network, and the assigned physical server.
---

# Data Source: intersight_hyperflex_node_profile
A configuration profile per node in the HyperFlex cluster.
It defines node settings such as IP address configuration for hypervisor management network, storage data network, HyperFlex management network, and the assigned physical server.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_hyperflex_node_profile.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `description`:(string) Description of the profile. 
* `hxdp_data_ip`:(string) IP address for storage data network (Controller VM interface). 
* `hxdp_mgmt_ip`:(string) IP address for HyperFlex management network. 
* `hypervisor_control_ip`:(string) IP address for hypervisor control such as VM migration or pod management. 
* `hypervisor_data_ip`:(string) IP address for storage data network (Hypervisor interface). 
* `hypervisor_mgmt_ip`:(string) IP address for Hypervisor management network. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the concrete profile. 
* `type`:(string) Defines the type of the profile. Accepted value is instance.* `instance` - The profile defines the configuration for a specific instance of a target. 
 