# FabricConfigResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Profile** | Pointer to [**FabricSwitchProfileRelationship**](fabric.SwitchProfile.Relationship.md) |  | [optional] 
**ResultEntries** | Pointer to [**[]FabricConfigResultEntryRelationship**](fabric.ConfigResultEntry.Relationship.md) | An array of relationships to fabricConfigResultEntry resources. | [optional] 

## Methods

### NewFabricConfigResult

`func NewFabricConfigResult() *FabricConfigResult`

NewFabricConfigResult instantiates a new FabricConfigResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFabricConfigResultWithDefaults

`func NewFabricConfigResultWithDefaults() *FabricConfigResult`

NewFabricConfigResultWithDefaults instantiates a new FabricConfigResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProfile

`func (o *FabricConfigResult) GetProfile() FabricSwitchProfileRelationship`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *FabricConfigResult) GetProfileOk() (*FabricSwitchProfileRelationship, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *FabricConfigResult) SetProfile(v FabricSwitchProfileRelationship)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *FabricConfigResult) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### GetResultEntries

`func (o *FabricConfigResult) GetResultEntries() []FabricConfigResultEntryRelationship`

GetResultEntries returns the ResultEntries field if non-nil, zero value otherwise.

### GetResultEntriesOk

`func (o *FabricConfigResult) GetResultEntriesOk() (*[]FabricConfigResultEntryRelationship, bool)`

GetResultEntriesOk returns a tuple with the ResultEntries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultEntries

`func (o *FabricConfigResult) SetResultEntries(v []FabricConfigResultEntryRelationship)`

SetResultEntries sets ResultEntries field to given value.

### HasResultEntries

`func (o *FabricConfigResult) HasResultEntries() bool`

HasResultEntries returns a boolean if a field has been set.

### SetResultEntriesNil

`func (o *FabricConfigResult) SetResultEntriesNil(b bool)`

 SetResultEntriesNil sets the value for ResultEntries to be an explicit nil

### UnsetResultEntries
`func (o *FabricConfigResult) UnsetResultEntries()`

UnsetResultEntries ensures that no value is present for ResultEntries, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


