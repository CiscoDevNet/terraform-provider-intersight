---
subcategory: "niaapi"
layout: "intersight"
page_title: "Intersight: intersight_niaapi_apic_latest_maintained_release"
description: |-
  This contains the latest maintained release information for each release on APIC.
---

# Data Source: intersight_niaapi_apic_latest_maintained_release
This contains the latest maintained release information for each release on APIC.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_niaapi_apic_latest_maintained_release.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `maintained_release`:(string) Lastest maintained release. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `software_release`:(string) Software release version string. 
* `version_tag`:(string) Long lived version or short lived version. 
 