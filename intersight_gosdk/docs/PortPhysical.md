# PortPhysical

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**OperState** | Pointer to **string** | Operational state of this port (enabled/disabled). | [optional] [readonly] 
**OperStateQual** | Pointer to **string** | Reason for this port&#39;s Operational state. | [optional] [readonly] 
**PortId** | Pointer to **int64** | Switch physical port identifier. | [optional] [readonly] 
**Role** | Pointer to **string** | The role assigned to this port. | [optional] [readonly] 
**SlotId** | Pointer to **int64** | Switch expansion slot module identifier. | [optional] [readonly] 
**SwitchId** | Pointer to **string** | Switch Identifier that is local to a cluster. | [optional] [readonly] 

## Methods

### NewPortPhysical

`func NewPortPhysical(classId string, objectType string, ) *PortPhysical`

NewPortPhysical instantiates a new PortPhysical object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPortPhysicalWithDefaults

`func NewPortPhysicalWithDefaults() *PortPhysical`

NewPortPhysicalWithDefaults instantiates a new PortPhysical object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *PortPhysical) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *PortPhysical) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *PortPhysical) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *PortPhysical) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *PortPhysical) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *PortPhysical) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetOperState

`func (o *PortPhysical) GetOperState() string`

GetOperState returns the OperState field if non-nil, zero value otherwise.

### GetOperStateOk

`func (o *PortPhysical) GetOperStateOk() (*string, bool)`

GetOperStateOk returns a tuple with the OperState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperState

`func (o *PortPhysical) SetOperState(v string)`

SetOperState sets OperState field to given value.

### HasOperState

`func (o *PortPhysical) HasOperState() bool`

HasOperState returns a boolean if a field has been set.

### GetOperStateQual

`func (o *PortPhysical) GetOperStateQual() string`

GetOperStateQual returns the OperStateQual field if non-nil, zero value otherwise.

### GetOperStateQualOk

`func (o *PortPhysical) GetOperStateQualOk() (*string, bool)`

GetOperStateQualOk returns a tuple with the OperStateQual field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperStateQual

`func (o *PortPhysical) SetOperStateQual(v string)`

SetOperStateQual sets OperStateQual field to given value.

### HasOperStateQual

`func (o *PortPhysical) HasOperStateQual() bool`

HasOperStateQual returns a boolean if a field has been set.

### GetPortId

`func (o *PortPhysical) GetPortId() int64`

GetPortId returns the PortId field if non-nil, zero value otherwise.

### GetPortIdOk

`func (o *PortPhysical) GetPortIdOk() (*int64, bool)`

GetPortIdOk returns a tuple with the PortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortId

`func (o *PortPhysical) SetPortId(v int64)`

SetPortId sets PortId field to given value.

### HasPortId

`func (o *PortPhysical) HasPortId() bool`

HasPortId returns a boolean if a field has been set.

### GetRole

`func (o *PortPhysical) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *PortPhysical) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *PortPhysical) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *PortPhysical) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetSlotId

`func (o *PortPhysical) GetSlotId() int64`

GetSlotId returns the SlotId field if non-nil, zero value otherwise.

### GetSlotIdOk

`func (o *PortPhysical) GetSlotIdOk() (*int64, bool)`

GetSlotIdOk returns a tuple with the SlotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlotId

`func (o *PortPhysical) SetSlotId(v int64)`

SetSlotId sets SlotId field to given value.

### HasSlotId

`func (o *PortPhysical) HasSlotId() bool`

HasSlotId returns a boolean if a field has been set.

### GetSwitchId

`func (o *PortPhysical) GetSwitchId() string`

GetSwitchId returns the SwitchId field if non-nil, zero value otherwise.

### GetSwitchIdOk

`func (o *PortPhysical) GetSwitchIdOk() (*string, bool)`

GetSwitchIdOk returns a tuple with the SwitchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchId

`func (o *PortPhysical) SetSwitchId(v string)`

SetSwitchId sets SwitchId field to given value.

### HasSwitchId

`func (o *PortPhysical) HasSwitchId() bool`

HasSwitchId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


