# HyperflexBackupClusterAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hyperflex.BackupCluster"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hyperflex.BackupCluster"]
**BackupDataStore** | Pointer to **string** | Defines the backup source cluster and its references. | [optional] [readonly] 
**SrcCluster** | Pointer to [**HyperflexClusterRelationship**](HyperflexClusterRelationship.md) |  | [optional] 
**TgtCluster** | Pointer to [**HyperflexClusterRelationship**](HyperflexClusterRelationship.md) |  | [optional] 

## Methods

### NewHyperflexBackupClusterAllOf

`func NewHyperflexBackupClusterAllOf(classId string, objectType string, ) *HyperflexBackupClusterAllOf`

NewHyperflexBackupClusterAllOf instantiates a new HyperflexBackupClusterAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexBackupClusterAllOfWithDefaults

`func NewHyperflexBackupClusterAllOfWithDefaults() *HyperflexBackupClusterAllOf`

NewHyperflexBackupClusterAllOfWithDefaults instantiates a new HyperflexBackupClusterAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HyperflexBackupClusterAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HyperflexBackupClusterAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HyperflexBackupClusterAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HyperflexBackupClusterAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HyperflexBackupClusterAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HyperflexBackupClusterAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetBackupDataStore

`func (o *HyperflexBackupClusterAllOf) GetBackupDataStore() string`

GetBackupDataStore returns the BackupDataStore field if non-nil, zero value otherwise.

### GetBackupDataStoreOk

`func (o *HyperflexBackupClusterAllOf) GetBackupDataStoreOk() (*string, bool)`

GetBackupDataStoreOk returns a tuple with the BackupDataStore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupDataStore

`func (o *HyperflexBackupClusterAllOf) SetBackupDataStore(v string)`

SetBackupDataStore sets BackupDataStore field to given value.

### HasBackupDataStore

`func (o *HyperflexBackupClusterAllOf) HasBackupDataStore() bool`

HasBackupDataStore returns a boolean if a field has been set.

### GetSrcCluster

`func (o *HyperflexBackupClusterAllOf) GetSrcCluster() HyperflexClusterRelationship`

GetSrcCluster returns the SrcCluster field if non-nil, zero value otherwise.

### GetSrcClusterOk

`func (o *HyperflexBackupClusterAllOf) GetSrcClusterOk() (*HyperflexClusterRelationship, bool)`

GetSrcClusterOk returns a tuple with the SrcCluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcCluster

`func (o *HyperflexBackupClusterAllOf) SetSrcCluster(v HyperflexClusterRelationship)`

SetSrcCluster sets SrcCluster field to given value.

### HasSrcCluster

`func (o *HyperflexBackupClusterAllOf) HasSrcCluster() bool`

HasSrcCluster returns a boolean if a field has been set.

### GetTgtCluster

`func (o *HyperflexBackupClusterAllOf) GetTgtCluster() HyperflexClusterRelationship`

GetTgtCluster returns the TgtCluster field if non-nil, zero value otherwise.

### GetTgtClusterOk

`func (o *HyperflexBackupClusterAllOf) GetTgtClusterOk() (*HyperflexClusterRelationship, bool)`

GetTgtClusterOk returns a tuple with the TgtCluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTgtCluster

`func (o *HyperflexBackupClusterAllOf) SetTgtCluster(v HyperflexClusterRelationship)`

SetTgtCluster sets TgtCluster field to given value.

### HasTgtCluster

`func (o *HyperflexBackupClusterAllOf) HasTgtCluster() bool`

HasTgtCluster returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


