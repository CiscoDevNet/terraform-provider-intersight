---
subcategory: "ntp"
layout: "intersight"
page_title: "Intersight: intersight_ntp_ntp_server"
description: |-
        NtpServers represent configured Network Time Protocol (NTP) peers/servers on a managed network device. They provide read-only visibility into which NTP endpoints a device uses for time synchronization and how those peers are characterized (stratum, polling interval, VRF).
        #### Purpose
        Expose NTP configuration inventory for network devices so administrators can validate time-synchronization sources and troubleshoot time drift or connectivity issues.
        #### Key Concepts
        - **Time synchronization inventory:** Captures the configured NTP server/peer IP address and associated parameters.
        - **Server vs peer semantics:** The `type` distinguishes whether the configured endpoint is an NTP server or a peer.
        - **Quality and cadence indicators:** `stratum` and `poll` reflect the peer’s hierarchy level and polling interval behavior.
        - **Routing context:** `vrfName` indicates which VRF the device uses to reach the NTP endpoint.
        - **Device association:** Each NtpServer is tied to a specific registered device and inherits permissions from the network element context.

---

# Data Source: intersight_ntp_ntp_server
NtpServers represent configured Network Time Protocol (NTP) peers/servers on a managed network device. They provide read-only visibility into which NTP endpoints a device uses for time synchronization and how those peers are characterized (stratum, polling interval, VRF).
#### Purpose
Expose NTP configuration inventory for network devices so administrators can validate time-synchronization sources and troubleshoot time drift or connectivity issues.
#### Key Concepts
- **Time synchronization inventory:** Captures the configured NTP server/peer IP address and associated parameters.
- **Server vs peer semantics:** The `type` distinguishes whether the configured endpoint is an NTP server or a peer.
- **Quality and cadence indicators:** `stratum` and `poll` reflect the peer’s hierarchy level and polling interval behavior.
- **Routing context:** `vrfName` indicates which VRF the device uses to reach the NTP endpoint.
- **Device association:** Each NtpServer is tied to a specific registered device and inherits permissions from the network element context.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_ntp_ntp_server.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `poll`:(int) The default polling interval of the NTP server in seconds. 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
* `server_ip_address`:(string) The IP address of the NTP server. Supports both IPv4 or IPv6 address. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `stratum`:(int) The stratum level of the NTP server. 
* `type`:(string) It determines whether the IP address configured is server or peer.* `Server` - NTP configured is server type.* `Peer` - NTP configured is peer type. 
* `vrf_name`:(string) VRF name to be used by NTP Server. 
 
