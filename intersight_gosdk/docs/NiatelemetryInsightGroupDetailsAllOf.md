# NiatelemetryInsightGroupDetailsAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "niatelemetry.InsightGroupDetails"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "niatelemetry.InsightGroupDetails"]
**AlertRulesCount** | Pointer to **int64** | Alert rules count of the Insight group. | [optional] 
**AnalysisSettingsStatus** | Pointer to **string** | Analysis setting status of the Insight group. | [optional] 
**BugScanSettingsStatus** | Pointer to **string** | Bug scan settings status of the Insight group. | [optional] 
**DeltaAnalysisJobCount** | Pointer to **int64** | Delta analysis job count of the Insight group. | [optional] 
**EmailSettingsCount** | Pointer to **int64** | Email settings count of the Insight group. | [optional] 
**FlowSettingsCount** | Pointer to **int64** | Flow setting count of the Insight group. | [optional] 
**GroupName** | Pointer to **string** | Name of the Insight group. | [optional] 
**KafkaSettingsCount** | Pointer to **int64** | Kafka settings count of the Insight group. | [optional] 
**MicroBurstSettingsStatus** | Pointer to **string** | Microburst setting status of the Insight group. | [optional] 
**PrechangeAnalysisCount** | Pointer to **int64** | Prechange analysis count of the Insight group. | [optional] 
**TacCollectionConfigCount** | Pointer to **int64** | TAC collection config count of the Insight group. | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewNiatelemetryInsightGroupDetailsAllOf

`func NewNiatelemetryInsightGroupDetailsAllOf(classId string, objectType string, ) *NiatelemetryInsightGroupDetailsAllOf`

NewNiatelemetryInsightGroupDetailsAllOf instantiates a new NiatelemetryInsightGroupDetailsAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiatelemetryInsightGroupDetailsAllOfWithDefaults

`func NewNiatelemetryInsightGroupDetailsAllOfWithDefaults() *NiatelemetryInsightGroupDetailsAllOf`

NewNiatelemetryInsightGroupDetailsAllOfWithDefaults instantiates a new NiatelemetryInsightGroupDetailsAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NiatelemetryInsightGroupDetailsAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NiatelemetryInsightGroupDetailsAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NiatelemetryInsightGroupDetailsAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NiatelemetryInsightGroupDetailsAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NiatelemetryInsightGroupDetailsAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NiatelemetryInsightGroupDetailsAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAlertRulesCount

`func (o *NiatelemetryInsightGroupDetailsAllOf) GetAlertRulesCount() int64`

GetAlertRulesCount returns the AlertRulesCount field if non-nil, zero value otherwise.

### GetAlertRulesCountOk

`func (o *NiatelemetryInsightGroupDetailsAllOf) GetAlertRulesCountOk() (*int64, bool)`

GetAlertRulesCountOk returns a tuple with the AlertRulesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertRulesCount

`func (o *NiatelemetryInsightGroupDetailsAllOf) SetAlertRulesCount(v int64)`

SetAlertRulesCount sets AlertRulesCount field to given value.

### HasAlertRulesCount

`func (o *NiatelemetryInsightGroupDetailsAllOf) HasAlertRulesCount() bool`

HasAlertRulesCount returns a boolean if a field has been set.

### GetAnalysisSettingsStatus

`func (o *NiatelemetryInsightGroupDetailsAllOf) GetAnalysisSettingsStatus() string`

GetAnalysisSettingsStatus returns the AnalysisSettingsStatus field if non-nil, zero value otherwise.

### GetAnalysisSettingsStatusOk

`func (o *NiatelemetryInsightGroupDetailsAllOf) GetAnalysisSettingsStatusOk() (*string, bool)`

GetAnalysisSettingsStatusOk returns a tuple with the AnalysisSettingsStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalysisSettingsStatus

`func (o *NiatelemetryInsightGroupDetailsAllOf) SetAnalysisSettingsStatus(v string)`

SetAnalysisSettingsStatus sets AnalysisSettingsStatus field to given value.

### HasAnalysisSettingsStatus

`func (o *NiatelemetryInsightGroupDetailsAllOf) HasAnalysisSettingsStatus() bool`

HasAnalysisSettingsStatus returns a boolean if a field has been set.

### GetBugScanSettingsStatus

`func (o *NiatelemetryInsightGroupDetailsAllOf) GetBugScanSettingsStatus() string`

GetBugScanSettingsStatus returns the BugScanSettingsStatus field if non-nil, zero value otherwise.

### GetBugScanSettingsStatusOk

`func (o *NiatelemetryInsightGroupDetailsAllOf) GetBugScanSettingsStatusOk() (*string, bool)`

GetBugScanSettingsStatusOk returns a tuple with the BugScanSettingsStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBugScanSettingsStatus

`func (o *NiatelemetryInsightGroupDetailsAllOf) SetBugScanSettingsStatus(v string)`

SetBugScanSettingsStatus sets BugScanSettingsStatus field to given value.

### HasBugScanSettingsStatus

`func (o *NiatelemetryInsightGroupDetailsAllOf) HasBugScanSettingsStatus() bool`

HasBugScanSettingsStatus returns a boolean if a field has been set.

### GetDeltaAnalysisJobCount

`func (o *NiatelemetryInsightGroupDetailsAllOf) GetDeltaAnalysisJobCount() int64`

GetDeltaAnalysisJobCount returns the DeltaAnalysisJobCount field if non-nil, zero value otherwise.

### GetDeltaAnalysisJobCountOk

`func (o *NiatelemetryInsightGroupDetailsAllOf) GetDeltaAnalysisJobCountOk() (*int64, bool)`

GetDeltaAnalysisJobCountOk returns a tuple with the DeltaAnalysisJobCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeltaAnalysisJobCount

`func (o *NiatelemetryInsightGroupDetailsAllOf) SetDeltaAnalysisJobCount(v int64)`

SetDeltaAnalysisJobCount sets DeltaAnalysisJobCount field to given value.

### HasDeltaAnalysisJobCount

`func (o *NiatelemetryInsightGroupDetailsAllOf) HasDeltaAnalysisJobCount() bool`

HasDeltaAnalysisJobCount returns a boolean if a field has been set.

### GetEmailSettingsCount

`func (o *NiatelemetryInsightGroupDetailsAllOf) GetEmailSettingsCount() int64`

GetEmailSettingsCount returns the EmailSettingsCount field if non-nil, zero value otherwise.

### GetEmailSettingsCountOk

`func (o *NiatelemetryInsightGroupDetailsAllOf) GetEmailSettingsCountOk() (*int64, bool)`

GetEmailSettingsCountOk returns a tuple with the EmailSettingsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailSettingsCount

`func (o *NiatelemetryInsightGroupDetailsAllOf) SetEmailSettingsCount(v int64)`

SetEmailSettingsCount sets EmailSettingsCount field to given value.

### HasEmailSettingsCount

`func (o *NiatelemetryInsightGroupDetailsAllOf) HasEmailSettingsCount() bool`

HasEmailSettingsCount returns a boolean if a field has been set.

### GetFlowSettingsCount

`func (o *NiatelemetryInsightGroupDetailsAllOf) GetFlowSettingsCount() int64`

GetFlowSettingsCount returns the FlowSettingsCount field if non-nil, zero value otherwise.

### GetFlowSettingsCountOk

`func (o *NiatelemetryInsightGroupDetailsAllOf) GetFlowSettingsCountOk() (*int64, bool)`

GetFlowSettingsCountOk returns a tuple with the FlowSettingsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowSettingsCount

`func (o *NiatelemetryInsightGroupDetailsAllOf) SetFlowSettingsCount(v int64)`

SetFlowSettingsCount sets FlowSettingsCount field to given value.

### HasFlowSettingsCount

`func (o *NiatelemetryInsightGroupDetailsAllOf) HasFlowSettingsCount() bool`

HasFlowSettingsCount returns a boolean if a field has been set.

### GetGroupName

`func (o *NiatelemetryInsightGroupDetailsAllOf) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *NiatelemetryInsightGroupDetailsAllOf) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *NiatelemetryInsightGroupDetailsAllOf) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.

### HasGroupName

`func (o *NiatelemetryInsightGroupDetailsAllOf) HasGroupName() bool`

HasGroupName returns a boolean if a field has been set.

### GetKafkaSettingsCount

`func (o *NiatelemetryInsightGroupDetailsAllOf) GetKafkaSettingsCount() int64`

GetKafkaSettingsCount returns the KafkaSettingsCount field if non-nil, zero value otherwise.

### GetKafkaSettingsCountOk

`func (o *NiatelemetryInsightGroupDetailsAllOf) GetKafkaSettingsCountOk() (*int64, bool)`

GetKafkaSettingsCountOk returns a tuple with the KafkaSettingsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKafkaSettingsCount

`func (o *NiatelemetryInsightGroupDetailsAllOf) SetKafkaSettingsCount(v int64)`

SetKafkaSettingsCount sets KafkaSettingsCount field to given value.

### HasKafkaSettingsCount

`func (o *NiatelemetryInsightGroupDetailsAllOf) HasKafkaSettingsCount() bool`

HasKafkaSettingsCount returns a boolean if a field has been set.

### GetMicroBurstSettingsStatus

`func (o *NiatelemetryInsightGroupDetailsAllOf) GetMicroBurstSettingsStatus() string`

GetMicroBurstSettingsStatus returns the MicroBurstSettingsStatus field if non-nil, zero value otherwise.

### GetMicroBurstSettingsStatusOk

`func (o *NiatelemetryInsightGroupDetailsAllOf) GetMicroBurstSettingsStatusOk() (*string, bool)`

GetMicroBurstSettingsStatusOk returns a tuple with the MicroBurstSettingsStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMicroBurstSettingsStatus

`func (o *NiatelemetryInsightGroupDetailsAllOf) SetMicroBurstSettingsStatus(v string)`

SetMicroBurstSettingsStatus sets MicroBurstSettingsStatus field to given value.

### HasMicroBurstSettingsStatus

`func (o *NiatelemetryInsightGroupDetailsAllOf) HasMicroBurstSettingsStatus() bool`

HasMicroBurstSettingsStatus returns a boolean if a field has been set.

### GetPrechangeAnalysisCount

`func (o *NiatelemetryInsightGroupDetailsAllOf) GetPrechangeAnalysisCount() int64`

GetPrechangeAnalysisCount returns the PrechangeAnalysisCount field if non-nil, zero value otherwise.

### GetPrechangeAnalysisCountOk

`func (o *NiatelemetryInsightGroupDetailsAllOf) GetPrechangeAnalysisCountOk() (*int64, bool)`

GetPrechangeAnalysisCountOk returns a tuple with the PrechangeAnalysisCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrechangeAnalysisCount

`func (o *NiatelemetryInsightGroupDetailsAllOf) SetPrechangeAnalysisCount(v int64)`

SetPrechangeAnalysisCount sets PrechangeAnalysisCount field to given value.

### HasPrechangeAnalysisCount

`func (o *NiatelemetryInsightGroupDetailsAllOf) HasPrechangeAnalysisCount() bool`

HasPrechangeAnalysisCount returns a boolean if a field has been set.

### GetTacCollectionConfigCount

`func (o *NiatelemetryInsightGroupDetailsAllOf) GetTacCollectionConfigCount() int64`

GetTacCollectionConfigCount returns the TacCollectionConfigCount field if non-nil, zero value otherwise.

### GetTacCollectionConfigCountOk

`func (o *NiatelemetryInsightGroupDetailsAllOf) GetTacCollectionConfigCountOk() (*int64, bool)`

GetTacCollectionConfigCountOk returns a tuple with the TacCollectionConfigCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTacCollectionConfigCount

`func (o *NiatelemetryInsightGroupDetailsAllOf) SetTacCollectionConfigCount(v int64)`

SetTacCollectionConfigCount sets TacCollectionConfigCount field to given value.

### HasTacCollectionConfigCount

`func (o *NiatelemetryInsightGroupDetailsAllOf) HasTacCollectionConfigCount() bool`

HasTacCollectionConfigCount returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *NiatelemetryInsightGroupDetailsAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *NiatelemetryInsightGroupDetailsAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *NiatelemetryInsightGroupDetailsAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *NiatelemetryInsightGroupDetailsAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


