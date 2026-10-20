# IamPasswordHistorySetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "iam.PasswordHistorySetting"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "iam.PasswordHistorySetting"]
**HistoryIndex** | Pointer to **int64** | Index of the next passwordHistory collection. | [optional] [readonly] [default to 0]
**LastModificationTime** | Pointer to **time.Time** | Time when the password was last modified. | [optional] [readonly] 
**Username** | Pointer to **string** | User for which this password history is applicable. | [optional] [readonly] 

## Methods

### NewIamPasswordHistorySetting

`func NewIamPasswordHistorySetting(classId string, objectType string, ) *IamPasswordHistorySetting`

NewIamPasswordHistorySetting instantiates a new IamPasswordHistorySetting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamPasswordHistorySettingWithDefaults

`func NewIamPasswordHistorySettingWithDefaults() *IamPasswordHistorySetting`

NewIamPasswordHistorySettingWithDefaults instantiates a new IamPasswordHistorySetting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *IamPasswordHistorySetting) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *IamPasswordHistorySetting) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *IamPasswordHistorySetting) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *IamPasswordHistorySetting) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *IamPasswordHistorySetting) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *IamPasswordHistorySetting) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetHistoryIndex

`func (o *IamPasswordHistorySetting) GetHistoryIndex() int64`

GetHistoryIndex returns the HistoryIndex field if non-nil, zero value otherwise.

### GetHistoryIndexOk

`func (o *IamPasswordHistorySetting) GetHistoryIndexOk() (*int64, bool)`

GetHistoryIndexOk returns a tuple with the HistoryIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistoryIndex

`func (o *IamPasswordHistorySetting) SetHistoryIndex(v int64)`

SetHistoryIndex sets HistoryIndex field to given value.

### HasHistoryIndex

`func (o *IamPasswordHistorySetting) HasHistoryIndex() bool`

HasHistoryIndex returns a boolean if a field has been set.

### GetLastModificationTime

`func (o *IamPasswordHistorySetting) GetLastModificationTime() time.Time`

GetLastModificationTime returns the LastModificationTime field if non-nil, zero value otherwise.

### GetLastModificationTimeOk

`func (o *IamPasswordHistorySetting) GetLastModificationTimeOk() (*time.Time, bool)`

GetLastModificationTimeOk returns a tuple with the LastModificationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModificationTime

`func (o *IamPasswordHistorySetting) SetLastModificationTime(v time.Time)`

SetLastModificationTime sets LastModificationTime field to given value.

### HasLastModificationTime

`func (o *IamPasswordHistorySetting) HasLastModificationTime() bool`

HasLastModificationTime returns a boolean if a field has been set.

### GetUsername

`func (o *IamPasswordHistorySetting) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *IamPasswordHistorySetting) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *IamPasswordHistorySetting) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *IamPasswordHistorySetting) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


