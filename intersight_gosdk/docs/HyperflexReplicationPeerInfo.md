# HyperflexReplicationPeerInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hyperflex.ReplicationPeerInfo"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hyperflex.ReplicationPeerInfo"]
**Datastores** | Pointer to [**[]HyperflexReplicationPlatDatastorePair**](HyperflexReplicationPlatDatastorePair.md) |  | [optional] 
**Dcip** | Pointer to **string** | Data Cluster IP for the replication peer. | [optional] [readonly] 
**Er** | Pointer to [**NullableHyperflexEntityReference**](HyperflexEntityReference.md) |  | [optional] 
**Mcip** | Pointer to **string** | Management Cluster IP for the replication peer. | [optional] [readonly] 
**Ports** | Pointer to [**[]HyperflexPortTypeToPortNumberMap**](HyperflexPortTypeToPortNumberMap.md) |  | [optional] 
**ReplCip** | Pointer to **string** | Replication Cluster IP for the replication peer. | [optional] [readonly] 
**Status** | Pointer to **string** | Peer Cluster Status for the replication peer. | [optional] [readonly] 
**StatusDetails** | Pointer to **string** | Peer Cluster Status Details for the replication peer. | [optional] [readonly] 

## Methods

### NewHyperflexReplicationPeerInfo

`func NewHyperflexReplicationPeerInfo(classId string, objectType string, ) *HyperflexReplicationPeerInfo`

NewHyperflexReplicationPeerInfo instantiates a new HyperflexReplicationPeerInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexReplicationPeerInfoWithDefaults

`func NewHyperflexReplicationPeerInfoWithDefaults() *HyperflexReplicationPeerInfo`

NewHyperflexReplicationPeerInfoWithDefaults instantiates a new HyperflexReplicationPeerInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HyperflexReplicationPeerInfo) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HyperflexReplicationPeerInfo) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HyperflexReplicationPeerInfo) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HyperflexReplicationPeerInfo) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HyperflexReplicationPeerInfo) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HyperflexReplicationPeerInfo) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDatastores

`func (o *HyperflexReplicationPeerInfo) GetDatastores() []HyperflexReplicationPlatDatastorePair`

GetDatastores returns the Datastores field if non-nil, zero value otherwise.

### GetDatastoresOk

`func (o *HyperflexReplicationPeerInfo) GetDatastoresOk() (*[]HyperflexReplicationPlatDatastorePair, bool)`

GetDatastoresOk returns a tuple with the Datastores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastores

`func (o *HyperflexReplicationPeerInfo) SetDatastores(v []HyperflexReplicationPlatDatastorePair)`

SetDatastores sets Datastores field to given value.

### HasDatastores

`func (o *HyperflexReplicationPeerInfo) HasDatastores() bool`

HasDatastores returns a boolean if a field has been set.

### SetDatastoresNil

`func (o *HyperflexReplicationPeerInfo) SetDatastoresNil(b bool)`

 SetDatastoresNil sets the value for Datastores to be an explicit nil

### UnsetDatastores
`func (o *HyperflexReplicationPeerInfo) UnsetDatastores()`

UnsetDatastores ensures that no value is present for Datastores, not even an explicit nil
### GetDcip

`func (o *HyperflexReplicationPeerInfo) GetDcip() string`

GetDcip returns the Dcip field if non-nil, zero value otherwise.

### GetDcipOk

`func (o *HyperflexReplicationPeerInfo) GetDcipOk() (*string, bool)`

GetDcipOk returns a tuple with the Dcip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDcip

`func (o *HyperflexReplicationPeerInfo) SetDcip(v string)`

SetDcip sets Dcip field to given value.

### HasDcip

`func (o *HyperflexReplicationPeerInfo) HasDcip() bool`

HasDcip returns a boolean if a field has been set.

### GetEr

`func (o *HyperflexReplicationPeerInfo) GetEr() HyperflexEntityReference`

GetEr returns the Er field if non-nil, zero value otherwise.

### GetErOk

`func (o *HyperflexReplicationPeerInfo) GetErOk() (*HyperflexEntityReference, bool)`

GetErOk returns a tuple with the Er field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEr

`func (o *HyperflexReplicationPeerInfo) SetEr(v HyperflexEntityReference)`

SetEr sets Er field to given value.

### HasEr

`func (o *HyperflexReplicationPeerInfo) HasEr() bool`

HasEr returns a boolean if a field has been set.

### SetErNil

`func (o *HyperflexReplicationPeerInfo) SetErNil(b bool)`

 SetErNil sets the value for Er to be an explicit nil

### UnsetEr
`func (o *HyperflexReplicationPeerInfo) UnsetEr()`

UnsetEr ensures that no value is present for Er, not even an explicit nil
### GetMcip

`func (o *HyperflexReplicationPeerInfo) GetMcip() string`

GetMcip returns the Mcip field if non-nil, zero value otherwise.

### GetMcipOk

`func (o *HyperflexReplicationPeerInfo) GetMcipOk() (*string, bool)`

GetMcipOk returns a tuple with the Mcip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMcip

`func (o *HyperflexReplicationPeerInfo) SetMcip(v string)`

SetMcip sets Mcip field to given value.

### HasMcip

`func (o *HyperflexReplicationPeerInfo) HasMcip() bool`

HasMcip returns a boolean if a field has been set.

### GetPorts

`func (o *HyperflexReplicationPeerInfo) GetPorts() []HyperflexPortTypeToPortNumberMap`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *HyperflexReplicationPeerInfo) GetPortsOk() (*[]HyperflexPortTypeToPortNumberMap, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *HyperflexReplicationPeerInfo) SetPorts(v []HyperflexPortTypeToPortNumberMap)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *HyperflexReplicationPeerInfo) HasPorts() bool`

HasPorts returns a boolean if a field has been set.

### SetPortsNil

`func (o *HyperflexReplicationPeerInfo) SetPortsNil(b bool)`

 SetPortsNil sets the value for Ports to be an explicit nil

### UnsetPorts
`func (o *HyperflexReplicationPeerInfo) UnsetPorts()`

UnsetPorts ensures that no value is present for Ports, not even an explicit nil
### GetReplCip

`func (o *HyperflexReplicationPeerInfo) GetReplCip() string`

GetReplCip returns the ReplCip field if non-nil, zero value otherwise.

### GetReplCipOk

`func (o *HyperflexReplicationPeerInfo) GetReplCipOk() (*string, bool)`

GetReplCipOk returns a tuple with the ReplCip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplCip

`func (o *HyperflexReplicationPeerInfo) SetReplCip(v string)`

SetReplCip sets ReplCip field to given value.

### HasReplCip

`func (o *HyperflexReplicationPeerInfo) HasReplCip() bool`

HasReplCip returns a boolean if a field has been set.

### GetStatus

`func (o *HyperflexReplicationPeerInfo) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *HyperflexReplicationPeerInfo) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *HyperflexReplicationPeerInfo) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *HyperflexReplicationPeerInfo) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusDetails

`func (o *HyperflexReplicationPeerInfo) GetStatusDetails() string`

GetStatusDetails returns the StatusDetails field if non-nil, zero value otherwise.

### GetStatusDetailsOk

`func (o *HyperflexReplicationPeerInfo) GetStatusDetailsOk() (*string, bool)`

GetStatusDetailsOk returns a tuple with the StatusDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusDetails

`func (o *HyperflexReplicationPeerInfo) SetStatusDetails(v string)`

SetStatusDetails sets StatusDetails field to given value.

### HasStatusDetails

`func (o *HyperflexReplicationPeerInfo) HasStatusDetails() bool`

HasStatusDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


