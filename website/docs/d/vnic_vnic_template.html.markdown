---
subcategory: "vnic"
layout: "intersight"
page_title: "Intersight: intersight_vnic_vnic_template"
description: |-
        The vNIC template consists of the common vNIC configuration, which can be reused across multiple vNICs. vNICs can be created from the template using the Derive operation. Additionally, an existing vNIC can be attached to a template to use the configuration set in the template.
        To derive vNICs from a vNIC template, you must use the asynchronous /v1/bulk/MoCloners bulk API.
        Deriving vNICs from a vNIC Template
        URL: /v1/bulk/MoCloners
        Method: POST
        Headers:
        - Key: prefer
        Value: respond-async
        Body: >
        {
        Sources: [
        {
        Moid: xxxxx,
        ObjectType: vnic.VnicTemplate
        }
        ],
        Targets: [
        {
        Name: test-vn2,
        MacAddressType: POOL,
        Placement: {
        AutoSlotId: true,
        AutoPciLink: true
        },
        LanConnectivityPolicy: aaaaa,
        ObjectType: vnic.EthIf
        }
        ],
        WorkflowNameSuffix: Derive vNIC from a Template,
        Organization: {
        Moid: ooooo,
        ObjectType: organization.Organization
        }
        }
        The API response includes the AsyncResult->bulk.Result MO reference. API clients must poll the bulk.Result MO to monitor the status of the operation.
        The bulk.Result object also includes a reference to a monitoring workflow, which is accessible from the 'Requests' tab in the UI.
        Managing Template Updates
        When the vNIC template is updated, the system initiates an automatic call to the /v1/bulk/MoMergers API to synchronize the template changes to all derived vNIC instances asynchronously. The status of the sync operation can be obtained by performing a query on vnic EthIf MO -
        $filter=SrcTemplate.Moid eq 'xxx'&$apply=groupby ( (TemplateSyncStatus), aggregate ($count as count))
        Override Option for vNIC Templates
        When enabled, this feature allows the configuration of the derived vNIC to override the configuration available in the template during vNIC create or update.
        You can query the list of overridable properties for a vNIC template from the Template Catalog. Use the following catalog API query:
        URL: /v1/capability/TemplateCatalogs
        Query: $filter= (Name eq ‘VnicTemplate’)&$select=AllowedOverrideList
        Once a property is overridden on a derived vNIC, it will be added to the ‘OverriddenList’ property on the vnic EthIf MO:
        URL: /v1/vnic/EthIfs
        Query: $select=SrcTemplate, OverriddenList
        When override is disabled on the template, derived vNIC will have same configuration as the template.

---

# Data Source: intersight_vnic_vnic_template
The vNIC template consists of the common vNIC configuration, which can be reused across multiple vNICs. vNICs can be created from the template using the Derive operation. Additionally, an existing vNIC can be attached to a template to use the configuration set in the template.
To derive vNICs from a vNIC template, you must use the asynchronous /v1/bulk/MoCloners bulk API.
Deriving vNICs from a vNIC Template
URL: /v1/bulk/MoCloners
Method: POST
Headers:
  - Key: "prefer"
    Value: "respond-async"
Body: >
 {
    "Sources": [
      {
        "Moid": "xxxxx",
        "ObjectType": "vnic.VnicTemplate"
     }
    ],
    "Targets": [
      {
        "Name": "test-vn2",
        "MacAddressType": "POOL",
        "Placement": {
            "AutoSlotId": true,
            "AutoPciLink": true
        },
        "LanConnectivityPolicy": "aaaaa",
        "ObjectType": "vnic.EthIf"
     }
    ],
    "WorkflowNameSuffix": "Derive vNIC from a Template",
    "Organization": {
        "Moid": "ooooo",
        "ObjectType": "organization.Organization"
    }
}
The API response includes the "AsyncResult"->bulk.Result MO reference. API clients must poll the bulk.Result MO to monitor the status of the operation.
The bulk.Result object also includes a reference to a monitoring workflow, which is accessible from the 'Requests' tab in the UI.
Managing Template Updates
When the vNIC template is updated, the system initiates an automatic call to the /v1/bulk/MoMergers API to synchronize the template changes to all derived vNIC instances asynchronously. The status of the sync operation can be obtained by performing a query on vnic EthIf MO -
$filter=SrcTemplate.Moid eq 'xxx'&$apply=groupby ( (TemplateSyncStatus), aggregate ($count as count))
Override Option for vNIC Templates
When enabled, this feature allows the configuration of the derived vNIC to override the configuration available in the template during vNIC create or update.
You can query the list of overridable properties for a vNIC template from the Template Catalog. Use the following catalog API query:
URL: /v1/capability/TemplateCatalogs
Query: $filter= (Name eq ‘VnicTemplate’)&$select=AllowedOverrideList
Once a property is overridden on a derived vNIC, it will be added to the ‘OverriddenList’ property on the vnic EthIf MO:
URL: /v1/vnic/EthIfs
Query: $select=SrcTemplate, OverriddenList
When override is disabled on the template, derived vNIC will have same configuration as the template.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_vnic_vnic_template.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `description`:(string) Description of the vNIC template. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `enable_override`:(bool) When enabled, the configuration of the derived instances may override the template configuration. 
* `failover_enabled`:(bool) Enabling failover ensures that traffic from the vNIC automatically fails over to the secondary Fabric Interconnect, in case the specified Fabric Interconnect path goes down. Failover applies only to Cisco VICs that are connected to a Fabric Interconnect cluster. 
* `lcp_usage_count`:(int) The count of the Lan Connectivity Policies using vNIC template. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the vNIC template. 
* `peer_vnic_name`:(string) Name of the peer vNIC which belongs to the peer FI. 
* `pin_group_name`:(string) Pingroup name associated to vNIC for static pinning. LCP deploy will resolve pingroup name and fetches the correspoding uplink port/port channel to pin the vNIC traffic. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `switch_id`:(string) The fabric port to which the vNICs will be associated.* `None` - Fabric Id is not set to either A or B for the standalone case where the server is not connected to Fabric Interconnects. The value 'None' should be used.* `A` - Fabric A of the FI cluster.* `B` - Fabric B of the FI cluster. 
* `usage_count`:(int) The number of objects derived from a Template MO instance. 
 
