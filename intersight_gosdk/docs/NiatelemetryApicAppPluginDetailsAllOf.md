# NiatelemetryApicAppPluginDetailsAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "niatelemetry.ApicAppPluginDetails"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "niatelemetry.ApicAppPluginDetails"]
**AppId** | Pointer to **string** | App Id of the plugin in APIC. | [optional] 
**Permission** | Pointer to **string** | Permissions to the app plugin in APIC. | [optional] 
**PermissionLevel** | Pointer to **string** | Permission Level of the app plugin in APIC. | [optional] 
**RecordType** | Pointer to **string** | Type of record DCNM / APIC / SE. This determines the type of platform where inventory was collected. | [optional] 
**RecordVersion** | Pointer to **string** | Version of record being pushed. This determines what was the API version for data available from the device. | [optional] 
**SiteName** | Pointer to **string** | Name of the APIC site from which this data is being collected. | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewNiatelemetryApicAppPluginDetailsAllOf

`func NewNiatelemetryApicAppPluginDetailsAllOf(classId string, objectType string, ) *NiatelemetryApicAppPluginDetailsAllOf`

NewNiatelemetryApicAppPluginDetailsAllOf instantiates a new NiatelemetryApicAppPluginDetailsAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiatelemetryApicAppPluginDetailsAllOfWithDefaults

`func NewNiatelemetryApicAppPluginDetailsAllOfWithDefaults() *NiatelemetryApicAppPluginDetailsAllOf`

NewNiatelemetryApicAppPluginDetailsAllOfWithDefaults instantiates a new NiatelemetryApicAppPluginDetailsAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NiatelemetryApicAppPluginDetailsAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NiatelemetryApicAppPluginDetailsAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NiatelemetryApicAppPluginDetailsAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NiatelemetryApicAppPluginDetailsAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NiatelemetryApicAppPluginDetailsAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NiatelemetryApicAppPluginDetailsAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAppId

`func (o *NiatelemetryApicAppPluginDetailsAllOf) GetAppId() string`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *NiatelemetryApicAppPluginDetailsAllOf) GetAppIdOk() (*string, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *NiatelemetryApicAppPluginDetailsAllOf) SetAppId(v string)`

SetAppId sets AppId field to given value.

### HasAppId

`func (o *NiatelemetryApicAppPluginDetailsAllOf) HasAppId() bool`

HasAppId returns a boolean if a field has been set.

### GetPermission

`func (o *NiatelemetryApicAppPluginDetailsAllOf) GetPermission() string`

GetPermission returns the Permission field if non-nil, zero value otherwise.

### GetPermissionOk

`func (o *NiatelemetryApicAppPluginDetailsAllOf) GetPermissionOk() (*string, bool)`

GetPermissionOk returns a tuple with the Permission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermission

`func (o *NiatelemetryApicAppPluginDetailsAllOf) SetPermission(v string)`

SetPermission sets Permission field to given value.

### HasPermission

`func (o *NiatelemetryApicAppPluginDetailsAllOf) HasPermission() bool`

HasPermission returns a boolean if a field has been set.

### GetPermissionLevel

`func (o *NiatelemetryApicAppPluginDetailsAllOf) GetPermissionLevel() string`

GetPermissionLevel returns the PermissionLevel field if non-nil, zero value otherwise.

### GetPermissionLevelOk

`func (o *NiatelemetryApicAppPluginDetailsAllOf) GetPermissionLevelOk() (*string, bool)`

GetPermissionLevelOk returns a tuple with the PermissionLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionLevel

`func (o *NiatelemetryApicAppPluginDetailsAllOf) SetPermissionLevel(v string)`

SetPermissionLevel sets PermissionLevel field to given value.

### HasPermissionLevel

`func (o *NiatelemetryApicAppPluginDetailsAllOf) HasPermissionLevel() bool`

HasPermissionLevel returns a boolean if a field has been set.

### GetRecordType

`func (o *NiatelemetryApicAppPluginDetailsAllOf) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *NiatelemetryApicAppPluginDetailsAllOf) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *NiatelemetryApicAppPluginDetailsAllOf) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *NiatelemetryApicAppPluginDetailsAllOf) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetRecordVersion

`func (o *NiatelemetryApicAppPluginDetailsAllOf) GetRecordVersion() string`

GetRecordVersion returns the RecordVersion field if non-nil, zero value otherwise.

### GetRecordVersionOk

`func (o *NiatelemetryApicAppPluginDetailsAllOf) GetRecordVersionOk() (*string, bool)`

GetRecordVersionOk returns a tuple with the RecordVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordVersion

`func (o *NiatelemetryApicAppPluginDetailsAllOf) SetRecordVersion(v string)`

SetRecordVersion sets RecordVersion field to given value.

### HasRecordVersion

`func (o *NiatelemetryApicAppPluginDetailsAllOf) HasRecordVersion() bool`

HasRecordVersion returns a boolean if a field has been set.

### GetSiteName

`func (o *NiatelemetryApicAppPluginDetailsAllOf) GetSiteName() string`

GetSiteName returns the SiteName field if non-nil, zero value otherwise.

### GetSiteNameOk

`func (o *NiatelemetryApicAppPluginDetailsAllOf) GetSiteNameOk() (*string, bool)`

GetSiteNameOk returns a tuple with the SiteName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteName

`func (o *NiatelemetryApicAppPluginDetailsAllOf) SetSiteName(v string)`

SetSiteName sets SiteName field to given value.

### HasSiteName

`func (o *NiatelemetryApicAppPluginDetailsAllOf) HasSiteName() bool`

HasSiteName returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *NiatelemetryApicAppPluginDetailsAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *NiatelemetryApicAppPluginDetailsAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *NiatelemetryApicAppPluginDetailsAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *NiatelemetryApicAppPluginDetailsAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


