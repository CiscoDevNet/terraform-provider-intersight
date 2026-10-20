# NiatelemetryAnomaly

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "niatelemetry.Anomaly"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "niatelemetry.Anomaly"]
**AlertTypeId** | Pointer to **string** | Identifier of the alert type that triggered this notification. | [optional] [readonly] 
**Category** | Pointer to **string** | Classification of the anomaly. | [optional] [readonly] 
**DeepLink** | Pointer to **string** | Deep link to investigate the anomaly in the source system. | [optional] [readonly] 
**Description** | Pointer to **string** | Concise explanation of the anomaly. | [optional] [readonly] 
**Details** | Pointer to **string** | Additional contextual details for triage. | [optional] [readonly] 
**EntityId** | Pointer to **string** | Identifier of the impacted entity in the source system. | [optional] [readonly] 
**EntityType** | Pointer to **string** | Impacted entity type information derived from anomaly objects. | [optional] [readonly] 
**EventTime** | Pointer to **string** | Time when the event occurred in the source system. | [optional] [readonly] 
**FabricName** | Pointer to **string** | The name of the fabric that generated the anomaly. | [optional] [readonly] 
**FirstRaisedTime** | Pointer to **string** | First time the anomaly was raised. | [optional] [readonly] 
**ExtId** | Pointer to **string** | Globally unique identifier for the anomaly notification. | [optional] [readonly] 
**LastRaisedTime** | Pointer to **string** | Most recent transition to open or reopened state. | [optional] [readonly] 
**LastUpdatedTime** | Pointer to **string** | Most recent lifecycle change time. | [optional] [readonly] 
**NotificationType** | Pointer to **string** | Type of notification, such as anomaly or advisory. | [optional] [readonly] 
**Product** | Pointer to **string** | Product that generated the notification. | [optional] [readonly] 
**Region** | Pointer to **string** | Region for data sovereignty. | [optional] [readonly] 
**Severity** | Pointer to **string** | Normalized severity level. | [optional] [readonly] 
**Title** | Pointer to **string** | Short human-readable summary of the notification. | [optional] [readonly] 
**Fabric** | Pointer to [**NullableNiatelemetryFabricRelationship**](NiatelemetryFabricRelationship.md) |  | [optional] 

## Methods

### NewNiatelemetryAnomaly

`func NewNiatelemetryAnomaly(classId string, objectType string, ) *NiatelemetryAnomaly`

NewNiatelemetryAnomaly instantiates a new NiatelemetryAnomaly object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiatelemetryAnomalyWithDefaults

`func NewNiatelemetryAnomalyWithDefaults() *NiatelemetryAnomaly`

NewNiatelemetryAnomalyWithDefaults instantiates a new NiatelemetryAnomaly object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NiatelemetryAnomaly) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NiatelemetryAnomaly) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NiatelemetryAnomaly) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NiatelemetryAnomaly) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NiatelemetryAnomaly) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NiatelemetryAnomaly) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAlertTypeId

`func (o *NiatelemetryAnomaly) GetAlertTypeId() string`

GetAlertTypeId returns the AlertTypeId field if non-nil, zero value otherwise.

### GetAlertTypeIdOk

`func (o *NiatelemetryAnomaly) GetAlertTypeIdOk() (*string, bool)`

GetAlertTypeIdOk returns a tuple with the AlertTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertTypeId

`func (o *NiatelemetryAnomaly) SetAlertTypeId(v string)`

SetAlertTypeId sets AlertTypeId field to given value.

### HasAlertTypeId

`func (o *NiatelemetryAnomaly) HasAlertTypeId() bool`

HasAlertTypeId returns a boolean if a field has been set.

### GetCategory

`func (o *NiatelemetryAnomaly) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *NiatelemetryAnomaly) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *NiatelemetryAnomaly) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *NiatelemetryAnomaly) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetDeepLink

`func (o *NiatelemetryAnomaly) GetDeepLink() string`

GetDeepLink returns the DeepLink field if non-nil, zero value otherwise.

### GetDeepLinkOk

`func (o *NiatelemetryAnomaly) GetDeepLinkOk() (*string, bool)`

GetDeepLinkOk returns a tuple with the DeepLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeepLink

`func (o *NiatelemetryAnomaly) SetDeepLink(v string)`

SetDeepLink sets DeepLink field to given value.

### HasDeepLink

`func (o *NiatelemetryAnomaly) HasDeepLink() bool`

HasDeepLink returns a boolean if a field has been set.

### GetDescription

`func (o *NiatelemetryAnomaly) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NiatelemetryAnomaly) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *NiatelemetryAnomaly) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *NiatelemetryAnomaly) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDetails

`func (o *NiatelemetryAnomaly) GetDetails() string`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *NiatelemetryAnomaly) GetDetailsOk() (*string, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *NiatelemetryAnomaly) SetDetails(v string)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *NiatelemetryAnomaly) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetEntityId

`func (o *NiatelemetryAnomaly) GetEntityId() string`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *NiatelemetryAnomaly) GetEntityIdOk() (*string, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *NiatelemetryAnomaly) SetEntityId(v string)`

SetEntityId sets EntityId field to given value.

### HasEntityId

`func (o *NiatelemetryAnomaly) HasEntityId() bool`

HasEntityId returns a boolean if a field has been set.

### GetEntityType

`func (o *NiatelemetryAnomaly) GetEntityType() string`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *NiatelemetryAnomaly) GetEntityTypeOk() (*string, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *NiatelemetryAnomaly) SetEntityType(v string)`

SetEntityType sets EntityType field to given value.

### HasEntityType

`func (o *NiatelemetryAnomaly) HasEntityType() bool`

HasEntityType returns a boolean if a field has been set.

### GetEventTime

`func (o *NiatelemetryAnomaly) GetEventTime() string`

GetEventTime returns the EventTime field if non-nil, zero value otherwise.

### GetEventTimeOk

`func (o *NiatelemetryAnomaly) GetEventTimeOk() (*string, bool)`

GetEventTimeOk returns a tuple with the EventTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTime

`func (o *NiatelemetryAnomaly) SetEventTime(v string)`

SetEventTime sets EventTime field to given value.

### HasEventTime

`func (o *NiatelemetryAnomaly) HasEventTime() bool`

HasEventTime returns a boolean if a field has been set.

### GetFabricName

`func (o *NiatelemetryAnomaly) GetFabricName() string`

GetFabricName returns the FabricName field if non-nil, zero value otherwise.

### GetFabricNameOk

`func (o *NiatelemetryAnomaly) GetFabricNameOk() (*string, bool)`

GetFabricNameOk returns a tuple with the FabricName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFabricName

`func (o *NiatelemetryAnomaly) SetFabricName(v string)`

SetFabricName sets FabricName field to given value.

### HasFabricName

`func (o *NiatelemetryAnomaly) HasFabricName() bool`

HasFabricName returns a boolean if a field has been set.

### GetFirstRaisedTime

`func (o *NiatelemetryAnomaly) GetFirstRaisedTime() string`

GetFirstRaisedTime returns the FirstRaisedTime field if non-nil, zero value otherwise.

### GetFirstRaisedTimeOk

`func (o *NiatelemetryAnomaly) GetFirstRaisedTimeOk() (*string, bool)`

GetFirstRaisedTimeOk returns a tuple with the FirstRaisedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstRaisedTime

`func (o *NiatelemetryAnomaly) SetFirstRaisedTime(v string)`

SetFirstRaisedTime sets FirstRaisedTime field to given value.

### HasFirstRaisedTime

`func (o *NiatelemetryAnomaly) HasFirstRaisedTime() bool`

HasFirstRaisedTime returns a boolean if a field has been set.

### GetExtId

`func (o *NiatelemetryAnomaly) GetExtId() string`

GetExtId returns the ExtId field if non-nil, zero value otherwise.

### GetExtIdOk

`func (o *NiatelemetryAnomaly) GetExtIdOk() (*string, bool)`

GetExtIdOk returns a tuple with the ExtId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtId

`func (o *NiatelemetryAnomaly) SetExtId(v string)`

SetExtId sets ExtId field to given value.

### HasExtId

`func (o *NiatelemetryAnomaly) HasExtId() bool`

HasExtId returns a boolean if a field has been set.

### GetLastRaisedTime

`func (o *NiatelemetryAnomaly) GetLastRaisedTime() string`

GetLastRaisedTime returns the LastRaisedTime field if non-nil, zero value otherwise.

### GetLastRaisedTimeOk

`func (o *NiatelemetryAnomaly) GetLastRaisedTimeOk() (*string, bool)`

GetLastRaisedTimeOk returns a tuple with the LastRaisedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRaisedTime

`func (o *NiatelemetryAnomaly) SetLastRaisedTime(v string)`

SetLastRaisedTime sets LastRaisedTime field to given value.

### HasLastRaisedTime

`func (o *NiatelemetryAnomaly) HasLastRaisedTime() bool`

HasLastRaisedTime returns a boolean if a field has been set.

### GetLastUpdatedTime

`func (o *NiatelemetryAnomaly) GetLastUpdatedTime() string`

GetLastUpdatedTime returns the LastUpdatedTime field if non-nil, zero value otherwise.

### GetLastUpdatedTimeOk

`func (o *NiatelemetryAnomaly) GetLastUpdatedTimeOk() (*string, bool)`

GetLastUpdatedTimeOk returns a tuple with the LastUpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedTime

`func (o *NiatelemetryAnomaly) SetLastUpdatedTime(v string)`

SetLastUpdatedTime sets LastUpdatedTime field to given value.

### HasLastUpdatedTime

`func (o *NiatelemetryAnomaly) HasLastUpdatedTime() bool`

HasLastUpdatedTime returns a boolean if a field has been set.

### GetNotificationType

`func (o *NiatelemetryAnomaly) GetNotificationType() string`

GetNotificationType returns the NotificationType field if non-nil, zero value otherwise.

### GetNotificationTypeOk

`func (o *NiatelemetryAnomaly) GetNotificationTypeOk() (*string, bool)`

GetNotificationTypeOk returns a tuple with the NotificationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationType

`func (o *NiatelemetryAnomaly) SetNotificationType(v string)`

SetNotificationType sets NotificationType field to given value.

### HasNotificationType

`func (o *NiatelemetryAnomaly) HasNotificationType() bool`

HasNotificationType returns a boolean if a field has been set.

### GetProduct

`func (o *NiatelemetryAnomaly) GetProduct() string`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *NiatelemetryAnomaly) GetProductOk() (*string, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *NiatelemetryAnomaly) SetProduct(v string)`

SetProduct sets Product field to given value.

### HasProduct

`func (o *NiatelemetryAnomaly) HasProduct() bool`

HasProduct returns a boolean if a field has been set.

### GetRegion

`func (o *NiatelemetryAnomaly) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *NiatelemetryAnomaly) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *NiatelemetryAnomaly) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *NiatelemetryAnomaly) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetSeverity

`func (o *NiatelemetryAnomaly) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *NiatelemetryAnomaly) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *NiatelemetryAnomaly) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *NiatelemetryAnomaly) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetTitle

`func (o *NiatelemetryAnomaly) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *NiatelemetryAnomaly) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *NiatelemetryAnomaly) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *NiatelemetryAnomaly) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetFabric

`func (o *NiatelemetryAnomaly) GetFabric() NiatelemetryFabricRelationship`

GetFabric returns the Fabric field if non-nil, zero value otherwise.

### GetFabricOk

`func (o *NiatelemetryAnomaly) GetFabricOk() (*NiatelemetryFabricRelationship, bool)`

GetFabricOk returns a tuple with the Fabric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFabric

`func (o *NiatelemetryAnomaly) SetFabric(v NiatelemetryFabricRelationship)`

SetFabric sets Fabric field to given value.

### HasFabric

`func (o *NiatelemetryAnomaly) HasFabric() bool`

HasFabric returns a boolean if a field has been set.

### SetFabricNil

`func (o *NiatelemetryAnomaly) SetFabricNil(b bool)`

 SetFabricNil sets the value for Fabric to be an explicit nil

### UnsetFabric
`func (o *NiatelemetryAnomaly) UnsetFabric()`

UnsetFabric ensures that no value is present for Fabric, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


