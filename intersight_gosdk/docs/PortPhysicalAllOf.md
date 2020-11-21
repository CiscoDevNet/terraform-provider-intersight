# PortPhysicalAllOf

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

### NewPortPhysicalAllOf

`func NewPortPhysicalAllOf(classId string, objectType string, ) *PortPhysicalAllOf`

NewPortPhysicalAllOf instantiates a new PortPhysicalAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPortPhysicalAllOfWithDefaults

`func NewPortPhysicalAllOfWithDefaults() *PortPhysicalAllOf`

NewPortPhysicalAllOfWithDefaults instantiates a new PortPhysicalAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *PortPhysicalAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *PortPhysicalAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *PortPhysicalAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *PortPhysicalAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *PortPhysicalAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *PortPhysicalAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetOperState

`func (o *PortPhysicalAllOf) GetOperState() string`

GetOperState returns the OperState field if non-nil, zero value otherwise.

### GetOperStateOk

`func (o *PortPhysicalAllOf) GetOperStateOk() (*string, bool)`

GetOperStateOk returns a tuple with the OperState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperState

`func (o *PortPhysicalAllOf) SetOperState(v string)`

SetOperState sets OperState field to given value.

### HasOperState

`func (o *PortPhysicalAllOf) HasOperState() bool`

HasOperState returns a boolean if a field has been set.

### GetOperStateQual

`func (o *PortPhysicalAllOf) GetOperStateQual() string`

GetOperStateQual returns the OperStateQual field if non-nil, zero value otherwise.

### GetOperStateQualOk

`func (o *PortPhysicalAllOf) GetOperStateQualOk() (*string, bool)`

GetOperStateQualOk returns a tuple with the OperStateQual field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperStateQual

`func (o *PortPhysicalAllOf) SetOperStateQual(v string)`

SetOperStateQual sets OperStateQual field to given value.

### HasOperStateQual

`func (o *PortPhysicalAllOf) HasOperStateQual() bool`

HasOperStateQual returns a boolean if a field has been set.

### GetPortId

`func (o *PortPhysicalAllOf) GetPortId() int64`

GetPortId returns the PortId field if non-nil, zero value otherwise.

### GetPortIdOk

`func (o *PortPhysicalAllOf) GetPortIdOk() (*int64, bool)`

GetPortIdOk returns a tuple with the PortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortId

`func (o *PortPhysicalAllOf) SetPortId(v int64)`

SetPortId sets PortId field to given value.

### HasPortId

`func (o *PortPhysicalAllOf) HasPortId() bool`

HasPortId returns a boolean if a field has been set.

### GetRole

`func (o *PortPhysicalAllOf) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *PortPhysicalAllOf) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *PortPhysicalAllOf) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *PortPhysicalAllOf) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetSlotId

`func (o *PortPhysicalAllOf) GetSlotId() int64`

GetSlotId returns the SlotId field if non-nil, zero value otherwise.

### GetSlotIdOk

`func (o *PortPhysicalAllOf) GetSlotIdOk() (*int64, bool)`

GetSlotIdOk returns a tuple with the SlotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlotId

`func (o *PortPhysicalAllOf) SetSlotId(v int64)`

SetSlotId sets SlotId field to given value.

### HasSlotId

`func (o *PortPhysicalAllOf) HasSlotId() bool`

HasSlotId returns a boolean if a field has been set.

### GetSwitchId

`func (o *PortPhysicalAllOf) GetSwitchId() string`

GetSwitchId returns the SwitchId field if non-nil, zero value otherwise.

### GetSwitchIdOk

`func (o *PortPhysicalAllOf) GetSwitchIdOk() (*string, bool)`

GetSwitchIdOk returns a tuple with the SwitchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchId

`func (o *PortPhysicalAllOf) SetSwitchId(v string)`

SetSwitchId sets SwitchId field to given value.

### HasSwitchId

`func (o *PortPhysicalAllOf) HasSwitchId() bool`

HasSwitchId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


