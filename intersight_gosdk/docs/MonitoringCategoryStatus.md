# MonitoringCategoryStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "monitoring.CategoryStatus"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "monitoring.CategoryStatus"]
**CategoryLabel** | Pointer to **string** | Name of the category for which status is being reported (for e.g. &#39;Licensing&#39;, &#39;Advisories&#39;). | [optional] 
**Deeplink** | Pointer to **string** | Link to the corresponding category specific page in Intersight to get additional information and troubleshoot. for e.g. &#39;Alarms&#39; category would have the deeplink as &#39;https://intersight.com/an/cond/alarms/active&#39;. | [optional] 
**Details** | Pointer to **string** | Additional information regarding category status that may be displayed in UI related to the category status. Optional and currently unused. | [optional] 
**SourceId** | Pointer to **string** | Additional parameter to be used for traceability and troubleshooting, currently unused. | [optional] 
**Status** | Pointer to **string** | Aggregated status of the category being reported. For e.g., a given Intersight account might have a combination of high and low severity Advisories applicable to managed devices. However, even if one of the devices is impacted by a high severity Advisories, the category status is reported as &#39;critical&#39;. * &#x60;Unknown&#x60; - The current status for the respective category cannot be determined. This may be due to some intermittent issue or due to the fact that the user making the request does not have appropriate privileges/entitlements for the concerned category. * &#x60;Critical&#x60; - The Intersight account is impacted by a high severity issue  for the applicable category that will need to be addressed immediately. for e.g. critical status for &#39;Advisories&#39; category would mean that the Intersight account is impacted by an urgent field notice or a vulnerability with high severity level. * &#x60;Warning&#x60; - The Intersight account is impacted by a high severity issue  for the applicable category that will need to be addressed soon. for e.g. warning status for &#39;Advisories&#39; category might mean that the Intersight account is impacted by an a vulnerability with moderate severity level. * &#x60;Healthy&#x60; - The Intersight account is not impacted by any high severity issue for the applicable category. | [optional] [default to "Unknown"]

## Methods

### NewMonitoringCategoryStatus

`func NewMonitoringCategoryStatus(classId string, objectType string, ) *MonitoringCategoryStatus`

NewMonitoringCategoryStatus instantiates a new MonitoringCategoryStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitoringCategoryStatusWithDefaults

`func NewMonitoringCategoryStatusWithDefaults() *MonitoringCategoryStatus`

NewMonitoringCategoryStatusWithDefaults instantiates a new MonitoringCategoryStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *MonitoringCategoryStatus) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *MonitoringCategoryStatus) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *MonitoringCategoryStatus) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *MonitoringCategoryStatus) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *MonitoringCategoryStatus) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *MonitoringCategoryStatus) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCategoryLabel

`func (o *MonitoringCategoryStatus) GetCategoryLabel() string`

GetCategoryLabel returns the CategoryLabel field if non-nil, zero value otherwise.

### GetCategoryLabelOk

`func (o *MonitoringCategoryStatus) GetCategoryLabelOk() (*string, bool)`

GetCategoryLabelOk returns a tuple with the CategoryLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryLabel

`func (o *MonitoringCategoryStatus) SetCategoryLabel(v string)`

SetCategoryLabel sets CategoryLabel field to given value.

### HasCategoryLabel

`func (o *MonitoringCategoryStatus) HasCategoryLabel() bool`

HasCategoryLabel returns a boolean if a field has been set.

### GetDeeplink

`func (o *MonitoringCategoryStatus) GetDeeplink() string`

GetDeeplink returns the Deeplink field if non-nil, zero value otherwise.

### GetDeeplinkOk

`func (o *MonitoringCategoryStatus) GetDeeplinkOk() (*string, bool)`

GetDeeplinkOk returns a tuple with the Deeplink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeeplink

`func (o *MonitoringCategoryStatus) SetDeeplink(v string)`

SetDeeplink sets Deeplink field to given value.

### HasDeeplink

`func (o *MonitoringCategoryStatus) HasDeeplink() bool`

HasDeeplink returns a boolean if a field has been set.

### GetDetails

`func (o *MonitoringCategoryStatus) GetDetails() string`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *MonitoringCategoryStatus) GetDetailsOk() (*string, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *MonitoringCategoryStatus) SetDetails(v string)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *MonitoringCategoryStatus) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetSourceId

`func (o *MonitoringCategoryStatus) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *MonitoringCategoryStatus) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *MonitoringCategoryStatus) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.

### HasSourceId

`func (o *MonitoringCategoryStatus) HasSourceId() bool`

HasSourceId returns a boolean if a field has been set.

### GetStatus

`func (o *MonitoringCategoryStatus) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MonitoringCategoryStatus) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MonitoringCategoryStatus) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MonitoringCategoryStatus) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


