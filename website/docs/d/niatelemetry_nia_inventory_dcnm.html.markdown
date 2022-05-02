---
subcategory: "niatelemetry"
layout: "intersight"
page_title: "Intersight: intersight_niatelemetry_nia_inventory_dcnm"
description: |-
        Inventory Object available for DCNM-scoped attributes.

---

# Data Source: intersight_niatelemetry_nia_inventory_dcnm
Inventory Object available for DCNM-scoped attributes.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_niatelemetry_nia_inventory_dcnm.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `controller_health`:(int) Health of controller on DCNM. 
* `create_time`:(string) The time when this managed object was created. 
* `dev`:(bool) Returns the value of the dev Field. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `epld_image_count`:(int) Number of EPLD images uploaded to DCNM. 
* `ha_enabled`:(bool) Returns the value of the haEnabled field. 
* `ha_replication_status`:(string) Returns the value of the haReplicationStatus field. 
* `install`:(string) Returns the value of the install field. 
* `installation_type`:(string) Installation type of controller on DCNM. 
* `installation_type_description`:(string) Installation type description of controller on DCNM. 
* `is_isn_configured`:(bool) Returns true if ISN is configured. 
* `is_media_controller`:(bool) Returns the value of the isMediaController field. 
* `is_smart_license_enabled`:(bool) Returns true if the Smart license is enabled and is in use. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `mode`:(string) Mode of controller on DCNM. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `ndfc_fabric_name`:(string) NDFC name information of the setup. 
* `ndfc_oper_state`:(string) NDFC status information for the setup. 
* `num_dcnm_site`:(int) Returns the number of DCNM site fabrics. 
* `num_fabrics`:(int) Returns total number of fabrics in DCNM set-up. 
* `num_fabrics_in_msd`:(int) Returns the number of fabrics in msd. 
* `num_ingress_replication_fabrics`:(int) Returns the number of fabrics that have ingress replication type. 
* `num_local_users`:(int) Returns the number of local users other than admin user. 
* `num_msd`:(int) Returns the number of MSD fabrics. 
* `num_svi_vrf_count`:(int) Returns the number of svi interfaces configured for VRF vlans. 
* `num_trm_enabled_count`:(int) Returns the number of links where TRM is enabled. 
* `num_upg_users`:(int) Number of users who have upgrade privileges excluding the admin. 
* `nxos_image_count`:(int) Number of NXOS images uploaded to DCNM. 
* `outofband_ip`:(string) Out of band IP of controller on DCNM. 
* `serial`:(string) Serial number of device being inventoried. The serial number is unique per device. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `site_name`:(string) Name of fabric domain of the controller. 
* `underlay_peering_active_links_count`:(int) Returns the number of underlay peering active links. 
* `upg_job_count`:(int) Number of upgrade jobs configured on DCNM. 
* `nr_version`:(string) Returns the value of the version field. 
 
