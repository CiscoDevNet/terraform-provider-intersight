# CapabilityServerModelsCapabilityDef

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "capability.ServerModelsCapabilityDef"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "capability.ServerModelsCapabilityDef"]
**Models** | Pointer to **[]string** |  | [optional] 
**ServerType** | Pointer to **string** | Type of the server. Example, BladeM6, RackM5. | [optional] 

## Methods

### NewCapabilityServerModelsCapabilityDef

`func NewCapabilityServerModelsCapabilityDef(classId string, objectType string, ) *CapabilityServerModelsCapabilityDef`

NewCapabilityServerModelsCapabilityDef instantiates a new CapabilityServerModelsCapabilityDef object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCapabilityServerModelsCapabilityDefWithDefaults

`func NewCapabilityServerModelsCapabilityDefWithDefaults() *CapabilityServerModelsCapabilityDef`

NewCapabilityServerModelsCapabilityDefWithDefaults instantiates a new CapabilityServerModelsCapabilityDef object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *CapabilityServerModelsCapabilityDef) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *CapabilityServerModelsCapabilityDef) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *CapabilityServerModelsCapabilityDef) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *CapabilityServerModelsCapabilityDef) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CapabilityServerModelsCapabilityDef) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CapabilityServerModelsCapabilityDef) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetModels

`func (o *CapabilityServerModelsCapabilityDef) GetModels() []string`

GetModels returns the Models field if non-nil, zero value otherwise.

### GetModelsOk

`func (o *CapabilityServerModelsCapabilityDef) GetModelsOk() (*[]string, bool)`

GetModelsOk returns a tuple with the Models field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModels

`func (o *CapabilityServerModelsCapabilityDef) SetModels(v []string)`

SetModels sets Models field to given value.

### HasModels

`func (o *CapabilityServerModelsCapabilityDef) HasModels() bool`

HasModels returns a boolean if a field has been set.

### SetModelsNil

`func (o *CapabilityServerModelsCapabilityDef) SetModelsNil(b bool)`

 SetModelsNil sets the value for Models to be an explicit nil

### UnsetModels
`func (o *CapabilityServerModelsCapabilityDef) UnsetModels()`

UnsetModels ensures that no value is present for Models, not even an explicit nil
### GetServerType

`func (o *CapabilityServerModelsCapabilityDef) GetServerType() string`

GetServerType returns the ServerType field if non-nil, zero value otherwise.

### GetServerTypeOk

`func (o *CapabilityServerModelsCapabilityDef) GetServerTypeOk() (*string, bool)`

GetServerTypeOk returns a tuple with the ServerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerType

`func (o *CapabilityServerModelsCapabilityDef) SetServerType(v string)`

SetServerType sets ServerType field to given value.

### HasServerType

`func (o *CapabilityServerModelsCapabilityDef) HasServerType() bool`

HasServerType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


