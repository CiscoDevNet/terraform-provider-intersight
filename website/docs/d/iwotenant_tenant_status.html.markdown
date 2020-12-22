---
subcategory: "iwotenant"
layout: "intersight"
page_title: "Intersight: intersight_iwotenant_tenant_status"
description: |-
  When an Intersight customer activates IWO license, kubernetes resources need to configured, vault policies need to be setup, etc. to run IWO services as pods. The provisioning of these resources takes a few minutes. Any feature dependent on the IWO APIs will fail till these services are healthy. TenantStatus MO provides status and health of the tenants, so user can be notified when tenant creation is in progress or has failed.
---

# Data Source: intersight_iwotenant_tenant_status
When an Intersight customer activates IWO license, kubernetes resources need to configured, vault policies need to be setup, etc. to run IWO services as pods. The provisioning of these resources takes a few minutes. Any feature dependent on the IWO APIs will fail till these services are healthy. TenantStatus MO provides status and health of the tenants, so user can be notified when tenant creation is in progress or has failed.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `deploy_status`:(string) The deployStatus provides the current status of this deployment.* `NotStarted` - The workflow to deploy the tenant cluster has not yet started.* `InProgress` - The workflow to deploy the tenant cluster in progress. All the tasks required for thesuccessful tenant creation are running.* `Completed` - The workflow to deploy the tenant cluster has completed and health checks have passed.* `Failed` - The workflow to deploy the tenant cluster has failed. Detailed reason for the failure isprovided from Tenant.deployStatus. 
* `iwo_id`:(string) The iwoId uniquely identifies a IWO tenant. The iwoId is used as part of namespace, (logical) database names, policies in vault and many others. As of now, accountMoid has to be provided as the iwoId. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `reference_time`:(string) During IWO tenant upgrade (or reconfiguration), deployStatus is set to InProgress and referenceTime set to current time. When tenant upgrade (or reconfiguration) does not complete within a pre-defined time using this as reference, deployStatus is set as Failed. 
