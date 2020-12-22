---
subcategory: "smtp"
layout: "intersight"
page_title: "Intersight: intersight_smtp_policy"
description: |-
  Name that identifies the SMTP Policy.
---

# Data Source: intersight_smtp_policy
Name that identifies the SMTP Policy.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `description`:(string) Description of the policy. 
* `enabled`:(bool) If enabled, controls the state of the SMTP client service on the managed device. 
* `min_severity`:(string) Minimum fault severity level to receive email notifications. Email notifications are sent for all faults whose severity is equal to or greater than the chosen level.* `critical` - Minimum severity to report is critical.* `condition` - Minimum severity to report is informational.* `warning` - Minimum severity to report is warning.* `minor` - Minimum severity to report is minor.* `major` - Minimum severity to report is major. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the concrete policy. 
* `sender_email`:(string) The email address entered here will be displayed as the from address (mail received from address) of all the SMTP mail alerts that are received. If not configured, the hostname of the server is used in the from address field. 
* `smtp_port`:(int) Port number used by the SMTP server for outgoing SMTP communication. 
* `smtp_server`:(string) IP address or hostname of the SMTP server. The SMTP server is used by the managed device to send email notifications. 
