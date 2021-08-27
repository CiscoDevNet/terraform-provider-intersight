# ApplianceAppStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "appliance.AppStatus"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "appliance.AppStatus"]
**ApiStatuses** | Pointer to [**[]ApplianceApiStatus**](ApplianceApiStatus.md) |  | [optional] 
**AppLabel** | Pointer to **string** | Unique label to identify the application. | [optional] [readonly] 
**OperationalStatus** | Pointer to **string** | Operational status of the application. Operational status is based on the result of the status checks. If result of any check is Critical, then its value is Impaired. Otherwise, if result of any check is Warning, then its value is AttentionNeeded. If all checks are OK, then its value is Operational. * &#x60;Unknown&#x60; - Operational status of the Intersight Appliance entity is Unknown. * &#x60;Operational&#x60; - Operational status of the Intersight Appliance entity is Operational. * &#x60;Impaired&#x60; - Operational status of the Intersight Appliance entity is Impaired. * &#x60;AttentionNeeded&#x60; - Operational status of the Intersight Appliance entity is AttentionNeeded. | [optional] [readonly] [default to "Unknown"]
**ReadyCount** | Pointer to **int64** | Number of replicas ready.  The number of instances of the application currently ready to perform its intended functions. | [optional] [readonly] 
**ReplicaCount** | Pointer to **int64** | Number of replicas provisioned. The number of instances of the application provisioned to run on the Intersight appliance. | [optional] [readonly] 
**RestartCount1Hour** | Pointer to **int64** | Number of instance restarts in the last hour. | [optional] [readonly] 
**RestartCount24Hours** | Pointer to **int64** | Number of instance restarts in the last 24 hours. | [optional] [readonly] 
**RestartCount5Mins** | Pointer to **int64** | Number of instance restarts in the last 5 minutes. | [optional] [readonly] 
**RestartCountTotal** | Pointer to **int64** | Total number of restarts since last deployment. | [optional] [readonly] 
**RunningCount** | Pointer to **int64** | Number of replicas running. The number of instances of the application currently running. | [optional] [readonly] 
**StatusChecks** | Pointer to [**[]ApplianceStatusCheck**](ApplianceStatusCheck.md) |  | [optional] 
**GroupStatus** | Pointer to [**ApplianceGroupStatusRelationship**](ApplianceGroupStatusRelationship.md) |  | [optional] 
**SystemStatus** | Pointer to [**ApplianceSystemStatusRelationship**](ApplianceSystemStatusRelationship.md) |  | [optional] 

## Methods

### NewApplianceAppStatus

`func NewApplianceAppStatus(classId string, objectType string, ) *ApplianceAppStatus`

NewApplianceAppStatus instantiates a new ApplianceAppStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplianceAppStatusWithDefaults

`func NewApplianceAppStatusWithDefaults() *ApplianceAppStatus`

NewApplianceAppStatusWithDefaults instantiates a new ApplianceAppStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ApplianceAppStatus) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ApplianceAppStatus) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ApplianceAppStatus) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ApplianceAppStatus) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ApplianceAppStatus) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ApplianceAppStatus) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetApiStatuses

`func (o *ApplianceAppStatus) GetApiStatuses() []ApplianceApiStatus`

GetApiStatuses returns the ApiStatuses field if non-nil, zero value otherwise.

### GetApiStatusesOk

`func (o *ApplianceAppStatus) GetApiStatusesOk() (*[]ApplianceApiStatus, bool)`

GetApiStatusesOk returns a tuple with the ApiStatuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiStatuses

`func (o *ApplianceAppStatus) SetApiStatuses(v []ApplianceApiStatus)`

SetApiStatuses sets ApiStatuses field to given value.

### HasApiStatuses

`func (o *ApplianceAppStatus) HasApiStatuses() bool`

HasApiStatuses returns a boolean if a field has been set.

### SetApiStatusesNil

`func (o *ApplianceAppStatus) SetApiStatusesNil(b bool)`

 SetApiStatusesNil sets the value for ApiStatuses to be an explicit nil

### UnsetApiStatuses
`func (o *ApplianceAppStatus) UnsetApiStatuses()`

UnsetApiStatuses ensures that no value is present for ApiStatuses, not even an explicit nil
### GetAppLabel

`func (o *ApplianceAppStatus) GetAppLabel() string`

GetAppLabel returns the AppLabel field if non-nil, zero value otherwise.

### GetAppLabelOk

`func (o *ApplianceAppStatus) GetAppLabelOk() (*string, bool)`

GetAppLabelOk returns a tuple with the AppLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppLabel

`func (o *ApplianceAppStatus) SetAppLabel(v string)`

SetAppLabel sets AppLabel field to given value.

### HasAppLabel

`func (o *ApplianceAppStatus) HasAppLabel() bool`

HasAppLabel returns a boolean if a field has been set.

### GetOperationalStatus

`func (o *ApplianceAppStatus) GetOperationalStatus() string`

GetOperationalStatus returns the OperationalStatus field if non-nil, zero value otherwise.

### GetOperationalStatusOk

`func (o *ApplianceAppStatus) GetOperationalStatusOk() (*string, bool)`

GetOperationalStatusOk returns a tuple with the OperationalStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationalStatus

`func (o *ApplianceAppStatus) SetOperationalStatus(v string)`

SetOperationalStatus sets OperationalStatus field to given value.

### HasOperationalStatus

`func (o *ApplianceAppStatus) HasOperationalStatus() bool`

HasOperationalStatus returns a boolean if a field has been set.

### GetReadyCount

`func (o *ApplianceAppStatus) GetReadyCount() int64`

GetReadyCount returns the ReadyCount field if non-nil, zero value otherwise.

### GetReadyCountOk

`func (o *ApplianceAppStatus) GetReadyCountOk() (*int64, bool)`

GetReadyCountOk returns a tuple with the ReadyCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadyCount

`func (o *ApplianceAppStatus) SetReadyCount(v int64)`

SetReadyCount sets ReadyCount field to given value.

### HasReadyCount

`func (o *ApplianceAppStatus) HasReadyCount() bool`

HasReadyCount returns a boolean if a field has been set.

### GetReplicaCount

`func (o *ApplianceAppStatus) GetReplicaCount() int64`

GetReplicaCount returns the ReplicaCount field if non-nil, zero value otherwise.

### GetReplicaCountOk

`func (o *ApplianceAppStatus) GetReplicaCountOk() (*int64, bool)`

GetReplicaCountOk returns a tuple with the ReplicaCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicaCount

`func (o *ApplianceAppStatus) SetReplicaCount(v int64)`

SetReplicaCount sets ReplicaCount field to given value.

### HasReplicaCount

`func (o *ApplianceAppStatus) HasReplicaCount() bool`

HasReplicaCount returns a boolean if a field has been set.

### GetRestartCount1Hour

`func (o *ApplianceAppStatus) GetRestartCount1Hour() int64`

GetRestartCount1Hour returns the RestartCount1Hour field if non-nil, zero value otherwise.

### GetRestartCount1HourOk

`func (o *ApplianceAppStatus) GetRestartCount1HourOk() (*int64, bool)`

GetRestartCount1HourOk returns a tuple with the RestartCount1Hour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestartCount1Hour

`func (o *ApplianceAppStatus) SetRestartCount1Hour(v int64)`

SetRestartCount1Hour sets RestartCount1Hour field to given value.

### HasRestartCount1Hour

`func (o *ApplianceAppStatus) HasRestartCount1Hour() bool`

HasRestartCount1Hour returns a boolean if a field has been set.

### GetRestartCount24Hours

`func (o *ApplianceAppStatus) GetRestartCount24Hours() int64`

GetRestartCount24Hours returns the RestartCount24Hours field if non-nil, zero value otherwise.

### GetRestartCount24HoursOk

`func (o *ApplianceAppStatus) GetRestartCount24HoursOk() (*int64, bool)`

GetRestartCount24HoursOk returns a tuple with the RestartCount24Hours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestartCount24Hours

`func (o *ApplianceAppStatus) SetRestartCount24Hours(v int64)`

SetRestartCount24Hours sets RestartCount24Hours field to given value.

### HasRestartCount24Hours

`func (o *ApplianceAppStatus) HasRestartCount24Hours() bool`

HasRestartCount24Hours returns a boolean if a field has been set.

### GetRestartCount5Mins

`func (o *ApplianceAppStatus) GetRestartCount5Mins() int64`

GetRestartCount5Mins returns the RestartCount5Mins field if non-nil, zero value otherwise.

### GetRestartCount5MinsOk

`func (o *ApplianceAppStatus) GetRestartCount5MinsOk() (*int64, bool)`

GetRestartCount5MinsOk returns a tuple with the RestartCount5Mins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestartCount5Mins

`func (o *ApplianceAppStatus) SetRestartCount5Mins(v int64)`

SetRestartCount5Mins sets RestartCount5Mins field to given value.

### HasRestartCount5Mins

`func (o *ApplianceAppStatus) HasRestartCount5Mins() bool`

HasRestartCount5Mins returns a boolean if a field has been set.

### GetRestartCountTotal

`func (o *ApplianceAppStatus) GetRestartCountTotal() int64`

GetRestartCountTotal returns the RestartCountTotal field if non-nil, zero value otherwise.

### GetRestartCountTotalOk

`func (o *ApplianceAppStatus) GetRestartCountTotalOk() (*int64, bool)`

GetRestartCountTotalOk returns a tuple with the RestartCountTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestartCountTotal

`func (o *ApplianceAppStatus) SetRestartCountTotal(v int64)`

SetRestartCountTotal sets RestartCountTotal field to given value.

### HasRestartCountTotal

`func (o *ApplianceAppStatus) HasRestartCountTotal() bool`

HasRestartCountTotal returns a boolean if a field has been set.

### GetRunningCount

`func (o *ApplianceAppStatus) GetRunningCount() int64`

GetRunningCount returns the RunningCount field if non-nil, zero value otherwise.

### GetRunningCountOk

`func (o *ApplianceAppStatus) GetRunningCountOk() (*int64, bool)`

GetRunningCountOk returns a tuple with the RunningCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunningCount

`func (o *ApplianceAppStatus) SetRunningCount(v int64)`

SetRunningCount sets RunningCount field to given value.

### HasRunningCount

`func (o *ApplianceAppStatus) HasRunningCount() bool`

HasRunningCount returns a boolean if a field has been set.

### GetStatusChecks

`func (o *ApplianceAppStatus) GetStatusChecks() []ApplianceStatusCheck`

GetStatusChecks returns the StatusChecks field if non-nil, zero value otherwise.

### GetStatusChecksOk

`func (o *ApplianceAppStatus) GetStatusChecksOk() (*[]ApplianceStatusCheck, bool)`

GetStatusChecksOk returns a tuple with the StatusChecks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusChecks

`func (o *ApplianceAppStatus) SetStatusChecks(v []ApplianceStatusCheck)`

SetStatusChecks sets StatusChecks field to given value.

### HasStatusChecks

`func (o *ApplianceAppStatus) HasStatusChecks() bool`

HasStatusChecks returns a boolean if a field has been set.

### SetStatusChecksNil

`func (o *ApplianceAppStatus) SetStatusChecksNil(b bool)`

 SetStatusChecksNil sets the value for StatusChecks to be an explicit nil

### UnsetStatusChecks
`func (o *ApplianceAppStatus) UnsetStatusChecks()`

UnsetStatusChecks ensures that no value is present for StatusChecks, not even an explicit nil
### GetGroupStatus

`func (o *ApplianceAppStatus) GetGroupStatus() ApplianceGroupStatusRelationship`

GetGroupStatus returns the GroupStatus field if non-nil, zero value otherwise.

### GetGroupStatusOk

`func (o *ApplianceAppStatus) GetGroupStatusOk() (*ApplianceGroupStatusRelationship, bool)`

GetGroupStatusOk returns a tuple with the GroupStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupStatus

`func (o *ApplianceAppStatus) SetGroupStatus(v ApplianceGroupStatusRelationship)`

SetGroupStatus sets GroupStatus field to given value.

### HasGroupStatus

`func (o *ApplianceAppStatus) HasGroupStatus() bool`

HasGroupStatus returns a boolean if a field has been set.

### GetSystemStatus

`func (o *ApplianceAppStatus) GetSystemStatus() ApplianceSystemStatusRelationship`

GetSystemStatus returns the SystemStatus field if non-nil, zero value otherwise.

### GetSystemStatusOk

`func (o *ApplianceAppStatus) GetSystemStatusOk() (*ApplianceSystemStatusRelationship, bool)`

GetSystemStatusOk returns a tuple with the SystemStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemStatus

`func (o *ApplianceAppStatus) SetSystemStatus(v ApplianceSystemStatusRelationship)`

SetSystemStatus sets SystemStatus field to given value.

### HasSystemStatus

`func (o *ApplianceAppStatus) HasSystemStatus() bool`

HasSystemStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


