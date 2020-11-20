# HyperflexHxResiliencyInfoDtAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hyperflex.HxResiliencyInfoDt"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hyperflex.HxResiliencyInfoDt"]
**DataReplicationFactor** | Pointer to **string** |  | [optional] [readonly] [default to "ONE_COPY"]
**HddFailuresTolerable** | Pointer to **int64** |  | [optional] [readonly] 
**Messages** | Pointer to **[]string** |  | [optional] 
**NodeFailuresTolerable** | Pointer to **int64** |  | [optional] [readonly] 
**PolicyCompliance** | Pointer to **string** |  | [optional] [readonly] [default to "UNKNOWN"]
**ResiliencyState** | Pointer to **string** |  | [optional] [readonly] [default to "UNKNOWN"]
**SsdFailuresTolerable** | Pointer to **int64** |  | [optional] [readonly] 

## Methods

### NewHyperflexHxResiliencyInfoDtAllOf

`func NewHyperflexHxResiliencyInfoDtAllOf(classId string, objectType string, ) *HyperflexHxResiliencyInfoDtAllOf`

NewHyperflexHxResiliencyInfoDtAllOf instantiates a new HyperflexHxResiliencyInfoDtAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexHxResiliencyInfoDtAllOfWithDefaults

`func NewHyperflexHxResiliencyInfoDtAllOfWithDefaults() *HyperflexHxResiliencyInfoDtAllOf`

NewHyperflexHxResiliencyInfoDtAllOfWithDefaults instantiates a new HyperflexHxResiliencyInfoDtAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HyperflexHxResiliencyInfoDtAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HyperflexHxResiliencyInfoDtAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HyperflexHxResiliencyInfoDtAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HyperflexHxResiliencyInfoDtAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HyperflexHxResiliencyInfoDtAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HyperflexHxResiliencyInfoDtAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDataReplicationFactor

`func (o *HyperflexHxResiliencyInfoDtAllOf) GetDataReplicationFactor() string`

GetDataReplicationFactor returns the DataReplicationFactor field if non-nil, zero value otherwise.

### GetDataReplicationFactorOk

`func (o *HyperflexHxResiliencyInfoDtAllOf) GetDataReplicationFactorOk() (*string, bool)`

GetDataReplicationFactorOk returns a tuple with the DataReplicationFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataReplicationFactor

`func (o *HyperflexHxResiliencyInfoDtAllOf) SetDataReplicationFactor(v string)`

SetDataReplicationFactor sets DataReplicationFactor field to given value.

### HasDataReplicationFactor

`func (o *HyperflexHxResiliencyInfoDtAllOf) HasDataReplicationFactor() bool`

HasDataReplicationFactor returns a boolean if a field has been set.

### GetHddFailuresTolerable

`func (o *HyperflexHxResiliencyInfoDtAllOf) GetHddFailuresTolerable() int64`

GetHddFailuresTolerable returns the HddFailuresTolerable field if non-nil, zero value otherwise.

### GetHddFailuresTolerableOk

`func (o *HyperflexHxResiliencyInfoDtAllOf) GetHddFailuresTolerableOk() (*int64, bool)`

GetHddFailuresTolerableOk returns a tuple with the HddFailuresTolerable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHddFailuresTolerable

`func (o *HyperflexHxResiliencyInfoDtAllOf) SetHddFailuresTolerable(v int64)`

SetHddFailuresTolerable sets HddFailuresTolerable field to given value.

### HasHddFailuresTolerable

`func (o *HyperflexHxResiliencyInfoDtAllOf) HasHddFailuresTolerable() bool`

HasHddFailuresTolerable returns a boolean if a field has been set.

### GetMessages

`func (o *HyperflexHxResiliencyInfoDtAllOf) GetMessages() []string`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *HyperflexHxResiliencyInfoDtAllOf) GetMessagesOk() (*[]string, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *HyperflexHxResiliencyInfoDtAllOf) SetMessages(v []string)`

SetMessages sets Messages field to given value.

### HasMessages

`func (o *HyperflexHxResiliencyInfoDtAllOf) HasMessages() bool`

HasMessages returns a boolean if a field has been set.

### SetMessagesNil

`func (o *HyperflexHxResiliencyInfoDtAllOf) SetMessagesNil(b bool)`

 SetMessagesNil sets the value for Messages to be an explicit nil

### UnsetMessages
`func (o *HyperflexHxResiliencyInfoDtAllOf) UnsetMessages()`

UnsetMessages ensures that no value is present for Messages, not even an explicit nil
### GetNodeFailuresTolerable

`func (o *HyperflexHxResiliencyInfoDtAllOf) GetNodeFailuresTolerable() int64`

GetNodeFailuresTolerable returns the NodeFailuresTolerable field if non-nil, zero value otherwise.

### GetNodeFailuresTolerableOk

`func (o *HyperflexHxResiliencyInfoDtAllOf) GetNodeFailuresTolerableOk() (*int64, bool)`

GetNodeFailuresTolerableOk returns a tuple with the NodeFailuresTolerable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeFailuresTolerable

`func (o *HyperflexHxResiliencyInfoDtAllOf) SetNodeFailuresTolerable(v int64)`

SetNodeFailuresTolerable sets NodeFailuresTolerable field to given value.

### HasNodeFailuresTolerable

`func (o *HyperflexHxResiliencyInfoDtAllOf) HasNodeFailuresTolerable() bool`

HasNodeFailuresTolerable returns a boolean if a field has been set.

### GetPolicyCompliance

`func (o *HyperflexHxResiliencyInfoDtAllOf) GetPolicyCompliance() string`

GetPolicyCompliance returns the PolicyCompliance field if non-nil, zero value otherwise.

### GetPolicyComplianceOk

`func (o *HyperflexHxResiliencyInfoDtAllOf) GetPolicyComplianceOk() (*string, bool)`

GetPolicyComplianceOk returns a tuple with the PolicyCompliance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyCompliance

`func (o *HyperflexHxResiliencyInfoDtAllOf) SetPolicyCompliance(v string)`

SetPolicyCompliance sets PolicyCompliance field to given value.

### HasPolicyCompliance

`func (o *HyperflexHxResiliencyInfoDtAllOf) HasPolicyCompliance() bool`

HasPolicyCompliance returns a boolean if a field has been set.

### GetResiliencyState

`func (o *HyperflexHxResiliencyInfoDtAllOf) GetResiliencyState() string`

GetResiliencyState returns the ResiliencyState field if non-nil, zero value otherwise.

### GetResiliencyStateOk

`func (o *HyperflexHxResiliencyInfoDtAllOf) GetResiliencyStateOk() (*string, bool)`

GetResiliencyStateOk returns a tuple with the ResiliencyState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResiliencyState

`func (o *HyperflexHxResiliencyInfoDtAllOf) SetResiliencyState(v string)`

SetResiliencyState sets ResiliencyState field to given value.

### HasResiliencyState

`func (o *HyperflexHxResiliencyInfoDtAllOf) HasResiliencyState() bool`

HasResiliencyState returns a boolean if a field has been set.

### GetSsdFailuresTolerable

`func (o *HyperflexHxResiliencyInfoDtAllOf) GetSsdFailuresTolerable() int64`

GetSsdFailuresTolerable returns the SsdFailuresTolerable field if non-nil, zero value otherwise.

### GetSsdFailuresTolerableOk

`func (o *HyperflexHxResiliencyInfoDtAllOf) GetSsdFailuresTolerableOk() (*int64, bool)`

GetSsdFailuresTolerableOk returns a tuple with the SsdFailuresTolerable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsdFailuresTolerable

`func (o *HyperflexHxResiliencyInfoDtAllOf) SetSsdFailuresTolerable(v int64)`

SetSsdFailuresTolerable sets SsdFailuresTolerable field to given value.

### HasSsdFailuresTolerable

`func (o *HyperflexHxResiliencyInfoDtAllOf) HasSsdFailuresTolerable() bool`

HasSsdFailuresTolerable returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


