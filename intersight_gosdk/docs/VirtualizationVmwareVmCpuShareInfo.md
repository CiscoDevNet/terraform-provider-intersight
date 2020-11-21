# VirtualizationVmwareVmCpuShareInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "virtualization.VmwareVmCpuShareInfo"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "virtualization.VmwareVmCpuShareInfo"]
**CpuLimit** | Pointer to **int64** | Upper limit on CPU allocation (MHz). | [optional] 
**CpuOverheadLimit** | Pointer to **int64** | Amount of CPU for virtualization overhead. | [optional] 
**CpuReservation** | Pointer to **int64** | Guaranteed minimum allocation of CPU resource (MHz). | [optional] 
**CpuShares** | Pointer to **int64** | Shows the relative importance of a VM. There is no unit for this value. It is a relative measure based on the settings for other resource pools. For more information, see VMware documentation. | [optional] 

## Methods

### NewVirtualizationVmwareVmCpuShareInfo

`func NewVirtualizationVmwareVmCpuShareInfo(classId string, objectType string, ) *VirtualizationVmwareVmCpuShareInfo`

NewVirtualizationVmwareVmCpuShareInfo instantiates a new VirtualizationVmwareVmCpuShareInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualizationVmwareVmCpuShareInfoWithDefaults

`func NewVirtualizationVmwareVmCpuShareInfoWithDefaults() *VirtualizationVmwareVmCpuShareInfo`

NewVirtualizationVmwareVmCpuShareInfoWithDefaults instantiates a new VirtualizationVmwareVmCpuShareInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *VirtualizationVmwareVmCpuShareInfo) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *VirtualizationVmwareVmCpuShareInfo) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *VirtualizationVmwareVmCpuShareInfo) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *VirtualizationVmwareVmCpuShareInfo) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VirtualizationVmwareVmCpuShareInfo) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VirtualizationVmwareVmCpuShareInfo) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCpuLimit

`func (o *VirtualizationVmwareVmCpuShareInfo) GetCpuLimit() int64`

GetCpuLimit returns the CpuLimit field if non-nil, zero value otherwise.

### GetCpuLimitOk

`func (o *VirtualizationVmwareVmCpuShareInfo) GetCpuLimitOk() (*int64, bool)`

GetCpuLimitOk returns a tuple with the CpuLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuLimit

`func (o *VirtualizationVmwareVmCpuShareInfo) SetCpuLimit(v int64)`

SetCpuLimit sets CpuLimit field to given value.

### HasCpuLimit

`func (o *VirtualizationVmwareVmCpuShareInfo) HasCpuLimit() bool`

HasCpuLimit returns a boolean if a field has been set.

### GetCpuOverheadLimit

`func (o *VirtualizationVmwareVmCpuShareInfo) GetCpuOverheadLimit() int64`

GetCpuOverheadLimit returns the CpuOverheadLimit field if non-nil, zero value otherwise.

### GetCpuOverheadLimitOk

`func (o *VirtualizationVmwareVmCpuShareInfo) GetCpuOverheadLimitOk() (*int64, bool)`

GetCpuOverheadLimitOk returns a tuple with the CpuOverheadLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuOverheadLimit

`func (o *VirtualizationVmwareVmCpuShareInfo) SetCpuOverheadLimit(v int64)`

SetCpuOverheadLimit sets CpuOverheadLimit field to given value.

### HasCpuOverheadLimit

`func (o *VirtualizationVmwareVmCpuShareInfo) HasCpuOverheadLimit() bool`

HasCpuOverheadLimit returns a boolean if a field has been set.

### GetCpuReservation

`func (o *VirtualizationVmwareVmCpuShareInfo) GetCpuReservation() int64`

GetCpuReservation returns the CpuReservation field if non-nil, zero value otherwise.

### GetCpuReservationOk

`func (o *VirtualizationVmwareVmCpuShareInfo) GetCpuReservationOk() (*int64, bool)`

GetCpuReservationOk returns a tuple with the CpuReservation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuReservation

`func (o *VirtualizationVmwareVmCpuShareInfo) SetCpuReservation(v int64)`

SetCpuReservation sets CpuReservation field to given value.

### HasCpuReservation

`func (o *VirtualizationVmwareVmCpuShareInfo) HasCpuReservation() bool`

HasCpuReservation returns a boolean if a field has been set.

### GetCpuShares

`func (o *VirtualizationVmwareVmCpuShareInfo) GetCpuShares() int64`

GetCpuShares returns the CpuShares field if non-nil, zero value otherwise.

### GetCpuSharesOk

`func (o *VirtualizationVmwareVmCpuShareInfo) GetCpuSharesOk() (*int64, bool)`

GetCpuSharesOk returns a tuple with the CpuShares field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuShares

`func (o *VirtualizationVmwareVmCpuShareInfo) SetCpuShares(v int64)`

SetCpuShares sets CpuShares field to given value.

### HasCpuShares

`func (o *VirtualizationVmwareVmCpuShareInfo) HasCpuShares() bool`

HasCpuShares returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


