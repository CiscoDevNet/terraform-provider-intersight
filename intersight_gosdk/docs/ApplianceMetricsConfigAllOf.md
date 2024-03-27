# ApplianceMetricsConfigAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "appliance.MetricsConfig"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "appliance.MetricsConfig"]
**CurrentEndpointCount** | Pointer to **int64** | Number of discovered endpoints from where metrics is being collected currently. | [optional] [readonly] [default to 0]
**EndpointUsagePercent** | Pointer to **int64** | Usage percentage of the discovered endpoints. | [optional] [readonly] [default to 0]
**LastDisabledDate** | Pointer to **time.Time** | Disabled date of the metrics collection feature. | [optional] [readonly] 
**LastEnabledDate** | Pointer to **time.Time** | Enabled date of the metrics collection feature. | [optional] [readonly] 
**MaxEndpointCount** | Pointer to **int64** | The maximum number of supported endpoints for an appliance deployment type. | [optional] [readonly] [default to 0]
**StatusMessage** | Pointer to **string** | The overall metrics collection Status based on resource constraints. | [optional] [readonly] 
**SystemEnabled** | Pointer to **bool** | Metric collection state defined by the system. | [optional] [readonly] [default to false]
**UserEnabled** | Pointer to **bool** | Configured metric collection state by the account administrator. | [optional] [default to false]
**Account** | Pointer to [**IamAccountRelationship**](IamAccountRelationship.md) |  | [optional] 

## Methods

### NewApplianceMetricsConfigAllOf

`func NewApplianceMetricsConfigAllOf(classId string, objectType string, ) *ApplianceMetricsConfigAllOf`

NewApplianceMetricsConfigAllOf instantiates a new ApplianceMetricsConfigAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplianceMetricsConfigAllOfWithDefaults

`func NewApplianceMetricsConfigAllOfWithDefaults() *ApplianceMetricsConfigAllOf`

NewApplianceMetricsConfigAllOfWithDefaults instantiates a new ApplianceMetricsConfigAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ApplianceMetricsConfigAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ApplianceMetricsConfigAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ApplianceMetricsConfigAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ApplianceMetricsConfigAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ApplianceMetricsConfigAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ApplianceMetricsConfigAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCurrentEndpointCount

`func (o *ApplianceMetricsConfigAllOf) GetCurrentEndpointCount() int64`

GetCurrentEndpointCount returns the CurrentEndpointCount field if non-nil, zero value otherwise.

### GetCurrentEndpointCountOk

`func (o *ApplianceMetricsConfigAllOf) GetCurrentEndpointCountOk() (*int64, bool)`

GetCurrentEndpointCountOk returns a tuple with the CurrentEndpointCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentEndpointCount

`func (o *ApplianceMetricsConfigAllOf) SetCurrentEndpointCount(v int64)`

SetCurrentEndpointCount sets CurrentEndpointCount field to given value.

### HasCurrentEndpointCount

`func (o *ApplianceMetricsConfigAllOf) HasCurrentEndpointCount() bool`

HasCurrentEndpointCount returns a boolean if a field has been set.

### GetEndpointUsagePercent

`func (o *ApplianceMetricsConfigAllOf) GetEndpointUsagePercent() int64`

GetEndpointUsagePercent returns the EndpointUsagePercent field if non-nil, zero value otherwise.

### GetEndpointUsagePercentOk

`func (o *ApplianceMetricsConfigAllOf) GetEndpointUsagePercentOk() (*int64, bool)`

GetEndpointUsagePercentOk returns a tuple with the EndpointUsagePercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointUsagePercent

`func (o *ApplianceMetricsConfigAllOf) SetEndpointUsagePercent(v int64)`

SetEndpointUsagePercent sets EndpointUsagePercent field to given value.

### HasEndpointUsagePercent

`func (o *ApplianceMetricsConfigAllOf) HasEndpointUsagePercent() bool`

HasEndpointUsagePercent returns a boolean if a field has been set.

### GetLastDisabledDate

`func (o *ApplianceMetricsConfigAllOf) GetLastDisabledDate() time.Time`

GetLastDisabledDate returns the LastDisabledDate field if non-nil, zero value otherwise.

### GetLastDisabledDateOk

`func (o *ApplianceMetricsConfigAllOf) GetLastDisabledDateOk() (*time.Time, bool)`

GetLastDisabledDateOk returns a tuple with the LastDisabledDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastDisabledDate

`func (o *ApplianceMetricsConfigAllOf) SetLastDisabledDate(v time.Time)`

SetLastDisabledDate sets LastDisabledDate field to given value.

### HasLastDisabledDate

`func (o *ApplianceMetricsConfigAllOf) HasLastDisabledDate() bool`

HasLastDisabledDate returns a boolean if a field has been set.

### GetLastEnabledDate

`func (o *ApplianceMetricsConfigAllOf) GetLastEnabledDate() time.Time`

GetLastEnabledDate returns the LastEnabledDate field if non-nil, zero value otherwise.

### GetLastEnabledDateOk

`func (o *ApplianceMetricsConfigAllOf) GetLastEnabledDateOk() (*time.Time, bool)`

GetLastEnabledDateOk returns a tuple with the LastEnabledDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastEnabledDate

`func (o *ApplianceMetricsConfigAllOf) SetLastEnabledDate(v time.Time)`

SetLastEnabledDate sets LastEnabledDate field to given value.

### HasLastEnabledDate

`func (o *ApplianceMetricsConfigAllOf) HasLastEnabledDate() bool`

HasLastEnabledDate returns a boolean if a field has been set.

### GetMaxEndpointCount

`func (o *ApplianceMetricsConfigAllOf) GetMaxEndpointCount() int64`

GetMaxEndpointCount returns the MaxEndpointCount field if non-nil, zero value otherwise.

### GetMaxEndpointCountOk

`func (o *ApplianceMetricsConfigAllOf) GetMaxEndpointCountOk() (*int64, bool)`

GetMaxEndpointCountOk returns a tuple with the MaxEndpointCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxEndpointCount

`func (o *ApplianceMetricsConfigAllOf) SetMaxEndpointCount(v int64)`

SetMaxEndpointCount sets MaxEndpointCount field to given value.

### HasMaxEndpointCount

`func (o *ApplianceMetricsConfigAllOf) HasMaxEndpointCount() bool`

HasMaxEndpointCount returns a boolean if a field has been set.

### GetStatusMessage

`func (o *ApplianceMetricsConfigAllOf) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *ApplianceMetricsConfigAllOf) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *ApplianceMetricsConfigAllOf) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *ApplianceMetricsConfigAllOf) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### GetSystemEnabled

`func (o *ApplianceMetricsConfigAllOf) GetSystemEnabled() bool`

GetSystemEnabled returns the SystemEnabled field if non-nil, zero value otherwise.

### GetSystemEnabledOk

`func (o *ApplianceMetricsConfigAllOf) GetSystemEnabledOk() (*bool, bool)`

GetSystemEnabledOk returns a tuple with the SystemEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemEnabled

`func (o *ApplianceMetricsConfigAllOf) SetSystemEnabled(v bool)`

SetSystemEnabled sets SystemEnabled field to given value.

### HasSystemEnabled

`func (o *ApplianceMetricsConfigAllOf) HasSystemEnabled() bool`

HasSystemEnabled returns a boolean if a field has been set.

### GetUserEnabled

`func (o *ApplianceMetricsConfigAllOf) GetUserEnabled() bool`

GetUserEnabled returns the UserEnabled field if non-nil, zero value otherwise.

### GetUserEnabledOk

`func (o *ApplianceMetricsConfigAllOf) GetUserEnabledOk() (*bool, bool)`

GetUserEnabledOk returns a tuple with the UserEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserEnabled

`func (o *ApplianceMetricsConfigAllOf) SetUserEnabled(v bool)`

SetUserEnabled sets UserEnabled field to given value.

### HasUserEnabled

`func (o *ApplianceMetricsConfigAllOf) HasUserEnabled() bool`

HasUserEnabled returns a boolean if a field has been set.

### GetAccount

`func (o *ApplianceMetricsConfigAllOf) GetAccount() IamAccountRelationship`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *ApplianceMetricsConfigAllOf) GetAccountOk() (*IamAccountRelationship, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *ApplianceMetricsConfigAllOf) SetAccount(v IamAccountRelationship)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *ApplianceMetricsConfigAllOf) HasAccount() bool`

HasAccount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


