# HyperflexBackupCluster

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hyperflex.BackupCluster"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hyperflex.BackupCluster"]
**BackupDataStore** | Pointer to **string** | Defines the backup source cluster and its references. | [optional] [readonly] 
**SrcCluster** | Pointer to [**HyperflexClusterRelationship**](HyperflexClusterRelationship.md) |  | [optional] 
**TgtCluster** | Pointer to [**HyperflexClusterRelationship**](HyperflexClusterRelationship.md) |  | [optional] 

## Methods

### NewHyperflexBackupCluster

`func NewHyperflexBackupCluster(classId string, objectType string, ) *HyperflexBackupCluster`

NewHyperflexBackupCluster instantiates a new HyperflexBackupCluster object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexBackupClusterWithDefaults

`func NewHyperflexBackupClusterWithDefaults() *HyperflexBackupCluster`

NewHyperflexBackupClusterWithDefaults instantiates a new HyperflexBackupCluster object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HyperflexBackupCluster) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HyperflexBackupCluster) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HyperflexBackupCluster) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HyperflexBackupCluster) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HyperflexBackupCluster) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HyperflexBackupCluster) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetBackupDataStore

`func (o *HyperflexBackupCluster) GetBackupDataStore() string`

GetBackupDataStore returns the BackupDataStore field if non-nil, zero value otherwise.

### GetBackupDataStoreOk

`func (o *HyperflexBackupCluster) GetBackupDataStoreOk() (*string, bool)`

GetBackupDataStoreOk returns a tuple with the BackupDataStore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupDataStore

`func (o *HyperflexBackupCluster) SetBackupDataStore(v string)`

SetBackupDataStore sets BackupDataStore field to given value.

### HasBackupDataStore

`func (o *HyperflexBackupCluster) HasBackupDataStore() bool`

HasBackupDataStore returns a boolean if a field has been set.

### GetSrcCluster

`func (o *HyperflexBackupCluster) GetSrcCluster() HyperflexClusterRelationship`

GetSrcCluster returns the SrcCluster field if non-nil, zero value otherwise.

### GetSrcClusterOk

`func (o *HyperflexBackupCluster) GetSrcClusterOk() (*HyperflexClusterRelationship, bool)`

GetSrcClusterOk returns a tuple with the SrcCluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcCluster

`func (o *HyperflexBackupCluster) SetSrcCluster(v HyperflexClusterRelationship)`

SetSrcCluster sets SrcCluster field to given value.

### HasSrcCluster

`func (o *HyperflexBackupCluster) HasSrcCluster() bool`

HasSrcCluster returns a boolean if a field has been set.

### GetTgtCluster

`func (o *HyperflexBackupCluster) GetTgtCluster() HyperflexClusterRelationship`

GetTgtCluster returns the TgtCluster field if non-nil, zero value otherwise.

### GetTgtClusterOk

`func (o *HyperflexBackupCluster) GetTgtClusterOk() (*HyperflexClusterRelationship, bool)`

GetTgtClusterOk returns a tuple with the TgtCluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTgtCluster

`func (o *HyperflexBackupCluster) SetTgtCluster(v HyperflexClusterRelationship)`

SetTgtCluster sets TgtCluster field to given value.

### HasTgtCluster

`func (o *HyperflexBackupCluster) HasTgtCluster() bool`

HasTgtCluster returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


