---
subcategory: "asset"
layout: "intersight"
page_title: "Intersight: intersight_asset_claim_token"
description: |-
        The ClaimToken object serves as a secure mechanism for onboarding and registering devices to an Intersight user account. It facilitates a streamlined connection process, allowing device connectors to authenticate and associate themselves with a specific user environment without requiring complex manual configuration or external dial-out procedures.
        #### Purpose
        A ClaimToken acts as a bridge between a physical or virtual device and the Intersight platform. Supplies the necessary credentials and instructions for a device connector to establish a secure link to the user's account, ensuring that device claiming is performed in a controlled and authorized manner.
        #### Key Concepts
        - **Simplified Onboarding:** Enables the claiming of devices into an Intersight account by providing a pre-configured token, removing the need for the appliance to perform manual outbound configuration. - **Security and Verification:** Incorporates cryptographic signatures to verify the legitimacy of the claim, ensuring that only authorized devices can be associated with an account. - **Access Control:** Management of claim tokens is restricted to specific privilege sets, such as Account Administrators and users with dedicated Claim Devices or Claim targets permissions, ensuring that only authorized personnel can initiate device registration. - **Lifecycle Management:** Includes built-in expiration logic to maintain security, ensuring that tokens remain valid only for a defined period and reducing the risk associated with long-lived or unused tokens.

---

# Data Source: intersight_asset_claim_token
The ClaimToken object serves as a secure mechanism for onboarding and registering devices to an Intersight user account. It facilitates a streamlined connection process, allowing device connectors to authenticate and associate themselves with a specific user environment without requiring complex manual configuration or external dial-out procedures.
#### Purpose
A ClaimToken acts as a bridge between a physical or virtual device and the Intersight platform. Supplies the necessary credentials and instructions for a device connector to establish a secure link to the user's account, ensuring that device claiming is performed in a controlled and authorized manner.
#### Key Concepts
- **Simplified Onboarding:** Enables the claiming of devices into an Intersight account by providing a pre-configured token, removing the need for the appliance to perform manual outbound configuration. - **Security and Verification:** Incorporates cryptographic signatures to verify the legitimacy of the claim, ensuring that only authorized devices can be associated with an account. - **Access Control:** Management of claim tokens is restricted to specific privilege sets, such as Account Administrators and users with dedicated Claim Devices or Claim targets permissions, ensuring that only authorized personnel can initiate device registration. - **Lifecycle Management:** Includes built-in expiration logic to maintain security, ensuring that tokens remain valid only for a defined period and reducing the risk associated with long-lived or unused tokens.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_asset_claim_token.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `claim`:(string) The claim generated to configure the device to connect and claim to the user account. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `expiration_date`:(string) The expiry time of the claim token, after this time any generated claims from this token are no longer valid. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
