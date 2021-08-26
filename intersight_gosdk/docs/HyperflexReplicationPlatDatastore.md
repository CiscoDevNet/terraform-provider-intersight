# HyperflexReplicationPlatDatastore

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hyperflex.ReplicationPlatDatastore"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hyperflex.ReplicationPlatDatastore"]
**ClusterEr** | Pointer to [**NullableHyperflexEntityReference**](HyperflexEntityReference.md) |  | [optional] 
**DatastoreEr** | Pointer to [**NullableHyperflexEntityReference**](HyperflexEntityReference.md) |  | [optional] 

## Methods

### NewHyperflexReplicationPlatDatastore

`func NewHyperflexReplicationPlatDatastore(classId string, objectType string, ) *HyperflexReplicationPlatDatastore`

NewHyperflexReplicationPlatDatastore instantiates a new HyperflexReplicationPlatDatastore object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexReplicationPlatDatastoreWithDefaults

`func NewHyperflexReplicationPlatDatastoreWithDefaults() *HyperflexReplicationPlatDatastore`

NewHyperflexReplicationPlatDatastoreWithDefaults instantiates a new HyperflexReplicationPlatDatastore object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HyperflexReplicationPlatDatastore) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HyperflexReplicationPlatDatastore) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HyperflexReplicationPlatDatastore) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HyperflexReplicationPlatDatastore) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HyperflexReplicationPlatDatastore) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HyperflexReplicationPlatDatastore) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetClusterEr

`func (o *HyperflexReplicationPlatDatastore) GetClusterEr() HyperflexEntityReference`

GetClusterEr returns the ClusterEr field if non-nil, zero value otherwise.

### GetClusterErOk

`func (o *HyperflexReplicationPlatDatastore) GetClusterErOk() (*HyperflexEntityReference, bool)`

GetClusterErOk returns a tuple with the ClusterEr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterEr

`func (o *HyperflexReplicationPlatDatastore) SetClusterEr(v HyperflexEntityReference)`

SetClusterEr sets ClusterEr field to given value.

### HasClusterEr

`func (o *HyperflexReplicationPlatDatastore) HasClusterEr() bool`

HasClusterEr returns a boolean if a field has been set.

### SetClusterErNil

`func (o *HyperflexReplicationPlatDatastore) SetClusterErNil(b bool)`

 SetClusterErNil sets the value for ClusterEr to be an explicit nil

### UnsetClusterEr
`func (o *HyperflexReplicationPlatDatastore) UnsetClusterEr()`

UnsetClusterEr ensures that no value is present for ClusterEr, not even an explicit nil
### GetDatastoreEr

`func (o *HyperflexReplicationPlatDatastore) GetDatastoreEr() HyperflexEntityReference`

GetDatastoreEr returns the DatastoreEr field if non-nil, zero value otherwise.

### GetDatastoreErOk

`func (o *HyperflexReplicationPlatDatastore) GetDatastoreErOk() (*HyperflexEntityReference, bool)`

GetDatastoreErOk returns a tuple with the DatastoreEr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastoreEr

`func (o *HyperflexReplicationPlatDatastore) SetDatastoreEr(v HyperflexEntityReference)`

SetDatastoreEr sets DatastoreEr field to given value.

### HasDatastoreEr

`func (o *HyperflexReplicationPlatDatastore) HasDatastoreEr() bool`

HasDatastoreEr returns a boolean if a field has been set.

### SetDatastoreErNil

`func (o *HyperflexReplicationPlatDatastore) SetDatastoreErNil(b bool)`

 SetDatastoreErNil sets the value for DatastoreEr to be an explicit nil

### UnsetDatastoreEr
`func (o *HyperflexReplicationPlatDatastore) UnsetDatastoreEr()`

UnsetDatastoreEr ensures that no value is present for DatastoreEr, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


