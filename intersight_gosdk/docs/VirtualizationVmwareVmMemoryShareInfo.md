# VirtualizationVmwareVmMemoryShareInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "virtualization.VmwareVmMemoryShareInfo"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "virtualization.VmwareVmMemoryShareInfo"]
**MemLimit** | Pointer to **int64** | Limit on the memory sharing imposed (in Mbytes). | [optional] 
**MemOverheadLimit** | Pointer to **int64** | Limit on memory overhead imposed (in Mbytes). | [optional] 
**MemReservation** | Pointer to **int64** | Similar to CPU reservations (Mbytes). | [optional] 
**MemShares** | Pointer to **int64** | Similar to CPU Shares but applicable to memory. There is no unit for this value. It is a relative measure based on the settings for other resource pools. | [optional] 

## Methods

### NewVirtualizationVmwareVmMemoryShareInfo

`func NewVirtualizationVmwareVmMemoryShareInfo(classId string, objectType string, ) *VirtualizationVmwareVmMemoryShareInfo`

NewVirtualizationVmwareVmMemoryShareInfo instantiates a new VirtualizationVmwareVmMemoryShareInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualizationVmwareVmMemoryShareInfoWithDefaults

`func NewVirtualizationVmwareVmMemoryShareInfoWithDefaults() *VirtualizationVmwareVmMemoryShareInfo`

NewVirtualizationVmwareVmMemoryShareInfoWithDefaults instantiates a new VirtualizationVmwareVmMemoryShareInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *VirtualizationVmwareVmMemoryShareInfo) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *VirtualizationVmwareVmMemoryShareInfo) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *VirtualizationVmwareVmMemoryShareInfo) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *VirtualizationVmwareVmMemoryShareInfo) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VirtualizationVmwareVmMemoryShareInfo) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VirtualizationVmwareVmMemoryShareInfo) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetMemLimit

`func (o *VirtualizationVmwareVmMemoryShareInfo) GetMemLimit() int64`

GetMemLimit returns the MemLimit field if non-nil, zero value otherwise.

### GetMemLimitOk

`func (o *VirtualizationVmwareVmMemoryShareInfo) GetMemLimitOk() (*int64, bool)`

GetMemLimitOk returns a tuple with the MemLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemLimit

`func (o *VirtualizationVmwareVmMemoryShareInfo) SetMemLimit(v int64)`

SetMemLimit sets MemLimit field to given value.

### HasMemLimit

`func (o *VirtualizationVmwareVmMemoryShareInfo) HasMemLimit() bool`

HasMemLimit returns a boolean if a field has been set.

### GetMemOverheadLimit

`func (o *VirtualizationVmwareVmMemoryShareInfo) GetMemOverheadLimit() int64`

GetMemOverheadLimit returns the MemOverheadLimit field if non-nil, zero value otherwise.

### GetMemOverheadLimitOk

`func (o *VirtualizationVmwareVmMemoryShareInfo) GetMemOverheadLimitOk() (*int64, bool)`

GetMemOverheadLimitOk returns a tuple with the MemOverheadLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemOverheadLimit

`func (o *VirtualizationVmwareVmMemoryShareInfo) SetMemOverheadLimit(v int64)`

SetMemOverheadLimit sets MemOverheadLimit field to given value.

### HasMemOverheadLimit

`func (o *VirtualizationVmwareVmMemoryShareInfo) HasMemOverheadLimit() bool`

HasMemOverheadLimit returns a boolean if a field has been set.

### GetMemReservation

`func (o *VirtualizationVmwareVmMemoryShareInfo) GetMemReservation() int64`

GetMemReservation returns the MemReservation field if non-nil, zero value otherwise.

### GetMemReservationOk

`func (o *VirtualizationVmwareVmMemoryShareInfo) GetMemReservationOk() (*int64, bool)`

GetMemReservationOk returns a tuple with the MemReservation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemReservation

`func (o *VirtualizationVmwareVmMemoryShareInfo) SetMemReservation(v int64)`

SetMemReservation sets MemReservation field to given value.

### HasMemReservation

`func (o *VirtualizationVmwareVmMemoryShareInfo) HasMemReservation() bool`

HasMemReservation returns a boolean if a field has been set.

### GetMemShares

`func (o *VirtualizationVmwareVmMemoryShareInfo) GetMemShares() int64`

GetMemShares returns the MemShares field if non-nil, zero value otherwise.

### GetMemSharesOk

`func (o *VirtualizationVmwareVmMemoryShareInfo) GetMemSharesOk() (*int64, bool)`

GetMemSharesOk returns a tuple with the MemShares field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemShares

`func (o *VirtualizationVmwareVmMemoryShareInfo) SetMemShares(v int64)`

SetMemShares sets MemShares field to given value.

### HasMemShares

`func (o *VirtualizationVmwareVmMemoryShareInfo) HasMemShares() bool`

HasMemShares returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


