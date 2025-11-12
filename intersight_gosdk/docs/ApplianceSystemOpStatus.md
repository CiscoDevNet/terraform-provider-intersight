# ApplianceSystemOpStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "appliance.SystemOpStatus"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "appliance.SystemOpStatus"]
**OperationalStatus** | Pointer to **string** | Operational status of the Intersight Appliance. Operational status is based on the result of the status checks. If result of any check is Critical, then its value is Impaired. Otherwise, if result of any check is Warning, then its value is AttentionNeeded. If all checks are OK, then its value is Operational. * &#x60;Unknown&#x60; - The status of the appliance node is unknown. * &#x60;Operational&#x60; - The appliance node is operational. * &#x60;Impaired&#x60; - The appliance node is impaired. * &#x60;AttentionNeeded&#x60; - The appliance node needs attention. * &#x60;ReadyToJoin&#x60; - The node is ready to be added to a standalone Intersight Appliance to form a cluster. * &#x60;OutOfService&#x60; - The user has taken this node (part of a cluster) to out of service. * &#x60;ReadyForReplacement&#x60; - The cluster node is ready to be replaced. * &#x60;ReplacementInProgress&#x60; - The cluster node replacement is in progress. * &#x60;ReplacementFailed&#x60; - There was a failure during the cluster node replacement. * &#x60;WorkerNodeInstInProgress&#x60; - The worker node installation is in progress. * &#x60;WorkerNodeInstSuccess&#x60; - The worker node installation succeeded. * &#x60;WorkerNodeInstFailed&#x60; - The worker node installation failed. | [optional] [readonly] [default to "Unknown"]
**Account** | Pointer to [**NullableIamAccountRelationship**](IamAccountRelationship.md) |  | [optional] 
**AppOpStatuses** | Pointer to [**[]ApplianceAppOpStatusRelationship**](ApplianceAppOpStatusRelationship.md) | An array of relationships to applianceAppOpStatus resources. | [optional] [readonly] 
**GroupOpStatuses** | Pointer to [**[]ApplianceGroupOpStatusRelationship**](ApplianceGroupOpStatusRelationship.md) | An array of relationships to applianceGroupOpStatus resources. | [optional] [readonly] 
**RegisteredDevice** | Pointer to [**NullableAssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 
**SystemInfo** | Pointer to [**NullableApplianceSystemInfoRelationship**](ApplianceSystemInfoRelationship.md) |  | [optional] 

## Methods

### NewApplianceSystemOpStatus

`func NewApplianceSystemOpStatus(classId string, objectType string, ) *ApplianceSystemOpStatus`

NewApplianceSystemOpStatus instantiates a new ApplianceSystemOpStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplianceSystemOpStatusWithDefaults

`func NewApplianceSystemOpStatusWithDefaults() *ApplianceSystemOpStatus`

NewApplianceSystemOpStatusWithDefaults instantiates a new ApplianceSystemOpStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ApplianceSystemOpStatus) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ApplianceSystemOpStatus) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ApplianceSystemOpStatus) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ApplianceSystemOpStatus) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ApplianceSystemOpStatus) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ApplianceSystemOpStatus) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetOperationalStatus

`func (o *ApplianceSystemOpStatus) GetOperationalStatus() string`

GetOperationalStatus returns the OperationalStatus field if non-nil, zero value otherwise.

### GetOperationalStatusOk

`func (o *ApplianceSystemOpStatus) GetOperationalStatusOk() (*string, bool)`

GetOperationalStatusOk returns a tuple with the OperationalStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationalStatus

`func (o *ApplianceSystemOpStatus) SetOperationalStatus(v string)`

SetOperationalStatus sets OperationalStatus field to given value.

### HasOperationalStatus

`func (o *ApplianceSystemOpStatus) HasOperationalStatus() bool`

HasOperationalStatus returns a boolean if a field has been set.

### GetAccount

`func (o *ApplianceSystemOpStatus) GetAccount() IamAccountRelationship`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *ApplianceSystemOpStatus) GetAccountOk() (*IamAccountRelationship, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *ApplianceSystemOpStatus) SetAccount(v IamAccountRelationship)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *ApplianceSystemOpStatus) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### SetAccountNil

`func (o *ApplianceSystemOpStatus) SetAccountNil(b bool)`

 SetAccountNil sets the value for Account to be an explicit nil

### UnsetAccount
`func (o *ApplianceSystemOpStatus) UnsetAccount()`

UnsetAccount ensures that no value is present for Account, not even an explicit nil
### GetAppOpStatuses

`func (o *ApplianceSystemOpStatus) GetAppOpStatuses() []ApplianceAppOpStatusRelationship`

GetAppOpStatuses returns the AppOpStatuses field if non-nil, zero value otherwise.

### GetAppOpStatusesOk

`func (o *ApplianceSystemOpStatus) GetAppOpStatusesOk() (*[]ApplianceAppOpStatusRelationship, bool)`

GetAppOpStatusesOk returns a tuple with the AppOpStatuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppOpStatuses

`func (o *ApplianceSystemOpStatus) SetAppOpStatuses(v []ApplianceAppOpStatusRelationship)`

SetAppOpStatuses sets AppOpStatuses field to given value.

### HasAppOpStatuses

`func (o *ApplianceSystemOpStatus) HasAppOpStatuses() bool`

HasAppOpStatuses returns a boolean if a field has been set.

### SetAppOpStatusesNil

`func (o *ApplianceSystemOpStatus) SetAppOpStatusesNil(b bool)`

 SetAppOpStatusesNil sets the value for AppOpStatuses to be an explicit nil

### UnsetAppOpStatuses
`func (o *ApplianceSystemOpStatus) UnsetAppOpStatuses()`

UnsetAppOpStatuses ensures that no value is present for AppOpStatuses, not even an explicit nil
### GetGroupOpStatuses

`func (o *ApplianceSystemOpStatus) GetGroupOpStatuses() []ApplianceGroupOpStatusRelationship`

GetGroupOpStatuses returns the GroupOpStatuses field if non-nil, zero value otherwise.

### GetGroupOpStatusesOk

`func (o *ApplianceSystemOpStatus) GetGroupOpStatusesOk() (*[]ApplianceGroupOpStatusRelationship, bool)`

GetGroupOpStatusesOk returns a tuple with the GroupOpStatuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupOpStatuses

`func (o *ApplianceSystemOpStatus) SetGroupOpStatuses(v []ApplianceGroupOpStatusRelationship)`

SetGroupOpStatuses sets GroupOpStatuses field to given value.

### HasGroupOpStatuses

`func (o *ApplianceSystemOpStatus) HasGroupOpStatuses() bool`

HasGroupOpStatuses returns a boolean if a field has been set.

### SetGroupOpStatusesNil

`func (o *ApplianceSystemOpStatus) SetGroupOpStatusesNil(b bool)`

 SetGroupOpStatusesNil sets the value for GroupOpStatuses to be an explicit nil

### UnsetGroupOpStatuses
`func (o *ApplianceSystemOpStatus) UnsetGroupOpStatuses()`

UnsetGroupOpStatuses ensures that no value is present for GroupOpStatuses, not even an explicit nil
### GetRegisteredDevice

`func (o *ApplianceSystemOpStatus) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *ApplianceSystemOpStatus) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *ApplianceSystemOpStatus) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *ApplianceSystemOpStatus) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### SetRegisteredDeviceNil

`func (o *ApplianceSystemOpStatus) SetRegisteredDeviceNil(b bool)`

 SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil

### UnsetRegisteredDevice
`func (o *ApplianceSystemOpStatus) UnsetRegisteredDevice()`

UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil
### GetSystemInfo

`func (o *ApplianceSystemOpStatus) GetSystemInfo() ApplianceSystemInfoRelationship`

GetSystemInfo returns the SystemInfo field if non-nil, zero value otherwise.

### GetSystemInfoOk

`func (o *ApplianceSystemOpStatus) GetSystemInfoOk() (*ApplianceSystemInfoRelationship, bool)`

GetSystemInfoOk returns a tuple with the SystemInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemInfo

`func (o *ApplianceSystemOpStatus) SetSystemInfo(v ApplianceSystemInfoRelationship)`

SetSystemInfo sets SystemInfo field to given value.

### HasSystemInfo

`func (o *ApplianceSystemOpStatus) HasSystemInfo() bool`

HasSystemInfo returns a boolean if a field has been set.

### SetSystemInfoNil

`func (o *ApplianceSystemOpStatus) SetSystemInfoNil(b bool)`

 SetSystemInfoNil sets the value for SystemInfo to be an explicit nil

### UnsetSystemInfo
`func (o *ApplianceSystemOpStatus) UnsetSystemInfo()`

UnsetSystemInfo ensures that no value is present for SystemInfo, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


