
---
layout: "intersight"
page_title: "Intersight: intersight_hyperflex_node_config_policy"
sidebar_current: "docs-intersight-resource-hyperflex-node-config-policy"
description: |-
  A policy specifying node details such as management and storage data IP ranges. For HyperFlex Edge, storage data IP range is pre-defined.
---

# Resource: intersight_hyperflex._node_config_policy
A policy specifying node details such as management and storage data IP ranges. For HyperFlex Edge, storage data IP range is pre-defined.
## Argument Reference
The following arguments are supported:
* `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `cluster_profiles`:(Array) An array of relationships to hyperflexClusterProfile resources. 
This complex property has following sub-properties:
  + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
  + `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data. 
  + `moid`:(string)(Computed) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string)(Computed) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `data_ip_range`:(Array with Maximum of one item) - The range of storage data IPs to be assigned to the nodes. 
This complex property has following sub-properties:
  + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
  + `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
  + `end_addr`:(string) The end IPv4 address of the range. 
  + `gateway`:(string) The default gateway for the start and end IPv4 addresses. 
  + `netmask`:(string) The netmask specified in dot decimal notation.The start address, end address, and gateway must all be within the network specified by this netmask. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `start_addr`:(string) The start IPv4 address of the range. 
* `description`:(string) Description of the policy. 
* `hxdp_ip_range`:(Array with Maximum of one item) - The range of storage management IPs to be assigned to the nodes. 
This complex property has following sub-properties:
  + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
  + `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
  + `end_addr`:(string) The end IPv4 address of the range. 
  + `gateway`:(string) The default gateway for the start and end IPv4 addresses. 
  + `netmask`:(string) The netmask specified in dot decimal notation.The start address, end address, and gateway must all be within the network specified by this netmask. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `start_addr`:(string) The start IPv4 address of the range. 
* `mgmt_ip_range`:(Array with Maximum of one item) - The range of management IPs to be assigned to the nodes. 
This complex property has following sub-properties:
  + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
  + `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
  + `end_addr`:(string) The end IPv4 address of the range. 
  + `gateway`:(string) The default gateway for the start and end IPv4 addresses. 
  + `netmask`:(string) The netmask specified in dot decimal notation.The start address, end address, and gateway must all be within the network specified by this netmask. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `start_addr`:(string) The start IPv4 address of the range. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the concrete policy. 
* `node_name_prefix`:(string) The node name prefix that is used to automatically generate the default hostname for each server.A dash (-) will be appended to the prefix followed by the node number to form a hostname.This default naming scheme can be manually overridden in the node configuration.The maximum length of a prefix is 60, must only contain alphanumeric characters or dash (-), and muststart with an alphanumeric character. 
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
