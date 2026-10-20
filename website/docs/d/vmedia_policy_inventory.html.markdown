---
subcategory: "vmedia"
layout: "intersight"
page_title: "Intersight: intersight_vmedia_policy_inventory"
description: |-
        The Virtual Media Policy is a reusable policy that configures the virtual media service on a server's management controller, allowing for the remote mounting of software images.
        #### Purpose
        The purpose of this policy is to enable and standardize the use of virtual media, which is essential for tasks like remote operating system installation, running diagnostic tools, or booting from recovery images. The policy defines how the virtual media service operates and specifies the remote locations of the images to be mounted.
        #### Key Concepts
        - **Remote Image Mounting:** The core function is to create mappings to remote image files (ISO or IMG), making them appear as local CD/DVD or HDD drives to the server.
        - **Multiple Protocol Support:** It supports mounting images from remote servers using various protocols, including NFS, CIFS, HTTP, and HTTPS.
        - **Service Control:** The policy allows administrators to enable or disable the virtual media service and to enforce encryption for all virtual media communications.
        - **Profile-Based Application:** It is attached to a Server Profile to apply a consistent virtual media configuration and set of mappings to multiple servers.
        - **Authentication:** For protocols like CIFS, the policy securely stores the username and password required to access the remote file share.

---

# Data Source: intersight_vmedia_policy_inventory
The Virtual Media Policy is a reusable policy that configures the virtual media service on a server's management controller, allowing for the remote mounting of software images.
#### Purpose
The purpose of this policy is to enable and standardize the use of virtual media, which is essential for tasks like remote operating system installation, running diagnostic tools, or booting from recovery images. The policy defines how the virtual media service operates and specifies the remote locations of the images to be mounted.
#### Key Concepts
- **Remote Image Mounting:** The core function is to create mappings to remote image files (ISO or IMG), making them appear as local CD/DVD or HDD drives to the server.
- **Multiple Protocol Support:** It supports mounting images from remote servers using various protocols, including NFS, CIFS, HTTP, and HTTPS.
- **Service Control:** The policy allows administrators to enable or disable the virtual media service and to enforce encryption for all virtual media communications.
- **Profile-Based Application:** It is attached to a Server Profile to apply a consistent virtual media configuration and set of mappings to multiple servers.
- **Authentication:** For protocols like CIFS, the policy securely stores the username and password required to access the remote file share.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_vmedia_policy_inventory.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `description`:(string) Description of the policy. 
* `device_mo_id`:(string) Device ID of the entity from where inventory is reported. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `enabled`:(bool) State of the Virtual Media service on the endpoint. 
* `encryption`:(bool) If enabled, allows encryption of all Virtual Media communications. Please note that this can no longer be disabled for servers running versions 4.2 and above. 
* `low_power_usb`:(bool) If enabled, the virtual drives appear on the boot selection menu after mapping the image and rebooting the host. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the inventoried policy object. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
