# ContentComplexTypeAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The unique name of this complex type within the grammar specification. | [optional] 
**Parameters** | Pointer to [**[]ContentBaseParameter**](content.BaseParameter.md) |  | [optional] 

## Methods

### NewContentComplexTypeAllOf

`func NewContentComplexTypeAllOf() *ContentComplexTypeAllOf`

NewContentComplexTypeAllOf instantiates a new ContentComplexTypeAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContentComplexTypeAllOfWithDefaults

`func NewContentComplexTypeAllOfWithDefaults() *ContentComplexTypeAllOf`

NewContentComplexTypeAllOfWithDefaults instantiates a new ContentComplexTypeAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ContentComplexTypeAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ContentComplexTypeAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ContentComplexTypeAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ContentComplexTypeAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetParameters

`func (o *ContentComplexTypeAllOf) GetParameters() []ContentBaseParameter`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *ContentComplexTypeAllOf) GetParametersOk() (*[]ContentBaseParameter, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *ContentComplexTypeAllOf) SetParameters(v []ContentBaseParameter)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *ContentComplexTypeAllOf) HasParameters() bool`

HasParameters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


