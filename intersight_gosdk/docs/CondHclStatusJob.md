# CondHclStatusJob

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ManagedObject** | Pointer to [**InventoryBaseRelationship**](inventory.Base.Relationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 

## Methods

### NewCondHclStatusJob

`func NewCondHclStatusJob() *CondHclStatusJob`

NewCondHclStatusJob instantiates a new CondHclStatusJob object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCondHclStatusJobWithDefaults

`func NewCondHclStatusJobWithDefaults() *CondHclStatusJob`

NewCondHclStatusJobWithDefaults instantiates a new CondHclStatusJob object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetManagedObject

`func (o *CondHclStatusJob) GetManagedObject() InventoryBaseRelationship`

GetManagedObject returns the ManagedObject field if non-nil, zero value otherwise.

### GetManagedObjectOk

`func (o *CondHclStatusJob) GetManagedObjectOk() (*InventoryBaseRelationship, bool)`

GetManagedObjectOk returns a tuple with the ManagedObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedObject

`func (o *CondHclStatusJob) SetManagedObject(v InventoryBaseRelationship)`

SetManagedObject sets ManagedObject field to given value.

### HasManagedObject

`func (o *CondHclStatusJob) HasManagedObject() bool`

HasManagedObject returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *CondHclStatusJob) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *CondHclStatusJob) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *CondHclStatusJob) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *CondHclStatusJob) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


