---
subcategory: "asset"
layout: "intersight"
page_title: "Intersight: intersight_asset_deployment"
description: |-
        Contains information about Deployments associated with consumption-based subscriptions. We listen to messages sent by Cisco Install Base and create/update an instance of this object.

---

# Data Source: intersight_asset_deployment
Contains information about Deployments associated with consumption-based subscriptions. We listen to messages sent by Cisco Install Base and create/update an instance of this object.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_asset_deployment.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `deployment_ref_id`:(string) Identifies the consumption-based subscription's deployment. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `end_date`:(string) End Date for the consumption-based subscription's deployment. 
* `license_type`:(string) Active license tier for the purchased Cisco device's deployment.* `Base` - Base as a License type. It is default license type.* `Essential` - Essential as a License type.* `Standard` - Standard as a License type.* `Advantage` - Advantage as a License type.* `Premier` - Premier as a License type.* `IWO-Essential` - IWO-Essential as a License type.* `IWO-Advantage` - IWO-Advantage as a License type.* `IWO-Premier` - IWO-Premier as a License type.* `IKS-Advantage` - IKS-Advantage as a License type. 
* `mlb_offer_type`:(string) Identifier for consumption based pricing. MLB refers to MultiLine Bundle. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `start_date`:(string) Start Date for the consumption-based subscription's deployment. 
* `subscription_ref_id`:(string) Identifies the consumption-based subscription. 
 
