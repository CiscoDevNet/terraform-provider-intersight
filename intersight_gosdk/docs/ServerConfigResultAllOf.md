# ServerConfigResultAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "server.ConfigResult"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "server.ConfigResult"]
**Profile** | Pointer to [**ServerProfileRelationship**](ServerProfileRelationship.md) |  | [optional] 
**ResultEntries** | Pointer to [**[]ServerConfigResultEntryRelationship**](ServerConfigResultEntryRelationship.md) | An array of relationships to serverConfigResultEntry resources. | [optional] 

## Methods

### NewServerConfigResultAllOf

`func NewServerConfigResultAllOf(classId string, objectType string, ) *ServerConfigResultAllOf`

NewServerConfigResultAllOf instantiates a new ServerConfigResultAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerConfigResultAllOfWithDefaults

`func NewServerConfigResultAllOfWithDefaults() *ServerConfigResultAllOf`

NewServerConfigResultAllOfWithDefaults instantiates a new ServerConfigResultAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ServerConfigResultAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ServerConfigResultAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ServerConfigResultAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ServerConfigResultAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ServerConfigResultAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ServerConfigResultAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


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


