# AssetDeviceStatistics

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

### NewAssetDeviceStatistics

`func NewAssetDeviceStatistics(classId string, objectType string, ) *AssetDeviceStatistics`

NewAssetDeviceStatistics instantiates a new AssetDeviceStatistics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetDeviceStatisticsWithDefaults

`func NewAssetDeviceStatisticsWithDefaults() *AssetDeviceStatistics`

NewAssetDeviceStatisticsWithDefaults instantiates a new AssetDeviceStatistics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *AssetDeviceStatistics) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *AssetDeviceStatistics) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *AssetDeviceStatistics) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *AssetDeviceStatistics) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *AssetDeviceStatistics) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *AssetDeviceStatistics) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetClusterName

`func (o *AssetDeviceStatistics) GetClusterName() string`

GetClusterName returns the ClusterName field if non-nil, zero value otherwise.

### GetClusterNameOk

`func (o *AssetDeviceStatistics) GetClusterNameOk() (*string, bool)`

GetClusterNameOk returns a tuple with the ClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterName

`func (o *AssetDeviceStatistics) SetClusterName(v string)`

SetClusterName sets ClusterName field to given value.

### HasClusterName

`func (o *AssetDeviceStatistics) HasClusterName() bool`

HasClusterName returns a boolean if a field has been set.

### GetConnected

`func (o *AssetDeviceStatistics) GetConnected() int64`

GetConnected returns the Connected field if non-nil, zero value otherwise.

### GetConnectedOk

`func (o *AssetDeviceStatistics) GetConnectedOk() (*int64, bool)`

GetConnectedOk returns a tuple with the Connected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnected

`func (o *AssetDeviceStatistics) SetConnected(v int64)`

SetConnected sets Connected field to given value.

### HasConnected

`func (o *AssetDeviceStatistics) HasConnected() bool`

HasConnected returns a boolean if a field has been set.

### GetMembershipRatio

`func (o *AssetDeviceStatistics) GetMembershipRatio() float32`

GetMembershipRatio returns the MembershipRatio field if non-nil, zero value otherwise.

### GetMembershipRatioOk

`func (o *AssetDeviceStatistics) GetMembershipRatioOk() (*float32, bool)`

GetMembershipRatioOk returns a tuple with the MembershipRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembershipRatio

`func (o *AssetDeviceStatistics) SetMembershipRatio(v float32)`

SetMembershipRatio sets MembershipRatio field to given value.

### HasMembershipRatio

`func (o *AssetDeviceStatistics) HasMembershipRatio() bool`

HasMembershipRatio returns a boolean if a field has been set.

### GetMemoryMirroringFactor

`func (o *AssetDeviceStatistics) GetMemoryMirroringFactor() float32`

GetMemoryMirroringFactor returns the MemoryMirroringFactor field if non-nil, zero value otherwise.

### GetMemoryMirroringFactorOk

`func (o *AssetDeviceStatistics) GetMemoryMirroringFactorOk() (*float32, bool)`

GetMemoryMirroringFactorOk returns a tuple with the MemoryMirroringFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryMirroringFactor

`func (o *AssetDeviceStatistics) SetMemoryMirroringFactor(v float32)`

SetMemoryMirroringFactor sets MemoryMirroringFactor field to given value.

### HasMemoryMirroringFactor

`func (o *AssetDeviceStatistics) HasMemoryMirroringFactor() bool`

HasMemoryMirroringFactor returns a boolean if a field has been set.

### GetVmHost

`func (o *AssetDeviceStatistics) GetVmHost() AssetVmHost`

GetVmHost returns the VmHost field if non-nil, zero value otherwise.

### GetVmHostOk

`func (o *AssetDeviceStatistics) GetVmHostOk() (*AssetVmHost, bool)`

GetVmHostOk returns a tuple with the VmHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmHost

`func (o *AssetDeviceStatistics) SetVmHost(v AssetVmHost)`

SetVmHost sets VmHost field to given value.

### HasVmHost

`func (o *AssetDeviceStatistics) HasVmHost() bool`

HasVmHost returns a boolean if a field has been set.

### SetVmHostNil

`func (o *AssetDeviceStatistics) SetVmHostNil(b bool)`

 SetVmHostNil sets the value for VmHost to be an explicit nil

### UnsetVmHost
`func (o *AssetDeviceStatistics) UnsetVmHost()`

UnsetVmHost ensures that no value is present for VmHost, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


