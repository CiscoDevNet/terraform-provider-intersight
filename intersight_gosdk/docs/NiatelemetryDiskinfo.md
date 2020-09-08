# NiatelemetryDiskinfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Free** | Pointer to **int64** | The free disk capacity, currently the type of this field is set to integer. This determines how much memory is free in Bytes. | [optional] 
**Name** | Pointer to **string** | Disk Name used to identified the disk usage record. This determines the name of the disk partition that is inventoried. | [optional] 
**Total** | Pointer to **int64** | The total disk capacity, it should be the sum of free and used, currently the type of this field is set to integer. This determines the total memory for this partition. | [optional] 
**Used** | Pointer to **int64** | The used disk capacity, currently the type of this field is set to integer. This determines how much memory is used in Bytes. | [optional] 

## Methods

### NewNiatelemetryDiskinfo

`func NewNiatelemetryDiskinfo() *NiatelemetryDiskinfo`

NewNiatelemetryDiskinfo instantiates a new NiatelemetryDiskinfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiatelemetryDiskinfoWithDefaults

`func NewNiatelemetryDiskinfoWithDefaults() *NiatelemetryDiskinfo`

NewNiatelemetryDiskinfoWithDefaults instantiates a new NiatelemetryDiskinfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFree

`func (o *NiatelemetryDiskinfo) GetFree() int64`

GetFree returns the Free field if non-nil, zero value otherwise.

### GetFreeOk

`func (o *NiatelemetryDiskinfo) GetFreeOk() (*int64, bool)`

GetFreeOk returns a tuple with the Free field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFree

`func (o *NiatelemetryDiskinfo) SetFree(v int64)`

SetFree sets Free field to given value.

### HasFree

`func (o *NiatelemetryDiskinfo) HasFree() bool`

HasFree returns a boolean if a field has been set.

### GetName

`func (o *NiatelemetryDiskinfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NiatelemetryDiskinfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NiatelemetryDiskinfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NiatelemetryDiskinfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTotal

`func (o *NiatelemetryDiskinfo) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *NiatelemetryDiskinfo) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *NiatelemetryDiskinfo) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *NiatelemetryDiskinfo) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetUsed

`func (o *NiatelemetryDiskinfo) GetUsed() int64`

GetUsed returns the Used field if non-nil, zero value otherwise.

### GetUsedOk

`func (o *NiatelemetryDiskinfo) GetUsedOk() (*int64, bool)`

GetUsedOk returns a tuple with the Used field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsed

`func (o *NiatelemetryDiskinfo) SetUsed(v int64)`

SetUsed sets Used field to given value.

### HasUsed

`func (o *NiatelemetryDiskinfo) HasUsed() bool`

HasUsed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


