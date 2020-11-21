# FirmwareHttpServer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "firmware.HttpServer"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "firmware.HttpServer"]
**LocationLink** | Pointer to **string** | HTTP/HTTPS link to the image. Accepted formats HTTP[s]://server-hostname/share/image or HTTP[s]://serverip/share/image. For a successful upgrade, the remote share server must have network connectivity with the CIMC of the selected devices. | [optional] 
**MountOptions** | Pointer to **string** | Mount option as configured on the HTTP server. Empty if nothing is configured. | [optional] 

## Methods

### NewFirmwareHttpServer

`func NewFirmwareHttpServer(classId string, objectType string, ) *FirmwareHttpServer`

NewFirmwareHttpServer instantiates a new FirmwareHttpServer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirmwareHttpServerWithDefaults

`func NewFirmwareHttpServerWithDefaults() *FirmwareHttpServer`

NewFirmwareHttpServerWithDefaults instantiates a new FirmwareHttpServer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *FirmwareHttpServer) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *FirmwareHttpServer) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *FirmwareHttpServer) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *FirmwareHttpServer) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *FirmwareHttpServer) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *FirmwareHttpServer) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetLocationLink

`func (o *FirmwareHttpServer) GetLocationLink() string`

GetLocationLink returns the LocationLink field if non-nil, zero value otherwise.

### GetLocationLinkOk

`func (o *FirmwareHttpServer) GetLocationLinkOk() (*string, bool)`

GetLocationLinkOk returns a tuple with the LocationLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationLink

`func (o *FirmwareHttpServer) SetLocationLink(v string)`

SetLocationLink sets LocationLink field to given value.

### HasLocationLink

`func (o *FirmwareHttpServer) HasLocationLink() bool`

HasLocationLink returns a boolean if a field has been set.

### GetMountOptions

`func (o *FirmwareHttpServer) GetMountOptions() string`

GetMountOptions returns the MountOptions field if non-nil, zero value otherwise.

### GetMountOptionsOk

`func (o *FirmwareHttpServer) GetMountOptionsOk() (*string, bool)`

GetMountOptionsOk returns a tuple with the MountOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountOptions

`func (o *FirmwareHttpServer) SetMountOptions(v string)`

SetMountOptions sets MountOptions field to given value.

### HasMountOptions

`func (o *FirmwareHttpServer) HasMountOptions() bool`

HasMountOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


