---
subcategory: "hyperflex"
layout: "intersight"
page_title: "Intersight: intersight_hyperflex_proxy_setting_policy"
description: |-
        ProxySettingPolicies define the HTTP proxy configuration used by the HyperFlex installation workflow and by HyperFlex storage controller VMs. This policy is intended for environments where direct internet access from servers (including CIMC) and controller VMs is restricted and must traverse an HTTP proxy.
        #### Purpose
        Provide a reusable, attachable policy that supplies required proxy settings to HyperFlex cluster profile operations—especially installation, expansion, and software redeploy—so HyperFlex components can reach external services through an approved proxy.
        #### Key Concepts
        - **Installation and runtime consumption**: Used both during HyperFlex installation and by storage controller VMs after deployment to enable proxied connectivity.
        - **Required for proxied environments**: Specifically targets environments where server/CIMC/controller VM internet access is secured via an HTTP proxy.
        - **Multi-platform applicability**: Tagged for both FI-managed and standalone management modes (`managementplatform.fi` and `managementplatform.standalone`).
        - **Profile attachability**: Marked as attachable to a HyperFlex cluster profile, aligning it with profile-driven configuration.
        - **Lifecycle operations support**: Tags indicate it is relevant for expansion and software redeploy workflows (`config.expansion`, `config.swredeploy`).
        - **Core proxy endpoint configuration**: `hostname` is a required, read-write setting representing the proxy server FQDN or IP address.

---

# Data Source: intersight_hyperflex_proxy_setting_policy
ProxySettingPolicies define the HTTP proxy configuration used by the HyperFlex installation workflow and by HyperFlex storage controller VMs. This policy is intended for environments where direct internet access from servers (including CIMC) and controller VMs is restricted and must traverse an HTTP proxy.
#### Purpose
Provide a reusable, attachable policy that supplies required proxy settings to HyperFlex cluster profile operations—especially installation, expansion, and software redeploy—so HyperFlex components can reach external services through an approved proxy.
#### Key Concepts
- **Installation and runtime consumption**: Used both during HyperFlex installation and by storage controller VMs after deployment to enable proxied connectivity.
- **Required for proxied environments**: Specifically targets environments where server/CIMC/controller VM internet access is secured via an HTTP proxy.
- **Multi-platform applicability**: Tagged for both FI-managed and standalone management modes (`managementplatform.fi` and `managementplatform.standalone`).
- **Profile attachability**: Marked as attachable to a HyperFlex cluster profile, aligning it with profile-driven configuration.
- **Lifecycle operations support**: Tags indicate it is relevant for expansion and software redeploy workflows (`config.expansion`, `config.swredeploy`).
- **Core proxy endpoint configuration**: `hostname` is a required, read-write setting representing the proxy server FQDN or IP address.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_hyperflex_proxy_setting_policy.<custom_name>.results[i].<propertyname>`.
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
 
