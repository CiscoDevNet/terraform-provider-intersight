# MotemplateActionEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "motemplate.ActionEntry"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "motemplate.ActionEntry"]
**Params** | Pointer to [**[]MotemplateActionParam**](MotemplateActionParam.md) |  | [optional] 
**Type** | Pointer to **string** | The action type to be executed. * &#x60;Sync&#x60; - The action to merge values from the template to its derived objects. * &#x60;Deploy&#x60; - The action to execute deploy action on all the objects derived from the template that is mainly applicable for the various profile types. * &#x60;Detach&#x60; - The action to detach the current derived object from its attached template. * &#x60;Attach&#x60; - The action to attach the current object to the specified template. | [optional] [default to "Sync"]

## Methods

### NewMotemplateActionEntry

`func NewMotemplateActionEntry(classId string, objectType string, ) *MotemplateActionEntry`

NewMotemplateActionEntry instantiates a new MotemplateActionEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMotemplateActionEntryWithDefaults

`func NewMotemplateActionEntryWithDefaults() *MotemplateActionEntry`

NewMotemplateActionEntryWithDefaults instantiates a new MotemplateActionEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *MotemplateActionEntry) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *MotemplateActionEntry) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *MotemplateActionEntry) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *MotemplateActionEntry) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *MotemplateActionEntry) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *MotemplateActionEntry) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetParams

`func (o *MotemplateActionEntry) GetParams() []MotemplateActionParam`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *MotemplateActionEntry) GetParamsOk() (*[]MotemplateActionParam, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *MotemplateActionEntry) SetParams(v []MotemplateActionParam)`

SetParams sets Params field to given value.

### HasParams

`func (o *MotemplateActionEntry) HasParams() bool`

HasParams returns a boolean if a field has been set.

### SetParamsNil

`func (o *MotemplateActionEntry) SetParamsNil(b bool)`

 SetParamsNil sets the value for Params to be an explicit nil

### UnsetParams
`func (o *MotemplateActionEntry) UnsetParams()`

UnsetParams ensures that no value is present for Params, not even an explicit nil
### GetType

`func (o *MotemplateActionEntry) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MotemplateActionEntry) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MotemplateActionEntry) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *MotemplateActionEntry) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


