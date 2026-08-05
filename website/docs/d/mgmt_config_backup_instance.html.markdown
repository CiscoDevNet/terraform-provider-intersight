---
subcategory: "mgmt"
layout: "intersight"
page_title: "Intersight: intersight_mgmt_config_backup_instance"
description: |-
        The ConfigBackupInstance object represents a stored backup artifact in the system repository. It is the consumable backup unit that can be protected, downloaded, and later used as a restore source.
        #### Purpose
        ConfigBackupInstance exists to provide a durable, managed representation of a backup file produced by a backup operation (or created from an imported file). It is the primary object customers interact with when securing and retrieving backup archives.
        #### Key Concepts
        - **Repository-Backed Artifact:** Represents a backup stored in the system, including its readiness and usability for restore operations.
        - **Controlled Download Workflow:** Supports generating short-lived download URLs and managing the actions required to prepare the archive for retrieval.
        - **Encryption Enablement:** Designed to support customer-controlled encryption flows prior to download, enabling stronger handling of backup confidentiality.
        - **Enable Retention:** Provides retention locking to prevent deletion and rollover, supporting compliance-oriented backup handling.
        - **Operational Status Signaling:** Exposes clear state transitions (e.g., ready, corrupted, encryption states) so customers can automate decisions around backup usage.

---

# Data Source: intersight_mgmt_config_backup_instance
The ConfigBackupInstance object represents a stored backup artifact in the system repository. It is the consumable backup unit that can be protected, downloaded, and later used as a restore source.
#### Purpose
ConfigBackupInstance exists to provide a durable, managed representation of a backup file produced by a backup operation (or created from an imported file). It is the primary object customers interact with when securing and retrieving backup archives.
#### Key Concepts
- **Repository-Backed Artifact:** Represents a backup stored in the system, including its readiness and usability for restore operations.
- **Controlled Download Workflow:** Supports generating short-lived download URLs and managing the actions required to prepare the archive for retrieval.
- **Encryption Enablement:** Designed to support customer-controlled encryption flows prior to download, enabling stronger handling of backup confidentiality.
- **Enable Retention:** Provides retention locking to prevent deletion and rollover, supporting compliance-oriented backup handling.
- **Operational Status Signaling:** Exposes clear state transitions (e.g., ready, corrupted, encryption states) so customers can automate decisions around backup usage.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_mgmt_config_backup_instance.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `admin_action`:(string) Action to be performed on the backup instance.* `None` - Nothing to be done. The default state of the backup instance when there is no action to be performed.* `EncryptWithPublicKey` - Encrypt the backup using the user given public key and make it available for download.* `EncryptWithPassword` - Encrypt the backup using the user provided password and make it available for download.* `GenerateDownloadUrl` - Generate a pre-signed URL that can be used to download a backup archive. 
* `backup_creation_time`:(string) The time at which the backup file was created. 
* `backup_name`:(string) Name of the backup for which the backup instance is created. 
* `backup_source`:(string) Indicates whether the backup was created locally or imported.* `Local` - A local Intersight location.* `Uploaded` - A local location where the backup file is uploaded as mgmt.ConfigBackupFile MO. Intersight creates the backup from the uploaded location when this location type is set.* `Remote` - A remote location hosted in the user's datacenter. 
* `checksum`:(string) SHA-256 checksum of the encrypted backup archive, returned as a hex-encoded string. Users can compare the checksum of a downloaded archive with this value to verify file integrity and confirm the archive was not corrupted or altered in transit. 
* `create_time`:(string) The time when this managed object was created. 
* `description`:(string) Description of the backup operation that created this instance. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `download_url`:(string) Pre-signed URL that can be used to download the backup archive. This URL is generated when the adminAction is set to GenerateDownloadUrl. The URL is valid for 5 minutes. 
* `is_user_password_set`:(bool) Indicates whether the value of the 'userPassword' property has been set. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `path`:(string) Path to the backup file in the Intersight repository. 
* `preserve_identities`:(bool) The flag set by the user during a configuration backup to preserve static or dynamic IDs assigned to an export item. 
* `retention_lock`:(bool) When set, ensures that the backup archive is protected from deletion and rollover operations. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `size`:(int) Size of the backup instance, represents size of the compressed backup file. 
* `source_account_moid`:(string) The moid of the account from which the backup was taken. 
* `source_name`:(string) The name of the account from which the backup was taken. 
* `status`:(string) Status of the backup instance.* `Unknown` - The status is not known. When the backup instances are created, they are in this state until the backup is ready.* `Ready` - Ready for use in an import operation.* `RemovalInProgress` - The backup is being removed.* `Corrupted` - The backup is corrupted and is not usable.* `EncryptionInProgress` - The backup is being encrypted using the user provided public key.* `EncryptionFailed` - The backup could not be encrypted using the user provided public key.* `Encrypted` - The backup has been encrypted using the user provided public key. 
* `system_admin_triggered`:(bool) Set when a system administrator initiated the backup operation. 
* `user_password`:(string) The password provided by the user to encrypt the backup before download. This password is not persisted in Intersight and this field will be reset after the backup archive is encrypted. The password encrypts the backup file before download. The user needs to provide the same password to decrypt when they upload the backup file to restore. 
* `user_public_key`:(string) The public key provided by the user to encrypt the backup before download. The key needs to be in PKCS#1 format encoded in PEM. 
 
