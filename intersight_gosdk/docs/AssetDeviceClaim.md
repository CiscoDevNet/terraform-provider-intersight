# AssetDeviceClaim

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "asset.DeviceClaim"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "asset.DeviceClaim"]
**DeviceUpdates** | Pointer to [**[]AssetConnectionControlMessage**](AssetConnectionControlMessage.md) |  | [optional] 
**SecurityToken** | Pointer to **string** | Obtained from the device connector management UI or API (REST endpoint &#39;/connector/SecurityTokens&#39;). | [optional] 
**SerialNumber** | Pointer to **string** | Obtained from the device connector management UI or API (REST endpoint &#39;/connector/DeviceIdentifiers&#39;). | [optional] 
**Account** | Pointer to [**IamAccountRelationship**](IamAccountRelationship.md) |  | [optional] 
**Device** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewAssetDeviceClaim

`func NewAssetDeviceClaim(classId string, objectType string, ) *AssetDeviceClaim`

NewAssetDeviceClaim instantiates a new AssetDeviceClaim object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetDeviceClaimWithDefaults

`func NewAssetDeviceClaimWithDefaults() *AssetDeviceClaim`

NewAssetDeviceClaimWithDefaults instantiates a new AssetDeviceClaim object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *AssetDeviceClaim) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *AssetDeviceClaim) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *AssetDeviceClaim) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *AssetDeviceClaim) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *AssetDeviceClaim) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *AssetDeviceClaim) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDeviceUpdates

`func (o *AssetDeviceClaim) GetDeviceUpdates() []AssetConnectionControlMessage`

GetDeviceUpdates returns the DeviceUpdates field if non-nil, zero value otherwise.

### GetDeviceUpdatesOk

`func (o *AssetDeviceClaim) GetDeviceUpdatesOk() (*[]AssetConnectionControlMessage, bool)`

GetDeviceUpdatesOk returns a tuple with the DeviceUpdates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceUpdates

`func (o *AssetDeviceClaim) SetDeviceUpdates(v []AssetConnectionControlMessage)`

SetDeviceUpdates sets DeviceUpdates field to given value.

### HasDeviceUpdates

`func (o *AssetDeviceClaim) HasDeviceUpdates() bool`

HasDeviceUpdates returns a boolean if a field has been set.

### SetDeviceUpdatesNil

`func (o *AssetDeviceClaim) SetDeviceUpdatesNil(b bool)`

 SetDeviceUpdatesNil sets the value for DeviceUpdates to be an explicit nil

### UnsetDeviceUpdates
`func (o *AssetDeviceClaim) UnsetDeviceUpdates()`

UnsetDeviceUpdates ensures that no value is present for DeviceUpdates, not even an explicit nil
### GetSecurityToken

`func (o *AssetDeviceClaim) GetSecurityToken() string`

GetSecurityToken returns the SecurityToken field if non-nil, zero value otherwise.

### GetSecurityTokenOk

`func (o *AssetDeviceClaim) GetSecurityTokenOk() (*string, bool)`

GetSecurityTokenOk returns a tuple with the SecurityToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityToken

`func (o *AssetDeviceClaim) SetSecurityToken(v string)`

SetSecurityToken sets SecurityToken field to given value.

### HasSecurityToken

`func (o *AssetDeviceClaim) HasSecurityToken() bool`

HasSecurityToken returns a boolean if a field has been set.

### GetSerialNumber

`func (o *AssetDeviceClaim) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *AssetDeviceClaim) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *AssetDeviceClaim) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *AssetDeviceClaim) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### GetAccount

`func (o *AssetDeviceClaim) GetAccount() IamAccountRelationship`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *AssetDeviceClaim) GetAccountOk() (*IamAccountRelationship, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *AssetDeviceClaim) SetAccount(v IamAccountRelationship)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *AssetDeviceClaim) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetDevice

`func (o *AssetDeviceClaim) GetDevice() AssetDeviceRegistrationRelationship`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *AssetDeviceClaim) GetDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *AssetDeviceClaim) SetDevice(v AssetDeviceRegistrationRelationship)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *AssetDeviceClaim) HasDevice() bool`

HasDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


