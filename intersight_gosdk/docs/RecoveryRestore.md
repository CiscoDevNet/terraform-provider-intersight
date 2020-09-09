# RecoveryRestore

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfigParams** | Pointer to [**RecoveryConfigParams**](recovery.ConfigParams.md) |  | [optional] 
**BackupInfo** | Pointer to [**RecoveryAbstractBackupInfoRelationship**](recovery.AbstractBackupInfo.Relationship.md) |  | [optional] 
**Device** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](organization.Organization.Relationship.md) |  | [optional] 
**Workflow** | Pointer to [**WorkflowWorkflowInfoRelationship**](workflow.WorkflowInfo.Relationship.md) |  | [optional] 

## Methods

### NewRecoveryRestore

`func NewRecoveryRestore() *RecoveryRestore`

NewRecoveryRestore instantiates a new RecoveryRestore object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoveryRestoreWithDefaults

`func NewRecoveryRestoreWithDefaults() *RecoveryRestore`

NewRecoveryRestoreWithDefaults instantiates a new RecoveryRestore object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfigParams

`func (o *RecoveryRestore) GetConfigParams() RecoveryConfigParams`

GetConfigParams returns the ConfigParams field if non-nil, zero value otherwise.

### GetConfigParamsOk

`func (o *RecoveryRestore) GetConfigParamsOk() (*RecoveryConfigParams, bool)`

GetConfigParamsOk returns a tuple with the ConfigParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigParams

`func (o *RecoveryRestore) SetConfigParams(v RecoveryConfigParams)`

SetConfigParams sets ConfigParams field to given value.

### HasConfigParams

`func (o *RecoveryRestore) HasConfigParams() bool`

HasConfigParams returns a boolean if a field has been set.

### GetBackupInfo

`func (o *RecoveryRestore) GetBackupInfo() RecoveryAbstractBackupInfoRelationship`

GetBackupInfo returns the BackupInfo field if non-nil, zero value otherwise.

### GetBackupInfoOk

`func (o *RecoveryRestore) GetBackupInfoOk() (*RecoveryAbstractBackupInfoRelationship, bool)`

GetBackupInfoOk returns a tuple with the BackupInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupInfo

`func (o *RecoveryRestore) SetBackupInfo(v RecoveryAbstractBackupInfoRelationship)`

SetBackupInfo sets BackupInfo field to given value.

### HasBackupInfo

`func (o *RecoveryRestore) HasBackupInfo() bool`

HasBackupInfo returns a boolean if a field has been set.

### GetDevice

`func (o *RecoveryRestore) GetDevice() AssetDeviceRegistrationRelationship`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *RecoveryRestore) GetDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *RecoveryRestore) SetDevice(v AssetDeviceRegistrationRelationship)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *RecoveryRestore) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### GetOrganization

`func (o *RecoveryRestore) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *RecoveryRestore) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *RecoveryRestore) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *RecoveryRestore) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetWorkflow

`func (o *RecoveryRestore) GetWorkflow() WorkflowWorkflowInfoRelationship`

GetWorkflow returns the Workflow field if non-nil, zero value otherwise.

### GetWorkflowOk

`func (o *RecoveryRestore) GetWorkflowOk() (*WorkflowWorkflowInfoRelationship, bool)`

GetWorkflowOk returns a tuple with the Workflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflow

`func (o *RecoveryRestore) SetWorkflow(v WorkflowWorkflowInfoRelationship)`

SetWorkflow sets Workflow field to given value.

### HasWorkflow

`func (o *RecoveryRestore) HasWorkflow() bool`

HasWorkflow returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


