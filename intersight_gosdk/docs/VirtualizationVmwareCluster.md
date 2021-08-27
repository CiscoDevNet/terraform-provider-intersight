# VirtualizationVmwareCluster

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "virtualization.VmwareCluster"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "virtualization.VmwareCluster"]
**CpuOverCommitment** | Pointer to **int64** | CPU over commitment associated with this cluster. | [optional] 
**DatastoreCount** | Pointer to **int64** | Count of all datastores associated with this cluster. | [optional] 
**InventoryPath** | Pointer to **string** | Inventory path of the cluster. | [optional] 
**Datacenter** | Pointer to [**VirtualizationVmwareDatacenterRelationship**](VirtualizationVmwareDatacenterRelationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewVirtualizationVmwareCluster

`func NewVirtualizationVmwareCluster(classId string, objectType string, ) *VirtualizationVmwareCluster`

NewVirtualizationVmwareCluster instantiates a new VirtualizationVmwareCluster object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualizationVmwareClusterWithDefaults

`func NewVirtualizationVmwareClusterWithDefaults() *VirtualizationVmwareCluster`

NewVirtualizationVmwareClusterWithDefaults instantiates a new VirtualizationVmwareCluster object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *VirtualizationVmwareCluster) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *VirtualizationVmwareCluster) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *VirtualizationVmwareCluster) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *VirtualizationVmwareCluster) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VirtualizationVmwareCluster) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VirtualizationVmwareCluster) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCpuOverCommitment

`func (o *VirtualizationVmwareCluster) GetCpuOverCommitment() int64`

GetCpuOverCommitment returns the CpuOverCommitment field if non-nil, zero value otherwise.

### GetCpuOverCommitmentOk

`func (o *VirtualizationVmwareCluster) GetCpuOverCommitmentOk() (*int64, bool)`

GetCpuOverCommitmentOk returns a tuple with the CpuOverCommitment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuOverCommitment

`func (o *VirtualizationVmwareCluster) SetCpuOverCommitment(v int64)`

SetCpuOverCommitment sets CpuOverCommitment field to given value.

### HasCpuOverCommitment

`func (o *VirtualizationVmwareCluster) HasCpuOverCommitment() bool`

HasCpuOverCommitment returns a boolean if a field has been set.

### GetDatastoreCount

`func (o *VirtualizationVmwareCluster) GetDatastoreCount() int64`

GetDatastoreCount returns the DatastoreCount field if non-nil, zero value otherwise.

### GetDatastoreCountOk

`func (o *VirtualizationVmwareCluster) GetDatastoreCountOk() (*int64, bool)`

GetDatastoreCountOk returns a tuple with the DatastoreCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastoreCount

`func (o *VirtualizationVmwareCluster) SetDatastoreCount(v int64)`

SetDatastoreCount sets DatastoreCount field to given value.

### HasDatastoreCount

`func (o *VirtualizationVmwareCluster) HasDatastoreCount() bool`

HasDatastoreCount returns a boolean if a field has been set.

### GetInventoryPath

`func (o *VirtualizationVmwareCluster) GetInventoryPath() string`

GetInventoryPath returns the InventoryPath field if non-nil, zero value otherwise.

### GetInventoryPathOk

`func (o *VirtualizationVmwareCluster) GetInventoryPathOk() (*string, bool)`

GetInventoryPathOk returns a tuple with the InventoryPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryPath

`func (o *VirtualizationVmwareCluster) SetInventoryPath(v string)`

SetInventoryPath sets InventoryPath field to given value.

### HasInventoryPath

`func (o *VirtualizationVmwareCluster) HasInventoryPath() bool`

HasInventoryPath returns a boolean if a field has been set.

### GetDatacenter

`func (o *VirtualizationVmwareCluster) GetDatacenter() VirtualizationVmwareDatacenterRelationship`

GetDatacenter returns the Datacenter field if non-nil, zero value otherwise.

### GetDatacenterOk

`func (o *VirtualizationVmwareCluster) GetDatacenterOk() (*VirtualizationVmwareDatacenterRelationship, bool)`

GetDatacenterOk returns a tuple with the Datacenter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatacenter

`func (o *VirtualizationVmwareCluster) SetDatacenter(v VirtualizationVmwareDatacenterRelationship)`

SetDatacenter sets Datacenter field to given value.

### HasDatacenter

`func (o *VirtualizationVmwareCluster) HasDatacenter() bool`

HasDatacenter returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *VirtualizationVmwareCluster) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *VirtualizationVmwareCluster) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *VirtualizationVmwareCluster) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *VirtualizationVmwareCluster) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


