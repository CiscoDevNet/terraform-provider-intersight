---
subcategory: "iwotenant"
layout: "intersight"
page_title: "Intersight: intersight_iwotenant_tenant_status"
description: |-
        TenantStatuses provide a customer-visible status and health indicator for IWO tenant provisioning. Because IWO features depend on the tenant namespace becoming healthy, this object exposes whether tenant creation is still in progress, completed, or has failed, with a reference timestamp used to detect timeouts during upgrades/reconfiguration.
        #### Purpose
        Expose the operational readiness of an IWO tenant to users and dependent features, enabling UI/automation to inform customers when the tenant is not yet usable or requires attention.
        #### Key Concepts
        - **Provisioning lifecycle state:** Reports deployment progress using a simplified status model (NotStarted/InProgress/Completed/Failed).
        - **Timeout/reference tracking:** Uses `referenceTime` during upgrade/reconfiguration to detect when an operation exceeds expected time bounds.
        - **Tenant identity anchor:** Uses `iwoId` as the key identifier aligned to the tenant namespace/account context.
        - **Correlated relationships:** Links to both the account and the underlying Tenant object for navigation and lifecycle coupling.

---

# Data Source: intersight_iwotenant_tenant_status
TenantStatuses provide a customer-visible status and health indicator for IWO tenant provisioning. Because IWO features depend on the tenant namespace becoming healthy, this object exposes whether tenant creation is still in progress, completed, or has failed, with a reference timestamp used to detect timeouts during upgrades/reconfiguration.
#### Purpose
Expose the operational readiness of an IWO tenant to users and dependent features, enabling UI/automation to inform customers when the tenant is not yet usable or requires attention.
#### Key Concepts
- **Provisioning lifecycle state:** Reports deployment progress using a simplified status model (NotStarted/InProgress/Completed/Failed).
- **Timeout/reference tracking:** Uses `referenceTime` during upgrade/reconfiguration to detect when an operation exceeds expected time bounds.
- **Tenant identity anchor:** Uses `iwoId` as the key identifier aligned to the tenant namespace/account context.
- **Correlated relationships:** Links to both the account and the underlying Tenant object for navigation and lifecycle coupling.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_iwotenant_tenant_status.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `deploy_status`:(string) The deployStatus provides the current status of this deployment.* `NotStarted` - The workflow to deploy the tenant cluster has not yet started.* `InProgress` - The workflow to deploy the tenant cluster in progress. All the tasks required for thesuccessful tenant creation are running.* `Completed` - The workflow to deploy the tenant cluster has completed and health checks have passed.* `Failed` - The workflow to deploy the tenant cluster has failed. Detailed reason for the failure isprovided from Tenant.deployStatus. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `iwo_id`:(string) The iwoId uniquely identifies a IWO tenant. The iwoId is used as part of namespace, (logical) database names, policies in vault and many others. As of now, accountMoid has to be provided as the iwoId. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `reference_time`:(string) During IWO tenant upgrade (or reconfiguration), deployStatus is set to InProgress and referenceTime set to current time. When tenant upgrade (or reconfiguration) does not complete within a pre-defined time using this as reference, deployStatus is set as Failed. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
