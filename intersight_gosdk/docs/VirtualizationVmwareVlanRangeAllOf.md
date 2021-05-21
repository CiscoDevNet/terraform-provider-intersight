# VirtualizationVmwareVlanRangeAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "virtualization.VmwareVlanRange"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "virtualization.VmwareVlanRange"]
**VlanRangeEnd** | Pointer to **int64** | End value of VLAN range for the trunk port. The valid range is from 0 to 4094. | [optional] 
**VlanRangeStart** | Pointer to **int64** | Start value of VLAN range for the trunk port. The valid range is from 0 to 4094. | [optional] 

## Methods

### NewVirtualizationVmwareVlanRangeAllOf

`func NewVirtualizationVmwareVlanRangeAllOf(classId string, objectType string, ) *VirtualizationVmwareVlanRangeAllOf`

NewVirtualizationVmwareVlanRangeAllOf instantiates a new VirtualizationVmwareVlanRangeAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualizationVmwareVlanRangeAllOfWithDefaults

`func NewVirtualizationVmwareVlanRangeAllOfWithDefaults() *VirtualizationVmwareVlanRangeAllOf`

NewVirtualizationVmwareVlanRangeAllOfWithDefaults instantiates a new VirtualizationVmwareVlanRangeAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *VirtualizationVmwareVlanRangeAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *VirtualizationVmwareVlanRangeAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *VirtualizationVmwareVlanRangeAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *VirtualizationVmwareVlanRangeAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VirtualizationVmwareVlanRangeAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VirtualizationVmwareVlanRangeAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetVlanRangeEnd

`func (o *VirtualizationVmwareVlanRangeAllOf) GetVlanRangeEnd() int64`

GetVlanRangeEnd returns the VlanRangeEnd field if non-nil, zero value otherwise.

### GetVlanRangeEndOk

`func (o *VirtualizationVmwareVlanRangeAllOf) GetVlanRangeEndOk() (*int64, bool)`

GetVlanRangeEndOk returns a tuple with the VlanRangeEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanRangeEnd

`func (o *VirtualizationVmwareVlanRangeAllOf) SetVlanRangeEnd(v int64)`

SetVlanRangeEnd sets VlanRangeEnd field to given value.

### HasVlanRangeEnd

`func (o *VirtualizationVmwareVlanRangeAllOf) HasVlanRangeEnd() bool`

HasVlanRangeEnd returns a boolean if a field has been set.

### GetVlanRangeStart

`func (o *VirtualizationVmwareVlanRangeAllOf) GetVlanRangeStart() int64`

GetVlanRangeStart returns the VlanRangeStart field if non-nil, zero value otherwise.

### GetVlanRangeStartOk

`func (o *VirtualizationVmwareVlanRangeAllOf) GetVlanRangeStartOk() (*int64, bool)`

GetVlanRangeStartOk returns a tuple with the VlanRangeStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanRangeStart

`func (o *VirtualizationVmwareVlanRangeAllOf) SetVlanRangeStart(v int64)`

SetVlanRangeStart sets VlanRangeStart field to given value.

### HasVlanRangeStart

`func (o *VirtualizationVmwareVlanRangeAllOf) HasVlanRangeStart() bool`

HasVlanRangeStart returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


