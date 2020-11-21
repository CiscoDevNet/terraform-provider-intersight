# FirmwareHttpServerAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "firmware.HttpServer"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "firmware.HttpServer"]
**LocationLink** | Pointer to **string** | HTTP/HTTPS link to the image. Accepted formats HTTP[s]://server-hostname/share/image or HTTP[s]://serverip/share/image. For a successful upgrade, the remote share server must have network connectivity with the CIMC of the selected devices. | [optional] 
**MountOptions** | Pointer to **string** | Mount option as configured on the HTTP server. Empty if nothing is configured. | [optional] 

## Methods

### NewFirmwareHttpServerAllOf

`func NewFirmwareHttpServerAllOf(classId string, objectType string, ) *FirmwareHttpServerAllOf`

NewFirmwareHttpServerAllOf instantiates a new FirmwareHttpServerAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirmwareHttpServerAllOfWithDefaults

`func NewFirmwareHttpServerAllOfWithDefaults() *FirmwareHttpServerAllOf`

NewFirmwareHttpServerAllOfWithDefaults instantiates a new FirmwareHttpServerAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *FirmwareHttpServerAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *FirmwareHttpServerAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *FirmwareHttpServerAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *FirmwareHttpServerAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *FirmwareHttpServerAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *FirmwareHttpServerAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetLocationLink

`func (o *FirmwareHttpServerAllOf) GetLocationLink() string`

GetLocationLink returns the LocationLink field if non-nil, zero value otherwise.

### GetLocationLinkOk

`func (o *FirmwareHttpServerAllOf) GetLocationLinkOk() (*string, bool)`

GetLocationLinkOk returns a tuple with the LocationLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationLink

`func (o *FirmwareHttpServerAllOf) SetLocationLink(v string)`

SetLocationLink sets LocationLink field to given value.

### HasLocationLink

`func (o *FirmwareHttpServerAllOf) HasLocationLink() bool`

HasLocationLink returns a boolean if a field has been set.

### GetMountOptions

`func (o *FirmwareHttpServerAllOf) GetMountOptions() string`

GetMountOptions returns the MountOptions field if non-nil, zero value otherwise.

### GetMountOptionsOk

`func (o *FirmwareHttpServerAllOf) GetMountOptionsOk() (*string, bool)`

GetMountOptionsOk returns a tuple with the MountOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountOptions

`func (o *FirmwareHttpServerAllOf) SetMountOptions(v string)`

SetMountOptions sets MountOptions field to given value.

### HasMountOptions

`func (o *FirmwareHttpServerAllOf) HasMountOptions() bool`

HasMountOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


