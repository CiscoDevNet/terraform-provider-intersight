# ApplianceSystemStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "appliance.SystemStatus"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "appliance.SystemStatus"]
**OperationalStatus** | Pointer to **string** | Operational status of the Intersight Appliance. Operational status is based on the result of the status checks. If result of any check is Critical, then its value is Impaired. Otherwise, if result of any check is Warning, then its value is AttentionNeeded. If all checks are OK, then its value is Operational. * &#x60;Unknown&#x60; - Operational status of the Intersight Appliance entity is Unknown. * &#x60;Operational&#x60; - Operational status of the Intersight Appliance entity is Operational. * &#x60;Impaired&#x60; - Operational status of the Intersight Appliance entity is Impaired. * &#x60;AttentionNeeded&#x60; - Operational status of the Intersight Appliance entity is AttentionNeeded. | [optional] [readonly] [default to "Unknown"]
**StatusChecks** | Pointer to [**[]ApplianceStatusCheck**](ApplianceStatusCheck.md) |  | [optional] 
**AppStatuses** | Pointer to [**[]ApplianceAppStatusRelationship**](ApplianceAppStatusRelationship.md) | An array of relationships to applianceAppStatus resources. | [optional] [readonly] 
**GroupStatuses** | Pointer to [**[]ApplianceGroupStatusRelationship**](ApplianceGroupStatusRelationship.md) | An array of relationships to applianceGroupStatus resources. | [optional] [readonly] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 
**SystemInfo** | Pointer to [**ApplianceSystemInfoRelationship**](ApplianceSystemInfoRelationship.md) |  | [optional] 

## Methods

### NewApplianceSystemStatus

`func NewApplianceSystemStatus(classId string, objectType string, ) *ApplianceSystemStatus`

NewApplianceSystemStatus instantiates a new ApplianceSystemStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplianceSystemStatusWithDefaults

`func NewApplianceSystemStatusWithDefaults() *ApplianceSystemStatus`

NewApplianceSystemStatusWithDefaults instantiates a new ApplianceSystemStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ApplianceSystemStatus) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ApplianceSystemStatus) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ApplianceSystemStatus) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ApplianceSystemStatus) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ApplianceSystemStatus) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ApplianceSystemStatus) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetOperationalStatus

`func (o *ApplianceSystemStatus) GetOperationalStatus() string`

GetOperationalStatus returns the OperationalStatus field if non-nil, zero value otherwise.

### GetOperationalStatusOk

`func (o *ApplianceSystemStatus) GetOperationalStatusOk() (*string, bool)`

GetOperationalStatusOk returns a tuple with the OperationalStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationalStatus

`func (o *ApplianceSystemStatus) SetOperationalStatus(v string)`

SetOperationalStatus sets OperationalStatus field to given value.

### HasOperationalStatus

`func (o *ApplianceSystemStatus) HasOperationalStatus() bool`

HasOperationalStatus returns a boolean if a field has been set.

### GetStatusChecks

`func (o *ApplianceSystemStatus) GetStatusChecks() []ApplianceStatusCheck`

GetStatusChecks returns the StatusChecks field if non-nil, zero value otherwise.

### GetStatusChecksOk

`func (o *ApplianceSystemStatus) GetStatusChecksOk() (*[]ApplianceStatusCheck, bool)`

GetStatusChecksOk returns a tuple with the StatusChecks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusChecks

`func (o *ApplianceSystemStatus) SetStatusChecks(v []ApplianceStatusCheck)`

SetStatusChecks sets StatusChecks field to given value.

### HasStatusChecks

`func (o *ApplianceSystemStatus) HasStatusChecks() bool`

HasStatusChecks returns a boolean if a field has been set.

### SetStatusChecksNil

`func (o *ApplianceSystemStatus) SetStatusChecksNil(b bool)`

 SetStatusChecksNil sets the value for StatusChecks to be an explicit nil

### UnsetStatusChecks
`func (o *ApplianceSystemStatus) UnsetStatusChecks()`

UnsetStatusChecks ensures that no value is present for StatusChecks, not even an explicit nil
### GetAppStatuses

`func (o *ApplianceSystemStatus) GetAppStatuses() []ApplianceAppStatusRelationship`

GetAppStatuses returns the AppStatuses field if non-nil, zero value otherwise.

### GetAppStatusesOk

`func (o *ApplianceSystemStatus) GetAppStatusesOk() (*[]ApplianceAppStatusRelationship, bool)`

GetAppStatusesOk returns a tuple with the AppStatuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppStatuses

`func (o *ApplianceSystemStatus) SetAppStatuses(v []ApplianceAppStatusRelationship)`

SetAppStatuses sets AppStatuses field to given value.

### HasAppStatuses

`func (o *ApplianceSystemStatus) HasAppStatuses() bool`

HasAppStatuses returns a boolean if a field has been set.

### SetAppStatusesNil

`func (o *ApplianceSystemStatus) SetAppStatusesNil(b bool)`

 SetAppStatusesNil sets the value for AppStatuses to be an explicit nil

### UnsetAppStatuses
`func (o *ApplianceSystemStatus) UnsetAppStatuses()`

UnsetAppStatuses ensures that no value is present for AppStatuses, not even an explicit nil
### GetGroupStatuses

`func (o *ApplianceSystemStatus) GetGroupStatuses() []ApplianceGroupStatusRelationship`

GetGroupStatuses returns the GroupStatuses field if non-nil, zero value otherwise.

### GetGroupStatusesOk

`func (o *ApplianceSystemStatus) GetGroupStatusesOk() (*[]ApplianceGroupStatusRelationship, bool)`

GetGroupStatusesOk returns a tuple with the GroupStatuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupStatuses

`func (o *ApplianceSystemStatus) SetGroupStatuses(v []ApplianceGroupStatusRelationship)`

SetGroupStatuses sets GroupStatuses field to given value.

### HasGroupStatuses

`func (o *ApplianceSystemStatus) HasGroupStatuses() bool`

HasGroupStatuses returns a boolean if a field has been set.

### SetGroupStatusesNil

`func (o *ApplianceSystemStatus) SetGroupStatusesNil(b bool)`

 SetGroupStatusesNil sets the value for GroupStatuses to be an explicit nil

### UnsetGroupStatuses
`func (o *ApplianceSystemStatus) UnsetGroupStatuses()`

UnsetGroupStatuses ensures that no value is present for GroupStatuses, not even an explicit nil
### GetRegisteredDevice

`func (o *ApplianceSystemStatus) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *ApplianceSystemStatus) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *ApplianceSystemStatus) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *ApplianceSystemStatus) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### GetSystemInfo

`func (o *ApplianceSystemStatus) GetSystemInfo() ApplianceSystemInfoRelationship`

GetSystemInfo returns the SystemInfo field if non-nil, zero value otherwise.

### GetSystemInfoOk

`func (o *ApplianceSystemStatus) GetSystemInfoOk() (*ApplianceSystemInfoRelationship, bool)`

GetSystemInfoOk returns a tuple with the SystemInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemInfo

`func (o *ApplianceSystemStatus) SetSystemInfo(v ApplianceSystemInfoRelationship)`

SetSystemInfo sets SystemInfo field to given value.

### HasSystemInfo

`func (o *ApplianceSystemStatus) HasSystemInfo() bool`

HasSystemInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


