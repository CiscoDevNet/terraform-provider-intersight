# ContentBaseParameterAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcceptSingleValue** | Pointer to **bool** | The flag that allows single values in content to be extracted as a single element collection in case the parameter is of Collection type. This flag is applicable for parameters of type Collection only. | [optional] 
**ComplexType** | Pointer to **string** | The name of the complex type definition in case this is a complex parameter. The content.Grammar object must have a complex type, content.ComplexType, defined with the specified name in types collection property. | [optional] 
**ItemType** | Pointer to **string** | The type of the collection item in case this is a collection parameter. * &#x60;simple&#x60; - The parameter value to be extracted is of the type simple. All the common scalar typessuch as int, bool, string, etc are represented by the simple enum. * &#x60;complex&#x60; - The parameter value to be extracted is a complex parameter that itself isanother collection of simple/complex parameters. * &#x60;collection&#x60; - The parameter value to be extracted is a collection parameter whose item typeshall be either simple type or complex type. | [optional] [default to "simple"]
**Name** | Pointer to **string** | The name of the parameter. | [optional] 
**Path** | Pointer to **string** | The content specific path information that identifies the parameter value within the content. The value is usually a XPath or JSONPath or a regular expression in case of text content. | [optional] 
**Type** | Pointer to **string** | The type of the parameter. Accepted values are simple, complex, collection. * &#x60;simple&#x60; - The parameter value to be extracted is of the type simple. All the common scalar typessuch as int, bool, string, etc are represented by the simple enum. * &#x60;complex&#x60; - The parameter value to be extracted is a complex parameter that itself isanother collection of simple/complex parameters. * &#x60;collection&#x60; - The parameter value to be extracted is a collection parameter whose item typeshall be either simple type or complex type. | [optional] [default to "simple"]

## Methods

### NewContentBaseParameterAllOf

`func NewContentBaseParameterAllOf() *ContentBaseParameterAllOf`

NewContentBaseParameterAllOf instantiates a new ContentBaseParameterAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContentBaseParameterAllOfWithDefaults

`func NewContentBaseParameterAllOfWithDefaults() *ContentBaseParameterAllOf`

NewContentBaseParameterAllOfWithDefaults instantiates a new ContentBaseParameterAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcceptSingleValue

`func (o *ContentBaseParameterAllOf) GetAcceptSingleValue() bool`

GetAcceptSingleValue returns the AcceptSingleValue field if non-nil, zero value otherwise.

### GetAcceptSingleValueOk

`func (o *ContentBaseParameterAllOf) GetAcceptSingleValueOk() (*bool, bool)`

GetAcceptSingleValueOk returns a tuple with the AcceptSingleValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptSingleValue

`func (o *ContentBaseParameterAllOf) SetAcceptSingleValue(v bool)`

SetAcceptSingleValue sets AcceptSingleValue field to given value.

### HasAcceptSingleValue

`func (o *ContentBaseParameterAllOf) HasAcceptSingleValue() bool`

HasAcceptSingleValue returns a boolean if a field has been set.

### GetComplexType

`func (o *ContentBaseParameterAllOf) GetComplexType() string`

GetComplexType returns the ComplexType field if non-nil, zero value otherwise.

### GetComplexTypeOk

`func (o *ContentBaseParameterAllOf) GetComplexTypeOk() (*string, bool)`

GetComplexTypeOk returns a tuple with the ComplexType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplexType

`func (o *ContentBaseParameterAllOf) SetComplexType(v string)`

SetComplexType sets ComplexType field to given value.

### HasComplexType

`func (o *ContentBaseParameterAllOf) HasComplexType() bool`

HasComplexType returns a boolean if a field has been set.

### GetItemType

`func (o *ContentBaseParameterAllOf) GetItemType() string`

GetItemType returns the ItemType field if non-nil, zero value otherwise.

### GetItemTypeOk

`func (o *ContentBaseParameterAllOf) GetItemTypeOk() (*string, bool)`

GetItemTypeOk returns a tuple with the ItemType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemType

`func (o *ContentBaseParameterAllOf) SetItemType(v string)`

SetItemType sets ItemType field to given value.

### HasItemType

`func (o *ContentBaseParameterAllOf) HasItemType() bool`

HasItemType returns a boolean if a field has been set.

### GetName

`func (o *ContentBaseParameterAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ContentBaseParameterAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ContentBaseParameterAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ContentBaseParameterAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPath

`func (o *ContentBaseParameterAllOf) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *ContentBaseParameterAllOf) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *ContentBaseParameterAllOf) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *ContentBaseParameterAllOf) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetType

`func (o *ContentBaseParameterAllOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ContentBaseParameterAllOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ContentBaseParameterAllOf) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ContentBaseParameterAllOf) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


