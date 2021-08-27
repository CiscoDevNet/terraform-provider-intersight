# FirmwareDirectDownload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "firmware.DirectDownload"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "firmware.DirectDownload"]
**HttpServer** | Pointer to [**NullableFirmwareHttpServer**](FirmwareHttpServer.md) |  | [optional] 
**ImageSource** | Pointer to **string** | Source type referring the image to be downloaded from CCO or from a local HTTPS server. * &#x60;cisco&#x60; - Image to be downloaded from Cisco software repository. * &#x60;localHttp&#x60; - Image to be downloaded from a https server. | [optional] [default to "cisco"]
**IsPasswordSet** | Pointer to **bool** | Indicates whether the value of the &#39;password&#39; property has been set. | [optional] [readonly] [default to false]
**Password** | Pointer to **string** | Password as configured on the local https server. | [optional] 
**Upgradeoption** | Pointer to **string** | Option to control the upgrade, e.g., sd_upgrade_mount_only - download the image into sd and upgrade wait for the server on-next boot. * &#x60;sd_upgrade_mount_only&#x60; - Direct upgrade SD upgrade mount only. * &#x60;sd_download_only&#x60; - Direct upgrade SD download only. * &#x60;sd_upgrade_only&#x60; - Direct upgrade SD upgrade only. * &#x60;sd_upgrade_full&#x60; - Direct upgrade SD upgrade full. * &#x60;download_only&#x60; - Direct upgrade image download only. * &#x60;upgrade_full&#x60; - The upgrade downloads or mounts the image, and reboots immediately for an upgrade. * &#x60;upgrade_mount_only&#x60; - The upgrade downloads or mounts the image. The upgrade happens in next reboot. * &#x60;chassis_upgrade_full&#x60; - Direct upgrade chassis upgrade full. | [optional] [default to "sd_upgrade_mount_only"]
**Username** | Pointer to **string** | Username as configured on the local https server. | [optional] 

## Methods

### NewFirmwareDirectDownload

`func NewFirmwareDirectDownload(classId string, objectType string, ) *FirmwareDirectDownload`

NewFirmwareDirectDownload instantiates a new FirmwareDirectDownload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirmwareDirectDownloadWithDefaults

`func NewFirmwareDirectDownloadWithDefaults() *FirmwareDirectDownload`

NewFirmwareDirectDownloadWithDefaults instantiates a new FirmwareDirectDownload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *FirmwareDirectDownload) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *FirmwareDirectDownload) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *FirmwareDirectDownload) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *FirmwareDirectDownload) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *FirmwareDirectDownload) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *FirmwareDirectDownload) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetHttpServer

`func (o *FirmwareDirectDownload) GetHttpServer() FirmwareHttpServer`

GetHttpServer returns the HttpServer field if non-nil, zero value otherwise.

### GetHttpServerOk

`func (o *FirmwareDirectDownload) GetHttpServerOk() (*FirmwareHttpServer, bool)`

GetHttpServerOk returns a tuple with the HttpServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpServer

`func (o *FirmwareDirectDownload) SetHttpServer(v FirmwareHttpServer)`

SetHttpServer sets HttpServer field to given value.

### HasHttpServer

`func (o *FirmwareDirectDownload) HasHttpServer() bool`

HasHttpServer returns a boolean if a field has been set.

### SetHttpServerNil

`func (o *FirmwareDirectDownload) SetHttpServerNil(b bool)`

 SetHttpServerNil sets the value for HttpServer to be an explicit nil

### UnsetHttpServer
`func (o *FirmwareDirectDownload) UnsetHttpServer()`

UnsetHttpServer ensures that no value is present for HttpServer, not even an explicit nil
### GetImageSource

`func (o *FirmwareDirectDownload) GetImageSource() string`

GetImageSource returns the ImageSource field if non-nil, zero value otherwise.

### GetImageSourceOk

`func (o *FirmwareDirectDownload) GetImageSourceOk() (*string, bool)`

GetImageSourceOk returns a tuple with the ImageSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageSource

`func (o *FirmwareDirectDownload) SetImageSource(v string)`

SetImageSource sets ImageSource field to given value.

### HasImageSource

`func (o *FirmwareDirectDownload) HasImageSource() bool`

HasImageSource returns a boolean if a field has been set.

### GetIsPasswordSet

`func (o *FirmwareDirectDownload) GetIsPasswordSet() bool`

GetIsPasswordSet returns the IsPasswordSet field if non-nil, zero value otherwise.

### GetIsPasswordSetOk

`func (o *FirmwareDirectDownload) GetIsPasswordSetOk() (*bool, bool)`

GetIsPasswordSetOk returns a tuple with the IsPasswordSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPasswordSet

`func (o *FirmwareDirectDownload) SetIsPasswordSet(v bool)`

SetIsPasswordSet sets IsPasswordSet field to given value.

### HasIsPasswordSet

`func (o *FirmwareDirectDownload) HasIsPasswordSet() bool`

HasIsPasswordSet returns a boolean if a field has been set.

### GetPassword

`func (o *FirmwareDirectDownload) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *FirmwareDirectDownload) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *FirmwareDirectDownload) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *FirmwareDirectDownload) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetUpgradeoption

`func (o *FirmwareDirectDownload) GetUpgradeoption() string`

GetUpgradeoption returns the Upgradeoption field if non-nil, zero value otherwise.

### GetUpgradeoptionOk

`func (o *FirmwareDirectDownload) GetUpgradeoptionOk() (*string, bool)`

GetUpgradeoptionOk returns a tuple with the Upgradeoption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradeoption

`func (o *FirmwareDirectDownload) SetUpgradeoption(v string)`

SetUpgradeoption sets Upgradeoption field to given value.

### HasUpgradeoption

`func (o *FirmwareDirectDownload) HasUpgradeoption() bool`

HasUpgradeoption returns a boolean if a field has been set.

### GetUsername

`func (o *FirmwareDirectDownload) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *FirmwareDirectDownload) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *FirmwareDirectDownload) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *FirmwareDirectDownload) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


