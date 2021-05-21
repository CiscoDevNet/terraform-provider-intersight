# VirtualizationVmwareVlanRange

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "virtualization.VmwareVlanRange"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "virtualization.VmwareVlanRange"]
**VlanRangeEnd** | Pointer to **int64** | End value of VLAN range for the trunk port. The valid range is from 0 to 4094. | [optional] 
**VlanRangeStart** | Pointer to **int64** | Start value of VLAN range for the trunk port. The valid range is from 0 to 4094. | [optional] 

## Methods

### NewVirtualizationVmwareVlanRange

`func NewVirtualizationVmwareVlanRange(classId string, objectType string, ) *VirtualizationVmwareVlanRange`

NewVirtualizationVmwareVlanRange instantiates a new VirtualizationVmwareVlanRange object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualizationVmwareVlanRangeWithDefaults

`func NewVirtualizationVmwareVlanRangeWithDefaults() *VirtualizationVmwareVlanRange`

NewVirtualizationVmwareVlanRangeWithDefaults instantiates a new VirtualizationVmwareVlanRange object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *VirtualizationVmwareVlanRange) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *VirtualizationVmwareVlanRange) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *VirtualizationVmwareVlanRange) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *VirtualizationVmwareVlanRange) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VirtualizationVmwareVlanRange) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VirtualizationVmwareVlanRange) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetVlanRangeEnd

`func (o *VirtualizationVmwareVlanRange) GetVlanRangeEnd() int64`

GetVlanRangeEnd returns the VlanRangeEnd field if non-nil, zero value otherwise.

### GetVlanRangeEndOk

`func (o *VirtualizationVmwareVlanRange) GetVlanRangeEndOk() (*int64, bool)`

GetVlanRangeEndOk returns a tuple with the VlanRangeEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanRangeEnd

`func (o *VirtualizationVmwareVlanRange) SetVlanRangeEnd(v int64)`

SetVlanRangeEnd sets VlanRangeEnd field to given value.

### HasVlanRangeEnd

`func (o *VirtualizationVmwareVlanRange) HasVlanRangeEnd() bool`

HasVlanRangeEnd returns a boolean if a field has been set.

### GetVlanRangeStart

`func (o *VirtualizationVmwareVlanRange) GetVlanRangeStart() int64`

GetVlanRangeStart returns the VlanRangeStart field if non-nil, zero value otherwise.

### GetVlanRangeStartOk

`func (o *VirtualizationVmwareVlanRange) GetVlanRangeStartOk() (*int64, bool)`

GetVlanRangeStartOk returns a tuple with the VlanRangeStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanRangeStart

`func (o *VirtualizationVmwareVlanRange) SetVlanRangeStart(v int64)`

SetVlanRangeStart sets VlanRangeStart field to given value.

### HasVlanRangeStart

`func (o *VirtualizationVmwareVlanRange) HasVlanRangeStart() bool`

HasVlanRangeStart returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


