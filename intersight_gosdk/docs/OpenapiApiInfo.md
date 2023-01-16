# OpenapiApiInfo

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

### NewOpenapiApiInfo

`func NewOpenapiApiInfo(classId string, objectType string, ) *OpenapiApiInfo`

NewOpenapiApiInfo instantiates a new OpenapiApiInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpenapiApiInfoWithDefaults

`func NewOpenapiApiInfoWithDefaults() *OpenapiApiInfo`

NewOpenapiApiInfoWithDefaults instantiates a new OpenapiApiInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *OpenapiApiInfo) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *OpenapiApiInfo) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *OpenapiApiInfo) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *OpenapiApiInfo) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *OpenapiApiInfo) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *OpenapiApiInfo) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetApiPathObjectIdentifier

`func (o *OpenapiApiInfo) GetApiPathObjectIdentifier() MoMoRef`

GetApiPathObjectIdentifier returns the ApiPathObjectIdentifier field if non-nil, zero value otherwise.

### GetApiPathObjectIdentifierOk

`func (o *OpenapiApiInfo) GetApiPathObjectIdentifierOk() (*MoMoRef, bool)`

GetApiPathObjectIdentifierOk returns a tuple with the ApiPathObjectIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiPathObjectIdentifier

`func (o *OpenapiApiInfo) SetApiPathObjectIdentifier(v MoMoRef)`

SetApiPathObjectIdentifier sets ApiPathObjectIdentifier field to given value.

### HasApiPathObjectIdentifier

`func (o *OpenapiApiInfo) HasApiPathObjectIdentifier() bool`

HasApiPathObjectIdentifier returns a boolean if a field has been set.

### GetDescription

`func (o *OpenapiApiInfo) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OpenapiApiInfo) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OpenapiApiInfo) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *OpenapiApiInfo) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayLabel

`func (o *OpenapiApiInfo) GetDisplayLabel() string`

GetDisplayLabel returns the DisplayLabel field if non-nil, zero value otherwise.

### GetDisplayLabelOk

`func (o *OpenapiApiInfo) GetDisplayLabelOk() (*string, bool)`

GetDisplayLabelOk returns a tuple with the DisplayLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayLabel

`func (o *OpenapiApiInfo) SetDisplayLabel(v string)`

SetDisplayLabel sets DisplayLabel field to given value.

### HasDisplayLabel

`func (o *OpenapiApiInfo) HasDisplayLabel() bool`

HasDisplayLabel returns a boolean if a field has been set.

### GetMethod

`func (o *OpenapiApiInfo) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *OpenapiApiInfo) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *OpenapiApiInfo) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *OpenapiApiInfo) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetName

`func (o *OpenapiApiInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OpenapiApiInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OpenapiApiInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OpenapiApiInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPath

`func (o *OpenapiApiInfo) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *OpenapiApiInfo) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *OpenapiApiInfo) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *OpenapiApiInfo) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetValidationError

`func (o *OpenapiApiInfo) GetValidationError() string`

GetValidationError returns the ValidationError field if non-nil, zero value otherwise.

### GetValidationErrorOk

`func (o *OpenapiApiInfo) GetValidationErrorOk() (*string, bool)`

GetValidationErrorOk returns a tuple with the ValidationError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationError

`func (o *OpenapiApiInfo) SetValidationError(v string)`

SetValidationError sets ValidationError field to given value.

### HasValidationError

`func (o *OpenapiApiInfo) HasValidationError() bool`

HasValidationError returns a boolean if a field has been set.

### GetValidationStatus

`func (o *OpenapiApiInfo) GetValidationStatus() string`

GetValidationStatus returns the ValidationStatus field if non-nil, zero value otherwise.

### GetValidationStatusOk

`func (o *OpenapiApiInfo) GetValidationStatusOk() (*string, bool)`

GetValidationStatusOk returns a tuple with the ValidationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationStatus

`func (o *OpenapiApiInfo) SetValidationStatus(v string)`

SetValidationStatus sets ValidationStatus field to given value.

### HasValidationStatus

`func (o *OpenapiApiInfo) HasValidationStatus() bool`

HasValidationStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


