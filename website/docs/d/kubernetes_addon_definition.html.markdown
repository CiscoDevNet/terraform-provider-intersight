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
The following arguments can be used to get data of already created objects in Intersight appliance:
* `chart_url`:(string) Description of the addon component. 
* `default_install_strategy`:(string) Default installation strategy for the release.* `InstallOnly` - Only install in green field. No action in case of failure or removal.* `NoAction` - No install action performed.* `Always` - Attempt install if chart is not already installed. 
* `default_namespace`:(string) Default namespace to install the release. 
* `default_upgrade_strategy`:(string) Default upgrade strategy for the release.* `UpgradeOnly` - Attempt upgrade if chart or overrides options change, no action on upgrade failure.* `NoAction` - This choice enables No upgrades to be performed.* `ReinstallOnFailure` - Attempt upgrade first. Remove and install on upgrade failure.* `AlwaysReinstall` - Always remove older release and reinstall. 
* `description`:(string) Description of the addon component. 
* `digest`:(string) Digest used to verify the integrity of an addon. 
* `icon_url`:(string) Icon used to represent the addon in UI. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of an addon component. 
* `nr_version`:(string) Version of the addon component. 
