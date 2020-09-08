# HyperflexHxResiliencyInfoDt

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataReplicationFactor** | Pointer to **string** |  | [optional] [readonly] [default to "ONE_COPY"]
**HddFailuresTolerable** | Pointer to **int64** |  | [optional] [readonly] 
**Messages** | Pointer to **[]string** |  | [optional] 
**NodeFailuresTolerable** | Pointer to **int64** |  | [optional] [readonly] 
**PolicyCompliance** | Pointer to **string** |  | [optional] [readonly] [default to "UNKNOWN"]
**ResiliencyState** | Pointer to **string** |  | [optional] [readonly] [default to "UNKNOWN"]
**SsdFailuresTolerable** | Pointer to **int64** |  | [optional] [readonly] 

## Methods

### NewHyperflexHxResiliencyInfoDt

`func NewHyperflexHxResiliencyInfoDt() *HyperflexHxResiliencyInfoDt`

NewHyperflexHxResiliencyInfoDt instantiates a new HyperflexHxResiliencyInfoDt object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexHxResiliencyInfoDtWithDefaults

`func NewHyperflexHxResiliencyInfoDtWithDefaults() *HyperflexHxResiliencyInfoDt`

NewHyperflexHxResiliencyInfoDtWithDefaults instantiates a new HyperflexHxResiliencyInfoDt object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

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


