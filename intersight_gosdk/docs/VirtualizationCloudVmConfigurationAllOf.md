# VirtualizationCloudVmConfigurationAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | [default to "virtualization.AwsVmConfiguration"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | [default to "virtualization.AwsVmConfiguration"]
**AvailabilityZoneId** | Pointer to **string** | Availability zone for the VM. | [optional] 
**Compute** | Pointer to [**NullableVirtualizationCloudVmComputeConfiguration**](VirtualizationCloudVmComputeConfiguration.md) |  | [optional] 
**ImageId** | Pointer to **string** | Virtual machine image used by this VM. | [optional] 
**KeyPairName** | Pointer to **string** | Keypair for accessing the VM. | [optional] 
**Network** | Pointer to [**NullableVirtualizationCloudVmNetworkConfiguration**](VirtualizationCloudVmNetworkConfiguration.md) |  | [optional] 
**RegionId** | Pointer to **string** | Region where the VM instance is created. | [optional] 
**SecurityGroups** | Pointer to **[]string** |  | [optional] 
**Storage** | Pointer to [**NullableVirtualizationCloudVmStorageConfiguration**](VirtualizationCloudVmStorageConfiguration.md) |  | [optional] 
**VmId** | Pointer to **string** | Unique Identifier of the cloud VM. | [optional] 

## Methods

### NewVirtualizationCloudVmConfigurationAllOf

`func NewVirtualizationCloudVmConfigurationAllOf(classId string, objectType string, ) *VirtualizationCloudVmConfigurationAllOf`

NewVirtualizationCloudVmConfigurationAllOf instantiates a new VirtualizationCloudVmConfigurationAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualizationCloudVmConfigurationAllOfWithDefaults

`func NewVirtualizationCloudVmConfigurationAllOfWithDefaults() *VirtualizationCloudVmConfigurationAllOf`

NewVirtualizationCloudVmConfigurationAllOfWithDefaults instantiates a new VirtualizationCloudVmConfigurationAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *VirtualizationCloudVmConfigurationAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *VirtualizationCloudVmConfigurationAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *VirtualizationCloudVmConfigurationAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *VirtualizationCloudVmConfigurationAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VirtualizationCloudVmConfigurationAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VirtualizationCloudVmConfigurationAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAvailabilityZoneId

`func (o *VirtualizationCloudVmConfigurationAllOf) GetAvailabilityZoneId() string`

GetAvailabilityZoneId returns the AvailabilityZoneId field if non-nil, zero value otherwise.

### GetAvailabilityZoneIdOk

`func (o *VirtualizationCloudVmConfigurationAllOf) GetAvailabilityZoneIdOk() (*string, bool)`

GetAvailabilityZoneIdOk returns a tuple with the AvailabilityZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZoneId

`func (o *VirtualizationCloudVmConfigurationAllOf) SetAvailabilityZoneId(v string)`

SetAvailabilityZoneId sets AvailabilityZoneId field to given value.

### HasAvailabilityZoneId

`func (o *VirtualizationCloudVmConfigurationAllOf) HasAvailabilityZoneId() bool`

HasAvailabilityZoneId returns a boolean if a field has been set.

### GetCompute

`func (o *VirtualizationCloudVmConfigurationAllOf) GetCompute() VirtualizationCloudVmComputeConfiguration`

GetCompute returns the Compute field if non-nil, zero value otherwise.

### GetComputeOk

`func (o *VirtualizationCloudVmConfigurationAllOf) GetComputeOk() (*VirtualizationCloudVmComputeConfiguration, bool)`

GetComputeOk returns a tuple with the Compute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompute

`func (o *VirtualizationCloudVmConfigurationAllOf) SetCompute(v VirtualizationCloudVmComputeConfiguration)`

SetCompute sets Compute field to given value.

### HasCompute

`func (o *VirtualizationCloudVmConfigurationAllOf) HasCompute() bool`

HasCompute returns a boolean if a field has been set.

### SetComputeNil

`func (o *VirtualizationCloudVmConfigurationAllOf) SetComputeNil(b bool)`

 SetComputeNil sets the value for Compute to be an explicit nil

### UnsetCompute
`func (o *VirtualizationCloudVmConfigurationAllOf) UnsetCompute()`

UnsetCompute ensures that no value is present for Compute, not even an explicit nil
### GetImageId

`func (o *VirtualizationCloudVmConfigurationAllOf) GetImageId() string`

GetImageId returns the ImageId field if non-nil, zero value otherwise.

### GetImageIdOk

`func (o *VirtualizationCloudVmConfigurationAllOf) GetImageIdOk() (*string, bool)`

GetImageIdOk returns a tuple with the ImageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageId

`func (o *VirtualizationCloudVmConfigurationAllOf) SetImageId(v string)`

SetImageId sets ImageId field to given value.

### HasImageId

`func (o *VirtualizationCloudVmConfigurationAllOf) HasImageId() bool`

HasImageId returns a boolean if a field has been set.

### GetKeyPairName

`func (o *VirtualizationCloudVmConfigurationAllOf) GetKeyPairName() string`

GetKeyPairName returns the KeyPairName field if non-nil, zero value otherwise.

### GetKeyPairNameOk

`func (o *VirtualizationCloudVmConfigurationAllOf) GetKeyPairNameOk() (*string, bool)`

GetKeyPairNameOk returns a tuple with the KeyPairName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyPairName

`func (o *VirtualizationCloudVmConfigurationAllOf) SetKeyPairName(v string)`

SetKeyPairName sets KeyPairName field to given value.

### HasKeyPairName

`func (o *VirtualizationCloudVmConfigurationAllOf) HasKeyPairName() bool`

HasKeyPairName returns a boolean if a field has been set.

### GetNetwork

`func (o *VirtualizationCloudVmConfigurationAllOf) GetNetwork() VirtualizationCloudVmNetworkConfiguration`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *VirtualizationCloudVmConfigurationAllOf) GetNetworkOk() (*VirtualizationCloudVmNetworkConfiguration, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *VirtualizationCloudVmConfigurationAllOf) SetNetwork(v VirtualizationCloudVmNetworkConfiguration)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *VirtualizationCloudVmConfigurationAllOf) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### SetNetworkNil

`func (o *VirtualizationCloudVmConfigurationAllOf) SetNetworkNil(b bool)`

 SetNetworkNil sets the value for Network to be an explicit nil

### UnsetNetwork
`func (o *VirtualizationCloudVmConfigurationAllOf) UnsetNetwork()`

UnsetNetwork ensures that no value is present for Network, not even an explicit nil
### GetRegionId

`func (o *VirtualizationCloudVmConfigurationAllOf) GetRegionId() string`

GetRegionId returns the RegionId field if non-nil, zero value otherwise.

### GetRegionIdOk

`func (o *VirtualizationCloudVmConfigurationAllOf) GetRegionIdOk() (*string, bool)`

GetRegionIdOk returns a tuple with the RegionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionId

`func (o *VirtualizationCloudVmConfigurationAllOf) SetRegionId(v string)`

SetRegionId sets RegionId field to given value.

### HasRegionId

`func (o *VirtualizationCloudVmConfigurationAllOf) HasRegionId() bool`

HasRegionId returns a boolean if a field has been set.

### GetSecurityGroups

`func (o *VirtualizationCloudVmConfigurationAllOf) GetSecurityGroups() []string`

GetSecurityGroups returns the SecurityGroups field if non-nil, zero value otherwise.

### GetSecurityGroupsOk

`func (o *VirtualizationCloudVmConfigurationAllOf) GetSecurityGroupsOk() (*[]string, bool)`

GetSecurityGroupsOk returns a tuple with the SecurityGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroups

`func (o *VirtualizationCloudVmConfigurationAllOf) SetSecurityGroups(v []string)`

SetSecurityGroups sets SecurityGroups field to given value.

### HasSecurityGroups

`func (o *VirtualizationCloudVmConfigurationAllOf) HasSecurityGroups() bool`

HasSecurityGroups returns a boolean if a field has been set.

### SetSecurityGroupsNil

`func (o *VirtualizationCloudVmConfigurationAllOf) SetSecurityGroupsNil(b bool)`

 SetSecurityGroupsNil sets the value for SecurityGroups to be an explicit nil

### UnsetSecurityGroups
`func (o *VirtualizationCloudVmConfigurationAllOf) UnsetSecurityGroups()`

UnsetSecurityGroups ensures that no value is present for SecurityGroups, not even an explicit nil
### GetStorage

`func (o *VirtualizationCloudVmConfigurationAllOf) GetStorage() VirtualizationCloudVmStorageConfiguration`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *VirtualizationCloudVmConfigurationAllOf) GetStorageOk() (*VirtualizationCloudVmStorageConfiguration, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *VirtualizationCloudVmConfigurationAllOf) SetStorage(v VirtualizationCloudVmStorageConfiguration)`

SetStorage sets Storage field to given value.

### HasStorage

`func (o *VirtualizationCloudVmConfigurationAllOf) HasStorage() bool`

HasStorage returns a boolean if a field has been set.

### SetStorageNil

`func (o *VirtualizationCloudVmConfigurationAllOf) SetStorageNil(b bool)`

 SetStorageNil sets the value for Storage to be an explicit nil

### UnsetStorage
`func (o *VirtualizationCloudVmConfigurationAllOf) UnsetStorage()`

UnsetStorage ensures that no value is present for Storage, not even an explicit nil
### GetVmId

`func (o *VirtualizationCloudVmConfigurationAllOf) GetVmId() string`

GetVmId returns the VmId field if non-nil, zero value otherwise.

### GetVmIdOk

`func (o *VirtualizationCloudVmConfigurationAllOf) GetVmIdOk() (*string, bool)`

GetVmIdOk returns a tuple with the VmId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmId

`func (o *VirtualizationCloudVmConfigurationAllOf) SetVmId(v string)`

SetVmId sets VmId field to given value.

### HasVmId

`func (o *VirtualizationCloudVmConfigurationAllOf) HasVmId() bool`

HasVmId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


