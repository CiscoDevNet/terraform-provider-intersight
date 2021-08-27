# HyperflexHealthAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hyperflex.Health"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hyperflex.Health"]
**ArbitrationServiceState** | Pointer to **string** | The status of the HyperFlex cluster&#39;s connection to the Intersight arbitration service. The arbitration service state is only applicable to 2-node edge clusters. * &#x60;NOT_AVAILABLE&#x60; - The cluster does not require a connection to the arbitration service. * &#x60;UNKNOWN&#x60; - The cluster&#39;s connection state to the arbitration service cannot be determined. * &#x60;ONLINE&#x60; - The cluster is connected to the arbitration service. * &#x60;OFFLINE&#x60; - The cluster is disconnected from the arbitration service. | [optional] [readonly] [default to "NOT_AVAILABLE"]
**DataReplicationCompliance** | Pointer to **string** | The HyperFlex cluster&#39;s compliance to the configured replication factor. It indicates that the compliance has degraded if the number of copies of data is reduced. * &#x60;UNKNOWN&#x60; - The replication compliance of the HyperFlex cluster is not known. * &#x60;COMPLIANT&#x60; - The HyperFlex cluster is compliant with the replication policy. All data on the cluster is replicated according to the configured replication factor. * &#x60;NON_COMPLIANT&#x60; - The HyperFlex cluster is not compliant with the replication policy. Some data on the cluster is not replicated in accordance with the configured replication factor. | [optional] [readonly] [default to "UNKNOWN"]
**ResiliencyDetails** | Pointer to [**NullableHyperflexHxResiliencyInfoDt**](HyperflexHxResiliencyInfoDt.md) |  | [optional] 
**State** | Pointer to **string** | The operational status of the HyperFlex cluster. * &#x60;UNKNOWN&#x60; - The operational status of the cluster cannot be determined. * &#x60;ONLINE&#x60; - The HyperFlex cluster is online and is performing IO operations. * &#x60;OFFLINE&#x60; - The HyperFlex cluster is offline and is not ready to perform IO operations. * &#x60;ENOSPACE&#x60; - The HyperFlex cluster is out of available storage capacity and cannot perform write transactions. * &#x60;READONLY&#x60; - The HyperFlex cluster is not accepting write transactions, but can still display static cluster information. | [optional] [readonly] [default to "UNKNOWN"]
**Uuid** | Pointer to **string** | The unique identifier for the cluster. | [optional] [readonly] 
**ZkHealth** | Pointer to **string** | The health status of the HyperFlex cluster&#39;s zookeeper ensemble. * &#x60;NOT_AVAILABLE&#x60; - The operational status of the ZK ensemble is not provided by the HyperFlex cluster. * &#x60;UNKNOWN&#x60; - The operational status of the ZK ensemble cannot be determined. * &#x60;ONLINE&#x60; - The ZK ensemble is online and operational. * &#x60;OFFLINE&#x60; - The ZK ensemble is offline and not operational. | [optional] [readonly] [default to "NOT_AVAILABLE"]
**ZoneResiliencyList** | Pointer to [**[]HyperflexHxZoneResiliencyInfoDt**](HyperflexHxZoneResiliencyInfoDt.md) |  | [optional] 
**Cluster** | Pointer to [**HyperflexClusterRelationship**](HyperflexClusterRelationship.md) |  | [optional] 

## Methods

### NewHyperflexHealthAllOf

`func NewHyperflexHealthAllOf(classId string, objectType string, ) *HyperflexHealthAllOf`

NewHyperflexHealthAllOf instantiates a new HyperflexHealthAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexHealthAllOfWithDefaults

`func NewHyperflexHealthAllOfWithDefaults() *HyperflexHealthAllOf`

NewHyperflexHealthAllOfWithDefaults instantiates a new HyperflexHealthAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HyperflexHealthAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HyperflexHealthAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HyperflexHealthAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HyperflexHealthAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HyperflexHealthAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HyperflexHealthAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetArbitrationServiceState

`func (o *HyperflexHealthAllOf) GetArbitrationServiceState() string`

GetArbitrationServiceState returns the ArbitrationServiceState field if non-nil, zero value otherwise.

### GetArbitrationServiceStateOk

`func (o *HyperflexHealthAllOf) GetArbitrationServiceStateOk() (*string, bool)`

GetArbitrationServiceStateOk returns a tuple with the ArbitrationServiceState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArbitrationServiceState

`func (o *HyperflexHealthAllOf) SetArbitrationServiceState(v string)`

SetArbitrationServiceState sets ArbitrationServiceState field to given value.

### HasArbitrationServiceState

`func (o *HyperflexHealthAllOf) HasArbitrationServiceState() bool`

HasArbitrationServiceState returns a boolean if a field has been set.

### GetDataReplicationCompliance

`func (o *HyperflexHealthAllOf) GetDataReplicationCompliance() string`

GetDataReplicationCompliance returns the DataReplicationCompliance field if non-nil, zero value otherwise.

### GetDataReplicationComplianceOk

`func (o *HyperflexHealthAllOf) GetDataReplicationComplianceOk() (*string, bool)`

GetDataReplicationComplianceOk returns a tuple with the DataReplicationCompliance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataReplicationCompliance

`func (o *HyperflexHealthAllOf) SetDataReplicationCompliance(v string)`

SetDataReplicationCompliance sets DataReplicationCompliance field to given value.

### HasDataReplicationCompliance

`func (o *HyperflexHealthAllOf) HasDataReplicationCompliance() bool`

HasDataReplicationCompliance returns a boolean if a field has been set.

### GetResiliencyDetails

`func (o *HyperflexHealthAllOf) GetResiliencyDetails() HyperflexHxResiliencyInfoDt`

GetResiliencyDetails returns the ResiliencyDetails field if non-nil, zero value otherwise.

### GetResiliencyDetailsOk

`func (o *HyperflexHealthAllOf) GetResiliencyDetailsOk() (*HyperflexHxResiliencyInfoDt, bool)`

GetResiliencyDetailsOk returns a tuple with the ResiliencyDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResiliencyDetails

`func (o *HyperflexHealthAllOf) SetResiliencyDetails(v HyperflexHxResiliencyInfoDt)`

SetResiliencyDetails sets ResiliencyDetails field to given value.

### HasResiliencyDetails

`func (o *HyperflexHealthAllOf) HasResiliencyDetails() bool`

HasResiliencyDetails returns a boolean if a field has been set.

### SetResiliencyDetailsNil

`func (o *HyperflexHealthAllOf) SetResiliencyDetailsNil(b bool)`

 SetResiliencyDetailsNil sets the value for ResiliencyDetails to be an explicit nil

### UnsetResiliencyDetails
`func (o *HyperflexHealthAllOf) UnsetResiliencyDetails()`

UnsetResiliencyDetails ensures that no value is present for ResiliencyDetails, not even an explicit nil
### GetState

`func (o *HyperflexHealthAllOf) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *HyperflexHealthAllOf) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *HyperflexHealthAllOf) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *HyperflexHealthAllOf) HasState() bool`

HasState returns a boolean if a field has been set.

### GetUuid

`func (o *HyperflexHealthAllOf) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *HyperflexHealthAllOf) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *HyperflexHealthAllOf) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *HyperflexHealthAllOf) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetZkHealth

`func (o *HyperflexHealthAllOf) GetZkHealth() string`

GetZkHealth returns the ZkHealth field if non-nil, zero value otherwise.

### GetZkHealthOk

`func (o *HyperflexHealthAllOf) GetZkHealthOk() (*string, bool)`

GetZkHealthOk returns a tuple with the ZkHealth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZkHealth

`func (o *HyperflexHealthAllOf) SetZkHealth(v string)`

SetZkHealth sets ZkHealth field to given value.

### HasZkHealth

`func (o *HyperflexHealthAllOf) HasZkHealth() bool`

HasZkHealth returns a boolean if a field has been set.

### GetZoneResiliencyList

`func (o *HyperflexHealthAllOf) GetZoneResiliencyList() []HyperflexHxZoneResiliencyInfoDt`

GetZoneResiliencyList returns the ZoneResiliencyList field if non-nil, zero value otherwise.

### GetZoneResiliencyListOk

`func (o *HyperflexHealthAllOf) GetZoneResiliencyListOk() (*[]HyperflexHxZoneResiliencyInfoDt, bool)`

GetZoneResiliencyListOk returns a tuple with the ZoneResiliencyList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneResiliencyList

`func (o *HyperflexHealthAllOf) SetZoneResiliencyList(v []HyperflexHxZoneResiliencyInfoDt)`

SetZoneResiliencyList sets ZoneResiliencyList field to given value.

### HasZoneResiliencyList

`func (o *HyperflexHealthAllOf) HasZoneResiliencyList() bool`

HasZoneResiliencyList returns a boolean if a field has been set.

### SetZoneResiliencyListNil

`func (o *HyperflexHealthAllOf) SetZoneResiliencyListNil(b bool)`

 SetZoneResiliencyListNil sets the value for ZoneResiliencyList to be an explicit nil

### UnsetZoneResiliencyList
`func (o *HyperflexHealthAllOf) UnsetZoneResiliencyList()`

UnsetZoneResiliencyList ensures that no value is present for ZoneResiliencyList, not even an explicit nil
### GetCluster

`func (o *HyperflexHealthAllOf) GetCluster() HyperflexClusterRelationship`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *HyperflexHealthAllOf) GetClusterOk() (*HyperflexClusterRelationship, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *HyperflexHealthAllOf) SetCluster(v HyperflexClusterRelationship)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *HyperflexHealthAllOf) HasCluster() bool`

HasCluster returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


