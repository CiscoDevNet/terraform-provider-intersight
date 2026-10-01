---
subcategory: "kvm"
layout: "intersight"
page_title: "Intersight: intersight_kvm_session"
description: |-
        The Session object represents a virtual KVM (vKVM) session that provides Single Sign-On (SSO) access to the vKVM console of a managed server.
        #### Purpose
        The Session object is designed to facilitate secure, remote console access to servers. It supports both direct vKVM connections and tunneled connections, providing a flexible and secure way for administrators to interact with server hardware remotely.
        #### Key Concepts
        - **Single Sign-On (SSO) Access:** Provides seamless authentication to the vKVM console, reducing the need for manual login steps at the server level.
        - **Connection Flexibility:** Supports both direct access and tunneled access, allowing for secure connections even in restricted network environments.
        - **Role-Based Access Control:** Integrates with specific privilege sets, ensuring that only authorized users can launch vKVM sessions or manage/terminate existing sessions.
        - **Session Lifecycle:** Captures unique session identifiers, launch URLs, and one-time passwords (OTP) to maintain secure, temporary access tokens for console sessions.
        - **Infrastructure Integration:** Establishes clear relationships with the target server and the specific network tunnel used for the connection, ensuring auditability and proper resource mapping.

---

# Data Source: intersight_kvm_session
The Session object represents a virtual KVM (vKVM) session that provides Single Sign-On (SSO) access to the vKVM console of a managed server.
#### Purpose
The Session object is designed to facilitate secure, remote console access to servers. It supports both direct vKVM connections and tunneled connections, providing a flexible and secure way for administrators to interact with server hardware remotely.
#### Key Concepts
- **Single Sign-On (SSO) Access:** Provides seamless authentication to the vKVM console, reducing the need for manual login steps at the server level.
- **Connection Flexibility:** Supports both direct access and tunneled access, allowing for secure connections even in restricted network environments.
- **Role-Based Access Control:** Integrates with specific privilege sets, ensuring that only authorized users can launch vKVM sessions or manage/terminate existing sessions.
- **Session Lifecycle:** Captures unique session identifiers, launch URLs, and one-time passwords (OTP) to maintain secure, temporary access tokens for console sessions.
- **Infrastructure Integration:** Establishes clear relationships with the target server and the specific network tunnel used for the connection, ensuring auditability and proper resource mapping.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_kvm_session.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `client_ip_address`:(string) The user agent IP address from which the session is launched. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `end_time`:(string) The time at which the session ended. 
* `kvm_launch_url_path`:(string) One time URL that is used to launch the vKVM console. 
* `kvm_session_id`:(string) Unique ID of the KVM Session URI. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `one_time_password`:(string) Temporary one-time password for vKVM access. 
* `role`:(string) Role of the user who launched the session. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `sso_supported`:(bool) Indicates if vKVM SSO is supported on the server. 
* `status`:(string) The status of the session.* `Active` - The session is currently active.* `Ended` - The session has ended normally.* `Terminated` - The session was terminated by an admin. 
* `target_name`:(string) Name of target on which session is initiated. 
* `user_id_or_email`:(string) User ID or E-mail Address of the user who launched the session. 
* `username`:(string) Username used for vKVM access. 
 
