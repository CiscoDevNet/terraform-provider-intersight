# OpenapiApiMethodMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "openapi.ApiMethodMeta"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "openapi.ApiMethodMeta"]
**Description** | Pointer to **string** | The description of the given API. | [optional] [readonly] 
**DisplayLabel** | Pointer to **string** | The display label of the given API. | [optional] [readonly] 
**Method** | Pointer to **string** | The method type for the given API. * &#x60;GET&#x60; - Method type which indicates it is a GET API call. * &#x60;POST&#x60; - Method type which indicates it is a POST API call. * &#x60;PUT&#x60; - Method type which indicates it is a PUT API call. * &#x60;PATCH&#x60; - Method type which indicates it is a PATCH API call. * &#x60;DELETE&#x60; - Method type which indicates it is a DELETE API call. | [optional] [readonly] [default to "GET"]
**Name** | Pointer to **string** | The description of the given API. | [optional] [readonly] 
**Path** | Pointer to **string** | Path of the selected API endpoint. | [optional] [readonly] 
**Source** | Pointer to [**OpenapiProcessFileRelationship**](OpenapiProcessFileRelationship.md) |  | [optional] 

## Methods

### NewOpenapiApiMethodMeta

`func NewOpenapiApiMethodMeta(classId string, objectType string, ) *OpenapiApiMethodMeta`

NewOpenapiApiMethodMeta instantiates a new OpenapiApiMethodMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpenapiApiMethodMetaWithDefaults

`func NewOpenapiApiMethodMetaWithDefaults() *OpenapiApiMethodMeta`

NewOpenapiApiMethodMetaWithDefaults instantiates a new OpenapiApiMethodMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *OpenapiApiMethodMeta) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *OpenapiApiMethodMeta) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *OpenapiApiMethodMeta) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *OpenapiApiMethodMeta) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *OpenapiApiMethodMeta) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *OpenapiApiMethodMeta) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDescription

`func (o *OpenapiApiMethodMeta) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OpenapiApiMethodMeta) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OpenapiApiMethodMeta) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *OpenapiApiMethodMeta) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayLabel

`func (o *OpenapiApiMethodMeta) GetDisplayLabel() string`

GetDisplayLabel returns the DisplayLabel field if non-nil, zero value otherwise.

### GetDisplayLabelOk

`func (o *OpenapiApiMethodMeta) GetDisplayLabelOk() (*string, bool)`

GetDisplayLabelOk returns a tuple with the DisplayLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayLabel

`func (o *OpenapiApiMethodMeta) SetDisplayLabel(v string)`

SetDisplayLabel sets DisplayLabel field to given value.

### HasDisplayLabel

`func (o *OpenapiApiMethodMeta) HasDisplayLabel() bool`

HasDisplayLabel returns a boolean if a field has been set.

### GetMethod

`func (o *OpenapiApiMethodMeta) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *OpenapiApiMethodMeta) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *OpenapiApiMethodMeta) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *OpenapiApiMethodMeta) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetName

`func (o *OpenapiApiMethodMeta) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OpenapiApiMethodMeta) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OpenapiApiMethodMeta) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OpenapiApiMethodMeta) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPath

`func (o *OpenapiApiMethodMeta) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *OpenapiApiMethodMeta) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *OpenapiApiMethodMeta) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *OpenapiApiMethodMeta) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetSource

`func (o *OpenapiApiMethodMeta) GetSource() OpenapiProcessFileRelationship`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *OpenapiApiMethodMeta) GetSourceOk() (*OpenapiProcessFileRelationship, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *OpenapiApiMethodMeta) SetSource(v OpenapiProcessFileRelationship)`

SetSource sets Source field to given value.

### HasSource

`func (o *OpenapiApiMethodMeta) HasSource() bool`

HasSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


