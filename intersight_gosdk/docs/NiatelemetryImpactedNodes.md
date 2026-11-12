# NiatelemetryImpactedNodes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "niatelemetry.ImpactedNodes"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "niatelemetry.ImpactedNodes"]
**CompensatingControlFailureReason** | Pointer to **string** | Failure reason reported for the impacted node compensating control. | [optional] [readonly] 
**CompensatingControlLastDeployedMode** | Pointer to **string** | Last deployed compensating control mode for the impacted node. | [optional] [readonly] 
**CompensatingControlStatus** | Pointer to **string** | Compensating control status reported for the impacted node. | [optional] [readonly] 
**HitCount** | Pointer to **int64** | Count of advisory hits reported for the impacted node. | [optional] [readonly] 
**LiveProtectAgentStatus** | Pointer to **string** | Live Protect agent status reported for the impacted node. | [optional] [readonly] 
**LiveProtectStatus** | Pointer to **string** | Live Protect status reported for the impacted node. | [optional] [readonly] 
**SwitchId** | Pointer to **string** | Switch identifier reported for the impacted node. | [optional] [readonly] 
**SwitchName** | Pointer to **string** | Switch name reported for the impacted node. | [optional] [readonly] 

## Methods

### NewNiatelemetryImpactedNodes

`func NewNiatelemetryImpactedNodes(classId string, objectType string, ) *NiatelemetryImpactedNodes`

NewNiatelemetryImpactedNodes instantiates a new NiatelemetryImpactedNodes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiatelemetryImpactedNodesWithDefaults

`func NewNiatelemetryImpactedNodesWithDefaults() *NiatelemetryImpactedNodes`

NewNiatelemetryImpactedNodesWithDefaults instantiates a new NiatelemetryImpactedNodes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NiatelemetryImpactedNodes) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NiatelemetryImpactedNodes) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NiatelemetryImpactedNodes) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NiatelemetryImpactedNodes) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NiatelemetryImpactedNodes) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NiatelemetryImpactedNodes) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCompensatingControlFailureReason

`func (o *NiatelemetryImpactedNodes) GetCompensatingControlFailureReason() string`

GetCompensatingControlFailureReason returns the CompensatingControlFailureReason field if non-nil, zero value otherwise.

### GetCompensatingControlFailureReasonOk

`func (o *NiatelemetryImpactedNodes) GetCompensatingControlFailureReasonOk() (*string, bool)`

GetCompensatingControlFailureReasonOk returns a tuple with the CompensatingControlFailureReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompensatingControlFailureReason

`func (o *NiatelemetryImpactedNodes) SetCompensatingControlFailureReason(v string)`

SetCompensatingControlFailureReason sets CompensatingControlFailureReason field to given value.

### HasCompensatingControlFailureReason

`func (o *NiatelemetryImpactedNodes) HasCompensatingControlFailureReason() bool`

HasCompensatingControlFailureReason returns a boolean if a field has been set.

### GetCompensatingControlLastDeployedMode

`func (o *NiatelemetryImpactedNodes) GetCompensatingControlLastDeployedMode() string`

GetCompensatingControlLastDeployedMode returns the CompensatingControlLastDeployedMode field if non-nil, zero value otherwise.

### GetCompensatingControlLastDeployedModeOk

`func (o *NiatelemetryImpactedNodes) GetCompensatingControlLastDeployedModeOk() (*string, bool)`

GetCompensatingControlLastDeployedModeOk returns a tuple with the CompensatingControlLastDeployedMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompensatingControlLastDeployedMode

`func (o *NiatelemetryImpactedNodes) SetCompensatingControlLastDeployedMode(v string)`

SetCompensatingControlLastDeployedMode sets CompensatingControlLastDeployedMode field to given value.

### HasCompensatingControlLastDeployedMode

`func (o *NiatelemetryImpactedNodes) HasCompensatingControlLastDeployedMode() bool`

HasCompensatingControlLastDeployedMode returns a boolean if a field has been set.

### GetCompensatingControlStatus

`func (o *NiatelemetryImpactedNodes) GetCompensatingControlStatus() string`

GetCompensatingControlStatus returns the CompensatingControlStatus field if non-nil, zero value otherwise.

### GetCompensatingControlStatusOk

`func (o *NiatelemetryImpactedNodes) GetCompensatingControlStatusOk() (*string, bool)`

GetCompensatingControlStatusOk returns a tuple with the CompensatingControlStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompensatingControlStatus

`func (o *NiatelemetryImpactedNodes) SetCompensatingControlStatus(v string)`

SetCompensatingControlStatus sets CompensatingControlStatus field to given value.

### HasCompensatingControlStatus

`func (o *NiatelemetryImpactedNodes) HasCompensatingControlStatus() bool`

HasCompensatingControlStatus returns a boolean if a field has been set.

### GetHitCount

`func (o *NiatelemetryImpactedNodes) GetHitCount() int64`

GetHitCount returns the HitCount field if non-nil, zero value otherwise.

### GetHitCountOk

`func (o *NiatelemetryImpactedNodes) GetHitCountOk() (*int64, bool)`

GetHitCountOk returns a tuple with the HitCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHitCount

`func (o *NiatelemetryImpactedNodes) SetHitCount(v int64)`

SetHitCount sets HitCount field to given value.

### HasHitCount

`func (o *NiatelemetryImpactedNodes) HasHitCount() bool`

HasHitCount returns a boolean if a field has been set.

### GetLiveProtectAgentStatus

`func (o *NiatelemetryImpactedNodes) GetLiveProtectAgentStatus() string`

GetLiveProtectAgentStatus returns the LiveProtectAgentStatus field if non-nil, zero value otherwise.

### GetLiveProtectAgentStatusOk

`func (o *NiatelemetryImpactedNodes) GetLiveProtectAgentStatusOk() (*string, bool)`

GetLiveProtectAgentStatusOk returns a tuple with the LiveProtectAgentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiveProtectAgentStatus

`func (o *NiatelemetryImpactedNodes) SetLiveProtectAgentStatus(v string)`

SetLiveProtectAgentStatus sets LiveProtectAgentStatus field to given value.

### HasLiveProtectAgentStatus

`func (o *NiatelemetryImpactedNodes) HasLiveProtectAgentStatus() bool`

HasLiveProtectAgentStatus returns a boolean if a field has been set.

### GetLiveProtectStatus

`func (o *NiatelemetryImpactedNodes) GetLiveProtectStatus() string`

GetLiveProtectStatus returns the LiveProtectStatus field if non-nil, zero value otherwise.

### GetLiveProtectStatusOk

`func (o *NiatelemetryImpactedNodes) GetLiveProtectStatusOk() (*string, bool)`

GetLiveProtectStatusOk returns a tuple with the LiveProtectStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiveProtectStatus

`func (o *NiatelemetryImpactedNodes) SetLiveProtectStatus(v string)`

SetLiveProtectStatus sets LiveProtectStatus field to given value.

### HasLiveProtectStatus

`func (o *NiatelemetryImpactedNodes) HasLiveProtectStatus() bool`

HasLiveProtectStatus returns a boolean if a field has been set.

### GetSwitchId

`func (o *NiatelemetryImpactedNodes) GetSwitchId() string`

GetSwitchId returns the SwitchId field if non-nil, zero value otherwise.

### GetSwitchIdOk

`func (o *NiatelemetryImpactedNodes) GetSwitchIdOk() (*string, bool)`

GetSwitchIdOk returns a tuple with the SwitchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchId

`func (o *NiatelemetryImpactedNodes) SetSwitchId(v string)`

SetSwitchId sets SwitchId field to given value.

### HasSwitchId

`func (o *NiatelemetryImpactedNodes) HasSwitchId() bool`

HasSwitchId returns a boolean if a field has been set.

### GetSwitchName

`func (o *NiatelemetryImpactedNodes) GetSwitchName() string`

GetSwitchName returns the SwitchName field if non-nil, zero value otherwise.

### GetSwitchNameOk

`func (o *NiatelemetryImpactedNodes) GetSwitchNameOk() (*string, bool)`

GetSwitchNameOk returns a tuple with the SwitchName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchName

`func (o *NiatelemetryImpactedNodes) SetSwitchName(v string)`

SetSwitchName sets SwitchName field to given value.

### HasSwitchName

`func (o *NiatelemetryImpactedNodes) HasSwitchName() bool`

HasSwitchName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


