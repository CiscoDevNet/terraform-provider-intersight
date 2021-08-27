# KubernetesVirtualMachineInfrastructureProvider

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "kubernetes.VirtualMachineInfrastructureProvider"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "kubernetes.VirtualMachineInfrastructureProvider"]
**InfraConfig** | Pointer to [**NullableKubernetesBaseVirtualMachineInfraConfig**](KubernetesBaseVirtualMachineInfraConfig.md) |  | [optional] 
**InfraConfigPolicy** | Pointer to [**KubernetesVirtualMachineInfraConfigPolicyRelationship**](KubernetesVirtualMachineInfraConfigPolicyRelationship.md) |  | [optional] 
**InstanceType** | Pointer to [**KubernetesVirtualMachineInstanceTypeRelationship**](KubernetesVirtualMachineInstanceTypeRelationship.md) |  | [optional] 
**Target** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewKubernetesVirtualMachineInfrastructureProvider

`func NewKubernetesVirtualMachineInfrastructureProvider(classId string, objectType string, ) *KubernetesVirtualMachineInfrastructureProvider`

NewKubernetesVirtualMachineInfrastructureProvider instantiates a new KubernetesVirtualMachineInfrastructureProvider object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesVirtualMachineInfrastructureProviderWithDefaults

`func NewKubernetesVirtualMachineInfrastructureProviderWithDefaults() *KubernetesVirtualMachineInfrastructureProvider`

NewKubernetesVirtualMachineInfrastructureProviderWithDefaults instantiates a new KubernetesVirtualMachineInfrastructureProvider object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *KubernetesVirtualMachineInfrastructureProvider) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *KubernetesVirtualMachineInfrastructureProvider) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *KubernetesVirtualMachineInfrastructureProvider) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *KubernetesVirtualMachineInfrastructureProvider) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *KubernetesVirtualMachineInfrastructureProvider) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *KubernetesVirtualMachineInfrastructureProvider) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetInfraConfig

`func (o *KubernetesVirtualMachineInfrastructureProvider) GetInfraConfig() KubernetesBaseVirtualMachineInfraConfig`

GetInfraConfig returns the InfraConfig field if non-nil, zero value otherwise.

### GetInfraConfigOk

`func (o *KubernetesVirtualMachineInfrastructureProvider) GetInfraConfigOk() (*KubernetesBaseVirtualMachineInfraConfig, bool)`

GetInfraConfigOk returns a tuple with the InfraConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfraConfig

`func (o *KubernetesVirtualMachineInfrastructureProvider) SetInfraConfig(v KubernetesBaseVirtualMachineInfraConfig)`

SetInfraConfig sets InfraConfig field to given value.

### HasInfraConfig

`func (o *KubernetesVirtualMachineInfrastructureProvider) HasInfraConfig() bool`

HasInfraConfig returns a boolean if a field has been set.

### SetInfraConfigNil

`func (o *KubernetesVirtualMachineInfrastructureProvider) SetInfraConfigNil(b bool)`

 SetInfraConfigNil sets the value for InfraConfig to be an explicit nil

### UnsetInfraConfig
`func (o *KubernetesVirtualMachineInfrastructureProvider) UnsetInfraConfig()`

UnsetInfraConfig ensures that no value is present for InfraConfig, not even an explicit nil
### GetInfraConfigPolicy

`func (o *KubernetesVirtualMachineInfrastructureProvider) GetInfraConfigPolicy() KubernetesVirtualMachineInfraConfigPolicyRelationship`

GetInfraConfigPolicy returns the InfraConfigPolicy field if non-nil, zero value otherwise.

### GetInfraConfigPolicyOk

`func (o *KubernetesVirtualMachineInfrastructureProvider) GetInfraConfigPolicyOk() (*KubernetesVirtualMachineInfraConfigPolicyRelationship, bool)`

GetInfraConfigPolicyOk returns a tuple with the InfraConfigPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfraConfigPolicy

`func (o *KubernetesVirtualMachineInfrastructureProvider) SetInfraConfigPolicy(v KubernetesVirtualMachineInfraConfigPolicyRelationship)`

SetInfraConfigPolicy sets InfraConfigPolicy field to given value.

### HasInfraConfigPolicy

`func (o *KubernetesVirtualMachineInfrastructureProvider) HasInfraConfigPolicy() bool`

HasInfraConfigPolicy returns a boolean if a field has been set.

### GetInstanceType

`func (o *KubernetesVirtualMachineInfrastructureProvider) GetInstanceType() KubernetesVirtualMachineInstanceTypeRelationship`

GetInstanceType returns the InstanceType field if non-nil, zero value otherwise.

### GetInstanceTypeOk

`func (o *KubernetesVirtualMachineInfrastructureProvider) GetInstanceTypeOk() (*KubernetesVirtualMachineInstanceTypeRelationship, bool)`

GetInstanceTypeOk returns a tuple with the InstanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceType

`func (o *KubernetesVirtualMachineInfrastructureProvider) SetInstanceType(v KubernetesVirtualMachineInstanceTypeRelationship)`

SetInstanceType sets InstanceType field to given value.

### HasInstanceType

`func (o *KubernetesVirtualMachineInfrastructureProvider) HasInstanceType() bool`

HasInstanceType returns a boolean if a field has been set.

### GetTarget

`func (o *KubernetesVirtualMachineInfrastructureProvider) GetTarget() AssetDeviceRegistrationRelationship`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *KubernetesVirtualMachineInfrastructureProvider) GetTargetOk() (*AssetDeviceRegistrationRelationship, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *KubernetesVirtualMachineInfrastructureProvider) SetTarget(v AssetDeviceRegistrationRelationship)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *KubernetesVirtualMachineInfrastructureProvider) HasTarget() bool`

HasTarget returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


