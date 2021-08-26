# KubernetesVirtualMachineInfrastructureProviderAllOf

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

### NewKubernetesVirtualMachineInfrastructureProviderAllOf

`func NewKubernetesVirtualMachineInfrastructureProviderAllOf(classId string, objectType string, ) *KubernetesVirtualMachineInfrastructureProviderAllOf`

NewKubernetesVirtualMachineInfrastructureProviderAllOf instantiates a new KubernetesVirtualMachineInfrastructureProviderAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesVirtualMachineInfrastructureProviderAllOfWithDefaults

`func NewKubernetesVirtualMachineInfrastructureProviderAllOfWithDefaults() *KubernetesVirtualMachineInfrastructureProviderAllOf`

NewKubernetesVirtualMachineInfrastructureProviderAllOfWithDefaults instantiates a new KubernetesVirtualMachineInfrastructureProviderAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *KubernetesVirtualMachineInfrastructureProviderAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *KubernetesVirtualMachineInfrastructureProviderAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *KubernetesVirtualMachineInfrastructureProviderAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *KubernetesVirtualMachineInfrastructureProviderAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *KubernetesVirtualMachineInfrastructureProviderAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *KubernetesVirtualMachineInfrastructureProviderAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetInfraConfig

`func (o *KubernetesVirtualMachineInfrastructureProviderAllOf) GetInfraConfig() KubernetesBaseVirtualMachineInfraConfig`

GetInfraConfig returns the InfraConfig field if non-nil, zero value otherwise.

### GetInfraConfigOk

`func (o *KubernetesVirtualMachineInfrastructureProviderAllOf) GetInfraConfigOk() (*KubernetesBaseVirtualMachineInfraConfig, bool)`

GetInfraConfigOk returns a tuple with the InfraConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfraConfig

`func (o *KubernetesVirtualMachineInfrastructureProviderAllOf) SetInfraConfig(v KubernetesBaseVirtualMachineInfraConfig)`

SetInfraConfig sets InfraConfig field to given value.

### HasInfraConfig

`func (o *KubernetesVirtualMachineInfrastructureProviderAllOf) HasInfraConfig() bool`

HasInfraConfig returns a boolean if a field has been set.

### SetInfraConfigNil

`func (o *KubernetesVirtualMachineInfrastructureProviderAllOf) SetInfraConfigNil(b bool)`

 SetInfraConfigNil sets the value for InfraConfig to be an explicit nil

### UnsetInfraConfig
`func (o *KubernetesVirtualMachineInfrastructureProviderAllOf) UnsetInfraConfig()`

UnsetInfraConfig ensures that no value is present for InfraConfig, not even an explicit nil
### GetInfraConfigPolicy

`func (o *KubernetesVirtualMachineInfrastructureProviderAllOf) GetInfraConfigPolicy() KubernetesVirtualMachineInfraConfigPolicyRelationship`

GetInfraConfigPolicy returns the InfraConfigPolicy field if non-nil, zero value otherwise.

### GetInfraConfigPolicyOk

`func (o *KubernetesVirtualMachineInfrastructureProviderAllOf) GetInfraConfigPolicyOk() (*KubernetesVirtualMachineInfraConfigPolicyRelationship, bool)`

GetInfraConfigPolicyOk returns a tuple with the InfraConfigPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfraConfigPolicy

`func (o *KubernetesVirtualMachineInfrastructureProviderAllOf) SetInfraConfigPolicy(v KubernetesVirtualMachineInfraConfigPolicyRelationship)`

SetInfraConfigPolicy sets InfraConfigPolicy field to given value.

### HasInfraConfigPolicy

`func (o *KubernetesVirtualMachineInfrastructureProviderAllOf) HasInfraConfigPolicy() bool`

HasInfraConfigPolicy returns a boolean if a field has been set.

### GetInstanceType

`func (o *KubernetesVirtualMachineInfrastructureProviderAllOf) GetInstanceType() KubernetesVirtualMachineInstanceTypeRelationship`

GetInstanceType returns the InstanceType field if non-nil, zero value otherwise.

### GetInstanceTypeOk

`func (o *KubernetesVirtualMachineInfrastructureProviderAllOf) GetInstanceTypeOk() (*KubernetesVirtualMachineInstanceTypeRelationship, bool)`

GetInstanceTypeOk returns a tuple with the InstanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceType

`func (o *KubernetesVirtualMachineInfrastructureProviderAllOf) SetInstanceType(v KubernetesVirtualMachineInstanceTypeRelationship)`

SetInstanceType sets InstanceType field to given value.

### HasInstanceType

`func (o *KubernetesVirtualMachineInfrastructureProviderAllOf) HasInstanceType() bool`

HasInstanceType returns a boolean if a field has been set.

### GetTarget

`func (o *KubernetesVirtualMachineInfrastructureProviderAllOf) GetTarget() AssetDeviceRegistrationRelationship`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *KubernetesVirtualMachineInfrastructureProviderAllOf) GetTargetOk() (*AssetDeviceRegistrationRelationship, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *KubernetesVirtualMachineInfrastructureProviderAllOf) SetTarget(v AssetDeviceRegistrationRelationship)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *KubernetesVirtualMachineInfrastructureProviderAllOf) HasTarget() bool`

HasTarget returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


