# HyperflexReplicationPlatDatastorePair

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hyperflex.ReplicationPlatDatastorePair"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hyperflex.ReplicationPlatDatastorePair"]
**Ads** | Pointer to [**NullableHyperflexReplicationPlatDatastore**](HyperflexReplicationPlatDatastore.md) |  | [optional] 
**BackupOnly** | Pointer to **bool** | Boolean representing if this is a backup only pair. | [optional] [readonly] 
**Bds** | Pointer to [**NullableHyperflexReplicationPlatDatastore**](HyperflexReplicationPlatDatastore.md) |  | [optional] 
**Quiesce** | Pointer to **bool** | Boolean representing if this datastore pairing has quiesce snapshots enabled. | [optional] [readonly] 
**ReplicationIntervalInMinutes** | Pointer to **int64** | The replication interval for this pair in minutes. | [optional] [readonly] 
**Sourceds** | Pointer to [**NullableHyperflexReplicationPlatDatastore**](HyperflexReplicationPlatDatastore.md) |  | [optional] 
**StorageOnly** | Pointer to **bool** | HyperFlex datastore pair is used for storage only. | [optional] [readonly] 

## Methods

### NewHyperflexReplicationPlatDatastorePair

`func NewHyperflexReplicationPlatDatastorePair(classId string, objectType string, ) *HyperflexReplicationPlatDatastorePair`

NewHyperflexReplicationPlatDatastorePair instantiates a new HyperflexReplicationPlatDatastorePair object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexReplicationPlatDatastorePairWithDefaults

`func NewHyperflexReplicationPlatDatastorePairWithDefaults() *HyperflexReplicationPlatDatastorePair`

NewHyperflexReplicationPlatDatastorePairWithDefaults instantiates a new HyperflexReplicationPlatDatastorePair object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HyperflexReplicationPlatDatastorePair) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HyperflexReplicationPlatDatastorePair) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HyperflexReplicationPlatDatastorePair) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HyperflexReplicationPlatDatastorePair) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HyperflexReplicationPlatDatastorePair) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HyperflexReplicationPlatDatastorePair) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAds

`func (o *HyperflexReplicationPlatDatastorePair) GetAds() HyperflexReplicationPlatDatastore`

GetAds returns the Ads field if non-nil, zero value otherwise.

### GetAdsOk

`func (o *HyperflexReplicationPlatDatastorePair) GetAdsOk() (*HyperflexReplicationPlatDatastore, bool)`

GetAdsOk returns a tuple with the Ads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAds

`func (o *HyperflexReplicationPlatDatastorePair) SetAds(v HyperflexReplicationPlatDatastore)`

SetAds sets Ads field to given value.

### HasAds

`func (o *HyperflexReplicationPlatDatastorePair) HasAds() bool`

HasAds returns a boolean if a field has been set.

### SetAdsNil

`func (o *HyperflexReplicationPlatDatastorePair) SetAdsNil(b bool)`

 SetAdsNil sets the value for Ads to be an explicit nil

### UnsetAds
`func (o *HyperflexReplicationPlatDatastorePair) UnsetAds()`

UnsetAds ensures that no value is present for Ads, not even an explicit nil
### GetBackupOnly

`func (o *HyperflexReplicationPlatDatastorePair) GetBackupOnly() bool`

GetBackupOnly returns the BackupOnly field if non-nil, zero value otherwise.

### GetBackupOnlyOk

`func (o *HyperflexReplicationPlatDatastorePair) GetBackupOnlyOk() (*bool, bool)`

GetBackupOnlyOk returns a tuple with the BackupOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupOnly

`func (o *HyperflexReplicationPlatDatastorePair) SetBackupOnly(v bool)`

SetBackupOnly sets BackupOnly field to given value.

### HasBackupOnly

`func (o *HyperflexReplicationPlatDatastorePair) HasBackupOnly() bool`

HasBackupOnly returns a boolean if a field has been set.

### GetBds

`func (o *HyperflexReplicationPlatDatastorePair) GetBds() HyperflexReplicationPlatDatastore`

GetBds returns the Bds field if non-nil, zero value otherwise.

### GetBdsOk

`func (o *HyperflexReplicationPlatDatastorePair) GetBdsOk() (*HyperflexReplicationPlatDatastore, bool)`

GetBdsOk returns a tuple with the Bds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBds

`func (o *HyperflexReplicationPlatDatastorePair) SetBds(v HyperflexReplicationPlatDatastore)`

SetBds sets Bds field to given value.

### HasBds

`func (o *HyperflexReplicationPlatDatastorePair) HasBds() bool`

HasBds returns a boolean if a field has been set.

### SetBdsNil

`func (o *HyperflexReplicationPlatDatastorePair) SetBdsNil(b bool)`

 SetBdsNil sets the value for Bds to be an explicit nil

### UnsetBds
`func (o *HyperflexReplicationPlatDatastorePair) UnsetBds()`

UnsetBds ensures that no value is present for Bds, not even an explicit nil
### GetQuiesce

`func (o *HyperflexReplicationPlatDatastorePair) GetQuiesce() bool`

GetQuiesce returns the Quiesce field if non-nil, zero value otherwise.

### GetQuiesceOk

`func (o *HyperflexReplicationPlatDatastorePair) GetQuiesceOk() (*bool, bool)`

GetQuiesceOk returns a tuple with the Quiesce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuiesce

`func (o *HyperflexReplicationPlatDatastorePair) SetQuiesce(v bool)`

SetQuiesce sets Quiesce field to given value.

### HasQuiesce

`func (o *HyperflexReplicationPlatDatastorePair) HasQuiesce() bool`

HasQuiesce returns a boolean if a field has been set.

### GetReplicationIntervalInMinutes

`func (o *HyperflexReplicationPlatDatastorePair) GetReplicationIntervalInMinutes() int64`

GetReplicationIntervalInMinutes returns the ReplicationIntervalInMinutes field if non-nil, zero value otherwise.

### GetReplicationIntervalInMinutesOk

`func (o *HyperflexReplicationPlatDatastorePair) GetReplicationIntervalInMinutesOk() (*int64, bool)`

GetReplicationIntervalInMinutesOk returns a tuple with the ReplicationIntervalInMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationIntervalInMinutes

`func (o *HyperflexReplicationPlatDatastorePair) SetReplicationIntervalInMinutes(v int64)`

SetReplicationIntervalInMinutes sets ReplicationIntervalInMinutes field to given value.

### HasReplicationIntervalInMinutes

`func (o *HyperflexReplicationPlatDatastorePair) HasReplicationIntervalInMinutes() bool`

HasReplicationIntervalInMinutes returns a boolean if a field has been set.

### GetSourceds

`func (o *HyperflexReplicationPlatDatastorePair) GetSourceds() HyperflexReplicationPlatDatastore`

GetSourceds returns the Sourceds field if non-nil, zero value otherwise.

### GetSourcedsOk

`func (o *HyperflexReplicationPlatDatastorePair) GetSourcedsOk() (*HyperflexReplicationPlatDatastore, bool)`

GetSourcedsOk returns a tuple with the Sourceds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceds

`func (o *HyperflexReplicationPlatDatastorePair) SetSourceds(v HyperflexReplicationPlatDatastore)`

SetSourceds sets Sourceds field to given value.

### HasSourceds

`func (o *HyperflexReplicationPlatDatastorePair) HasSourceds() bool`

HasSourceds returns a boolean if a field has been set.

### SetSourcedsNil

`func (o *HyperflexReplicationPlatDatastorePair) SetSourcedsNil(b bool)`

 SetSourcedsNil sets the value for Sourceds to be an explicit nil

### UnsetSourceds
`func (o *HyperflexReplicationPlatDatastorePair) UnsetSourceds()`

UnsetSourceds ensures that no value is present for Sourceds, not even an explicit nil
### GetStorageOnly

`func (o *HyperflexReplicationPlatDatastorePair) GetStorageOnly() bool`

GetStorageOnly returns the StorageOnly field if non-nil, zero value otherwise.

### GetStorageOnlyOk

`func (o *HyperflexReplicationPlatDatastorePair) GetStorageOnlyOk() (*bool, bool)`

GetStorageOnlyOk returns a tuple with the StorageOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageOnly

`func (o *HyperflexReplicationPlatDatastorePair) SetStorageOnly(v bool)`

SetStorageOnly sets StorageOnly field to given value.

### HasStorageOnly

`func (o *HyperflexReplicationPlatDatastorePair) HasStorageOnly() bool`

HasStorageOnly returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


