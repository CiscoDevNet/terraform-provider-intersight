# HyperflexHxResiliencyInfoDt

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hyperflex.HxResiliencyInfoDt"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hyperflex.HxResiliencyInfoDt"]
**DataReplicationFactor** | Pointer to **string** | The number of copies of data replicated by the cluster. * &#x60;ONE_COPY&#x60; - The HyperFlex cluster does not replicate data. * &#x60;TWO_COPIES&#x60; - The HyperFlex cluster keeps 2 copies of data. * &#x60;THREE_COPIES&#x60; - The HyperFlex cluster keeps 3 copies of data. * &#x60;FOUR_COPIES&#x60; - The HyperFlex cluster keeps 4 copies of data. * &#x60;SIX_COPIES&#x60; - The HyperFlex cluster keeps 6 copies of data. | [optional] [readonly] [default to "ONE_COPY"]
**HddFailuresTolerable** | Pointer to **int64** | The number of persistent device disruptions the HyperFlex storage cluster can handle at this point in time. | [optional] [readonly] 
**Messages** | Pointer to **[]string** |  | [optional] 
**NodeFailuresTolerable** | Pointer to **int64** | The number of node disruptions the HyperFlex storage cluster can handle at this point in time. | [optional] [readonly] 
**PolicyCompliance** | Pointer to **string** | The state of the storage cluster&#39;s compliance with the cluster access policy. * &#x60;UNKNOWN&#x60; - The HyperFlex cluster&#39;s compliance with the data replication policy could not be determined. * &#x60;COMPLIANT&#x60; - The HyperFlex cluster is compliant with the data replication policy and data is replicated to the configured replication factor. * &#x60;NON_COMPLIANT&#x60; - The HyperFlex cluster is not compliant with the data replication policy because some data is not replicated. | [optional] [readonly] [default to "UNKNOWN"]
**ResiliencyState** | Pointer to **string** | The resiliency state of the storage platform. The resiliency state is the storage cluster&#39;s current ability to maintain data. * &#x60;UNKNOWN&#x60; - The resiliency status of the HyperFlex cluster cannot be determined, or the cluster is transitioning into ONLINE state. * &#x60;HEALTHY&#x60; - The HyperFlex cluster is healthy. The cluster is able to perform IO operations and data is available. * &#x60;WARNING&#x60; - The HyperFlex cluster or data availability is adversely affected. This can happen if there are node or storage device failures beyond the tolerable failure threshold. * &#x60;OFFLINE&#x60; - The HyperFlex cluster is offline and not performing IO operations. * &#x60;CRITICAL&#x60; - The HyperFlex cluster has severe faults that affect cluster and data availability. | [optional] [readonly] [default to "UNKNOWN"]
**SsdFailuresTolerable** | Pointer to **int64** | The number of cache device disruptions the HyperFlex storage cluster can handle at this point in time. | [optional] [readonly] 

## Methods

### NewHyperflexHxResiliencyInfoDt

`func NewHyperflexHxResiliencyInfoDt(classId string, objectType string, ) *HyperflexHxResiliencyInfoDt`

NewHyperflexHxResiliencyInfoDt instantiates a new HyperflexHxResiliencyInfoDt object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexHxResiliencyInfoDtWithDefaults

`func NewHyperflexHxResiliencyInfoDtWithDefaults() *HyperflexHxResiliencyInfoDt`

NewHyperflexHxResiliencyInfoDtWithDefaults instantiates a new HyperflexHxResiliencyInfoDt object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HyperflexHxResiliencyInfoDt) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HyperflexHxResiliencyInfoDt) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HyperflexHxResiliencyInfoDt) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HyperflexHxResiliencyInfoDt) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HyperflexHxResiliencyInfoDt) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HyperflexHxResiliencyInfoDt) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDataReplicationFactor

`func (o *HyperflexHxResiliencyInfoDt) GetDataReplicationFactor() string`

GetDataReplicationFactor returns the DataReplicationFactor field if non-nil, zero value otherwise.

### GetDataReplicationFactorOk

`func (o *HyperflexHxResiliencyInfoDt) GetDataReplicationFactorOk() (*string, bool)`

GetDataReplicationFactorOk returns a tuple with the DataReplicationFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataReplicationFactor

`func (o *HyperflexHxResiliencyInfoDt) SetDataReplicationFactor(v string)`

SetDataReplicationFactor sets DataReplicationFactor field to given value.

### HasDataReplicationFactor

`func (o *HyperflexHxResiliencyInfoDt) HasDataReplicationFactor() bool`

HasDataReplicationFactor returns a boolean if a field has been set.

### GetHddFailuresTolerable

`func (o *HyperflexHxResiliencyInfoDt) GetHddFailuresTolerable() int64`

GetHddFailuresTolerable returns the HddFailuresTolerable field if non-nil, zero value otherwise.

### GetHddFailuresTolerableOk

`func (o *HyperflexHxResiliencyInfoDt) GetHddFailuresTolerableOk() (*int64, bool)`

GetHddFailuresTolerableOk returns a tuple with the HddFailuresTolerable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHddFailuresTolerable

`func (o *HyperflexHxResiliencyInfoDt) SetHddFailuresTolerable(v int64)`

SetHddFailuresTolerable sets HddFailuresTolerable field to given value.

### HasHddFailuresTolerable

`func (o *HyperflexHxResiliencyInfoDt) HasHddFailuresTolerable() bool`

HasHddFailuresTolerable returns a boolean if a field has been set.

### GetMessages

`func (o *HyperflexHxResiliencyInfoDt) GetMessages() []string`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *HyperflexHxResiliencyInfoDt) GetMessagesOk() (*[]string, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *HyperflexHxResiliencyInfoDt) SetMessages(v []string)`

SetMessages sets Messages field to given value.

### HasMessages

`func (o *HyperflexHxResiliencyInfoDt) HasMessages() bool`

HasMessages returns a boolean if a field has been set.

### SetMessagesNil

`func (o *HyperflexHxResiliencyInfoDt) SetMessagesNil(b bool)`

 SetMessagesNil sets the value for Messages to be an explicit nil

### UnsetMessages
`func (o *HyperflexHxResiliencyInfoDt) UnsetMessages()`

UnsetMessages ensures that no value is present for Messages, not even an explicit nil
### GetNodeFailuresTolerable

`func (o *HyperflexHxResiliencyInfoDt) GetNodeFailuresTolerable() int64`

GetNodeFailuresTolerable returns the NodeFailuresTolerable field if non-nil, zero value otherwise.

### GetNodeFailuresTolerableOk

`func (o *HyperflexHxResiliencyInfoDt) GetNodeFailuresTolerableOk() (*int64, bool)`

GetNodeFailuresTolerableOk returns a tuple with the NodeFailuresTolerable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeFailuresTolerable

`func (o *HyperflexHxResiliencyInfoDt) SetNodeFailuresTolerable(v int64)`

SetNodeFailuresTolerable sets NodeFailuresTolerable field to given value.

### HasNodeFailuresTolerable

`func (o *HyperflexHxResiliencyInfoDt) HasNodeFailuresTolerable() bool`

HasNodeFailuresTolerable returns a boolean if a field has been set.

### GetPolicyCompliance

`func (o *HyperflexHxResiliencyInfoDt) GetPolicyCompliance() string`

GetPolicyCompliance returns the PolicyCompliance field if non-nil, zero value otherwise.

### GetPolicyComplianceOk

`func (o *HyperflexHxResiliencyInfoDt) GetPolicyComplianceOk() (*string, bool)`

GetPolicyComplianceOk returns a tuple with the PolicyCompliance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyCompliance

`func (o *HyperflexHxResiliencyInfoDt) SetPolicyCompliance(v string)`

SetPolicyCompliance sets PolicyCompliance field to given value.

### HasPolicyCompliance

`func (o *HyperflexHxResiliencyInfoDt) HasPolicyCompliance() bool`

HasPolicyCompliance returns a boolean if a field has been set.

### GetResiliencyState

`func (o *HyperflexHxResiliencyInfoDt) GetResiliencyState() string`

GetResiliencyState returns the ResiliencyState field if non-nil, zero value otherwise.

### GetResiliencyStateOk

`func (o *HyperflexHxResiliencyInfoDt) GetResiliencyStateOk() (*string, bool)`

GetResiliencyStateOk returns a tuple with the ResiliencyState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResiliencyState

`func (o *HyperflexHxResiliencyInfoDt) SetResiliencyState(v string)`

SetResiliencyState sets ResiliencyState field to given value.

### HasResiliencyState

`func (o *HyperflexHxResiliencyInfoDt) HasResiliencyState() bool`

HasResiliencyState returns a boolean if a field has been set.

### GetSsdFailuresTolerable

`func (o *HyperflexHxResiliencyInfoDt) GetSsdFailuresTolerable() int64`

GetSsdFailuresTolerable returns the SsdFailuresTolerable field if non-nil, zero value otherwise.

### GetSsdFailuresTolerableOk

`func (o *HyperflexHxResiliencyInfoDt) GetSsdFailuresTolerableOk() (*int64, bool)`

GetSsdFailuresTolerableOk returns a tuple with the SsdFailuresTolerable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsdFailuresTolerable

`func (o *HyperflexHxResiliencyInfoDt) SetSsdFailuresTolerable(v int64)`

SetSsdFailuresTolerable sets SsdFailuresTolerable field to given value.

### HasSsdFailuresTolerable

`func (o *HyperflexHxResiliencyInfoDt) HasSsdFailuresTolerable() bool`

HasSsdFailuresTolerable returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


