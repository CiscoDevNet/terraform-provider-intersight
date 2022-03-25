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
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_smtp_policy.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `auth_password`:(string) Authorization password for the process. 
* `create_time`:(string) The time when this managed object was created. 
* `description`:(string) Description of the policy. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `enable_auth`:(bool) If enabled, lets user input username and password. 
* `enable_tls`:(bool) If enabled, lets user input valid CA certificates for authorization. 
* `enabled`:(bool) If enabled, controls the state of the SMTP client service on the managed device. 
* `is_auth_password_set`:(bool) Indicates whether the value of the 'authPassword' property has been set. 
* `min_severity`:(string) Minimum fault severity level to receive email notifications. Email notifications are sent for all faults whose severity is equal to or greater than the chosen level.* `critical` - Minimum severity to report is critical.* `condition` - Minimum severity to report is informational.* `warning` - Minimum severity to report is warning.* `minor` - Minimum severity to report is minor.* `major` - Minimum severity to report is major. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the concrete policy. 
* `sender_email`:(string) The email address entered here will be displayed as the from address (mail received from address) of all the SMTP mail alerts that are received. If not configured, the hostname of the server is used in the from address field. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `smtp_port`:(int) Port number used by the SMTP server for outgoing SMTP communication. 
* `smtp_server`:(string) IP address or hostname of the SMTP server. The SMTP server is used by the managed device to send email notifications. 
* `user_name`:(string) SMTP username from which email notification is sent. 
 
