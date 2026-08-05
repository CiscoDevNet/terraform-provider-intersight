---
subcategory: "partnerintegration"
layout: "intersight"
page_title: "Intersight: intersight_partnerintegration_etl"
description: |-
        The Etls object defines the transformation logic used to translate platform-specific API outputs into Intersight-managed objects.
        #### Purpose
        It serves as the transformation engine for inventory data, enabling the system to ingest data from various platforms and normalize it into a common model.
        #### Key Concepts
        - **Data Transformation:** Uses YAML-based models to map raw API data to managed objects.
        - **Normalization:** Ensures consistent data structure across different integrated platforms.
        - **Inventory Integration:** Links transformation definitions to specific inventory collections.

---

# Data Source: intersight_partnerintegration_etl
The Etls object defines the transformation logic used to translate platform-specific API outputs into Intersight-managed objects.
#### Purpose
It serves as the transformation engine for inventory data, enabling the system to ingest data from various platforms and normalize it into a common model.
#### Key Concepts
- **Data Transformation:** Uses YAML-based models to map raw API data to managed objects.
- **Normalization:** Ensures consistent data structure across different integrated platforms.
- **Inventory Integration:** Links transformation definitions to specific inventory collections.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_partnerintegration_etl.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Placeholder name for the ETL. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
