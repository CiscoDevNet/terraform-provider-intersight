---
subcategory: "hyperflex"
layout: "intersight"
page_title: "Intersight: intersight_hyperflex_hxdp_version"
description: |-
        HyperFlex software version records represent the desired HyperFlex software stack versions to be applied to a HyperFlex cluster, including the HyperFlex Data Platform, server firmware, and hypervisor versions. They also capture which components are intended to be upgraded and provide linkage to bundle metadata for the selected software release.
        #### Purpose
        Provide a version intent record that HyperFlex cluster profile workflows can use to drive consistent installation and upgrade behavior across the HyperFlex Data Platform, server firmware, and hypervisor components, while enabling correlation to the underlying distributable and bundle metadata.
        #### Key Concepts
        - **Multi-component version intent**: Captures target versions across major stack layers:
        - **HyperFlex Data Platform version**: The target platform software version.
        - `serverFirmwareVersion` / `serverFirmwareVersions` (UCS firmware baseline and per-server-type details)
        - `hypervisorVersion` (applies across cluster nodes)
        - **Upgrade scoping**: `upgradeTypes` indicates which components are in scope for an upgrade operation (enabling partial vs full stack upgrades).
        - **Validation for firmware strings**: `serverFirmwareVersion` includes a regex constraint supporting both release-style and internal build version formats.
        - **Profile consumption and reuse**: `clusterProfiles` relates the version intent to the set of cluster profiles that use it (with `ondelete: unset`, so profiles can outlive a particular version record).
        - **Bundle metadata correlation**: Associates the selected component versions with the corresponding distributable and bundle metadata.

---

# Data Source: intersight_hyperflex_hxdp_version
HyperFlex software version records represent the desired HyperFlex software stack versions to be applied to a HyperFlex cluster, including the HyperFlex Data Platform, server firmware, and hypervisor versions. They also capture which components are intended to be upgraded and provide linkage to bundle metadata for the selected software release.
#### Purpose
Provide a version intent record that HyperFlex cluster profile workflows can use to drive consistent installation and upgrade behavior across the HyperFlex Data Platform, server firmware, and hypervisor components, while enabling correlation to the underlying distributable and bundle metadata.
#### Key Concepts
- **Multi-component version intent**: Captures target versions across major stack layers:
  - **HyperFlex Data Platform version**: The target platform software version.
  - `serverFirmwareVersion` / `serverFirmwareVersions` (UCS firmware baseline and per-server-type details)
  - `hypervisorVersion` (applies across cluster nodes)
- **Upgrade scoping**: `upgradeTypes` indicates which components are in scope for an upgrade operation (enabling partial vs full stack upgrades).
- **Validation for firmware strings**: `serverFirmwareVersion` includes a regex constraint supporting both release-style and internal build version formats.
- **Profile consumption and reuse**: `clusterProfiles` relates the version intent to the set of cluster profiles that use it (with `ondelete: unset`, so profiles can outlive a particular version record).
- **Bundle metadata correlation**: Associates the selected component versions with the corresponding distributable and bundle metadata.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_hyperflex_hxdp_version.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `nr_version`:(string) The HyperFlex Data Platform version. 
 
