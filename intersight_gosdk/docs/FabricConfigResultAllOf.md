# FabricConfigResultAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Profile** | Pointer to [**FabricSwitchProfileRelationship**](fabric.SwitchProfile.Relationship.md) |  | [optional] 
**ResultEntries** | Pointer to [**[]FabricConfigResultEntryRelationship**](fabric.ConfigResultEntry.Relationship.md) | An array of relationships to fabricConfigResultEntry resources. | [optional] 

## Methods

### NewFabricConfigResultAllOf

`func NewFabricConfigResultAllOf() *FabricConfigResultAllOf`

NewFabricConfigResultAllOf instantiates a new FabricConfigResultAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFabricConfigResultAllOfWithDefaults

`func NewFabricConfigResultAllOfWithDefaults() *FabricConfigResultAllOf`

NewFabricConfigResultAllOfWithDefaults instantiates a new FabricConfigResultAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProfile

`func (o *FabricConfigResultAllOf) GetProfile() FabricSwitchProfileRelationship`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *FabricConfigResultAllOf) GetProfileOk() (*FabricSwitchProfileRelationship, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *FabricConfigResultAllOf) SetProfile(v FabricSwitchProfileRelationship)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *FabricConfigResultAllOf) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### GetResultEntries

`func (o *FabricConfigResultAllOf) GetResultEntries() []FabricConfigResultEntryRelationship`

GetResultEntries returns the ResultEntries field if non-nil, zero value otherwise.

### GetResultEntriesOk

`func (o *FabricConfigResultAllOf) GetResultEntriesOk() (*[]FabricConfigResultEntryRelationship, bool)`

GetResultEntriesOk returns a tuple with the ResultEntries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultEntries

`func (o *FabricConfigResultAllOf) SetResultEntries(v []FabricConfigResultEntryRelationship)`

SetResultEntries sets ResultEntries field to given value.

### HasResultEntries

`func (o *FabricConfigResultAllOf) HasResultEntries() bool`

HasResultEntries returns a boolean if a field has been set.

### SetResultEntriesNil

`func (o *FabricConfigResultAllOf) SetResultEntriesNil(b bool)`

 SetResultEntriesNil sets the value for ResultEntries to be an explicit nil

### UnsetResultEntries
`func (o *FabricConfigResultAllOf) UnsetResultEntries()`

UnsetResultEntries ensures that no value is present for ResultEntries, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


