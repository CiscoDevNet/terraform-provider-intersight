---
subcategory: "tam"
layout: "intersight"
page_title: "Intersight: intersight_tam_advisory_info"
description: |-
  State of an advisory in the context of a given account. Used to capture a given account's preferences regarding  associated advisory.
---

# Data Source: intersight_tam_advisory_info
State of an advisory in the context of a given account. Used to capture a given account's preferences regarding  associated advisory.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `state`:(string) Current state of the advisory for the owner. Indicates if the user is interested in getting updates for the advisory.* `active` - Advisory is currently active and the user wants to receive updates for this advisory.* `acknowledged` - Advisory is seen and acknowledged by the user and she no longer wants to recieve updates. 
