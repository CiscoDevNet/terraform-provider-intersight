---
subcategory: "kubernetes"
layout: "intersight"
page_title: "Intersight: intersight_kubernetes_addon_definition"
description: |-
  An addon that can be added to any Kubernetes cluster.
---

# Resource: intersight_kubernetes_addon_definition
An addon that can be added to any Kubernetes cluster.
## Argument Reference
The following arguments are supported:
* `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
* `catalog`:(Array with Maximum of one item) - A reference to a kubernetesCatalog resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
  + `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data. 
  + `moid`:(string)(Computed) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string)(Computed) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `chart_url`:(string) Description of the addon component. 
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data. 
* `default_install_strategy`:(string) Default installation strategy for the release.* `InstallOnly` - Only install in green field. No action in case of failure or removal.* `NoAction` - No install action performed.* `Always` - Attempt install if chart is not already installed. 
* `default_namespace`:(string) Default namespace to install the release. 
* `default_upgrade_strategy`:(string) Default upgrade strategy for the release.* `UpgradeOnly` - Attempt upgrade if chart or overrides options change, no action on upgrade failure.* `NoAction` - This choice enables No upgrades to be performed.* `ReinstallOnFailure` - Attempt upgrade first. Remove and install on upgrade failure.* `AlwaysReinstall` - Always remove older release and reinstall. 
* `description`:(string) Description of the addon component. 
* `digest`:(string) Digest used to verify the integrity of an addon. 
* `icon_url`:(string) Icon used to represent the addon in UI. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of an addon component. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `organization`:(Array with Maximum of one item) - A reference to a organizationOrganization resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
  + `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data. 
  + `moid`:(string)(Computed) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string)(Computed) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `tags`:(Array)
This complex property has following sub-properties:
  + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
  + `key`:(string) The string representation of a tag key. 
  + `value`:(string) The string representation of a tag value. 
* `nr_version`:(string)(Computed) Version of the addon component. 


## Import
`intersight_kubernetes_addon_definition` can be imported using the Moid of the object, e.g.
```
$ terraform import intersight_kubernetes_addon_definition.example 1234567890987654321abcde
```