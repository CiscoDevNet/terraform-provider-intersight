---
subcategory: "monitoring"
layout: "intersight"
page_title: "Intersight: intersight_monitoring_health_status"
description: |-
        HealthStatus provides a high-level, aggregated view of overall Intersight component health for a given Intersight account user. It is designed to quickly indicate whether there are issues that require user attention, and to be easily consumed by external dashboards for an at-a-glance status display. The model conforms to the health data API schema published as part of Cisco PlatformSuite.
        #### Purpose
        Expose a single, dashboard-friendly health signal summarizing key Intersight categories—currently including **Licensing**, **Advisories**, and **Alarms**—so users and external monitoring systems can detect account-impacting conditions without querying multiple feature-specific APIs.
        #### Key Concepts
        - **Category-based aggregation**: `categoryStatus` provides a list of health categories (e.g., Advisories, Licensing, Alarms) with their individual statuses and supporting details.
        - **Schema compliance**: `healthDataSchemaVersion` identifies the version of the PlatformSuite-compliant health schema the response adheres to.
        - **Multi-source federation**: `source` is set to `Intersight` to enable external dashboards to combine health signals from multiple platforms (for example, Intersight and other Cisco sources) while preserving attribution.
        - **Near-real-time reporting**: `statusTimeStamp` indicates when the health summary was generated; consumers should account for a potential lag of up to ~5 minutes versus real-time state.

---

# Data Source: intersight_monitoring_health_status
HealthStatus provides a high-level, aggregated view of overall Intersight component health for a given Intersight account user. It is designed to quickly indicate whether there are issues that require user attention, and to be easily consumed by external dashboards for an at-a-glance status display. The model conforms to the health data API schema published as part of Cisco PlatformSuite.
  #### Purpose
  Expose a single, dashboard-friendly health signal summarizing key Intersight categories—currently including **Licensing**, **Advisories**, and **Alarms**—so users and external monitoring systems can detect account-impacting conditions without querying multiple feature-specific APIs.
  #### Key Concepts
  - **Category-based aggregation**: `categoryStatus` provides a list of health categories (e.g., Advisories, Licensing, Alarms) with their individual statuses and supporting details.
  - **Schema compliance**: `healthDataSchemaVersion` identifies the version of the PlatformSuite-compliant health schema the response adheres to.
  - **Multi-source federation**: `source` is set to `Intersight` to enable external dashboards to combine health signals from multiple platforms (for example, Intersight and other Cisco sources) while preserving attribution.
  - **Near-real-time reporting**: `statusTimeStamp` indicates when the health summary was generated; consumers should account for a potential lag of up to ~5 minutes versus real-time state.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_monitoring_health_status.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `health_data_schema_version`:(string) Version of compliant health data API schema. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `nr_source`:(string) Set as 'Intersight'. Especially useful in cases such as when this API is consumed by an external dashboard. This field allows such dashboards to aggregate health status across multiple  sources (Intersight, Meraki etc.). 
* `status_time_stamp`:(string) Time stamp when the status was generated. The status reported by this API may lag the real time status by up to 5 minutes. 
 
