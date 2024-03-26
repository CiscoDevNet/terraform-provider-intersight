---
subcategory: "resourcepool"
layout: "intersight"
page_title: "Intersight: intersight_resourcepool_membership_reservation"
description: |-
        A MembershipReservation is created when a resource that belongs to a pool is decommissioned. The MembershipReservation is mapped to pools that have the resource chosen for decommissioning. This MembershipReservation will be utilized in the future to pre-provision servers based on membership serials.
        The system automatically generates the membership during decommissioning, and its permissions are updated for all organizations associated with the pools. Users of the pool have the authority to remove the pool from the membership. The membership will be removed either during recommissioning or when all the associated pools are deleted.

---

# Data Source: intersight_resourcepool_membership_reservation
A MembershipReservation is created when a resource that belongs to a pool is decommissioned. The MembershipReservation is mapped to pools that have the resource chosen for decommissioning. This MembershipReservation will be utilized in the future to pre-provision servers based on membership serials.
The system automatically generates the membership during decommissioning, and its permissions are updated for all organizations associated with the pools. Users of the pool have the authority to remove the pool from the membership. The membership will be removed either during recommissioning or when all the associated pools are deleted.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_resourcepool_membership_reservation.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `description`:(string) Details of the use case for which the reservation was created, such as decommissioning. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `expiration`:(string) The resource reservation includes an expiration date and a timestamp indicating when this management object will be cleared. The expiration date is set during the decommissioning process and is maintained for a period of 3 months. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `reservation_selector`:(string) The unique identification of the resource is based on the resource OData string, which is mentioned as part of the ReservationSelector. For example, 'Serial eq 'EM6259AE6B'. 
* `resource_type`:(string) The type of resource that is placed into resource groups or pools. Resource Type can be either 'compute.Blade' or 'compute.RackUnit for pools. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `status`:(string) The reservation status can be in the 'Created', 'Processing', 'Failed', or 'Finished' state.* `Created` - By default, a reservation is in Created status.* `Processing` - A reservation is changed to Processing status for appliance mode resource claim requests.* `Failed` - A reservation is changed to Failed status if the validations on resources, resource groups fails.* `Finished` - A reservation is changed to Finished status if the validations on resources, resource groups are successful. The resource moids in reservation will be added to resource groups using OData filters. 
 
