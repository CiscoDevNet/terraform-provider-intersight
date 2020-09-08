# FirmwareDirectDownloadAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HttpServer** | Pointer to [**FirmwareHttpServer**](firmware.HttpServer.md) |  | [optional] 
**ImageSource** | Pointer to **string** | Source type referring the image to be downloaded from CCO or from a local HTTPS server. * &#x60;cisco&#x60; - Image to be downloaded from Cisco software repository. * &#x60;localHttp&#x60; - Image to be downloaded from a https server. | [optional] [default to "cisco"]
**IsPasswordSet** | Pointer to **bool** | Indicates whether the value of the &#39;password&#39; property has been set. | [optional] [readonly] 
**Password** | Pointer to **string** | Password as configured on the local https server. | [optional] 
**Upgradeoption** | Pointer to **string** | Option to control the upgrade, e.g., sd_upgrade_mount_only - download the image into sd and upgrade wait for the server on-next boot. * &#x60;sd_upgrade_mount_only&#x60; - Direct upgrade SD upgrade mount only. * &#x60;sd_download_only&#x60; - Direct upgrade SD download only. * &#x60;sd_upgrade_only&#x60; - Direct upgrade SD upgrade only. * &#x60;sd_upgrade_full&#x60; - Direct upgrade SD upgrade full. * &#x60;upgrade_full&#x60; - The upgrade downloads or mounts the image, and reboots immediately for an upgrade. * &#x60;upgrade_mount_only&#x60; - The upgrade downloads or mounts the image. The upgrade happens in next reboot. * &#x60;chassis_upgrade_full&#x60; - Direct upgrade chassis upgrade full. | [optional] [default to "sd_upgrade_mount_only"]
**Username** | Pointer to **string** | Username as configured on the local https server. | [optional] 

## Methods

### NewFirmwareDirectDownloadAllOf

`func NewFirmwareDirectDownloadAllOf() *FirmwareDirectDownloadAllOf`

NewFirmwareDirectDownloadAllOf instantiates a new FirmwareDirectDownloadAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirmwareDirectDownloadAllOfWithDefaults

`func NewFirmwareDirectDownloadAllOfWithDefaults() *FirmwareDirectDownloadAllOf`

NewFirmwareDirectDownloadAllOfWithDefaults instantiates a new FirmwareDirectDownloadAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHttpServer

`func (o *FirmwareDirectDownloadAllOf) GetHttpServer() FirmwareHttpServer`

GetHttpServer returns the HttpServer field if non-nil, zero value otherwise.

### GetHttpServerOk

`func (o *FirmwareDirectDownloadAllOf) GetHttpServerOk() (*FirmwareHttpServer, bool)`

GetHttpServerOk returns a tuple with the HttpServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpServer

`func (o *FirmwareDirectDownloadAllOf) SetHttpServer(v FirmwareHttpServer)`

SetHttpServer sets HttpServer field to given value.

### HasHttpServer

`func (o *FirmwareDirectDownloadAllOf) HasHttpServer() bool`

HasHttpServer returns a boolean if a field has been set.

### GetImageSource

`func (o *FirmwareDirectDownloadAllOf) GetImageSource() string`

GetImageSource returns the ImageSource field if non-nil, zero value otherwise.

### GetImageSourceOk

`func (o *FirmwareDirectDownloadAllOf) GetImageSourceOk() (*string, bool)`

GetImageSourceOk returns a tuple with the ImageSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageSource

`func (o *FirmwareDirectDownloadAllOf) SetImageSource(v string)`

SetImageSource sets ImageSource field to given value.

### HasImageSource

`func (o *FirmwareDirectDownloadAllOf) HasImageSource() bool`

HasImageSource returns a boolean if a field has been set.

### GetIsPasswordSet

`func (o *FirmwareDirectDownloadAllOf) GetIsPasswordSet() bool`

GetIsPasswordSet returns the IsPasswordSet field if non-nil, zero value otherwise.

### GetIsPasswordSetOk

`func (o *FirmwareDirectDownloadAllOf) GetIsPasswordSetOk() (*bool, bool)`

GetIsPasswordSetOk returns a tuple with the IsPasswordSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPasswordSet

`func (o *FirmwareDirectDownloadAllOf) SetIsPasswordSet(v bool)`

SetIsPasswordSet sets IsPasswordSet field to given value.

### HasIsPasswordSet

`func (o *FirmwareDirectDownloadAllOf) HasIsPasswordSet() bool`

HasIsPasswordSet returns a boolean if a field has been set.

### GetPassword

`func (o *FirmwareDirectDownloadAllOf) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *FirmwareDirectDownloadAllOf) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *FirmwareDirectDownloadAllOf) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *FirmwareDirectDownloadAllOf) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetUpgradeoption

`func (o *FirmwareDirectDownloadAllOf) GetUpgradeoption() string`

GetUpgradeoption returns the Upgradeoption field if non-nil, zero value otherwise.

### GetUpgradeoptionOk

`func (o *FirmwareDirectDownloadAllOf) GetUpgradeoptionOk() (*string, bool)`

GetUpgradeoptionOk returns a tuple with the Upgradeoption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradeoption

`func (o *FirmwareDirectDownloadAllOf) SetUpgradeoption(v string)`

SetUpgradeoption sets Upgradeoption field to given value.

### HasUpgradeoption

`func (o *FirmwareDirectDownloadAllOf) HasUpgradeoption() bool`

HasUpgradeoption returns a boolean if a field has been set.

### GetUsername

`func (o *FirmwareDirectDownloadAllOf) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *FirmwareDirectDownloadAllOf) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *FirmwareDirectDownloadAllOf) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *FirmwareDirectDownloadAllOf) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


