# StorageNetAppNode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "storage.NetAppNode"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "storage.NetAppNode"]
**AvgPerformanceMetrics** | Pointer to [**NullableStorageBasePerformanceMetricsAverage**](StorageBasePerformanceMetricsAverage.md) | Average performance metrics data for a NetApp storage resource over a given period of time. | [optional] 
**CdpdEnabled** | Pointer to **string** | Storage node option for cdpd state. * &#x60;unknown&#x60; - The cdpd option is unknown on the node. * &#x60;on&#x60; - The cdpd option is enabled on the node. * &#x60;off&#x60; - The cdpd option is disabled on the node. | [optional] [readonly] [default to "unknown"]
**Health** | Pointer to **bool** | The health of the NetApp Node. | [optional] [readonly] 
**HighAvailability** | Pointer to [**NullableStorageNetAppHighAvailability**](StorageNetAppHighAvailability.md) |  | [optional] 
**Key** | Pointer to **string** | Unique identifier of NetApp Node across data center. | [optional] [readonly] 
**State** | Pointer to **string** | The state of the NetApp Node. | [optional] [readonly] 
**Systemid** | Pointer to **string** | The system id of the NetApp Node. | [optional] [readonly] 
**Uuid** | Pointer to **string** | Universally unique identifier of NetApp Node. | [optional] [readonly] 
**Array** | Pointer to [**NullableStorageNetAppClusterRelationship**](StorageNetAppClusterRelationship.md) |  | [optional] 
**Events** | Pointer to [**[]StorageNetAppNodeEventRelationship**](StorageNetAppNodeEventRelationship.md) | An array of relationships to storageNetAppNodeEvent resources. | [optional] [readonly] 

## Methods

### NewStorageNetAppNode

`func NewStorageNetAppNode(classId string, objectType string, ) *StorageNetAppNode`

NewStorageNetAppNode instantiates a new StorageNetAppNode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageNetAppNodeWithDefaults

`func NewStorageNetAppNodeWithDefaults() *StorageNetAppNode`

NewStorageNetAppNodeWithDefaults instantiates a new StorageNetAppNode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StorageNetAppNode) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageNetAppNode) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageNetAppNode) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StorageNetAppNode) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageNetAppNode) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageNetAppNode) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAvgPerformanceMetrics

`func (o *StorageNetAppNode) GetAvgPerformanceMetrics() StorageBasePerformanceMetricsAverage`

GetAvgPerformanceMetrics returns the AvgPerformanceMetrics field if non-nil, zero value otherwise.

### GetAvgPerformanceMetricsOk

`func (o *StorageNetAppNode) GetAvgPerformanceMetricsOk() (*StorageBasePerformanceMetricsAverage, bool)`

GetAvgPerformanceMetricsOk returns a tuple with the AvgPerformanceMetrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgPerformanceMetrics

`func (o *StorageNetAppNode) SetAvgPerformanceMetrics(v StorageBasePerformanceMetricsAverage)`

SetAvgPerformanceMetrics sets AvgPerformanceMetrics field to given value.

### HasAvgPerformanceMetrics

`func (o *StorageNetAppNode) HasAvgPerformanceMetrics() bool`

HasAvgPerformanceMetrics returns a boolean if a field has been set.

### SetAvgPerformanceMetricsNil

`func (o *StorageNetAppNode) SetAvgPerformanceMetricsNil(b bool)`

 SetAvgPerformanceMetricsNil sets the value for AvgPerformanceMetrics to be an explicit nil

### UnsetAvgPerformanceMetrics
`func (o *StorageNetAppNode) UnsetAvgPerformanceMetrics()`

UnsetAvgPerformanceMetrics ensures that no value is present for AvgPerformanceMetrics, not even an explicit nil
### GetCdpdEnabled

`func (o *StorageNetAppNode) GetCdpdEnabled() string`

GetCdpdEnabled returns the CdpdEnabled field if non-nil, zero value otherwise.

### GetCdpdEnabledOk

`func (o *StorageNetAppNode) GetCdpdEnabledOk() (*string, bool)`

GetCdpdEnabledOk returns a tuple with the CdpdEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCdpdEnabled

`func (o *StorageNetAppNode) SetCdpdEnabled(v string)`

SetCdpdEnabled sets CdpdEnabled field to given value.

### HasCdpdEnabled

`func (o *StorageNetAppNode) HasCdpdEnabled() bool`

HasCdpdEnabled returns a boolean if a field has been set.

### GetHealth

`func (o *StorageNetAppNode) GetHealth() bool`

GetHealth returns the Health field if non-nil, zero value otherwise.

### GetHealthOk

`func (o *StorageNetAppNode) GetHealthOk() (*bool, bool)`

GetHealthOk returns a tuple with the Health field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealth

`func (o *StorageNetAppNode) SetHealth(v bool)`

SetHealth sets Health field to given value.

### HasHealth

`func (o *StorageNetAppNode) HasHealth() bool`

HasHealth returns a boolean if a field has been set.

### GetHighAvailability

`func (o *StorageNetAppNode) GetHighAvailability() StorageNetAppHighAvailability`

GetHighAvailability returns the HighAvailability field if non-nil, zero value otherwise.

### GetHighAvailabilityOk

`func (o *StorageNetAppNode) GetHighAvailabilityOk() (*StorageNetAppHighAvailability, bool)`

GetHighAvailabilityOk returns a tuple with the HighAvailability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighAvailability

`func (o *StorageNetAppNode) SetHighAvailability(v StorageNetAppHighAvailability)`

SetHighAvailability sets HighAvailability field to given value.

### HasHighAvailability

`func (o *StorageNetAppNode) HasHighAvailability() bool`

HasHighAvailability returns a boolean if a field has been set.

### SetHighAvailabilityNil

`func (o *StorageNetAppNode) SetHighAvailabilityNil(b bool)`

 SetHighAvailabilityNil sets the value for HighAvailability to be an explicit nil

### UnsetHighAvailability
`func (o *StorageNetAppNode) UnsetHighAvailability()`

UnsetHighAvailability ensures that no value is present for HighAvailability, not even an explicit nil
### GetKey

`func (o *StorageNetAppNode) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *StorageNetAppNode) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *StorageNetAppNode) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *StorageNetAppNode) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetState

`func (o *StorageNetAppNode) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *StorageNetAppNode) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *StorageNetAppNode) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *StorageNetAppNode) HasState() bool`

HasState returns a boolean if a field has been set.

### GetSystemid

`func (o *StorageNetAppNode) GetSystemid() string`

GetSystemid returns the Systemid field if non-nil, zero value otherwise.

### GetSystemidOk

`func (o *StorageNetAppNode) GetSystemidOk() (*string, bool)`

GetSystemidOk returns a tuple with the Systemid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemid

`func (o *StorageNetAppNode) SetSystemid(v string)`

SetSystemid sets Systemid field to given value.

### HasSystemid

`func (o *StorageNetAppNode) HasSystemid() bool`

HasSystemid returns a boolean if a field has been set.

### GetUuid

`func (o *StorageNetAppNode) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *StorageNetAppNode) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *StorageNetAppNode) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *StorageNetAppNode) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetArray

`func (o *StorageNetAppNode) GetArray() StorageNetAppClusterRelationship`

GetArray returns the Array field if non-nil, zero value otherwise.

### GetArrayOk

`func (o *StorageNetAppNode) GetArrayOk() (*StorageNetAppClusterRelationship, bool)`

GetArrayOk returns a tuple with the Array field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArray

`func (o *StorageNetAppNode) SetArray(v StorageNetAppClusterRelationship)`

SetArray sets Array field to given value.

### HasArray

`func (o *StorageNetAppNode) HasArray() bool`

HasArray returns a boolean if a field has been set.

### SetArrayNil

`func (o *StorageNetAppNode) SetArrayNil(b bool)`

 SetArrayNil sets the value for Array to be an explicit nil

### UnsetArray
`func (o *StorageNetAppNode) UnsetArray()`

UnsetArray ensures that no value is present for Array, not even an explicit nil
### GetEvents

`func (o *StorageNetAppNode) GetEvents() []StorageNetAppNodeEventRelationship`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *StorageNetAppNode) GetEventsOk() (*[]StorageNetAppNodeEventRelationship, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *StorageNetAppNode) SetEvents(v []StorageNetAppNodeEventRelationship)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *StorageNetAppNode) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### SetEventsNil

`func (o *StorageNetAppNode) SetEventsNil(b bool)`

 SetEventsNil sets the value for Events to be an explicit nil

### UnsetEvents
`func (o *StorageNetAppNode) UnsetEvents()`

UnsetEvents ensures that no value is present for Events, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


