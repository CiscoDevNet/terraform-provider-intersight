# VirtualizationVmwareResourceAllocationSystemTrafficTypesAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "virtualization.VmwareResourceAllocationSystemTrafficTypes"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "virtualization.VmwareResourceAllocationSystemTrafficTypes"]
**Limit** | Pointer to **int64** | The maximum allowed usage for a traffic class belonging to this resource pool per host physical NIC. The utilization of a traffic class will not exceed the specified limit even if there are available network resources. If this value is unset or set to -1 in an update operation, then there is no limit on the network resource usage (only bounded by available resource and shares). Units are in Mbits/sec. | [optional] 
**Reservation** | Pointer to **int64** | Amount of bandwidth resource that is guaranteed available to the host infrastructure traffic class. If the utilization is less than the reservation, the extra bandwidth is used for other host infrastructure traffic class types. | [optional] 
**Shares** | Pointer to **string** | Network share. The value is used as a relative weight in competing for shared bandwidth, in case of resource contention. * &#x60;low&#x60; - Share of the traffic in the overall flow through a physical adapter. * &#x60;high&#x60; - Share of the traffic in the overall flow through a physical adapter. * &#x60;normal&#x60; - Share of the traffic in the overall flow through a physical adapter. * &#x60;custom&#x60; - Share of the traffic in the overall flow through a physical adapter. | [optional] [default to "low"]
**SharesValue** | Pointer to **int32** | The number of shares allocated. Used to determine resource allocation in case of resource contention. Shares value is only set if level is set to custom. If level is not set to custom, this value is ignored. Therefore, only shares with custom values can be compared. | [optional] 
**TrafficType** | Pointer to **string** | Key of the host infrastructure resource. | [optional] 

## Methods

### NewVirtualizationVmwareResourceAllocationSystemTrafficTypesAllOf

`func NewVirtualizationVmwareResourceAllocationSystemTrafficTypesAllOf(classId string, objectType string, ) *VirtualizationVmwareResourceAllocationSystemTrafficTypesAllOf`

NewVirtualizationVmwareResourceAllocationSystemTrafficTypesAllOf instantiates a new VirtualizationVmwareResourceAllocationSystemTrafficTypesAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualizationVmwareResourceAllocationSystemTrafficTypesAllOfWithDefaults

`func NewVirtualizationVmwareResourceAllocationSystemTrafficTypesAllOfWithDefaults() *VirtualizationVmwareResourceAllocationSystemTrafficTypesAllOf`

NewVirtualizationVmwareResourceAllocationSystemTrafficTypesAllOfWithDefaults instantiates a new VirtualizationVmwareResourceAllocationSystemTrafficTypesAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *VirtualizationVmwareResourceAllocationSystemTrafficTypesAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *VirtualizationVmwareResourceAllocationSystemTrafficTypesAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *VirtualizationVmwareResourceAllocationSystemTrafficTypesAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *VirtualizationVmwareResourceAllocationSystemTrafficTypesAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VirtualizationVmwareResourceAllocationSystemTrafficTypesAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VirtualizationVmwareResourceAllocationSystemTrafficTypesAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetLimit

`func (o *VirtualizationVmwareResourceAllocationSystemTrafficTypesAllOf) GetLimit() int64`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *VirtualizationVmwareResourceAllocationSystemTrafficTypesAllOf) GetLimitOk() (*int64, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *VirtualizationVmwareResourceAllocationSystemTrafficTypesAllOf) SetLimit(v int64)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *VirtualizationVmwareResourceAllocationSystemTrafficTypesAllOf) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetReservation

`func (o *VirtualizationVmwareResourceAllocationSystemTrafficTypesAllOf) GetReservation() int64`

GetReservation returns the Reservation field if non-nil, zero value otherwise.

### GetReservationOk

`func (o *VirtualizationVmwareResourceAllocationSystemTrafficTypesAllOf) GetReservationOk() (*int64, bool)`

GetReservationOk returns a tuple with the Reservation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservation

`func (o *VirtualizationVmwareResourceAllocationSystemTrafficTypesAllOf) SetReservation(v int64)`

SetReservation sets Reservation field to given value.

### HasReservation

`func (o *VirtualizationVmwareResourceAllocationSystemTrafficTypesAllOf) HasReservation() bool`

HasReservation returns a boolean if a field has been set.

### GetShares

`func (o *VirtualizationVmwareResourceAllocationSystemTrafficTypesAllOf) GetShares() string`

GetShares returns the Shares field if non-nil, zero value otherwise.

### GetSharesOk

`func (o *VirtualizationVmwareResourceAllocationSystemTrafficTypesAllOf) GetSharesOk() (*string, bool)`

GetSharesOk returns a tuple with the Shares field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShares

`func (o *VirtualizationVmwareResourceAllocationSystemTrafficTypesAllOf) SetShares(v string)`

SetShares sets Shares field to given value.

### HasShares

`func (o *VirtualizationVmwareResourceAllocationSystemTrafficTypesAllOf) HasShares() bool`

HasShares returns a boolean if a field has been set.

### GetSharesValue

`func (o *VirtualizationVmwareResourceAllocationSystemTrafficTypesAllOf) GetSharesValue() int32`

GetSharesValue returns the SharesValue field if non-nil, zero value otherwise.

### GetSharesValueOk

`func (o *VirtualizationVmwareResourceAllocationSystemTrafficTypesAllOf) GetSharesValueOk() (*int32, bool)`

GetSharesValueOk returns a tuple with the SharesValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharesValue

`func (o *VirtualizationVmwareResourceAllocationSystemTrafficTypesAllOf) SetSharesValue(v int32)`

SetSharesValue sets SharesValue field to given value.

### HasSharesValue

`func (o *VirtualizationVmwareResourceAllocationSystemTrafficTypesAllOf) HasSharesValue() bool`

HasSharesValue returns a boolean if a field has been set.

### GetTrafficType

`func (o *VirtualizationVmwareResourceAllocationSystemTrafficTypesAllOf) GetTrafficType() string`

GetTrafficType returns the TrafficType field if non-nil, zero value otherwise.

### GetTrafficTypeOk

`func (o *VirtualizationVmwareResourceAllocationSystemTrafficTypesAllOf) GetTrafficTypeOk() (*string, bool)`

GetTrafficTypeOk returns a tuple with the TrafficType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficType

`func (o *VirtualizationVmwareResourceAllocationSystemTrafficTypesAllOf) SetTrafficType(v string)`

SetTrafficType sets TrafficType field to given value.

### HasTrafficType

`func (o *VirtualizationVmwareResourceAllocationSystemTrafficTypesAllOf) HasTrafficType() bool`

HasTrafficType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


