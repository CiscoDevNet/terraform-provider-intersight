# MgmtConfigBackupOperation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "mgmt.ConfigBackupOperation"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "mgmt.ConfigBackupOperation"]
**BackupAccount** | Pointer to **string** | The moid of the account from which the backup was taken. | [optional] [readonly] 
**BackupAccountDomainGroupMoid** | Pointer to **string** | The domain group moid of the account from which the backup was taken. | [optional] [readonly] 
**CategoryExportData** | Pointer to [**[]MgmtCategoryExportData**](MgmtCategoryExportData.md) |  | [optional] 
**Description** | Pointer to **string** | Description of the backup operation. | [optional] 
**Destination** | Pointer to **string** | Indicates whether the backup was created locally or imported. * &#x60;Local&#x60; - A local Intersight location. * &#x60;Uploaded&#x60; - A local location where the backup file is uploaded as mgmt.ConfigBackupFile MO. Intersight creates the backup from the uploaded location when this location type is set. * &#x60;Remote&#x60; - A remote location hosted in the user&#39;s datacenter. | [optional] [default to "Local"]
**EndTime** | Pointer to **time.Time** | End date and time of the backup operation. | [optional] [readonly] 
**IsAesKeySet** | Pointer to **bool** | Indicates whether the value of the &#39;aesKey&#39; property has been set. | [optional] [readonly] [default to false]
**Name** | Pointer to **string** | User provided identifier for the backup operation. | [optional] 
**Options** | Pointer to [**NullableMgmtBackupOptions**](MgmtBackupOptions.md) |  | [optional] 
**RetentionLock** | Pointer to **bool** | When set, ensures that the backup archive is protected from deletion and rollover operations. The value for retention lock is in sync with the backup instance created as part of this operation. | [optional] [readonly] 
**RolloverInitiated** | Pointer to **bool** | Internal property used to indicate that a rollover has been initiated as part of this operation. | [optional] [readonly] 
**StartTime** | Pointer to **time.Time** | Start date and time of the backup operation. | [optional] [readonly] 
**Status** | Pointer to **string** | The current status of the backup operation. * &#x60;Scheduled&#x60; - Backup or restore process has been scheduled. * &#x60;NotInitiated&#x60; - Backup or restore process has not been initiated. * &#x60;InProgress&#x60; - Backup or restore process is in progress. * &#x60;Failed&#x60; - Backup or restore process has failed. * &#x60;Completed&#x60; - Backup or restore process has completed. * &#x60;NotComplete&#x60; - Backup or restore process has partially imported configuration. | [optional] [readonly] [default to "Scheduled"]
**SystemAdminTriggered** | Pointer to **bool** | Set when a system administrator initiated the backup operation. | [optional] [readonly] 
**UserPassword** | Pointer to **string** | The password provided by the user to decrypt a password encrypted backup file during import operation.  This password is used to decrypt the backup archive before validation when importing a password encrypted backup. This field is optional and only required when importing a password encrypted backup file. | [optional] 
**Account** | Pointer to [**NullableIamAccountRelationship**](IamAccountRelationship.md) |  | [optional] 
**BackupInstance** | Pointer to [**NullableMgmtConfigBackupInstanceRelationship**](MgmtConfigBackupInstanceRelationship.md) |  | [optional] 
**OrgBackups** | Pointer to [**[]MgmtOrgBackupOperationRelationship**](MgmtOrgBackupOperationRelationship.md) | An array of relationships to mgmtOrgBackupOperation resources. | [optional] [readonly] 
**UserBackupFile** | Pointer to [**NullableMgmtConfigBackupFileRelationship**](MgmtConfigBackupFileRelationship.md) |  | [optional] 
**WorkflowInfo** | Pointer to [**NullableWorkflowWorkflowInfoRelationship**](WorkflowWorkflowInfoRelationship.md) |  | [optional] 

## Methods

### NewMgmtConfigBackupOperation

`func NewMgmtConfigBackupOperation(classId string, objectType string, ) *MgmtConfigBackupOperation`

NewMgmtConfigBackupOperation instantiates a new MgmtConfigBackupOperation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMgmtConfigBackupOperationWithDefaults

`func NewMgmtConfigBackupOperationWithDefaults() *MgmtConfigBackupOperation`

NewMgmtConfigBackupOperationWithDefaults instantiates a new MgmtConfigBackupOperation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *MgmtConfigBackupOperation) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *MgmtConfigBackupOperation) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *MgmtConfigBackupOperation) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *MgmtConfigBackupOperation) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *MgmtConfigBackupOperation) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *MgmtConfigBackupOperation) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetBackupAccount

`func (o *MgmtConfigBackupOperation) GetBackupAccount() string`

GetBackupAccount returns the BackupAccount field if non-nil, zero value otherwise.

### GetBackupAccountOk

`func (o *MgmtConfigBackupOperation) GetBackupAccountOk() (*string, bool)`

GetBackupAccountOk returns a tuple with the BackupAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupAccount

`func (o *MgmtConfigBackupOperation) SetBackupAccount(v string)`

SetBackupAccount sets BackupAccount field to given value.

### HasBackupAccount

`func (o *MgmtConfigBackupOperation) HasBackupAccount() bool`

HasBackupAccount returns a boolean if a field has been set.

### GetBackupAccountDomainGroupMoid

`func (o *MgmtConfigBackupOperation) GetBackupAccountDomainGroupMoid() string`

GetBackupAccountDomainGroupMoid returns the BackupAccountDomainGroupMoid field if non-nil, zero value otherwise.

### GetBackupAccountDomainGroupMoidOk

`func (o *MgmtConfigBackupOperation) GetBackupAccountDomainGroupMoidOk() (*string, bool)`

GetBackupAccountDomainGroupMoidOk returns a tuple with the BackupAccountDomainGroupMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupAccountDomainGroupMoid

`func (o *MgmtConfigBackupOperation) SetBackupAccountDomainGroupMoid(v string)`

SetBackupAccountDomainGroupMoid sets BackupAccountDomainGroupMoid field to given value.

### HasBackupAccountDomainGroupMoid

`func (o *MgmtConfigBackupOperation) HasBackupAccountDomainGroupMoid() bool`

HasBackupAccountDomainGroupMoid returns a boolean if a field has been set.

### GetCategoryExportData

`func (o *MgmtConfigBackupOperation) GetCategoryExportData() []MgmtCategoryExportData`

GetCategoryExportData returns the CategoryExportData field if non-nil, zero value otherwise.

### GetCategoryExportDataOk

`func (o *MgmtConfigBackupOperation) GetCategoryExportDataOk() (*[]MgmtCategoryExportData, bool)`

GetCategoryExportDataOk returns a tuple with the CategoryExportData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryExportData

`func (o *MgmtConfigBackupOperation) SetCategoryExportData(v []MgmtCategoryExportData)`

SetCategoryExportData sets CategoryExportData field to given value.

### HasCategoryExportData

`func (o *MgmtConfigBackupOperation) HasCategoryExportData() bool`

HasCategoryExportData returns a boolean if a field has been set.

### SetCategoryExportDataNil

`func (o *MgmtConfigBackupOperation) SetCategoryExportDataNil(b bool)`

 SetCategoryExportDataNil sets the value for CategoryExportData to be an explicit nil

### UnsetCategoryExportData
`func (o *MgmtConfigBackupOperation) UnsetCategoryExportData()`

UnsetCategoryExportData ensures that no value is present for CategoryExportData, not even an explicit nil
### GetDescription

`func (o *MgmtConfigBackupOperation) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MgmtConfigBackupOperation) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MgmtConfigBackupOperation) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MgmtConfigBackupOperation) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDestination

`func (o *MgmtConfigBackupOperation) GetDestination() string`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *MgmtConfigBackupOperation) GetDestinationOk() (*string, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *MgmtConfigBackupOperation) SetDestination(v string)`

SetDestination sets Destination field to given value.

### HasDestination

`func (o *MgmtConfigBackupOperation) HasDestination() bool`

HasDestination returns a boolean if a field has been set.

### GetEndTime

`func (o *MgmtConfigBackupOperation) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *MgmtConfigBackupOperation) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *MgmtConfigBackupOperation) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *MgmtConfigBackupOperation) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetIsAesKeySet

`func (o *MgmtConfigBackupOperation) GetIsAesKeySet() bool`

GetIsAesKeySet returns the IsAesKeySet field if non-nil, zero value otherwise.

### GetIsAesKeySetOk

`func (o *MgmtConfigBackupOperation) GetIsAesKeySetOk() (*bool, bool)`

GetIsAesKeySetOk returns a tuple with the IsAesKeySet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAesKeySet

`func (o *MgmtConfigBackupOperation) SetIsAesKeySet(v bool)`

SetIsAesKeySet sets IsAesKeySet field to given value.

### HasIsAesKeySet

`func (o *MgmtConfigBackupOperation) HasIsAesKeySet() bool`

HasIsAesKeySet returns a boolean if a field has been set.

### GetName

`func (o *MgmtConfigBackupOperation) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MgmtConfigBackupOperation) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MgmtConfigBackupOperation) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MgmtConfigBackupOperation) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOptions

`func (o *MgmtConfigBackupOperation) GetOptions() MgmtBackupOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *MgmtConfigBackupOperation) GetOptionsOk() (*MgmtBackupOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *MgmtConfigBackupOperation) SetOptions(v MgmtBackupOptions)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *MgmtConfigBackupOperation) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### SetOptionsNil

`func (o *MgmtConfigBackupOperation) SetOptionsNil(b bool)`

 SetOptionsNil sets the value for Options to be an explicit nil

### UnsetOptions
`func (o *MgmtConfigBackupOperation) UnsetOptions()`

UnsetOptions ensures that no value is present for Options, not even an explicit nil
### GetRetentionLock

`func (o *MgmtConfigBackupOperation) GetRetentionLock() bool`

GetRetentionLock returns the RetentionLock field if non-nil, zero value otherwise.

### GetRetentionLockOk

`func (o *MgmtConfigBackupOperation) GetRetentionLockOk() (*bool, bool)`

GetRetentionLockOk returns a tuple with the RetentionLock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionLock

`func (o *MgmtConfigBackupOperation) SetRetentionLock(v bool)`

SetRetentionLock sets RetentionLock field to given value.

### HasRetentionLock

`func (o *MgmtConfigBackupOperation) HasRetentionLock() bool`

HasRetentionLock returns a boolean if a field has been set.

### GetRolloverInitiated

`func (o *MgmtConfigBackupOperation) GetRolloverInitiated() bool`

GetRolloverInitiated returns the RolloverInitiated field if non-nil, zero value otherwise.

### GetRolloverInitiatedOk

`func (o *MgmtConfigBackupOperation) GetRolloverInitiatedOk() (*bool, bool)`

GetRolloverInitiatedOk returns a tuple with the RolloverInitiated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRolloverInitiated

`func (o *MgmtConfigBackupOperation) SetRolloverInitiated(v bool)`

SetRolloverInitiated sets RolloverInitiated field to given value.

### HasRolloverInitiated

`func (o *MgmtConfigBackupOperation) HasRolloverInitiated() bool`

HasRolloverInitiated returns a boolean if a field has been set.

### GetStartTime

`func (o *MgmtConfigBackupOperation) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *MgmtConfigBackupOperation) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *MgmtConfigBackupOperation) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *MgmtConfigBackupOperation) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetStatus

`func (o *MgmtConfigBackupOperation) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MgmtConfigBackupOperation) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MgmtConfigBackupOperation) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MgmtConfigBackupOperation) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSystemAdminTriggered

`func (o *MgmtConfigBackupOperation) GetSystemAdminTriggered() bool`

GetSystemAdminTriggered returns the SystemAdminTriggered field if non-nil, zero value otherwise.

### GetSystemAdminTriggeredOk

`func (o *MgmtConfigBackupOperation) GetSystemAdminTriggeredOk() (*bool, bool)`

GetSystemAdminTriggeredOk returns a tuple with the SystemAdminTriggered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemAdminTriggered

`func (o *MgmtConfigBackupOperation) SetSystemAdminTriggered(v bool)`

SetSystemAdminTriggered sets SystemAdminTriggered field to given value.

### HasSystemAdminTriggered

`func (o *MgmtConfigBackupOperation) HasSystemAdminTriggered() bool`

HasSystemAdminTriggered returns a boolean if a field has been set.

### GetUserPassword

`func (o *MgmtConfigBackupOperation) GetUserPassword() string`

GetUserPassword returns the UserPassword field if non-nil, zero value otherwise.

### GetUserPasswordOk

`func (o *MgmtConfigBackupOperation) GetUserPasswordOk() (*string, bool)`

GetUserPasswordOk returns a tuple with the UserPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserPassword

`func (o *MgmtConfigBackupOperation) SetUserPassword(v string)`

SetUserPassword sets UserPassword field to given value.

### HasUserPassword

`func (o *MgmtConfigBackupOperation) HasUserPassword() bool`

HasUserPassword returns a boolean if a field has been set.

### GetAccount

`func (o *MgmtConfigBackupOperation) GetAccount() IamAccountRelationship`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *MgmtConfigBackupOperation) GetAccountOk() (*IamAccountRelationship, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *MgmtConfigBackupOperation) SetAccount(v IamAccountRelationship)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *MgmtConfigBackupOperation) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### SetAccountNil

`func (o *MgmtConfigBackupOperation) SetAccountNil(b bool)`

 SetAccountNil sets the value for Account to be an explicit nil

### UnsetAccount
`func (o *MgmtConfigBackupOperation) UnsetAccount()`

UnsetAccount ensures that no value is present for Account, not even an explicit nil
### GetBackupInstance

`func (o *MgmtConfigBackupOperation) GetBackupInstance() MgmtConfigBackupInstanceRelationship`

GetBackupInstance returns the BackupInstance field if non-nil, zero value otherwise.

### GetBackupInstanceOk

`func (o *MgmtConfigBackupOperation) GetBackupInstanceOk() (*MgmtConfigBackupInstanceRelationship, bool)`

GetBackupInstanceOk returns a tuple with the BackupInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupInstance

`func (o *MgmtConfigBackupOperation) SetBackupInstance(v MgmtConfigBackupInstanceRelationship)`

SetBackupInstance sets BackupInstance field to given value.

### HasBackupInstance

`func (o *MgmtConfigBackupOperation) HasBackupInstance() bool`

HasBackupInstance returns a boolean if a field has been set.

### SetBackupInstanceNil

`func (o *MgmtConfigBackupOperation) SetBackupInstanceNil(b bool)`

 SetBackupInstanceNil sets the value for BackupInstance to be an explicit nil

### UnsetBackupInstance
`func (o *MgmtConfigBackupOperation) UnsetBackupInstance()`

UnsetBackupInstance ensures that no value is present for BackupInstance, not even an explicit nil
### GetOrgBackups

`func (o *MgmtConfigBackupOperation) GetOrgBackups() []MgmtOrgBackupOperationRelationship`

GetOrgBackups returns the OrgBackups field if non-nil, zero value otherwise.

### GetOrgBackupsOk

`func (o *MgmtConfigBackupOperation) GetOrgBackupsOk() (*[]MgmtOrgBackupOperationRelationship, bool)`

GetOrgBackupsOk returns a tuple with the OrgBackups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgBackups

`func (o *MgmtConfigBackupOperation) SetOrgBackups(v []MgmtOrgBackupOperationRelationship)`

SetOrgBackups sets OrgBackups field to given value.

### HasOrgBackups

`func (o *MgmtConfigBackupOperation) HasOrgBackups() bool`

HasOrgBackups returns a boolean if a field has been set.

### SetOrgBackupsNil

`func (o *MgmtConfigBackupOperation) SetOrgBackupsNil(b bool)`

 SetOrgBackupsNil sets the value for OrgBackups to be an explicit nil

### UnsetOrgBackups
`func (o *MgmtConfigBackupOperation) UnsetOrgBackups()`

UnsetOrgBackups ensures that no value is present for OrgBackups, not even an explicit nil
### GetUserBackupFile

`func (o *MgmtConfigBackupOperation) GetUserBackupFile() MgmtConfigBackupFileRelationship`

GetUserBackupFile returns the UserBackupFile field if non-nil, zero value otherwise.

### GetUserBackupFileOk

`func (o *MgmtConfigBackupOperation) GetUserBackupFileOk() (*MgmtConfigBackupFileRelationship, bool)`

GetUserBackupFileOk returns a tuple with the UserBackupFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserBackupFile

`func (o *MgmtConfigBackupOperation) SetUserBackupFile(v MgmtConfigBackupFileRelationship)`

SetUserBackupFile sets UserBackupFile field to given value.

### HasUserBackupFile

`func (o *MgmtConfigBackupOperation) HasUserBackupFile() bool`

HasUserBackupFile returns a boolean if a field has been set.

### SetUserBackupFileNil

`func (o *MgmtConfigBackupOperation) SetUserBackupFileNil(b bool)`

 SetUserBackupFileNil sets the value for UserBackupFile to be an explicit nil

### UnsetUserBackupFile
`func (o *MgmtConfigBackupOperation) UnsetUserBackupFile()`

UnsetUserBackupFile ensures that no value is present for UserBackupFile, not even an explicit nil
### GetWorkflowInfo

`func (o *MgmtConfigBackupOperation) GetWorkflowInfo() WorkflowWorkflowInfoRelationship`

GetWorkflowInfo returns the WorkflowInfo field if non-nil, zero value otherwise.

### GetWorkflowInfoOk

`func (o *MgmtConfigBackupOperation) GetWorkflowInfoOk() (*WorkflowWorkflowInfoRelationship, bool)`

GetWorkflowInfoOk returns a tuple with the WorkflowInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowInfo

`func (o *MgmtConfigBackupOperation) SetWorkflowInfo(v WorkflowWorkflowInfoRelationship)`

SetWorkflowInfo sets WorkflowInfo field to given value.

### HasWorkflowInfo

`func (o *MgmtConfigBackupOperation) HasWorkflowInfo() bool`

HasWorkflowInfo returns a boolean if a field has been set.

### SetWorkflowInfoNil

`func (o *MgmtConfigBackupOperation) SetWorkflowInfoNil(b bool)`

 SetWorkflowInfoNil sets the value for WorkflowInfo to be an explicit nil

### UnsetWorkflowInfo
`func (o *MgmtConfigBackupOperation) UnsetWorkflowInfo()`

UnsetWorkflowInfo ensures that no value is present for WorkflowInfo, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


