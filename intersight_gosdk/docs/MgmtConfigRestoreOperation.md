# MgmtConfigRestoreOperation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "mgmt.ConfigRestoreOperation"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "mgmt.ConfigRestoreOperation"]
**AcctCategoryImportData** | Pointer to [**[]MgmtCategoryImportData**](MgmtCategoryImportData.md) |  | [optional] 
**BackupAccountMoid** | Pointer to **string** | The moid of the account from which the backup was taken. | [optional] [readonly] 
**BackupName** | Pointer to **string** | Name of the backup that needs to be restored. | [optional] [readonly] 
**BackupSource** | Pointer to **string** | Indicates whether the backup was created locally or imported. * &#x60;Local&#x60; - A local Intersight location. * &#x60;Uploaded&#x60; - A local location where the backup file is uploaded as mgmt.ConfigBackupFile MO. Intersight creates the backup from the uploaded location when this location type is set. * &#x60;Remote&#x60; - A remote location hosted in the user&#39;s datacenter. | [optional] [readonly] [default to "Local"]
**EndTime** | Pointer to **time.Time** | End date and time of the restore process. | [optional] [readonly] 
**IgnoreSecureProperties** | Pointer to **bool** | If set to true, secure properties will not be decrypted, and empty values will be used for these properties. As a result, if a secure property field is mandatory, the MO will be marked as incomplete. | [optional] [default to false]
**IsAesKeySet** | Pointer to **bool** | Indicates whether the value of the &#39;aesKey&#39; property has been set. | [optional] [readonly] [default to false]
**Name** | Pointer to **string** | User provided identifier for the restore operation. | [optional] 
**Options** | Pointer to [**NullableMgmtRestoreOptions**](MgmtRestoreOptions.md) |  | [optional] 
**Source** | Pointer to **string** | The type of source from which the backup - local or remote - needs to be restored. * &#x60;Local&#x60; - A local Intersight location. * &#x60;Uploaded&#x60; - A local location where the backup file is uploaded as mgmt.ConfigBackupFile MO. Intersight creates the backup from the uploaded location when this location type is set. * &#x60;Remote&#x60; - A remote location hosted in the user&#39;s datacenter. | [optional] [default to "Local"]
**StartTime** | Pointer to **time.Time** | Start date and time of the restore process. | [optional] [readonly] 
**Status** | Pointer to **string** | Status of the restore operation. * &#x60;Scheduled&#x60; - Backup or restore process has been scheduled. * &#x60;NotInitiated&#x60; - Backup or restore process has not been initiated. * &#x60;InProgress&#x60; - Backup or restore process is in progress. * &#x60;Failed&#x60; - Backup or restore process has failed. * &#x60;Completed&#x60; - Backup or restore process has completed. * &#x60;NotComplete&#x60; - Backup or restore process has partially imported configuration. | [optional] [readonly] [default to "Scheduled"]
**Account** | Pointer to [**NullableIamAccountRelationship**](IamAccountRelationship.md) |  | [optional] 
**AcctImportJobs** | Pointer to [**[]BulkRequestRelationship**](BulkRequestRelationship.md) | An array of relationships to bulkRequest resources. | [optional] [readonly] 
**LocalSource** | Pointer to [**NullableMgmtConfigBackupInstanceRelationship**](MgmtConfigBackupInstanceRelationship.md) |  | [optional] 
**OrgImportJobs** | Pointer to [**[]BulkRequestRelationship**](BulkRequestRelationship.md) | An array of relationships to bulkRequest resources. | [optional] [readonly] 
**OrgRestores** | Pointer to [**[]MgmtOrgRestoreOperationRelationship**](MgmtOrgRestoreOperationRelationship.md) | An array of relationships to mgmtOrgRestoreOperation resources. | [optional] [readonly] 
**WorkflowInfo** | Pointer to [**NullableWorkflowWorkflowInfoRelationship**](WorkflowWorkflowInfoRelationship.md) |  | [optional] 

## Methods

### NewMgmtConfigRestoreOperation

`func NewMgmtConfigRestoreOperation(classId string, objectType string, ) *MgmtConfigRestoreOperation`

NewMgmtConfigRestoreOperation instantiates a new MgmtConfigRestoreOperation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMgmtConfigRestoreOperationWithDefaults

`func NewMgmtConfigRestoreOperationWithDefaults() *MgmtConfigRestoreOperation`

NewMgmtConfigRestoreOperationWithDefaults instantiates a new MgmtConfigRestoreOperation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *MgmtConfigRestoreOperation) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *MgmtConfigRestoreOperation) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *MgmtConfigRestoreOperation) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *MgmtConfigRestoreOperation) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *MgmtConfigRestoreOperation) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *MgmtConfigRestoreOperation) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAcctCategoryImportData

`func (o *MgmtConfigRestoreOperation) GetAcctCategoryImportData() []MgmtCategoryImportData`

GetAcctCategoryImportData returns the AcctCategoryImportData field if non-nil, zero value otherwise.

### GetAcctCategoryImportDataOk

`func (o *MgmtConfigRestoreOperation) GetAcctCategoryImportDataOk() (*[]MgmtCategoryImportData, bool)`

GetAcctCategoryImportDataOk returns a tuple with the AcctCategoryImportData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcctCategoryImportData

`func (o *MgmtConfigRestoreOperation) SetAcctCategoryImportData(v []MgmtCategoryImportData)`

SetAcctCategoryImportData sets AcctCategoryImportData field to given value.

### HasAcctCategoryImportData

`func (o *MgmtConfigRestoreOperation) HasAcctCategoryImportData() bool`

HasAcctCategoryImportData returns a boolean if a field has been set.

### SetAcctCategoryImportDataNil

`func (o *MgmtConfigRestoreOperation) SetAcctCategoryImportDataNil(b bool)`

 SetAcctCategoryImportDataNil sets the value for AcctCategoryImportData to be an explicit nil

### UnsetAcctCategoryImportData
`func (o *MgmtConfigRestoreOperation) UnsetAcctCategoryImportData()`

UnsetAcctCategoryImportData ensures that no value is present for AcctCategoryImportData, not even an explicit nil
### GetBackupAccountMoid

`func (o *MgmtConfigRestoreOperation) GetBackupAccountMoid() string`

GetBackupAccountMoid returns the BackupAccountMoid field if non-nil, zero value otherwise.

### GetBackupAccountMoidOk

`func (o *MgmtConfigRestoreOperation) GetBackupAccountMoidOk() (*string, bool)`

GetBackupAccountMoidOk returns a tuple with the BackupAccountMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupAccountMoid

`func (o *MgmtConfigRestoreOperation) SetBackupAccountMoid(v string)`

SetBackupAccountMoid sets BackupAccountMoid field to given value.

### HasBackupAccountMoid

`func (o *MgmtConfigRestoreOperation) HasBackupAccountMoid() bool`

HasBackupAccountMoid returns a boolean if a field has been set.

### GetBackupName

`func (o *MgmtConfigRestoreOperation) GetBackupName() string`

GetBackupName returns the BackupName field if non-nil, zero value otherwise.

### GetBackupNameOk

`func (o *MgmtConfigRestoreOperation) GetBackupNameOk() (*string, bool)`

GetBackupNameOk returns a tuple with the BackupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupName

`func (o *MgmtConfigRestoreOperation) SetBackupName(v string)`

SetBackupName sets BackupName field to given value.

### HasBackupName

`func (o *MgmtConfigRestoreOperation) HasBackupName() bool`

HasBackupName returns a boolean if a field has been set.

### GetBackupSource

`func (o *MgmtConfigRestoreOperation) GetBackupSource() string`

GetBackupSource returns the BackupSource field if non-nil, zero value otherwise.

### GetBackupSourceOk

`func (o *MgmtConfigRestoreOperation) GetBackupSourceOk() (*string, bool)`

GetBackupSourceOk returns a tuple with the BackupSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupSource

`func (o *MgmtConfigRestoreOperation) SetBackupSource(v string)`

SetBackupSource sets BackupSource field to given value.

### HasBackupSource

`func (o *MgmtConfigRestoreOperation) HasBackupSource() bool`

HasBackupSource returns a boolean if a field has been set.

### GetEndTime

`func (o *MgmtConfigRestoreOperation) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *MgmtConfigRestoreOperation) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *MgmtConfigRestoreOperation) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *MgmtConfigRestoreOperation) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetIgnoreSecureProperties

`func (o *MgmtConfigRestoreOperation) GetIgnoreSecureProperties() bool`

GetIgnoreSecureProperties returns the IgnoreSecureProperties field if non-nil, zero value otherwise.

### GetIgnoreSecurePropertiesOk

`func (o *MgmtConfigRestoreOperation) GetIgnoreSecurePropertiesOk() (*bool, bool)`

GetIgnoreSecurePropertiesOk returns a tuple with the IgnoreSecureProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreSecureProperties

`func (o *MgmtConfigRestoreOperation) SetIgnoreSecureProperties(v bool)`

SetIgnoreSecureProperties sets IgnoreSecureProperties field to given value.

### HasIgnoreSecureProperties

`func (o *MgmtConfigRestoreOperation) HasIgnoreSecureProperties() bool`

HasIgnoreSecureProperties returns a boolean if a field has been set.

### GetIsAesKeySet

`func (o *MgmtConfigRestoreOperation) GetIsAesKeySet() bool`

GetIsAesKeySet returns the IsAesKeySet field if non-nil, zero value otherwise.

### GetIsAesKeySetOk

`func (o *MgmtConfigRestoreOperation) GetIsAesKeySetOk() (*bool, bool)`

GetIsAesKeySetOk returns a tuple with the IsAesKeySet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAesKeySet

`func (o *MgmtConfigRestoreOperation) SetIsAesKeySet(v bool)`

SetIsAesKeySet sets IsAesKeySet field to given value.

### HasIsAesKeySet

`func (o *MgmtConfigRestoreOperation) HasIsAesKeySet() bool`

HasIsAesKeySet returns a boolean if a field has been set.

### GetName

`func (o *MgmtConfigRestoreOperation) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MgmtConfigRestoreOperation) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MgmtConfigRestoreOperation) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MgmtConfigRestoreOperation) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOptions

`func (o *MgmtConfigRestoreOperation) GetOptions() MgmtRestoreOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *MgmtConfigRestoreOperation) GetOptionsOk() (*MgmtRestoreOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *MgmtConfigRestoreOperation) SetOptions(v MgmtRestoreOptions)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *MgmtConfigRestoreOperation) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### SetOptionsNil

`func (o *MgmtConfigRestoreOperation) SetOptionsNil(b bool)`

 SetOptionsNil sets the value for Options to be an explicit nil

### UnsetOptions
`func (o *MgmtConfigRestoreOperation) UnsetOptions()`

UnsetOptions ensures that no value is present for Options, not even an explicit nil
### GetSource

`func (o *MgmtConfigRestoreOperation) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *MgmtConfigRestoreOperation) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *MgmtConfigRestoreOperation) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *MgmtConfigRestoreOperation) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetStartTime

`func (o *MgmtConfigRestoreOperation) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *MgmtConfigRestoreOperation) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *MgmtConfigRestoreOperation) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *MgmtConfigRestoreOperation) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetStatus

`func (o *MgmtConfigRestoreOperation) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MgmtConfigRestoreOperation) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MgmtConfigRestoreOperation) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MgmtConfigRestoreOperation) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetAccount

`func (o *MgmtConfigRestoreOperation) GetAccount() IamAccountRelationship`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *MgmtConfigRestoreOperation) GetAccountOk() (*IamAccountRelationship, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *MgmtConfigRestoreOperation) SetAccount(v IamAccountRelationship)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *MgmtConfigRestoreOperation) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### SetAccountNil

`func (o *MgmtConfigRestoreOperation) SetAccountNil(b bool)`

 SetAccountNil sets the value for Account to be an explicit nil

### UnsetAccount
`func (o *MgmtConfigRestoreOperation) UnsetAccount()`

UnsetAccount ensures that no value is present for Account, not even an explicit nil
### GetAcctImportJobs

`func (o *MgmtConfigRestoreOperation) GetAcctImportJobs() []BulkRequestRelationship`

GetAcctImportJobs returns the AcctImportJobs field if non-nil, zero value otherwise.

### GetAcctImportJobsOk

`func (o *MgmtConfigRestoreOperation) GetAcctImportJobsOk() (*[]BulkRequestRelationship, bool)`

GetAcctImportJobsOk returns a tuple with the AcctImportJobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcctImportJobs

`func (o *MgmtConfigRestoreOperation) SetAcctImportJobs(v []BulkRequestRelationship)`

SetAcctImportJobs sets AcctImportJobs field to given value.

### HasAcctImportJobs

`func (o *MgmtConfigRestoreOperation) HasAcctImportJobs() bool`

HasAcctImportJobs returns a boolean if a field has been set.

### SetAcctImportJobsNil

`func (o *MgmtConfigRestoreOperation) SetAcctImportJobsNil(b bool)`

 SetAcctImportJobsNil sets the value for AcctImportJobs to be an explicit nil

### UnsetAcctImportJobs
`func (o *MgmtConfigRestoreOperation) UnsetAcctImportJobs()`

UnsetAcctImportJobs ensures that no value is present for AcctImportJobs, not even an explicit nil
### GetLocalSource

`func (o *MgmtConfigRestoreOperation) GetLocalSource() MgmtConfigBackupInstanceRelationship`

GetLocalSource returns the LocalSource field if non-nil, zero value otherwise.

### GetLocalSourceOk

`func (o *MgmtConfigRestoreOperation) GetLocalSourceOk() (*MgmtConfigBackupInstanceRelationship, bool)`

GetLocalSourceOk returns a tuple with the LocalSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalSource

`func (o *MgmtConfigRestoreOperation) SetLocalSource(v MgmtConfigBackupInstanceRelationship)`

SetLocalSource sets LocalSource field to given value.

### HasLocalSource

`func (o *MgmtConfigRestoreOperation) HasLocalSource() bool`

HasLocalSource returns a boolean if a field has been set.

### SetLocalSourceNil

`func (o *MgmtConfigRestoreOperation) SetLocalSourceNil(b bool)`

 SetLocalSourceNil sets the value for LocalSource to be an explicit nil

### UnsetLocalSource
`func (o *MgmtConfigRestoreOperation) UnsetLocalSource()`

UnsetLocalSource ensures that no value is present for LocalSource, not even an explicit nil
### GetOrgImportJobs

`func (o *MgmtConfigRestoreOperation) GetOrgImportJobs() []BulkRequestRelationship`

GetOrgImportJobs returns the OrgImportJobs field if non-nil, zero value otherwise.

### GetOrgImportJobsOk

`func (o *MgmtConfigRestoreOperation) GetOrgImportJobsOk() (*[]BulkRequestRelationship, bool)`

GetOrgImportJobsOk returns a tuple with the OrgImportJobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgImportJobs

`func (o *MgmtConfigRestoreOperation) SetOrgImportJobs(v []BulkRequestRelationship)`

SetOrgImportJobs sets OrgImportJobs field to given value.

### HasOrgImportJobs

`func (o *MgmtConfigRestoreOperation) HasOrgImportJobs() bool`

HasOrgImportJobs returns a boolean if a field has been set.

### SetOrgImportJobsNil

`func (o *MgmtConfigRestoreOperation) SetOrgImportJobsNil(b bool)`

 SetOrgImportJobsNil sets the value for OrgImportJobs to be an explicit nil

### UnsetOrgImportJobs
`func (o *MgmtConfigRestoreOperation) UnsetOrgImportJobs()`

UnsetOrgImportJobs ensures that no value is present for OrgImportJobs, not even an explicit nil
### GetOrgRestores

`func (o *MgmtConfigRestoreOperation) GetOrgRestores() []MgmtOrgRestoreOperationRelationship`

GetOrgRestores returns the OrgRestores field if non-nil, zero value otherwise.

### GetOrgRestoresOk

`func (o *MgmtConfigRestoreOperation) GetOrgRestoresOk() (*[]MgmtOrgRestoreOperationRelationship, bool)`

GetOrgRestoresOk returns a tuple with the OrgRestores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgRestores

`func (o *MgmtConfigRestoreOperation) SetOrgRestores(v []MgmtOrgRestoreOperationRelationship)`

SetOrgRestores sets OrgRestores field to given value.

### HasOrgRestores

`func (o *MgmtConfigRestoreOperation) HasOrgRestores() bool`

HasOrgRestores returns a boolean if a field has been set.

### SetOrgRestoresNil

`func (o *MgmtConfigRestoreOperation) SetOrgRestoresNil(b bool)`

 SetOrgRestoresNil sets the value for OrgRestores to be an explicit nil

### UnsetOrgRestores
`func (o *MgmtConfigRestoreOperation) UnsetOrgRestores()`

UnsetOrgRestores ensures that no value is present for OrgRestores, not even an explicit nil
### GetWorkflowInfo

`func (o *MgmtConfigRestoreOperation) GetWorkflowInfo() WorkflowWorkflowInfoRelationship`

GetWorkflowInfo returns the WorkflowInfo field if non-nil, zero value otherwise.

### GetWorkflowInfoOk

`func (o *MgmtConfigRestoreOperation) GetWorkflowInfoOk() (*WorkflowWorkflowInfoRelationship, bool)`

GetWorkflowInfoOk returns a tuple with the WorkflowInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowInfo

`func (o *MgmtConfigRestoreOperation) SetWorkflowInfo(v WorkflowWorkflowInfoRelationship)`

SetWorkflowInfo sets WorkflowInfo field to given value.

### HasWorkflowInfo

`func (o *MgmtConfigRestoreOperation) HasWorkflowInfo() bool`

HasWorkflowInfo returns a boolean if a field has been set.

### SetWorkflowInfoNil

`func (o *MgmtConfigRestoreOperation) SetWorkflowInfoNil(b bool)`

 SetWorkflowInfoNil sets the value for WorkflowInfo to be an explicit nil

### UnsetWorkflowInfo
`func (o *MgmtConfigRestoreOperation) UnsetWorkflowInfo()`

UnsetWorkflowInfo ensures that no value is present for WorkflowInfo, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


