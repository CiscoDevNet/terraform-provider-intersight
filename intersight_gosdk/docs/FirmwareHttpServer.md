# FirmwareHttpServer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LocationLink** | Pointer to **string** | HTTP/HTTPS link to the image. Accepted formats HTTP[s]://server-hostname/share/image or HTTP[s]://serverip/share/image. For a successful upgrade, the remote share server must have network connectivity with the CIMC of the selected devices. | [optional] 
**MountOptions** | Pointer to **string** | Mount option as configured on the HTTP server. Empty if nothing is configured. | [optional] 

## Methods

### NewFirmwareHttpServer

`func NewFirmwareHttpServer() *FirmwareHttpServer`

NewFirmwareHttpServer instantiates a new FirmwareHttpServer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirmwareHttpServerWithDefaults

`func NewFirmwareHttpServerWithDefaults() *FirmwareHttpServer`

NewFirmwareHttpServerWithDefaults instantiates a new FirmwareHttpServer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

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


