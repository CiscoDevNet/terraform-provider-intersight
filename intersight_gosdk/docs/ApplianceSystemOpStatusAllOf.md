# ApplianceSystemOpStatusAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "appliance.SystemOpStatus"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "appliance.SystemOpStatus"]
**OperationalStatus** | Pointer to **string** | Operational status of the Intersight Appliance. Operational status is based on the result of the status checks. If result of any check is Critical, then its value is Impaired. Otherwise, if result of any check is Warning, then its value is AttentionNeeded. If all checks are OK, then its value is Operational. * &#x60;Unknown&#x60; - The status of the appliance node is unknown. * &#x60;Operational&#x60; - The appliance node is operational. * &#x60;Impaired&#x60; - The appliance node is impaired. * &#x60;AttentionNeeded&#x60; - The appliance node needs attention. * &#x60;ReadyToJoin&#x60; - The node is ready to be added to a standalone Intersight Appliance to form a cluster. * &#x60;OutOfService&#x60; - The user has taken this node (part of a cluster) to out of service. * &#x60;ReadyForReplacement&#x60; - The cluster node is ready to be replaced. * &#x60;ReplacementInProgress&#x60; - The cluster node replacement is in progress. * &#x60;ReplacementFailed&#x60; - There was a failure during the cluster node replacement. | [optional] [readonly] [default to "Unknown"]
**Account** | Pointer to [**IamAccountRelationship**](IamAccountRelationship.md) |  | [optional] 
**AppOpStatuses** | Pointer to [**[]ApplianceAppOpStatusRelationship**](ApplianceAppOpStatusRelationship.md) | An array of relationships to applianceAppOpStatus resources. | [optional] [readonly] 
**GroupOpStatuses** | Pointer to [**[]ApplianceGroupOpStatusRelationship**](ApplianceGroupOpStatusRelationship.md) | An array of relationships to applianceGroupOpStatus resources. | [optional] [readonly] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 
**SystemInfo** | Pointer to [**ApplianceSystemInfoRelationship**](ApplianceSystemInfoRelationship.md) |  | [optional] 

## Methods

### NewApplianceSystemOpStatusAllOf

`func NewApplianceSystemOpStatusAllOf(classId string, objectType string, ) *ApplianceSystemOpStatusAllOf`

NewApplianceSystemOpStatusAllOf instantiates a new ApplianceSystemOpStatusAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplianceSystemOpStatusAllOfWithDefaults

`func NewApplianceSystemOpStatusAllOfWithDefaults() *ApplianceSystemOpStatusAllOf`

NewApplianceSystemOpStatusAllOfWithDefaults instantiates a new ApplianceSystemOpStatusAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ApplianceSystemOpStatusAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ApplianceSystemOpStatusAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ApplianceSystemOpStatusAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ApplianceSystemOpStatusAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ApplianceSystemOpStatusAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ApplianceSystemOpStatusAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetOperationalStatus

`func (o *ApplianceSystemOpStatusAllOf) GetOperationalStatus() string`

GetOperationalStatus returns the OperationalStatus field if non-nil, zero value otherwise.

### GetOperationalStatusOk

`func (o *ApplianceSystemOpStatusAllOf) GetOperationalStatusOk() (*string, bool)`

GetOperationalStatusOk returns a tuple with the OperationalStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationalStatus

`func (o *ApplianceSystemOpStatusAllOf) SetOperationalStatus(v string)`

SetOperationalStatus sets OperationalStatus field to given value.

### HasOperationalStatus

`func (o *ApplianceSystemOpStatusAllOf) HasOperationalStatus() bool`

HasOperationalStatus returns a boolean if a field has been set.

### GetAccount

`func (o *ApplianceSystemOpStatusAllOf) GetAccount() IamAccountRelationship`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *ApplianceSystemOpStatusAllOf) GetAccountOk() (*IamAccountRelationship, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *ApplianceSystemOpStatusAllOf) SetAccount(v IamAccountRelationship)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *ApplianceSystemOpStatusAllOf) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetAppOpStatuses

`func (o *ApplianceSystemOpStatusAllOf) GetAppOpStatuses() []ApplianceAppOpStatusRelationship`

GetAppOpStatuses returns the AppOpStatuses field if non-nil, zero value otherwise.

### GetAppOpStatusesOk

`func (o *ApplianceSystemOpStatusAllOf) GetAppOpStatusesOk() (*[]ApplianceAppOpStatusRelationship, bool)`

GetAppOpStatusesOk returns a tuple with the AppOpStatuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppOpStatuses

`func (o *ApplianceSystemOpStatusAllOf) SetAppOpStatuses(v []ApplianceAppOpStatusRelationship)`

SetAppOpStatuses sets AppOpStatuses field to given value.

### HasAppOpStatuses

`func (o *ApplianceSystemOpStatusAllOf) HasAppOpStatuses() bool`

HasAppOpStatuses returns a boolean if a field has been set.

### SetAppOpStatusesNil

`func (o *ApplianceSystemOpStatusAllOf) SetAppOpStatusesNil(b bool)`

 SetAppOpStatusesNil sets the value for AppOpStatuses to be an explicit nil

### UnsetAppOpStatuses
`func (o *ApplianceSystemOpStatusAllOf) UnsetAppOpStatuses()`

UnsetAppOpStatuses ensures that no value is present for AppOpStatuses, not even an explicit nil
### GetGroupOpStatuses

`func (o *ApplianceSystemOpStatusAllOf) GetGroupOpStatuses() []ApplianceGroupOpStatusRelationship`

GetGroupOpStatuses returns the GroupOpStatuses field if non-nil, zero value otherwise.

### GetGroupOpStatusesOk

`func (o *ApplianceSystemOpStatusAllOf) GetGroupOpStatusesOk() (*[]ApplianceGroupOpStatusRelationship, bool)`

GetGroupOpStatusesOk returns a tuple with the GroupOpStatuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupOpStatuses

`func (o *ApplianceSystemOpStatusAllOf) SetGroupOpStatuses(v []ApplianceGroupOpStatusRelationship)`

SetGroupOpStatuses sets GroupOpStatuses field to given value.

### HasGroupOpStatuses

`func (o *ApplianceSystemOpStatusAllOf) HasGroupOpStatuses() bool`

HasGroupOpStatuses returns a boolean if a field has been set.

### SetGroupOpStatusesNil

`func (o *ApplianceSystemOpStatusAllOf) SetGroupOpStatusesNil(b bool)`

 SetGroupOpStatusesNil sets the value for GroupOpStatuses to be an explicit nil

### UnsetGroupOpStatuses
`func (o *ApplianceSystemOpStatusAllOf) UnsetGroupOpStatuses()`

UnsetGroupOpStatuses ensures that no value is present for GroupOpStatuses, not even an explicit nil
### GetRegisteredDevice

`func (o *ApplianceSystemOpStatusAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *ApplianceSystemOpStatusAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *ApplianceSystemOpStatusAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *ApplianceSystemOpStatusAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### GetSystemInfo

`func (o *ApplianceSystemOpStatusAllOf) GetSystemInfo() ApplianceSystemInfoRelationship`

GetSystemInfo returns the SystemInfo field if non-nil, zero value otherwise.

### GetSystemInfoOk

`func (o *ApplianceSystemOpStatusAllOf) GetSystemInfoOk() (*ApplianceSystemInfoRelationship, bool)`

GetSystemInfoOk returns a tuple with the SystemInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemInfo

`func (o *ApplianceSystemOpStatusAllOf) SetSystemInfo(v ApplianceSystemInfoRelationship)`

SetSystemInfo sets SystemInfo field to given value.

### HasSystemInfo

`func (o *ApplianceSystemOpStatusAllOf) HasSystemInfo() bool`

HasSystemInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


