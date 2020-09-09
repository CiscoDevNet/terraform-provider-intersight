# HyperflexConfigResultAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfigProgress** | Pointer to **string** | The progress percentage of the running configuration or workflow. | [optional] 
**Duration** | Pointer to **string** | The duration of the running configuration or workflow. | [optional] 
**StartTime** | Pointer to **string** | The start time of the configuration or workflow. | [optional] 
**ClusterProfile** | Pointer to [**HyperflexClusterProfileRelationship**](hyperflex.ClusterProfile.Relationship.md) |  | [optional] 
**ResultEntries** | Pointer to [**[]HyperflexConfigResultEntryRelationship**](hyperflex.ConfigResultEntry.Relationship.md) | An array of relationships to hyperflexConfigResultEntry resources. | [optional] 

## Methods

### NewHyperflexConfigResultAllOf

`func NewHyperflexConfigResultAllOf() *HyperflexConfigResultAllOf`

NewHyperflexConfigResultAllOf instantiates a new HyperflexConfigResultAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexConfigResultAllOfWithDefaults

`func NewHyperflexConfigResultAllOfWithDefaults() *HyperflexConfigResultAllOf`

NewHyperflexConfigResultAllOfWithDefaults instantiates a new HyperflexConfigResultAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfigProgress

`func (o *HyperflexConfigResultAllOf) GetConfigProgress() string`

GetConfigProgress returns the ConfigProgress field if non-nil, zero value otherwise.

### GetConfigProgressOk

`func (o *HyperflexConfigResultAllOf) GetConfigProgressOk() (*string, bool)`

GetConfigProgressOk returns a tuple with the ConfigProgress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigProgress

`func (o *HyperflexConfigResultAllOf) SetConfigProgress(v string)`

SetConfigProgress sets ConfigProgress field to given value.

### HasConfigProgress

`func (o *HyperflexConfigResultAllOf) HasConfigProgress() bool`

HasConfigProgress returns a boolean if a field has been set.

### GetDuration

`func (o *HyperflexConfigResultAllOf) GetDuration() string`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *HyperflexConfigResultAllOf) GetDurationOk() (*string, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *HyperflexConfigResultAllOf) SetDuration(v string)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *HyperflexConfigResultAllOf) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetStartTime

`func (o *HyperflexConfigResultAllOf) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *HyperflexConfigResultAllOf) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *HyperflexConfigResultAllOf) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *HyperflexConfigResultAllOf) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetClusterProfile

`func (o *HyperflexConfigResultAllOf) GetClusterProfile() HyperflexClusterProfileRelationship`

GetClusterProfile returns the ClusterProfile field if non-nil, zero value otherwise.

### GetClusterProfileOk

`func (o *HyperflexConfigResultAllOf) GetClusterProfileOk() (*HyperflexClusterProfileRelationship, bool)`

GetClusterProfileOk returns a tuple with the ClusterProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterProfile

`func (o *HyperflexConfigResultAllOf) SetClusterProfile(v HyperflexClusterProfileRelationship)`

SetClusterProfile sets ClusterProfile field to given value.

### HasClusterProfile

`func (o *HyperflexConfigResultAllOf) HasClusterProfile() bool`

HasClusterProfile returns a boolean if a field has been set.

### GetResultEntries

`func (o *HyperflexConfigResultAllOf) GetResultEntries() []HyperflexConfigResultEntryRelationship`

GetResultEntries returns the ResultEntries field if non-nil, zero value otherwise.

### GetResultEntriesOk

`func (o *HyperflexConfigResultAllOf) GetResultEntriesOk() (*[]HyperflexConfigResultEntryRelationship, bool)`

GetResultEntriesOk returns a tuple with the ResultEntries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultEntries

`func (o *HyperflexConfigResultAllOf) SetResultEntries(v []HyperflexConfigResultEntryRelationship)`

SetResultEntries sets ResultEntries field to given value.

### HasResultEntries

`func (o *HyperflexConfigResultAllOf) HasResultEntries() bool`

HasResultEntries returns a boolean if a field has been set.

### SetResultEntriesNil

`func (o *HyperflexConfigResultAllOf) SetResultEntriesNil(b bool)`

 SetResultEntriesNil sets the value for ResultEntries to be an explicit nil

### UnsetResultEntries
`func (o *HyperflexConfigResultAllOf) UnsetResultEntries()`

UnsetResultEntries ensures that no value is present for ResultEntries, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


