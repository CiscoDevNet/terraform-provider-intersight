# AssetDeviceStatisticsAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "asset.DeviceStatistics"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "asset.DeviceStatistics"]
**ClusterName** | Pointer to **string** | Name of the cluster. It is specified only for HyperFlex based devices. | [optional] [readonly] 
**Connected** | Pointer to **int64** | The status of the persistent connection between the device connector and Intersight, for HyperFlex or UCS device. 1 represents being connected and 0 represents being disconnected. | [optional] [readonly] 
**MembershipRatio** | Pointer to **float32** | Defines the average proportion of resources used by the device within the cluster. example in a cluster having 3 nodes, the membershipRatio of each node is 1/3 or 0.33. It is specified only for HyperFlex based devices. | [optional] [readonly] 
**MemoryMirroringFactor** | Pointer to **float32** | Memory Reliability, availability and serviceability (RAS) factor. | [optional] [readonly] 
**VmHost** | Pointer to [**NullableAssetVmHost**](asset.VmHost.md) |  | [optional] 

## Methods

### NewAssetDeviceStatisticsAllOf

`func NewAssetDeviceStatisticsAllOf(classId string, objectType string, ) *AssetDeviceStatisticsAllOf`

NewAssetDeviceStatisticsAllOf instantiates a new AssetDeviceStatisticsAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetDeviceStatisticsAllOfWithDefaults

`func NewAssetDeviceStatisticsAllOfWithDefaults() *AssetDeviceStatisticsAllOf`

NewAssetDeviceStatisticsAllOfWithDefaults instantiates a new AssetDeviceStatisticsAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *AssetDeviceStatisticsAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *AssetDeviceStatisticsAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *AssetDeviceStatisticsAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *AssetDeviceStatisticsAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *AssetDeviceStatisticsAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *AssetDeviceStatisticsAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetClusterName

`func (o *AssetDeviceStatisticsAllOf) GetClusterName() string`

GetClusterName returns the ClusterName field if non-nil, zero value otherwise.

### GetClusterNameOk

`func (o *AssetDeviceStatisticsAllOf) GetClusterNameOk() (*string, bool)`

GetClusterNameOk returns a tuple with the ClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterName

`func (o *AssetDeviceStatisticsAllOf) SetClusterName(v string)`

SetClusterName sets ClusterName field to given value.

### HasClusterName

`func (o *AssetDeviceStatisticsAllOf) HasClusterName() bool`

HasClusterName returns a boolean if a field has been set.

### GetConnected

`func (o *AssetDeviceStatisticsAllOf) GetConnected() int64`

GetConnected returns the Connected field if non-nil, zero value otherwise.

### GetConnectedOk

`func (o *AssetDeviceStatisticsAllOf) GetConnectedOk() (*int64, bool)`

GetConnectedOk returns a tuple with the Connected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnected

`func (o *AssetDeviceStatisticsAllOf) SetConnected(v int64)`

SetConnected sets Connected field to given value.

### HasConnected

`func (o *AssetDeviceStatisticsAllOf) HasConnected() bool`

HasConnected returns a boolean if a field has been set.

### GetMembershipRatio

`func (o *AssetDeviceStatisticsAllOf) GetMembershipRatio() float32`

GetMembershipRatio returns the MembershipRatio field if non-nil, zero value otherwise.

### GetMembershipRatioOk

`func (o *AssetDeviceStatisticsAllOf) GetMembershipRatioOk() (*float32, bool)`

GetMembershipRatioOk returns a tuple with the MembershipRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembershipRatio

`func (o *AssetDeviceStatisticsAllOf) SetMembershipRatio(v float32)`

SetMembershipRatio sets MembershipRatio field to given value.

### HasMembershipRatio

`func (o *AssetDeviceStatisticsAllOf) HasMembershipRatio() bool`

HasMembershipRatio returns a boolean if a field has been set.

### GetMemoryMirroringFactor

`func (o *AssetDeviceStatisticsAllOf) GetMemoryMirroringFactor() float32`

GetMemoryMirroringFactor returns the MemoryMirroringFactor field if non-nil, zero value otherwise.

### GetMemoryMirroringFactorOk

`func (o *AssetDeviceStatisticsAllOf) GetMemoryMirroringFactorOk() (*float32, bool)`

GetMemoryMirroringFactorOk returns a tuple with the MemoryMirroringFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryMirroringFactor

`func (o *AssetDeviceStatisticsAllOf) SetMemoryMirroringFactor(v float32)`

SetMemoryMirroringFactor sets MemoryMirroringFactor field to given value.

### HasMemoryMirroringFactor

`func (o *AssetDeviceStatisticsAllOf) HasMemoryMirroringFactor() bool`

HasMemoryMirroringFactor returns a boolean if a field has been set.

### GetVmHost

`func (o *AssetDeviceStatisticsAllOf) GetVmHost() AssetVmHost`

GetVmHost returns the VmHost field if non-nil, zero value otherwise.

### GetVmHostOk

`func (o *AssetDeviceStatisticsAllOf) GetVmHostOk() (*AssetVmHost, bool)`

GetVmHostOk returns a tuple with the VmHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmHost

`func (o *AssetDeviceStatisticsAllOf) SetVmHost(v AssetVmHost)`

SetVmHost sets VmHost field to given value.

### HasVmHost

`func (o *AssetDeviceStatisticsAllOf) HasVmHost() bool`

HasVmHost returns a boolean if a field has been set.

### SetVmHostNil

`func (o *AssetDeviceStatisticsAllOf) SetVmHostNil(b bool)`

 SetVmHostNil sets the value for VmHost to be an explicit nil

### UnsetVmHost
`func (o *AssetDeviceStatisticsAllOf) UnsetVmHost()`

UnsetVmHost ensures that no value is present for VmHost, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


