---
subcategory: "hyperflex"
layout: "intersight"
page_title: "Intersight: intersight_hyperflex_service_auth_token"
description: |-
  A Service auth token entity that represents HyperFlex Data Platform service AAA token.
---

# Data Source: intersight_hyperflex_service_auth_token
A Service auth token entity that represents HyperFlex Data Platform service AAA token.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_hyperflex_service_auth_token.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `client_id`:(string) Client Id or tenant Id of the entity that uses the service auth token. 
* `create_time`:(string) The time when this managed object was created. 
* `csi_version`:(string) Version of Container Storage Interface (CSI) that the tokenOwner is associated with. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `service_auth_token`:(string) Service auth token that has been created by HyperFlex cluster. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `status`:(string) Represents status of ervice auth claim or revocation.* `Unknown` - Unknown claim state of the service auth token.* `Claiming` - The service auth token claim is in progress.* `Claimed` - The service auth token has been successfully claimed.* `FailedToClaim` - Cannot claim the service auth token on the underlying HyperFlex cluster.* `Revoking` - The service auth token revocation is in progress.* `Revoked` - The service auth token revocation has been successfully revoked.* `FailedToRevoke` - Cannot revoke the service auth token on the underlying HyperFlex cluster. 
 