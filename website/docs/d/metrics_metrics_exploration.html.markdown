---
subcategory: "metrics"
layout: "intersight"
page_title: "Intersight: intersight_metrics_metrics_exploration"
description: |-
        The MetricsExploration object is pivotal for querying and analyzing metrics data within the system. This provides a structured approach to defining and executing metrics queries for data exploration and visualization.
        #### Purpose
        MetricsExploration enables users to define complex queries and visualize metrics data for deeper analysis and actionable insights. It supports interactive, user-driven exploration to identify trends, patterns, and opportunities for data-driven decisions.
        #### Key Concepts
        - **Flexible Querying:** Supports a wide range of query configurations and criteria, allowing users to tailor their metrics exploration.
        - **Visualization Options:** Offers diverse visualization configurations, enhancing the interpretability and presentation of metrics data.
        - **Access Control:** Provides controlled access to metrics data queries, ensuring secure and authorized exploration by privileged users.

---

# Data Source: intersight_metrics_metrics_exploration
The MetricsExploration object is pivotal for querying and analyzing metrics data within the system. This provides a structured approach to defining and executing metrics queries for data exploration and visualization.
#### Purpose
MetricsExploration enables users to define complex queries and visualize metrics data for deeper analysis and actionable insights. It supports interactive, user-driven exploration to identify trends, patterns, and opportunities for data-driven decisions.
#### Key Concepts
- **Flexible Querying:** Supports a wide range of query configurations and criteria, allowing users to tailor their metrics exploration.
- **Visualization Options:** Offers diverse visualization configurations, enhancing the interpretability and presentation of metrics data.
- **Access Control:** Provides controlled access to metrics data queries, ensuring secure and authorized exploration by privileged users.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_metrics_metrics_exploration.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `description`:(string) User specified description of the MetricsExploration. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `granularity`:(string) Time unit for aggregating the metrics. 
* `is_widget`:(bool) Set to true when the MetricsExploration is presented as a Dashlet widget. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) User specified name of the MetricsExploration. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
