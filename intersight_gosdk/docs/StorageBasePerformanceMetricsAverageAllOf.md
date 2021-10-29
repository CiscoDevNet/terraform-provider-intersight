# StorageBasePerformanceMetricsAverageAllOf

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

### NewStorageBasePerformanceMetricsAverageAllOf

`func NewStorageBasePerformanceMetricsAverageAllOf(classId string, objectType string, ) *StorageBasePerformanceMetricsAverageAllOf`

NewStorageBasePerformanceMetricsAverageAllOf instantiates a new StorageBasePerformanceMetricsAverageAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageBasePerformanceMetricsAverageAllOfWithDefaults

`func NewStorageBasePerformanceMetricsAverageAllOfWithDefaults() *StorageBasePerformanceMetricsAverageAllOf`

NewStorageBasePerformanceMetricsAverageAllOfWithDefaults instantiates a new StorageBasePerformanceMetricsAverageAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StorageBasePerformanceMetricsAverageAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageBasePerformanceMetricsAverageAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageBasePerformanceMetricsAverageAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StorageBasePerformanceMetricsAverageAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageBasePerformanceMetricsAverageAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageBasePerformanceMetricsAverageAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetIops

`func (o *StorageBasePerformanceMetricsAverageAllOf) GetIops() float64`

GetIops returns the Iops field if non-nil, zero value otherwise.

### GetIopsOk

`func (o *StorageBasePerformanceMetricsAverageAllOf) GetIopsOk() (*float64, bool)`

GetIopsOk returns a tuple with the Iops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIops

`func (o *StorageBasePerformanceMetricsAverageAllOf) SetIops(v float64)`

SetIops sets Iops field to given value.

### HasIops

`func (o *StorageBasePerformanceMetricsAverageAllOf) HasIops() bool`

HasIops returns a boolean if a field has been set.

### GetLatency

`func (o *StorageBasePerformanceMetricsAverageAllOf) GetLatency() float64`

GetLatency returns the Latency field if non-nil, zero value otherwise.

### GetLatencyOk

`func (o *StorageBasePerformanceMetricsAverageAllOf) GetLatencyOk() (*float64, bool)`

GetLatencyOk returns a tuple with the Latency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatency

`func (o *StorageBasePerformanceMetricsAverageAllOf) SetLatency(v float64)`

SetLatency sets Latency field to given value.

### HasLatency

`func (o *StorageBasePerformanceMetricsAverageAllOf) HasLatency() bool`

HasLatency returns a boolean if a field has been set.

### GetPeriod

`func (o *StorageBasePerformanceMetricsAverageAllOf) GetPeriod() int64`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *StorageBasePerformanceMetricsAverageAllOf) GetPeriodOk() (*int64, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *StorageBasePerformanceMetricsAverageAllOf) SetPeriod(v int64)`

SetPeriod sets Period field to given value.

### HasPeriod

`func (o *StorageBasePerformanceMetricsAverageAllOf) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.

### GetThroughput

`func (o *StorageBasePerformanceMetricsAverageAllOf) GetThroughput() float64`

GetThroughput returns the Throughput field if non-nil, zero value otherwise.

### GetThroughputOk

`func (o *StorageBasePerformanceMetricsAverageAllOf) GetThroughputOk() (*float64, bool)`

GetThroughputOk returns a tuple with the Throughput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThroughput

`func (o *StorageBasePerformanceMetricsAverageAllOf) SetThroughput(v float64)`

SetThroughput sets Throughput field to given value.

### HasThroughput

`func (o *StorageBasePerformanceMetricsAverageAllOf) HasThroughput() bool`

HasThroughput returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


