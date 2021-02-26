---
subcategory: "niatelemetry"
layout: "intersight"
page_title: "Intersight: intersight_niatelemetry_apic_ui_page_counts"
description: |-
  Object to capture the UI page counts in APIC.
---

# Data Source: intersight_niatelemetry_apic_ui_page_counts
Object to capture the UI page counts in APIC.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `page_count`:(int) Number of times that the user has opened this page. 
* `page_name`:(string) Name of the page for which page count is recorded. 
* `record_type`:(string) Type of record DCNM / APIC / SE. This determines the type of platform where inventory was collected. 
* `record_version`:(string) Version of record being pushed. This determines what was the API version for data available from the device. 
* `site_name`:(string) Name of the APIC site from which this data is being collected. 
