# KubernetesBaseVirtualMachineInfraConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**Interfaces** | Pointer to **[]string** |  | [optional] 

## Methods

### NewKubernetesBaseVirtualMachineInfraConfig

`func NewKubernetesBaseVirtualMachineInfraConfig(classId string, objectType string, ) *KubernetesBaseVirtualMachineInfraConfig`

NewKubernetesBaseVirtualMachineInfraConfig instantiates a new KubernetesBaseVirtualMachineInfraConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesBaseVirtualMachineInfraConfigWithDefaults

`func NewKubernetesBaseVirtualMachineInfraConfigWithDefaults() *KubernetesBaseVirtualMachineInfraConfig`

NewKubernetesBaseVirtualMachineInfraConfigWithDefaults instantiates a new KubernetesBaseVirtualMachineInfraConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *KubernetesBaseVirtualMachineInfraConfig) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *KubernetesBaseVirtualMachineInfraConfig) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *KubernetesBaseVirtualMachineInfraConfig) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *KubernetesBaseVirtualMachineInfraConfig) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *KubernetesBaseVirtualMachineInfraConfig) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *KubernetesBaseVirtualMachineInfraConfig) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetInterfaces

`func (o *KubernetesBaseVirtualMachineInfraConfig) GetInterfaces() []string`

GetInterfaces returns the Interfaces field if non-nil, zero value otherwise.

### GetInterfacesOk

`func (o *KubernetesBaseVirtualMachineInfraConfig) GetInterfacesOk() (*[]string, bool)`

GetInterfacesOk returns a tuple with the Interfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaces

`func (o *KubernetesBaseVirtualMachineInfraConfig) SetInterfaces(v []string)`

SetInterfaces sets Interfaces field to given value.

### HasInterfaces

`func (o *KubernetesBaseVirtualMachineInfraConfig) HasInterfaces() bool`

HasInterfaces returns a boolean if a field has been set.

### SetInterfacesNil

`func (o *KubernetesBaseVirtualMachineInfraConfig) SetInterfacesNil(b bool)`

 SetInterfacesNil sets the value for Interfaces to be an explicit nil

### UnsetInterfaces
`func (o *KubernetesBaseVirtualMachineInfraConfig) UnsetInterfaces()`

UnsetInterfaces ensures that no value is present for Interfaces, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


