# ApplianceAppOpStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "appliance.AppOpStatus"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "appliance.AppOpStatus"]
**ApiStatuses** | Pointer to [**[]ApplianceApiStatus**](ApplianceApiStatus.md) |  | [optional] 
**AppLabel** | Pointer to **string** | Unique label to identify the application. | [optional] [readonly] 
**OperationalStatus** | Pointer to **string** | Operational status of the application. Operational status is based on the result of the status checks. If result of any check is Critical, then its value is Impaired. Otherwise, if result of any check is Warning, then its value is AttentionNeeded. If all checks are OK, then its value is Operational. * &#x60;Unknown&#x60; - The status of the appliance node is unknown. * &#x60;Operational&#x60; - The appliance node is operational. * &#x60;Impaired&#x60; - The appliance node is impaired. * &#x60;AttentionNeeded&#x60; - The appliance node needs attention. * &#x60;ReadyToJoin&#x60; - The node is ready to be added to a standalone Intersight Appliance to form a cluster. * &#x60;OutOfService&#x60; - The user has taken this node (part of a cluster) to out of service. * &#x60;ReadyForReplacement&#x60; - The cluster node is ready to be replaced. * &#x60;ReplacementInProgress&#x60; - The cluster node replacement is in progress. * &#x60;ReplacementFailed&#x60; - There was a failure during the cluster node replacement. | [optional] [readonly] [default to "Unknown"]
**ReadyCount** | Pointer to **int64** | Number of replicas ready.  The number of instances of the application currently ready to perform its intended functions. | [optional] [readonly] 
**ReplicaCount** | Pointer to **int64** | Number of replicas provisioned. The number of instances of the application provisioned to run on the Intersight appliance. | [optional] [readonly] 
**RestartCount1Hour** | Pointer to **int64** | Number of instance restarts in the last hour. | [optional] [readonly] 
**RestartCount24Hours** | Pointer to **int64** | Number of instance restarts in the last 24 hours. | [optional] [readonly] 
**RestartCount5Mins** | Pointer to **int64** | Number of instance restarts in the last 5 minutes. | [optional] [readonly] 
**RestartCountTotal** | Pointer to **int64** | Total number of restarts since last deployment. | [optional] [readonly] 
**RunningCount** | Pointer to **int64** | Number of replicas running. The number of instances of the application currently running. | [optional] [readonly] 
**GroupOpStatus** | Pointer to [**NullableApplianceGroupOpStatusRelationship**](ApplianceGroupOpStatusRelationship.md) |  | [optional] 
**SystemOpStatus** | Pointer to [**NullableApplianceSystemOpStatusRelationship**](ApplianceSystemOpStatusRelationship.md) |  | [optional] 

## Methods

### NewApplianceAppOpStatus

`func NewApplianceAppOpStatus(classId string, objectType string, ) *ApplianceAppOpStatus`

NewApplianceAppOpStatus instantiates a new ApplianceAppOpStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplianceAppOpStatusWithDefaults

`func NewApplianceAppOpStatusWithDefaults() *ApplianceAppOpStatus`

NewApplianceAppOpStatusWithDefaults instantiates a new ApplianceAppOpStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ApplianceAppOpStatus) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ApplianceAppOpStatus) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ApplianceAppOpStatus) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ApplianceAppOpStatus) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ApplianceAppOpStatus) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ApplianceAppOpStatus) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetApiStatuses

`func (o *ApplianceAppOpStatus) GetApiStatuses() []ApplianceApiStatus`

GetApiStatuses returns the ApiStatuses field if non-nil, zero value otherwise.

### GetApiStatusesOk

`func (o *ApplianceAppOpStatus) GetApiStatusesOk() (*[]ApplianceApiStatus, bool)`

GetApiStatusesOk returns a tuple with the ApiStatuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiStatuses

`func (o *ApplianceAppOpStatus) SetApiStatuses(v []ApplianceApiStatus)`

SetApiStatuses sets ApiStatuses field to given value.

### HasApiStatuses

`func (o *ApplianceAppOpStatus) HasApiStatuses() bool`

HasApiStatuses returns a boolean if a field has been set.

### SetApiStatusesNil

`func (o *ApplianceAppOpStatus) SetApiStatusesNil(b bool)`

 SetApiStatusesNil sets the value for ApiStatuses to be an explicit nil

### UnsetApiStatuses
`func (o *ApplianceAppOpStatus) UnsetApiStatuses()`

UnsetApiStatuses ensures that no value is present for ApiStatuses, not even an explicit nil
### GetAppLabel

`func (o *ApplianceAppOpStatus) GetAppLabel() string`

GetAppLabel returns the AppLabel field if non-nil, zero value otherwise.

### GetAppLabelOk

`func (o *ApplianceAppOpStatus) GetAppLabelOk() (*string, bool)`

GetAppLabelOk returns a tuple with the AppLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppLabel

`func (o *ApplianceAppOpStatus) SetAppLabel(v string)`

SetAppLabel sets AppLabel field to given value.

### HasAppLabel

`func (o *ApplianceAppOpStatus) HasAppLabel() bool`

HasAppLabel returns a boolean if a field has been set.

### GetOperationalStatus

`func (o *ApplianceAppOpStatus) GetOperationalStatus() string`

GetOperationalStatus returns the OperationalStatus field if non-nil, zero value otherwise.

### GetOperationalStatusOk

`func (o *ApplianceAppOpStatus) GetOperationalStatusOk() (*string, bool)`

GetOperationalStatusOk returns a tuple with the OperationalStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationalStatus

`func (o *ApplianceAppOpStatus) SetOperationalStatus(v string)`

SetOperationalStatus sets OperationalStatus field to given value.

### HasOperationalStatus

`func (o *ApplianceAppOpStatus) HasOperationalStatus() bool`

HasOperationalStatus returns a boolean if a field has been set.

### GetReadyCount

`func (o *ApplianceAppOpStatus) GetReadyCount() int64`

GetReadyCount returns the ReadyCount field if non-nil, zero value otherwise.

### GetReadyCountOk

`func (o *ApplianceAppOpStatus) GetReadyCountOk() (*int64, bool)`

GetReadyCountOk returns a tuple with the ReadyCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadyCount

`func (o *ApplianceAppOpStatus) SetReadyCount(v int64)`

SetReadyCount sets ReadyCount field to given value.

### HasReadyCount

`func (o *ApplianceAppOpStatus) HasReadyCount() bool`

HasReadyCount returns a boolean if a field has been set.

### GetReplicaCount

`func (o *ApplianceAppOpStatus) GetReplicaCount() int64`

GetReplicaCount returns the ReplicaCount field if non-nil, zero value otherwise.

### GetReplicaCountOk

`func (o *ApplianceAppOpStatus) GetReplicaCountOk() (*int64, bool)`

GetReplicaCountOk returns a tuple with the ReplicaCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicaCount

`func (o *ApplianceAppOpStatus) SetReplicaCount(v int64)`

SetReplicaCount sets ReplicaCount field to given value.

### HasReplicaCount

`func (o *ApplianceAppOpStatus) HasReplicaCount() bool`

HasReplicaCount returns a boolean if a field has been set.

### GetRestartCount1Hour

`func (o *ApplianceAppOpStatus) GetRestartCount1Hour() int64`

GetRestartCount1Hour returns the RestartCount1Hour field if non-nil, zero value otherwise.

### GetRestartCount1HourOk

`func (o *ApplianceAppOpStatus) GetRestartCount1HourOk() (*int64, bool)`

GetRestartCount1HourOk returns a tuple with the RestartCount1Hour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestartCount1Hour

`func (o *ApplianceAppOpStatus) SetRestartCount1Hour(v int64)`

SetRestartCount1Hour sets RestartCount1Hour field to given value.

### HasRestartCount1Hour

`func (o *ApplianceAppOpStatus) HasRestartCount1Hour() bool`

HasRestartCount1Hour returns a boolean if a field has been set.

### GetRestartCount24Hours

`func (o *ApplianceAppOpStatus) GetRestartCount24Hours() int64`

GetRestartCount24Hours returns the RestartCount24Hours field if non-nil, zero value otherwise.

### GetRestartCount24HoursOk

`func (o *ApplianceAppOpStatus) GetRestartCount24HoursOk() (*int64, bool)`

GetRestartCount24HoursOk returns a tuple with the RestartCount24Hours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestartCount24Hours

`func (o *ApplianceAppOpStatus) SetRestartCount24Hours(v int64)`

SetRestartCount24Hours sets RestartCount24Hours field to given value.

### HasRestartCount24Hours

`func (o *ApplianceAppOpStatus) HasRestartCount24Hours() bool`

HasRestartCount24Hours returns a boolean if a field has been set.

### GetRestartCount5Mins

`func (o *ApplianceAppOpStatus) GetRestartCount5Mins() int64`

GetRestartCount5Mins returns the RestartCount5Mins field if non-nil, zero value otherwise.

### GetRestartCount5MinsOk

`func (o *ApplianceAppOpStatus) GetRestartCount5MinsOk() (*int64, bool)`

GetRestartCount5MinsOk returns a tuple with the RestartCount5Mins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestartCount5Mins

`func (o *ApplianceAppOpStatus) SetRestartCount5Mins(v int64)`

SetRestartCount5Mins sets RestartCount5Mins field to given value.

### HasRestartCount5Mins

`func (o *ApplianceAppOpStatus) HasRestartCount5Mins() bool`

HasRestartCount5Mins returns a boolean if a field has been set.

### GetRestartCountTotal

`func (o *ApplianceAppOpStatus) GetRestartCountTotal() int64`

GetRestartCountTotal returns the RestartCountTotal field if non-nil, zero value otherwise.

### GetRestartCountTotalOk

`func (o *ApplianceAppOpStatus) GetRestartCountTotalOk() (*int64, bool)`

GetRestartCountTotalOk returns a tuple with the RestartCountTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestartCountTotal

`func (o *ApplianceAppOpStatus) SetRestartCountTotal(v int64)`

SetRestartCountTotal sets RestartCountTotal field to given value.

### HasRestartCountTotal

`func (o *ApplianceAppOpStatus) HasRestartCountTotal() bool`

HasRestartCountTotal returns a boolean if a field has been set.

### GetRunningCount

`func (o *ApplianceAppOpStatus) GetRunningCount() int64`

GetRunningCount returns the RunningCount field if non-nil, zero value otherwise.

### GetRunningCountOk

`func (o *ApplianceAppOpStatus) GetRunningCountOk() (*int64, bool)`

GetRunningCountOk returns a tuple with the RunningCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunningCount

`func (o *ApplianceAppOpStatus) SetRunningCount(v int64)`

SetRunningCount sets RunningCount field to given value.

### HasRunningCount

`func (o *ApplianceAppOpStatus) HasRunningCount() bool`

HasRunningCount returns a boolean if a field has been set.

### GetGroupOpStatus

`func (o *ApplianceAppOpStatus) GetGroupOpStatus() ApplianceGroupOpStatusRelationship`

GetGroupOpStatus returns the GroupOpStatus field if non-nil, zero value otherwise.

### GetGroupOpStatusOk

`func (o *ApplianceAppOpStatus) GetGroupOpStatusOk() (*ApplianceGroupOpStatusRelationship, bool)`

GetGroupOpStatusOk returns a tuple with the GroupOpStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupOpStatus

`func (o *ApplianceAppOpStatus) SetGroupOpStatus(v ApplianceGroupOpStatusRelationship)`

SetGroupOpStatus sets GroupOpStatus field to given value.

### HasGroupOpStatus

`func (o *ApplianceAppOpStatus) HasGroupOpStatus() bool`

HasGroupOpStatus returns a boolean if a field has been set.

### SetGroupOpStatusNil

`func (o *ApplianceAppOpStatus) SetGroupOpStatusNil(b bool)`

 SetGroupOpStatusNil sets the value for GroupOpStatus to be an explicit nil

### UnsetGroupOpStatus
`func (o *ApplianceAppOpStatus) UnsetGroupOpStatus()`

UnsetGroupOpStatus ensures that no value is present for GroupOpStatus, not even an explicit nil
### GetSystemOpStatus

`func (o *ApplianceAppOpStatus) GetSystemOpStatus() ApplianceSystemOpStatusRelationship`

GetSystemOpStatus returns the SystemOpStatus field if non-nil, zero value otherwise.

### GetSystemOpStatusOk

`func (o *ApplianceAppOpStatus) GetSystemOpStatusOk() (*ApplianceSystemOpStatusRelationship, bool)`

GetSystemOpStatusOk returns a tuple with the SystemOpStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemOpStatus

`func (o *ApplianceAppOpStatus) SetSystemOpStatus(v ApplianceSystemOpStatusRelationship)`

SetSystemOpStatus sets SystemOpStatus field to given value.

### HasSystemOpStatus

`func (o *ApplianceAppOpStatus) HasSystemOpStatus() bool`

HasSystemOpStatus returns a boolean if a field has been set.

### SetSystemOpStatusNil

`func (o *ApplianceAppOpStatus) SetSystemOpStatusNil(b bool)`

 SetSystemOpStatusNil sets the value for SystemOpStatus to be an explicit nil

### UnsetSystemOpStatus
`func (o *ApplianceAppOpStatus) UnsetSystemOpStatus()`

UnsetSystemOpStatus ensures that no value is present for SystemOpStatus, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


