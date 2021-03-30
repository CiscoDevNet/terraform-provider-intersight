# AssetVmHost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "asset.VmHost"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "asset.VmHost"]
**AccountMoid** | Pointer to **string** | Reference to virtualization target account ID. | [optional] [readonly] 
**ClusterIdentity** | Pointer to **string** | Reference to virtualization cluster identity. | [optional] [readonly] 
**ClusterMoid** | Pointer to **string** | Reference to virtualization cluster ID. | [optional] [readonly] 
**ClusterName** | Pointer to **string** | Reference to virtualization cluster name. | [optional] [readonly] 
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


### GetAccountMoid

`func (o *AssetVmHost) GetAccountMoid() string`

GetAccountMoid returns the AccountMoid field if non-nil, zero value otherwise.

### GetAccountMoidOk

`func (o *AssetVmHost) GetAccountMoidOk() (*string, bool)`

GetAccountMoidOk returns a tuple with the AccountMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountMoid

`func (o *AssetVmHost) SetAccountMoid(v string)`

SetAccountMoid sets AccountMoid field to given value.

### HasAccountMoid

`func (o *AssetVmHost) HasAccountMoid() bool`

HasAccountMoid returns a boolean if a field has been set.

### GetClusterIdentity

`func (o *AssetVmHost) GetClusterIdentity() string`

GetClusterIdentity returns the ClusterIdentity field if non-nil, zero value otherwise.

### GetClusterIdentityOk

`func (o *AssetVmHost) GetClusterIdentityOk() (*string, bool)`

GetClusterIdentityOk returns a tuple with the ClusterIdentity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterIdentity

`func (o *AssetVmHost) SetClusterIdentity(v string)`

SetClusterIdentity sets ClusterIdentity field to given value.

### HasClusterIdentity

`func (o *AssetVmHost) HasClusterIdentity() bool`

HasClusterIdentity returns a boolean if a field has been set.

### GetClusterMoid

`func (o *AssetVmHost) GetClusterMoid() string`

GetClusterMoid returns the ClusterMoid field if non-nil, zero value otherwise.

### GetClusterMoidOk

`func (o *AssetVmHost) GetClusterMoidOk() (*string, bool)`

GetClusterMoidOk returns a tuple with the ClusterMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterMoid

`func (o *AssetVmHost) SetClusterMoid(v string)`

SetClusterMoid sets ClusterMoid field to given value.

### HasClusterMoid

`func (o *AssetVmHost) HasClusterMoid() bool`

HasClusterMoid returns a boolean if a field has been set.

### GetClusterName

`func (o *AssetVmHost) GetClusterName() string`

GetClusterName returns the ClusterName field if non-nil, zero value otherwise.

### GetClusterNameOk

`func (o *AssetVmHost) GetClusterNameOk() (*string, bool)`

GetClusterNameOk returns a tuple with the ClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterName

`func (o *AssetVmHost) SetClusterName(v string)`

SetClusterName sets ClusterName field to given value.

### HasClusterName

`func (o *AssetVmHost) HasClusterName() bool`

HasClusterName returns a boolean if a field has been set.

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


