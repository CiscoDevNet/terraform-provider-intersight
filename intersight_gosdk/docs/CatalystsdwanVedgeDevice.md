# CatalystsdwanVedgeDevice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "catalystsdwan.VedgeDevice"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "catalystsdwan.VedgeDevice"]
**ConfigStatusMessage** | Pointer to **string** | The Catalyst SDWAN device config status message. | [optional] 
**DeviceState** | Pointer to **string** | The Catalyst SDWAN device state. | [optional] 
**HostName** | Pointer to **string** | The Catalyst SDWAN device host name. | [optional] 
**PlatformFamily** | Pointer to **string** | The Catalyst SDWAN device platform family. | [optional] 
**Reachability** | Pointer to **string** | The Catalyst SDWAN device reachability. | [optional] 
**SiteId** | Pointer to **string** | The Catalyst SDWAN device site id. | [optional] 
**SiteName** | Pointer to **string** | The Catalyst SDWAN device site name. | [optional] 
**SpOrganizationName** | Pointer to **string** | The Catalyst SDWAN device sp organization name. | [optional] 
**SystemIp** | Pointer to **string** | The Catalyst SDWAN device system IP. | [optional] 
**TemplateStatus** | Pointer to **string** | The Catalyst SDWAN device template status. | [optional] 
**Validity** | Pointer to **string** | The Catalyst SDWAN device validity. | [optional] 
**RegisteredDevice** | Pointer to [**NullableAssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewCatalystsdwanVedgeDevice

`func NewCatalystsdwanVedgeDevice(classId string, objectType string, ) *CatalystsdwanVedgeDevice`

NewCatalystsdwanVedgeDevice instantiates a new CatalystsdwanVedgeDevice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalystsdwanVedgeDeviceWithDefaults

`func NewCatalystsdwanVedgeDeviceWithDefaults() *CatalystsdwanVedgeDevice`

NewCatalystsdwanVedgeDeviceWithDefaults instantiates a new CatalystsdwanVedgeDevice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *CatalystsdwanVedgeDevice) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *CatalystsdwanVedgeDevice) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *CatalystsdwanVedgeDevice) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *CatalystsdwanVedgeDevice) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CatalystsdwanVedgeDevice) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CatalystsdwanVedgeDevice) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetConfigStatusMessage

`func (o *CatalystsdwanVedgeDevice) GetConfigStatusMessage() string`

GetConfigStatusMessage returns the ConfigStatusMessage field if non-nil, zero value otherwise.

### GetConfigStatusMessageOk

`func (o *CatalystsdwanVedgeDevice) GetConfigStatusMessageOk() (*string, bool)`

GetConfigStatusMessageOk returns a tuple with the ConfigStatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigStatusMessage

`func (o *CatalystsdwanVedgeDevice) SetConfigStatusMessage(v string)`

SetConfigStatusMessage sets ConfigStatusMessage field to given value.

### HasConfigStatusMessage

`func (o *CatalystsdwanVedgeDevice) HasConfigStatusMessage() bool`

HasConfigStatusMessage returns a boolean if a field has been set.

### GetDeviceState

`func (o *CatalystsdwanVedgeDevice) GetDeviceState() string`

GetDeviceState returns the DeviceState field if non-nil, zero value otherwise.

### GetDeviceStateOk

`func (o *CatalystsdwanVedgeDevice) GetDeviceStateOk() (*string, bool)`

GetDeviceStateOk returns a tuple with the DeviceState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceState

`func (o *CatalystsdwanVedgeDevice) SetDeviceState(v string)`

SetDeviceState sets DeviceState field to given value.

### HasDeviceState

`func (o *CatalystsdwanVedgeDevice) HasDeviceState() bool`

HasDeviceState returns a boolean if a field has been set.

### GetHostName

`func (o *CatalystsdwanVedgeDevice) GetHostName() string`

GetHostName returns the HostName field if non-nil, zero value otherwise.

### GetHostNameOk

`func (o *CatalystsdwanVedgeDevice) GetHostNameOk() (*string, bool)`

GetHostNameOk returns a tuple with the HostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostName

`func (o *CatalystsdwanVedgeDevice) SetHostName(v string)`

SetHostName sets HostName field to given value.

### HasHostName

`func (o *CatalystsdwanVedgeDevice) HasHostName() bool`

HasHostName returns a boolean if a field has been set.

### GetPlatformFamily

`func (o *CatalystsdwanVedgeDevice) GetPlatformFamily() string`

GetPlatformFamily returns the PlatformFamily field if non-nil, zero value otherwise.

### GetPlatformFamilyOk

`func (o *CatalystsdwanVedgeDevice) GetPlatformFamilyOk() (*string, bool)`

GetPlatformFamilyOk returns a tuple with the PlatformFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformFamily

`func (o *CatalystsdwanVedgeDevice) SetPlatformFamily(v string)`

SetPlatformFamily sets PlatformFamily field to given value.

### HasPlatformFamily

`func (o *CatalystsdwanVedgeDevice) HasPlatformFamily() bool`

HasPlatformFamily returns a boolean if a field has been set.

### GetReachability

`func (o *CatalystsdwanVedgeDevice) GetReachability() string`

GetReachability returns the Reachability field if non-nil, zero value otherwise.

### GetReachabilityOk

`func (o *CatalystsdwanVedgeDevice) GetReachabilityOk() (*string, bool)`

GetReachabilityOk returns a tuple with the Reachability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReachability

`func (o *CatalystsdwanVedgeDevice) SetReachability(v string)`

SetReachability sets Reachability field to given value.

### HasReachability

`func (o *CatalystsdwanVedgeDevice) HasReachability() bool`

HasReachability returns a boolean if a field has been set.

### GetSiteId

`func (o *CatalystsdwanVedgeDevice) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *CatalystsdwanVedgeDevice) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *CatalystsdwanVedgeDevice) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *CatalystsdwanVedgeDevice) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetSiteName

`func (o *CatalystsdwanVedgeDevice) GetSiteName() string`

GetSiteName returns the SiteName field if non-nil, zero value otherwise.

### GetSiteNameOk

`func (o *CatalystsdwanVedgeDevice) GetSiteNameOk() (*string, bool)`

GetSiteNameOk returns a tuple with the SiteName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteName

`func (o *CatalystsdwanVedgeDevice) SetSiteName(v string)`

SetSiteName sets SiteName field to given value.

### HasSiteName

`func (o *CatalystsdwanVedgeDevice) HasSiteName() bool`

HasSiteName returns a boolean if a field has been set.

### GetSpOrganizationName

`func (o *CatalystsdwanVedgeDevice) GetSpOrganizationName() string`

GetSpOrganizationName returns the SpOrganizationName field if non-nil, zero value otherwise.

### GetSpOrganizationNameOk

`func (o *CatalystsdwanVedgeDevice) GetSpOrganizationNameOk() (*string, bool)`

GetSpOrganizationNameOk returns a tuple with the SpOrganizationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpOrganizationName

`func (o *CatalystsdwanVedgeDevice) SetSpOrganizationName(v string)`

SetSpOrganizationName sets SpOrganizationName field to given value.

### HasSpOrganizationName

`func (o *CatalystsdwanVedgeDevice) HasSpOrganizationName() bool`

HasSpOrganizationName returns a boolean if a field has been set.

### GetSystemIp

`func (o *CatalystsdwanVedgeDevice) GetSystemIp() string`

GetSystemIp returns the SystemIp field if non-nil, zero value otherwise.

### GetSystemIpOk

`func (o *CatalystsdwanVedgeDevice) GetSystemIpOk() (*string, bool)`

GetSystemIpOk returns a tuple with the SystemIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemIp

`func (o *CatalystsdwanVedgeDevice) SetSystemIp(v string)`

SetSystemIp sets SystemIp field to given value.

### HasSystemIp

`func (o *CatalystsdwanVedgeDevice) HasSystemIp() bool`

HasSystemIp returns a boolean if a field has been set.

### GetTemplateStatus

`func (o *CatalystsdwanVedgeDevice) GetTemplateStatus() string`

GetTemplateStatus returns the TemplateStatus field if non-nil, zero value otherwise.

### GetTemplateStatusOk

`func (o *CatalystsdwanVedgeDevice) GetTemplateStatusOk() (*string, bool)`

GetTemplateStatusOk returns a tuple with the TemplateStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateStatus

`func (o *CatalystsdwanVedgeDevice) SetTemplateStatus(v string)`

SetTemplateStatus sets TemplateStatus field to given value.

### HasTemplateStatus

`func (o *CatalystsdwanVedgeDevice) HasTemplateStatus() bool`

HasTemplateStatus returns a boolean if a field has been set.

### GetValidity

`func (o *CatalystsdwanVedgeDevice) GetValidity() string`

GetValidity returns the Validity field if non-nil, zero value otherwise.

### GetValidityOk

`func (o *CatalystsdwanVedgeDevice) GetValidityOk() (*string, bool)`

GetValidityOk returns a tuple with the Validity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidity

`func (o *CatalystsdwanVedgeDevice) SetValidity(v string)`

SetValidity sets Validity field to given value.

### HasValidity

`func (o *CatalystsdwanVedgeDevice) HasValidity() bool`

HasValidity returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *CatalystsdwanVedgeDevice) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *CatalystsdwanVedgeDevice) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *CatalystsdwanVedgeDevice) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *CatalystsdwanVedgeDevice) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### SetRegisteredDeviceNil

`func (o *CatalystsdwanVedgeDevice) SetRegisteredDeviceNil(b bool)`

 SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil

### UnsetRegisteredDevice
`func (o *CatalystsdwanVedgeDevice) UnsetRegisteredDevice()`

UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


