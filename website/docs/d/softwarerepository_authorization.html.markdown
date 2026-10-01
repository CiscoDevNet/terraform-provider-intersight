---
subcategory: "softwarerepository"
layout: "intersight"
page_title: "Intersight: intersight_softwarerepository_authorization"
description: |-
        The Authorizations object tracks user consent for Intersight to contact external software repositories (e.g., cisco.com) to download images on the user's behalf.
        #### Purpose
        It manages the secure storage of credentials and authentication modes required to access external repositories, ensuring that software downloads are authorized and secure.
        #### Key Concepts
        - **Credential Management:** Securely manages usernames and passwords for external repositories.
        - **Authentication Modes:** Defines the OAuth2 mechanisms used for secure communication.
        - **Repository Access:** Authorizes Intersight to act on behalf of the user account for software retrieval.

---

# Data Source: intersight_softwarerepository_authorization
The Authorizations object tracks user consent for Intersight to contact external software repositories (e.g., cisco.com) to download images on the user's behalf.
#### Purpose
It manages the secure storage of credentials and authentication modes required to access external repositories, ensuring that software downloads are authorized and secure.
#### Key Concepts
- **Credential Management:** Securely manages usernames and passwords for external repositories.
- **Authentication Modes:** Defines the OAuth2 mechanisms used for secure communication.
- **Repository Access:** Authorizes Intersight to act on behalf of the user account for software retrieval.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_softwarerepository_authorization.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `is_asdv4_alarm_dismissed`:(bool) The state of the alarm dismissal for the 'IsAsdDialogDismissed' alarm. 
* `is_password_set`:(bool) Indicates whether the value of the 'password' property has been set. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `password`:(string) The password that will be used by Intersight to create OAuth2 tokens for interacting with the external repository, on the user account's behalf. 
* `repository_type`:(string) The external repository for which this authorization has been provided. The only supported repository today is cisco.com.* `Cisco` - External repository hosted on cisco.com.* `IntersightCloud` - Repository hosted by the Intersight Cloud.* `LocalMachine` - The file is available on the local client machine. Used as an upload source type.* `NetworkShare` - External repository in the customer datacenter. This will typically be a file server. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `user_id`:(string) The username that will be used by Intersight to create OAuth2 tokens for interacting with the external repository, on the user account's behalf. 
* `nr_version`:(string) The Automated Software Distribution version of the authorization MO.* `V3` - The client is running Automated Software Distribution V3.* `V4` - The client is running Automated Software Distribution V4. 
 
