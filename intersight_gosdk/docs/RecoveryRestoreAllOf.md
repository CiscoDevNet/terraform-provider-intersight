# RecoveryRestoreAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "recovery.Restore"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "recovery.Restore"]
**ConfigParams** | Pointer to [**NullableRecoveryConfigParams**](RecoveryConfigParams.md) |  | [optional] 
**BackupInfo** | Pointer to [**RecoveryAbstractBackupInfoRelationship**](RecoveryAbstractBackupInfoRelationship.md) |  | [optional] 
**Device** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](OrganizationOrganizationRelationship.md) |  | [optional] 
**Workflow** | Pointer to [**WorkflowWorkflowInfoRelationship**](WorkflowWorkflowInfoRelationship.md) |  | [optional] 

## Methods

### NewRecoveryRestoreAllOf

`func NewRecoveryRestoreAllOf(classId string, objectType string, ) *RecoveryRestoreAllOf`

NewRecoveryRestoreAllOf instantiates a new RecoveryRestoreAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoveryRestoreAllOfWithDefaults

`func NewRecoveryRestoreAllOfWithDefaults() *RecoveryRestoreAllOf`

NewRecoveryRestoreAllOfWithDefaults instantiates a new RecoveryRestoreAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *RecoveryRestoreAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *RecoveryRestoreAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *RecoveryRestoreAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *RecoveryRestoreAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *RecoveryRestoreAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *RecoveryRestoreAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetConfigParams

`func (o *RecoveryRestoreAllOf) GetConfigParams() RecoveryConfigParams`

GetConfigParams returns the ConfigParams field if non-nil, zero value otherwise.

### GetConfigParamsOk

`func (o *RecoveryRestoreAllOf) GetConfigParamsOk() (*RecoveryConfigParams, bool)`

GetConfigParamsOk returns a tuple with the ConfigParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigParams

`func (o *RecoveryRestoreAllOf) SetConfigParams(v RecoveryConfigParams)`

SetConfigParams sets ConfigParams field to given value.

### HasConfigParams

`func (o *RecoveryRestoreAllOf) HasConfigParams() bool`

HasConfigParams returns a boolean if a field has been set.

### SetConfigParamsNil

`func (o *RecoveryRestoreAllOf) SetConfigParamsNil(b bool)`

 SetConfigParamsNil sets the value for ConfigParams to be an explicit nil

### UnsetConfigParams
`func (o *RecoveryRestoreAllOf) UnsetConfigParams()`

UnsetConfigParams ensures that no value is present for ConfigParams, not even an explicit nil
### GetBackupInfo

`func (o *RecoveryRestoreAllOf) GetBackupInfo() RecoveryAbstractBackupInfoRelationship`

GetBackupInfo returns the BackupInfo field if non-nil, zero value otherwise.

### GetBackupInfoOk

`func (o *RecoveryRestoreAllOf) GetBackupInfoOk() (*RecoveryAbstractBackupInfoRelationship, bool)`

GetBackupInfoOk returns a tuple with the BackupInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupInfo

`func (o *RecoveryRestoreAllOf) SetBackupInfo(v RecoveryAbstractBackupInfoRelationship)`

SetBackupInfo sets BackupInfo field to given value.

### HasBackupInfo

`func (o *RecoveryRestoreAllOf) HasBackupInfo() bool`

HasBackupInfo returns a boolean if a field has been set.

### GetDevice

`func (o *RecoveryRestoreAllOf) GetDevice() AssetDeviceRegistrationRelationship`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *RecoveryRestoreAllOf) GetDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *RecoveryRestoreAllOf) SetDevice(v AssetDeviceRegistrationRelationship)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *RecoveryRestoreAllOf) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### GetOrganization

`func (o *RecoveryRestoreAllOf) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *RecoveryRestoreAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *RecoveryRestoreAllOf) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *RecoveryRestoreAllOf) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetWorkflow

`func (o *RecoveryRestoreAllOf) GetWorkflow() WorkflowWorkflowInfoRelationship`

GetWorkflow returns the Workflow field if non-nil, zero value otherwise.

### GetWorkflowOk

`func (o *RecoveryRestoreAllOf) GetWorkflowOk() (*WorkflowWorkflowInfoRelationship, bool)`

GetWorkflowOk returns a tuple with the Workflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflow

`func (o *RecoveryRestoreAllOf) SetWorkflow(v WorkflowWorkflowInfoRelationship)`

SetWorkflow sets Workflow field to given value.

### HasWorkflow

`func (o *RecoveryRestoreAllOf) HasWorkflow() bool`

HasWorkflow returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


