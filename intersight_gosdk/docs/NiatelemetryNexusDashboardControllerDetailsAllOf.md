# NiatelemetryNexusDashboardControllerDetailsAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "niatelemetry.NexusDashboardControllerDetails"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "niatelemetry.NexusDashboardControllerDetails"]
**SiteHealth** | Pointer to **int64** | Health of the site serviced by ND. | [optional] 
**SiteName** | Pointer to **string** | Name of fabric domain of the controller. | [optional] 
**VersionOfController** | Pointer to **string** | Version of the controller serviced by ND. | [optional] 
**NexusDashboard** | Pointer to [**NiatelemetryNexusDashboardsRelationship**](niatelemetry.NexusDashboards.Relationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 

## Methods

### NewNiatelemetryNexusDashboardControllerDetailsAllOf

`func NewNiatelemetryNexusDashboardControllerDetailsAllOf(classId string, objectType string, ) *NiatelemetryNexusDashboardControllerDetailsAllOf`

NewNiatelemetryNexusDashboardControllerDetailsAllOf instantiates a new NiatelemetryNexusDashboardControllerDetailsAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiatelemetryNexusDashboardControllerDetailsAllOfWithDefaults

`func NewNiatelemetryNexusDashboardControllerDetailsAllOfWithDefaults() *NiatelemetryNexusDashboardControllerDetailsAllOf`

NewNiatelemetryNexusDashboardControllerDetailsAllOfWithDefaults instantiates a new NiatelemetryNexusDashboardControllerDetailsAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NiatelemetryNexusDashboardControllerDetailsAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NiatelemetryNexusDashboardControllerDetailsAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NiatelemetryNexusDashboardControllerDetailsAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NiatelemetryNexusDashboardControllerDetailsAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NiatelemetryNexusDashboardControllerDetailsAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NiatelemetryNexusDashboardControllerDetailsAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetSiteHealth

`func (o *NiatelemetryNexusDashboardControllerDetailsAllOf) GetSiteHealth() int64`

GetSiteHealth returns the SiteHealth field if non-nil, zero value otherwise.

### GetSiteHealthOk

`func (o *NiatelemetryNexusDashboardControllerDetailsAllOf) GetSiteHealthOk() (*int64, bool)`

GetSiteHealthOk returns a tuple with the SiteHealth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteHealth

`func (o *NiatelemetryNexusDashboardControllerDetailsAllOf) SetSiteHealth(v int64)`

SetSiteHealth sets SiteHealth field to given value.

### HasSiteHealth

`func (o *NiatelemetryNexusDashboardControllerDetailsAllOf) HasSiteHealth() bool`

HasSiteHealth returns a boolean if a field has been set.

### GetSiteName

`func (o *NiatelemetryNexusDashboardControllerDetailsAllOf) GetSiteName() string`

GetSiteName returns the SiteName field if non-nil, zero value otherwise.

### GetSiteNameOk

`func (o *NiatelemetryNexusDashboardControllerDetailsAllOf) GetSiteNameOk() (*string, bool)`

GetSiteNameOk returns a tuple with the SiteName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteName

`func (o *NiatelemetryNexusDashboardControllerDetailsAllOf) SetSiteName(v string)`

SetSiteName sets SiteName field to given value.

### HasSiteName

`func (o *NiatelemetryNexusDashboardControllerDetailsAllOf) HasSiteName() bool`

HasSiteName returns a boolean if a field has been set.

### GetVersionOfController

`func (o *NiatelemetryNexusDashboardControllerDetailsAllOf) GetVersionOfController() string`

GetVersionOfController returns the VersionOfController field if non-nil, zero value otherwise.

### GetVersionOfControllerOk

`func (o *NiatelemetryNexusDashboardControllerDetailsAllOf) GetVersionOfControllerOk() (*string, bool)`

GetVersionOfControllerOk returns a tuple with the VersionOfController field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionOfController

`func (o *NiatelemetryNexusDashboardControllerDetailsAllOf) SetVersionOfController(v string)`

SetVersionOfController sets VersionOfController field to given value.

### HasVersionOfController

`func (o *NiatelemetryNexusDashboardControllerDetailsAllOf) HasVersionOfController() bool`

HasVersionOfController returns a boolean if a field has been set.

### GetNexusDashboard

`func (o *NiatelemetryNexusDashboardControllerDetailsAllOf) GetNexusDashboard() NiatelemetryNexusDashboardsRelationship`

GetNexusDashboard returns the NexusDashboard field if non-nil, zero value otherwise.

### GetNexusDashboardOk

`func (o *NiatelemetryNexusDashboardControllerDetailsAllOf) GetNexusDashboardOk() (*NiatelemetryNexusDashboardsRelationship, bool)`

GetNexusDashboardOk returns a tuple with the NexusDashboard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNexusDashboard

`func (o *NiatelemetryNexusDashboardControllerDetailsAllOf) SetNexusDashboard(v NiatelemetryNexusDashboardsRelationship)`

SetNexusDashboard sets NexusDashboard field to given value.

### HasNexusDashboard

`func (o *NiatelemetryNexusDashboardControllerDetailsAllOf) HasNexusDashboard() bool`

HasNexusDashboard returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *NiatelemetryNexusDashboardControllerDetailsAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *NiatelemetryNexusDashboardControllerDetailsAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *NiatelemetryNexusDashboardControllerDetailsAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *NiatelemetryNexusDashboardControllerDetailsAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


