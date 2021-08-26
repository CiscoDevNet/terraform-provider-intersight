# HyperflexReplicationPlatDatastorePairAllOf

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

### NewHyperflexReplicationPlatDatastorePairAllOf

`func NewHyperflexReplicationPlatDatastorePairAllOf(classId string, objectType string, ) *HyperflexReplicationPlatDatastorePairAllOf`

NewHyperflexReplicationPlatDatastorePairAllOf instantiates a new HyperflexReplicationPlatDatastorePairAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexReplicationPlatDatastorePairAllOfWithDefaults

`func NewHyperflexReplicationPlatDatastorePairAllOfWithDefaults() *HyperflexReplicationPlatDatastorePairAllOf`

NewHyperflexReplicationPlatDatastorePairAllOfWithDefaults instantiates a new HyperflexReplicationPlatDatastorePairAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HyperflexReplicationPlatDatastorePairAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HyperflexReplicationPlatDatastorePairAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HyperflexReplicationPlatDatastorePairAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HyperflexReplicationPlatDatastorePairAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HyperflexReplicationPlatDatastorePairAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HyperflexReplicationPlatDatastorePairAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAds

`func (o *HyperflexReplicationPlatDatastorePairAllOf) GetAds() HyperflexReplicationPlatDatastore`

GetAds returns the Ads field if non-nil, zero value otherwise.

### GetAdsOk

`func (o *HyperflexReplicationPlatDatastorePairAllOf) GetAdsOk() (*HyperflexReplicationPlatDatastore, bool)`

GetAdsOk returns a tuple with the Ads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAds

`func (o *HyperflexReplicationPlatDatastorePairAllOf) SetAds(v HyperflexReplicationPlatDatastore)`

SetAds sets Ads field to given value.

### HasAds

`func (o *HyperflexReplicationPlatDatastorePairAllOf) HasAds() bool`

HasAds returns a boolean if a field has been set.

### SetAdsNil

`func (o *HyperflexReplicationPlatDatastorePairAllOf) SetAdsNil(b bool)`

 SetAdsNil sets the value for Ads to be an explicit nil

### UnsetAds
`func (o *HyperflexReplicationPlatDatastorePairAllOf) UnsetAds()`

UnsetAds ensures that no value is present for Ads, not even an explicit nil
### GetBackupOnly

`func (o *HyperflexReplicationPlatDatastorePairAllOf) GetBackupOnly() bool`

GetBackupOnly returns the BackupOnly field if non-nil, zero value otherwise.

### GetBackupOnlyOk

`func (o *HyperflexReplicationPlatDatastorePairAllOf) GetBackupOnlyOk() (*bool, bool)`

GetBackupOnlyOk returns a tuple with the BackupOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupOnly

`func (o *HyperflexReplicationPlatDatastorePairAllOf) SetBackupOnly(v bool)`

SetBackupOnly sets BackupOnly field to given value.

### HasBackupOnly

`func (o *HyperflexReplicationPlatDatastorePairAllOf) HasBackupOnly() bool`

HasBackupOnly returns a boolean if a field has been set.

### GetBds

`func (o *HyperflexReplicationPlatDatastorePairAllOf) GetBds() HyperflexReplicationPlatDatastore`

GetBds returns the Bds field if non-nil, zero value otherwise.

### GetBdsOk

`func (o *HyperflexReplicationPlatDatastorePairAllOf) GetBdsOk() (*HyperflexReplicationPlatDatastore, bool)`

GetBdsOk returns a tuple with the Bds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBds

`func (o *HyperflexReplicationPlatDatastorePairAllOf) SetBds(v HyperflexReplicationPlatDatastore)`

SetBds sets Bds field to given value.

### HasBds

`func (o *HyperflexReplicationPlatDatastorePairAllOf) HasBds() bool`

HasBds returns a boolean if a field has been set.

### SetBdsNil

`func (o *HyperflexReplicationPlatDatastorePairAllOf) SetBdsNil(b bool)`

 SetBdsNil sets the value for Bds to be an explicit nil

### UnsetBds
`func (o *HyperflexReplicationPlatDatastorePairAllOf) UnsetBds()`

UnsetBds ensures that no value is present for Bds, not even an explicit nil
### GetQuiesce

`func (o *HyperflexReplicationPlatDatastorePairAllOf) GetQuiesce() bool`

GetQuiesce returns the Quiesce field if non-nil, zero value otherwise.

### GetQuiesceOk

`func (o *HyperflexReplicationPlatDatastorePairAllOf) GetQuiesceOk() (*bool, bool)`

GetQuiesceOk returns a tuple with the Quiesce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuiesce

`func (o *HyperflexReplicationPlatDatastorePairAllOf) SetQuiesce(v bool)`

SetQuiesce sets Quiesce field to given value.

### HasQuiesce

`func (o *HyperflexReplicationPlatDatastorePairAllOf) HasQuiesce() bool`

HasQuiesce returns a boolean if a field has been set.

### GetReplicationIntervalInMinutes

`func (o *HyperflexReplicationPlatDatastorePairAllOf) GetReplicationIntervalInMinutes() int64`

GetReplicationIntervalInMinutes returns the ReplicationIntervalInMinutes field if non-nil, zero value otherwise.

### GetReplicationIntervalInMinutesOk

`func (o *HyperflexReplicationPlatDatastorePairAllOf) GetReplicationIntervalInMinutesOk() (*int64, bool)`

GetReplicationIntervalInMinutesOk returns a tuple with the ReplicationIntervalInMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationIntervalInMinutes

`func (o *HyperflexReplicationPlatDatastorePairAllOf) SetReplicationIntervalInMinutes(v int64)`

SetReplicationIntervalInMinutes sets ReplicationIntervalInMinutes field to given value.

### HasReplicationIntervalInMinutes

`func (o *HyperflexReplicationPlatDatastorePairAllOf) HasReplicationIntervalInMinutes() bool`

HasReplicationIntervalInMinutes returns a boolean if a field has been set.

### GetSourceds

`func (o *HyperflexReplicationPlatDatastorePairAllOf) GetSourceds() HyperflexReplicationPlatDatastore`

GetSourceds returns the Sourceds field if non-nil, zero value otherwise.

### GetSourcedsOk

`func (o *HyperflexReplicationPlatDatastorePairAllOf) GetSourcedsOk() (*HyperflexReplicationPlatDatastore, bool)`

GetSourcedsOk returns a tuple with the Sourceds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceds

`func (o *HyperflexReplicationPlatDatastorePairAllOf) SetSourceds(v HyperflexReplicationPlatDatastore)`

SetSourceds sets Sourceds field to given value.

### HasSourceds

`func (o *HyperflexReplicationPlatDatastorePairAllOf) HasSourceds() bool`

HasSourceds returns a boolean if a field has been set.

### SetSourcedsNil

`func (o *HyperflexReplicationPlatDatastorePairAllOf) SetSourcedsNil(b bool)`

 SetSourcedsNil sets the value for Sourceds to be an explicit nil

### UnsetSourceds
`func (o *HyperflexReplicationPlatDatastorePairAllOf) UnsetSourceds()`

UnsetSourceds ensures that no value is present for Sourceds, not even an explicit nil
### GetStorageOnly

`func (o *HyperflexReplicationPlatDatastorePairAllOf) GetStorageOnly() bool`

GetStorageOnly returns the StorageOnly field if non-nil, zero value otherwise.

### GetStorageOnlyOk

`func (o *HyperflexReplicationPlatDatastorePairAllOf) GetStorageOnlyOk() (*bool, bool)`

GetStorageOnlyOk returns a tuple with the StorageOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageOnly

`func (o *HyperflexReplicationPlatDatastorePairAllOf) SetStorageOnly(v bool)`

SetStorageOnly sets StorageOnly field to given value.

### HasStorageOnly

`func (o *HyperflexReplicationPlatDatastorePairAllOf) HasStorageOnly() bool`

HasStorageOnly returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


