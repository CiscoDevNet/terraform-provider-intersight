# ServerConfigResultEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfigResult** | Pointer to [**ServerConfigResultRelationship**](server.ConfigResult.Relationship.md) |  | [optional] 

## Methods

### NewServerConfigResultEntry

`func NewServerConfigResultEntry() *ServerConfigResultEntry`

NewServerConfigResultEntry instantiates a new ServerConfigResultEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerConfigResultEntryWithDefaults

`func NewServerConfigResultEntryWithDefaults() *ServerConfigResultEntry`

NewServerConfigResultEntryWithDefaults instantiates a new ServerConfigResultEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfigResult

`func (o *ServerConfigResultEntry) GetConfigResult() ServerConfigResultRelationship`

GetConfigResult returns the ConfigResult field if non-nil, zero value otherwise.

### GetConfigResultOk

`func (o *ServerConfigResultEntry) GetConfigResultOk() (*ServerConfigResultRelationship, bool)`

GetConfigResultOk returns a tuple with the ConfigResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigResult

`func (o *ServerConfigResultEntry) SetConfigResult(v ServerConfigResultRelationship)`

SetConfigResult sets ConfigResult field to given value.

### HasConfigResult

`func (o *ServerConfigResultEntry) HasConfigResult() bool`

HasConfigResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


