# MgmtRestoreOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "mgmt.RestoreOptions"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "mgmt.RestoreOptions"]
**ActionOnError** | Pointer to **string** | The action to be taken if an error is encountered during the restore operation. * &#x60;Continue&#x60; - The restore operation will continue on encountering an error.Any failure will be tagged as incomplete and the operation will continue importing other configurations.User shall review the partially imported configuration and fix later. * &#x60;Stop&#x60; - The restore operation will stop on encountering an error.Configurations that were changed as part of the restore operation till the error was encountered will remain in the system. * &#x60;Rollback&#x60; - Configurations that were changed as part of the restore operation till the error was encountered will be rolled back. | [optional] [default to "Continue"]
**OrgsToRestore** | Pointer to **[]string** |  | [optional] 
**RestoreAction** | Pointer to **string** | Mechanism to be used for updating configurations as part of a restore operation. * &#x60;Merge&#x60; - Configurations in the backup will be merged with existing configurations in the target system. * &#x60;Replace&#x60; - Configurations in the target system will be replaced with the configurations in the backup.Configurations that are present only in the target system and not in the backup will be removed. | [optional] [default to "Merge"]

## Methods

### NewMgmtRestoreOptions

`func NewMgmtRestoreOptions(classId string, objectType string, ) *MgmtRestoreOptions`

NewMgmtRestoreOptions instantiates a new MgmtRestoreOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMgmtRestoreOptionsWithDefaults

`func NewMgmtRestoreOptionsWithDefaults() *MgmtRestoreOptions`

NewMgmtRestoreOptionsWithDefaults instantiates a new MgmtRestoreOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *MgmtRestoreOptions) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *MgmtRestoreOptions) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *MgmtRestoreOptions) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *MgmtRestoreOptions) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *MgmtRestoreOptions) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *MgmtRestoreOptions) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetActionOnError

`func (o *MgmtRestoreOptions) GetActionOnError() string`

GetActionOnError returns the ActionOnError field if non-nil, zero value otherwise.

### GetActionOnErrorOk

`func (o *MgmtRestoreOptions) GetActionOnErrorOk() (*string, bool)`

GetActionOnErrorOk returns a tuple with the ActionOnError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionOnError

`func (o *MgmtRestoreOptions) SetActionOnError(v string)`

SetActionOnError sets ActionOnError field to given value.

### HasActionOnError

`func (o *MgmtRestoreOptions) HasActionOnError() bool`

HasActionOnError returns a boolean if a field has been set.

### GetOrgsToRestore

`func (o *MgmtRestoreOptions) GetOrgsToRestore() []string`

GetOrgsToRestore returns the OrgsToRestore field if non-nil, zero value otherwise.

### GetOrgsToRestoreOk

`func (o *MgmtRestoreOptions) GetOrgsToRestoreOk() (*[]string, bool)`

GetOrgsToRestoreOk returns a tuple with the OrgsToRestore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgsToRestore

`func (o *MgmtRestoreOptions) SetOrgsToRestore(v []string)`

SetOrgsToRestore sets OrgsToRestore field to given value.

### HasOrgsToRestore

`func (o *MgmtRestoreOptions) HasOrgsToRestore() bool`

HasOrgsToRestore returns a boolean if a field has been set.

### SetOrgsToRestoreNil

`func (o *MgmtRestoreOptions) SetOrgsToRestoreNil(b bool)`

 SetOrgsToRestoreNil sets the value for OrgsToRestore to be an explicit nil

### UnsetOrgsToRestore
`func (o *MgmtRestoreOptions) UnsetOrgsToRestore()`

UnsetOrgsToRestore ensures that no value is present for OrgsToRestore, not even an explicit nil
### GetRestoreAction

`func (o *MgmtRestoreOptions) GetRestoreAction() string`

GetRestoreAction returns the RestoreAction field if non-nil, zero value otherwise.

### GetRestoreActionOk

`func (o *MgmtRestoreOptions) GetRestoreActionOk() (*string, bool)`

GetRestoreActionOk returns a tuple with the RestoreAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreAction

`func (o *MgmtRestoreOptions) SetRestoreAction(v string)`

SetRestoreAction sets RestoreAction field to given value.

### HasRestoreAction

`func (o *MgmtRestoreOptions) HasRestoreAction() bool`

HasRestoreAction returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


