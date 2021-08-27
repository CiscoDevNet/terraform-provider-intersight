# CondHclStatusJobAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "cond.HclStatusJob"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "cond.HclStatusJob"]
**ManagedObject** | Pointer to [**InventoryBaseRelationship**](InventoryBaseRelationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewCondHclStatusJobAllOf

`func NewCondHclStatusJobAllOf(classId string, objectType string, ) *CondHclStatusJobAllOf`

NewCondHclStatusJobAllOf instantiates a new CondHclStatusJobAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCondHclStatusJobAllOfWithDefaults

`func NewCondHclStatusJobAllOfWithDefaults() *CondHclStatusJobAllOf`

NewCondHclStatusJobAllOfWithDefaults instantiates a new CondHclStatusJobAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *CondHclStatusJobAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *CondHclStatusJobAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *CondHclStatusJobAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *CondHclStatusJobAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CondHclStatusJobAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CondHclStatusJobAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetManagedObject

`func (o *CondHclStatusJobAllOf) GetManagedObject() InventoryBaseRelationship`

GetManagedObject returns the ManagedObject field if non-nil, zero value otherwise.

### GetManagedObjectOk

`func (o *CondHclStatusJobAllOf) GetManagedObjectOk() (*InventoryBaseRelationship, bool)`

GetManagedObjectOk returns a tuple with the ManagedObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedObject

`func (o *CondHclStatusJobAllOf) SetManagedObject(v InventoryBaseRelationship)`

SetManagedObject sets ManagedObject field to given value.

### HasManagedObject

`func (o *CondHclStatusJobAllOf) HasManagedObject() bool`

HasManagedObject returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *CondHclStatusJobAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *CondHclStatusJobAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *CondHclStatusJobAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *CondHclStatusJobAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


