---
subcategory: "hyperflex"
layout: "intersight"
page_title: "Intersight: intersight_hyperflex_key_encryption_key"
description: |-
        Key Encryption Key and associated resource parameters specification.

---

# Data Source: intersight_hyperflex_key_encryption_key
Key Encryption Key and associated resource parameters specification.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_hyperflex_key_encryption_key.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `is_account_recovery`:(bool) Account recovery scenario flag determining parameter population from internal MO, with possibility of ignoring of input parameters when this value is true. 
* `is_kek_set`:(bool) Indicates whether the value of the 'kek' property has been set. 
* `is_passphrase_set`:(bool) Indicates whether the value of the 'passphrase' property has been set. 
* `iteration`:(int) Number of hash iterations to run. 
* `kek`:(string) Key Encryption Key used to encrypt the DEK's on the HyperFlex cluster. 
* `key_id`:(string) Resource ID and time for Kek retrieval. 
* `key_state`:(string) Last known state of the Key Encryption Key.* `NEW` - Newly created Key Encryption Key (KEK).* `ACTIVE` - Deployed Key Encryption Key on active resources.* `INACTIVE` - Inactive and unused Key Encryption Key.* `INPROGRESS` - Unconfirmed Key Encryption Key usage on Intersight platform. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `passphrase`:(string) Initial passphrase for encryption policy, requiring a minimum of 12 characters, including 1 lowercase, 1 uppercase, and 1 numeric character. 
* `resource_type`:(string) Resource type for key application.* `CLUSTER` - Cluster specific encryption per HyperFlex cluster.* `DATASTORE` - Data store encryption on the HyperFlex cluster.* `DRIVE` - Drive specific encryption on the HyperFlex cluster. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `transit_kek`:(string) Temporary copy of KEK used for transfer to remote device endpoint, not persisted anywhere. 
 
