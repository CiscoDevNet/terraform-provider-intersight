---
subcategory: "kubernetes"
layout: "intersight"
page_title: "Intersight: intersight_kubernetes_addon"
description: |-
  An object that describes an instance of an addon definition with install/upgrade strategies and optional overrides.
---

# Data Source: intersight_kubernetes_addon
An object that describes an instance of an addon definition with install/upgrade strategies and optional overrides.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `install_strategy`:(string) Addon install strategy to determine whether an addon is installed if not present.* `InstallOnly` - Only install in green field. No action in case of failure or removal.* `NoAction` - No install action performed.* `Always` - Attempt install if chart is not already installed. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of addon to be installed on a Kubernetes cluster. 
* `overrides`:(string) Properties that can be overridden for an addon. 
* `upgrade_strategy`:(string) Addon upgrade strategy to determine whether an addon configuration is overwritten on upgrade.* `UpgradeOnly` - Attempt upgrade if chart or overrides options change, no action on upgrade failure.* `NoAction` - This choice enables No upgrades to be performed.* `ReinstallOnFailure` - Attempt upgrade first. Remove and install on upgrade failure.* `AlwaysReinstall` - Always remove older release and reinstall. 
