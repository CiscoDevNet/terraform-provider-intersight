---
subcategory: "kubernetes"
layout: "intersight"
page_title: "Intersight: intersight_kubernetes_addon"
description: |-
  An object that describes an instance of an addon definition with install/upgrade strategies and optional overrides.
---

# Resource: intersight_kubernetes_addon
An object that describes an instance of an addon definition with install/upgrade strategies and optional overrides.
## Argument Reference
The following arguments are supported:
* `addon_definition`:(HashMap) - A reference to a kubernetesAddonDefinition resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string)(Computed) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string)(Computed) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `install_strategy`:(string) Addon install strategy to determine whether an addon is installed if not present.* `InstallOnly` - Only install in green field. No action in case of failure or removal.* `NoAction` - No install action performed.* `Always` - Attempt install if chart is not already installed. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of addon to be installed on a Kubernetes cluster. 
* `organization`:(HashMap) - A reference to a organizationOrganization resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string)(Computed) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string)(Computed) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `override_sets`:(Array)
This complex property has following sub-properties:
  + `key`:(string) Key or property name in a key/value pair. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `value`:(string) Property value in a key/value pair. 
* `overrides`:(string) Properties that can be overridden for an addon. 
* `tags`:(Array)
This complex property has following sub-properties:
  + `key`:(string) The string representation of a tag key. 
  + `value`:(string) The string representation of a tag value. 
* `upgrade_strategy`:(string) Addon upgrade strategy to determine whether an addon configuration is overwritten on upgrade.* `UpgradeOnly` - Attempt upgrade if chart or overrides options change, no action on upgrade failure.* `NoAction` - This choice enables No upgrades to be performed.* `ReinstallOnFailure` - Attempt upgrade first. Remove and install on upgrade failure.* `AlwaysReinstall` - Always remove older release and reinstall. 


## Import
`intersight_kubernetes_addon` can be imported using the Moid of the object, e.g.
```
$ terraform import intersight_kubernetes_addon.example 1234567890987654321abcde
``` 