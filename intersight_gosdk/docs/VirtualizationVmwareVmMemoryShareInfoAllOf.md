# VirtualizationVmwareVmMemoryShareInfoAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MemLimit** | Pointer to **int64** | Limit on the memory sharing imposed (in Mbytes). | [optional] 
**MemOverheadLimit** | Pointer to **int64** | Limit on memory overhead imposed (in Mbytes). | [optional] 
**MemReservation** | Pointer to **int64** | Similar to CPU reservations (Mbytes). | [optional] 
**MemShares** | Pointer to **int64** | Similar to CPU Shares but applicable to memory. There is no unit for this value. It is a relative measure based on the settings for other resource pools. | [optional] 

## Methods

### NewVirtualizationVmwareVmMemoryShareInfoAllOf

`func NewVirtualizationVmwareVmMemoryShareInfoAllOf() *VirtualizationVmwareVmMemoryShareInfoAllOf`

NewVirtualizationVmwareVmMemoryShareInfoAllOf instantiates a new VirtualizationVmwareVmMemoryShareInfoAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualizationVmwareVmMemoryShareInfoAllOfWithDefaults

`func NewVirtualizationVmwareVmMemoryShareInfoAllOfWithDefaults() *VirtualizationVmwareVmMemoryShareInfoAllOf`

NewVirtualizationVmwareVmMemoryShareInfoAllOfWithDefaults instantiates a new VirtualizationVmwareVmMemoryShareInfoAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMemLimit

`func (o *VirtualizationVmwareVmMemoryShareInfoAllOf) GetMemLimit() int64`

GetMemLimit returns the MemLimit field if non-nil, zero value otherwise.

### GetMemLimitOk

`func (o *VirtualizationVmwareVmMemoryShareInfoAllOf) GetMemLimitOk() (*int64, bool)`

GetMemLimitOk returns a tuple with the MemLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemLimit

`func (o *VirtualizationVmwareVmMemoryShareInfoAllOf) SetMemLimit(v int64)`

SetMemLimit sets MemLimit field to given value.

### HasMemLimit

`func (o *VirtualizationVmwareVmMemoryShareInfoAllOf) HasMemLimit() bool`

HasMemLimit returns a boolean if a field has been set.

### GetMemOverheadLimit

`func (o *VirtualizationVmwareVmMemoryShareInfoAllOf) GetMemOverheadLimit() int64`

GetMemOverheadLimit returns the MemOverheadLimit field if non-nil, zero value otherwise.

### GetMemOverheadLimitOk

`func (o *VirtualizationVmwareVmMemoryShareInfoAllOf) GetMemOverheadLimitOk() (*int64, bool)`

GetMemOverheadLimitOk returns a tuple with the MemOverheadLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemOverheadLimit

`func (o *VirtualizationVmwareVmMemoryShareInfoAllOf) SetMemOverheadLimit(v int64)`

SetMemOverheadLimit sets MemOverheadLimit field to given value.

### HasMemOverheadLimit

`func (o *VirtualizationVmwareVmMemoryShareInfoAllOf) HasMemOverheadLimit() bool`

HasMemOverheadLimit returns a boolean if a field has been set.

### GetMemReservation

`func (o *VirtualizationVmwareVmMemoryShareInfoAllOf) GetMemReservation() int64`

GetMemReservation returns the MemReservation field if non-nil, zero value otherwise.

### GetMemReservationOk

`func (o *VirtualizationVmwareVmMemoryShareInfoAllOf) GetMemReservationOk() (*int64, bool)`

GetMemReservationOk returns a tuple with the MemReservation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemReservation

`func (o *VirtualizationVmwareVmMemoryShareInfoAllOf) SetMemReservation(v int64)`

SetMemReservation sets MemReservation field to given value.

### HasMemReservation

`func (o *VirtualizationVmwareVmMemoryShareInfoAllOf) HasMemReservation() bool`

HasMemReservation returns a boolean if a field has been set.

### GetMemShares

`func (o *VirtualizationVmwareVmMemoryShareInfoAllOf) GetMemShares() int64`

GetMemShares returns the MemShares field if non-nil, zero value otherwise.

### GetMemSharesOk

`func (o *VirtualizationVmwareVmMemoryShareInfoAllOf) GetMemSharesOk() (*int64, bool)`

GetMemSharesOk returns a tuple with the MemShares field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemShares

`func (o *VirtualizationVmwareVmMemoryShareInfoAllOf) SetMemShares(v int64)`

SetMemShares sets MemShares field to given value.

### HasMemShares

`func (o *VirtualizationVmwareVmMemoryShareInfoAllOf) HasMemShares() bool`

HasMemShares returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


