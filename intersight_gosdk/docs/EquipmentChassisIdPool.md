# EquipmentChassisIdPool

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "equipment.ChassisIdPool"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "equipment.ChassisIdPool"]
**PreferredIds** | Pointer to **[]int64** |  | [optional] 
**DeviceRegistration** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewEquipmentChassisIdPool

`func NewEquipmentChassisIdPool(classId string, objectType string, ) *EquipmentChassisIdPool`

NewEquipmentChassisIdPool instantiates a new EquipmentChassisIdPool object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEquipmentChassisIdPoolWithDefaults

`func NewEquipmentChassisIdPoolWithDefaults() *EquipmentChassisIdPool`

NewEquipmentChassisIdPoolWithDefaults instantiates a new EquipmentChassisIdPool object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *EquipmentChassisIdPool) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *EquipmentChassisIdPool) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *EquipmentChassisIdPool) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *EquipmentChassisIdPool) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *EquipmentChassisIdPool) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *EquipmentChassisIdPool) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetPreferredIds

`func (o *EquipmentChassisIdPool) GetPreferredIds() []int64`

GetPreferredIds returns the PreferredIds field if non-nil, zero value otherwise.

### GetPreferredIdsOk

`func (o *EquipmentChassisIdPool) GetPreferredIdsOk() (*[]int64, bool)`

GetPreferredIdsOk returns a tuple with the PreferredIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredIds

`func (o *EquipmentChassisIdPool) SetPreferredIds(v []int64)`

SetPreferredIds sets PreferredIds field to given value.

### HasPreferredIds

`func (o *EquipmentChassisIdPool) HasPreferredIds() bool`

HasPreferredIds returns a boolean if a field has been set.

### SetPreferredIdsNil

`func (o *EquipmentChassisIdPool) SetPreferredIdsNil(b bool)`

 SetPreferredIdsNil sets the value for PreferredIds to be an explicit nil

### UnsetPreferredIds
`func (o *EquipmentChassisIdPool) UnsetPreferredIds()`

UnsetPreferredIds ensures that no value is present for PreferredIds, not even an explicit nil
### GetDeviceRegistration

`func (o *EquipmentChassisIdPool) GetDeviceRegistration() AssetDeviceRegistrationRelationship`

GetDeviceRegistration returns the DeviceRegistration field if non-nil, zero value otherwise.

### GetDeviceRegistrationOk

`func (o *EquipmentChassisIdPool) GetDeviceRegistrationOk() (*AssetDeviceRegistrationRelationship, bool)`

GetDeviceRegistrationOk returns a tuple with the DeviceRegistration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceRegistration

`func (o *EquipmentChassisIdPool) SetDeviceRegistration(v AssetDeviceRegistrationRelationship)`

SetDeviceRegistration sets DeviceRegistration field to given value.

### HasDeviceRegistration

`func (o *EquipmentChassisIdPool) HasDeviceRegistration() bool`

HasDeviceRegistration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


