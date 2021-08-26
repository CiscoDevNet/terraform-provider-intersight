# VirtualizationVmwareFolder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "virtualization.VmwareFolder"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "virtualization.VmwareFolder"]
**Internal** | Pointer to **bool** | If a folder is internal, it will be set to true. | [optional] 
**InventoryPath** | Pointer to **string** | Inventory path to the folder. Example - /DC/myFolder. | [optional] 
**TypeofFolder** | Pointer to **string** | Determines the type of folder. e.g. vCenter folder, VM and Templete Folder, StorageFolder, NetworkFolder, Host and Cluster Folder. * &#x60;Unknown&#x60; - The type of the folder is unknown. It may not represent that the folder does not exist but indicates that something might be wrong. * &#x60;VMTemplateFolder&#x60; - The folder contains VMs and VM templates. * &#x60;StorageFolder&#x60; - The folder contains storage devices. * &#x60;HostClusterFolder&#x60; - The folder contains hosts and clusters. * &#x60;NetworkFolder&#x60; - The folder contains network items. * &#x60;VcenterFolder&#x60; - The folder created under a vCenter or vCenter folder. | [optional] [default to "Unknown"]
**Datacenter** | Pointer to [**VirtualizationVmwareDatacenterRelationship**](VirtualizationVmwareDatacenterRelationship.md) |  | [optional] 
**HypervisorManager** | Pointer to [**VirtualizationVmwareVcenterRelationship**](VirtualizationVmwareVcenterRelationship.md) |  | [optional] 
**VmwareFolder** | Pointer to [**VirtualizationVmwareFolderRelationship**](VirtualizationVmwareFolderRelationship.md) |  | [optional] 

## Methods

### NewVirtualizationVmwareFolder

`func NewVirtualizationVmwareFolder(classId string, objectType string, ) *VirtualizationVmwareFolder`

NewVirtualizationVmwareFolder instantiates a new VirtualizationVmwareFolder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualizationVmwareFolderWithDefaults

`func NewVirtualizationVmwareFolderWithDefaults() *VirtualizationVmwareFolder`

NewVirtualizationVmwareFolderWithDefaults instantiates a new VirtualizationVmwareFolder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *VirtualizationVmwareFolder) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *VirtualizationVmwareFolder) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *VirtualizationVmwareFolder) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *VirtualizationVmwareFolder) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VirtualizationVmwareFolder) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VirtualizationVmwareFolder) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetInternal

`func (o *VirtualizationVmwareFolder) GetInternal() bool`

GetInternal returns the Internal field if non-nil, zero value otherwise.

### GetInternalOk

`func (o *VirtualizationVmwareFolder) GetInternalOk() (*bool, bool)`

GetInternalOk returns a tuple with the Internal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternal

`func (o *VirtualizationVmwareFolder) SetInternal(v bool)`

SetInternal sets Internal field to given value.

### HasInternal

`func (o *VirtualizationVmwareFolder) HasInternal() bool`

HasInternal returns a boolean if a field has been set.

### GetInventoryPath

`func (o *VirtualizationVmwareFolder) GetInventoryPath() string`

GetInventoryPath returns the InventoryPath field if non-nil, zero value otherwise.

### GetInventoryPathOk

`func (o *VirtualizationVmwareFolder) GetInventoryPathOk() (*string, bool)`

GetInventoryPathOk returns a tuple with the InventoryPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryPath

`func (o *VirtualizationVmwareFolder) SetInventoryPath(v string)`

SetInventoryPath sets InventoryPath field to given value.

### HasInventoryPath

`func (o *VirtualizationVmwareFolder) HasInventoryPath() bool`

HasInventoryPath returns a boolean if a field has been set.

### GetTypeofFolder

`func (o *VirtualizationVmwareFolder) GetTypeofFolder() string`

GetTypeofFolder returns the TypeofFolder field if non-nil, zero value otherwise.

### GetTypeofFolderOk

`func (o *VirtualizationVmwareFolder) GetTypeofFolderOk() (*string, bool)`

GetTypeofFolderOk returns a tuple with the TypeofFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeofFolder

`func (o *VirtualizationVmwareFolder) SetTypeofFolder(v string)`

SetTypeofFolder sets TypeofFolder field to given value.

### HasTypeofFolder

`func (o *VirtualizationVmwareFolder) HasTypeofFolder() bool`

HasTypeofFolder returns a boolean if a field has been set.

### GetDatacenter

`func (o *VirtualizationVmwareFolder) GetDatacenter() VirtualizationVmwareDatacenterRelationship`

GetDatacenter returns the Datacenter field if non-nil, zero value otherwise.

### GetDatacenterOk

`func (o *VirtualizationVmwareFolder) GetDatacenterOk() (*VirtualizationVmwareDatacenterRelationship, bool)`

GetDatacenterOk returns a tuple with the Datacenter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatacenter

`func (o *VirtualizationVmwareFolder) SetDatacenter(v VirtualizationVmwareDatacenterRelationship)`

SetDatacenter sets Datacenter field to given value.

### HasDatacenter

`func (o *VirtualizationVmwareFolder) HasDatacenter() bool`

HasDatacenter returns a boolean if a field has been set.

### GetHypervisorManager

`func (o *VirtualizationVmwareFolder) GetHypervisorManager() VirtualizationVmwareVcenterRelationship`

GetHypervisorManager returns the HypervisorManager field if non-nil, zero value otherwise.

### GetHypervisorManagerOk

`func (o *VirtualizationVmwareFolder) GetHypervisorManagerOk() (*VirtualizationVmwareVcenterRelationship, bool)`

GetHypervisorManagerOk returns a tuple with the HypervisorManager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHypervisorManager

`func (o *VirtualizationVmwareFolder) SetHypervisorManager(v VirtualizationVmwareVcenterRelationship)`

SetHypervisorManager sets HypervisorManager field to given value.

### HasHypervisorManager

`func (o *VirtualizationVmwareFolder) HasHypervisorManager() bool`

HasHypervisorManager returns a boolean if a field has been set.

### GetVmwareFolder

`func (o *VirtualizationVmwareFolder) GetVmwareFolder() VirtualizationVmwareFolderRelationship`

GetVmwareFolder returns the VmwareFolder field if non-nil, zero value otherwise.

### GetVmwareFolderOk

`func (o *VirtualizationVmwareFolder) GetVmwareFolderOk() (*VirtualizationVmwareFolderRelationship, bool)`

GetVmwareFolderOk returns a tuple with the VmwareFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmwareFolder

`func (o *VirtualizationVmwareFolder) SetVmwareFolder(v VirtualizationVmwareFolderRelationship)`

SetVmwareFolder sets VmwareFolder field to given value.

### HasVmwareFolder

`func (o *VirtualizationVmwareFolder) HasVmwareFolder() bool`

HasVmwareFolder returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


