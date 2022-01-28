---
subcategory: "hyperflex"
layout: "intersight"
page_title: "Intersight: intersight_hyperflex_key_encryption_key"
description: |-
  Specifies a key encryption Key and parameters for the associated resource.
---

# Data Source: intersight_hyperflex_key_encryption_key
Specifies a key encryption Key and parameters for the associated resource.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_hyperflex_key_encryption_key.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `is_account_recovery`:(bool) This defines whether we need to operate in an account recovery scenario or not. If yes, then most of the parameters will be populated from an internal MO. So, some of the input parameters MAY be ignored, if this value is set to true. 
* `is_kek_set`:(bool) Indicates whether the value of the 'kek' property has been set. 
* `is_passphrase_set`:(bool) Indicates whether the value of the 'passphrase' property has been set. 
* `iteration`:(int) Number of iterations we want the hash to be run. 
* `kek`:(string) Key encryption key used to encrypt the DEK's on the HyperFlex cluster. 
* `key_id`:(string) Resource id + time of creation used for retrieving the KEK. 
* `key_state`:(string) Last known Key encryption key state for this Key.* `NEW` - Key Encryption key is newly created.* `ACTIVE` - Key Encryption key is deployed on active resource.* `INACTIVE` - Key Encryption key is inactive and not used.* `INPROGRESS` - Key Encryption key is in a state where it was used on Intersight but did not receive confirmation from platform of success/failure. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `passphrase`:(string) Initial passphrase for the encryption policy, password must contain a minimum of 12 characters, with at least 1 lowercase, 1 uppercase, 1 numeric. 
* `resource_type`:(string) Resource type on which this key will be applied.* `CLUSTER` - Encryption is per HyperFlex cluster.* `DATASTORE` - Encryption is per dataStore on the HyperFlex cluster.* `DRIVE` - Encryption is per drive on the HyperFlex cluster. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `transit_kek`:(string) Copy of Key encryption key, which is used for sending the key over to the remote device endpoint. It is not persisited anywhere. 
 