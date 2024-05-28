# ApplianceDeviceClusterInstall

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "appliance.DeviceClusterInstall"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "appliance.DeviceClusterInstall"]
**CompletedPhases** | Pointer to [**[]ApplianceClusterInstallPhase**](ApplianceClusterInstallPhase.md) |  | [optional] 
**CurrentPhase** | Pointer to [**NullableApplianceClusterInstallPhase**](ApplianceClusterInstallPhase.md) |  | [optional] 
**ElapsedTime** | Pointer to **int64** | Elapsed time in seconds during the software install. | [optional] [readonly] 
**EndTime** | Pointer to **time.Time** | End date of the software install. | [optional] [readonly] 
**ErrorCode** | Pointer to **int64** | Error code for Intersight Appliance&#39;s software install. In case of failure - this code will help decide if software install can be retried. | [optional] [readonly] 
**Messages** | Pointer to **[]string** |  | [optional] 
**NodeInfo** | Pointer to [**[]ApplianceNodeIpInfo**](ApplianceNodeIpInfo.md) |  | [optional] 
**StartTime** | Pointer to **time.Time** | Start date of the software install. UI can modify startTime to re-schedule an install. | [optional] 
**Status** | Pointer to **string** | Status of the Intersight Appliance&#39;s software install. * &#x60;NotReady&#x60; - Cluster is not ready. Install cannot be triggered. * &#x60;Ready&#x60; - Cluster is ready. Install can be triggered. * &#x60;InProgress&#x60; - Install is currently in progress. * &#x60;Success&#x60; - Install was run and succeeded. * &#x60;Fail&#x60; - Install was run and failed. | [optional] [readonly] [default to "NotReady"]
**TotalNodes** | Pointer to **int64** | Total number of nodes in the system. | [optional] [readonly] 
**TotalPhases** | Pointer to **int64** | TotalPhase represents the total number of the install phases for one install. | [optional] [readonly] 
**RegisteredDevice** | Pointer to [**NullableAssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewApplianceDeviceClusterInstall

`func NewApplianceDeviceClusterInstall(classId string, objectType string, ) *ApplianceDeviceClusterInstall`

NewApplianceDeviceClusterInstall instantiates a new ApplianceDeviceClusterInstall object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplianceDeviceClusterInstallWithDefaults

`func NewApplianceDeviceClusterInstallWithDefaults() *ApplianceDeviceClusterInstall`

NewApplianceDeviceClusterInstallWithDefaults instantiates a new ApplianceDeviceClusterInstall object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ApplianceDeviceClusterInstall) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ApplianceDeviceClusterInstall) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ApplianceDeviceClusterInstall) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ApplianceDeviceClusterInstall) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ApplianceDeviceClusterInstall) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ApplianceDeviceClusterInstall) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCompletedPhases

`func (o *ApplianceDeviceClusterInstall) GetCompletedPhases() []ApplianceClusterInstallPhase`

GetCompletedPhases returns the CompletedPhases field if non-nil, zero value otherwise.

### GetCompletedPhasesOk

`func (o *ApplianceDeviceClusterInstall) GetCompletedPhasesOk() (*[]ApplianceClusterInstallPhase, bool)`

GetCompletedPhasesOk returns a tuple with the CompletedPhases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedPhases

`func (o *ApplianceDeviceClusterInstall) SetCompletedPhases(v []ApplianceClusterInstallPhase)`

SetCompletedPhases sets CompletedPhases field to given value.

### HasCompletedPhases

`func (o *ApplianceDeviceClusterInstall) HasCompletedPhases() bool`

HasCompletedPhases returns a boolean if a field has been set.

### SetCompletedPhasesNil

`func (o *ApplianceDeviceClusterInstall) SetCompletedPhasesNil(b bool)`

 SetCompletedPhasesNil sets the value for CompletedPhases to be an explicit nil

### UnsetCompletedPhases
`func (o *ApplianceDeviceClusterInstall) UnsetCompletedPhases()`

UnsetCompletedPhases ensures that no value is present for CompletedPhases, not even an explicit nil
### GetCurrentPhase

`func (o *ApplianceDeviceClusterInstall) GetCurrentPhase() ApplianceClusterInstallPhase`

GetCurrentPhase returns the CurrentPhase field if non-nil, zero value otherwise.

### GetCurrentPhaseOk

`func (o *ApplianceDeviceClusterInstall) GetCurrentPhaseOk() (*ApplianceClusterInstallPhase, bool)`

GetCurrentPhaseOk returns a tuple with the CurrentPhase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPhase

`func (o *ApplianceDeviceClusterInstall) SetCurrentPhase(v ApplianceClusterInstallPhase)`

SetCurrentPhase sets CurrentPhase field to given value.

### HasCurrentPhase

`func (o *ApplianceDeviceClusterInstall) HasCurrentPhase() bool`

HasCurrentPhase returns a boolean if a field has been set.

### SetCurrentPhaseNil

`func (o *ApplianceDeviceClusterInstall) SetCurrentPhaseNil(b bool)`

 SetCurrentPhaseNil sets the value for CurrentPhase to be an explicit nil

### UnsetCurrentPhase
`func (o *ApplianceDeviceClusterInstall) UnsetCurrentPhase()`

UnsetCurrentPhase ensures that no value is present for CurrentPhase, not even an explicit nil
### GetElapsedTime

`func (o *ApplianceDeviceClusterInstall) GetElapsedTime() int64`

GetElapsedTime returns the ElapsedTime field if non-nil, zero value otherwise.

### GetElapsedTimeOk

`func (o *ApplianceDeviceClusterInstall) GetElapsedTimeOk() (*int64, bool)`

GetElapsedTimeOk returns a tuple with the ElapsedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElapsedTime

`func (o *ApplianceDeviceClusterInstall) SetElapsedTime(v int64)`

SetElapsedTime sets ElapsedTime field to given value.

### HasElapsedTime

`func (o *ApplianceDeviceClusterInstall) HasElapsedTime() bool`

HasElapsedTime returns a boolean if a field has been set.

### GetEndTime

`func (o *ApplianceDeviceClusterInstall) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *ApplianceDeviceClusterInstall) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *ApplianceDeviceClusterInstall) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *ApplianceDeviceClusterInstall) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetErrorCode

`func (o *ApplianceDeviceClusterInstall) GetErrorCode() int64`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *ApplianceDeviceClusterInstall) GetErrorCodeOk() (*int64, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *ApplianceDeviceClusterInstall) SetErrorCode(v int64)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *ApplianceDeviceClusterInstall) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetMessages

`func (o *ApplianceDeviceClusterInstall) GetMessages() []string`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *ApplianceDeviceClusterInstall) GetMessagesOk() (*[]string, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *ApplianceDeviceClusterInstall) SetMessages(v []string)`

SetMessages sets Messages field to given value.

### HasMessages

`func (o *ApplianceDeviceClusterInstall) HasMessages() bool`

HasMessages returns a boolean if a field has been set.

### SetMessagesNil

`func (o *ApplianceDeviceClusterInstall) SetMessagesNil(b bool)`

 SetMessagesNil sets the value for Messages to be an explicit nil

### UnsetMessages
`func (o *ApplianceDeviceClusterInstall) UnsetMessages()`

UnsetMessages ensures that no value is present for Messages, not even an explicit nil
### GetNodeInfo

`func (o *ApplianceDeviceClusterInstall) GetNodeInfo() []ApplianceNodeIpInfo`

GetNodeInfo returns the NodeInfo field if non-nil, zero value otherwise.

### GetNodeInfoOk

`func (o *ApplianceDeviceClusterInstall) GetNodeInfoOk() (*[]ApplianceNodeIpInfo, bool)`

GetNodeInfoOk returns a tuple with the NodeInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeInfo

`func (o *ApplianceDeviceClusterInstall) SetNodeInfo(v []ApplianceNodeIpInfo)`

SetNodeInfo sets NodeInfo field to given value.

### HasNodeInfo

`func (o *ApplianceDeviceClusterInstall) HasNodeInfo() bool`

HasNodeInfo returns a boolean if a field has been set.

### SetNodeInfoNil

`func (o *ApplianceDeviceClusterInstall) SetNodeInfoNil(b bool)`

 SetNodeInfoNil sets the value for NodeInfo to be an explicit nil

### UnsetNodeInfo
`func (o *ApplianceDeviceClusterInstall) UnsetNodeInfo()`

UnsetNodeInfo ensures that no value is present for NodeInfo, not even an explicit nil
### GetStartTime

`func (o *ApplianceDeviceClusterInstall) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *ApplianceDeviceClusterInstall) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *ApplianceDeviceClusterInstall) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *ApplianceDeviceClusterInstall) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetStatus

`func (o *ApplianceDeviceClusterInstall) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ApplianceDeviceClusterInstall) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ApplianceDeviceClusterInstall) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ApplianceDeviceClusterInstall) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTotalNodes

`func (o *ApplianceDeviceClusterInstall) GetTotalNodes() int64`

GetTotalNodes returns the TotalNodes field if non-nil, zero value otherwise.

### GetTotalNodesOk

`func (o *ApplianceDeviceClusterInstall) GetTotalNodesOk() (*int64, bool)`

GetTotalNodesOk returns a tuple with the TotalNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalNodes

`func (o *ApplianceDeviceClusterInstall) SetTotalNodes(v int64)`

SetTotalNodes sets TotalNodes field to given value.

### HasTotalNodes

`func (o *ApplianceDeviceClusterInstall) HasTotalNodes() bool`

HasTotalNodes returns a boolean if a field has been set.

### GetTotalPhases

`func (o *ApplianceDeviceClusterInstall) GetTotalPhases() int64`

GetTotalPhases returns the TotalPhases field if non-nil, zero value otherwise.

### GetTotalPhasesOk

`func (o *ApplianceDeviceClusterInstall) GetTotalPhasesOk() (*int64, bool)`

GetTotalPhasesOk returns a tuple with the TotalPhases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPhases

`func (o *ApplianceDeviceClusterInstall) SetTotalPhases(v int64)`

SetTotalPhases sets TotalPhases field to given value.

### HasTotalPhases

`func (o *ApplianceDeviceClusterInstall) HasTotalPhases() bool`

HasTotalPhases returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *ApplianceDeviceClusterInstall) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *ApplianceDeviceClusterInstall) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *ApplianceDeviceClusterInstall) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *ApplianceDeviceClusterInstall) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### SetRegisteredDeviceNil

`func (o *ApplianceDeviceClusterInstall) SetRegisteredDeviceNil(b bool)`

 SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil

### UnsetRegisteredDevice
`func (o *ApplianceDeviceClusterInstall) UnsetRegisteredDevice()`

UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


