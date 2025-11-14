---
subcategory: "smtp"
layout: "intersight"
page_title: "Intersight: intersight_smtp_policy_test"
description: |-
        ### Overview
        The PolicyTest object is designed to assist with the testing of SMTP policy compliance within Intersight Appliance, providing a mechanism for Account Administrators to verify the effectiveness of their email notification configurations.
        #### Purpose
        PolicyTest serves as a tool for sending test emails to validate SMTP policy settings. It ensures that email notifications are configured correctly, allowing administrators to troubleshoot and refine their notification strategies.
        #### Key Concepts
        - **SMTP Policy Verification** - Enables testing of SMTP configurations by sending test emails, helping administrators ensure compliance and correct settings.
        - **Account Administrator Access** - Specifically available to Account Administrators, enabling them to manage email settings and perform tests.
        - **Troubleshooting Support** - Provides detailed feedback on test results, including error specifics, to assist in  resolving SMTP policy issues.

---

# Data Source: intersight_smtp_policy_test
### Overview
The PolicyTest object is designed to assist with the testing of SMTP policy compliance within Intersight Appliance, providing a mechanism for Account Administrators to verify the effectiveness of their email notification configurations.
#### Purpose
PolicyTest serves as a tool for sending test emails to validate SMTP policy settings. It ensures that email notifications are configured correctly, allowing administrators to troubleshoot and refine their notification strategies.
#### Key Concepts
- **SMTP Policy Verification** - Enables testing of SMTP configurations by sending test emails, helping administrators ensure compliance and correct settings.
- **Account Administrator Access** - Specifically available to Account Administrators, enabling them to manage email settings and perform tests.
- **Troubleshooting Support** - Provides detailed feedback on test results, including error specifics, to assist in  resolving SMTP policy issues.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_smtp_policy_test.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `enabled`:(bool) If enabled, an attempt will be made to send a test email notification to the configured recipients. 
* `last_status_details`:(string) Information on the last status, including any encountered error specifics, to assist users in troubleshooting SMTP policy issues. 
* `last_successful_attempt_time`:(string) The date and time of the most recent successful attempt to send a test email notification. 
* `last_test_status`:(string) Status of the last test email notification.* `` - Notification has not been tried.* `failed` - Sending notification failed.* `success` - Sending notification successful. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
