# IssueMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "issue.Message"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "issue.Message"]
**Message** | Pointer to **string** | A parameterized message such as \&quot;The temperature is %[1]v.\&quot; where &#39;hot&#39; could be substituted for %[1]v. | [optional] 
**Parameters** | Pointer to **[]string** |  | [optional] 

## Methods

### NewIssueMessage

`func NewIssueMessage(classId string, objectType string, ) *IssueMessage`

NewIssueMessage instantiates a new IssueMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIssueMessageWithDefaults

`func NewIssueMessageWithDefaults() *IssueMessage`

NewIssueMessageWithDefaults instantiates a new IssueMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *IssueMessage) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *IssueMessage) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *IssueMessage) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *IssueMessage) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *IssueMessage) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *IssueMessage) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetMessage

`func (o *IssueMessage) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *IssueMessage) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *IssueMessage) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *IssueMessage) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetParameters

`func (o *IssueMessage) GetParameters() []string`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *IssueMessage) GetParametersOk() (*[]string, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *IssueMessage) SetParameters(v []string)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *IssueMessage) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### SetParametersNil

`func (o *IssueMessage) SetParametersNil(b bool)`

 SetParametersNil sets the value for Parameters to be an explicit nil

### UnsetParameters
`func (o *IssueMessage) UnsetParameters()`

UnsetParameters ensures that no value is present for Parameters, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


