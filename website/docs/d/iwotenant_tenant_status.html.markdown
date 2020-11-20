
---
layout: "intersight"
page_title: "Intersight: intersight_iwotenant_tenant_status"
sidebar_current: "docs-intersight-data-source-iwotenant-tenant-status"
description: |-
When an Intersight customer activates IWO license, kubernetes resources need to configured, vault policies need to be setup, etc. to run IWO services as pods. The provisioning of these resources takes a few minutes. Any feature dependent on the IWO APIs will fail till these services are healthy. TenantStatus MO provides status and health of the tenants, so user can be notified when tenant creation is in progress or has failed.
---

# Data Source: intersight_iwotenant._tenant_status
When an Intersight customer activates IWO license, kubernetes resources need to configured, vault policies need to be setup, etc. to run IWO services as pods. The provisioning of these resources takes a few minutes. Any feature dependent on the IWO APIs will fail till these services are healthy. TenantStatus MO provides status and health of the tenants, so user can be notified when tenant creation is in progress or has failed.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `deploy_status`:(string) The deployStatus provides the current status of this deployment.* `NotStarted` - The workflow to deploy the tenant cluster has not yet started.* `InProgress` - The workflow to deploy the tenant cluster in progress. All the tasks required for thesuccessful tenant creation are running.* `Completed` - The workflow to deploy the tenant cluster has completed and health checks have passed.* `Failed` - The workflow to deploy the tenant cluster has failed. Detailed reason for the failure isprovided from Tenant.deployStatus. 
* `iwo_id`:(string) The iwoId uniquely identifies a IWO tenant. The iwoId is used as part of namespace, (logical) database names, policies in vault and many others. As of now, accountMoid has to be provided as the iwoId. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
