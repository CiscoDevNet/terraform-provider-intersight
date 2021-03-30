---
subcategory: "kubernetes"
layout: "intersight"
page_title: "Intersight: intersight_kubernetes_cluster_addon_profile"
description: |-
  A profile that describes a collection of Addons for a particular cluster.
---

# Resource: intersight_kubernetes_cluster_addon_profile
A profile that describes a collection of Addons for a particular cluster.
## Argument Reference
The following arguments are supported:
* `addons`:(Array)
This complex property has following sub-properties:
  + `addon_configuration`:(HashMap) - Addon configuration settings that are specific to a single cluster. 
This complex property has following sub-properties:
    + `install_strategy`:(string) Addon install strategy to determine whether an addon is installed if not present.* `None` - Unspecified install strategy.* `NoAction` - No install action performed.* `InstallOnly` - Only install in green field. No action in case of failure or removal.* `Always` - Attempt install if chart is not already installed. 
    + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
    + `override_sets`:(Array)
This complex property has following sub-properties:
    + `key`:(string) Key or property name in a key/value pair. 
    + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
    + `value`:(string) Property value in a key/value pair. 
  + `overrides`:(string) Properties that can be overridden for an addon. 
  + `release_name`:(string) Name for the helm release. 
  + `release_namespace`:(string) Namespace for the helm release. 
  + `upgrade_strategy`:(string) Addon upgrade strategy to determine whether an addon configuration is overwritten on upgrade.* `None` - Unspecified upgrade strategy.* `NoAction` - This choice enables No upgrades to be performed.* `UpgradeOnly` - Attempt upgrade if chart or overrides options change, no action on upgrade failure.* `ReinstallOnFailure` - Attempt upgrade first. Remove and install on upgrade failure.* `AlwaysReinstall` - Always remove older release and reinstall. 
  + `addon_policy`:(HashMap) - Addon policy associated with the addon. 
This complex property has following sub-properties:
    + `moid`:(string)(Computed) The Moid of the referenced REST resource. 
    + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
    + `selector`:(string)(Computed) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
  + `name`:(string) Name of addon to be installed on a Kubernetes cluster. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
* `associated_cluster`:(HashMap) - A reference to a kubernetesCluster resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string)(Computed) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string)(Computed) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the cluster addon profile. 
* `organization`:(HashMap) - A reference to a organizationOrganization resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string)(Computed) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string)(Computed) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `tags`:(Array)
This complex property has following sub-properties:
  + `key`:(string) The string representation of a tag key. 
  + `value`:(string) The string representation of a tag value. 


## Import
`intersight_kubernetes_cluster_addon_profile` can be imported using the Moid of the object, e.g.
```
$ terraform import intersight_kubernetes_cluster_addon_profile.example 1234567890987654321abcde
``` 