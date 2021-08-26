# ServerConfigResultEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "server.ConfigResultEntry"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "server.ConfigResultEntry"]
**ConfigResult** | Pointer to [**ServerConfigResultRelationship**](ServerConfigResultRelationship.md) |  | [optional] 

## Methods

### NewServerConfigResultEntry

`func NewServerConfigResultEntry(classId string, objectType string, ) *ServerConfigResultEntry`

NewServerConfigResultEntry instantiates a new ServerConfigResultEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerConfigResultEntryWithDefaults

`func NewServerConfigResultEntryWithDefaults() *ServerConfigResultEntry`

NewServerConfigResultEntryWithDefaults instantiates a new ServerConfigResultEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ServerConfigResultEntry) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ServerConfigResultEntry) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ServerConfigResultEntry) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ServerConfigResultEntry) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ServerConfigResultEntry) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ServerConfigResultEntry) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


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


