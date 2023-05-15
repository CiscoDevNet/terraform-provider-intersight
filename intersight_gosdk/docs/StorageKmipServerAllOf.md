# StorageKmipServerAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "storage.KmipServer"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "storage.KmipServer"]
**EnableDriveSecurity** | Pointer to **bool** | Enable the selected KMIP Server configuration for encryption. This flag just enables the drive security and only after remote key setting configured, the actual encryption will be done. | [optional] 
**IpAddress** | Pointer to **string** | The IP address of the KMIP server. It could be an IPv4 address, an IPv6 address, or a hostname. Hostnames are valid only when Inband is configured for the CIMC address. | [optional] 
**Port** | Pointer to **int64** | The port to which the KMIP client should connect. | [optional] [default to 5696]
**Timeout** | Pointer to **int64** | The timeout before which the KMIP client should connect. | [optional] [default to 60]

## Methods

### NewStorageKmipServerAllOf

`func NewStorageKmipServerAllOf(classId string, objectType string, ) *StorageKmipServerAllOf`

NewStorageKmipServerAllOf instantiates a new StorageKmipServerAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageKmipServerAllOfWithDefaults

`func NewStorageKmipServerAllOfWithDefaults() *StorageKmipServerAllOf`

NewStorageKmipServerAllOfWithDefaults instantiates a new StorageKmipServerAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StorageKmipServerAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageKmipServerAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageKmipServerAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StorageKmipServerAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageKmipServerAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageKmipServerAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetEnableDriveSecurity

`func (o *StorageKmipServerAllOf) GetEnableDriveSecurity() bool`

GetEnableDriveSecurity returns the EnableDriveSecurity field if non-nil, zero value otherwise.

### GetEnableDriveSecurityOk

`func (o *StorageKmipServerAllOf) GetEnableDriveSecurityOk() (*bool, bool)`

GetEnableDriveSecurityOk returns a tuple with the EnableDriveSecurity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableDriveSecurity

`func (o *StorageKmipServerAllOf) SetEnableDriveSecurity(v bool)`

SetEnableDriveSecurity sets EnableDriveSecurity field to given value.

### HasEnableDriveSecurity

`func (o *StorageKmipServerAllOf) HasEnableDriveSecurity() bool`

HasEnableDriveSecurity returns a boolean if a field has been set.

### GetIpAddress

`func (o *StorageKmipServerAllOf) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *StorageKmipServerAllOf) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *StorageKmipServerAllOf) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *StorageKmipServerAllOf) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetPort

`func (o *StorageKmipServerAllOf) GetPort() int64`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *StorageKmipServerAllOf) GetPortOk() (*int64, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *StorageKmipServerAllOf) SetPort(v int64)`

SetPort sets Port field to given value.

### HasPort

`func (o *StorageKmipServerAllOf) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetTimeout

`func (o *StorageKmipServerAllOf) GetTimeout() int64`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *StorageKmipServerAllOf) GetTimeoutOk() (*int64, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *StorageKmipServerAllOf) SetTimeout(v int64)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *StorageKmipServerAllOf) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


