# HyperflexClusterBackupPolicyInventory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hyperflex.ClusterBackupPolicyInventory"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hyperflex.ClusterBackupPolicyInventory"]
**Action** | Pointer to **string** | Validate, Deploy, Prepare, Commit, or Abort the Backup Policy. Allowed values are \&quot;VALIDATE\&quot;, \&quot;DEPLOY\&quot;, \&quot;PREPARE\&quot;, \&quot;COMMIT\&quot;, \&quot;ABORT\&quot;. * &#x60;VALIDATE&#x60; - Check for problems in policy request without deploying. * &#x60;DEPLOY&#x60; - Remove the policy.  Only allowed when cleanup field is true. * &#x60;PREPARE&#x60; - Prepare the policy for subsequent \&quot;COMMIT\&quot; or \&quot;ABORT\&quot;.  Only allowed when cleanup field is false. * &#x60;COMMIT&#x60; - Commit the prepared policy.  Only allowed when cleanup field is false. * &#x60;ABORT&#x60; - Abort the prepared policy.  Only allowed when cleanup field is false. | [optional] [readonly] [default to "VALIDATE"]
**Cleanup** | Pointer to **bool** | If true, remove the policy. Otherwise proceed with the specified policy action. | [optional] [readonly] 
**IsSource** | Pointer to **bool** | Indicates if the HyperFlex Cluster is the source cluster or the target cluster. When set to true, the HyperFlex Cluster is the source for VM backups. When set to false, the HyperFlex Cluster is the target cluster where VM backups are saved. | [optional] [readonly] 
**JobDetails** | Pointer to **string** | Details from associated HyperFlex job execution. | [optional] [readonly] 
**JobExceptionMessage** | Pointer to **string** | Job Exception message from associated HyperFlex job execution. | [optional] [readonly] 
**JobId** | Pointer to **string** | Job ID of associated HyperFlex job. | [optional] [readonly] 
**JobState** | Pointer to **string** | State of the associated HyperFlex job. When present possible values are \&quot;RUNNING\&quot;, \&quot;COMPLETED\&quot; or \&quot;EXCEPTION\&quot;. * &#x60;RUNNING&#x60; - HyperFlex policy job is currently \&quot;RUNNING\&quot;. * &#x60;COMPLETED&#x60; - HyperFlex policy job completed successfully. * &#x60;EXCEPTION&#x60; - HyperFlex policy job failed. | [optional] [readonly] [default to "RUNNING"]
**PolicyMoid** | Pointer to **string** | Intersight HyperFlex Cluster Backup Policy MOID. | [optional] [readonly] 
**RequestId** | Pointer to **string** | Unique request ID allowing retry of the same logical request following a transient communication failure. | [optional] [readonly] 
**Settings** | Pointer to [**NullableHyperflexBackupPolicySettings**](HyperflexBackupPolicySettings.md) |  | [optional] 
**SourceUuid** | Pointer to **string** | UUID of the source HyperFlex Cluster. | [optional] [readonly] 
**TargetUuid** | Pointer to **string** | UUID of the target HyperFlex Cluster. | [optional] [readonly] 
**Version** | Pointer to **int64** | Version of the Backup Policy. | [optional] [readonly] 
**SrcCluster** | Pointer to [**HyperflexClusterRelationship**](HyperflexClusterRelationship.md) |  | [optional] 
**TgtCluster** | Pointer to [**HyperflexClusterRelationship**](HyperflexClusterRelationship.md) |  | [optional] 

## Methods

### NewHyperflexClusterBackupPolicyInventory

`func NewHyperflexClusterBackupPolicyInventory(classId string, objectType string, ) *HyperflexClusterBackupPolicyInventory`

NewHyperflexClusterBackupPolicyInventory instantiates a new HyperflexClusterBackupPolicyInventory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexClusterBackupPolicyInventoryWithDefaults

`func NewHyperflexClusterBackupPolicyInventoryWithDefaults() *HyperflexClusterBackupPolicyInventory`

NewHyperflexClusterBackupPolicyInventoryWithDefaults instantiates a new HyperflexClusterBackupPolicyInventory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HyperflexClusterBackupPolicyInventory) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HyperflexClusterBackupPolicyInventory) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HyperflexClusterBackupPolicyInventory) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HyperflexClusterBackupPolicyInventory) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HyperflexClusterBackupPolicyInventory) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HyperflexClusterBackupPolicyInventory) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAction

`func (o *HyperflexClusterBackupPolicyInventory) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *HyperflexClusterBackupPolicyInventory) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *HyperflexClusterBackupPolicyInventory) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *HyperflexClusterBackupPolicyInventory) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetCleanup

`func (o *HyperflexClusterBackupPolicyInventory) GetCleanup() bool`

GetCleanup returns the Cleanup field if non-nil, zero value otherwise.

### GetCleanupOk

`func (o *HyperflexClusterBackupPolicyInventory) GetCleanupOk() (*bool, bool)`

GetCleanupOk returns a tuple with the Cleanup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCleanup

`func (o *HyperflexClusterBackupPolicyInventory) SetCleanup(v bool)`

SetCleanup sets Cleanup field to given value.

### HasCleanup

`func (o *HyperflexClusterBackupPolicyInventory) HasCleanup() bool`

HasCleanup returns a boolean if a field has been set.

### GetIsSource

`func (o *HyperflexClusterBackupPolicyInventory) GetIsSource() bool`

GetIsSource returns the IsSource field if non-nil, zero value otherwise.

### GetIsSourceOk

`func (o *HyperflexClusterBackupPolicyInventory) GetIsSourceOk() (*bool, bool)`

GetIsSourceOk returns a tuple with the IsSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSource

`func (o *HyperflexClusterBackupPolicyInventory) SetIsSource(v bool)`

SetIsSource sets IsSource field to given value.

### HasIsSource

`func (o *HyperflexClusterBackupPolicyInventory) HasIsSource() bool`

HasIsSource returns a boolean if a field has been set.

### GetJobDetails

`func (o *HyperflexClusterBackupPolicyInventory) GetJobDetails() string`

GetJobDetails returns the JobDetails field if non-nil, zero value otherwise.

### GetJobDetailsOk

`func (o *HyperflexClusterBackupPolicyInventory) GetJobDetailsOk() (*string, bool)`

GetJobDetailsOk returns a tuple with the JobDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobDetails

`func (o *HyperflexClusterBackupPolicyInventory) SetJobDetails(v string)`

SetJobDetails sets JobDetails field to given value.

### HasJobDetails

`func (o *HyperflexClusterBackupPolicyInventory) HasJobDetails() bool`

HasJobDetails returns a boolean if a field has been set.

### GetJobExceptionMessage

`func (o *HyperflexClusterBackupPolicyInventory) GetJobExceptionMessage() string`

GetJobExceptionMessage returns the JobExceptionMessage field if non-nil, zero value otherwise.

### GetJobExceptionMessageOk

`func (o *HyperflexClusterBackupPolicyInventory) GetJobExceptionMessageOk() (*string, bool)`

GetJobExceptionMessageOk returns a tuple with the JobExceptionMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobExceptionMessage

`func (o *HyperflexClusterBackupPolicyInventory) SetJobExceptionMessage(v string)`

SetJobExceptionMessage sets JobExceptionMessage field to given value.

### HasJobExceptionMessage

`func (o *HyperflexClusterBackupPolicyInventory) HasJobExceptionMessage() bool`

HasJobExceptionMessage returns a boolean if a field has been set.

### GetJobId

`func (o *HyperflexClusterBackupPolicyInventory) GetJobId() string`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *HyperflexClusterBackupPolicyInventory) GetJobIdOk() (*string, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *HyperflexClusterBackupPolicyInventory) SetJobId(v string)`

SetJobId sets JobId field to given value.

### HasJobId

`func (o *HyperflexClusterBackupPolicyInventory) HasJobId() bool`

HasJobId returns a boolean if a field has been set.

### GetJobState

`func (o *HyperflexClusterBackupPolicyInventory) GetJobState() string`

GetJobState returns the JobState field if non-nil, zero value otherwise.

### GetJobStateOk

`func (o *HyperflexClusterBackupPolicyInventory) GetJobStateOk() (*string, bool)`

GetJobStateOk returns a tuple with the JobState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobState

`func (o *HyperflexClusterBackupPolicyInventory) SetJobState(v string)`

SetJobState sets JobState field to given value.

### HasJobState

`func (o *HyperflexClusterBackupPolicyInventory) HasJobState() bool`

HasJobState returns a boolean if a field has been set.

### GetPolicyMoid

`func (o *HyperflexClusterBackupPolicyInventory) GetPolicyMoid() string`

GetPolicyMoid returns the PolicyMoid field if non-nil, zero value otherwise.

### GetPolicyMoidOk

`func (o *HyperflexClusterBackupPolicyInventory) GetPolicyMoidOk() (*string, bool)`

GetPolicyMoidOk returns a tuple with the PolicyMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyMoid

`func (o *HyperflexClusterBackupPolicyInventory) SetPolicyMoid(v string)`

SetPolicyMoid sets PolicyMoid field to given value.

### HasPolicyMoid

`func (o *HyperflexClusterBackupPolicyInventory) HasPolicyMoid() bool`

HasPolicyMoid returns a boolean if a field has been set.

### GetRequestId

`func (o *HyperflexClusterBackupPolicyInventory) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *HyperflexClusterBackupPolicyInventory) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *HyperflexClusterBackupPolicyInventory) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *HyperflexClusterBackupPolicyInventory) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetSettings

`func (o *HyperflexClusterBackupPolicyInventory) GetSettings() HyperflexBackupPolicySettings`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *HyperflexClusterBackupPolicyInventory) GetSettingsOk() (*HyperflexBackupPolicySettings, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *HyperflexClusterBackupPolicyInventory) SetSettings(v HyperflexBackupPolicySettings)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *HyperflexClusterBackupPolicyInventory) HasSettings() bool`

HasSettings returns a boolean if a field has been set.

### SetSettingsNil

`func (o *HyperflexClusterBackupPolicyInventory) SetSettingsNil(b bool)`

 SetSettingsNil sets the value for Settings to be an explicit nil

### UnsetSettings
`func (o *HyperflexClusterBackupPolicyInventory) UnsetSettings()`

UnsetSettings ensures that no value is present for Settings, not even an explicit nil
### GetSourceUuid

`func (o *HyperflexClusterBackupPolicyInventory) GetSourceUuid() string`

GetSourceUuid returns the SourceUuid field if non-nil, zero value otherwise.

### GetSourceUuidOk

`func (o *HyperflexClusterBackupPolicyInventory) GetSourceUuidOk() (*string, bool)`

GetSourceUuidOk returns a tuple with the SourceUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceUuid

`func (o *HyperflexClusterBackupPolicyInventory) SetSourceUuid(v string)`

SetSourceUuid sets SourceUuid field to given value.

### HasSourceUuid

`func (o *HyperflexClusterBackupPolicyInventory) HasSourceUuid() bool`

HasSourceUuid returns a boolean if a field has been set.

### GetTargetUuid

`func (o *HyperflexClusterBackupPolicyInventory) GetTargetUuid() string`

GetTargetUuid returns the TargetUuid field if non-nil, zero value otherwise.

### GetTargetUuidOk

`func (o *HyperflexClusterBackupPolicyInventory) GetTargetUuidOk() (*string, bool)`

GetTargetUuidOk returns a tuple with the TargetUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetUuid

`func (o *HyperflexClusterBackupPolicyInventory) SetTargetUuid(v string)`

SetTargetUuid sets TargetUuid field to given value.

### HasTargetUuid

`func (o *HyperflexClusterBackupPolicyInventory) HasTargetUuid() bool`

HasTargetUuid returns a boolean if a field has been set.

### GetVersion

`func (o *HyperflexClusterBackupPolicyInventory) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *HyperflexClusterBackupPolicyInventory) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *HyperflexClusterBackupPolicyInventory) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *HyperflexClusterBackupPolicyInventory) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetSrcCluster

`func (o *HyperflexClusterBackupPolicyInventory) GetSrcCluster() HyperflexClusterRelationship`

GetSrcCluster returns the SrcCluster field if non-nil, zero value otherwise.

### GetSrcClusterOk

`func (o *HyperflexClusterBackupPolicyInventory) GetSrcClusterOk() (*HyperflexClusterRelationship, bool)`

GetSrcClusterOk returns a tuple with the SrcCluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcCluster

`func (o *HyperflexClusterBackupPolicyInventory) SetSrcCluster(v HyperflexClusterRelationship)`

SetSrcCluster sets SrcCluster field to given value.

### HasSrcCluster

`func (o *HyperflexClusterBackupPolicyInventory) HasSrcCluster() bool`

HasSrcCluster returns a boolean if a field has been set.

### GetTgtCluster

`func (o *HyperflexClusterBackupPolicyInventory) GetTgtCluster() HyperflexClusterRelationship`

GetTgtCluster returns the TgtCluster field if non-nil, zero value otherwise.

### GetTgtClusterOk

`func (o *HyperflexClusterBackupPolicyInventory) GetTgtClusterOk() (*HyperflexClusterRelationship, bool)`

GetTgtClusterOk returns a tuple with the TgtCluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTgtCluster

`func (o *HyperflexClusterBackupPolicyInventory) SetTgtCluster(v HyperflexClusterRelationship)`

SetTgtCluster sets TgtCluster field to given value.

### HasTgtCluster

`func (o *HyperflexClusterBackupPolicyInventory) HasTgtCluster() bool`

HasTgtCluster returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


