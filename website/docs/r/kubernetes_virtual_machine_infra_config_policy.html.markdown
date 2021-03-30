---
subcategory: "kubernetes"
layout: "intersight"
page_title: "Intersight: intersight_kubernetes_virtual_machine_infra_config_policy"
description: |-
  A policy specifying compute, storage and network infrastructure configuration for a Virtual Machine.
---

# Resource: intersight_kubernetes_virtual_machine_infra_config_policy
A policy specifying compute, storage and network infrastructure configuration for a Virtual Machine.
## Argument Reference
The following arguments are supported:
* `description`:(string) Description of the policy. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the concrete policy. 
* `organization`:(HashMap) - A reference to a organizationOrganization resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string)(Computed) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string)(Computed) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `profiles`:(Array) An array of relationships to kubernetesVirtualMachineInfrastructureProvider resources. 
This complex property has following sub-properties:
  + `moid`:(string)(Computed) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string)(Computed) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `tags`:(Array)
This complex property has following sub-properties:
  + `key`:(string) The string representation of a tag key. 
  + `value`:(string) The string representation of a tag value. 
* `target`:(HashMap) - A reference to a assetDeviceRegistration resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string)(Computed) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string)(Computed) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `vm_config`:(HashMap) - Virtual machine infrastucture provider allocation properties. 
This complex property has following sub-properties:
  + `additional_properties`:(JSON) - Additional Properties as per object type, can be added as JSON using `jsonencode()`. Allowed Types are: [kubernetes.EsxiVirtualMachineInfraConfig](#kubernetesEsxiVirtualMachineInfraConfig)
[kubernetes.HyperFlexApVirtualMachineInfraConfig](#kubernetesHyperFlexApVirtualMachineInfraConfig)
  + `interfaces`:
                (Array of schema.TypeString) -
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. 


## Import
`intersight_kubernetes_virtual_machine_infra_config_policy` can be imported using the Moid of the object, e.g.
```
$ terraform import intersight_kubernetes_virtual_machine_infra_config_policy.example 1234567890987654321abcde
```
## Allowed Types in `AdditionalProperties`
 
### [kubernetes.EsxiVirtualMachineInfraConfig](#argument-reference)
Infrastructure provider allocation configuration for ESXi virtual machine Kubernetes nodes.
* `cluster`:(string) Name of the vSphere cluster on which the virtual machines are created. 
* `datastore`:(string) Name of the datasore on which the virtual machine disks are created. 
* `is_passphrase_set`:(bool)(Computed) Indicates whether the value of the 'passphrase' property has been set. 
* `passphrase`:(string) Passphrase for the vcenter user. 
* `resource_pool`:(string) Name of the vSphere resource pool on which the virtual machines are created. 

### [kubernetes.HyperFlexApVirtualMachineInfraConfig](#argument-reference)
Infrastructure provider allocation configuration for HyperFlex Application platform virtual machine Kubernetes nodes.
  