---
subcategory: "iam"
layout: "intersight"
page_title: "Intersight: intersight_iam_banner_message"
description: |-
  Configuration for the banner message, including title, contents, and display toggle.
---

# Data Source: intersight_iam_banner_message
Configuration for the banner message, including title, contents, and display toggle.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_iam_banner_message.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `banner_contents`:(string) Contents of the banner message. 
* `banner_display`:(bool) Whether or not to display the banner message. 
* `banner_title`:(string) Title of the banner message. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
 