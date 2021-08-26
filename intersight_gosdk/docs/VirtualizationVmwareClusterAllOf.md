# VirtualizationVmwareClusterAllOf

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

### NewVirtualizationVmwareClusterAllOf

`func NewVirtualizationVmwareClusterAllOf(classId string, objectType string, ) *VirtualizationVmwareClusterAllOf`

NewVirtualizationVmwareClusterAllOf instantiates a new VirtualizationVmwareClusterAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualizationVmwareClusterAllOfWithDefaults

`func NewVirtualizationVmwareClusterAllOfWithDefaults() *VirtualizationVmwareClusterAllOf`

NewVirtualizationVmwareClusterAllOfWithDefaults instantiates a new VirtualizationVmwareClusterAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *VirtualizationVmwareClusterAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *VirtualizationVmwareClusterAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *VirtualizationVmwareClusterAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *VirtualizationVmwareClusterAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VirtualizationVmwareClusterAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VirtualizationVmwareClusterAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCpuOverCommitment

`func (o *VirtualizationVmwareClusterAllOf) GetCpuOverCommitment() int64`

GetCpuOverCommitment returns the CpuOverCommitment field if non-nil, zero value otherwise.

### GetCpuOverCommitmentOk

`func (o *VirtualizationVmwareClusterAllOf) GetCpuOverCommitmentOk() (*int64, bool)`

GetCpuOverCommitmentOk returns a tuple with the CpuOverCommitment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuOverCommitment

`func (o *VirtualizationVmwareClusterAllOf) SetCpuOverCommitment(v int64)`

SetCpuOverCommitment sets CpuOverCommitment field to given value.

### HasCpuOverCommitment

`func (o *VirtualizationVmwareClusterAllOf) HasCpuOverCommitment() bool`

HasCpuOverCommitment returns a boolean if a field has been set.

### GetDatastoreCount

`func (o *VirtualizationVmwareClusterAllOf) GetDatastoreCount() int64`

GetDatastoreCount returns the DatastoreCount field if non-nil, zero value otherwise.

### GetDatastoreCountOk

`func (o *VirtualizationVmwareClusterAllOf) GetDatastoreCountOk() (*int64, bool)`

GetDatastoreCountOk returns a tuple with the DatastoreCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastoreCount

`func (o *VirtualizationVmwareClusterAllOf) SetDatastoreCount(v int64)`

SetDatastoreCount sets DatastoreCount field to given value.

### HasDatastoreCount

`func (o *VirtualizationVmwareClusterAllOf) HasDatastoreCount() bool`

HasDatastoreCount returns a boolean if a field has been set.

### GetInventoryPath

`func (o *VirtualizationVmwareClusterAllOf) GetInventoryPath() string`

GetInventoryPath returns the InventoryPath field if non-nil, zero value otherwise.

### GetInventoryPathOk

`func (o *VirtualizationVmwareClusterAllOf) GetInventoryPathOk() (*string, bool)`

GetInventoryPathOk returns a tuple with the InventoryPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryPath

`func (o *VirtualizationVmwareClusterAllOf) SetInventoryPath(v string)`

SetInventoryPath sets InventoryPath field to given value.

### HasInventoryPath

`func (o *VirtualizationVmwareClusterAllOf) HasInventoryPath() bool`

HasInventoryPath returns a boolean if a field has been set.

### GetDatacenter

`func (o *VirtualizationVmwareClusterAllOf) GetDatacenter() VirtualizationVmwareDatacenterRelationship`

GetDatacenter returns the Datacenter field if non-nil, zero value otherwise.

### GetDatacenterOk

`func (o *VirtualizationVmwareClusterAllOf) GetDatacenterOk() (*VirtualizationVmwareDatacenterRelationship, bool)`

GetDatacenterOk returns a tuple with the Datacenter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatacenter

`func (o *VirtualizationVmwareClusterAllOf) SetDatacenter(v VirtualizationVmwareDatacenterRelationship)`

SetDatacenter sets Datacenter field to given value.

### HasDatacenter

`func (o *VirtualizationVmwareClusterAllOf) HasDatacenter() bool`

HasDatacenter returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *VirtualizationVmwareClusterAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *VirtualizationVmwareClusterAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *VirtualizationVmwareClusterAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *VirtualizationVmwareClusterAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


