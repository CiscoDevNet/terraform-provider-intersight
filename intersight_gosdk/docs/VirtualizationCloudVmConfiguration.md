# VirtualizationCloudVmConfiguration

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

### NewVirtualizationCloudVmConfiguration

`func NewVirtualizationCloudVmConfiguration(classId string, objectType string, ) *VirtualizationCloudVmConfiguration`

NewVirtualizationCloudVmConfiguration instantiates a new VirtualizationCloudVmConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualizationCloudVmConfigurationWithDefaults

`func NewVirtualizationCloudVmConfigurationWithDefaults() *VirtualizationCloudVmConfiguration`

NewVirtualizationCloudVmConfigurationWithDefaults instantiates a new VirtualizationCloudVmConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *VirtualizationCloudVmConfiguration) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *VirtualizationCloudVmConfiguration) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *VirtualizationCloudVmConfiguration) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *VirtualizationCloudVmConfiguration) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VirtualizationCloudVmConfiguration) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VirtualizationCloudVmConfiguration) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAvailabilityZoneId

`func (o *VirtualizationCloudVmConfiguration) GetAvailabilityZoneId() string`

GetAvailabilityZoneId returns the AvailabilityZoneId field if non-nil, zero value otherwise.

### GetAvailabilityZoneIdOk

`func (o *VirtualizationCloudVmConfiguration) GetAvailabilityZoneIdOk() (*string, bool)`

GetAvailabilityZoneIdOk returns a tuple with the AvailabilityZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZoneId

`func (o *VirtualizationCloudVmConfiguration) SetAvailabilityZoneId(v string)`

SetAvailabilityZoneId sets AvailabilityZoneId field to given value.

### HasAvailabilityZoneId

`func (o *VirtualizationCloudVmConfiguration) HasAvailabilityZoneId() bool`

HasAvailabilityZoneId returns a boolean if a field has been set.

### GetCompute

`func (o *VirtualizationCloudVmConfiguration) GetCompute() VirtualizationCloudVmComputeConfiguration`

GetCompute returns the Compute field if non-nil, zero value otherwise.

### GetComputeOk

`func (o *VirtualizationCloudVmConfiguration) GetComputeOk() (*VirtualizationCloudVmComputeConfiguration, bool)`

GetComputeOk returns a tuple with the Compute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompute

`func (o *VirtualizationCloudVmConfiguration) SetCompute(v VirtualizationCloudVmComputeConfiguration)`

SetCompute sets Compute field to given value.

### HasCompute

`func (o *VirtualizationCloudVmConfiguration) HasCompute() bool`

HasCompute returns a boolean if a field has been set.

### SetComputeNil

`func (o *VirtualizationCloudVmConfiguration) SetComputeNil(b bool)`

 SetComputeNil sets the value for Compute to be an explicit nil

### UnsetCompute
`func (o *VirtualizationCloudVmConfiguration) UnsetCompute()`

UnsetCompute ensures that no value is present for Compute, not even an explicit nil
### GetImageId

`func (o *VirtualizationCloudVmConfiguration) GetImageId() string`

GetImageId returns the ImageId field if non-nil, zero value otherwise.

### GetImageIdOk

`func (o *VirtualizationCloudVmConfiguration) GetImageIdOk() (*string, bool)`

GetImageIdOk returns a tuple with the ImageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageId

`func (o *VirtualizationCloudVmConfiguration) SetImageId(v string)`

SetImageId sets ImageId field to given value.

### HasImageId

`func (o *VirtualizationCloudVmConfiguration) HasImageId() bool`

HasImageId returns a boolean if a field has been set.

### GetKeyPairName

`func (o *VirtualizationCloudVmConfiguration) GetKeyPairName() string`

GetKeyPairName returns the KeyPairName field if non-nil, zero value otherwise.

### GetKeyPairNameOk

`func (o *VirtualizationCloudVmConfiguration) GetKeyPairNameOk() (*string, bool)`

GetKeyPairNameOk returns a tuple with the KeyPairName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyPairName

`func (o *VirtualizationCloudVmConfiguration) SetKeyPairName(v string)`

SetKeyPairName sets KeyPairName field to given value.

### HasKeyPairName

`func (o *VirtualizationCloudVmConfiguration) HasKeyPairName() bool`

HasKeyPairName returns a boolean if a field has been set.

### GetNetwork

`func (o *VirtualizationCloudVmConfiguration) GetNetwork() VirtualizationCloudVmNetworkConfiguration`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *VirtualizationCloudVmConfiguration) GetNetworkOk() (*VirtualizationCloudVmNetworkConfiguration, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *VirtualizationCloudVmConfiguration) SetNetwork(v VirtualizationCloudVmNetworkConfiguration)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *VirtualizationCloudVmConfiguration) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### SetNetworkNil

`func (o *VirtualizationCloudVmConfiguration) SetNetworkNil(b bool)`

 SetNetworkNil sets the value for Network to be an explicit nil

### UnsetNetwork
`func (o *VirtualizationCloudVmConfiguration) UnsetNetwork()`

UnsetNetwork ensures that no value is present for Network, not even an explicit nil
### GetRegionId

`func (o *VirtualizationCloudVmConfiguration) GetRegionId() string`

GetRegionId returns the RegionId field if non-nil, zero value otherwise.

### GetRegionIdOk

`func (o *VirtualizationCloudVmConfiguration) GetRegionIdOk() (*string, bool)`

GetRegionIdOk returns a tuple with the RegionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionId

`func (o *VirtualizationCloudVmConfiguration) SetRegionId(v string)`

SetRegionId sets RegionId field to given value.

### HasRegionId

`func (o *VirtualizationCloudVmConfiguration) HasRegionId() bool`

HasRegionId returns a boolean if a field has been set.

### GetSecurityGroups

`func (o *VirtualizationCloudVmConfiguration) GetSecurityGroups() []string`

GetSecurityGroups returns the SecurityGroups field if non-nil, zero value otherwise.

### GetSecurityGroupsOk

`func (o *VirtualizationCloudVmConfiguration) GetSecurityGroupsOk() (*[]string, bool)`

GetSecurityGroupsOk returns a tuple with the SecurityGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroups

`func (o *VirtualizationCloudVmConfiguration) SetSecurityGroups(v []string)`

SetSecurityGroups sets SecurityGroups field to given value.

### HasSecurityGroups

`func (o *VirtualizationCloudVmConfiguration) HasSecurityGroups() bool`

HasSecurityGroups returns a boolean if a field has been set.

### SetSecurityGroupsNil

`func (o *VirtualizationCloudVmConfiguration) SetSecurityGroupsNil(b bool)`

 SetSecurityGroupsNil sets the value for SecurityGroups to be an explicit nil

### UnsetSecurityGroups
`func (o *VirtualizationCloudVmConfiguration) UnsetSecurityGroups()`

UnsetSecurityGroups ensures that no value is present for SecurityGroups, not even an explicit nil
### GetStorage

`func (o *VirtualizationCloudVmConfiguration) GetStorage() VirtualizationCloudVmStorageConfiguration`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *VirtualizationCloudVmConfiguration) GetStorageOk() (*VirtualizationCloudVmStorageConfiguration, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *VirtualizationCloudVmConfiguration) SetStorage(v VirtualizationCloudVmStorageConfiguration)`

SetStorage sets Storage field to given value.

### HasStorage

`func (o *VirtualizationCloudVmConfiguration) HasStorage() bool`

HasStorage returns a boolean if a field has been set.

### SetStorageNil

`func (o *VirtualizationCloudVmConfiguration) SetStorageNil(b bool)`

 SetStorageNil sets the value for Storage to be an explicit nil

### UnsetStorage
`func (o *VirtualizationCloudVmConfiguration) UnsetStorage()`

UnsetStorage ensures that no value is present for Storage, not even an explicit nil
### GetVmId

`func (o *VirtualizationCloudVmConfiguration) GetVmId() string`

GetVmId returns the VmId field if non-nil, zero value otherwise.

### GetVmIdOk

`func (o *VirtualizationCloudVmConfiguration) GetVmIdOk() (*string, bool)`

GetVmIdOk returns a tuple with the VmId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmId

`func (o *VirtualizationCloudVmConfiguration) SetVmId(v string)`

SetVmId sets VmId field to given value.

### HasVmId

`func (o *VirtualizationCloudVmConfiguration) HasVmId() bool`

HasVmId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


