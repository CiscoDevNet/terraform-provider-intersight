# VirtualizationCpuAllocation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "virtualization.CpuAllocation"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "virtualization.CpuAllocation"]
**Free** | Pointer to **int64** | Free CPU count on the entity. | [optional] 
**Reserved** | Pointer to **int64** | Reserved CPU count on the entity. These reserved CPUs can be used for system purposes. | [optional] 
**Total** | Pointer to **int64** | Total number of CPU available on the entity. | [optional] 
**Used** | Pointer to **int64** | Used or allocated CPU count on the entity. | [optional] 

## Methods

### NewVirtualizationCpuAllocation

`func NewVirtualizationCpuAllocation(classId string, objectType string, ) *VirtualizationCpuAllocation`

NewVirtualizationCpuAllocation instantiates a new VirtualizationCpuAllocation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualizationCpuAllocationWithDefaults

`func NewVirtualizationCpuAllocationWithDefaults() *VirtualizationCpuAllocation`

NewVirtualizationCpuAllocationWithDefaults instantiates a new VirtualizationCpuAllocation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *VirtualizationCpuAllocation) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *VirtualizationCpuAllocation) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *VirtualizationCpuAllocation) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *VirtualizationCpuAllocation) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VirtualizationCpuAllocation) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VirtualizationCpuAllocation) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetFree

`func (o *VirtualizationCpuAllocation) GetFree() int64`

GetFree returns the Free field if non-nil, zero value otherwise.

### GetFreeOk

`func (o *VirtualizationCpuAllocation) GetFreeOk() (*int64, bool)`

GetFreeOk returns a tuple with the Free field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFree

`func (o *VirtualizationCpuAllocation) SetFree(v int64)`

SetFree sets Free field to given value.

### HasFree

`func (o *VirtualizationCpuAllocation) HasFree() bool`

HasFree returns a boolean if a field has been set.

### GetReserved

`func (o *VirtualizationCpuAllocation) GetReserved() int64`

GetReserved returns the Reserved field if non-nil, zero value otherwise.

### GetReservedOk

`func (o *VirtualizationCpuAllocation) GetReservedOk() (*int64, bool)`

GetReservedOk returns a tuple with the Reserved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReserved

`func (o *VirtualizationCpuAllocation) SetReserved(v int64)`

SetReserved sets Reserved field to given value.

### HasReserved

`func (o *VirtualizationCpuAllocation) HasReserved() bool`

HasReserved returns a boolean if a field has been set.

### GetTotal

`func (o *VirtualizationCpuAllocation) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *VirtualizationCpuAllocation) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *VirtualizationCpuAllocation) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *VirtualizationCpuAllocation) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetUsed

`func (o *VirtualizationCpuAllocation) GetUsed() int64`

GetUsed returns the Used field if non-nil, zero value otherwise.

### GetUsedOk

`func (o *VirtualizationCpuAllocation) GetUsedOk() (*int64, bool)`

GetUsedOk returns a tuple with the Used field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsed

`func (o *VirtualizationCpuAllocation) SetUsed(v int64)`

SetUsed sets Used field to given value.

### HasUsed

`func (o *VirtualizationCpuAllocation) HasUsed() bool`

HasUsed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


