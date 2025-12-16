---
subcategory: "iwotenant"
layout: "intersight"
page_title: "Intersight: intersight_iwotenant_tenant_customization"
description: |-
        ### Overview
        The TenantCustomization object allows end-users to tailor their IWO tenant configurations, enhancing functionality through options like enabling data extractors.
        #### Purpose
        The TenantCustomization object provides mechanisms for users to customize their tenant environments, offering flexibility in data handling and feature management.
        #### Key Concepts
        - **Data Management:** - Enables the activation of data extractors, facilitating reporting and analysis within the tenant's namespace.
        - **Customization Flexibility:** - Provides options for users to enhance tenant capabilities according to their specific requirements.
        - **Account Linkage:** - Maintains associations with relevant accounts and tenants, ensuring consistent customization across linked entities.

---

# Data Source: intersight_iwotenant_tenant_customization
### Overview
The TenantCustomization object allows end-users to tailor their IWO tenant configurations, enhancing functionality through options like enabling data extractors.
#### Purpose
The TenantCustomization object provides mechanisms for users to customize their tenant environments, offering flexibility in data handling and feature management. 
#### Key Concepts
- **Data Management:** - Enables the activation of data extractors, facilitating reporting and analysis within the tenant's namespace.
- **Customization Flexibility:** - Provides options for users to enhance tenant capabilities according to their specific requirements.
- **Account Linkage:** - Maintains associations with relevant accounts and tenants, ensuring consistent customization across linked entities.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_iwotenant_tenant_customization.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `enable_data_extractor`:(bool) Enables IWO tenant data to be available for reporting.  This will start 'extractor' pod. 
* `is_write_user_access_key_id_set`:(bool) Indicates whether the value of the 'writeUserAccessKeyId' property has been set. 
* `is_write_user_secret_access_key_set`:(bool) Indicates whether the value of the 'writeUserSecretAccessKey' property has been set. 
* `iwo_id`:(string) The iwoId uniquely identifies a IWO tenant. The iwoId is used as part of namespace, (logical) database names, policies in vault and many others. As of now, accountMoid has to be provided as the iwoId. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `msk_server_for_data_extractor`:(string) MSK cluster endpoint that data extractor can send reporting data to. This  MS cluster in turn populates data into tables in Redshift cluster. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `write_user_access_key_id`:(string) AWS access key Id to write data to redshift.  Refer to AWS cloud formation stack 'Output' of the tenant. 
* `write_user_secret_access_key`:(string) AWS secret access key to write data to redshift.  Refer to AWS cloud formation stack 'Output' of the tenant. 
 
