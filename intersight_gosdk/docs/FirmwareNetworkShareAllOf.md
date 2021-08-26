# FirmwareNetworkShareAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "firmware.NetworkShare"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "firmware.NetworkShare"]
**CifsServer** | Pointer to [**NullableFirmwareCifsServer**](FirmwareCifsServer.md) |  | [optional] 
**HttpServer** | Pointer to [**NullableFirmwareHttpServer**](FirmwareHttpServer.md) |  | [optional] 
**IsPasswordSet** | Pointer to **bool** | Indicates whether the value of the &#39;password&#39; property has been set. | [optional] [readonly] [default to false]
**MapType** | Pointer to **string** | File server protocols such as CIFS, NFS, WWW for HTTP (S) that hosts the image. * &#x60;nfs&#x60; - File server protocol used is NFS. * &#x60;cifs&#x60; - File server protocol used is CIFS. * &#x60;www&#x60; - File server protocol used is WWW. | [optional] [default to "nfs"]
**NfsServer** | Pointer to [**NullableFirmwareNfsServer**](FirmwareNfsServer.md) |  | [optional] 
**Password** | Pointer to **string** | Password as configured on the file server. | [optional] 
**Upgradeoption** | Pointer to **string** | Option to control the upgrade operation. Some examples, 1) nw_upgrade_mount_only - mount the image from a file server and run the upgrade on the next server boot and 2) nw_upgrade_full - mount the image and immediately run the upgrade. * &#x60;nw_upgrade_full&#x60; - Network upgrade option for full upgrade. * &#x60;nw_upgrade_mount_only&#x60; - Network upgrade mount only. | [optional] [default to "nw_upgrade_full"]
**Username** | Pointer to **string** | Username as configured on the file server. | [optional] 

## Methods

### NewFirmwareNetworkShareAllOf

`func NewFirmwareNetworkShareAllOf(classId string, objectType string, ) *FirmwareNetworkShareAllOf`

NewFirmwareNetworkShareAllOf instantiates a new FirmwareNetworkShareAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirmwareNetworkShareAllOfWithDefaults

`func NewFirmwareNetworkShareAllOfWithDefaults() *FirmwareNetworkShareAllOf`

NewFirmwareNetworkShareAllOfWithDefaults instantiates a new FirmwareNetworkShareAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *FirmwareNetworkShareAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *FirmwareNetworkShareAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *FirmwareNetworkShareAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *FirmwareNetworkShareAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *FirmwareNetworkShareAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *FirmwareNetworkShareAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCifsServer

`func (o *FirmwareNetworkShareAllOf) GetCifsServer() FirmwareCifsServer`

GetCifsServer returns the CifsServer field if non-nil, zero value otherwise.

### GetCifsServerOk

`func (o *FirmwareNetworkShareAllOf) GetCifsServerOk() (*FirmwareCifsServer, bool)`

GetCifsServerOk returns a tuple with the CifsServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCifsServer

`func (o *FirmwareNetworkShareAllOf) SetCifsServer(v FirmwareCifsServer)`

SetCifsServer sets CifsServer field to given value.

### HasCifsServer

`func (o *FirmwareNetworkShareAllOf) HasCifsServer() bool`

HasCifsServer returns a boolean if a field has been set.

### SetCifsServerNil

`func (o *FirmwareNetworkShareAllOf) SetCifsServerNil(b bool)`

 SetCifsServerNil sets the value for CifsServer to be an explicit nil

### UnsetCifsServer
`func (o *FirmwareNetworkShareAllOf) UnsetCifsServer()`

UnsetCifsServer ensures that no value is present for CifsServer, not even an explicit nil
### GetHttpServer

`func (o *FirmwareNetworkShareAllOf) GetHttpServer() FirmwareHttpServer`

GetHttpServer returns the HttpServer field if non-nil, zero value otherwise.

### GetHttpServerOk

`func (o *FirmwareNetworkShareAllOf) GetHttpServerOk() (*FirmwareHttpServer, bool)`

GetHttpServerOk returns a tuple with the HttpServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpServer

`func (o *FirmwareNetworkShareAllOf) SetHttpServer(v FirmwareHttpServer)`

SetHttpServer sets HttpServer field to given value.

### HasHttpServer

`func (o *FirmwareNetworkShareAllOf) HasHttpServer() bool`

HasHttpServer returns a boolean if a field has been set.

### SetHttpServerNil

`func (o *FirmwareNetworkShareAllOf) SetHttpServerNil(b bool)`

 SetHttpServerNil sets the value for HttpServer to be an explicit nil

### UnsetHttpServer
`func (o *FirmwareNetworkShareAllOf) UnsetHttpServer()`

UnsetHttpServer ensures that no value is present for HttpServer, not even an explicit nil
### GetIsPasswordSet

`func (o *FirmwareNetworkShareAllOf) GetIsPasswordSet() bool`

GetIsPasswordSet returns the IsPasswordSet field if non-nil, zero value otherwise.

### GetIsPasswordSetOk

`func (o *FirmwareNetworkShareAllOf) GetIsPasswordSetOk() (*bool, bool)`

GetIsPasswordSetOk returns a tuple with the IsPasswordSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPasswordSet

`func (o *FirmwareNetworkShareAllOf) SetIsPasswordSet(v bool)`

SetIsPasswordSet sets IsPasswordSet field to given value.

### HasIsPasswordSet

`func (o *FirmwareNetworkShareAllOf) HasIsPasswordSet() bool`

HasIsPasswordSet returns a boolean if a field has been set.

### GetMapType

`func (o *FirmwareNetworkShareAllOf) GetMapType() string`

GetMapType returns the MapType field if non-nil, zero value otherwise.

### GetMapTypeOk

`func (o *FirmwareNetworkShareAllOf) GetMapTypeOk() (*string, bool)`

GetMapTypeOk returns a tuple with the MapType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMapType

`func (o *FirmwareNetworkShareAllOf) SetMapType(v string)`

SetMapType sets MapType field to given value.

### HasMapType

`func (o *FirmwareNetworkShareAllOf) HasMapType() bool`

HasMapType returns a boolean if a field has been set.

### GetNfsServer

`func (o *FirmwareNetworkShareAllOf) GetNfsServer() FirmwareNfsServer`

GetNfsServer returns the NfsServer field if non-nil, zero value otherwise.

### GetNfsServerOk

`func (o *FirmwareNetworkShareAllOf) GetNfsServerOk() (*FirmwareNfsServer, bool)`

GetNfsServerOk returns a tuple with the NfsServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfsServer

`func (o *FirmwareNetworkShareAllOf) SetNfsServer(v FirmwareNfsServer)`

SetNfsServer sets NfsServer field to given value.

### HasNfsServer

`func (o *FirmwareNetworkShareAllOf) HasNfsServer() bool`

HasNfsServer returns a boolean if a field has been set.

### SetNfsServerNil

`func (o *FirmwareNetworkShareAllOf) SetNfsServerNil(b bool)`

 SetNfsServerNil sets the value for NfsServer to be an explicit nil

### UnsetNfsServer
`func (o *FirmwareNetworkShareAllOf) UnsetNfsServer()`

UnsetNfsServer ensures that no value is present for NfsServer, not even an explicit nil
### GetPassword

`func (o *FirmwareNetworkShareAllOf) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *FirmwareNetworkShareAllOf) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *FirmwareNetworkShareAllOf) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *FirmwareNetworkShareAllOf) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetUpgradeoption

`func (o *FirmwareNetworkShareAllOf) GetUpgradeoption() string`

GetUpgradeoption returns the Upgradeoption field if non-nil, zero value otherwise.

### GetUpgradeoptionOk

`func (o *FirmwareNetworkShareAllOf) GetUpgradeoptionOk() (*string, bool)`

GetUpgradeoptionOk returns a tuple with the Upgradeoption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradeoption

`func (o *FirmwareNetworkShareAllOf) SetUpgradeoption(v string)`

SetUpgradeoption sets Upgradeoption field to given value.

### HasUpgradeoption

`func (o *FirmwareNetworkShareAllOf) HasUpgradeoption() bool`

HasUpgradeoption returns a boolean if a field has been set.

### GetUsername

`func (o *FirmwareNetworkShareAllOf) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *FirmwareNetworkShareAllOf) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *FirmwareNetworkShareAllOf) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *FirmwareNetworkShareAllOf) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


