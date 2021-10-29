# StorageNetAppNodeAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "storage.NetAppNode"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "storage.NetAppNode"]
**AvgPerformanceMetrics** | Pointer to [**StorageNetAppPerformanceMetricsAverage**](StorageNetAppPerformanceMetricsAverage.md) |  | [optional] 
**Health** | Pointer to **bool** | The health of the NetApp Node. | [optional] [readonly] 
**HighAvailability** | Pointer to [**NullableStorageNetAppHighAvailability**](StorageNetAppHighAvailability.md) |  | [optional] 
**Key** | Pointer to **string** | Unique identifier of NetApp Node across data center. | [optional] [readonly] 
**Systemid** | Pointer to **string** | The system id of the NetApp Node. | [optional] [readonly] 
**Uuid** | Pointer to **string** | Universally unique identifier of NetApp Node. | [optional] [readonly] 
**Array** | Pointer to [**StorageNetAppClusterRelationship**](StorageNetAppClusterRelationship.md) |  | [optional] 

## Methods

### NewStorageNetAppNodeAllOf

`func NewStorageNetAppNodeAllOf(classId string, objectType string, ) *StorageNetAppNodeAllOf`

NewStorageNetAppNodeAllOf instantiates a new StorageNetAppNodeAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageNetAppNodeAllOfWithDefaults

`func NewStorageNetAppNodeAllOfWithDefaults() *StorageNetAppNodeAllOf`

NewStorageNetAppNodeAllOfWithDefaults instantiates a new StorageNetAppNodeAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StorageNetAppNodeAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageNetAppNodeAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageNetAppNodeAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StorageNetAppNodeAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageNetAppNodeAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageNetAppNodeAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAvgPerformanceMetrics

`func (o *StorageNetAppNodeAllOf) GetAvgPerformanceMetrics() StorageNetAppPerformanceMetricsAverage`

GetAvgPerformanceMetrics returns the AvgPerformanceMetrics field if non-nil, zero value otherwise.

### GetAvgPerformanceMetricsOk

`func (o *StorageNetAppNodeAllOf) GetAvgPerformanceMetricsOk() (*StorageNetAppPerformanceMetricsAverage, bool)`

GetAvgPerformanceMetricsOk returns a tuple with the AvgPerformanceMetrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgPerformanceMetrics

`func (o *StorageNetAppNodeAllOf) SetAvgPerformanceMetrics(v StorageNetAppPerformanceMetricsAverage)`

SetAvgPerformanceMetrics sets AvgPerformanceMetrics field to given value.

### HasAvgPerformanceMetrics

`func (o *StorageNetAppNodeAllOf) HasAvgPerformanceMetrics() bool`

HasAvgPerformanceMetrics returns a boolean if a field has been set.

### GetHealth

`func (o *StorageNetAppNodeAllOf) GetHealth() bool`

GetHealth returns the Health field if non-nil, zero value otherwise.

### GetHealthOk

`func (o *StorageNetAppNodeAllOf) GetHealthOk() (*bool, bool)`

GetHealthOk returns a tuple with the Health field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealth

`func (o *StorageNetAppNodeAllOf) SetHealth(v bool)`

SetHealth sets Health field to given value.

### HasHealth

`func (o *StorageNetAppNodeAllOf) HasHealth() bool`

HasHealth returns a boolean if a field has been set.

### GetHighAvailability

`func (o *StorageNetAppNodeAllOf) GetHighAvailability() StorageNetAppHighAvailability`

GetHighAvailability returns the HighAvailability field if non-nil, zero value otherwise.

### GetHighAvailabilityOk

`func (o *StorageNetAppNodeAllOf) GetHighAvailabilityOk() (*StorageNetAppHighAvailability, bool)`

GetHighAvailabilityOk returns a tuple with the HighAvailability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighAvailability

`func (o *StorageNetAppNodeAllOf) SetHighAvailability(v StorageNetAppHighAvailability)`

SetHighAvailability sets HighAvailability field to given value.

### HasHighAvailability

`func (o *StorageNetAppNodeAllOf) HasHighAvailability() bool`

HasHighAvailability returns a boolean if a field has been set.

### SetHighAvailabilityNil

`func (o *StorageNetAppNodeAllOf) SetHighAvailabilityNil(b bool)`

 SetHighAvailabilityNil sets the value for HighAvailability to be an explicit nil

### UnsetHighAvailability
`func (o *StorageNetAppNodeAllOf) UnsetHighAvailability()`

UnsetHighAvailability ensures that no value is present for HighAvailability, not even an explicit nil
### GetKey

`func (o *StorageNetAppNodeAllOf) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *StorageNetAppNodeAllOf) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *StorageNetAppNodeAllOf) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *StorageNetAppNodeAllOf) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetSystemid

`func (o *StorageNetAppNodeAllOf) GetSystemid() string`

GetSystemid returns the Systemid field if non-nil, zero value otherwise.

### GetSystemidOk

`func (o *StorageNetAppNodeAllOf) GetSystemidOk() (*string, bool)`

GetSystemidOk returns a tuple with the Systemid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemid

`func (o *StorageNetAppNodeAllOf) SetSystemid(v string)`

SetSystemid sets Systemid field to given value.

### HasSystemid

`func (o *StorageNetAppNodeAllOf) HasSystemid() bool`

HasSystemid returns a boolean if a field has been set.

### GetUuid

`func (o *StorageNetAppNodeAllOf) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *StorageNetAppNodeAllOf) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *StorageNetAppNodeAllOf) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *StorageNetAppNodeAllOf) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetArray

`func (o *StorageNetAppNodeAllOf) GetArray() StorageNetAppClusterRelationship`

GetArray returns the Array field if non-nil, zero value otherwise.

### GetArrayOk

`func (o *StorageNetAppNodeAllOf) GetArrayOk() (*StorageNetAppClusterRelationship, bool)`

GetArrayOk returns a tuple with the Array field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArray

`func (o *StorageNetAppNodeAllOf) SetArray(v StorageNetAppClusterRelationship)`

SetArray sets Array field to given value.

### HasArray

`func (o *StorageNetAppNodeAllOf) HasArray() bool`

HasArray returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


