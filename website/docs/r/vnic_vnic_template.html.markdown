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

# Resource: intersight_vnic_vnic_template
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
The following arguments are supported:
* `account_moid`:(string)(ReadOnly) The Account ID for this managed object. 
* `ancestors`:(Array)(ReadOnly) An array of relationships to moBaseMo resources. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `cdn`:(HashMap) - Consistent Device Naming configuration for the virtual NIC. 
This complex property has following sub-properties:
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `nr_source`:(string) Source of the CDN. It can either be user specified or be the same as the vNIC name.* `vnic` - Source of the CDN is the same as the vNIC name.* `user` - Source of the CDN is specified by the user. 
  + `value`:(string) The CDN value entered in case of user defined mode. 
* `create_time`:(string)(ReadOnly) The time when this managed object was created. 
* `description`:(string) Description of the vNIC template. 
* `domain_group_moid`:(string)(ReadOnly) The DomainGroup ID for this managed object. 
* `enable_override`:(bool) When enabled, the configuration of the derived instances may override the template configuration. 
* `eth_adapter_policy`:(HashMap) - A reference to a vnicEthAdapterPolicy resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `eth_network_policy`:(HashMap) - A reference to a vnicEthNetworkPolicy resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `eth_qos_policy`:(HashMap) - A reference to a vnicEthQosPolicy resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `fabric_eth_network_control_policy`:(HashMap) - A reference to a fabricEthNetworkControlPolicy resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `fabric_eth_network_group_policy`:(Array) An array of relationships to fabricEthNetworkGroupPolicy resources. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `failover_enabled`:(bool) Enabling failover ensures that traffic from the vNIC automatically fails over to the secondary Fabric Interconnect, in case the specified Fabric Interconnect path goes down. Failover applies only to Cisco VICs that are connected to a Fabric Interconnect cluster. 
* `iscsi_boot_policy`:(HashMap) - A reference to a vnicIscsiBootPolicy resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `lcp_usage_count`:(int)(ReadOnly) The count of the Lan Connectivity Policies using vNIC template. 
* `mac_pool`:(HashMap) - A reference to a macpoolPool resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `mod_time`:(string)(ReadOnly) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the vNIC template. 
* `organization`:(HashMap) - A reference to a organizationOrganization resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `owners`:
                (Array of schema.TypeString) -(ReadOnly)
* `parent`:(HashMap) -(ReadOnly) A reference to a moBaseMo resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `peer_vnic_name`:(string) Name of the peer vNIC which belongs to the peer FI. 
* `permission_resources`:(Array)(ReadOnly) An array of relationships to moBaseMo resources. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `pin_group_name`:(string) Pingroup name associated to vNIC for static pinning. LCP deploy will resolve pingroup name and fetches the correspoding uplink port/port channel to pin the vNIC traffic. 
* `shared_scope`:(string)(ReadOnly) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `sriov_settings`:(HashMap) - Single Root Input Output Virtualization (SR-IOV) Settings that enable one physical ethernet socket to appear as multiple NICs to the hypervisor. 
This complex property has following sub-properties:
  + `comp_count_per_vf`:(int) Completion Queue resources per Virtual Function (VF). 
  + `enabled`:(bool) If enabled, sets Single Root Input Output Virtualization (SR-IOV) on this vNIC. 
  + `int_count_per_vf`:(int) Interrupt Count resources per Virtual Function (VF). 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `rx_count_per_vf`:(int) Receive Queue resources per Virtual Function (VF). 
  + `tx_count_per_vf`:(int) Transmit Queue resources per Virtual Function (VF). 
  + `vf_count`:(int) Number of Virtual Functions (VF) to be created for this vNIC. Valid values are 1 to 64 when SR-IOV is enabled. 
* `switch_id`:(string) The fabric port to which the vNICs will be associated.* `None` - Fabric Id is not set to either A or B for the standalone case where the server is not connected to Fabric Interconnects. The value 'None' should be used.* `A` - Fabric A of the FI cluster.* `B` - Fabric B of the FI cluster. 
* `tags`:(Array)
This complex property has following sub-properties:
  + `key`:(string) The string representation of a tag key. 
  + `value`:(string) The string representation of a tag value. 
* `template_actions`:(Array)
This complex property has following sub-properties:
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `params`:(Array)
This complex property has following sub-properties:
    + `name`:(string) The action parameter identifier. The supported values are SyncType and SyncTimer for the template sync action.* `None` - The default parameter that implies that no action parameter is required for the template action.* `SyncType` - The parameter that describes the type of sync action such as SyncOne or SyncFailed supported on any template or derived object.* `SyncTimer` - The parameter for the initial delay in seconds after which the sync action must be executed. The supported range is from 0 to 60 seconds.* `OverriddenList` - The parameter applicable in attach operation indicating the configurations that must override the template configurations. 
    + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
    + `value`:(string) The action parameter value is based on the action parameter type. Supported action parameters and their values are-a) Name - SyncType, Supported Values - SyncFailed, SyncOne.b) Name - SyncTimer, Supported Values - 0 to 60 seconds.c) Name - OverriddenList, Supported Values - Comma Separated list of overridable configurations. 
  + `type`:(string) The action type to be executed.* `Sync` - The action to merge values from the template to its derived objects.* `Deploy` - The action to execute deploy action on all the objects derived from the template that is mainly applicable for the various profile types.* `Detach` - The action to detach the current derived object from its attached template.* `Attach` - The action to attach the current object to the specified template. 
* `update_status`:(string)(ReadOnly) The template sync status with all derived objects.* `None` - The Enum value represents that the object is not attached to any template.* `OK` - The Enum value represents that the object values are in sync with attached template.* `Scheduled` - The Enum value represents that the object sync from attached template is scheduled from template.* `InProgress` - The Enum value represents that the object sync with the attached template is in progress.* `OutOfSync` - The Enum value represents that the object values are not in sync with attached template. 
* `usage_count`:(int)(ReadOnly) The number of objects derived from a Template MO instance. 
* `usnic_settings`:(HashMap) - User Space NIC Settings that enable low-latency and higher throughput by bypassing the kernel layer when sending/receiving packets. 
This complex property has following sub-properties:
  + `cos`:(int) Class of Service to be used for traffic on the usNIC. 
  + `nr_count`:(int) Number of usNIC interfaces to be created. When usNIC is enabled, the valid values are from 1 to 225. When usNIC is disabled, the default value is 0. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `usnic_adapter_policy`:(string) Ethernet Adapter policy to be associated with the usNICs. 
* `version_context`:(HashMap) -(ReadOnly) The versioning info for this managed object. 
This complex property has following sub-properties:
  + `interested_mos`:(Array)
This complex property has following sub-properties:
    + `moid`:(string) The Moid of the referenced REST resource. 
    + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
    + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
  + `marked_for_deletion`:(bool)(ReadOnly) The flag to indicate if snapshot is marked for deletion or not. If flag is set then snapshot will be removed after the successful deployment of the policy. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `ref_mo`:(HashMap) -(ReadOnly) A reference to the original Managed Object. 
This complex property has following sub-properties:
    + `moid`:(string) The Moid of the referenced REST resource. 
    + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
    + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
  + `timestamp`:(string)(ReadOnly) The time this versioned Managed Object was created. 
  + `nr_version`:(string)(ReadOnly) The version of the Managed Object, e.g. an incrementing number or a hash id. 
  + `version_type`:(string)(ReadOnly) Specifies type of version. Currently the only supported value is \ Configured\ that is used to keep track of snapshots of policies and profiles that are intendedto be configured to target endpoints.* `Modified` - Version created every time an object is modified.* `Configured` - Version created every time an object is configured to the service profile.* `Deployed` - Version created for objects related to a service profile when it is deployed. 
* `vmq_settings`:(HashMap) - Virtual Machine Queue Settings for the virtual interface that allow efficient transfer of network traffic to the guest OS. 
This complex property has following sub-properties:
  + `enabled`:(bool) Enables VMQ feature on the virtual interface. 
  + `multi_queue_support`:(bool) Enables Virtual Machine Multi-Queue feature on the virtual interface. VMMQ allows configuration of multiple I/O queues for a single VM and thus distributes traffic across multiple CPU cores in a VM. 
  + `num_interrupts`:(int) The number of interrupt resources to be allocated. Recommended value is the number of CPU threads or logical processors available in the server. 
  + `num_sub_vnics`:(int) The number of sub vNICs to be created. 
  + `num_vmqs`:(int) The number of hardware Virtual Machine Queues to be allocated. The number of VMQs per adapter must be one more than the maximum number of VM NICs. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `vmmq_adapter_policy`:(string) Ethernet Adapter policy to be associated with the Sub vNICs. The Transmit Queue and Receive Queue resource value of VMMQ adapter policy should be greater than or equal to the configured number of sub vNICs. 


## Import
`intersight_vnic_vnic_template` can be imported using the Moid of the object, e.g.
```
$ terraform import intersight_vnic_vnic_template.example 1234567890987654321abcde
``` 
