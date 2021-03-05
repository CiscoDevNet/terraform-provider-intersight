---
subcategory: "appliance"
layout: "intersight"
page_title: "Intersight: intersight_appliance_release_note"
description: |-
  ReleaseUpdate managed the object provides the information preview (new features and bug fixes) for one pending upgrade.
---

# Data Source: intersight_appliance_release_note
ReleaseUpdate managed the object provides the information preview (new features and bug fixes) for one pending upgrade.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_appliance_release_note.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `nr_version`:(string) Version number of the pending upgrade. 
 