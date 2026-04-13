---
subcategory: "cloud"
layout: "intersight"
page_title: "Intersight: intersight_cloud_tfc_agentpool"
description: |-
        The TfcAgentpool object is integral to managing groups of agents within Terraform Cloud, enabling communication with isolated, private, or on-premises infrastructure.
        #### Purpose
        Intersight discovers the Agent Pools in the Terraform Cloud account whose credentials are specified in the Terraform Cloud target.
        #### Key Concepts
        - **Agent Grouping:** Represents a collection of Terraform Cloud agents used to execute runs against private, isolated, or on-premises environments.
        - **Credential-Based Discovery:** Intersight identifies available agent pools using the credentials specified in the Terraform Cloud target.

---

# Data Source: intersight_cloud_tfc_agentpool
The TfcAgentpool object is integral to managing groups of agents within Terraform Cloud, enabling communication with isolated, private, or on-premises infrastructure. 
#### Purpose
Intersight discovers the Agent Pools in the Terraform Cloud account whose credentials are specified in the Terraform Cloud target.
#### Key Concepts
- **Agent Grouping:** Represents a collection of Terraform Cloud agents used to execute runs against private, isolated, or on-premises environments.
- **Credential-Based Discovery:** Intersight identifies available agent pools using the credentials specified in the Terraform Cloud target.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_cloud_tfc_agentpool.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `identity`:(string) The ID of the Agent Pool. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) The name of the agent pool. 
* `num_active_agents`:(int) The number of active agents used by this pool. The total active agent are sum of idle, busy and unknown agent counts. 
* `num_tokens`:(int) The number of Tokens in this agent Pool. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
