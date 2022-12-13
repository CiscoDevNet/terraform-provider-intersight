# ApplianceClusterInstallBase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**CompletedPhases** | Pointer to [**[]ApplianceClusterInstallPhase**](ApplianceClusterInstallPhase.md) |  | [optional] 
**CurrentPhase** | Pointer to [**NullableApplianceClusterInstallPhase**](ApplianceClusterInstallPhase.md) |  | [optional] 
**ElapsedTime** | Pointer to **int64** | Elapsed time in seconds during the software install. | [optional] [readonly] 
**EndTime** | Pointer to **time.Time** | End date of the software install. | [optional] [readonly] 
**ErrorCode** | Pointer to **int64** | Error code for Intersight Appliance&#39;s software install. In case of failure - this code will help decide if software install can be retried. | [optional] [readonly] 
**Messages** | Pointer to **[]string** |  | [optional] 
**NodeInfo** | Pointer to [**[]ApplianceNodeIpInfo**](ApplianceNodeIpInfo.md) |  | [optional] 
**RemoteDns** | Pointer to **string** | Round robin DNS address, which should be able to resolve the hostnames of all the nodes in the cluster. | [optional] 
**StartTime** | Pointer to **time.Time** | Start date of the software install. UI can modify startTime to re-schedule an install. | [optional] 
**Status** | Pointer to **string** | Status of the Intersight Appliance&#39;s software install. * &#x60;NotReady&#x60; - Cluster is not ready. Install cannot be triggered. * &#x60;Ready&#x60; - Cluster is ready. Install can be triggered. * &#x60;InProgress&#x60; - Install is currently in progress. * &#x60;Success&#x60; - Install was run and succeeded. * &#x60;Fail&#x60; - Install was run and failed. | [optional] [readonly] [default to "NotReady"]
**TotalNodes** | Pointer to **int64** | Total number of nodes in the system. | [optional] [readonly] 
**TotalPhases** | Pointer to **int64** | TotalPhase represents the total number of the install phases for one install. | [optional] [readonly] 
**Vip** | Pointer to [**NullableApplianceNodeIpInfo**](ApplianceNodeIpInfo.md) |  | [optional] 
**Account** | Pointer to [**IamAccountRelationship**](IamAccountRelationship.md) |  | [optional] 

## Methods

### NewApplianceClusterInstallBase

`func NewApplianceClusterInstallBase(classId string, objectType string, ) *ApplianceClusterInstallBase`

NewApplianceClusterInstallBase instantiates a new ApplianceClusterInstallBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplianceClusterInstallBaseWithDefaults

`func NewApplianceClusterInstallBaseWithDefaults() *ApplianceClusterInstallBase`

NewApplianceClusterInstallBaseWithDefaults instantiates a new ApplianceClusterInstallBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ApplianceClusterInstallBase) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ApplianceClusterInstallBase) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ApplianceClusterInstallBase) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ApplianceClusterInstallBase) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ApplianceClusterInstallBase) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ApplianceClusterInstallBase) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCompletedPhases

`func (o *ApplianceClusterInstallBase) GetCompletedPhases() []ApplianceClusterInstallPhase`

GetCompletedPhases returns the CompletedPhases field if non-nil, zero value otherwise.

### GetCompletedPhasesOk

`func (o *ApplianceClusterInstallBase) GetCompletedPhasesOk() (*[]ApplianceClusterInstallPhase, bool)`

GetCompletedPhasesOk returns a tuple with the CompletedPhases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedPhases

`func (o *ApplianceClusterInstallBase) SetCompletedPhases(v []ApplianceClusterInstallPhase)`

SetCompletedPhases sets CompletedPhases field to given value.

### HasCompletedPhases

`func (o *ApplianceClusterInstallBase) HasCompletedPhases() bool`

HasCompletedPhases returns a boolean if a field has been set.

### SetCompletedPhasesNil

`func (o *ApplianceClusterInstallBase) SetCompletedPhasesNil(b bool)`

 SetCompletedPhasesNil sets the value for CompletedPhases to be an explicit nil

### UnsetCompletedPhases
`func (o *ApplianceClusterInstallBase) UnsetCompletedPhases()`

UnsetCompletedPhases ensures that no value is present for CompletedPhases, not even an explicit nil
### GetCurrentPhase

`func (o *ApplianceClusterInstallBase) GetCurrentPhase() ApplianceClusterInstallPhase`

GetCurrentPhase returns the CurrentPhase field if non-nil, zero value otherwise.

### GetCurrentPhaseOk

`func (o *ApplianceClusterInstallBase) GetCurrentPhaseOk() (*ApplianceClusterInstallPhase, bool)`

GetCurrentPhaseOk returns a tuple with the CurrentPhase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPhase

`func (o *ApplianceClusterInstallBase) SetCurrentPhase(v ApplianceClusterInstallPhase)`

SetCurrentPhase sets CurrentPhase field to given value.

### HasCurrentPhase

`func (o *ApplianceClusterInstallBase) HasCurrentPhase() bool`

HasCurrentPhase returns a boolean if a field has been set.

### SetCurrentPhaseNil

`func (o *ApplianceClusterInstallBase) SetCurrentPhaseNil(b bool)`

 SetCurrentPhaseNil sets the value for CurrentPhase to be an explicit nil

### UnsetCurrentPhase
`func (o *ApplianceClusterInstallBase) UnsetCurrentPhase()`

UnsetCurrentPhase ensures that no value is present for CurrentPhase, not even an explicit nil
### GetElapsedTime

`func (o *ApplianceClusterInstallBase) GetElapsedTime() int64`

GetElapsedTime returns the ElapsedTime field if non-nil, zero value otherwise.

### GetElapsedTimeOk

`func (o *ApplianceClusterInstallBase) GetElapsedTimeOk() (*int64, bool)`

GetElapsedTimeOk returns a tuple with the ElapsedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElapsedTime

`func (o *ApplianceClusterInstallBase) SetElapsedTime(v int64)`

SetElapsedTime sets ElapsedTime field to given value.

### HasElapsedTime

`func (o *ApplianceClusterInstallBase) HasElapsedTime() bool`

HasElapsedTime returns a boolean if a field has been set.

### GetEndTime

`func (o *ApplianceClusterInstallBase) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *ApplianceClusterInstallBase) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *ApplianceClusterInstallBase) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *ApplianceClusterInstallBase) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetErrorCode

`func (o *ApplianceClusterInstallBase) GetErrorCode() int64`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *ApplianceClusterInstallBase) GetErrorCodeOk() (*int64, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *ApplianceClusterInstallBase) SetErrorCode(v int64)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *ApplianceClusterInstallBase) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetMessages

`func (o *ApplianceClusterInstallBase) GetMessages() []string`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *ApplianceClusterInstallBase) GetMessagesOk() (*[]string, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *ApplianceClusterInstallBase) SetMessages(v []string)`

SetMessages sets Messages field to given value.

### HasMessages

`func (o *ApplianceClusterInstallBase) HasMessages() bool`

HasMessages returns a boolean if a field has been set.

### SetMessagesNil

`func (o *ApplianceClusterInstallBase) SetMessagesNil(b bool)`

 SetMessagesNil sets the value for Messages to be an explicit nil

### UnsetMessages
`func (o *ApplianceClusterInstallBase) UnsetMessages()`

UnsetMessages ensures that no value is present for Messages, not even an explicit nil
### GetNodeInfo

`func (o *ApplianceClusterInstallBase) GetNodeInfo() []ApplianceNodeIpInfo`

GetNodeInfo returns the NodeInfo field if non-nil, zero value otherwise.

### GetNodeInfoOk

`func (o *ApplianceClusterInstallBase) GetNodeInfoOk() (*[]ApplianceNodeIpInfo, bool)`

GetNodeInfoOk returns a tuple with the NodeInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeInfo

`func (o *ApplianceClusterInstallBase) SetNodeInfo(v []ApplianceNodeIpInfo)`

SetNodeInfo sets NodeInfo field to given value.

### HasNodeInfo

`func (o *ApplianceClusterInstallBase) HasNodeInfo() bool`

HasNodeInfo returns a boolean if a field has been set.

### SetNodeInfoNil

`func (o *ApplianceClusterInstallBase) SetNodeInfoNil(b bool)`

 SetNodeInfoNil sets the value for NodeInfo to be an explicit nil

### UnsetNodeInfo
`func (o *ApplianceClusterInstallBase) UnsetNodeInfo()`

UnsetNodeInfo ensures that no value is present for NodeInfo, not even an explicit nil
### GetRemoteDns

`func (o *ApplianceClusterInstallBase) GetRemoteDns() string`

GetRemoteDns returns the RemoteDns field if non-nil, zero value otherwise.

### GetRemoteDnsOk

`func (o *ApplianceClusterInstallBase) GetRemoteDnsOk() (*string, bool)`

GetRemoteDnsOk returns a tuple with the RemoteDns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteDns

`func (o *ApplianceClusterInstallBase) SetRemoteDns(v string)`

SetRemoteDns sets RemoteDns field to given value.

### HasRemoteDns

`func (o *ApplianceClusterInstallBase) HasRemoteDns() bool`

HasRemoteDns returns a boolean if a field has been set.

### GetStartTime

`func (o *ApplianceClusterInstallBase) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *ApplianceClusterInstallBase) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *ApplianceClusterInstallBase) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *ApplianceClusterInstallBase) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetStatus

`func (o *ApplianceClusterInstallBase) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ApplianceClusterInstallBase) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ApplianceClusterInstallBase) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ApplianceClusterInstallBase) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTotalNodes

`func (o *ApplianceClusterInstallBase) GetTotalNodes() int64`

GetTotalNodes returns the TotalNodes field if non-nil, zero value otherwise.

### GetTotalNodesOk

`func (o *ApplianceClusterInstallBase) GetTotalNodesOk() (*int64, bool)`

GetTotalNodesOk returns a tuple with the TotalNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalNodes

`func (o *ApplianceClusterInstallBase) SetTotalNodes(v int64)`

SetTotalNodes sets TotalNodes field to given value.

### HasTotalNodes

`func (o *ApplianceClusterInstallBase) HasTotalNodes() bool`

HasTotalNodes returns a boolean if a field has been set.

### GetTotalPhases

`func (o *ApplianceClusterInstallBase) GetTotalPhases() int64`

GetTotalPhases returns the TotalPhases field if non-nil, zero value otherwise.

### GetTotalPhasesOk

`func (o *ApplianceClusterInstallBase) GetTotalPhasesOk() (*int64, bool)`

GetTotalPhasesOk returns a tuple with the TotalPhases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPhases

`func (o *ApplianceClusterInstallBase) SetTotalPhases(v int64)`

SetTotalPhases sets TotalPhases field to given value.

### HasTotalPhases

`func (o *ApplianceClusterInstallBase) HasTotalPhases() bool`

HasTotalPhases returns a boolean if a field has been set.

### GetVip

`func (o *ApplianceClusterInstallBase) GetVip() ApplianceNodeIpInfo`

GetVip returns the Vip field if non-nil, zero value otherwise.

### GetVipOk

`func (o *ApplianceClusterInstallBase) GetVipOk() (*ApplianceNodeIpInfo, bool)`

GetVipOk returns a tuple with the Vip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVip

`func (o *ApplianceClusterInstallBase) SetVip(v ApplianceNodeIpInfo)`

SetVip sets Vip field to given value.

### HasVip

`func (o *ApplianceClusterInstallBase) HasVip() bool`

HasVip returns a boolean if a field has been set.

### SetVipNil

`func (o *ApplianceClusterInstallBase) SetVipNil(b bool)`

 SetVipNil sets the value for Vip to be an explicit nil

### UnsetVip
`func (o *ApplianceClusterInstallBase) UnsetVip()`

UnsetVip ensures that no value is present for Vip, not even an explicit nil
### GetAccount

`func (o *ApplianceClusterInstallBase) GetAccount() IamAccountRelationship`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *ApplianceClusterInstallBase) GetAccountOk() (*IamAccountRelationship, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *ApplianceClusterInstallBase) SetAccount(v IamAccountRelationship)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *ApplianceClusterInstallBase) HasAccount() bool`

HasAccount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


