# EquipmentChassisIdentity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IoCardIdentityList** | Pointer to [**[]EquipmentIoCardIdentity**](equipment.IoCardIdentity.md) |  | [optional] 
**Chassis** | Pointer to [**EquipmentChassisRelationship**](equipment.Chassis.Relationship.md) |  | [optional] 

## Methods

### NewEquipmentChassisIdentity

`func NewEquipmentChassisIdentity() *EquipmentChassisIdentity`

NewEquipmentChassisIdentity instantiates a new EquipmentChassisIdentity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEquipmentChassisIdentityWithDefaults

`func NewEquipmentChassisIdentityWithDefaults() *EquipmentChassisIdentity`

NewEquipmentChassisIdentityWithDefaults instantiates a new EquipmentChassisIdentity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


