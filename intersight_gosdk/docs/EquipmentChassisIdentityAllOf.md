# EquipmentChassisIdentityAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "equipment.ChassisIdentity"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "equipment.ChassisIdentity"]
**IoCardIdentityList** | Pointer to [**[]EquipmentIoCardIdentity**](EquipmentIoCardIdentity.md) |  | [optional] 
**Chassis** | Pointer to [**EquipmentChassisRelationship**](EquipmentChassisRelationship.md) |  | [optional] 

## Methods

### NewEquipmentChassisIdentityAllOf

`func NewEquipmentChassisIdentityAllOf(classId string, objectType string, ) *EquipmentChassisIdentityAllOf`

NewEquipmentChassisIdentityAllOf instantiates a new EquipmentChassisIdentityAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEquipmentChassisIdentityAllOfWithDefaults

`func NewEquipmentChassisIdentityAllOfWithDefaults() *EquipmentChassisIdentityAllOf`

NewEquipmentChassisIdentityAllOfWithDefaults instantiates a new EquipmentChassisIdentityAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *EquipmentChassisIdentityAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *EquipmentChassisIdentityAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *EquipmentChassisIdentityAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *EquipmentChassisIdentityAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *EquipmentChassisIdentityAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *EquipmentChassisIdentityAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetIoCardIdentityList

`func (o *EquipmentChassisIdentityAllOf) GetIoCardIdentityList() []EquipmentIoCardIdentity`

GetIoCardIdentityList returns the IoCardIdentityList field if non-nil, zero value otherwise.

### GetIoCardIdentityListOk

`func (o *EquipmentChassisIdentityAllOf) GetIoCardIdentityListOk() (*[]EquipmentIoCardIdentity, bool)`

GetIoCardIdentityListOk returns a tuple with the IoCardIdentityList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIoCardIdentityList

`func (o *EquipmentChassisIdentityAllOf) SetIoCardIdentityList(v []EquipmentIoCardIdentity)`

SetIoCardIdentityList sets IoCardIdentityList field to given value.

### HasIoCardIdentityList

`func (o *EquipmentChassisIdentityAllOf) HasIoCardIdentityList() bool`

HasIoCardIdentityList returns a boolean if a field has been set.

### SetIoCardIdentityListNil

`func (o *EquipmentChassisIdentityAllOf) SetIoCardIdentityListNil(b bool)`

 SetIoCardIdentityListNil sets the value for IoCardIdentityList to be an explicit nil

### UnsetIoCardIdentityList
`func (o *EquipmentChassisIdentityAllOf) UnsetIoCardIdentityList()`

UnsetIoCardIdentityList ensures that no value is present for IoCardIdentityList, not even an explicit nil
### GetChassis

`func (o *EquipmentChassisIdentityAllOf) GetChassis() EquipmentChassisRelationship`

GetChassis returns the Chassis field if non-nil, zero value otherwise.

### GetChassisOk

`func (o *EquipmentChassisIdentityAllOf) GetChassisOk() (*EquipmentChassisRelationship, bool)`

GetChassisOk returns a tuple with the Chassis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChassis

`func (o *EquipmentChassisIdentityAllOf) SetChassis(v EquipmentChassisRelationship)`

SetChassis sets Chassis field to given value.

### HasChassis

`func (o *EquipmentChassisIdentityAllOf) HasChassis() bool`

HasChassis returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


