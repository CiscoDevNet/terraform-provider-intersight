# ContentTextParameterAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsDelimiter** | Pointer to **bool** | Data to be extracted from text content can be simple type or complex type or collection of simple/complex types. Complex types are group of simple or complex type. Delimiter is required to stop parsing list and complex data types. isDelimiter specifies whether given TextParameter is a delimiter or regular rule to capture the text data. | [optional] 
**IsNextCaptureOnSameLine** | Pointer to **bool** | Set to true of the next value to capture resides on the same text line of current match. By default textFSM engine gets the next text line on finding the first match. | [optional] 
**RegexLine** | Pointer to **string** | Regular expression of the line containing the data to be extracted from text content. | [optional] 

## Methods

### NewContentTextParameterAllOf

`func NewContentTextParameterAllOf() *ContentTextParameterAllOf`

NewContentTextParameterAllOf instantiates a new ContentTextParameterAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContentTextParameterAllOfWithDefaults

`func NewContentTextParameterAllOfWithDefaults() *ContentTextParameterAllOf`

NewContentTextParameterAllOfWithDefaults instantiates a new ContentTextParameterAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsDelimiter

`func (o *ContentTextParameterAllOf) GetIsDelimiter() bool`

GetIsDelimiter returns the IsDelimiter field if non-nil, zero value otherwise.

### GetIsDelimiterOk

`func (o *ContentTextParameterAllOf) GetIsDelimiterOk() (*bool, bool)`

GetIsDelimiterOk returns a tuple with the IsDelimiter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDelimiter

`func (o *ContentTextParameterAllOf) SetIsDelimiter(v bool)`

SetIsDelimiter sets IsDelimiter field to given value.

### HasIsDelimiter

`func (o *ContentTextParameterAllOf) HasIsDelimiter() bool`

HasIsDelimiter returns a boolean if a field has been set.

### GetIsNextCaptureOnSameLine

`func (o *ContentTextParameterAllOf) GetIsNextCaptureOnSameLine() bool`

GetIsNextCaptureOnSameLine returns the IsNextCaptureOnSameLine field if non-nil, zero value otherwise.

### GetIsNextCaptureOnSameLineOk

`func (o *ContentTextParameterAllOf) GetIsNextCaptureOnSameLineOk() (*bool, bool)`

GetIsNextCaptureOnSameLineOk returns a tuple with the IsNextCaptureOnSameLine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsNextCaptureOnSameLine

`func (o *ContentTextParameterAllOf) SetIsNextCaptureOnSameLine(v bool)`

SetIsNextCaptureOnSameLine sets IsNextCaptureOnSameLine field to given value.

### HasIsNextCaptureOnSameLine

`func (o *ContentTextParameterAllOf) HasIsNextCaptureOnSameLine() bool`

HasIsNextCaptureOnSameLine returns a boolean if a field has been set.

### GetRegexLine

`func (o *ContentTextParameterAllOf) GetRegexLine() string`

GetRegexLine returns the RegexLine field if non-nil, zero value otherwise.

### GetRegexLineOk

`func (o *ContentTextParameterAllOf) GetRegexLineOk() (*string, bool)`

GetRegexLineOk returns a tuple with the RegexLine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegexLine

`func (o *ContentTextParameterAllOf) SetRegexLine(v string)`

SetRegexLine sets RegexLine field to given value.

### HasRegexLine

`func (o *ContentTextParameterAllOf) HasRegexLine() bool`

HasRegexLine returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


