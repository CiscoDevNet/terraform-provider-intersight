# EquipmentExpanderModuleIdentity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "equipment.ExpanderModuleIdentity"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "equipment.ExpanderModuleIdentity"]
**ChassisId** | Pointer to **int64** | Chassis Identifier of an expander module. | [optional] [readonly] 
**SlotId** | Pointer to **int64** | Chassis slot number of an expander module. | [optional] [readonly] 
**ExpanderModule** | Pointer to [**NullableEquipmentExpanderModuleRelationship**](EquipmentExpanderModuleRelationship.md) |  | [optional] 

## Methods

### NewEquipmentExpanderModuleIdentity

`func NewEquipmentExpanderModuleIdentity(classId string, objectType string, ) *EquipmentExpanderModuleIdentity`

NewEquipmentExpanderModuleIdentity instantiates a new EquipmentExpanderModuleIdentity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEquipmentExpanderModuleIdentityWithDefaults

`func NewEquipmentExpanderModuleIdentityWithDefaults() *EquipmentExpanderModuleIdentity`

NewEquipmentExpanderModuleIdentityWithDefaults instantiates a new EquipmentExpanderModuleIdentity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *EquipmentExpanderModuleIdentity) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *EquipmentExpanderModuleIdentity) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *EquipmentExpanderModuleIdentity) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *EquipmentExpanderModuleIdentity) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *EquipmentExpanderModuleIdentity) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *EquipmentExpanderModuleIdentity) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetChassisId

`func (o *EquipmentExpanderModuleIdentity) GetChassisId() int64`

GetChassisId returns the ChassisId field if non-nil, zero value otherwise.

### GetChassisIdOk

`func (o *EquipmentExpanderModuleIdentity) GetChassisIdOk() (*int64, bool)`

GetChassisIdOk returns a tuple with the ChassisId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChassisId

`func (o *EquipmentExpanderModuleIdentity) SetChassisId(v int64)`

SetChassisId sets ChassisId field to given value.

### HasChassisId

`func (o *EquipmentExpanderModuleIdentity) HasChassisId() bool`

HasChassisId returns a boolean if a field has been set.

### GetSlotId

`func (o *EquipmentExpanderModuleIdentity) GetSlotId() int64`

GetSlotId returns the SlotId field if non-nil, zero value otherwise.

### GetSlotIdOk

`func (o *EquipmentExpanderModuleIdentity) GetSlotIdOk() (*int64, bool)`

GetSlotIdOk returns a tuple with the SlotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlotId

`func (o *EquipmentExpanderModuleIdentity) SetSlotId(v int64)`

SetSlotId sets SlotId field to given value.

### HasSlotId

`func (o *EquipmentExpanderModuleIdentity) HasSlotId() bool`

HasSlotId returns a boolean if a field has been set.

### GetExpanderModule

`func (o *EquipmentExpanderModuleIdentity) GetExpanderModule() EquipmentExpanderModuleRelationship`

GetExpanderModule returns the ExpanderModule field if non-nil, zero value otherwise.

### GetExpanderModuleOk

`func (o *EquipmentExpanderModuleIdentity) GetExpanderModuleOk() (*EquipmentExpanderModuleRelationship, bool)`

GetExpanderModuleOk returns a tuple with the ExpanderModule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpanderModule

`func (o *EquipmentExpanderModuleIdentity) SetExpanderModule(v EquipmentExpanderModuleRelationship)`

SetExpanderModule sets ExpanderModule field to given value.

### HasExpanderModule

`func (o *EquipmentExpanderModuleIdentity) HasExpanderModule() bool`

HasExpanderModule returns a boolean if a field has been set.

### SetExpanderModuleNil

`func (o *EquipmentExpanderModuleIdentity) SetExpanderModuleNil(b bool)`

 SetExpanderModuleNil sets the value for ExpanderModule to be an explicit nil

### UnsetExpanderModule
`func (o *EquipmentExpanderModuleIdentity) UnsetExpanderModule()`

UnsetExpanderModule ensures that no value is present for ExpanderModule, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


