# IssueDeviceTag

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "issue.DeviceTag"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "issue.DeviceTag"]
**Key** | Pointer to **string** | The string representation of a tag key. | [optional] 
**Value** | Pointer to **string** | The string representation of a tag value. | [optional] 

## Methods

### NewIssueDeviceTag

`func NewIssueDeviceTag(classId string, objectType string, ) *IssueDeviceTag`

NewIssueDeviceTag instantiates a new IssueDeviceTag object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIssueDeviceTagWithDefaults

`func NewIssueDeviceTagWithDefaults() *IssueDeviceTag`

NewIssueDeviceTagWithDefaults instantiates a new IssueDeviceTag object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *IssueDeviceTag) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *IssueDeviceTag) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *IssueDeviceTag) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *IssueDeviceTag) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *IssueDeviceTag) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *IssueDeviceTag) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetKey

`func (o *IssueDeviceTag) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *IssueDeviceTag) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *IssueDeviceTag) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *IssueDeviceTag) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetValue

`func (o *IssueDeviceTag) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *IssueDeviceTag) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *IssueDeviceTag) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *IssueDeviceTag) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


