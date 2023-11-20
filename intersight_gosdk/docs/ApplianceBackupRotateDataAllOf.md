# ApplianceBackupRotateDataAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "appliance.BackupRotateData"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "appliance.BackupRotateData"]
**BackupTime** | Pointer to **time.Time** | The time at which the backup was taken. | [optional] [readonly] 
**Account** | Pointer to [**IamAccountRelationship**](IamAccountRelationship.md) |  | [optional] 

## Methods

### NewApplianceBackupRotateDataAllOf

`func NewApplianceBackupRotateDataAllOf(classId string, objectType string, ) *ApplianceBackupRotateDataAllOf`

NewApplianceBackupRotateDataAllOf instantiates a new ApplianceBackupRotateDataAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplianceBackupRotateDataAllOfWithDefaults

`func NewApplianceBackupRotateDataAllOfWithDefaults() *ApplianceBackupRotateDataAllOf`

NewApplianceBackupRotateDataAllOfWithDefaults instantiates a new ApplianceBackupRotateDataAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ApplianceBackupRotateDataAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ApplianceBackupRotateDataAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ApplianceBackupRotateDataAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ApplianceBackupRotateDataAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ApplianceBackupRotateDataAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ApplianceBackupRotateDataAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetBackupTime

`func (o *ApplianceBackupRotateDataAllOf) GetBackupTime() time.Time`

GetBackupTime returns the BackupTime field if non-nil, zero value otherwise.

### GetBackupTimeOk

`func (o *ApplianceBackupRotateDataAllOf) GetBackupTimeOk() (*time.Time, bool)`

GetBackupTimeOk returns a tuple with the BackupTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupTime

`func (o *ApplianceBackupRotateDataAllOf) SetBackupTime(v time.Time)`

SetBackupTime sets BackupTime field to given value.

### HasBackupTime

`func (o *ApplianceBackupRotateDataAllOf) HasBackupTime() bool`

HasBackupTime returns a boolean if a field has been set.

### GetAccount

`func (o *ApplianceBackupRotateDataAllOf) GetAccount() IamAccountRelationship`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *ApplianceBackupRotateDataAllOf) GetAccountOk() (*IamAccountRelationship, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *ApplianceBackupRotateDataAllOf) SetAccount(v IamAccountRelationship)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *ApplianceBackupRotateDataAllOf) HasAccount() bool`

HasAccount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


