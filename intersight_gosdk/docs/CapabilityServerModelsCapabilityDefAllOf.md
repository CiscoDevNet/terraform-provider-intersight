# CapabilityServerModelsCapabilityDefAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "capability.ServerModelsCapabilityDef"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "capability.ServerModelsCapabilityDef"]
**Models** | Pointer to **[]string** |  | [optional] 
**ServerType** | Pointer to **string** | Type of the server. Example, BladeM6, RackM5. | [optional] 

## Methods

### NewCapabilityServerModelsCapabilityDefAllOf

`func NewCapabilityServerModelsCapabilityDefAllOf(classId string, objectType string, ) *CapabilityServerModelsCapabilityDefAllOf`

NewCapabilityServerModelsCapabilityDefAllOf instantiates a new CapabilityServerModelsCapabilityDefAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCapabilityServerModelsCapabilityDefAllOfWithDefaults

`func NewCapabilityServerModelsCapabilityDefAllOfWithDefaults() *CapabilityServerModelsCapabilityDefAllOf`

NewCapabilityServerModelsCapabilityDefAllOfWithDefaults instantiates a new CapabilityServerModelsCapabilityDefAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *CapabilityServerModelsCapabilityDefAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *CapabilityServerModelsCapabilityDefAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *CapabilityServerModelsCapabilityDefAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *CapabilityServerModelsCapabilityDefAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CapabilityServerModelsCapabilityDefAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CapabilityServerModelsCapabilityDefAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetModels

`func (o *CapabilityServerModelsCapabilityDefAllOf) GetModels() []string`

GetModels returns the Models field if non-nil, zero value otherwise.

### GetModelsOk

`func (o *CapabilityServerModelsCapabilityDefAllOf) GetModelsOk() (*[]string, bool)`

GetModelsOk returns a tuple with the Models field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModels

`func (o *CapabilityServerModelsCapabilityDefAllOf) SetModels(v []string)`

SetModels sets Models field to given value.

### HasModels

`func (o *CapabilityServerModelsCapabilityDefAllOf) HasModels() bool`

HasModels returns a boolean if a field has been set.

### SetModelsNil

`func (o *CapabilityServerModelsCapabilityDefAllOf) SetModelsNil(b bool)`

 SetModelsNil sets the value for Models to be an explicit nil

### UnsetModels
`func (o *CapabilityServerModelsCapabilityDefAllOf) UnsetModels()`

UnsetModels ensures that no value is present for Models, not even an explicit nil
### GetServerType

`func (o *CapabilityServerModelsCapabilityDefAllOf) GetServerType() string`

GetServerType returns the ServerType field if non-nil, zero value otherwise.

### GetServerTypeOk

`func (o *CapabilityServerModelsCapabilityDefAllOf) GetServerTypeOk() (*string, bool)`

GetServerTypeOk returns a tuple with the ServerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerType

`func (o *CapabilityServerModelsCapabilityDefAllOf) SetServerType(v string)`

SetServerType sets ServerType field to given value.

### HasServerType

`func (o *CapabilityServerModelsCapabilityDefAllOf) HasServerType() bool`

HasServerType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


