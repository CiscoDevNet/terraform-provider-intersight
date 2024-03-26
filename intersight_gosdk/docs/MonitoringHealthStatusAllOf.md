# MonitoringHealthStatusAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "monitoring.HealthStatus"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "monitoring.HealthStatus"]
**CategoryStatus** | Pointer to [**[]MonitoringCategoryStatus**](MonitoringCategoryStatus.md) |  | [optional] 
**HealthDataSchemaVersion** | Pointer to **string** | Version of compliant health data API schema. | [optional] [readonly] 
**Source** | Pointer to **string** | Set as &#39;Intersight&#39;. Especially useful in cases such as when this API is consumed by an external dashboard. This field allows such dashboards to aggregate health status across multiple  sources (Intersight, Meraki etc.). | [optional] [readonly] 
**StatusTimeStamp** | Pointer to **time.Time** | Time stamp when the status was generated. The status reported by this API may lag the real time status by up to 5 minutes. | [optional] [readonly] 
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](OrganizationOrganizationRelationship.md) |  | [optional] 

## Methods

### NewMonitoringHealthStatusAllOf

`func NewMonitoringHealthStatusAllOf(classId string, objectType string, ) *MonitoringHealthStatusAllOf`

NewMonitoringHealthStatusAllOf instantiates a new MonitoringHealthStatusAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitoringHealthStatusAllOfWithDefaults

`func NewMonitoringHealthStatusAllOfWithDefaults() *MonitoringHealthStatusAllOf`

NewMonitoringHealthStatusAllOfWithDefaults instantiates a new MonitoringHealthStatusAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *MonitoringHealthStatusAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *MonitoringHealthStatusAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *MonitoringHealthStatusAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *MonitoringHealthStatusAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *MonitoringHealthStatusAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *MonitoringHealthStatusAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCategoryStatus

`func (o *MonitoringHealthStatusAllOf) GetCategoryStatus() []MonitoringCategoryStatus`

GetCategoryStatus returns the CategoryStatus field if non-nil, zero value otherwise.

### GetCategoryStatusOk

`func (o *MonitoringHealthStatusAllOf) GetCategoryStatusOk() (*[]MonitoringCategoryStatus, bool)`

GetCategoryStatusOk returns a tuple with the CategoryStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryStatus

`func (o *MonitoringHealthStatusAllOf) SetCategoryStatus(v []MonitoringCategoryStatus)`

SetCategoryStatus sets CategoryStatus field to given value.

### HasCategoryStatus

`func (o *MonitoringHealthStatusAllOf) HasCategoryStatus() bool`

HasCategoryStatus returns a boolean if a field has been set.

### SetCategoryStatusNil

`func (o *MonitoringHealthStatusAllOf) SetCategoryStatusNil(b bool)`

 SetCategoryStatusNil sets the value for CategoryStatus to be an explicit nil

### UnsetCategoryStatus
`func (o *MonitoringHealthStatusAllOf) UnsetCategoryStatus()`

UnsetCategoryStatus ensures that no value is present for CategoryStatus, not even an explicit nil
### GetHealthDataSchemaVersion

`func (o *MonitoringHealthStatusAllOf) GetHealthDataSchemaVersion() string`

GetHealthDataSchemaVersion returns the HealthDataSchemaVersion field if non-nil, zero value otherwise.

### GetHealthDataSchemaVersionOk

`func (o *MonitoringHealthStatusAllOf) GetHealthDataSchemaVersionOk() (*string, bool)`

GetHealthDataSchemaVersionOk returns a tuple with the HealthDataSchemaVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthDataSchemaVersion

`func (o *MonitoringHealthStatusAllOf) SetHealthDataSchemaVersion(v string)`

SetHealthDataSchemaVersion sets HealthDataSchemaVersion field to given value.

### HasHealthDataSchemaVersion

`func (o *MonitoringHealthStatusAllOf) HasHealthDataSchemaVersion() bool`

HasHealthDataSchemaVersion returns a boolean if a field has been set.

### GetSource

`func (o *MonitoringHealthStatusAllOf) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *MonitoringHealthStatusAllOf) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *MonitoringHealthStatusAllOf) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *MonitoringHealthStatusAllOf) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetStatusTimeStamp

`func (o *MonitoringHealthStatusAllOf) GetStatusTimeStamp() time.Time`

GetStatusTimeStamp returns the StatusTimeStamp field if non-nil, zero value otherwise.

### GetStatusTimeStampOk

`func (o *MonitoringHealthStatusAllOf) GetStatusTimeStampOk() (*time.Time, bool)`

GetStatusTimeStampOk returns a tuple with the StatusTimeStamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusTimeStamp

`func (o *MonitoringHealthStatusAllOf) SetStatusTimeStamp(v time.Time)`

SetStatusTimeStamp sets StatusTimeStamp field to given value.

### HasStatusTimeStamp

`func (o *MonitoringHealthStatusAllOf) HasStatusTimeStamp() bool`

HasStatusTimeStamp returns a boolean if a field has been set.

### GetOrganization

`func (o *MonitoringHealthStatusAllOf) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *MonitoringHealthStatusAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *MonitoringHealthStatusAllOf) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *MonitoringHealthStatusAllOf) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


