# MgmtConfigBackupFile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "mgmt.ConfigBackupFile"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "mgmt.ConfigBackupFile"]
**ConfigurationFileUploadId** | Pointer to **string** | A unique tracking ID for the file being uploaded. | [optional] [readonly] 
**Account** | Pointer to [**NullableIamAccountRelationship**](IamAccountRelationship.md) |  | [optional] 

## Methods

### NewMgmtConfigBackupFile

`func NewMgmtConfigBackupFile(classId string, objectType string, ) *MgmtConfigBackupFile`

NewMgmtConfigBackupFile instantiates a new MgmtConfigBackupFile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMgmtConfigBackupFileWithDefaults

`func NewMgmtConfigBackupFileWithDefaults() *MgmtConfigBackupFile`

NewMgmtConfigBackupFileWithDefaults instantiates a new MgmtConfigBackupFile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *MgmtConfigBackupFile) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *MgmtConfigBackupFile) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *MgmtConfigBackupFile) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *MgmtConfigBackupFile) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *MgmtConfigBackupFile) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *MgmtConfigBackupFile) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetConfigurationFileUploadId

`func (o *MgmtConfigBackupFile) GetConfigurationFileUploadId() string`

GetConfigurationFileUploadId returns the ConfigurationFileUploadId field if non-nil, zero value otherwise.

### GetConfigurationFileUploadIdOk

`func (o *MgmtConfigBackupFile) GetConfigurationFileUploadIdOk() (*string, bool)`

GetConfigurationFileUploadIdOk returns a tuple with the ConfigurationFileUploadId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationFileUploadId

`func (o *MgmtConfigBackupFile) SetConfigurationFileUploadId(v string)`

SetConfigurationFileUploadId sets ConfigurationFileUploadId field to given value.

### HasConfigurationFileUploadId

`func (o *MgmtConfigBackupFile) HasConfigurationFileUploadId() bool`

HasConfigurationFileUploadId returns a boolean if a field has been set.

### GetAccount

`func (o *MgmtConfigBackupFile) GetAccount() IamAccountRelationship`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *MgmtConfigBackupFile) GetAccountOk() (*IamAccountRelationship, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *MgmtConfigBackupFile) SetAccount(v IamAccountRelationship)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *MgmtConfigBackupFile) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### SetAccountNil

`func (o *MgmtConfigBackupFile) SetAccountNil(b bool)`

 SetAccountNil sets the value for Account to be an explicit nil

### UnsetAccount
`func (o *MgmtConfigBackupFile) UnsetAccount()`

UnsetAccount ensures that no value is present for Account, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


