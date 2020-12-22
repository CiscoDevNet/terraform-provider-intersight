# ContentTextParameter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "content.TextParameter"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "content.TextParameter"]
**IsDelimiter** | Pointer to **bool** | Data to be extracted from text content can be simple type or complex type or collection of simple/complex types. Complex types are group of simple or complex type. Delimiter is required to stop parsing list and complex data types. isDelimiter specifies whether given TextParameter is a delimiter or regular rule to capture the text data. | [optional] [default to false]
**IsNextCaptureOnSameLine** | Pointer to **bool** | Set to true of the next value to capture resides on the same text line of current match. By default textFSM engine gets the next text line on finding the first match. | [optional] [default to false]
**RegexLine** | Pointer to **string** | Regular expression of the line containing the data to be extracted from text content. | [optional] 

## Methods

### NewContentTextParameter

`func NewContentTextParameter(classId string, objectType string, ) *ContentTextParameter`

NewContentTextParameter instantiates a new ContentTextParameter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContentTextParameterWithDefaults

`func NewContentTextParameterWithDefaults() *ContentTextParameter`

NewContentTextParameterWithDefaults instantiates a new ContentTextParameter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ContentTextParameter) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ContentTextParameter) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ContentTextParameter) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ContentTextParameter) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ContentTextParameter) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ContentTextParameter) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetIsDelimiter

`func (o *ContentTextParameter) GetIsDelimiter() bool`

GetIsDelimiter returns the IsDelimiter field if non-nil, zero value otherwise.

### GetIsDelimiterOk

`func (o *ContentTextParameter) GetIsDelimiterOk() (*bool, bool)`

GetIsDelimiterOk returns a tuple with the IsDelimiter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDelimiter

`func (o *ContentTextParameter) SetIsDelimiter(v bool)`

SetIsDelimiter sets IsDelimiter field to given value.

### HasIsDelimiter

`func (o *ContentTextParameter) HasIsDelimiter() bool`

HasIsDelimiter returns a boolean if a field has been set.

### GetIsNextCaptureOnSameLine

`func (o *ContentTextParameter) GetIsNextCaptureOnSameLine() bool`

GetIsNextCaptureOnSameLine returns the IsNextCaptureOnSameLine field if non-nil, zero value otherwise.

### GetIsNextCaptureOnSameLineOk

`func (o *ContentTextParameter) GetIsNextCaptureOnSameLineOk() (*bool, bool)`

GetIsNextCaptureOnSameLineOk returns a tuple with the IsNextCaptureOnSameLine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsNextCaptureOnSameLine

`func (o *ContentTextParameter) SetIsNextCaptureOnSameLine(v bool)`

SetIsNextCaptureOnSameLine sets IsNextCaptureOnSameLine field to given value.

### HasIsNextCaptureOnSameLine

`func (o *ContentTextParameter) HasIsNextCaptureOnSameLine() bool`

HasIsNextCaptureOnSameLine returns a boolean if a field has been set.

### GetRegexLine

`func (o *ContentTextParameter) GetRegexLine() string`

GetRegexLine returns the RegexLine field if non-nil, zero value otherwise.

### GetRegexLineOk

`func (o *ContentTextParameter) GetRegexLineOk() (*string, bool)`

GetRegexLineOk returns a tuple with the RegexLine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegexLine

`func (o *ContentTextParameter) SetRegexLine(v string)`

SetRegexLine sets RegexLine field to given value.

### HasRegexLine

`func (o *ContentTextParameter) HasRegexLine() bool`

HasRegexLine returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


