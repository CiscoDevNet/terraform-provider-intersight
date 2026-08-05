---
subcategory: "hyperflex"
layout: "intersight"
page_title: "Intersight: intersight_hyperflex_witness_configuration"
description: |-
        The WitnessConfiguration object pertains to the witness configuration of the HyperFlex cluster, especially in Cisco HyperFlex Edge deployments. It ensures high availability through witness-based arbitration mechanisms.
        #### Purpose
        WitnessConfiguration is critical in maintaining the operational integrity of HyperFlex Edge deployments by providing arbitration capabilities in case of node failures or network partitions. It ensures the cluster continues to function effectively under adverse conditions.
        #### Key Concepts
        - **High Availability:** Supports HA arbitration implementations to maintain cluster operations during single-node failures or network issues.
        - **Witness Types:** Offers flexibility with the option to use the Cisco Intersight Invisible Cloud Witness or a locally deployed witness.
        - **Configuration Management:** Allows administrators to configure and manage witness settings, ensuring proper cluster arbitration and failover mechanisms.
        - **Integration:** Seamlessly integrates with the HyperFlex cluster's infrastructure, providing essential support for high availability and resilience.

---

# Data Source: intersight_hyperflex_witness_configuration
The WitnessConfiguration object pertains to the witness configuration of the HyperFlex cluster, especially in Cisco HyperFlex Edge deployments. It ensures high availability through witness-based arbitration mechanisms.
#### Purpose
WitnessConfiguration is critical in maintaining the operational integrity of HyperFlex Edge deployments by providing arbitration capabilities in case of node failures or network partitions. It ensures the cluster continues to function effectively under adverse conditions.
#### Key Concepts
- **High Availability:** Supports HA arbitration implementations to maintain cluster operations during single-node failures or network issues.
- **Witness Types:** Offers flexibility with the option to use the Cisco Intersight Invisible Cloud Witness or a locally deployed witness.
- **Configuration Management:** Allows administrators to configure and manage witness settings, ensuring proper cluster arbitration and failover mechanisms.
- **Integration:** Seamlessly integrates with the HyperFlex cluster's infrastructure, providing essential support for high availability and resilience.
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
 
