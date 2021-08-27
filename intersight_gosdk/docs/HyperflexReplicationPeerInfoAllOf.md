# HyperflexReplicationPeerInfoAllOf

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

### NewHyperflexReplicationPeerInfoAllOf

`func NewHyperflexReplicationPeerInfoAllOf(classId string, objectType string, ) *HyperflexReplicationPeerInfoAllOf`

NewHyperflexReplicationPeerInfoAllOf instantiates a new HyperflexReplicationPeerInfoAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexReplicationPeerInfoAllOfWithDefaults

`func NewHyperflexReplicationPeerInfoAllOfWithDefaults() *HyperflexReplicationPeerInfoAllOf`

NewHyperflexReplicationPeerInfoAllOfWithDefaults instantiates a new HyperflexReplicationPeerInfoAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HyperflexReplicationPeerInfoAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HyperflexReplicationPeerInfoAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HyperflexReplicationPeerInfoAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HyperflexReplicationPeerInfoAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HyperflexReplicationPeerInfoAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HyperflexReplicationPeerInfoAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDatastores

`func (o *HyperflexReplicationPeerInfoAllOf) GetDatastores() []HyperflexReplicationPlatDatastorePair`

GetDatastores returns the Datastores field if non-nil, zero value otherwise.

### GetDatastoresOk

`func (o *HyperflexReplicationPeerInfoAllOf) GetDatastoresOk() (*[]HyperflexReplicationPlatDatastorePair, bool)`

GetDatastoresOk returns a tuple with the Datastores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastores

`func (o *HyperflexReplicationPeerInfoAllOf) SetDatastores(v []HyperflexReplicationPlatDatastorePair)`

SetDatastores sets Datastores field to given value.

### HasDatastores

`func (o *HyperflexReplicationPeerInfoAllOf) HasDatastores() bool`

HasDatastores returns a boolean if a field has been set.

### SetDatastoresNil

`func (o *HyperflexReplicationPeerInfoAllOf) SetDatastoresNil(b bool)`

 SetDatastoresNil sets the value for Datastores to be an explicit nil

### UnsetDatastores
`func (o *HyperflexReplicationPeerInfoAllOf) UnsetDatastores()`

UnsetDatastores ensures that no value is present for Datastores, not even an explicit nil
### GetDcip

`func (o *HyperflexReplicationPeerInfoAllOf) GetDcip() string`

GetDcip returns the Dcip field if non-nil, zero value otherwise.

### GetDcipOk

`func (o *HyperflexReplicationPeerInfoAllOf) GetDcipOk() (*string, bool)`

GetDcipOk returns a tuple with the Dcip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDcip

`func (o *HyperflexReplicationPeerInfoAllOf) SetDcip(v string)`

SetDcip sets Dcip field to given value.

### HasDcip

`func (o *HyperflexReplicationPeerInfoAllOf) HasDcip() bool`

HasDcip returns a boolean if a field has been set.

### GetEr

`func (o *HyperflexReplicationPeerInfoAllOf) GetEr() HyperflexEntityReference`

GetEr returns the Er field if non-nil, zero value otherwise.

### GetErOk

`func (o *HyperflexReplicationPeerInfoAllOf) GetErOk() (*HyperflexEntityReference, bool)`

GetErOk returns a tuple with the Er field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEr

`func (o *HyperflexReplicationPeerInfoAllOf) SetEr(v HyperflexEntityReference)`

SetEr sets Er field to given value.

### HasEr

`func (o *HyperflexReplicationPeerInfoAllOf) HasEr() bool`

HasEr returns a boolean if a field has been set.

### SetErNil

`func (o *HyperflexReplicationPeerInfoAllOf) SetErNil(b bool)`

 SetErNil sets the value for Er to be an explicit nil

### UnsetEr
`func (o *HyperflexReplicationPeerInfoAllOf) UnsetEr()`

UnsetEr ensures that no value is present for Er, not even an explicit nil
### GetMcip

`func (o *HyperflexReplicationPeerInfoAllOf) GetMcip() string`

GetMcip returns the Mcip field if non-nil, zero value otherwise.

### GetMcipOk

`func (o *HyperflexReplicationPeerInfoAllOf) GetMcipOk() (*string, bool)`

GetMcipOk returns a tuple with the Mcip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMcip

`func (o *HyperflexReplicationPeerInfoAllOf) SetMcip(v string)`

SetMcip sets Mcip field to given value.

### HasMcip

`func (o *HyperflexReplicationPeerInfoAllOf) HasMcip() bool`

HasMcip returns a boolean if a field has been set.

### GetPorts

`func (o *HyperflexReplicationPeerInfoAllOf) GetPorts() []HyperflexPortTypeToPortNumberMap`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *HyperflexReplicationPeerInfoAllOf) GetPortsOk() (*[]HyperflexPortTypeToPortNumberMap, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *HyperflexReplicationPeerInfoAllOf) SetPorts(v []HyperflexPortTypeToPortNumberMap)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *HyperflexReplicationPeerInfoAllOf) HasPorts() bool`

HasPorts returns a boolean if a field has been set.

### SetPortsNil

`func (o *HyperflexReplicationPeerInfoAllOf) SetPortsNil(b bool)`

 SetPortsNil sets the value for Ports to be an explicit nil

### UnsetPorts
`func (o *HyperflexReplicationPeerInfoAllOf) UnsetPorts()`

UnsetPorts ensures that no value is present for Ports, not even an explicit nil
### GetReplCip

`func (o *HyperflexReplicationPeerInfoAllOf) GetReplCip() string`

GetReplCip returns the ReplCip field if non-nil, zero value otherwise.

### GetReplCipOk

`func (o *HyperflexReplicationPeerInfoAllOf) GetReplCipOk() (*string, bool)`

GetReplCipOk returns a tuple with the ReplCip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplCip

`func (o *HyperflexReplicationPeerInfoAllOf) SetReplCip(v string)`

SetReplCip sets ReplCip field to given value.

### HasReplCip

`func (o *HyperflexReplicationPeerInfoAllOf) HasReplCip() bool`

HasReplCip returns a boolean if a field has been set.

### GetStatus

`func (o *HyperflexReplicationPeerInfoAllOf) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *HyperflexReplicationPeerInfoAllOf) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *HyperflexReplicationPeerInfoAllOf) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *HyperflexReplicationPeerInfoAllOf) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusDetails

`func (o *HyperflexReplicationPeerInfoAllOf) GetStatusDetails() string`

GetStatusDetails returns the StatusDetails field if non-nil, zero value otherwise.

### GetStatusDetailsOk

`func (o *HyperflexReplicationPeerInfoAllOf) GetStatusDetailsOk() (*string, bool)`

GetStatusDetailsOk returns a tuple with the StatusDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusDetails

`func (o *HyperflexReplicationPeerInfoAllOf) SetStatusDetails(v string)`

SetStatusDetails sets StatusDetails field to given value.

### HasStatusDetails

`func (o *HyperflexReplicationPeerInfoAllOf) HasStatusDetails() bool`

HasStatusDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


