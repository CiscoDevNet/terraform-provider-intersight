# AssetSubscriptionDeviceContractInformation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "asset.SubscriptionDeviceContractInformation"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "asset.SubscriptionDeviceContractInformation"]
**DeviceId** | Pointer to **string** | Unique identifier of the Cisco device. | [optional] [readonly] 
**DeviceInformation** | Pointer to [**NullableAssetDeviceInformation**](AssetDeviceInformation.md) |  | [optional] 
**DevicePid** | Pointer to **string** | Product identifier for the specified Cisco device. It is used to distinguish between HyperFlex and UCS devices. | [optional] [readonly] 
**SubscriptionRefId** | Pointer to **string** | Identifies the consumption-based subscription. | [optional] [readonly] 
**DeviceContractInformation** | Pointer to [**AssetDeviceContractInformationRelationship**](AssetDeviceContractInformationRelationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewAssetSubscriptionDeviceContractInformation

`func NewAssetSubscriptionDeviceContractInformation(classId string, objectType string, ) *AssetSubscriptionDeviceContractInformation`

NewAssetSubscriptionDeviceContractInformation instantiates a new AssetSubscriptionDeviceContractInformation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetSubscriptionDeviceContractInformationWithDefaults

`func NewAssetSubscriptionDeviceContractInformationWithDefaults() *AssetSubscriptionDeviceContractInformation`

NewAssetSubscriptionDeviceContractInformationWithDefaults instantiates a new AssetSubscriptionDeviceContractInformation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *AssetSubscriptionDeviceContractInformation) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *AssetSubscriptionDeviceContractInformation) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *AssetSubscriptionDeviceContractInformation) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *AssetSubscriptionDeviceContractInformation) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *AssetSubscriptionDeviceContractInformation) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *AssetSubscriptionDeviceContractInformation) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDeviceId

`func (o *AssetSubscriptionDeviceContractInformation) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *AssetSubscriptionDeviceContractInformation) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *AssetSubscriptionDeviceContractInformation) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *AssetSubscriptionDeviceContractInformation) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### GetDeviceInformation

`func (o *AssetSubscriptionDeviceContractInformation) GetDeviceInformation() AssetDeviceInformation`

GetDeviceInformation returns the DeviceInformation field if non-nil, zero value otherwise.

### GetDeviceInformationOk

`func (o *AssetSubscriptionDeviceContractInformation) GetDeviceInformationOk() (*AssetDeviceInformation, bool)`

GetDeviceInformationOk returns a tuple with the DeviceInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceInformation

`func (o *AssetSubscriptionDeviceContractInformation) SetDeviceInformation(v AssetDeviceInformation)`

SetDeviceInformation sets DeviceInformation field to given value.

### HasDeviceInformation

`func (o *AssetSubscriptionDeviceContractInformation) HasDeviceInformation() bool`

HasDeviceInformation returns a boolean if a field has been set.

### SetDeviceInformationNil

`func (o *AssetSubscriptionDeviceContractInformation) SetDeviceInformationNil(b bool)`

 SetDeviceInformationNil sets the value for DeviceInformation to be an explicit nil

### UnsetDeviceInformation
`func (o *AssetSubscriptionDeviceContractInformation) UnsetDeviceInformation()`

UnsetDeviceInformation ensures that no value is present for DeviceInformation, not even an explicit nil
### GetDevicePid

`func (o *AssetSubscriptionDeviceContractInformation) GetDevicePid() string`

GetDevicePid returns the DevicePid field if non-nil, zero value otherwise.

### GetDevicePidOk

`func (o *AssetSubscriptionDeviceContractInformation) GetDevicePidOk() (*string, bool)`

GetDevicePidOk returns a tuple with the DevicePid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevicePid

`func (o *AssetSubscriptionDeviceContractInformation) SetDevicePid(v string)`

SetDevicePid sets DevicePid field to given value.

### HasDevicePid

`func (o *AssetSubscriptionDeviceContractInformation) HasDevicePid() bool`

HasDevicePid returns a boolean if a field has been set.

### GetSubscriptionRefId

`func (o *AssetSubscriptionDeviceContractInformation) GetSubscriptionRefId() string`

GetSubscriptionRefId returns the SubscriptionRefId field if non-nil, zero value otherwise.

### GetSubscriptionRefIdOk

`func (o *AssetSubscriptionDeviceContractInformation) GetSubscriptionRefIdOk() (*string, bool)`

GetSubscriptionRefIdOk returns a tuple with the SubscriptionRefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionRefId

`func (o *AssetSubscriptionDeviceContractInformation) SetSubscriptionRefId(v string)`

SetSubscriptionRefId sets SubscriptionRefId field to given value.

### HasSubscriptionRefId

`func (o *AssetSubscriptionDeviceContractInformation) HasSubscriptionRefId() bool`

HasSubscriptionRefId returns a boolean if a field has been set.

### GetDeviceContractInformation

`func (o *AssetSubscriptionDeviceContractInformation) GetDeviceContractInformation() AssetDeviceContractInformationRelationship`

GetDeviceContractInformation returns the DeviceContractInformation field if non-nil, zero value otherwise.

### GetDeviceContractInformationOk

`func (o *AssetSubscriptionDeviceContractInformation) GetDeviceContractInformationOk() (*AssetDeviceContractInformationRelationship, bool)`

GetDeviceContractInformationOk returns a tuple with the DeviceContractInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceContractInformation

`func (o *AssetSubscriptionDeviceContractInformation) SetDeviceContractInformation(v AssetDeviceContractInformationRelationship)`

SetDeviceContractInformation sets DeviceContractInformation field to given value.

### HasDeviceContractInformation

`func (o *AssetSubscriptionDeviceContractInformation) HasDeviceContractInformation() bool`

HasDeviceContractInformation returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *AssetSubscriptionDeviceContractInformation) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *AssetSubscriptionDeviceContractInformation) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *AssetSubscriptionDeviceContractInformation) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *AssetSubscriptionDeviceContractInformation) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


