# FirmwareNetworkShare

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CifsServer** | Pointer to [**FirmwareCifsServer**](firmware.CifsServer.md) |  | [optional] 
**HttpServer** | Pointer to [**FirmwareHttpServer**](firmware.HttpServer.md) |  | [optional] 
**IsPasswordSet** | Pointer to **bool** | Indicates whether the value of the &#39;password&#39; property has been set. | [optional] [readonly] 
**MapType** | Pointer to **string** | File server protocols such as CIFS, NFS, WWW for HTTP (S) that hosts the image. * &#x60;nfs&#x60; - File server protocol used is NFS. * &#x60;cifs&#x60; - File server protocol used is CIFS. * &#x60;www&#x60; - File server protocol used is WWW. | [optional] [default to "nfs"]
**NfsServer** | Pointer to [**FirmwareNfsServer**](firmware.NfsServer.md) |  | [optional] 
**Password** | Pointer to **string** | Password as configured on the file server. | [optional] 
**Upgradeoption** | Pointer to **string** | Option to control the upgrade operation. Some examples, 1) nw_upgrade_mount_only - mount the image from a file server and run the upgrade on the next server boot and 2) nw_upgrade_full - mount the image and immediately run the upgrade. * &#x60;nw_upgrade_full&#x60; - Network upgrade option for full upgrade. * &#x60;nw_upgrade_mount_only&#x60; - Network upgrade mount only. | [optional] [default to "nw_upgrade_full"]
**Username** | Pointer to **string** | Username as configured on the file server. | [optional] 

## Methods

### NewFirmwareNetworkShare

`func NewFirmwareNetworkShare() *FirmwareNetworkShare`

NewFirmwareNetworkShare instantiates a new FirmwareNetworkShare object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirmwareNetworkShareWithDefaults

`func NewFirmwareNetworkShareWithDefaults() *FirmwareNetworkShare`

NewFirmwareNetworkShareWithDefaults instantiates a new FirmwareNetworkShare object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCifsServer

`func (o *FirmwareNetworkShare) GetCifsServer() FirmwareCifsServer`

GetCifsServer returns the CifsServer field if non-nil, zero value otherwise.

### GetCifsServerOk

`func (o *FirmwareNetworkShare) GetCifsServerOk() (*FirmwareCifsServer, bool)`

GetCifsServerOk returns a tuple with the CifsServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCifsServer

`func (o *FirmwareNetworkShare) SetCifsServer(v FirmwareCifsServer)`

SetCifsServer sets CifsServer field to given value.

### HasCifsServer

`func (o *FirmwareNetworkShare) HasCifsServer() bool`

HasCifsServer returns a boolean if a field has been set.

### GetHttpServer

`func (o *FirmwareNetworkShare) GetHttpServer() FirmwareHttpServer`

GetHttpServer returns the HttpServer field if non-nil, zero value otherwise.

### GetHttpServerOk

`func (o *FirmwareNetworkShare) GetHttpServerOk() (*FirmwareHttpServer, bool)`

GetHttpServerOk returns a tuple with the HttpServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpServer

`func (o *FirmwareNetworkShare) SetHttpServer(v FirmwareHttpServer)`

SetHttpServer sets HttpServer field to given value.

### HasHttpServer

`func (o *FirmwareNetworkShare) HasHttpServer() bool`

HasHttpServer returns a boolean if a field has been set.

### GetIsPasswordSet

`func (o *FirmwareNetworkShare) GetIsPasswordSet() bool`

GetIsPasswordSet returns the IsPasswordSet field if non-nil, zero value otherwise.

### GetIsPasswordSetOk

`func (o *FirmwareNetworkShare) GetIsPasswordSetOk() (*bool, bool)`

GetIsPasswordSetOk returns a tuple with the IsPasswordSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPasswordSet

`func (o *FirmwareNetworkShare) SetIsPasswordSet(v bool)`

SetIsPasswordSet sets IsPasswordSet field to given value.

### HasIsPasswordSet

`func (o *FirmwareNetworkShare) HasIsPasswordSet() bool`

HasIsPasswordSet returns a boolean if a field has been set.

### GetMapType

`func (o *FirmwareNetworkShare) GetMapType() string`

GetMapType returns the MapType field if non-nil, zero value otherwise.

### GetMapTypeOk

`func (o *FirmwareNetworkShare) GetMapTypeOk() (*string, bool)`

GetMapTypeOk returns a tuple with the MapType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMapType

`func (o *FirmwareNetworkShare) SetMapType(v string)`

SetMapType sets MapType field to given value.

### HasMapType

`func (o *FirmwareNetworkShare) HasMapType() bool`

HasMapType returns a boolean if a field has been set.

### GetNfsServer

`func (o *FirmwareNetworkShare) GetNfsServer() FirmwareNfsServer`

GetNfsServer returns the NfsServer field if non-nil, zero value otherwise.

### GetNfsServerOk

`func (o *FirmwareNetworkShare) GetNfsServerOk() (*FirmwareNfsServer, bool)`

GetNfsServerOk returns a tuple with the NfsServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfsServer

`func (o *FirmwareNetworkShare) SetNfsServer(v FirmwareNfsServer)`

SetNfsServer sets NfsServer field to given value.

### HasNfsServer

`func (o *FirmwareNetworkShare) HasNfsServer() bool`

HasNfsServer returns a boolean if a field has been set.

### GetPassword

`func (o *FirmwareNetworkShare) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *FirmwareNetworkShare) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *FirmwareNetworkShare) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *FirmwareNetworkShare) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetUpgradeoption

`func (o *FirmwareNetworkShare) GetUpgradeoption() string`

GetUpgradeoption returns the Upgradeoption field if non-nil, zero value otherwise.

### GetUpgradeoptionOk

`func (o *FirmwareNetworkShare) GetUpgradeoptionOk() (*string, bool)`

GetUpgradeoptionOk returns a tuple with the Upgradeoption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradeoption

`func (o *FirmwareNetworkShare) SetUpgradeoption(v string)`

SetUpgradeoption sets Upgradeoption field to given value.

### HasUpgradeoption

`func (o *FirmwareNetworkShare) HasUpgradeoption() bool`

HasUpgradeoption returns a boolean if a field has been set.

### GetUsername

`func (o *FirmwareNetworkShare) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *FirmwareNetworkShare) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *FirmwareNetworkShare) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *FirmwareNetworkShare) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


