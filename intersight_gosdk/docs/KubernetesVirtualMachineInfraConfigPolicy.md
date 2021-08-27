# KubernetesVirtualMachineInfraConfigPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "kubernetes.VirtualMachineInfraConfigPolicy"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "kubernetes.VirtualMachineInfraConfigPolicy"]
**VmConfig** | Pointer to [**NullableKubernetesBaseVirtualMachineInfraConfig**](KubernetesBaseVirtualMachineInfraConfig.md) |  | [optional] 
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](OrganizationOrganizationRelationship.md) |  | [optional] 
**Profiles** | Pointer to [**[]KubernetesVirtualMachineInfrastructureProviderRelationship**](KubernetesVirtualMachineInfrastructureProviderRelationship.md) | An array of relationships to kubernetesVirtualMachineInfrastructureProvider resources. | [optional] 
**Target** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewKubernetesVirtualMachineInfraConfigPolicy

`func NewKubernetesVirtualMachineInfraConfigPolicy(classId string, objectType string, ) *KubernetesVirtualMachineInfraConfigPolicy`

NewKubernetesVirtualMachineInfraConfigPolicy instantiates a new KubernetesVirtualMachineInfraConfigPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesVirtualMachineInfraConfigPolicyWithDefaults

`func NewKubernetesVirtualMachineInfraConfigPolicyWithDefaults() *KubernetesVirtualMachineInfraConfigPolicy`

NewKubernetesVirtualMachineInfraConfigPolicyWithDefaults instantiates a new KubernetesVirtualMachineInfraConfigPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *KubernetesVirtualMachineInfraConfigPolicy) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *KubernetesVirtualMachineInfraConfigPolicy) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *KubernetesVirtualMachineInfraConfigPolicy) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *KubernetesVirtualMachineInfraConfigPolicy) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *KubernetesVirtualMachineInfraConfigPolicy) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *KubernetesVirtualMachineInfraConfigPolicy) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetVmConfig

`func (o *KubernetesVirtualMachineInfraConfigPolicy) GetVmConfig() KubernetesBaseVirtualMachineInfraConfig`

GetVmConfig returns the VmConfig field if non-nil, zero value otherwise.

### GetVmConfigOk

`func (o *KubernetesVirtualMachineInfraConfigPolicy) GetVmConfigOk() (*KubernetesBaseVirtualMachineInfraConfig, bool)`

GetVmConfigOk returns a tuple with the VmConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmConfig

`func (o *KubernetesVirtualMachineInfraConfigPolicy) SetVmConfig(v KubernetesBaseVirtualMachineInfraConfig)`

SetVmConfig sets VmConfig field to given value.

### HasVmConfig

`func (o *KubernetesVirtualMachineInfraConfigPolicy) HasVmConfig() bool`

HasVmConfig returns a boolean if a field has been set.

### SetVmConfigNil

`func (o *KubernetesVirtualMachineInfraConfigPolicy) SetVmConfigNil(b bool)`

 SetVmConfigNil sets the value for VmConfig to be an explicit nil

### UnsetVmConfig
`func (o *KubernetesVirtualMachineInfraConfigPolicy) UnsetVmConfig()`

UnsetVmConfig ensures that no value is present for VmConfig, not even an explicit nil
### GetOrganization

`func (o *KubernetesVirtualMachineInfraConfigPolicy) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *KubernetesVirtualMachineInfraConfigPolicy) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *KubernetesVirtualMachineInfraConfigPolicy) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *KubernetesVirtualMachineInfraConfigPolicy) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetProfiles

`func (o *KubernetesVirtualMachineInfraConfigPolicy) GetProfiles() []KubernetesVirtualMachineInfrastructureProviderRelationship`

GetProfiles returns the Profiles field if non-nil, zero value otherwise.

### GetProfilesOk

`func (o *KubernetesVirtualMachineInfraConfigPolicy) GetProfilesOk() (*[]KubernetesVirtualMachineInfrastructureProviderRelationship, bool)`

GetProfilesOk returns a tuple with the Profiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfiles

`func (o *KubernetesVirtualMachineInfraConfigPolicy) SetProfiles(v []KubernetesVirtualMachineInfrastructureProviderRelationship)`

SetProfiles sets Profiles field to given value.

### HasProfiles

`func (o *KubernetesVirtualMachineInfraConfigPolicy) HasProfiles() bool`

HasProfiles returns a boolean if a field has been set.

### SetProfilesNil

`func (o *KubernetesVirtualMachineInfraConfigPolicy) SetProfilesNil(b bool)`

 SetProfilesNil sets the value for Profiles to be an explicit nil

### UnsetProfiles
`func (o *KubernetesVirtualMachineInfraConfigPolicy) UnsetProfiles()`

UnsetProfiles ensures that no value is present for Profiles, not even an explicit nil
### GetTarget

`func (o *KubernetesVirtualMachineInfraConfigPolicy) GetTarget() AssetDeviceRegistrationRelationship`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *KubernetesVirtualMachineInfraConfigPolicy) GetTargetOk() (*AssetDeviceRegistrationRelationship, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *KubernetesVirtualMachineInfraConfigPolicy) SetTarget(v AssetDeviceRegistrationRelationship)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *KubernetesVirtualMachineInfraConfigPolicy) HasTarget() bool`

HasTarget returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


