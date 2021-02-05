---
subcategory: "syslog"
layout: "intersight"
page_title: "Intersight: intersight_syslog_policy"
description: |-
  The syslog policy configure the syslog server to receive CIMC log entries.
---

# Resource: intersight_syslog_policy
The syslog policy configure the syslog server to receive CIMC log entries.
## Argument Reference
The following arguments are supported:
* `description`:(string) Description of the policy. 
* `local_clients`:(Array)
This complex property has following sub-properties:
  + `min_severity`:(string) Lowest level of messages to be included in the local log.* `warning` - Use logging level warning for logs classified as warning.* `emergency` - Use logging level emergency for logs classified as emergency.* `alert` - Use logging level alert for logs classified as alert.* `critical` - Use logging level critical for logs classified as critical.* `error` - Use logging level error for logs classified as error.* `notice` - Use logging level notice for logs classified as notice.* `informational` - Use logging level informational for logs classified as informational.* `debug` - Use logging level debug for logs classified as debug. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the concrete policy. 
* `organization`:(HashMap) - A reference to a organizationOrganization resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string)(Computed) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string)(Computed) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `profiles`:(Array) An array of relationships to policyAbstractConfigProfile resources. 
This complex property has following sub-properties:
  + `moid`:(string)(Computed) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string)(Computed) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `remote_clients`:(Array)
This complex property has following sub-properties:
  + `enabled`:(bool) Enables/disables remote logging for the endpoint If enabled, log messages will be sent to the syslog server mentioned in the Hostname/IP Address field. 
  + `hostname`:(string) Hostname or IP Address of the syslog server where log should be stored. 
  + `min_severity`:(string) Lowest level of messages to be included in the remote log.* `warning` - Use logging level warning for logs classified as warning.* `emergency` - Use logging level emergency for logs classified as emergency.* `alert` - Use logging level alert for logs classified as alert.* `critical` - Use logging level critical for logs classified as critical.* `error` - Use logging level error for logs classified as error.* `notice` - Use logging level notice for logs classified as notice.* `informational` - Use logging level informational for logs classified as informational.* `debug` - Use logging level debug for logs classified as debug. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
  + `port`:(int) Port number used for logging on syslog server. 
  + `protocol`:(string) Transport layer protocol for transmission of log messages to syslog server.* `udp` - Use User Datagram Protocol (UDP) for syslog remote server connection.* `tcp` - Use Transmission Control Protocol (TCP) for syslog remote server connection. 
* `tags`:(Array)
This complex property has following sub-properties:
  + `key`:(string) The string representation of a tag key. 
  + `value`:(string) The string representation of a tag value. 

## Usage Example
### Resource Creation

```hcl
resource "intersight_syslog_policy" "syslog1" {
  name        = "syslog1"
  description = "demo syslog policy"
  local_clients {
    min_severity = "emergency"
    object_type  = "syslog.LocalFileLoggingClient"
  }
  remote_clients {
    enabled      = true
    hostname     = "10.10.10.10"
    port         = 514
    protocol     = "tcp"
    min_severity = "emergency"
    object_type  = "syslog.RemoteLoggingClient"
  }
  remote_clients {
    enabled      = true
    hostname     = "2001:0db8:0a0b:12f0:0000:0000:0000:0004"
    port         = 64000
    protocol     = "udp"
    min_severity = "emergency"
    object_type  = "syslog.RemoteLoggingClient"
  }
  profiles {
    moid        = intersight_server_profile.server1.id
    object_type = "server.Profile"
  }
  organization {
    object_type = "organization.Organization"
    moid = var.organization
  }
}
```

## Import
`intersight_syslog_policy` can be imported using the Moid of the object, e.g.
```
$ terraform import intersight_syslog_policy.example 1234567890987654321abcde
``` 