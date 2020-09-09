# ServerConfigResultAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Profile** | Pointer to [**ServerProfileRelationship**](server.Profile.Relationship.md) |  | [optional] 
**ResultEntries** | Pointer to [**[]ServerConfigResultEntryRelationship**](server.ConfigResultEntry.Relationship.md) | An array of relationships to serverConfigResultEntry resources. | [optional] 

## Methods

### NewServerConfigResultAllOf

`func NewServerConfigResultAllOf() *ServerConfigResultAllOf`

NewServerConfigResultAllOf instantiates a new ServerConfigResultAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerConfigResultAllOfWithDefaults

`func NewServerConfigResultAllOfWithDefaults() *ServerConfigResultAllOf`

NewServerConfigResultAllOfWithDefaults instantiates a new ServerConfigResultAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProfile

`func (o *ServerConfigResultAllOf) GetProfile() ServerProfileRelationship`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *ServerConfigResultAllOf) GetProfileOk() (*ServerProfileRelationship, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *ServerConfigResultAllOf) SetProfile(v ServerProfileRelationship)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *ServerConfigResultAllOf) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### GetResultEntries

`func (o *ServerConfigResultAllOf) GetResultEntries() []ServerConfigResultEntryRelationship`

GetResultEntries returns the ResultEntries field if non-nil, zero value otherwise.

### GetResultEntriesOk

`func (o *ServerConfigResultAllOf) GetResultEntriesOk() (*[]ServerConfigResultEntryRelationship, bool)`

GetResultEntriesOk returns a tuple with the ResultEntries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultEntries

`func (o *ServerConfigResultAllOf) SetResultEntries(v []ServerConfigResultEntryRelationship)`

SetResultEntries sets ResultEntries field to given value.

### HasResultEntries

`func (o *ServerConfigResultAllOf) HasResultEntries() bool`

HasResultEntries returns a boolean if a field has been set.

### SetResultEntriesNil

`func (o *ServerConfigResultAllOf) SetResultEntriesNil(b bool)`

 SetResultEntriesNil sets the value for ResultEntries to be an explicit nil

### UnsetResultEntries
`func (o *ServerConfigResultAllOf) UnsetResultEntries()`

UnsetResultEntries ensures that no value is present for ResultEntries, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


