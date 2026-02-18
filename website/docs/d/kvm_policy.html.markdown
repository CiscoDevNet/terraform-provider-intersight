---
subcategory: "kvm"
layout: "intersight"
page_title: "Intersight: intersight_kvm_policy"
description: |-
        The Virtual KVM Policy is a reusable policy used to configure the virtual Keyboard, Video, and Mouse (KVM) settings for a server.
        #### Purpose
        The purpose of this policy is to standardize and control remote console access to servers. It allows administrators to manage the vKVM service, define the number of concurrent sessions, configure network ports, and enforce encryption for secure remote management.
        #### Key Concepts
        - **Service Control:** The policy enables or disables the vKVM service on the endpoint.
        - **Session Management:** It allows administrators to set the maximum number of concurrent vKVM sessions, helping to manage resources and access.
        - **Security:** The policy can enforce encryption for all video information sent through the KVM, securing the remote console session.
        - **Tunneled KVM:** Supports enabling Tunneled vKVM, which allows secure KVM access through the Intersight platform without requiring direct network access to the server's management IP.
        - **Profile-Based Application:** It is attached to a Server Profile to apply the vKVM settings to the assigned physical server.

---

# Data Source: intersight_kvm_policy
The Virtual KVM Policy is a reusable policy used to configure the virtual Keyboard, Video, and Mouse (KVM) settings for a server.
#### Purpose
The purpose of this policy is to standardize and control remote console access to servers. It allows administrators to manage the vKVM service, define the number of concurrent sessions, configure network ports, and enforce encryption for secure remote management.
#### Key Concepts
- **Service Control:** The policy enables or disables the vKVM service on the endpoint.
- **Session Management:** It allows administrators to set the maximum number of concurrent vKVM sessions, helping to manage resources and access.
- **Security:** The policy can enforce encryption for all video information sent through the KVM, securing the remote console session.
- **Tunneled KVM:** Supports enabling Tunneled vKVM, which allows secure KVM access through the Intersight platform without requiring direct network access to the server's management IP.
- **Profile-Based Application:** It is attached to a Server Profile to apply the vKVM settings to the assigned physical server.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_kvm_policy.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `description`:(string) Description of the policy. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `enable_local_server_video`:(bool) If enabled, displays KVM session on any monitor attached to the server. 
* `enable_video_encryption`:(bool) If enabled, encrypts all video information sent through KVM. Please note that this can no longer be disabled for servers running versions 4.2 and above. 
* `enabled`:(bool) State of the vKVM service on the endpoint. 
* `maximum_sessions`:(int) The maximum number of concurrent KVM sessions allowed. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the concrete policy. 
* `remote_port`:(int) The port used for KVM communication. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `tunneled_kvm_enabled`:(bool) Enables Tunneled vKVM on the endpoint. Applicable only for Device Connectors that support Tunneled vKVM. 
 
