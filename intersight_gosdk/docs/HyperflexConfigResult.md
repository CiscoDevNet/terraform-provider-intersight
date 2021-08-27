# HyperflexConfigResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hyperflex.ConfigResult"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hyperflex.ConfigResult"]
**ConfigProgress** | Pointer to **string** | The progress percentage of the running configuration or workflow. | [optional] 
**Duration** | Pointer to **string** | The duration of the running configuration or workflow. | [optional] 
**StartTime** | Pointer to **string** | The start time of the configuration or workflow. | [optional] 
**ClusterProfile** | Pointer to [**HyperflexClusterProfileRelationship**](HyperflexClusterProfileRelationship.md) |  | [optional] 
**ResultEntries** | Pointer to [**[]HyperflexConfigResultEntryRelationship**](HyperflexConfigResultEntryRelationship.md) | An array of relationships to hyperflexConfigResultEntry resources. | [optional] 

## Methods

### NewHyperflexConfigResult

`func NewHyperflexConfigResult(classId string, objectType string, ) *HyperflexConfigResult`

NewHyperflexConfigResult instantiates a new HyperflexConfigResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexConfigResultWithDefaults

`func NewHyperflexConfigResultWithDefaults() *HyperflexConfigResult`

NewHyperflexConfigResultWithDefaults instantiates a new HyperflexConfigResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HyperflexConfigResult) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HyperflexConfigResult) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HyperflexConfigResult) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HyperflexConfigResult) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HyperflexConfigResult) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HyperflexConfigResult) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetConfigProgress

`func (o *HyperflexConfigResult) GetConfigProgress() string`

GetConfigProgress returns the ConfigProgress field if non-nil, zero value otherwise.

### GetConfigProgressOk

`func (o *HyperflexConfigResult) GetConfigProgressOk() (*string, bool)`

GetConfigProgressOk returns a tuple with the ConfigProgress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigProgress

`func (o *HyperflexConfigResult) SetConfigProgress(v string)`

SetConfigProgress sets ConfigProgress field to given value.

### HasConfigProgress

`func (o *HyperflexConfigResult) HasConfigProgress() bool`

HasConfigProgress returns a boolean if a field has been set.

### GetDuration

`func (o *HyperflexConfigResult) GetDuration() string`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *HyperflexConfigResult) GetDurationOk() (*string, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *HyperflexConfigResult) SetDuration(v string)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *HyperflexConfigResult) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetStartTime

`func (o *HyperflexConfigResult) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *HyperflexConfigResult) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *HyperflexConfigResult) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *HyperflexConfigResult) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetClusterProfile

`func (o *HyperflexConfigResult) GetClusterProfile() HyperflexClusterProfileRelationship`

GetClusterProfile returns the ClusterProfile field if non-nil, zero value otherwise.

### GetClusterProfileOk

`func (o *HyperflexConfigResult) GetClusterProfileOk() (*HyperflexClusterProfileRelationship, bool)`

GetClusterProfileOk returns a tuple with the ClusterProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterProfile

`func (o *HyperflexConfigResult) SetClusterProfile(v HyperflexClusterProfileRelationship)`

SetClusterProfile sets ClusterProfile field to given value.

### HasClusterProfile

`func (o *HyperflexConfigResult) HasClusterProfile() bool`

HasClusterProfile returns a boolean if a field has been set.

### GetResultEntries

`func (o *HyperflexConfigResult) GetResultEntries() []HyperflexConfigResultEntryRelationship`

GetResultEntries returns the ResultEntries field if non-nil, zero value otherwise.

### GetResultEntriesOk

`func (o *HyperflexConfigResult) GetResultEntriesOk() (*[]HyperflexConfigResultEntryRelationship, bool)`

GetResultEntriesOk returns a tuple with the ResultEntries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultEntries

`func (o *HyperflexConfigResult) SetResultEntries(v []HyperflexConfigResultEntryRelationship)`

SetResultEntries sets ResultEntries field to given value.

### HasResultEntries

`func (o *HyperflexConfigResult) HasResultEntries() bool`

HasResultEntries returns a boolean if a field has been set.

### SetResultEntriesNil

`func (o *HyperflexConfigResult) SetResultEntriesNil(b bool)`

 SetResultEntriesNil sets the value for ResultEntries to be an explicit nil

### UnsetResultEntries
`func (o *HyperflexConfigResult) UnsetResultEntries()`

UnsetResultEntries ensures that no value is present for ResultEntries, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


