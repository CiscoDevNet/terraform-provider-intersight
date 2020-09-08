# ContentComplexType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The unique name of this complex type within the grammar specification. | [optional] 
**Parameters** | Pointer to [**[]ContentBaseParameter**](content.BaseParameter.md) |  | [optional] 

## Methods

### NewContentComplexType

`func NewContentComplexType() *ContentComplexType`

NewContentComplexType instantiates a new ContentComplexType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContentComplexTypeWithDefaults

`func NewContentComplexTypeWithDefaults() *ContentComplexType`

NewContentComplexTypeWithDefaults instantiates a new ContentComplexType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ContentComplexType) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ContentComplexType) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ContentComplexType) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ContentComplexType) HasName() bool`

HasName returns a boolean if a field has been set.

### GetParameters

`func (o *ContentComplexType) GetParameters() []ContentBaseParameter`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *ContentComplexType) GetParametersOk() (*[]ContentBaseParameter, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *ContentComplexType) SetParameters(v []ContentBaseParameter)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *ContentComplexType) HasParameters() bool`

HasParameters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


