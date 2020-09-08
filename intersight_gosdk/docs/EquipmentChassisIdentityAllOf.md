# EquipmentChassisIdentityAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IoCardIdentityList** | Pointer to [**[]EquipmentIoCardIdentity**](equipment.IoCardIdentity.md) |  | [optional] 
**Chassis** | Pointer to [**EquipmentChassisRelationship**](equipment.Chassis.Relationship.md) |  | [optional] 

## Methods

### NewEquipmentChassisIdentityAllOf

`func NewEquipmentChassisIdentityAllOf() *EquipmentChassisIdentityAllOf`

NewEquipmentChassisIdentityAllOf instantiates a new EquipmentChassisIdentityAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEquipmentChassisIdentityAllOfWithDefaults

`func NewEquipmentChassisIdentityAllOfWithDefaults() *EquipmentChassisIdentityAllOf`

NewEquipmentChassisIdentityAllOfWithDefaults instantiates a new EquipmentChassisIdentityAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

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


