---
subcategory: "kubernetes"
layout: "intersight"
page_title: "Intersight: intersight_kubernetes_addon_definition"
description: |-
  An addon that can be added to any Kubernetes cluster.
---

# Data Source: intersight_kubernetes_addon_definition
An addon that can be added to any Kubernetes cluster.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_kubernetes_addon_definition.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `chart_url`:(string) Description of the addon component. 
* `create_time`:(string) The time when this managed object was created. 
* `default_install_strategy`:(string) Default installation strategy for the release.* `None` - Unspecified install strategy.* `NoAction` - No install action performed.* `InstallOnly` - Only install in green field. No action in case of failure or removal.* `Always` - Attempt install if chart is not already installed. 
* `default_namespace`:(string) Default namespace to install the release. 
* `default_upgrade_strategy`:(string) Default upgrade strategy for the release.* `None` - Unspecified upgrade strategy.* `NoAction` - This choice enables No upgrades to be performed.* `UpgradeOnly` - Attempt upgrade if chart or overrides options change, no action on upgrade failure.* `ReinstallOnFailure` - Attempt upgrade first. Remove and install on upgrade failure.* `AlwaysReinstall` - Always remove older release and reinstall. 
* `description`:(string) Description of the addon component. 
* `digest`:(string) Digest used to verify the integrity of an addon. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `icon_url`:(string) Icon used to represent the addon in UI. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of an addon component. 
* `overrides`:(string) Properties that can be overridden for an addon. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `nr_version`:(string) Version of the addon component. 
 