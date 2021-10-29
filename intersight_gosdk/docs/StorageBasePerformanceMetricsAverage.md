# StorageBasePerformanceMetricsAverage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | [default to "storage.NetAppPerformanceMetricsAverage"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | [default to "storage.NetAppPerformanceMetricsAverage"]
**Iops** | Pointer to **float64** | Rate of I/O operations observed at the storage object. | [optional] [readonly] 
**Latency** | Pointer to **float64** | Latency observed at the storage object. | [optional] [readonly] 
**Period** | Pointer to **int64** | Duration of periodic aggregation, in hours. | [optional] [readonly] 
**Throughput** | Pointer to **float64** | Throughput observed at the storage object. | [optional] [readonly] 

## Methods

### NewStorageBasePerformanceMetricsAverage

`func NewStorageBasePerformanceMetricsAverage(classId string, objectType string, ) *StorageBasePerformanceMetricsAverage`

NewStorageBasePerformanceMetricsAverage instantiates a new StorageBasePerformanceMetricsAverage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageBasePerformanceMetricsAverageWithDefaults

`func NewStorageBasePerformanceMetricsAverageWithDefaults() *StorageBasePerformanceMetricsAverage`

NewStorageBasePerformanceMetricsAverageWithDefaults instantiates a new StorageBasePerformanceMetricsAverage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StorageBasePerformanceMetricsAverage) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageBasePerformanceMetricsAverage) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageBasePerformanceMetricsAverage) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StorageBasePerformanceMetricsAverage) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageBasePerformanceMetricsAverage) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageBasePerformanceMetricsAverage) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetIops

`func (o *StorageBasePerformanceMetricsAverage) GetIops() float64`

GetIops returns the Iops field if non-nil, zero value otherwise.

### GetIopsOk

`func (o *StorageBasePerformanceMetricsAverage) GetIopsOk() (*float64, bool)`

GetIopsOk returns a tuple with the Iops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIops

`func (o *StorageBasePerformanceMetricsAverage) SetIops(v float64)`

SetIops sets Iops field to given value.

### HasIops

`func (o *StorageBasePerformanceMetricsAverage) HasIops() bool`

HasIops returns a boolean if a field has been set.

### GetLatency

`func (o *StorageBasePerformanceMetricsAverage) GetLatency() float64`

GetLatency returns the Latency field if non-nil, zero value otherwise.

### GetLatencyOk

`func (o *StorageBasePerformanceMetricsAverage) GetLatencyOk() (*float64, bool)`

GetLatencyOk returns a tuple with the Latency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatency

`func (o *StorageBasePerformanceMetricsAverage) SetLatency(v float64)`

SetLatency sets Latency field to given value.

### HasLatency

`func (o *StorageBasePerformanceMetricsAverage) HasLatency() bool`

HasLatency returns a boolean if a field has been set.

### GetPeriod

`func (o *StorageBasePerformanceMetricsAverage) GetPeriod() int64`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *StorageBasePerformanceMetricsAverage) GetPeriodOk() (*int64, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *StorageBasePerformanceMetricsAverage) SetPeriod(v int64)`

SetPeriod sets Period field to given value.

### HasPeriod

`func (o *StorageBasePerformanceMetricsAverage) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.

### GetThroughput

`func (o *StorageBasePerformanceMetricsAverage) GetThroughput() float64`

GetThroughput returns the Throughput field if non-nil, zero value otherwise.

### GetThroughputOk

`func (o *StorageBasePerformanceMetricsAverage) GetThroughputOk() (*float64, bool)`

GetThroughputOk returns a tuple with the Throughput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThroughput

`func (o *StorageBasePerformanceMetricsAverage) SetThroughput(v float64)`

SetThroughput sets Throughput field to given value.

### HasThroughput

`func (o *StorageBasePerformanceMetricsAverage) HasThroughput() bool`

HasThroughput returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


