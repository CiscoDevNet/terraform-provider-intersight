---
subcategory: "kvm"
layout: "intersight"
page_title: "Intersight: intersight_kvm_session"
description: |-
  Virtual KVM Session that provides Single Sign-On access to the vKVM console of the server. The vKVM access can be direct or can be tunneled by specifying the tunnel to be used for the access.
---

# Data Source: intersight_kvm_session
Virtual KVM Session that provides Single Sign-On access to the vKVM console of the server. The vKVM access can be direct or can be tunneled by specifying the tunnel to be used for the access.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_kvm_session.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `one_time_password`:(string) Temporary one-time password for KVM access. 
* `sso_supported`:(bool) Indicates if KVM SSO is supported on the server. 
* `username`:(string) Username used for KVM access. 
 