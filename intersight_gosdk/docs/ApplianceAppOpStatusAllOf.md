# ApplianceAppOpStatusAllOf

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
**GroupOpStatus** | Pointer to [**ApplianceGroupOpStatusRelationship**](ApplianceGroupOpStatusRelationship.md) |  | [optional] 
**SystemOpStatus** | Pointer to [**ApplianceSystemOpStatusRelationship**](ApplianceSystemOpStatusRelationship.md) |  | [optional] 

## Methods

### NewApplianceAppOpStatusAllOf

`func NewApplianceAppOpStatusAllOf(classId string, objectType string, ) *ApplianceAppOpStatusAllOf`

NewApplianceAppOpStatusAllOf instantiates a new ApplianceAppOpStatusAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplianceAppOpStatusAllOfWithDefaults

`func NewApplianceAppOpStatusAllOfWithDefaults() *ApplianceAppOpStatusAllOf`

NewApplianceAppOpStatusAllOfWithDefaults instantiates a new ApplianceAppOpStatusAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ApplianceAppOpStatusAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ApplianceAppOpStatusAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ApplianceAppOpStatusAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ApplianceAppOpStatusAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ApplianceAppOpStatusAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ApplianceAppOpStatusAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetApiStatuses

`func (o *ApplianceAppOpStatusAllOf) GetApiStatuses() []ApplianceApiStatus`

GetApiStatuses returns the ApiStatuses field if non-nil, zero value otherwise.

### GetApiStatusesOk

`func (o *ApplianceAppOpStatusAllOf) GetApiStatusesOk() (*[]ApplianceApiStatus, bool)`

GetApiStatusesOk returns a tuple with the ApiStatuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiStatuses

`func (o *ApplianceAppOpStatusAllOf) SetApiStatuses(v []ApplianceApiStatus)`

SetApiStatuses sets ApiStatuses field to given value.

### HasApiStatuses

`func (o *ApplianceAppOpStatusAllOf) HasApiStatuses() bool`

HasApiStatuses returns a boolean if a field has been set.

### SetApiStatusesNil

`func (o *ApplianceAppOpStatusAllOf) SetApiStatusesNil(b bool)`

 SetApiStatusesNil sets the value for ApiStatuses to be an explicit nil

### UnsetApiStatuses
`func (o *ApplianceAppOpStatusAllOf) UnsetApiStatuses()`

UnsetApiStatuses ensures that no value is present for ApiStatuses, not even an explicit nil
### GetAppLabel

`func (o *ApplianceAppOpStatusAllOf) GetAppLabel() string`

GetAppLabel returns the AppLabel field if non-nil, zero value otherwise.

### GetAppLabelOk

`func (o *ApplianceAppOpStatusAllOf) GetAppLabelOk() (*string, bool)`

GetAppLabelOk returns a tuple with the AppLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppLabel

`func (o *ApplianceAppOpStatusAllOf) SetAppLabel(v string)`

SetAppLabel sets AppLabel field to given value.

### HasAppLabel

`func (o *ApplianceAppOpStatusAllOf) HasAppLabel() bool`

HasAppLabel returns a boolean if a field has been set.

### GetOperationalStatus

`func (o *ApplianceAppOpStatusAllOf) GetOperationalStatus() string`

GetOperationalStatus returns the OperationalStatus field if non-nil, zero value otherwise.

### GetOperationalStatusOk

`func (o *ApplianceAppOpStatusAllOf) GetOperationalStatusOk() (*string, bool)`

GetOperationalStatusOk returns a tuple with the OperationalStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationalStatus

`func (o *ApplianceAppOpStatusAllOf) SetOperationalStatus(v string)`

SetOperationalStatus sets OperationalStatus field to given value.

### HasOperationalStatus

`func (o *ApplianceAppOpStatusAllOf) HasOperationalStatus() bool`

HasOperationalStatus returns a boolean if a field has been set.

### GetReadyCount

`func (o *ApplianceAppOpStatusAllOf) GetReadyCount() int64`

GetReadyCount returns the ReadyCount field if non-nil, zero value otherwise.

### GetReadyCountOk

`func (o *ApplianceAppOpStatusAllOf) GetReadyCountOk() (*int64, bool)`

GetReadyCountOk returns a tuple with the ReadyCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadyCount

`func (o *ApplianceAppOpStatusAllOf) SetReadyCount(v int64)`

SetReadyCount sets ReadyCount field to given value.

### HasReadyCount

`func (o *ApplianceAppOpStatusAllOf) HasReadyCount() bool`

HasReadyCount returns a boolean if a field has been set.

### GetReplicaCount

`func (o *ApplianceAppOpStatusAllOf) GetReplicaCount() int64`

GetReplicaCount returns the ReplicaCount field if non-nil, zero value otherwise.

### GetReplicaCountOk

`func (o *ApplianceAppOpStatusAllOf) GetReplicaCountOk() (*int64, bool)`

GetReplicaCountOk returns a tuple with the ReplicaCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicaCount

`func (o *ApplianceAppOpStatusAllOf) SetReplicaCount(v int64)`

SetReplicaCount sets ReplicaCount field to given value.

### HasReplicaCount

`func (o *ApplianceAppOpStatusAllOf) HasReplicaCount() bool`

HasReplicaCount returns a boolean if a field has been set.

### GetRestartCount1Hour

`func (o *ApplianceAppOpStatusAllOf) GetRestartCount1Hour() int64`

GetRestartCount1Hour returns the RestartCount1Hour field if non-nil, zero value otherwise.

### GetRestartCount1HourOk

`func (o *ApplianceAppOpStatusAllOf) GetRestartCount1HourOk() (*int64, bool)`

GetRestartCount1HourOk returns a tuple with the RestartCount1Hour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestartCount1Hour

`func (o *ApplianceAppOpStatusAllOf) SetRestartCount1Hour(v int64)`

SetRestartCount1Hour sets RestartCount1Hour field to given value.

### HasRestartCount1Hour

`func (o *ApplianceAppOpStatusAllOf) HasRestartCount1Hour() bool`

HasRestartCount1Hour returns a boolean if a field has been set.

### GetRestartCount24Hours

`func (o *ApplianceAppOpStatusAllOf) GetRestartCount24Hours() int64`

GetRestartCount24Hours returns the RestartCount24Hours field if non-nil, zero value otherwise.

### GetRestartCount24HoursOk

`func (o *ApplianceAppOpStatusAllOf) GetRestartCount24HoursOk() (*int64, bool)`

GetRestartCount24HoursOk returns a tuple with the RestartCount24Hours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestartCount24Hours

`func (o *ApplianceAppOpStatusAllOf) SetRestartCount24Hours(v int64)`

SetRestartCount24Hours sets RestartCount24Hours field to given value.

### HasRestartCount24Hours

`func (o *ApplianceAppOpStatusAllOf) HasRestartCount24Hours() bool`

HasRestartCount24Hours returns a boolean if a field has been set.

### GetRestartCount5Mins

`func (o *ApplianceAppOpStatusAllOf) GetRestartCount5Mins() int64`

GetRestartCount5Mins returns the RestartCount5Mins field if non-nil, zero value otherwise.

### GetRestartCount5MinsOk

`func (o *ApplianceAppOpStatusAllOf) GetRestartCount5MinsOk() (*int64, bool)`

GetRestartCount5MinsOk returns a tuple with the RestartCount5Mins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestartCount5Mins

`func (o *ApplianceAppOpStatusAllOf) SetRestartCount5Mins(v int64)`

SetRestartCount5Mins sets RestartCount5Mins field to given value.

### HasRestartCount5Mins

`func (o *ApplianceAppOpStatusAllOf) HasRestartCount5Mins() bool`

HasRestartCount5Mins returns a boolean if a field has been set.

### GetRestartCountTotal

`func (o *ApplianceAppOpStatusAllOf) GetRestartCountTotal() int64`

GetRestartCountTotal returns the RestartCountTotal field if non-nil, zero value otherwise.

### GetRestartCountTotalOk

`func (o *ApplianceAppOpStatusAllOf) GetRestartCountTotalOk() (*int64, bool)`

GetRestartCountTotalOk returns a tuple with the RestartCountTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestartCountTotal

`func (o *ApplianceAppOpStatusAllOf) SetRestartCountTotal(v int64)`

SetRestartCountTotal sets RestartCountTotal field to given value.

### HasRestartCountTotal

`func (o *ApplianceAppOpStatusAllOf) HasRestartCountTotal() bool`

HasRestartCountTotal returns a boolean if a field has been set.

### GetRunningCount

`func (o *ApplianceAppOpStatusAllOf) GetRunningCount() int64`

GetRunningCount returns the RunningCount field if non-nil, zero value otherwise.

### GetRunningCountOk

`func (o *ApplianceAppOpStatusAllOf) GetRunningCountOk() (*int64, bool)`

GetRunningCountOk returns a tuple with the RunningCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunningCount

`func (o *ApplianceAppOpStatusAllOf) SetRunningCount(v int64)`

SetRunningCount sets RunningCount field to given value.

### HasRunningCount

`func (o *ApplianceAppOpStatusAllOf) HasRunningCount() bool`

HasRunningCount returns a boolean if a field has been set.

### GetGroupOpStatus

`func (o *ApplianceAppOpStatusAllOf) GetGroupOpStatus() ApplianceGroupOpStatusRelationship`

GetGroupOpStatus returns the GroupOpStatus field if non-nil, zero value otherwise.

### GetGroupOpStatusOk

`func (o *ApplianceAppOpStatusAllOf) GetGroupOpStatusOk() (*ApplianceGroupOpStatusRelationship, bool)`

GetGroupOpStatusOk returns a tuple with the GroupOpStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupOpStatus

`func (o *ApplianceAppOpStatusAllOf) SetGroupOpStatus(v ApplianceGroupOpStatusRelationship)`

SetGroupOpStatus sets GroupOpStatus field to given value.

### HasGroupOpStatus

`func (o *ApplianceAppOpStatusAllOf) HasGroupOpStatus() bool`

HasGroupOpStatus returns a boolean if a field has been set.

### GetSystemOpStatus

`func (o *ApplianceAppOpStatusAllOf) GetSystemOpStatus() ApplianceSystemOpStatusRelationship`

GetSystemOpStatus returns the SystemOpStatus field if non-nil, zero value otherwise.

### GetSystemOpStatusOk

`func (o *ApplianceAppOpStatusAllOf) GetSystemOpStatusOk() (*ApplianceSystemOpStatusRelationship, bool)`

GetSystemOpStatusOk returns a tuple with the SystemOpStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemOpStatus

`func (o *ApplianceAppOpStatusAllOf) SetSystemOpStatus(v ApplianceSystemOpStatusRelationship)`

SetSystemOpStatus sets SystemOpStatus field to given value.

### HasSystemOpStatus

`func (o *ApplianceAppOpStatusAllOf) HasSystemOpStatus() bool`

HasSystemOpStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


