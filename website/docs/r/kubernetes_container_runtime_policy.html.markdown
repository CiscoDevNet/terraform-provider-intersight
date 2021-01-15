---
subcategory: "kubernetes"
layout: "intersight"
page_title: "Intersight: intersight_kubernetes_container_runtime_policy"
description: |-
  A policy specifying container runtime configuration, such as docker proxy, no proxy and bridge network IP.
---

# Resource: intersight_kubernetes_container_runtime_policy
A policy specifying container runtime configuration, such as docker proxy, no proxy and bridge network IP.
## Argument Reference
The following arguments are supported:
* `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `cluster_profiles`:(Array) An array of relationships to kubernetesClusterProfile resources. 
This complex property has following sub-properties:
  + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
  + `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data. 
  + `moid`:(string)(Computed) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string)(Computed) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `description`:(string) Description of the policy. 
* `docker_bridge_network_cidr`:(string) The DNS Search Domain Name. 
* `docker_http_proxy`:(Array with Maximum of one item) - The HTTP proxy configuration for docker. Refer to https://docs.docker.com/network/proxy/ for details. 
This complex property has following sub-properties:
  + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
  + `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
  + `hostname`:(string) HTTP/HTTPS Proxy server FQDN or IP. 
  + `is_password_set`:(bool)(Computed) Indicates whether the value of the 'password' property has been set. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
  + `password`:(string) The password for the HTTP/HTTPS Proxy. 
  + `port`:(int) The HTTP Proxy port number.The port number of the HTTP/HTTPS proxy must be between 1 and 65535, inclusive. 
  + `protocol`:(string) Protocol to use for the HTTP/HTTPS Proxy. 
  + `username`:(string) The username for the HTTP/HTTPS Proxy. 
* `docker_https_proxy`:(Array with Maximum of one item) - The https proxy configuration for docker. Refer to https://docs.docker.com/network/proxy/ for details. 
This complex property has following sub-properties:
  + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
  + `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
  + `hostname`:(string) HTTP/HTTPS Proxy server FQDN or IP. 
  + `is_password_set`:(bool)(Computed) Indicates whether the value of the 'password' property has been set. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
  + `password`:(string) The password for the HTTP/HTTPS Proxy. 
  + `port`:(int) The HTTP Proxy port number.The port number of the HTTP/HTTPS proxy must be between 1 and 65535, inclusive. 
  + `protocol`:(string) Protocol to use for the HTTP/HTTPS Proxy. 
  + `username`:(string) The username for the HTTP/HTTPS Proxy. 
* `docker_no_proxy`:
                (Array of schema.TypeString) -
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the concrete policy. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
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


## Import
`intersight_kubernetes_container_runtime_policy` can be imported using the Moid of the object, e.g.
```
$ terraform import intersight_kubernetes_container_runtime_policy.example 1234567890987654321abcde
```