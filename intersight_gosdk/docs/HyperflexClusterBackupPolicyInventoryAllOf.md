# HyperflexClusterBackupPolicyInventoryAllOf

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

### NewHyperflexClusterBackupPolicyInventoryAllOf

`func NewHyperflexClusterBackupPolicyInventoryAllOf(classId string, objectType string, ) *HyperflexClusterBackupPolicyInventoryAllOf`

NewHyperflexClusterBackupPolicyInventoryAllOf instantiates a new HyperflexClusterBackupPolicyInventoryAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexClusterBackupPolicyInventoryAllOfWithDefaults

`func NewHyperflexClusterBackupPolicyInventoryAllOfWithDefaults() *HyperflexClusterBackupPolicyInventoryAllOf`

NewHyperflexClusterBackupPolicyInventoryAllOfWithDefaults instantiates a new HyperflexClusterBackupPolicyInventoryAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HyperflexClusterBackupPolicyInventoryAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HyperflexClusterBackupPolicyInventoryAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HyperflexClusterBackupPolicyInventoryAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HyperflexClusterBackupPolicyInventoryAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HyperflexClusterBackupPolicyInventoryAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HyperflexClusterBackupPolicyInventoryAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAction

`func (o *HyperflexClusterBackupPolicyInventoryAllOf) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *HyperflexClusterBackupPolicyInventoryAllOf) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *HyperflexClusterBackupPolicyInventoryAllOf) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *HyperflexClusterBackupPolicyInventoryAllOf) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetCleanup

`func (o *HyperflexClusterBackupPolicyInventoryAllOf) GetCleanup() bool`

GetCleanup returns the Cleanup field if non-nil, zero value otherwise.

### GetCleanupOk

`func (o *HyperflexClusterBackupPolicyInventoryAllOf) GetCleanupOk() (*bool, bool)`

GetCleanupOk returns a tuple with the Cleanup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCleanup

`func (o *HyperflexClusterBackupPolicyInventoryAllOf) SetCleanup(v bool)`

SetCleanup sets Cleanup field to given value.

### HasCleanup

`func (o *HyperflexClusterBackupPolicyInventoryAllOf) HasCleanup() bool`

HasCleanup returns a boolean if a field has been set.

### GetIsSource

`func (o *HyperflexClusterBackupPolicyInventoryAllOf) GetIsSource() bool`

GetIsSource returns the IsSource field if non-nil, zero value otherwise.

### GetIsSourceOk

`func (o *HyperflexClusterBackupPolicyInventoryAllOf) GetIsSourceOk() (*bool, bool)`

GetIsSourceOk returns a tuple with the IsSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSource

`func (o *HyperflexClusterBackupPolicyInventoryAllOf) SetIsSource(v bool)`

SetIsSource sets IsSource field to given value.

### HasIsSource

`func (o *HyperflexClusterBackupPolicyInventoryAllOf) HasIsSource() bool`

HasIsSource returns a boolean if a field has been set.

### GetJobDetails

`func (o *HyperflexClusterBackupPolicyInventoryAllOf) GetJobDetails() string`

GetJobDetails returns the JobDetails field if non-nil, zero value otherwise.

### GetJobDetailsOk

`func (o *HyperflexClusterBackupPolicyInventoryAllOf) GetJobDetailsOk() (*string, bool)`

GetJobDetailsOk returns a tuple with the JobDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobDetails

`func (o *HyperflexClusterBackupPolicyInventoryAllOf) SetJobDetails(v string)`

SetJobDetails sets JobDetails field to given value.

### HasJobDetails

`func (o *HyperflexClusterBackupPolicyInventoryAllOf) HasJobDetails() bool`

HasJobDetails returns a boolean if a field has been set.

### GetJobExceptionMessage

`func (o *HyperflexClusterBackupPolicyInventoryAllOf) GetJobExceptionMessage() string`

GetJobExceptionMessage returns the JobExceptionMessage field if non-nil, zero value otherwise.

### GetJobExceptionMessageOk

`func (o *HyperflexClusterBackupPolicyInventoryAllOf) GetJobExceptionMessageOk() (*string, bool)`

GetJobExceptionMessageOk returns a tuple with the JobExceptionMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobExceptionMessage

`func (o *HyperflexClusterBackupPolicyInventoryAllOf) SetJobExceptionMessage(v string)`

SetJobExceptionMessage sets JobExceptionMessage field to given value.

### HasJobExceptionMessage

`func (o *HyperflexClusterBackupPolicyInventoryAllOf) HasJobExceptionMessage() bool`

HasJobExceptionMessage returns a boolean if a field has been set.

### GetJobId

`func (o *HyperflexClusterBackupPolicyInventoryAllOf) GetJobId() string`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *HyperflexClusterBackupPolicyInventoryAllOf) GetJobIdOk() (*string, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *HyperflexClusterBackupPolicyInventoryAllOf) SetJobId(v string)`

SetJobId sets JobId field to given value.

### HasJobId

`func (o *HyperflexClusterBackupPolicyInventoryAllOf) HasJobId() bool`

HasJobId returns a boolean if a field has been set.

### GetJobState

`func (o *HyperflexClusterBackupPolicyInventoryAllOf) GetJobState() string`

GetJobState returns the JobState field if non-nil, zero value otherwise.

### GetJobStateOk

`func (o *HyperflexClusterBackupPolicyInventoryAllOf) GetJobStateOk() (*string, bool)`

GetJobStateOk returns a tuple with the JobState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobState

`func (o *HyperflexClusterBackupPolicyInventoryAllOf) SetJobState(v string)`

SetJobState sets JobState field to given value.

### HasJobState

`func (o *HyperflexClusterBackupPolicyInventoryAllOf) HasJobState() bool`

HasJobState returns a boolean if a field has been set.

### GetPolicyMoid

`func (o *HyperflexClusterBackupPolicyInventoryAllOf) GetPolicyMoid() string`

GetPolicyMoid returns the PolicyMoid field if non-nil, zero value otherwise.

### GetPolicyMoidOk

`func (o *HyperflexClusterBackupPolicyInventoryAllOf) GetPolicyMoidOk() (*string, bool)`

GetPolicyMoidOk returns a tuple with the PolicyMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyMoid

`func (o *HyperflexClusterBackupPolicyInventoryAllOf) SetPolicyMoid(v string)`

SetPolicyMoid sets PolicyMoid field to given value.

### HasPolicyMoid

`func (o *HyperflexClusterBackupPolicyInventoryAllOf) HasPolicyMoid() bool`

HasPolicyMoid returns a boolean if a field has been set.

### GetRequestId

`func (o *HyperflexClusterBackupPolicyInventoryAllOf) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *HyperflexClusterBackupPolicyInventoryAllOf) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *HyperflexClusterBackupPolicyInventoryAllOf) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *HyperflexClusterBackupPolicyInventoryAllOf) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetSettings

`func (o *HyperflexClusterBackupPolicyInventoryAllOf) GetSettings() HyperflexBackupPolicySettings`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *HyperflexClusterBackupPolicyInventoryAllOf) GetSettingsOk() (*HyperflexBackupPolicySettings, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *HyperflexClusterBackupPolicyInventoryAllOf) SetSettings(v HyperflexBackupPolicySettings)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *HyperflexClusterBackupPolicyInventoryAllOf) HasSettings() bool`

HasSettings returns a boolean if a field has been set.

### SetSettingsNil

`func (o *HyperflexClusterBackupPolicyInventoryAllOf) SetSettingsNil(b bool)`

 SetSettingsNil sets the value for Settings to be an explicit nil

### UnsetSettings
`func (o *HyperflexClusterBackupPolicyInventoryAllOf) UnsetSettings()`

UnsetSettings ensures that no value is present for Settings, not even an explicit nil
### GetSourceUuid

`func (o *HyperflexClusterBackupPolicyInventoryAllOf) GetSourceUuid() string`

GetSourceUuid returns the SourceUuid field if non-nil, zero value otherwise.

### GetSourceUuidOk

`func (o *HyperflexClusterBackupPolicyInventoryAllOf) GetSourceUuidOk() (*string, bool)`

GetSourceUuidOk returns a tuple with the SourceUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceUuid

`func (o *HyperflexClusterBackupPolicyInventoryAllOf) SetSourceUuid(v string)`

SetSourceUuid sets SourceUuid field to given value.

### HasSourceUuid

`func (o *HyperflexClusterBackupPolicyInventoryAllOf) HasSourceUuid() bool`

HasSourceUuid returns a boolean if a field has been set.

### GetTargetUuid

`func (o *HyperflexClusterBackupPolicyInventoryAllOf) GetTargetUuid() string`

GetTargetUuid returns the TargetUuid field if non-nil, zero value otherwise.

### GetTargetUuidOk

`func (o *HyperflexClusterBackupPolicyInventoryAllOf) GetTargetUuidOk() (*string, bool)`

GetTargetUuidOk returns a tuple with the TargetUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetUuid

`func (o *HyperflexClusterBackupPolicyInventoryAllOf) SetTargetUuid(v string)`

SetTargetUuid sets TargetUuid field to given value.

### HasTargetUuid

`func (o *HyperflexClusterBackupPolicyInventoryAllOf) HasTargetUuid() bool`

HasTargetUuid returns a boolean if a field has been set.

### GetVersion

`func (o *HyperflexClusterBackupPolicyInventoryAllOf) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *HyperflexClusterBackupPolicyInventoryAllOf) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *HyperflexClusterBackupPolicyInventoryAllOf) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *HyperflexClusterBackupPolicyInventoryAllOf) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetSrcCluster

`func (o *HyperflexClusterBackupPolicyInventoryAllOf) GetSrcCluster() HyperflexClusterRelationship`

GetSrcCluster returns the SrcCluster field if non-nil, zero value otherwise.

### GetSrcClusterOk

`func (o *HyperflexClusterBackupPolicyInventoryAllOf) GetSrcClusterOk() (*HyperflexClusterRelationship, bool)`

GetSrcClusterOk returns a tuple with the SrcCluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcCluster

`func (o *HyperflexClusterBackupPolicyInventoryAllOf) SetSrcCluster(v HyperflexClusterRelationship)`

SetSrcCluster sets SrcCluster field to given value.

### HasSrcCluster

`func (o *HyperflexClusterBackupPolicyInventoryAllOf) HasSrcCluster() bool`

HasSrcCluster returns a boolean if a field has been set.

### GetTgtCluster

`func (o *HyperflexClusterBackupPolicyInventoryAllOf) GetTgtCluster() HyperflexClusterRelationship`

GetTgtCluster returns the TgtCluster field if non-nil, zero value otherwise.

### GetTgtClusterOk

`func (o *HyperflexClusterBackupPolicyInventoryAllOf) GetTgtClusterOk() (*HyperflexClusterRelationship, bool)`

GetTgtClusterOk returns a tuple with the TgtCluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTgtCluster

`func (o *HyperflexClusterBackupPolicyInventoryAllOf) SetTgtCluster(v HyperflexClusterRelationship)`

SetTgtCluster sets TgtCluster field to given value.

### HasTgtCluster

`func (o *HyperflexClusterBackupPolicyInventoryAllOf) HasTgtCluster() bool`

HasTgtCluster returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


