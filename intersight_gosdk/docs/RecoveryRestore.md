# RecoveryRestore

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "recovery.Restore"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "recovery.Restore"]
**ConfigParams** | Pointer to [**NullableRecoveryConfigParams**](RecoveryConfigParams.md) |  | [optional] 
**BackupInfo** | Pointer to [**NullableRecoveryAbstractBackupInfoRelationship**](RecoveryAbstractBackupInfoRelationship.md) |  | [optional] 
**Device** | Pointer to [**NullableAssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 
**Organization** | Pointer to [**NullableOrganizationOrganizationRelationship**](OrganizationOrganizationRelationship.md) |  | [optional] 
**Workflow** | Pointer to [**NullableWorkflowWorkflowInfoRelationship**](WorkflowWorkflowInfoRelationship.md) |  | [optional] 

## Methods

### NewRecoveryRestore

`func NewRecoveryRestore(classId string, objectType string, ) *RecoveryRestore`

NewRecoveryRestore instantiates a new RecoveryRestore object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoveryRestoreWithDefaults

`func NewRecoveryRestoreWithDefaults() *RecoveryRestore`

NewRecoveryRestoreWithDefaults instantiates a new RecoveryRestore object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *RecoveryRestore) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *RecoveryRestore) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *RecoveryRestore) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *RecoveryRestore) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *RecoveryRestore) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *RecoveryRestore) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


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

### SetConfigParamsNil

`func (o *RecoveryRestore) SetConfigParamsNil(b bool)`

 SetConfigParamsNil sets the value for ConfigParams to be an explicit nil

### UnsetConfigParams
`func (o *RecoveryRestore) UnsetConfigParams()`

UnsetConfigParams ensures that no value is present for ConfigParams, not even an explicit nil
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

### SetBackupInfoNil

`func (o *RecoveryRestore) SetBackupInfoNil(b bool)`

 SetBackupInfoNil sets the value for BackupInfo to be an explicit nil

### UnsetBackupInfo
`func (o *RecoveryRestore) UnsetBackupInfo()`

UnsetBackupInfo ensures that no value is present for BackupInfo, not even an explicit nil
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

### SetDeviceNil

`func (o *RecoveryRestore) SetDeviceNil(b bool)`

 SetDeviceNil sets the value for Device to be an explicit nil

### UnsetDevice
`func (o *RecoveryRestore) UnsetDevice()`

UnsetDevice ensures that no value is present for Device, not even an explicit nil
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

### SetOrganizationNil

`func (o *RecoveryRestore) SetOrganizationNil(b bool)`

 SetOrganizationNil sets the value for Organization to be an explicit nil

### UnsetOrganization
`func (o *RecoveryRestore) UnsetOrganization()`

UnsetOrganization ensures that no value is present for Organization, not even an explicit nil
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

### SetWorkflowNil

`func (o *RecoveryRestore) SetWorkflowNil(b bool)`

 SetWorkflowNil sets the value for Workflow to be an explicit nil

### UnsetWorkflow
`func (o *RecoveryRestore) UnsetWorkflow()`

UnsetWorkflow ensures that no value is present for Workflow, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


