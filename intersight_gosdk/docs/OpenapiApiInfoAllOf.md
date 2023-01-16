# OpenapiApiInfoAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "openapi.ApiInfo"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "openapi.ApiInfo"]
**ApiPathObjectIdentifier** | Pointer to [**MoMoRef**](MoMoRef.md) |  | [optional] 
**Description** | Pointer to **string** | A detailed description of the API. | [optional] 
**DisplayLabel** | Pointer to **string** | Display name of the selected API endpoint. | [optional] 
**Method** | Pointer to **string** | Method Type of the selected API. * &#x60;GET&#x60; - Method type which indicates it is a GET API call * &#x60;POST&#x60; - Method type which indicates it is a POST API call * &#x60;PUT&#x60; - Method type which indicates it is a PUT API call * &#x60;PATCH&#x60; - Method type which indicates it is a PATCH API call * &#x60;DELETE&#x60; - Method type which indicates it is a DELETE API call | [optional] [default to "GET"]
**Name** | Pointer to **string** | Name of the selected API endpoint. | [optional] 
**Path** | Pointer to **string** | API Path of the selected API endpoint. | [optional] 
**ValidationError** | Pointer to **string** | Validation error messages will be captured by this property. | [optional] [readonly] 
**ValidationStatus** | Pointer to **string** | Validation Status of selected API that indicates if the API validation passed or failed. * &#x60;none&#x60; - Indicates the default status. * &#x60;Success&#x60; - Denotes that the validation is successful. * &#x60;Failed&#x60; - Denotes that the validation failed. Validation could fail if information present in the OpenAPI specification is incomplete or missing. | [optional] [readonly] [default to "none"]

## Methods

### NewOpenapiApiInfoAllOf

`func NewOpenapiApiInfoAllOf(classId string, objectType string, ) *OpenapiApiInfoAllOf`

NewOpenapiApiInfoAllOf instantiates a new OpenapiApiInfoAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpenapiApiInfoAllOfWithDefaults

`func NewOpenapiApiInfoAllOfWithDefaults() *OpenapiApiInfoAllOf`

NewOpenapiApiInfoAllOfWithDefaults instantiates a new OpenapiApiInfoAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *OpenapiApiInfoAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *OpenapiApiInfoAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *OpenapiApiInfoAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *OpenapiApiInfoAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *OpenapiApiInfoAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *OpenapiApiInfoAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetApiPathObjectIdentifier

`func (o *OpenapiApiInfoAllOf) GetApiPathObjectIdentifier() MoMoRef`

GetApiPathObjectIdentifier returns the ApiPathObjectIdentifier field if non-nil, zero value otherwise.

### GetApiPathObjectIdentifierOk

`func (o *OpenapiApiInfoAllOf) GetApiPathObjectIdentifierOk() (*MoMoRef, bool)`

GetApiPathObjectIdentifierOk returns a tuple with the ApiPathObjectIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiPathObjectIdentifier

`func (o *OpenapiApiInfoAllOf) SetApiPathObjectIdentifier(v MoMoRef)`

SetApiPathObjectIdentifier sets ApiPathObjectIdentifier field to given value.

### HasApiPathObjectIdentifier

`func (o *OpenapiApiInfoAllOf) HasApiPathObjectIdentifier() bool`

HasApiPathObjectIdentifier returns a boolean if a field has been set.

### GetDescription

`func (o *OpenapiApiInfoAllOf) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OpenapiApiInfoAllOf) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OpenapiApiInfoAllOf) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *OpenapiApiInfoAllOf) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayLabel

`func (o *OpenapiApiInfoAllOf) GetDisplayLabel() string`

GetDisplayLabel returns the DisplayLabel field if non-nil, zero value otherwise.

### GetDisplayLabelOk

`func (o *OpenapiApiInfoAllOf) GetDisplayLabelOk() (*string, bool)`

GetDisplayLabelOk returns a tuple with the DisplayLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayLabel

`func (o *OpenapiApiInfoAllOf) SetDisplayLabel(v string)`

SetDisplayLabel sets DisplayLabel field to given value.

### HasDisplayLabel

`func (o *OpenapiApiInfoAllOf) HasDisplayLabel() bool`

HasDisplayLabel returns a boolean if a field has been set.

### GetMethod

`func (o *OpenapiApiInfoAllOf) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *OpenapiApiInfoAllOf) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *OpenapiApiInfoAllOf) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *OpenapiApiInfoAllOf) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetName

`func (o *OpenapiApiInfoAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OpenapiApiInfoAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OpenapiApiInfoAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OpenapiApiInfoAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPath

`func (o *OpenapiApiInfoAllOf) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *OpenapiApiInfoAllOf) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *OpenapiApiInfoAllOf) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *OpenapiApiInfoAllOf) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetValidationError

`func (o *OpenapiApiInfoAllOf) GetValidationError() string`

GetValidationError returns the ValidationError field if non-nil, zero value otherwise.

### GetValidationErrorOk

`func (o *OpenapiApiInfoAllOf) GetValidationErrorOk() (*string, bool)`

GetValidationErrorOk returns a tuple with the ValidationError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationError

`func (o *OpenapiApiInfoAllOf) SetValidationError(v string)`

SetValidationError sets ValidationError field to given value.

### HasValidationError

`func (o *OpenapiApiInfoAllOf) HasValidationError() bool`

HasValidationError returns a boolean if a field has been set.

### GetValidationStatus

`func (o *OpenapiApiInfoAllOf) GetValidationStatus() string`

GetValidationStatus returns the ValidationStatus field if non-nil, zero value otherwise.

### GetValidationStatusOk

`func (o *OpenapiApiInfoAllOf) GetValidationStatusOk() (*string, bool)`

GetValidationStatusOk returns a tuple with the ValidationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationStatus

`func (o *OpenapiApiInfoAllOf) SetValidationStatus(v string)`

SetValidationStatus sets ValidationStatus field to given value.

### HasValidationStatus

`func (o *OpenapiApiInfoAllOf) HasValidationStatus() bool`

HasValidationStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


