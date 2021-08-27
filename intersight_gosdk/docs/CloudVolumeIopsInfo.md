# CloudVolumeIopsInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "cloud.VolumeIopsInfo"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "cloud.VolumeIopsInfo"]
**IopsReadLimit** | Pointer to **int64** | Number of disk read commands that can be performed per second. | [optional] [readonly] 
**IopsWriteLimit** | Pointer to **int64** | Number of disk write commands that can be performed per second. | [optional] [readonly] 
**ThroughputReadLimit** | Pointer to **int64** | Data transfer rate limit from the disk, specified in mebibytes (MiB) per second. | [optional] [readonly] 
**ThroughputWriteLimit** | Pointer to **int64** | Data transfer rate limit to the disk, specified in mebibytes (MiB) per second. | [optional] [readonly] 

## Methods

### NewCloudVolumeIopsInfo

`func NewCloudVolumeIopsInfo(classId string, objectType string, ) *CloudVolumeIopsInfo`

NewCloudVolumeIopsInfo instantiates a new CloudVolumeIopsInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudVolumeIopsInfoWithDefaults

`func NewCloudVolumeIopsInfoWithDefaults() *CloudVolumeIopsInfo`

NewCloudVolumeIopsInfoWithDefaults instantiates a new CloudVolumeIopsInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *CloudVolumeIopsInfo) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *CloudVolumeIopsInfo) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *CloudVolumeIopsInfo) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *CloudVolumeIopsInfo) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CloudVolumeIopsInfo) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CloudVolumeIopsInfo) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetIopsReadLimit

`func (o *CloudVolumeIopsInfo) GetIopsReadLimit() int64`

GetIopsReadLimit returns the IopsReadLimit field if non-nil, zero value otherwise.

### GetIopsReadLimitOk

`func (o *CloudVolumeIopsInfo) GetIopsReadLimitOk() (*int64, bool)`

GetIopsReadLimitOk returns a tuple with the IopsReadLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIopsReadLimit

`func (o *CloudVolumeIopsInfo) SetIopsReadLimit(v int64)`

SetIopsReadLimit sets IopsReadLimit field to given value.

### HasIopsReadLimit

`func (o *CloudVolumeIopsInfo) HasIopsReadLimit() bool`

HasIopsReadLimit returns a boolean if a field has been set.

### GetIopsWriteLimit

`func (o *CloudVolumeIopsInfo) GetIopsWriteLimit() int64`

GetIopsWriteLimit returns the IopsWriteLimit field if non-nil, zero value otherwise.

### GetIopsWriteLimitOk

`func (o *CloudVolumeIopsInfo) GetIopsWriteLimitOk() (*int64, bool)`

GetIopsWriteLimitOk returns a tuple with the IopsWriteLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIopsWriteLimit

`func (o *CloudVolumeIopsInfo) SetIopsWriteLimit(v int64)`

SetIopsWriteLimit sets IopsWriteLimit field to given value.

### HasIopsWriteLimit

`func (o *CloudVolumeIopsInfo) HasIopsWriteLimit() bool`

HasIopsWriteLimit returns a boolean if a field has been set.

### GetThroughputReadLimit

`func (o *CloudVolumeIopsInfo) GetThroughputReadLimit() int64`

GetThroughputReadLimit returns the ThroughputReadLimit field if non-nil, zero value otherwise.

### GetThroughputReadLimitOk

`func (o *CloudVolumeIopsInfo) GetThroughputReadLimitOk() (*int64, bool)`

GetThroughputReadLimitOk returns a tuple with the ThroughputReadLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThroughputReadLimit

`func (o *CloudVolumeIopsInfo) SetThroughputReadLimit(v int64)`

SetThroughputReadLimit sets ThroughputReadLimit field to given value.

### HasThroughputReadLimit

`func (o *CloudVolumeIopsInfo) HasThroughputReadLimit() bool`

HasThroughputReadLimit returns a boolean if a field has been set.

### GetThroughputWriteLimit

`func (o *CloudVolumeIopsInfo) GetThroughputWriteLimit() int64`

GetThroughputWriteLimit returns the ThroughputWriteLimit field if non-nil, zero value otherwise.

### GetThroughputWriteLimitOk

`func (o *CloudVolumeIopsInfo) GetThroughputWriteLimitOk() (*int64, bool)`

GetThroughputWriteLimitOk returns a tuple with the ThroughputWriteLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThroughputWriteLimit

`func (o *CloudVolumeIopsInfo) SetThroughputWriteLimit(v int64)`

SetThroughputWriteLimit sets ThroughputWriteLimit field to given value.

### HasThroughputWriteLimit

`func (o *CloudVolumeIopsInfo) HasThroughputWriteLimit() bool`

HasThroughputWriteLimit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


