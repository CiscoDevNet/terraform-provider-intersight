# NiatelemetryDiskinfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "niatelemetry.Diskinfo"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "niatelemetry.Diskinfo"]
**Free** | Pointer to **int64** | The free disk capacity, currently the type of this field is set to integer. This determines how much memory is free in Bytes. | [optional] 
**Name** | Pointer to **string** | Disk Name used to identified the disk usage record. This determines the name of the disk partition that is inventoried. | [optional] 
**Total** | Pointer to **int64** | The total disk capacity, it should be the sum of free and used, currently the type of this field is set to integer. This determines the total memory for this partition. | [optional] 
**Used** | Pointer to **int64** | The used disk capacity, currently the type of this field is set to integer. This determines how much memory is used in Bytes. | [optional] 

## Methods

### NewNiatelemetryDiskinfo

`func NewNiatelemetryDiskinfo(classId string, objectType string, ) *NiatelemetryDiskinfo`

NewNiatelemetryDiskinfo instantiates a new NiatelemetryDiskinfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiatelemetryDiskinfoWithDefaults

`func NewNiatelemetryDiskinfoWithDefaults() *NiatelemetryDiskinfo`

NewNiatelemetryDiskinfoWithDefaults instantiates a new NiatelemetryDiskinfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NiatelemetryDiskinfo) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NiatelemetryDiskinfo) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NiatelemetryDiskinfo) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NiatelemetryDiskinfo) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NiatelemetryDiskinfo) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NiatelemetryDiskinfo) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


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


