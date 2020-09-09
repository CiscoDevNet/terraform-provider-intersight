# NiatelemetryDiskinfoAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Free** | Pointer to **int64** | The free disk capacity, currently the type of this field is set to integer. This determines how much memory is free in Bytes. | [optional] 
**Name** | Pointer to **string** | Disk Name used to identified the disk usage record. This determines the name of the disk partition that is inventoried. | [optional] 
**Total** | Pointer to **int64** | The total disk capacity, it should be the sum of free and used, currently the type of this field is set to integer. This determines the total memory for this partition. | [optional] 
**Used** | Pointer to **int64** | The used disk capacity, currently the type of this field is set to integer. This determines how much memory is used in Bytes. | [optional] 

## Methods

### NewNiatelemetryDiskinfoAllOf

`func NewNiatelemetryDiskinfoAllOf() *NiatelemetryDiskinfoAllOf`

NewNiatelemetryDiskinfoAllOf instantiates a new NiatelemetryDiskinfoAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiatelemetryDiskinfoAllOfWithDefaults

`func NewNiatelemetryDiskinfoAllOfWithDefaults() *NiatelemetryDiskinfoAllOf`

NewNiatelemetryDiskinfoAllOfWithDefaults instantiates a new NiatelemetryDiskinfoAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFree

`func (o *NiatelemetryDiskinfoAllOf) GetFree() int64`

GetFree returns the Free field if non-nil, zero value otherwise.

### GetFreeOk

`func (o *NiatelemetryDiskinfoAllOf) GetFreeOk() (*int64, bool)`

GetFreeOk returns a tuple with the Free field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFree

`func (o *NiatelemetryDiskinfoAllOf) SetFree(v int64)`

SetFree sets Free field to given value.

### HasFree

`func (o *NiatelemetryDiskinfoAllOf) HasFree() bool`

HasFree returns a boolean if a field has been set.

### GetName

`func (o *NiatelemetryDiskinfoAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NiatelemetryDiskinfoAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NiatelemetryDiskinfoAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NiatelemetryDiskinfoAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTotal

`func (o *NiatelemetryDiskinfoAllOf) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *NiatelemetryDiskinfoAllOf) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *NiatelemetryDiskinfoAllOf) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *NiatelemetryDiskinfoAllOf) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetUsed

`func (o *NiatelemetryDiskinfoAllOf) GetUsed() int64`

GetUsed returns the Used field if non-nil, zero value otherwise.

### GetUsedOk

`func (o *NiatelemetryDiskinfoAllOf) GetUsedOk() (*int64, bool)`

GetUsedOk returns a tuple with the Used field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsed

`func (o *NiatelemetryDiskinfoAllOf) SetUsed(v int64)`

SetUsed sets Used field to given value.

### HasUsed

`func (o *NiatelemetryDiskinfoAllOf) HasUsed() bool`

HasUsed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


