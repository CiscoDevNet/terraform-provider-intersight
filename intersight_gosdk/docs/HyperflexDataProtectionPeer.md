# HyperflexDataProtectionPeer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hyperflex.DataProtectionPeer"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hyperflex.DataProtectionPeer"]
**Er** | Pointer to [**NullableHyperflexEntityReference**](HyperflexEntityReference.md) |  | [optional] 
**PeerInfo** | Pointer to [**NullableHyperflexReplicationPeerInfo**](HyperflexReplicationPeerInfo.md) |  | [optional] 
**SrcCluster** | Pointer to [**HyperflexClusterRelationship**](HyperflexClusterRelationship.md) |  | [optional] 

## Methods

### NewHyperflexDataProtectionPeer

`func NewHyperflexDataProtectionPeer(classId string, objectType string, ) *HyperflexDataProtectionPeer`

NewHyperflexDataProtectionPeer instantiates a new HyperflexDataProtectionPeer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexDataProtectionPeerWithDefaults

`func NewHyperflexDataProtectionPeerWithDefaults() *HyperflexDataProtectionPeer`

NewHyperflexDataProtectionPeerWithDefaults instantiates a new HyperflexDataProtectionPeer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HyperflexDataProtectionPeer) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HyperflexDataProtectionPeer) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HyperflexDataProtectionPeer) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HyperflexDataProtectionPeer) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HyperflexDataProtectionPeer) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HyperflexDataProtectionPeer) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetEr

`func (o *HyperflexDataProtectionPeer) GetEr() HyperflexEntityReference`

GetEr returns the Er field if non-nil, zero value otherwise.

### GetErOk

`func (o *HyperflexDataProtectionPeer) GetErOk() (*HyperflexEntityReference, bool)`

GetErOk returns a tuple with the Er field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEr

`func (o *HyperflexDataProtectionPeer) SetEr(v HyperflexEntityReference)`

SetEr sets Er field to given value.

### HasEr

`func (o *HyperflexDataProtectionPeer) HasEr() bool`

HasEr returns a boolean if a field has been set.

### SetErNil

`func (o *HyperflexDataProtectionPeer) SetErNil(b bool)`

 SetErNil sets the value for Er to be an explicit nil

### UnsetEr
`func (o *HyperflexDataProtectionPeer) UnsetEr()`

UnsetEr ensures that no value is present for Er, not even an explicit nil
### GetPeerInfo

`func (o *HyperflexDataProtectionPeer) GetPeerInfo() HyperflexReplicationPeerInfo`

GetPeerInfo returns the PeerInfo field if non-nil, zero value otherwise.

### GetPeerInfoOk

`func (o *HyperflexDataProtectionPeer) GetPeerInfoOk() (*HyperflexReplicationPeerInfo, bool)`

GetPeerInfoOk returns a tuple with the PeerInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerInfo

`func (o *HyperflexDataProtectionPeer) SetPeerInfo(v HyperflexReplicationPeerInfo)`

SetPeerInfo sets PeerInfo field to given value.

### HasPeerInfo

`func (o *HyperflexDataProtectionPeer) HasPeerInfo() bool`

HasPeerInfo returns a boolean if a field has been set.

### SetPeerInfoNil

`func (o *HyperflexDataProtectionPeer) SetPeerInfoNil(b bool)`

 SetPeerInfoNil sets the value for PeerInfo to be an explicit nil

### UnsetPeerInfo
`func (o *HyperflexDataProtectionPeer) UnsetPeerInfo()`

UnsetPeerInfo ensures that no value is present for PeerInfo, not even an explicit nil
### GetSrcCluster

`func (o *HyperflexDataProtectionPeer) GetSrcCluster() HyperflexClusterRelationship`

GetSrcCluster returns the SrcCluster field if non-nil, zero value otherwise.

### GetSrcClusterOk

`func (o *HyperflexDataProtectionPeer) GetSrcClusterOk() (*HyperflexClusterRelationship, bool)`

GetSrcClusterOk returns a tuple with the SrcCluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcCluster

`func (o *HyperflexDataProtectionPeer) SetSrcCluster(v HyperflexClusterRelationship)`

SetSrcCluster sets SrcCluster field to given value.

### HasSrcCluster

`func (o *HyperflexDataProtectionPeer) HasSrcCluster() bool`

HasSrcCluster returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


