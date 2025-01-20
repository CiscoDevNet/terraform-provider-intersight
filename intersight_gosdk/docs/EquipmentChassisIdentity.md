# EquipmentChassisIdentity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "equipment.ChassisIdentity"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "equipment.ChassisIdentity"]
**IoCardIdentityList** | Pointer to [**[]EquipmentIoCardIdentity**](EquipmentIoCardIdentity.md) |  | [optional] 
**ResetToDefault** | Pointer to **string** | Determines the default configuration deployment state for a chassis. * &#x60;NA&#x60; - Configuration deployment is not applicable for the chassis. * &#x60;Pending&#x60; - Configuration is not deployed to either A or B side IO card of the chassis. * &#x60;Applied&#x60; - Configuration is deployed to both A and B side IO cards of the chassis. * &#x60;PendingA&#x60; - Configuration deployment is pending in the A side IO card of the chassis. * &#x60;PendingB&#x60; - Configuration deployment is pending in the B side IO card of the chassis. | [optional] [readonly] [default to "NA"]
**Chassis** | Pointer to [**NullableEquipmentChassisRelationship**](EquipmentChassisRelationship.md) |  | [optional] 

## Methods

### NewEquipmentChassisIdentity

`func NewEquipmentChassisIdentity(classId string, objectType string, ) *EquipmentChassisIdentity`

NewEquipmentChassisIdentity instantiates a new EquipmentChassisIdentity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEquipmentChassisIdentityWithDefaults

`func NewEquipmentChassisIdentityWithDefaults() *EquipmentChassisIdentity`

NewEquipmentChassisIdentityWithDefaults instantiates a new EquipmentChassisIdentity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *EquipmentChassisIdentity) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *EquipmentChassisIdentity) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *EquipmentChassisIdentity) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *EquipmentChassisIdentity) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *EquipmentChassisIdentity) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *EquipmentChassisIdentity) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetIoCardIdentityList

`func (o *EquipmentChassisIdentity) GetIoCardIdentityList() []EquipmentIoCardIdentity`

GetIoCardIdentityList returns the IoCardIdentityList field if non-nil, zero value otherwise.

### GetIoCardIdentityListOk

`func (o *EquipmentChassisIdentity) GetIoCardIdentityListOk() (*[]EquipmentIoCardIdentity, bool)`

GetIoCardIdentityListOk returns a tuple with the IoCardIdentityList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIoCardIdentityList

`func (o *EquipmentChassisIdentity) SetIoCardIdentityList(v []EquipmentIoCardIdentity)`

SetIoCardIdentityList sets IoCardIdentityList field to given value.

### HasIoCardIdentityList

`func (o *EquipmentChassisIdentity) HasIoCardIdentityList() bool`

HasIoCardIdentityList returns a boolean if a field has been set.

### SetIoCardIdentityListNil

`func (o *EquipmentChassisIdentity) SetIoCardIdentityListNil(b bool)`

 SetIoCardIdentityListNil sets the value for IoCardIdentityList to be an explicit nil

### UnsetIoCardIdentityList
`func (o *EquipmentChassisIdentity) UnsetIoCardIdentityList()`

UnsetIoCardIdentityList ensures that no value is present for IoCardIdentityList, not even an explicit nil
### GetResetToDefault

`func (o *EquipmentChassisIdentity) GetResetToDefault() string`

GetResetToDefault returns the ResetToDefault field if non-nil, zero value otherwise.

### GetResetToDefaultOk

`func (o *EquipmentChassisIdentity) GetResetToDefaultOk() (*string, bool)`

GetResetToDefaultOk returns a tuple with the ResetToDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResetToDefault

`func (o *EquipmentChassisIdentity) SetResetToDefault(v string)`

SetResetToDefault sets ResetToDefault field to given value.

### HasResetToDefault

`func (o *EquipmentChassisIdentity) HasResetToDefault() bool`

HasResetToDefault returns a boolean if a field has been set.

### GetChassis

`func (o *EquipmentChassisIdentity) GetChassis() EquipmentChassisRelationship`

GetChassis returns the Chassis field if non-nil, zero value otherwise.

### GetChassisOk

`func (o *EquipmentChassisIdentity) GetChassisOk() (*EquipmentChassisRelationship, bool)`

GetChassisOk returns a tuple with the Chassis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChassis

`func (o *EquipmentChassisIdentity) SetChassis(v EquipmentChassisRelationship)`

SetChassis sets Chassis field to given value.

### HasChassis

`func (o *EquipmentChassisIdentity) HasChassis() bool`

HasChassis returns a boolean if a field has been set.

### SetChassisNil

`func (o *EquipmentChassisIdentity) SetChassisNil(b bool)`

 SetChassisNil sets the value for Chassis to be an explicit nil

### UnsetChassis
`func (o *EquipmentChassisIdentity) UnsetChassis()`

UnsetChassis ensures that no value is present for Chassis, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


