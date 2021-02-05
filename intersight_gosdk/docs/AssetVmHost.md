# AssetVmHost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "asset.VmHost"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "asset.VmHost"]
**Connected** | Pointer to **int64** | The connection status of the host. 1 represents being connected and 0 represents being disconnected. | [optional] [readonly] 
**RegisteredDeviceMoid** | Pointer to **string** | Reference to virtualization target device ID. | [optional] [readonly] 

## Methods

### NewAssetVmHost

`func NewAssetVmHost(classId string, objectType string, ) *AssetVmHost`

NewAssetVmHost instantiates a new AssetVmHost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetVmHostWithDefaults

`func NewAssetVmHostWithDefaults() *AssetVmHost`

NewAssetVmHostWithDefaults instantiates a new AssetVmHost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *AssetVmHost) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *AssetVmHost) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *AssetVmHost) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *AssetVmHost) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *AssetVmHost) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *AssetVmHost) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetConnected

`func (o *AssetVmHost) GetConnected() int64`

GetConnected returns the Connected field if non-nil, zero value otherwise.

### GetConnectedOk

`func (o *AssetVmHost) GetConnectedOk() (*int64, bool)`

GetConnectedOk returns a tuple with the Connected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnected

`func (o *AssetVmHost) SetConnected(v int64)`

SetConnected sets Connected field to given value.

### HasConnected

`func (o *AssetVmHost) HasConnected() bool`

HasConnected returns a boolean if a field has been set.

### GetRegisteredDeviceMoid

`func (o *AssetVmHost) GetRegisteredDeviceMoid() string`

GetRegisteredDeviceMoid returns the RegisteredDeviceMoid field if non-nil, zero value otherwise.

### GetRegisteredDeviceMoidOk

`func (o *AssetVmHost) GetRegisteredDeviceMoidOk() (*string, bool)`

GetRegisteredDeviceMoidOk returns a tuple with the RegisteredDeviceMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDeviceMoid

`func (o *AssetVmHost) SetRegisteredDeviceMoid(v string)`

SetRegisteredDeviceMoid sets RegisteredDeviceMoid field to given value.

### HasRegisteredDeviceMoid

`func (o *AssetVmHost) HasRegisteredDeviceMoid() bool`

HasRegisteredDeviceMoid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


