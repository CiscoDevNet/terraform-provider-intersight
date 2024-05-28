# ApplianceBackupRotateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "appliance.BackupRotateData"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "appliance.BackupRotateData"]
**BackupTime** | Pointer to **time.Time** | The time at which the backup was taken. | [optional] [readonly] 
**Account** | Pointer to [**NullableIamAccountRelationship**](IamAccountRelationship.md) |  | [optional] 

## Methods

### NewApplianceBackupRotateData

`func NewApplianceBackupRotateData(classId string, objectType string, ) *ApplianceBackupRotateData`

NewApplianceBackupRotateData instantiates a new ApplianceBackupRotateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplianceBackupRotateDataWithDefaults

`func NewApplianceBackupRotateDataWithDefaults() *ApplianceBackupRotateData`

NewApplianceBackupRotateDataWithDefaults instantiates a new ApplianceBackupRotateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ApplianceBackupRotateData) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ApplianceBackupRotateData) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ApplianceBackupRotateData) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ApplianceBackupRotateData) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ApplianceBackupRotateData) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ApplianceBackupRotateData) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetBackupTime

`func (o *ApplianceBackupRotateData) GetBackupTime() time.Time`

GetBackupTime returns the BackupTime field if non-nil, zero value otherwise.

### GetBackupTimeOk

`func (o *ApplianceBackupRotateData) GetBackupTimeOk() (*time.Time, bool)`

GetBackupTimeOk returns a tuple with the BackupTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupTime

`func (o *ApplianceBackupRotateData) SetBackupTime(v time.Time)`

SetBackupTime sets BackupTime field to given value.

### HasBackupTime

`func (o *ApplianceBackupRotateData) HasBackupTime() bool`

HasBackupTime returns a boolean if a field has been set.

### GetAccount

`func (o *ApplianceBackupRotateData) GetAccount() IamAccountRelationship`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *ApplianceBackupRotateData) GetAccountOk() (*IamAccountRelationship, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *ApplianceBackupRotateData) SetAccount(v IamAccountRelationship)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *ApplianceBackupRotateData) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### SetAccountNil

`func (o *ApplianceBackupRotateData) SetAccountNil(b bool)`

 SetAccountNil sets the value for Account to be an explicit nil

### UnsetAccount
`func (o *ApplianceBackupRotateData) UnsetAccount()`

UnsetAccount ensures that no value is present for Account, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


