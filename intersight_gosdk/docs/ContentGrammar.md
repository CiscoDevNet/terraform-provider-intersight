# ContentGrammar

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ErrorParameters** | Pointer to [**[]ContentBaseParameter**](content.BaseParameter.md) |  | [optional] 
**Parameters** | Pointer to [**[]ContentBaseParameter**](content.BaseParameter.md) |  | [optional] 
**Types** | Pointer to [**[]ContentComplexType**](content.ComplexType.md) |  | [optional] 

## Methods

### NewContentGrammar

`func NewContentGrammar() *ContentGrammar`

NewContentGrammar instantiates a new ContentGrammar object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContentGrammarWithDefaults

`func NewContentGrammarWithDefaults() *ContentGrammar`

NewContentGrammarWithDefaults instantiates a new ContentGrammar object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrorParameters

`func (o *ContentGrammar) GetErrorParameters() []ContentBaseParameter`

GetErrorParameters returns the ErrorParameters field if non-nil, zero value otherwise.

### GetErrorParametersOk

`func (o *ContentGrammar) GetErrorParametersOk() (*[]ContentBaseParameter, bool)`

GetErrorParametersOk returns a tuple with the ErrorParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorParameters

`func (o *ContentGrammar) SetErrorParameters(v []ContentBaseParameter)`

SetErrorParameters sets ErrorParameters field to given value.

### HasErrorParameters

`func (o *ContentGrammar) HasErrorParameters() bool`

HasErrorParameters returns a boolean if a field has been set.

### GetParameters

`func (o *ContentGrammar) GetParameters() []ContentBaseParameter`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *ContentGrammar) GetParametersOk() (*[]ContentBaseParameter, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *ContentGrammar) SetParameters(v []ContentBaseParameter)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *ContentGrammar) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### GetTypes

`func (o *ContentGrammar) GetTypes() []ContentComplexType`

GetTypes returns the Types field if non-nil, zero value otherwise.

### GetTypesOk

`func (o *ContentGrammar) GetTypesOk() (*[]ContentComplexType, bool)`

GetTypesOk returns a tuple with the Types field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypes

`func (o *ContentGrammar) SetTypes(v []ContentComplexType)`

SetTypes sets Types field to given value.

### HasTypes

`func (o *ContentGrammar) HasTypes() bool`

HasTypes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


