# HyperflexProtectedClusterAllOf

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
**DatastoreStatistic** | Pointer to [**HyperflexDatastoreStatisticRelationship**](HyperflexDatastoreStatisticRelationship.md) |  | [optional] 
**SrcCluster** | Pointer to [**HyperflexClusterRelationship**](HyperflexClusterRelationship.md) |  | [optional] 
**TgtCluster** | Pointer to [**HyperflexClusterRelationship**](HyperflexClusterRelationship.md) |  | [optional] 

## Methods

### NewHyperflexProtectedClusterAllOf

`func NewHyperflexProtectedClusterAllOf(classId string, objectType string, ) *HyperflexProtectedClusterAllOf`

NewHyperflexProtectedClusterAllOf instantiates a new HyperflexProtectedClusterAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexProtectedClusterAllOfWithDefaults

`func NewHyperflexProtectedClusterAllOfWithDefaults() *HyperflexProtectedClusterAllOf`

NewHyperflexProtectedClusterAllOfWithDefaults instantiates a new HyperflexProtectedClusterAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HyperflexProtectedClusterAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HyperflexProtectedClusterAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HyperflexProtectedClusterAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HyperflexProtectedClusterAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HyperflexProtectedClusterAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HyperflexProtectedClusterAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetHxVersion

`func (o *HyperflexProtectedClusterAllOf) GetHxVersion() string`

GetHxVersion returns the HxVersion field if non-nil, zero value otherwise.

### GetHxVersionOk

`func (o *HyperflexProtectedClusterAllOf) GetHxVersionOk() (*string, bool)`

GetHxVersionOk returns a tuple with the HxVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHxVersion

`func (o *HyperflexProtectedClusterAllOf) SetHxVersion(v string)`

SetHxVersion sets HxVersion field to given value.

### HasHxVersion

`func (o *HyperflexProtectedClusterAllOf) HasHxVersion() bool`

HasHxVersion returns a boolean if a field has been set.

### GetHypervisorVersion

`func (o *HyperflexProtectedClusterAllOf) GetHypervisorVersion() string`

GetHypervisorVersion returns the HypervisorVersion field if non-nil, zero value otherwise.

### GetHypervisorVersionOk

`func (o *HyperflexProtectedClusterAllOf) GetHypervisorVersionOk() (*string, bool)`

GetHypervisorVersionOk returns a tuple with the HypervisorVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHypervisorVersion

`func (o *HyperflexProtectedClusterAllOf) SetHypervisorVersion(v string)`

SetHypervisorVersion sets HypervisorVersion field to given value.

### HasHypervisorVersion

`func (o *HyperflexProtectedClusterAllOf) HasHypervisorVersion() bool`

HasHypervisorVersion returns a boolean if a field has been set.

### GetProtectedDatastoreName

`func (o *HyperflexProtectedClusterAllOf) GetProtectedDatastoreName() string`

GetProtectedDatastoreName returns the ProtectedDatastoreName field if non-nil, zero value otherwise.

### GetProtectedDatastoreNameOk

`func (o *HyperflexProtectedClusterAllOf) GetProtectedDatastoreNameOk() (*string, bool)`

GetProtectedDatastoreNameOk returns a tuple with the ProtectedDatastoreName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectedDatastoreName

`func (o *HyperflexProtectedClusterAllOf) SetProtectedDatastoreName(v string)`

SetProtectedDatastoreName sets ProtectedDatastoreName field to given value.

### HasProtectedDatastoreName

`func (o *HyperflexProtectedClusterAllOf) HasProtectedDatastoreName() bool`

HasProtectedDatastoreName returns a boolean if a field has been set.

### GetProtectedVmsCount

`func (o *HyperflexProtectedClusterAllOf) GetProtectedVmsCount() int64`

GetProtectedVmsCount returns the ProtectedVmsCount field if non-nil, zero value otherwise.

### GetProtectedVmsCountOk

`func (o *HyperflexProtectedClusterAllOf) GetProtectedVmsCountOk() (*int64, bool)`

GetProtectedVmsCountOk returns a tuple with the ProtectedVmsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectedVmsCount

`func (o *HyperflexProtectedClusterAllOf) SetProtectedVmsCount(v int64)`

SetProtectedVmsCount sets ProtectedVmsCount field to given value.

### HasProtectedVmsCount

`func (o *HyperflexProtectedClusterAllOf) HasProtectedVmsCount() bool`

HasProtectedVmsCount returns a boolean if a field has been set.

### GetSourceClusterName

`func (o *HyperflexProtectedClusterAllOf) GetSourceClusterName() string`

GetSourceClusterName returns the SourceClusterName field if non-nil, zero value otherwise.

### GetSourceClusterNameOk

`func (o *HyperflexProtectedClusterAllOf) GetSourceClusterNameOk() (*string, bool)`

GetSourceClusterNameOk returns a tuple with the SourceClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceClusterName

`func (o *HyperflexProtectedClusterAllOf) SetSourceClusterName(v string)`

SetSourceClusterName sets SourceClusterName field to given value.

### HasSourceClusterName

`func (o *HyperflexProtectedClusterAllOf) HasSourceClusterName() bool`

HasSourceClusterName returns a boolean if a field has been set.

### GetTargetClusterName

`func (o *HyperflexProtectedClusterAllOf) GetTargetClusterName() string`

GetTargetClusterName returns the TargetClusterName field if non-nil, zero value otherwise.

### GetTargetClusterNameOk

`func (o *HyperflexProtectedClusterAllOf) GetTargetClusterNameOk() (*string, bool)`

GetTargetClusterNameOk returns a tuple with the TargetClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetClusterName

`func (o *HyperflexProtectedClusterAllOf) SetTargetClusterName(v string)`

SetTargetClusterName sets TargetClusterName field to given value.

### HasTargetClusterName

`func (o *HyperflexProtectedClusterAllOf) HasTargetClusterName() bool`

HasTargetClusterName returns a boolean if a field has been set.

### GetTargetDatastoreName

`func (o *HyperflexProtectedClusterAllOf) GetTargetDatastoreName() string`

GetTargetDatastoreName returns the TargetDatastoreName field if non-nil, zero value otherwise.

### GetTargetDatastoreNameOk

`func (o *HyperflexProtectedClusterAllOf) GetTargetDatastoreNameOk() (*string, bool)`

GetTargetDatastoreNameOk returns a tuple with the TargetDatastoreName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetDatastoreName

`func (o *HyperflexProtectedClusterAllOf) SetTargetDatastoreName(v string)`

SetTargetDatastoreName sets TargetDatastoreName field to given value.

### HasTargetDatastoreName

`func (o *HyperflexProtectedClusterAllOf) HasTargetDatastoreName() bool`

HasTargetDatastoreName returns a boolean if a field has been set.

### GetTargetDatastoreUtilization

`func (o *HyperflexProtectedClusterAllOf) GetTargetDatastoreUtilization() float32`

GetTargetDatastoreUtilization returns the TargetDatastoreUtilization field if non-nil, zero value otherwise.

### GetTargetDatastoreUtilizationOk

`func (o *HyperflexProtectedClusterAllOf) GetTargetDatastoreUtilizationOk() (*float32, bool)`

GetTargetDatastoreUtilizationOk returns a tuple with the TargetDatastoreUtilization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetDatastoreUtilization

`func (o *HyperflexProtectedClusterAllOf) SetTargetDatastoreUtilization(v float32)`

SetTargetDatastoreUtilization sets TargetDatastoreUtilization field to given value.

### HasTargetDatastoreUtilization

`func (o *HyperflexProtectedClusterAllOf) HasTargetDatastoreUtilization() bool`

HasTargetDatastoreUtilization returns a boolean if a field has been set.

### GetDatastoreStatistic

`func (o *HyperflexProtectedClusterAllOf) GetDatastoreStatistic() HyperflexDatastoreStatisticRelationship`

GetDatastoreStatistic returns the DatastoreStatistic field if non-nil, zero value otherwise.

### GetDatastoreStatisticOk

`func (o *HyperflexProtectedClusterAllOf) GetDatastoreStatisticOk() (*HyperflexDatastoreStatisticRelationship, bool)`

GetDatastoreStatisticOk returns a tuple with the DatastoreStatistic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastoreStatistic

`func (o *HyperflexProtectedClusterAllOf) SetDatastoreStatistic(v HyperflexDatastoreStatisticRelationship)`

SetDatastoreStatistic sets DatastoreStatistic field to given value.

### HasDatastoreStatistic

`func (o *HyperflexProtectedClusterAllOf) HasDatastoreStatistic() bool`

HasDatastoreStatistic returns a boolean if a field has been set.

### GetSrcCluster

`func (o *HyperflexProtectedClusterAllOf) GetSrcCluster() HyperflexClusterRelationship`

GetSrcCluster returns the SrcCluster field if non-nil, zero value otherwise.

### GetSrcClusterOk

`func (o *HyperflexProtectedClusterAllOf) GetSrcClusterOk() (*HyperflexClusterRelationship, bool)`

GetSrcClusterOk returns a tuple with the SrcCluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcCluster

`func (o *HyperflexProtectedClusterAllOf) SetSrcCluster(v HyperflexClusterRelationship)`

SetSrcCluster sets SrcCluster field to given value.

### HasSrcCluster

`func (o *HyperflexProtectedClusterAllOf) HasSrcCluster() bool`

HasSrcCluster returns a boolean if a field has been set.

### GetTgtCluster

`func (o *HyperflexProtectedClusterAllOf) GetTgtCluster() HyperflexClusterRelationship`

GetTgtCluster returns the TgtCluster field if non-nil, zero value otherwise.

### GetTgtClusterOk

`func (o *HyperflexProtectedClusterAllOf) GetTgtClusterOk() (*HyperflexClusterRelationship, bool)`

GetTgtClusterOk returns a tuple with the TgtCluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTgtCluster

`func (o *HyperflexProtectedClusterAllOf) SetTgtCluster(v HyperflexClusterRelationship)`

SetTgtCluster sets TgtCluster field to given value.

### HasTgtCluster

`func (o *HyperflexProtectedClusterAllOf) HasTgtCluster() bool`

HasTgtCluster returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


