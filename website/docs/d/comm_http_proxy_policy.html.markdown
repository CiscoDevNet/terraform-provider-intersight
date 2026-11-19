---
subcategory: "comm"
layout: "intersight"
page_title: "Intersight: intersight_comm_http_proxy_policy"
description: |-
        HttpProxyPolicies define the HTTP proxy configuration used by HyperFlex installation workflows and HyperFlex storage controller VMs. They provide a reusable, attachable policy that supplies proxy settings in environments where servers (including CIMC) and controller VMs must reach external services through an HTTP proxy.
        #### Purpose
        Enable consistent, profile-driven application of HTTP proxy settings for HyperFlex deployments across both FI-managed and standalone environments, supporting installation and ongoing controller-VM connectivity requirements.
        #### Key Concepts
        - **Proxy configuration intent**: Captures the desired HTTP proxy settings (via `comm.AbstractHttpProxyPolicy`) to be consumed by HyperFlex installation and controller VMs.
        - **Profile attachability and reuse**: Designed to be attached to `hyperflex.ClusterProfile`, and related via `clusterProfiles` for visibility into consumers.
        - **Multi-platform applicability**: Tagged for both FI and standalone management platforms (`managementplatform.fi`, `managementplatform.standalone`).
        - **Governed lifecycle**: Supports READ/CREATE/UPDATE/DELETE with privileges scoped to HyperFlex cluster profile management roles.
        - **Non-destructive detachment**: `clusterProfiles` uses `ondelete: unset`, allowing cluster profiles to persist even if a proxy policy is removed or replaced.

---

# Data Source: intersight_comm_http_proxy_policy
HttpProxyPolicies define the HTTP proxy configuration used by HyperFlex installation workflows and HyperFlex storage controller VMs. They provide a reusable, attachable policy that supplies proxy settings in environments where servers (including CIMC) and controller VMs must reach external services through an HTTP proxy.
#### Purpose
Enable consistent, profile-driven application of HTTP proxy settings for HyperFlex deployments across both FI-managed and standalone environments, supporting installation and ongoing controller-VM connectivity requirements.
#### Key Concepts
- **Proxy configuration intent**: Captures the desired HTTP proxy settings (via `comm.AbstractHttpProxyPolicy`) to be consumed by HyperFlex installation and controller VMs.
- **Profile attachability and reuse**: Designed to be attached to `hyperflex.ClusterProfile`, and related via `clusterProfiles` for visibility into consumers.
- **Multi-platform applicability**: Tagged for both FI and standalone management platforms (`managementplatform.fi`, `managementplatform.standalone`).
- **Governed lifecycle**: Supports READ/CREATE/UPDATE/DELETE with privileges scoped to HyperFlex cluster profile management roles.
- **Non-destructive detachment**: `clusterProfiles` uses `ondelete: unset`, allowing cluster profiles to persist even if a proxy policy is removed or replaced.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_comm_http_proxy_policy.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `description`:(string) Description of the policy. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `hostname`:(string) HTTP Proxy server FQDN or IP. 
* `is_password_set`:(bool) Indicates whether the value of the 'password' property has been set. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the concrete policy. 
* `password`:(string) The password for the HTTP Proxy. 
* `port`:(int) The HTTP Proxy port number.The port number of the HTTP proxy must be between 1 and 65535, inclusive. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `username`:(string) The username for the HTTP Proxy. 
 
