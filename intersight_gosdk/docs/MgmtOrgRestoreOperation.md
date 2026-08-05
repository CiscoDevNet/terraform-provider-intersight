# MgmtOrgRestoreOperation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "mgmt.OrgRestoreOperation"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "mgmt.OrgRestoreOperation"]
**CategoryImportData** | Pointer to [**[]MgmtCategoryImportData**](MgmtCategoryImportData.md) |  | [optional] 
**OrgName** | Pointer to **string** | Name of the organization for which the restore operation is being performed. | [optional] [readonly] 
**SequenceNumber** | Pointer to **int64** | The sequence number that determines the order in which organization configurations will be processed. | [optional] [readonly] 
**SharedOrgs** | Pointer to **[]string** |  | [optional] 
**Status** | Pointer to **string** | The current status of the restore operation. * &#x60;Scheduled&#x60; - Backup or restore process has been scheduled. * &#x60;NotInitiated&#x60; - Backup or restore process has not been initiated. * &#x60;InProgress&#x60; - Backup or restore process is in progress. * &#x60;Failed&#x60; - Backup or restore process has failed. * &#x60;Completed&#x60; - Backup or restore process has completed. * &#x60;NotComplete&#x60; - Backup or restore process has partially imported configuration. | [optional] [readonly] [default to "Scheduled"]
**ImportJobs** | Pointer to [**[]BulkRequestRelationship**](BulkRequestRelationship.md) | An array of relationships to bulkRequest resources. | [optional] [readonly] 
**Organization** | Pointer to [**NullableOrganizationOrganizationRelationship**](OrganizationOrganizationRelationship.md) |  | [optional] 
**RestoreOperation** | Pointer to [**NullableMgmtConfigRestoreOperationRelationship**](MgmtConfigRestoreOperationRelationship.md) |  | [optional] 

## Methods

### NewMgmtOrgRestoreOperation

`func NewMgmtOrgRestoreOperation(classId string, objectType string, ) *MgmtOrgRestoreOperation`

NewMgmtOrgRestoreOperation instantiates a new MgmtOrgRestoreOperation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMgmtOrgRestoreOperationWithDefaults

`func NewMgmtOrgRestoreOperationWithDefaults() *MgmtOrgRestoreOperation`

NewMgmtOrgRestoreOperationWithDefaults instantiates a new MgmtOrgRestoreOperation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *MgmtOrgRestoreOperation) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *MgmtOrgRestoreOperation) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *MgmtOrgRestoreOperation) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *MgmtOrgRestoreOperation) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *MgmtOrgRestoreOperation) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *MgmtOrgRestoreOperation) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCategoryImportData

`func (o *MgmtOrgRestoreOperation) GetCategoryImportData() []MgmtCategoryImportData`

GetCategoryImportData returns the CategoryImportData field if non-nil, zero value otherwise.

### GetCategoryImportDataOk

`func (o *MgmtOrgRestoreOperation) GetCategoryImportDataOk() (*[]MgmtCategoryImportData, bool)`

GetCategoryImportDataOk returns a tuple with the CategoryImportData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryImportData

`func (o *MgmtOrgRestoreOperation) SetCategoryImportData(v []MgmtCategoryImportData)`

SetCategoryImportData sets CategoryImportData field to given value.

### HasCategoryImportData

`func (o *MgmtOrgRestoreOperation) HasCategoryImportData() bool`

HasCategoryImportData returns a boolean if a field has been set.

### SetCategoryImportDataNil

`func (o *MgmtOrgRestoreOperation) SetCategoryImportDataNil(b bool)`

 SetCategoryImportDataNil sets the value for CategoryImportData to be an explicit nil

### UnsetCategoryImportData
`func (o *MgmtOrgRestoreOperation) UnsetCategoryImportData()`

UnsetCategoryImportData ensures that no value is present for CategoryImportData, not even an explicit nil
### GetOrgName

`func (o *MgmtOrgRestoreOperation) GetOrgName() string`

GetOrgName returns the OrgName field if non-nil, zero value otherwise.

### GetOrgNameOk

`func (o *MgmtOrgRestoreOperation) GetOrgNameOk() (*string, bool)`

GetOrgNameOk returns a tuple with the OrgName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgName

`func (o *MgmtOrgRestoreOperation) SetOrgName(v string)`

SetOrgName sets OrgName field to given value.

### HasOrgName

`func (o *MgmtOrgRestoreOperation) HasOrgName() bool`

HasOrgName returns a boolean if a field has been set.

### GetSequenceNumber

`func (o *MgmtOrgRestoreOperation) GetSequenceNumber() int64`

GetSequenceNumber returns the SequenceNumber field if non-nil, zero value otherwise.

### GetSequenceNumberOk

`func (o *MgmtOrgRestoreOperation) GetSequenceNumberOk() (*int64, bool)`

GetSequenceNumberOk returns a tuple with the SequenceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequenceNumber

`func (o *MgmtOrgRestoreOperation) SetSequenceNumber(v int64)`

SetSequenceNumber sets SequenceNumber field to given value.

### HasSequenceNumber

`func (o *MgmtOrgRestoreOperation) HasSequenceNumber() bool`

HasSequenceNumber returns a boolean if a field has been set.

### GetSharedOrgs

`func (o *MgmtOrgRestoreOperation) GetSharedOrgs() []string`

GetSharedOrgs returns the SharedOrgs field if non-nil, zero value otherwise.

### GetSharedOrgsOk

`func (o *MgmtOrgRestoreOperation) GetSharedOrgsOk() (*[]string, bool)`

GetSharedOrgsOk returns a tuple with the SharedOrgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedOrgs

`func (o *MgmtOrgRestoreOperation) SetSharedOrgs(v []string)`

SetSharedOrgs sets SharedOrgs field to given value.

### HasSharedOrgs

`func (o *MgmtOrgRestoreOperation) HasSharedOrgs() bool`

HasSharedOrgs returns a boolean if a field has been set.

### SetSharedOrgsNil

`func (o *MgmtOrgRestoreOperation) SetSharedOrgsNil(b bool)`

 SetSharedOrgsNil sets the value for SharedOrgs to be an explicit nil

### UnsetSharedOrgs
`func (o *MgmtOrgRestoreOperation) UnsetSharedOrgs()`

UnsetSharedOrgs ensures that no value is present for SharedOrgs, not even an explicit nil
### GetStatus

`func (o *MgmtOrgRestoreOperation) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MgmtOrgRestoreOperation) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MgmtOrgRestoreOperation) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MgmtOrgRestoreOperation) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetImportJobs

`func (o *MgmtOrgRestoreOperation) GetImportJobs() []BulkRequestRelationship`

GetImportJobs returns the ImportJobs field if non-nil, zero value otherwise.

### GetImportJobsOk

`func (o *MgmtOrgRestoreOperation) GetImportJobsOk() (*[]BulkRequestRelationship, bool)`

GetImportJobsOk returns a tuple with the ImportJobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportJobs

`func (o *MgmtOrgRestoreOperation) SetImportJobs(v []BulkRequestRelationship)`

SetImportJobs sets ImportJobs field to given value.

### HasImportJobs

`func (o *MgmtOrgRestoreOperation) HasImportJobs() bool`

HasImportJobs returns a boolean if a field has been set.

### SetImportJobsNil

`func (o *MgmtOrgRestoreOperation) SetImportJobsNil(b bool)`

 SetImportJobsNil sets the value for ImportJobs to be an explicit nil

### UnsetImportJobs
`func (o *MgmtOrgRestoreOperation) UnsetImportJobs()`

UnsetImportJobs ensures that no value is present for ImportJobs, not even an explicit nil
### GetOrganization

`func (o *MgmtOrgRestoreOperation) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *MgmtOrgRestoreOperation) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *MgmtOrgRestoreOperation) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *MgmtOrgRestoreOperation) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### SetOrganizationNil

`func (o *MgmtOrgRestoreOperation) SetOrganizationNil(b bool)`

 SetOrganizationNil sets the value for Organization to be an explicit nil

### UnsetOrganization
`func (o *MgmtOrgRestoreOperation) UnsetOrganization()`

UnsetOrganization ensures that no value is present for Organization, not even an explicit nil
### GetRestoreOperation

`func (o *MgmtOrgRestoreOperation) GetRestoreOperation() MgmtConfigRestoreOperationRelationship`

GetRestoreOperation returns the RestoreOperation field if non-nil, zero value otherwise.

### GetRestoreOperationOk

`func (o *MgmtOrgRestoreOperation) GetRestoreOperationOk() (*MgmtConfigRestoreOperationRelationship, bool)`

GetRestoreOperationOk returns a tuple with the RestoreOperation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreOperation

`func (o *MgmtOrgRestoreOperation) SetRestoreOperation(v MgmtConfigRestoreOperationRelationship)`

SetRestoreOperation sets RestoreOperation field to given value.

### HasRestoreOperation

`func (o *MgmtOrgRestoreOperation) HasRestoreOperation() bool`

HasRestoreOperation returns a boolean if a field has been set.

### SetRestoreOperationNil

`func (o *MgmtOrgRestoreOperation) SetRestoreOperationNil(b bool)`

 SetRestoreOperationNil sets the value for RestoreOperation to be an explicit nil

### UnsetRestoreOperation
`func (o *MgmtOrgRestoreOperation) UnsetRestoreOperation()`

UnsetRestoreOperation ensures that no value is present for RestoreOperation, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


