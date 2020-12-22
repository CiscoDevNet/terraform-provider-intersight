# ContentTextParameterAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "content.TextParameter"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "content.TextParameter"]
**IsDelimiter** | Pointer to **bool** | Data to be extracted from text content can be simple type or complex type or collection of simple/complex types. Complex types are group of simple or complex type. Delimiter is required to stop parsing list and complex data types. isDelimiter specifies whether given TextParameter is a delimiter or regular rule to capture the text data. | [optional] [default to false]
**IsNextCaptureOnSameLine** | Pointer to **bool** | Set to true of the next value to capture resides on the same text line of current match. By default textFSM engine gets the next text line on finding the first match. | [optional] [default to false]
**RegexLine** | Pointer to **string** | Regular expression of the line containing the data to be extracted from text content. | [optional] 

## Methods

### NewContentTextParameterAllOf

`func NewContentTextParameterAllOf(classId string, objectType string, ) *ContentTextParameterAllOf`

NewContentTextParameterAllOf instantiates a new ContentTextParameterAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContentTextParameterAllOfWithDefaults

`func NewContentTextParameterAllOfWithDefaults() *ContentTextParameterAllOf`

NewContentTextParameterAllOfWithDefaults instantiates a new ContentTextParameterAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ContentTextParameterAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ContentTextParameterAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ContentTextParameterAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ContentTextParameterAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ContentTextParameterAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ContentTextParameterAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


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


