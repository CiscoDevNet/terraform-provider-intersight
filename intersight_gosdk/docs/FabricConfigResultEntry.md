# FabricConfigResultEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "fabric.ConfigResultEntry"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "fabric.ConfigResultEntry"]
**ConfigResult** | Pointer to [**FabricConfigResultRelationship**](FabricConfigResultRelationship.md) |  | [optional] 

## Methods

### NewFabricConfigResultEntry

`func NewFabricConfigResultEntry(classId string, objectType string, ) *FabricConfigResultEntry`

NewFabricConfigResultEntry instantiates a new FabricConfigResultEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFabricConfigResultEntryWithDefaults

`func NewFabricConfigResultEntryWithDefaults() *FabricConfigResultEntry`

NewFabricConfigResultEntryWithDefaults instantiates a new FabricConfigResultEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *FabricConfigResultEntry) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *FabricConfigResultEntry) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *FabricConfigResultEntry) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *FabricConfigResultEntry) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *FabricConfigResultEntry) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *FabricConfigResultEntry) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetConfigResult

`func (o *FabricConfigResultEntry) GetConfigResult() FabricConfigResultRelationship`

GetConfigResult returns the ConfigResult field if non-nil, zero value otherwise.

### GetConfigResultOk

`func (o *FabricConfigResultEntry) GetConfigResultOk() (*FabricConfigResultRelationship, bool)`

GetConfigResultOk returns a tuple with the ConfigResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigResult

`func (o *FabricConfigResultEntry) SetConfigResult(v FabricConfigResultRelationship)`

SetConfigResult sets ConfigResult field to given value.

### HasConfigResult

`func (o *FabricConfigResultEntry) HasConfigResult() bool`

HasConfigResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


