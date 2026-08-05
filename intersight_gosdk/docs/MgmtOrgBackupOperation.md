# MgmtOrgBackupOperation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "mgmt.OrgBackupOperation"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "mgmt.OrgBackupOperation"]
**CategoryExportData** | Pointer to [**[]MgmtCategoryExportData**](MgmtCategoryExportData.md) |  | [optional] 
**OrgName** | Pointer to **string** | Name of the organization to be backed up. | [optional] [readonly] 
**PreserveIdentities** | Pointer to **bool** | The flag set by the user during a configuration backup to preserve static or dynamic IDs assigned to an export item. | [optional] 
**Account** | Pointer to [**NullableIamAccountRelationship**](IamAccountRelationship.md) |  | [optional] 
**BackupOperation** | Pointer to [**NullableMgmtConfigBackupOperationRelationship**](MgmtConfigBackupOperationRelationship.md) |  | [optional] 
**OrgExportJob** | Pointer to [**NullableBulkExportRelationship**](BulkExportRelationship.md) |  | [optional] 
**Organization** | Pointer to [**NullableOrganizationOrganizationRelationship**](OrganizationOrganizationRelationship.md) |  | [optional] 

## Methods

### NewMgmtOrgBackupOperation

`func NewMgmtOrgBackupOperation(classId string, objectType string, ) *MgmtOrgBackupOperation`

NewMgmtOrgBackupOperation instantiates a new MgmtOrgBackupOperation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMgmtOrgBackupOperationWithDefaults

`func NewMgmtOrgBackupOperationWithDefaults() *MgmtOrgBackupOperation`

NewMgmtOrgBackupOperationWithDefaults instantiates a new MgmtOrgBackupOperation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *MgmtOrgBackupOperation) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *MgmtOrgBackupOperation) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *MgmtOrgBackupOperation) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *MgmtOrgBackupOperation) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *MgmtOrgBackupOperation) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *MgmtOrgBackupOperation) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCategoryExportData

`func (o *MgmtOrgBackupOperation) GetCategoryExportData() []MgmtCategoryExportData`

GetCategoryExportData returns the CategoryExportData field if non-nil, zero value otherwise.

### GetCategoryExportDataOk

`func (o *MgmtOrgBackupOperation) GetCategoryExportDataOk() (*[]MgmtCategoryExportData, bool)`

GetCategoryExportDataOk returns a tuple with the CategoryExportData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryExportData

`func (o *MgmtOrgBackupOperation) SetCategoryExportData(v []MgmtCategoryExportData)`

SetCategoryExportData sets CategoryExportData field to given value.

### HasCategoryExportData

`func (o *MgmtOrgBackupOperation) HasCategoryExportData() bool`

HasCategoryExportData returns a boolean if a field has been set.

### SetCategoryExportDataNil

`func (o *MgmtOrgBackupOperation) SetCategoryExportDataNil(b bool)`

 SetCategoryExportDataNil sets the value for CategoryExportData to be an explicit nil

### UnsetCategoryExportData
`func (o *MgmtOrgBackupOperation) UnsetCategoryExportData()`

UnsetCategoryExportData ensures that no value is present for CategoryExportData, not even an explicit nil
### GetOrgName

`func (o *MgmtOrgBackupOperation) GetOrgName() string`

GetOrgName returns the OrgName field if non-nil, zero value otherwise.

### GetOrgNameOk

`func (o *MgmtOrgBackupOperation) GetOrgNameOk() (*string, bool)`

GetOrgNameOk returns a tuple with the OrgName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgName

`func (o *MgmtOrgBackupOperation) SetOrgName(v string)`

SetOrgName sets OrgName field to given value.

### HasOrgName

`func (o *MgmtOrgBackupOperation) HasOrgName() bool`

HasOrgName returns a boolean if a field has been set.

### GetPreserveIdentities

`func (o *MgmtOrgBackupOperation) GetPreserveIdentities() bool`

GetPreserveIdentities returns the PreserveIdentities field if non-nil, zero value otherwise.

### GetPreserveIdentitiesOk

`func (o *MgmtOrgBackupOperation) GetPreserveIdentitiesOk() (*bool, bool)`

GetPreserveIdentitiesOk returns a tuple with the PreserveIdentities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreserveIdentities

`func (o *MgmtOrgBackupOperation) SetPreserveIdentities(v bool)`

SetPreserveIdentities sets PreserveIdentities field to given value.

### HasPreserveIdentities

`func (o *MgmtOrgBackupOperation) HasPreserveIdentities() bool`

HasPreserveIdentities returns a boolean if a field has been set.

### GetAccount

`func (o *MgmtOrgBackupOperation) GetAccount() IamAccountRelationship`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *MgmtOrgBackupOperation) GetAccountOk() (*IamAccountRelationship, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *MgmtOrgBackupOperation) SetAccount(v IamAccountRelationship)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *MgmtOrgBackupOperation) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### SetAccountNil

`func (o *MgmtOrgBackupOperation) SetAccountNil(b bool)`

 SetAccountNil sets the value for Account to be an explicit nil

### UnsetAccount
`func (o *MgmtOrgBackupOperation) UnsetAccount()`

UnsetAccount ensures that no value is present for Account, not even an explicit nil
### GetBackupOperation

`func (o *MgmtOrgBackupOperation) GetBackupOperation() MgmtConfigBackupOperationRelationship`

GetBackupOperation returns the BackupOperation field if non-nil, zero value otherwise.

### GetBackupOperationOk

`func (o *MgmtOrgBackupOperation) GetBackupOperationOk() (*MgmtConfigBackupOperationRelationship, bool)`

GetBackupOperationOk returns a tuple with the BackupOperation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupOperation

`func (o *MgmtOrgBackupOperation) SetBackupOperation(v MgmtConfigBackupOperationRelationship)`

SetBackupOperation sets BackupOperation field to given value.

### HasBackupOperation

`func (o *MgmtOrgBackupOperation) HasBackupOperation() bool`

HasBackupOperation returns a boolean if a field has been set.

### SetBackupOperationNil

`func (o *MgmtOrgBackupOperation) SetBackupOperationNil(b bool)`

 SetBackupOperationNil sets the value for BackupOperation to be an explicit nil

### UnsetBackupOperation
`func (o *MgmtOrgBackupOperation) UnsetBackupOperation()`

UnsetBackupOperation ensures that no value is present for BackupOperation, not even an explicit nil
### GetOrgExportJob

`func (o *MgmtOrgBackupOperation) GetOrgExportJob() BulkExportRelationship`

GetOrgExportJob returns the OrgExportJob field if non-nil, zero value otherwise.

### GetOrgExportJobOk

`func (o *MgmtOrgBackupOperation) GetOrgExportJobOk() (*BulkExportRelationship, bool)`

GetOrgExportJobOk returns a tuple with the OrgExportJob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgExportJob

`func (o *MgmtOrgBackupOperation) SetOrgExportJob(v BulkExportRelationship)`

SetOrgExportJob sets OrgExportJob field to given value.

### HasOrgExportJob

`func (o *MgmtOrgBackupOperation) HasOrgExportJob() bool`

HasOrgExportJob returns a boolean if a field has been set.

### SetOrgExportJobNil

`func (o *MgmtOrgBackupOperation) SetOrgExportJobNil(b bool)`

 SetOrgExportJobNil sets the value for OrgExportJob to be an explicit nil

### UnsetOrgExportJob
`func (o *MgmtOrgBackupOperation) UnsetOrgExportJob()`

UnsetOrgExportJob ensures that no value is present for OrgExportJob, not even an explicit nil
### GetOrganization

`func (o *MgmtOrgBackupOperation) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *MgmtOrgBackupOperation) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *MgmtOrgBackupOperation) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *MgmtOrgBackupOperation) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### SetOrganizationNil

`func (o *MgmtOrgBackupOperation) SetOrganizationNil(b bool)`

 SetOrganizationNil sets the value for Organization to be an explicit nil

### UnsetOrganization
`func (o *MgmtOrgBackupOperation) UnsetOrganization()`

UnsetOrganization ensures that no value is present for Organization, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


