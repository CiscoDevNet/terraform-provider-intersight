# HyperflexProtectedCluster

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hyperflex.ProtectedCluster"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hyperflex.ProtectedCluster"]
**HxVersion** | Pointer to **string** | Version of the Hyperflex cluster. | [optional] [readonly] 
**HypervisorVersion** | Pointer to **string** | The version of hypervisor running on this cluster. | [optional] [readonly] 
**ProtectedDatastoreName** | Pointer to **string** | Name of the protected datastore. | [optional] [readonly] 
**ProtectedVmsCount** | Pointer to **int64** | Number of VMs protected on this cluster. | [optional] [readonly] 
**SourceClusterName** | Pointer to **string** | Name of the source cluster. | [optional] [readonly] 
**TargetClusterName** | Pointer to **string** | Name of the target cluster. | [optional] [readonly] 
**TargetDatastoreName** | Pointer to **string** | Name of the target datastore. | [optional] [readonly] 
**TargetDatastoreUtilization** | Pointer to **float32** | Percent usage of the datastore. | [optional] [readonly] 
**BackupPolicy** | Pointer to [**HyperflexClusterBackupPolicyInventoryRelationship**](HyperflexClusterBackupPolicyInventoryRelationship.md) |  | [optional] 
**DatastoreStatistic** | Pointer to [**HyperflexDatastoreStatisticRelationship**](HyperflexDatastoreStatisticRelationship.md) |  | [optional] 
**SrcCluster** | Pointer to [**HyperflexClusterRelationship**](HyperflexClusterRelationship.md) |  | [optional] 
**TgtCluster** | Pointer to [**HyperflexClusterRelationship**](HyperflexClusterRelationship.md) |  | [optional] 

## Methods

### NewHyperflexProtectedCluster

`func NewHyperflexProtectedCluster(classId string, objectType string, ) *HyperflexProtectedCluster`

NewHyperflexProtectedCluster instantiates a new HyperflexProtectedCluster object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexProtectedClusterWithDefaults

`func NewHyperflexProtectedClusterWithDefaults() *HyperflexProtectedCluster`

NewHyperflexProtectedClusterWithDefaults instantiates a new HyperflexProtectedCluster object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HyperflexProtectedCluster) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HyperflexProtectedCluster) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HyperflexProtectedCluster) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HyperflexProtectedCluster) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HyperflexProtectedCluster) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HyperflexProtectedCluster) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetHxVersion

`func (o *HyperflexProtectedCluster) GetHxVersion() string`

GetHxVersion returns the HxVersion field if non-nil, zero value otherwise.

### GetHxVersionOk

`func (o *HyperflexProtectedCluster) GetHxVersionOk() (*string, bool)`

GetHxVersionOk returns a tuple with the HxVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHxVersion

`func (o *HyperflexProtectedCluster) SetHxVersion(v string)`

SetHxVersion sets HxVersion field to given value.

### HasHxVersion

`func (o *HyperflexProtectedCluster) HasHxVersion() bool`

HasHxVersion returns a boolean if a field has been set.

### GetHypervisorVersion

`func (o *HyperflexProtectedCluster) GetHypervisorVersion() string`

GetHypervisorVersion returns the HypervisorVersion field if non-nil, zero value otherwise.

### GetHypervisorVersionOk

`func (o *HyperflexProtectedCluster) GetHypervisorVersionOk() (*string, bool)`

GetHypervisorVersionOk returns a tuple with the HypervisorVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHypervisorVersion

`func (o *HyperflexProtectedCluster) SetHypervisorVersion(v string)`

SetHypervisorVersion sets HypervisorVersion field to given value.

### HasHypervisorVersion

`func (o *HyperflexProtectedCluster) HasHypervisorVersion() bool`

HasHypervisorVersion returns a boolean if a field has been set.

### GetProtectedDatastoreName

`func (o *HyperflexProtectedCluster) GetProtectedDatastoreName() string`

GetProtectedDatastoreName returns the ProtectedDatastoreName field if non-nil, zero value otherwise.

### GetProtectedDatastoreNameOk

`func (o *HyperflexProtectedCluster) GetProtectedDatastoreNameOk() (*string, bool)`

GetProtectedDatastoreNameOk returns a tuple with the ProtectedDatastoreName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectedDatastoreName

`func (o *HyperflexProtectedCluster) SetProtectedDatastoreName(v string)`

SetProtectedDatastoreName sets ProtectedDatastoreName field to given value.

### HasProtectedDatastoreName

`func (o *HyperflexProtectedCluster) HasProtectedDatastoreName() bool`

HasProtectedDatastoreName returns a boolean if a field has been set.

### GetProtectedVmsCount

`func (o *HyperflexProtectedCluster) GetProtectedVmsCount() int64`

GetProtectedVmsCount returns the ProtectedVmsCount field if non-nil, zero value otherwise.

### GetProtectedVmsCountOk

`func (o *HyperflexProtectedCluster) GetProtectedVmsCountOk() (*int64, bool)`

GetProtectedVmsCountOk returns a tuple with the ProtectedVmsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectedVmsCount

`func (o *HyperflexProtectedCluster) SetProtectedVmsCount(v int64)`

SetProtectedVmsCount sets ProtectedVmsCount field to given value.

### HasProtectedVmsCount

`func (o *HyperflexProtectedCluster) HasProtectedVmsCount() bool`

HasProtectedVmsCount returns a boolean if a field has been set.

### GetSourceClusterName

`func (o *HyperflexProtectedCluster) GetSourceClusterName() string`

GetSourceClusterName returns the SourceClusterName field if non-nil, zero value otherwise.

### GetSourceClusterNameOk

`func (o *HyperflexProtectedCluster) GetSourceClusterNameOk() (*string, bool)`

GetSourceClusterNameOk returns a tuple with the SourceClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceClusterName

`func (o *HyperflexProtectedCluster) SetSourceClusterName(v string)`

SetSourceClusterName sets SourceClusterName field to given value.

### HasSourceClusterName

`func (o *HyperflexProtectedCluster) HasSourceClusterName() bool`

HasSourceClusterName returns a boolean if a field has been set.

### GetTargetClusterName

`func (o *HyperflexProtectedCluster) GetTargetClusterName() string`

GetTargetClusterName returns the TargetClusterName field if non-nil, zero value otherwise.

### GetTargetClusterNameOk

`func (o *HyperflexProtectedCluster) GetTargetClusterNameOk() (*string, bool)`

GetTargetClusterNameOk returns a tuple with the TargetClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetClusterName

`func (o *HyperflexProtectedCluster) SetTargetClusterName(v string)`

SetTargetClusterName sets TargetClusterName field to given value.

### HasTargetClusterName

`func (o *HyperflexProtectedCluster) HasTargetClusterName() bool`

HasTargetClusterName returns a boolean if a field has been set.

### GetTargetDatastoreName

`func (o *HyperflexProtectedCluster) GetTargetDatastoreName() string`

GetTargetDatastoreName returns the TargetDatastoreName field if non-nil, zero value otherwise.

### GetTargetDatastoreNameOk

`func (o *HyperflexProtectedCluster) GetTargetDatastoreNameOk() (*string, bool)`

GetTargetDatastoreNameOk returns a tuple with the TargetDatastoreName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetDatastoreName

`func (o *HyperflexProtectedCluster) SetTargetDatastoreName(v string)`

SetTargetDatastoreName sets TargetDatastoreName field to given value.

### HasTargetDatastoreName

`func (o *HyperflexProtectedCluster) HasTargetDatastoreName() bool`

HasTargetDatastoreName returns a boolean if a field has been set.

### GetTargetDatastoreUtilization

`func (o *HyperflexProtectedCluster) GetTargetDatastoreUtilization() float32`

GetTargetDatastoreUtilization returns the TargetDatastoreUtilization field if non-nil, zero value otherwise.

### GetTargetDatastoreUtilizationOk

`func (o *HyperflexProtectedCluster) GetTargetDatastoreUtilizationOk() (*float32, bool)`

GetTargetDatastoreUtilizationOk returns a tuple with the TargetDatastoreUtilization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetDatastoreUtilization

`func (o *HyperflexProtectedCluster) SetTargetDatastoreUtilization(v float32)`

SetTargetDatastoreUtilization sets TargetDatastoreUtilization field to given value.

### HasTargetDatastoreUtilization

`func (o *HyperflexProtectedCluster) HasTargetDatastoreUtilization() bool`

HasTargetDatastoreUtilization returns a boolean if a field has been set.

### GetBackupPolicy

`func (o *HyperflexProtectedCluster) GetBackupPolicy() HyperflexClusterBackupPolicyInventoryRelationship`

GetBackupPolicy returns the BackupPolicy field if non-nil, zero value otherwise.

### GetBackupPolicyOk

`func (o *HyperflexProtectedCluster) GetBackupPolicyOk() (*HyperflexClusterBackupPolicyInventoryRelationship, bool)`

GetBackupPolicyOk returns a tuple with the BackupPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupPolicy

`func (o *HyperflexProtectedCluster) SetBackupPolicy(v HyperflexClusterBackupPolicyInventoryRelationship)`

SetBackupPolicy sets BackupPolicy field to given value.

### HasBackupPolicy

`func (o *HyperflexProtectedCluster) HasBackupPolicy() bool`

HasBackupPolicy returns a boolean if a field has been set.

### GetDatastoreStatistic

`func (o *HyperflexProtectedCluster) GetDatastoreStatistic() HyperflexDatastoreStatisticRelationship`

GetDatastoreStatistic returns the DatastoreStatistic field if non-nil, zero value otherwise.

### GetDatastoreStatisticOk

`func (o *HyperflexProtectedCluster) GetDatastoreStatisticOk() (*HyperflexDatastoreStatisticRelationship, bool)`

GetDatastoreStatisticOk returns a tuple with the DatastoreStatistic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastoreStatistic

`func (o *HyperflexProtectedCluster) SetDatastoreStatistic(v HyperflexDatastoreStatisticRelationship)`

SetDatastoreStatistic sets DatastoreStatistic field to given value.

### HasDatastoreStatistic

`func (o *HyperflexProtectedCluster) HasDatastoreStatistic() bool`

HasDatastoreStatistic returns a boolean if a field has been set.

### GetSrcCluster

`func (o *HyperflexProtectedCluster) GetSrcCluster() HyperflexClusterRelationship`

GetSrcCluster returns the SrcCluster field if non-nil, zero value otherwise.

### GetSrcClusterOk

`func (o *HyperflexProtectedCluster) GetSrcClusterOk() (*HyperflexClusterRelationship, bool)`

GetSrcClusterOk returns a tuple with the SrcCluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcCluster

`func (o *HyperflexProtectedCluster) SetSrcCluster(v HyperflexClusterRelationship)`

SetSrcCluster sets SrcCluster field to given value.

### HasSrcCluster

`func (o *HyperflexProtectedCluster) HasSrcCluster() bool`

HasSrcCluster returns a boolean if a field has been set.

### GetTgtCluster

`func (o *HyperflexProtectedCluster) GetTgtCluster() HyperflexClusterRelationship`

GetTgtCluster returns the TgtCluster field if non-nil, zero value otherwise.

### GetTgtClusterOk

`func (o *HyperflexProtectedCluster) GetTgtClusterOk() (*HyperflexClusterRelationship, bool)`

GetTgtClusterOk returns a tuple with the TgtCluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTgtCluster

`func (o *HyperflexProtectedCluster) SetTgtCluster(v HyperflexClusterRelationship)`

SetTgtCluster sets TgtCluster field to given value.

### HasTgtCluster

`func (o *HyperflexProtectedCluster) HasTgtCluster() bool`

HasTgtCluster returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


