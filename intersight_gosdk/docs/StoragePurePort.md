# StoragePurePort

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "storage.PurePort"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "storage.PurePort"]
**Failover** | Pointer to **string** | Name of the port to which this port has failed over. | [optional] [readonly] 
**Nqn** | Pointer to **string** | The NVMe Qualified Name (NQN) associated with the host for ethernet port. | [optional] [readonly] 
**Portal** | Pointer to **string** | Ip address of iSCSI portal configured on the port. | [optional] [readonly] 
**Array** | Pointer to [**NullableStoragePureArrayRelationship**](StoragePureArrayRelationship.md) |  | [optional] 
**Controller** | Pointer to [**NullableStoragePureControllerRelationship**](StoragePureControllerRelationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**NullableAssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewStoragePurePort

`func NewStoragePurePort(classId string, objectType string, ) *StoragePurePort`

NewStoragePurePort instantiates a new StoragePurePort object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStoragePurePortWithDefaults

`func NewStoragePurePortWithDefaults() *StoragePurePort`

NewStoragePurePortWithDefaults instantiates a new StoragePurePort object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StoragePurePort) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StoragePurePort) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StoragePurePort) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StoragePurePort) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StoragePurePort) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StoragePurePort) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetFailover

`func (o *StoragePurePort) GetFailover() string`

GetFailover returns the Failover field if non-nil, zero value otherwise.

### GetFailoverOk

`func (o *StoragePurePort) GetFailoverOk() (*string, bool)`

GetFailoverOk returns a tuple with the Failover field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailover

`func (o *StoragePurePort) SetFailover(v string)`

SetFailover sets Failover field to given value.

### HasFailover

`func (o *StoragePurePort) HasFailover() bool`

HasFailover returns a boolean if a field has been set.

### GetNqn

`func (o *StoragePurePort) GetNqn() string`

GetNqn returns the Nqn field if non-nil, zero value otherwise.

### GetNqnOk

`func (o *StoragePurePort) GetNqnOk() (*string, bool)`

GetNqnOk returns a tuple with the Nqn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNqn

`func (o *StoragePurePort) SetNqn(v string)`

SetNqn sets Nqn field to given value.

### HasNqn

`func (o *StoragePurePort) HasNqn() bool`

HasNqn returns a boolean if a field has been set.

### GetPortal

`func (o *StoragePurePort) GetPortal() string`

GetPortal returns the Portal field if non-nil, zero value otherwise.

### GetPortalOk

`func (o *StoragePurePort) GetPortalOk() (*string, bool)`

GetPortalOk returns a tuple with the Portal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortal

`func (o *StoragePurePort) SetPortal(v string)`

SetPortal sets Portal field to given value.

### HasPortal

`func (o *StoragePurePort) HasPortal() bool`

HasPortal returns a boolean if a field has been set.

### GetArray

`func (o *StoragePurePort) GetArray() StoragePureArrayRelationship`

GetArray returns the Array field if non-nil, zero value otherwise.

### GetArrayOk

`func (o *StoragePurePort) GetArrayOk() (*StoragePureArrayRelationship, bool)`

GetArrayOk returns a tuple with the Array field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArray

`func (o *StoragePurePort) SetArray(v StoragePureArrayRelationship)`

SetArray sets Array field to given value.

### HasArray

`func (o *StoragePurePort) HasArray() bool`

HasArray returns a boolean if a field has been set.

### SetArrayNil

`func (o *StoragePurePort) SetArrayNil(b bool)`

 SetArrayNil sets the value for Array to be an explicit nil

### UnsetArray
`func (o *StoragePurePort) UnsetArray()`

UnsetArray ensures that no value is present for Array, not even an explicit nil
### GetController

`func (o *StoragePurePort) GetController() StoragePureControllerRelationship`

GetController returns the Controller field if non-nil, zero value otherwise.

### GetControllerOk

`func (o *StoragePurePort) GetControllerOk() (*StoragePureControllerRelationship, bool)`

GetControllerOk returns a tuple with the Controller field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetController

`func (o *StoragePurePort) SetController(v StoragePureControllerRelationship)`

SetController sets Controller field to given value.

### HasController

`func (o *StoragePurePort) HasController() bool`

HasController returns a boolean if a field has been set.

### SetControllerNil

`func (o *StoragePurePort) SetControllerNil(b bool)`

 SetControllerNil sets the value for Controller to be an explicit nil

### UnsetController
`func (o *StoragePurePort) UnsetController()`

UnsetController ensures that no value is present for Controller, not even an explicit nil
### GetRegisteredDevice

`func (o *StoragePurePort) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *StoragePurePort) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *StoragePurePort) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *StoragePurePort) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### SetRegisteredDeviceNil

`func (o *StoragePurePort) SetRegisteredDeviceNil(b bool)`

 SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil

### UnsetRegisteredDevice
`func (o *StoragePurePort) UnsetRegisteredDevice()`

UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


