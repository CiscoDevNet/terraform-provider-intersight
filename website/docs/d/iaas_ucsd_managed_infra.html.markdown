
---
layout: "intersight"
page_title: "Intersight: intersight_iaas_ucsd_managed_infra"
sidebar_current: "docs-intersight-data-source-iaas-ucsd-managed-infra"
description: |-
Describes about UCSD Managed infrastructure statistics.
---

# Data Source: intersight_iaas._ucsd_managed_infra
Describes about UCSD Managed infrastructure statistics.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `advanced_catalog_count`:(int) Total advanced catalogs in UCSD. 
* `bm_catalog_count`:(int) Total bare metal catalogs in UCSD. 
* `class_id`:(string) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
* `container_catalog_count`:(int) Total service container catalogs in UCSD. 
* `esxi_host_count`:(int) Total ESXi hosts in UCSD. 
* `external_group_count`:(int) Total external (Ldap) groups in UCSD. 
* `hyperv_host_count`:(int) Total HyperV hosts in UCSD. 
* `local_group_count`:(int) Total local groups in UCSD. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path. 
* `standard_catalog_count`:(int) Total standard catalogs in UCSD. 
* `user_count`:(int) Total user accounts in UCSD. 
* `vdc_count`:(int) Total virtual datacenters in UCSD. 
* `vm_count`:(int) Total Virtual machines in UCSD. 
