# VirtualizationVmwareDatacenterAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "virtualization.VmwareDatacenter"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "virtualization.VmwareDatacenter"]
**ClusterCount** | Pointer to **int64** | Count of all clusters associated with this DC. | [optional] 
**DatastoreCount** | Pointer to **int64** | Count of all datastores associated with this DC. | [optional] 
**HostCount** | Pointer to **int64** | Count of all hosts associated with this DC. | [optional] 
**InventoryPath** | Pointer to **string** | Inventory path of the DC. | [optional] 
**NetworkCount** | Pointer to **int64** | Count of all networks associated with this datacenter (DC). | [optional] 
**VmCount** | Pointer to **int64** | Count of all virtual machines (VMs) associated with this DC. | [optional] 
**HypervisorManager** | Pointer to [**VirtualizationVmwareVcenterRelationship**](VirtualizationVmwareVcenterRelationship.md) |  | [optional] 
**ParentFolder** | Pointer to [**VirtualizationVmwareFolderRelationship**](VirtualizationVmwareFolderRelationship.md) |  | [optional] 

## Methods

### NewVirtualizationVmwareDatacenterAllOf

`func NewVirtualizationVmwareDatacenterAllOf(classId string, objectType string, ) *VirtualizationVmwareDatacenterAllOf`

NewVirtualizationVmwareDatacenterAllOf instantiates a new VirtualizationVmwareDatacenterAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualizationVmwareDatacenterAllOfWithDefaults

`func NewVirtualizationVmwareDatacenterAllOfWithDefaults() *VirtualizationVmwareDatacenterAllOf`

NewVirtualizationVmwareDatacenterAllOfWithDefaults instantiates a new VirtualizationVmwareDatacenterAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *VirtualizationVmwareDatacenterAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *VirtualizationVmwareDatacenterAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *VirtualizationVmwareDatacenterAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *VirtualizationVmwareDatacenterAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VirtualizationVmwareDatacenterAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VirtualizationVmwareDatacenterAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetClusterCount

`func (o *VirtualizationVmwareDatacenterAllOf) GetClusterCount() int64`

GetClusterCount returns the ClusterCount field if non-nil, zero value otherwise.

### GetClusterCountOk

`func (o *VirtualizationVmwareDatacenterAllOf) GetClusterCountOk() (*int64, bool)`

GetClusterCountOk returns a tuple with the ClusterCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterCount

`func (o *VirtualizationVmwareDatacenterAllOf) SetClusterCount(v int64)`

SetClusterCount sets ClusterCount field to given value.

### HasClusterCount

`func (o *VirtualizationVmwareDatacenterAllOf) HasClusterCount() bool`

HasClusterCount returns a boolean if a field has been set.

### GetDatastoreCount

`func (o *VirtualizationVmwareDatacenterAllOf) GetDatastoreCount() int64`

GetDatastoreCount returns the DatastoreCount field if non-nil, zero value otherwise.

### GetDatastoreCountOk

`func (o *VirtualizationVmwareDatacenterAllOf) GetDatastoreCountOk() (*int64, bool)`

GetDatastoreCountOk returns a tuple with the DatastoreCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastoreCount

`func (o *VirtualizationVmwareDatacenterAllOf) SetDatastoreCount(v int64)`

SetDatastoreCount sets DatastoreCount field to given value.

### HasDatastoreCount

`func (o *VirtualizationVmwareDatacenterAllOf) HasDatastoreCount() bool`

HasDatastoreCount returns a boolean if a field has been set.

### GetHostCount

`func (o *VirtualizationVmwareDatacenterAllOf) GetHostCount() int64`

GetHostCount returns the HostCount field if non-nil, zero value otherwise.

### GetHostCountOk

`func (o *VirtualizationVmwareDatacenterAllOf) GetHostCountOk() (*int64, bool)`

GetHostCountOk returns a tuple with the HostCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostCount

`func (o *VirtualizationVmwareDatacenterAllOf) SetHostCount(v int64)`

SetHostCount sets HostCount field to given value.

### HasHostCount

`func (o *VirtualizationVmwareDatacenterAllOf) HasHostCount() bool`

HasHostCount returns a boolean if a field has been set.

### GetInventoryPath

`func (o *VirtualizationVmwareDatacenterAllOf) GetInventoryPath() string`

GetInventoryPath returns the InventoryPath field if non-nil, zero value otherwise.

### GetInventoryPathOk

`func (o *VirtualizationVmwareDatacenterAllOf) GetInventoryPathOk() (*string, bool)`

GetInventoryPathOk returns a tuple with the InventoryPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryPath

`func (o *VirtualizationVmwareDatacenterAllOf) SetInventoryPath(v string)`

SetInventoryPath sets InventoryPath field to given value.

### HasInventoryPath

`func (o *VirtualizationVmwareDatacenterAllOf) HasInventoryPath() bool`

HasInventoryPath returns a boolean if a field has been set.

### GetNetworkCount

`func (o *VirtualizationVmwareDatacenterAllOf) GetNetworkCount() int64`

GetNetworkCount returns the NetworkCount field if non-nil, zero value otherwise.

### GetNetworkCountOk

`func (o *VirtualizationVmwareDatacenterAllOf) GetNetworkCountOk() (*int64, bool)`

GetNetworkCountOk returns a tuple with the NetworkCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkCount

`func (o *VirtualizationVmwareDatacenterAllOf) SetNetworkCount(v int64)`

SetNetworkCount sets NetworkCount field to given value.

### HasNetworkCount

`func (o *VirtualizationVmwareDatacenterAllOf) HasNetworkCount() bool`

HasNetworkCount returns a boolean if a field has been set.

### GetVmCount

`func (o *VirtualizationVmwareDatacenterAllOf) GetVmCount() int64`

GetVmCount returns the VmCount field if non-nil, zero value otherwise.

### GetVmCountOk

`func (o *VirtualizationVmwareDatacenterAllOf) GetVmCountOk() (*int64, bool)`

GetVmCountOk returns a tuple with the VmCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmCount

`func (o *VirtualizationVmwareDatacenterAllOf) SetVmCount(v int64)`

SetVmCount sets VmCount field to given value.

### HasVmCount

`func (o *VirtualizationVmwareDatacenterAllOf) HasVmCount() bool`

HasVmCount returns a boolean if a field has been set.

### GetHypervisorManager

`func (o *VirtualizationVmwareDatacenterAllOf) GetHypervisorManager() VirtualizationVmwareVcenterRelationship`

GetHypervisorManager returns the HypervisorManager field if non-nil, zero value otherwise.

### GetHypervisorManagerOk

`func (o *VirtualizationVmwareDatacenterAllOf) GetHypervisorManagerOk() (*VirtualizationVmwareVcenterRelationship, bool)`

GetHypervisorManagerOk returns a tuple with the HypervisorManager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHypervisorManager

`func (o *VirtualizationVmwareDatacenterAllOf) SetHypervisorManager(v VirtualizationVmwareVcenterRelationship)`

SetHypervisorManager sets HypervisorManager field to given value.

### HasHypervisorManager

`func (o *VirtualizationVmwareDatacenterAllOf) HasHypervisorManager() bool`

HasHypervisorManager returns a boolean if a field has been set.

### GetParentFolder

`func (o *VirtualizationVmwareDatacenterAllOf) GetParentFolder() VirtualizationVmwareFolderRelationship`

GetParentFolder returns the ParentFolder field if non-nil, zero value otherwise.

### GetParentFolderOk

`func (o *VirtualizationVmwareDatacenterAllOf) GetParentFolderOk() (*VirtualizationVmwareFolderRelationship, bool)`

GetParentFolderOk returns a tuple with the ParentFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentFolder

`func (o *VirtualizationVmwareDatacenterAllOf) SetParentFolder(v VirtualizationVmwareFolderRelationship)`

SetParentFolder sets ParentFolder field to given value.

### HasParentFolder

`func (o *VirtualizationVmwareDatacenterAllOf) HasParentFolder() bool`

HasParentFolder returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


