---
subcategory: "vnic"
layout: "intersight"
page_title: "Intersight: intersight_vnic_vhba_template"
description: |-
        The vHBA template consists of the common vHBA configuration, which can be reused across multiple vHBAs. vHBAs can be created from the template using the Derive operation. Additionally, an existing vHBA can be attached to a template to use the configuration set in the template.
        To derive vHBAs from a vHBA template, you must use the asynchronous /v1/bulk/MoCloners bulk API.
        Deriving vHBAs from a vHBA Template
        URL: /v1/bulk/MoCloners
        Method: POST
        Headers:
        - Key: prefer
        Value: respond-async
        Body: >
        {
        Sources: [
        {
        Moid: xxxx,
        ObjectType: vnic.VhbaTemplate
        }
        ],
        Targets: [
        {
        Name: test-vh2,
        SanConnectivityPolicy: aaaaa,
        ObjectType: vnic.FcIf
        }
        ],
        WorkflowNameSuffix: Derive vHBA from a Template,
        Organization: {
        Moid: oooo,
        ObjectType: organization.Organization
        }
        }
        The API response includes the AsyncResult->bulk.Result MO reference. API clients must poll the bulk.Result MO to monitor the status of the operation.
        The bulk.Result object also includes a reference to a monitoring workflow, which is accessible from the 'Requests' tab in the UI.
        Managing Template Updates
        When the vHBA template is updated, the system initiates an automatic call to the /v1/bulk/MoMergers API to synchronize the template changes to all derived vHBA instances asynchronously. The status of the sync operation can be obtained by performing a query on vnic EthIf MO -
        $filter=SrcTemplate.Moid eq 'xxx'&$apply=groupby ( (TemplateSyncStatus), aggregate ($count as count))
        Override Option for vHBA Templates
        When enabled, this feature allows the configuration of the derived vHBA to override the configuration available in the template during vHBA create or update.
        You can query the list of overridable properties for a vHBA template from the Template Catalog. Use the following catalog API query:
        URL: /v1/capability/TemplateCatalogs
        Query: $filter= (Name eq ‘VhbaTemplate’)&$select=AllowedOverrideList
        Once a property is overridden on a derived vHBA, it will be added to the ‘OverriddenList’ property on the vnic FcIf MO.
        URL: /v1/vnic/FcIfs
        Query: $select=SrcTemplate, OverriddenList
        When override is disabled on the template, derived vHBA will have same configuration as the template.

---

# Data Source: intersight_vnic_vhba_template
The vHBA template consists of the common vHBA configuration, which can be reused across multiple vHBAs. vHBAs can be created from the template using the Derive operation. Additionally, an existing vHBA can be attached to a template to use the configuration set in the template.
To derive vHBAs from a vHBA template, you must use the asynchronous /v1/bulk/MoCloners bulk API.
Deriving vHBAs from a vHBA Template
URL: /v1/bulk/MoCloners
Method: POST
Headers:
  - Key: "prefer"
    Value: "respond-async"
Body: >
 {
    "Sources": [
      {
        "Moid": "xxxx",
        "ObjectType": "vnic.VhbaTemplate"
     }
    ],
    "Targets": [
      {
        "Name": "test-vh2",
        "SanConnectivityPolicy": "aaaaa",
        "ObjectType": "vnic.FcIf"
     }
    ],
    "WorkflowNameSuffix": "Derive vHBA from a Template",
    "Organization": {
        "Moid": "oooo",
        "ObjectType": "organization.Organization"
    }
}
The API response includes the "AsyncResult"->bulk.Result MO reference. API clients must poll the bulk.Result MO to monitor the status of the operation.
The bulk.Result object also includes a reference to a monitoring workflow, which is accessible from the 'Requests' tab in the UI.
Managing Template Updates
When the vHBA template is updated, the system initiates an automatic call to the /v1/bulk/MoMergers API to synchronize the template changes to all derived vHBA instances asynchronously. The status of the sync operation can be obtained by performing a query on vnic EthIf MO -
$filter=SrcTemplate.Moid eq 'xxx'&$apply=groupby ( (TemplateSyncStatus), aggregate ($count as count))
Override Option for vHBA Templates
When enabled, this feature allows the configuration of the derived vHBA to override the configuration available in the template during vHBA create or update.
You can query the list of overridable properties for a vHBA template from the Template Catalog. Use the following catalog API query:
URL: /v1/capability/TemplateCatalogs
Query: $filter= (Name eq ‘VhbaTemplate’)&$select=AllowedOverrideList
Once a property is overridden on a derived vHBA, it will be added to the ‘OverriddenList’ property on the vnic FcIf MO.
URL: /v1/vnic/FcIfs
Query: $select=SrcTemplate, OverriddenList
When override is disabled on the template, derived vHBA will have same configuration as the template.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_vnic_vhba_template.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `description`:(string) Description of the vHBA template. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `enable_override`:(bool) When enabled, the configuration of the derived instances may override the template configuration. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the virtual fibre channel interface. 
* `peer_vhba_name`:(string) Name of the peer vHBA which belongs to the peer FI. 
* `persistent_bindings`:(bool) Enables retention of LUN ID associations in memory until they are manually cleared. 
* `pin_group_name`:(string) Pingroup name associated to vfc for static pinning. SCP deploy will resolve pingroup name and fetches the correspoding uplink port/port channel to pin the vfc traffic. 
* `scp_usage_count`:(int) The count of the San Connectivity Policies using vHBA template. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `switch_id`:(string) The fabric port to which the vHBAs will be associated.* `None` - Fabric Id is not set to either A or B for the standalone case where the server is not connected to Fabric Interconnects. The value 'None' should be used.* `A` - Fabric A of the FI cluster.* `B` - Fabric B of the FI cluster. 
* `type`:(string) VHBA Type configuration for SAN Connectivity Policy. This configuration is supported only on Cisco VIC 14XX series and higher series of adapters.* `fc-initiator` - The default value set for vHBA Type Configuration. Fc-initiator specifies vHBA as a consumer of storage. Enables SCSI commands to transfer data and status information between host and target storage systems.* `fc-nvme-initiator` - Fc-nvme-initiator specifies vHBA as a consumer of storage. Enables NVMe-based message commands to transfer data and status information between host and target storage systems.* `fc-nvme-target` - Fc-nvme-target specifies vHBA as a provider of storage volumes to initiators. Enables NVMe-based message commands to transfer data and status information between host and target storage systems. Currently tech-preview, only enabled with an asynchronous driver.* `fc-target` - Fc-target specifies vHBA as a provider of storage volumes to initiators. Enables SCSI commands to transfer data and status information between host and target storage systems. fc-target is enabled only with an asynchronous driver. 
* `usage_count`:(int) The number of objects derived from a Template MO instance. 
 
