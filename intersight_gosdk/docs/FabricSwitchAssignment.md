# FabricSwitchAssignment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "fabric.SwitchAssignment"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "fabric.SwitchAssignment"]
**Enabled** | Pointer to **bool** | Indicates if this assignment is enabled. | [optional] 
**SwitchSerial** | Pointer to **string** | Serial number of the switch. | [optional] 

## Methods

### NewFabricSwitchAssignment

`func NewFabricSwitchAssignment(classId string, objectType string, ) *FabricSwitchAssignment`

NewFabricSwitchAssignment instantiates a new FabricSwitchAssignment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFabricSwitchAssignmentWithDefaults

`func NewFabricSwitchAssignmentWithDefaults() *FabricSwitchAssignment`

NewFabricSwitchAssignmentWithDefaults instantiates a new FabricSwitchAssignment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *FabricSwitchAssignment) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *FabricSwitchAssignment) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *FabricSwitchAssignment) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *FabricSwitchAssignment) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *FabricSwitchAssignment) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *FabricSwitchAssignment) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetEnabled

`func (o *FabricSwitchAssignment) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *FabricSwitchAssignment) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *FabricSwitchAssignment) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *FabricSwitchAssignment) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetSwitchSerial

`func (o *FabricSwitchAssignment) GetSwitchSerial() string`

GetSwitchSerial returns the SwitchSerial field if non-nil, zero value otherwise.

### GetSwitchSerialOk

`func (o *FabricSwitchAssignment) GetSwitchSerialOk() (*string, bool)`

GetSwitchSerialOk returns a tuple with the SwitchSerial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchSerial

`func (o *FabricSwitchAssignment) SetSwitchSerial(v string)`

SetSwitchSerial sets SwitchSerial field to given value.

### HasSwitchSerial

`func (o *FabricSwitchAssignment) HasSwitchSerial() bool`

HasSwitchSerial returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


