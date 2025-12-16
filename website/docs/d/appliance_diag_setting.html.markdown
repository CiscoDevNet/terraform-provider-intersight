---
subcategory: "appliance"
layout: "intersight"
page_title: "Intersight: intersight_appliance_diag_setting"
description: |-
        The DiagSetting object is used to manage the password of the Intersight Appliance's OS diagnostic user account. It is designed to ensure secure access to the appliance's virtual machine for diagnostic purposes.
        #### Purpose
        DiagSetting provides administrators with the ability to change the password of the diagnostic user account, facilitating secure management of the Intersight Appliance's diagnostics. It ensures that only authorized users can access diagnostic features.
        #### Key Concepts
        - **Password Management:** Allows administrators to update the diagnostic user account password, enhancing security.
        - **Access Control:** Restricted to account administrators, ensuring that sensitive operations are conducted securely.
        - **Diagnostic Access:** Supports secure login to the Intersight Appliance virtual machine for diagnostic purposes.

---

# Data Source: intersight_appliance_diag_setting
The DiagSetting object is used to manage the password of the Intersight Appliance's OS diagnostic user account. It is designed to ensure secure access to the appliance's virtual machine for diagnostic purposes.
#### Purpose
DiagSetting provides administrators with the ability to change the password of the diagnostic user account, facilitating secure management of the Intersight Appliance's diagnostics. It ensures that only authorized users can access diagnostic features.
#### Key Concepts
- **Password Management:** Allows administrators to update the diagnostic user account password, enhancing security.
- **Access Control:** Restricted to account administrators, ensuring that sensitive operations are conducted securely.
- **Diagnostic Access:** Supports secure login to the Intersight Appliance virtual machine for diagnostic purposes.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_appliance_diag_setting.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `is_password_set`:(bool) Indicates whether the value of the 'password' property has been set. 
* `message`:(string) Status message of the password change operation. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `password`:(string) Password of the Intersight Appliance's OS diagnostic user account. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
