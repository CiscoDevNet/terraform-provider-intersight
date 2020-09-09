# HyperflexHealth

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ArbitrationServiceState** | Pointer to **string** |  | [optional] [readonly] [default to "NOT_AVAILABLE"]
**DataReplicationCompliance** | Pointer to **string** |  | [optional] [readonly] [default to "UNKNOWN"]
**ResiliencyDetails** | Pointer to [**HyperflexHxResiliencyInfoDt**](hyperflex.HxResiliencyInfoDt.md) |  | [optional] 
**State** | Pointer to **string** |  | [optional] [readonly] [default to "UNKNOWN"]
**Uuid** | Pointer to **string** |  | [optional] [readonly] 
**ZkHealth** | Pointer to **string** |  | [optional] [readonly] [default to "NOT_AVAILABLE"]
**ZoneResiliencyList** | Pointer to [**[]HyperflexHxZoneResiliencyInfoDt**](hyperflex.HxZoneResiliencyInfoDt.md) |  | [optional] 
**Cluster** | Pointer to [**HyperflexClusterRelationship**](hyperflex.Cluster.Relationship.md) |  | [optional] 

## Methods

### NewHyperflexHealth

`func NewHyperflexHealth() *HyperflexHealth`

NewHyperflexHealth instantiates a new HyperflexHealth object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexHealthWithDefaults

`func NewHyperflexHealthWithDefaults() *HyperflexHealth`

NewHyperflexHealthWithDefaults instantiates a new HyperflexHealth object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArbitrationServiceState

`func (o *HyperflexHealth) GetArbitrationServiceState() string`

GetArbitrationServiceState returns the ArbitrationServiceState field if non-nil, zero value otherwise.

### GetArbitrationServiceStateOk

`func (o *HyperflexHealth) GetArbitrationServiceStateOk() (*string, bool)`

GetArbitrationServiceStateOk returns a tuple with the ArbitrationServiceState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArbitrationServiceState

`func (o *HyperflexHealth) SetArbitrationServiceState(v string)`

SetArbitrationServiceState sets ArbitrationServiceState field to given value.

### HasArbitrationServiceState

`func (o *HyperflexHealth) HasArbitrationServiceState() bool`

HasArbitrationServiceState returns a boolean if a field has been set.

### GetDataReplicationCompliance

`func (o *HyperflexHealth) GetDataReplicationCompliance() string`

GetDataReplicationCompliance returns the DataReplicationCompliance field if non-nil, zero value otherwise.

### GetDataReplicationComplianceOk

`func (o *HyperflexHealth) GetDataReplicationComplianceOk() (*string, bool)`

GetDataReplicationComplianceOk returns a tuple with the DataReplicationCompliance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataReplicationCompliance

`func (o *HyperflexHealth) SetDataReplicationCompliance(v string)`

SetDataReplicationCompliance sets DataReplicationCompliance field to given value.

### HasDataReplicationCompliance

`func (o *HyperflexHealth) HasDataReplicationCompliance() bool`

HasDataReplicationCompliance returns a boolean if a field has been set.

### GetResiliencyDetails

`func (o *HyperflexHealth) GetResiliencyDetails() HyperflexHxResiliencyInfoDt`

GetResiliencyDetails returns the ResiliencyDetails field if non-nil, zero value otherwise.

### GetResiliencyDetailsOk

`func (o *HyperflexHealth) GetResiliencyDetailsOk() (*HyperflexHxResiliencyInfoDt, bool)`

GetResiliencyDetailsOk returns a tuple with the ResiliencyDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResiliencyDetails

`func (o *HyperflexHealth) SetResiliencyDetails(v HyperflexHxResiliencyInfoDt)`

SetResiliencyDetails sets ResiliencyDetails field to given value.

### HasResiliencyDetails

`func (o *HyperflexHealth) HasResiliencyDetails() bool`

HasResiliencyDetails returns a boolean if a field has been set.

### GetState

`func (o *HyperflexHealth) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *HyperflexHealth) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *HyperflexHealth) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *HyperflexHealth) HasState() bool`

HasState returns a boolean if a field has been set.

### GetUuid

`func (o *HyperflexHealth) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *HyperflexHealth) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *HyperflexHealth) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *HyperflexHealth) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetZkHealth

`func (o *HyperflexHealth) GetZkHealth() string`

GetZkHealth returns the ZkHealth field if non-nil, zero value otherwise.

### GetZkHealthOk

`func (o *HyperflexHealth) GetZkHealthOk() (*string, bool)`

GetZkHealthOk returns a tuple with the ZkHealth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZkHealth

`func (o *HyperflexHealth) SetZkHealth(v string)`

SetZkHealth sets ZkHealth field to given value.

### HasZkHealth

`func (o *HyperflexHealth) HasZkHealth() bool`

HasZkHealth returns a boolean if a field has been set.

### GetZoneResiliencyList

`func (o *HyperflexHealth) GetZoneResiliencyList() []HyperflexHxZoneResiliencyInfoDt`

GetZoneResiliencyList returns the ZoneResiliencyList field if non-nil, zero value otherwise.

### GetZoneResiliencyListOk

`func (o *HyperflexHealth) GetZoneResiliencyListOk() (*[]HyperflexHxZoneResiliencyInfoDt, bool)`

GetZoneResiliencyListOk returns a tuple with the ZoneResiliencyList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneResiliencyList

`func (o *HyperflexHealth) SetZoneResiliencyList(v []HyperflexHxZoneResiliencyInfoDt)`

SetZoneResiliencyList sets ZoneResiliencyList field to given value.

### HasZoneResiliencyList

`func (o *HyperflexHealth) HasZoneResiliencyList() bool`

HasZoneResiliencyList returns a boolean if a field has been set.

### GetCluster

`func (o *HyperflexHealth) GetCluster() HyperflexClusterRelationship`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *HyperflexHealth) GetClusterOk() (*HyperflexClusterRelationship, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *HyperflexHealth) SetCluster(v HyperflexClusterRelationship)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *HyperflexHealth) HasCluster() bool`

HasCluster returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


