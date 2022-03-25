# NiatelemetryInsightGroupDetails

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
**InsightSites** | Pointer to [**[]NiatelemetrySites**](NiatelemetrySites.md) |  | [optional] 
**KafkaSettingsCount** | Pointer to **int64** | Kafka settings count of the Insight group. | [optional] 
**MicroBurstSettingsStatus** | Pointer to **string** | Microburst setting status of the Insight group. | [optional] 
**PrechangeAnalysisCount** | Pointer to **int64** | Prechange analysis count of the Insight group. | [optional] 
**TacCollectionConfigCount** | Pointer to **int64** | TAC collection config count of the Insight group. | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewNiatelemetryInsightGroupDetails

`func NewNiatelemetryInsightGroupDetails(classId string, objectType string, ) *NiatelemetryInsightGroupDetails`

NewNiatelemetryInsightGroupDetails instantiates a new NiatelemetryInsightGroupDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiatelemetryInsightGroupDetailsWithDefaults

`func NewNiatelemetryInsightGroupDetailsWithDefaults() *NiatelemetryInsightGroupDetails`

NewNiatelemetryInsightGroupDetailsWithDefaults instantiates a new NiatelemetryInsightGroupDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NiatelemetryInsightGroupDetails) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NiatelemetryInsightGroupDetails) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NiatelemetryInsightGroupDetails) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NiatelemetryInsightGroupDetails) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NiatelemetryInsightGroupDetails) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NiatelemetryInsightGroupDetails) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAlertRulesCount

`func (o *NiatelemetryInsightGroupDetails) GetAlertRulesCount() int64`

GetAlertRulesCount returns the AlertRulesCount field if non-nil, zero value otherwise.

### GetAlertRulesCountOk

`func (o *NiatelemetryInsightGroupDetails) GetAlertRulesCountOk() (*int64, bool)`

GetAlertRulesCountOk returns a tuple with the AlertRulesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertRulesCount

`func (o *NiatelemetryInsightGroupDetails) SetAlertRulesCount(v int64)`

SetAlertRulesCount sets AlertRulesCount field to given value.

### HasAlertRulesCount

`func (o *NiatelemetryInsightGroupDetails) HasAlertRulesCount() bool`

HasAlertRulesCount returns a boolean if a field has been set.

### GetAnalysisSettingsStatus

`func (o *NiatelemetryInsightGroupDetails) GetAnalysisSettingsStatus() string`

GetAnalysisSettingsStatus returns the AnalysisSettingsStatus field if non-nil, zero value otherwise.

### GetAnalysisSettingsStatusOk

`func (o *NiatelemetryInsightGroupDetails) GetAnalysisSettingsStatusOk() (*string, bool)`

GetAnalysisSettingsStatusOk returns a tuple with the AnalysisSettingsStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalysisSettingsStatus

`func (o *NiatelemetryInsightGroupDetails) SetAnalysisSettingsStatus(v string)`

SetAnalysisSettingsStatus sets AnalysisSettingsStatus field to given value.

### HasAnalysisSettingsStatus

`func (o *NiatelemetryInsightGroupDetails) HasAnalysisSettingsStatus() bool`

HasAnalysisSettingsStatus returns a boolean if a field has been set.

### GetBugScanSettingsStatus

`func (o *NiatelemetryInsightGroupDetails) GetBugScanSettingsStatus() string`

GetBugScanSettingsStatus returns the BugScanSettingsStatus field if non-nil, zero value otherwise.

### GetBugScanSettingsStatusOk

`func (o *NiatelemetryInsightGroupDetails) GetBugScanSettingsStatusOk() (*string, bool)`

GetBugScanSettingsStatusOk returns a tuple with the BugScanSettingsStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBugScanSettingsStatus

`func (o *NiatelemetryInsightGroupDetails) SetBugScanSettingsStatus(v string)`

SetBugScanSettingsStatus sets BugScanSettingsStatus field to given value.

### HasBugScanSettingsStatus

`func (o *NiatelemetryInsightGroupDetails) HasBugScanSettingsStatus() bool`

HasBugScanSettingsStatus returns a boolean if a field has been set.

### GetDeltaAnalysisJobCount

`func (o *NiatelemetryInsightGroupDetails) GetDeltaAnalysisJobCount() int64`

GetDeltaAnalysisJobCount returns the DeltaAnalysisJobCount field if non-nil, zero value otherwise.

### GetDeltaAnalysisJobCountOk

`func (o *NiatelemetryInsightGroupDetails) GetDeltaAnalysisJobCountOk() (*int64, bool)`

GetDeltaAnalysisJobCountOk returns a tuple with the DeltaAnalysisJobCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeltaAnalysisJobCount

`func (o *NiatelemetryInsightGroupDetails) SetDeltaAnalysisJobCount(v int64)`

SetDeltaAnalysisJobCount sets DeltaAnalysisJobCount field to given value.

### HasDeltaAnalysisJobCount

`func (o *NiatelemetryInsightGroupDetails) HasDeltaAnalysisJobCount() bool`

HasDeltaAnalysisJobCount returns a boolean if a field has been set.

### GetEmailSettingsCount

`func (o *NiatelemetryInsightGroupDetails) GetEmailSettingsCount() int64`

GetEmailSettingsCount returns the EmailSettingsCount field if non-nil, zero value otherwise.

### GetEmailSettingsCountOk

`func (o *NiatelemetryInsightGroupDetails) GetEmailSettingsCountOk() (*int64, bool)`

GetEmailSettingsCountOk returns a tuple with the EmailSettingsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailSettingsCount

`func (o *NiatelemetryInsightGroupDetails) SetEmailSettingsCount(v int64)`

SetEmailSettingsCount sets EmailSettingsCount field to given value.

### HasEmailSettingsCount

`func (o *NiatelemetryInsightGroupDetails) HasEmailSettingsCount() bool`

HasEmailSettingsCount returns a boolean if a field has been set.

### GetFlowSettingsCount

`func (o *NiatelemetryInsightGroupDetails) GetFlowSettingsCount() int64`

GetFlowSettingsCount returns the FlowSettingsCount field if non-nil, zero value otherwise.

### GetFlowSettingsCountOk

`func (o *NiatelemetryInsightGroupDetails) GetFlowSettingsCountOk() (*int64, bool)`

GetFlowSettingsCountOk returns a tuple with the FlowSettingsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowSettingsCount

`func (o *NiatelemetryInsightGroupDetails) SetFlowSettingsCount(v int64)`

SetFlowSettingsCount sets FlowSettingsCount field to given value.

### HasFlowSettingsCount

`func (o *NiatelemetryInsightGroupDetails) HasFlowSettingsCount() bool`

HasFlowSettingsCount returns a boolean if a field has been set.

### GetGroupName

`func (o *NiatelemetryInsightGroupDetails) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *NiatelemetryInsightGroupDetails) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *NiatelemetryInsightGroupDetails) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.

### HasGroupName

`func (o *NiatelemetryInsightGroupDetails) HasGroupName() bool`

HasGroupName returns a boolean if a field has been set.

### GetInsightSites

`func (o *NiatelemetryInsightGroupDetails) GetInsightSites() []NiatelemetrySites`

GetInsightSites returns the InsightSites field if non-nil, zero value otherwise.

### GetInsightSitesOk

`func (o *NiatelemetryInsightGroupDetails) GetInsightSitesOk() (*[]NiatelemetrySites, bool)`

GetInsightSitesOk returns a tuple with the InsightSites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsightSites

`func (o *NiatelemetryInsightGroupDetails) SetInsightSites(v []NiatelemetrySites)`

SetInsightSites sets InsightSites field to given value.

### HasInsightSites

`func (o *NiatelemetryInsightGroupDetails) HasInsightSites() bool`

HasInsightSites returns a boolean if a field has been set.

### SetInsightSitesNil

`func (o *NiatelemetryInsightGroupDetails) SetInsightSitesNil(b bool)`

 SetInsightSitesNil sets the value for InsightSites to be an explicit nil

### UnsetInsightSites
`func (o *NiatelemetryInsightGroupDetails) UnsetInsightSites()`

UnsetInsightSites ensures that no value is present for InsightSites, not even an explicit nil
### GetKafkaSettingsCount

`func (o *NiatelemetryInsightGroupDetails) GetKafkaSettingsCount() int64`

GetKafkaSettingsCount returns the KafkaSettingsCount field if non-nil, zero value otherwise.

### GetKafkaSettingsCountOk

`func (o *NiatelemetryInsightGroupDetails) GetKafkaSettingsCountOk() (*int64, bool)`

GetKafkaSettingsCountOk returns a tuple with the KafkaSettingsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKafkaSettingsCount

`func (o *NiatelemetryInsightGroupDetails) SetKafkaSettingsCount(v int64)`

SetKafkaSettingsCount sets KafkaSettingsCount field to given value.

### HasKafkaSettingsCount

`func (o *NiatelemetryInsightGroupDetails) HasKafkaSettingsCount() bool`

HasKafkaSettingsCount returns a boolean if a field has been set.

### GetMicroBurstSettingsStatus

`func (o *NiatelemetryInsightGroupDetails) GetMicroBurstSettingsStatus() string`

GetMicroBurstSettingsStatus returns the MicroBurstSettingsStatus field if non-nil, zero value otherwise.

### GetMicroBurstSettingsStatusOk

`func (o *NiatelemetryInsightGroupDetails) GetMicroBurstSettingsStatusOk() (*string, bool)`

GetMicroBurstSettingsStatusOk returns a tuple with the MicroBurstSettingsStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMicroBurstSettingsStatus

`func (o *NiatelemetryInsightGroupDetails) SetMicroBurstSettingsStatus(v string)`

SetMicroBurstSettingsStatus sets MicroBurstSettingsStatus field to given value.

### HasMicroBurstSettingsStatus

`func (o *NiatelemetryInsightGroupDetails) HasMicroBurstSettingsStatus() bool`

HasMicroBurstSettingsStatus returns a boolean if a field has been set.

### GetPrechangeAnalysisCount

`func (o *NiatelemetryInsightGroupDetails) GetPrechangeAnalysisCount() int64`

GetPrechangeAnalysisCount returns the PrechangeAnalysisCount field if non-nil, zero value otherwise.

### GetPrechangeAnalysisCountOk

`func (o *NiatelemetryInsightGroupDetails) GetPrechangeAnalysisCountOk() (*int64, bool)`

GetPrechangeAnalysisCountOk returns a tuple with the PrechangeAnalysisCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrechangeAnalysisCount

`func (o *NiatelemetryInsightGroupDetails) SetPrechangeAnalysisCount(v int64)`

SetPrechangeAnalysisCount sets PrechangeAnalysisCount field to given value.

### HasPrechangeAnalysisCount

`func (o *NiatelemetryInsightGroupDetails) HasPrechangeAnalysisCount() bool`

HasPrechangeAnalysisCount returns a boolean if a field has been set.

### GetTacCollectionConfigCount

`func (o *NiatelemetryInsightGroupDetails) GetTacCollectionConfigCount() int64`

GetTacCollectionConfigCount returns the TacCollectionConfigCount field if non-nil, zero value otherwise.

### GetTacCollectionConfigCountOk

`func (o *NiatelemetryInsightGroupDetails) GetTacCollectionConfigCountOk() (*int64, bool)`

GetTacCollectionConfigCountOk returns a tuple with the TacCollectionConfigCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTacCollectionConfigCount

`func (o *NiatelemetryInsightGroupDetails) SetTacCollectionConfigCount(v int64)`

SetTacCollectionConfigCount sets TacCollectionConfigCount field to given value.

### HasTacCollectionConfigCount

`func (o *NiatelemetryInsightGroupDetails) HasTacCollectionConfigCount() bool`

HasTacCollectionConfigCount returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *NiatelemetryInsightGroupDetails) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *NiatelemetryInsightGroupDetails) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *NiatelemetryInsightGroupDetails) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *NiatelemetryInsightGroupDetails) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


