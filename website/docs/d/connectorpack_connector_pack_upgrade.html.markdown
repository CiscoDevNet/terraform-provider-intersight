---
subcategory: "connectorpack"
layout: "intersight"
page_title: "Intersight: intersight_connectorpack_connector_pack_upgrade"
description: |-
        ConnectorPackUpgrades represent an operation request used to download and/or install connector packs on a target device (specifically UCS Director in this model). The object captures the requested operation type and links to both the target UCS Director instance and the workflow execution that performs the upgrade.
        #### Purpose
        Provide an API-driven mechanism to initiate and track connector pack download/install actions against a specific UCS Director target, enabling controlled upgrades via a workflow-backed operation.
        #### Key Concepts
        - **Operation request object**: Created to trigger a connector pack action; supports READ for visibility and DELETE for cleanup.
        - **Operation type control**: `connectorPackOpType` specifies what action should be performed on UCS Director (for example, download vs install, depending on the enum definition).
        - **Target binding (create-only)**: `ucsdInfo` identifies the UCS Director instance to which packs are pushed/installed and is `createonly`, preventing retargeting after creation.
        - **Lifecycle coupling to target**: `onpeerdelete: cascade` on `ucsdInfo` ties the operation record to the UCS Director object.
        - **Workflow-backed execution**: `workflow` references the runtime `workflow.WorkflowInfo` instance that carries out the upgrade operation, enabling tracking and troubleshooting of execution state.
        - **Role-governed access**: CREATE/DELETE is restricted to Account Administrator, while READ is available to common operational roles.

---

# Data Source: intersight_connectorpack_connector_pack_upgrade
ConnectorPackUpgrades represent an operation request used to download and/or install connector packs on a target device (specifically UCS Director in this model). The object captures the requested operation type and links to both the target UCS Director instance and the workflow execution that performs the upgrade.
#### Purpose
Provide an API-driven mechanism to initiate and track connector pack download/install actions against a specific UCS Director target, enabling controlled upgrades via a workflow-backed operation.
#### Key Concepts
- **Operation request object**: Created to trigger a connector pack action; supports READ for visibility and DELETE for cleanup.
- **Operation type control**: `connectorPackOpType` specifies what action should be performed on UCS Director (for example, download vs install, depending on the enum definition).
- **Target binding (create-only)**: `ucsdInfo` identifies the UCS Director instance to which packs are pushed/installed and is `createonly`, preventing retargeting after creation.
- **Lifecycle coupling to target**: `onpeerdelete: cascade` on `ucsdInfo` ties the operation record to the UCS Director object.
- **Workflow-backed execution**: `workflow` references the runtime `workflow.WorkflowInfo` instance that carries out the upgrade operation, enabling tracking and troubleshooting of execution state.
- **Role-governed access**: CREATE/DELETE is restricted to Account Administrator, while READ is available to common operational roles.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_connectorpack_connector_pack_upgrade.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `connector_pack_op_type`:(string) The type of operation to be performed on UCS Director.* `Install` - Installs the requisite connector packs on UCS Director.* `Push` - Pushes the requisite connector packs to UCS Director. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
