---
subcategory: "hyperflex"
layout: "intersight"
page_title: "Intersight: intersight_hyperflex_witness_configuration"
description: |-
  The witness configuration of the HyperFlex cluster.
Cisco HyperFlex Edge deployments use a witness in their HA arbitration implementations to continue functioning in case of single node failure or network partitions between nodes. HyperFlex administrators have the option to use two different types of witnesses, the Cisco Intersight Invisible Cloud Witness or a local witness deployed within the customers environment.
The type of witness and its corresponding configuration is configured within the HX Connect Device Connector UI component and configuration is stored locally on the device. Changes made locally in the device are reflected in Intersight on this object.
---

# Data Source: intersight_hyperflex_witness_configuration
The witness configuration of the HyperFlex cluster.
Cisco HyperFlex Edge deployments use a witness in their HA arbitration implementations to continue functioning in case of single node failure or network partitions between nodes. HyperFlex administrators have the option to use two different types of witnesses, the Cisco Intersight Invisible Cloud Witness or a local witness deployed within the customers environment.
The type of witness and its corresponding configuration is configured within the HX Connect Device Connector UI component and configuration is stored locally on the device. Changes made locally in the device are reflected in Intersight on this object.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_hyperflex_witness_configuration.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `connection_error`:(string) The detailed connection error to the external witness. Empty if status is connected. 
* `create_time`:(string) The time when this managed object was created. 
* `custom_witness_enabled`:(bool) Custom witness has been configured by user. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `fingerprint`:(string) The fingerprint of the witness server, identifies the revision of the witness servers database. Only applicable if custom witness has been enabled in the cluster, otherwise value is always empty. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `status`:(string) Status of the devices connection to the witness. Device will report status as either 'Connected' or 'NotConnected'. 
* `nr_version`:(string) The version of the custom witness server. Only applicable if custom witness has been enabled in the cluster, otherwise value is always empty. 
* `witness_url`:(string) URL of the witness endpoint, including IP/host and path. Only applicable if custom witness has been enabled in the cluster, otherwise value is always empty. 
 