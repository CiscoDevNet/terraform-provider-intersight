# KubernetesVirtualMachineInstanceType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "kubernetes.VirtualMachineInstanceType"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "kubernetes.VirtualMachineInstanceType"]
**Cpu** | Pointer to **int64** | Number of CPUs allocated to virtual machine. | [optional] [default to 4]
**DiskSize** | Pointer to **int64** | Ephemeral disk capacity to be provided with units example - 10Gi. | [optional] 
**Memory** | Pointer to **int64** | Virtual machine memory defined in mebibytes (MiB). | [optional] [default to 16384]
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](OrganizationOrganizationRelationship.md) |  | [optional] 
**Profiles** | Pointer to [**[]KubernetesVirtualMachineInfrastructureProviderRelationship**](KubernetesVirtualMachineInfrastructureProviderRelationship.md) | An array of relationships to kubernetesVirtualMachineInfrastructureProvider resources. | [optional] 

## Methods

### NewKubernetesVirtualMachineInstanceType

`func NewKubernetesVirtualMachineInstanceType(classId string, objectType string, ) *KubernetesVirtualMachineInstanceType`

NewKubernetesVirtualMachineInstanceType instantiates a new KubernetesVirtualMachineInstanceType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesVirtualMachineInstanceTypeWithDefaults

`func NewKubernetesVirtualMachineInstanceTypeWithDefaults() *KubernetesVirtualMachineInstanceType`

NewKubernetesVirtualMachineInstanceTypeWithDefaults instantiates a new KubernetesVirtualMachineInstanceType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *KubernetesVirtualMachineInstanceType) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *KubernetesVirtualMachineInstanceType) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *KubernetesVirtualMachineInstanceType) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *KubernetesVirtualMachineInstanceType) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *KubernetesVirtualMachineInstanceType) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *KubernetesVirtualMachineInstanceType) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCpu

`func (o *KubernetesVirtualMachineInstanceType) GetCpu() int64`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *KubernetesVirtualMachineInstanceType) GetCpuOk() (*int64, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *KubernetesVirtualMachineInstanceType) SetCpu(v int64)`

SetCpu sets Cpu field to given value.

### HasCpu

`func (o *KubernetesVirtualMachineInstanceType) HasCpu() bool`

HasCpu returns a boolean if a field has been set.

### GetDiskSize

`func (o *KubernetesVirtualMachineInstanceType) GetDiskSize() int64`

GetDiskSize returns the DiskSize field if non-nil, zero value otherwise.

### GetDiskSizeOk

`func (o *KubernetesVirtualMachineInstanceType) GetDiskSizeOk() (*int64, bool)`

GetDiskSizeOk returns a tuple with the DiskSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskSize

`func (o *KubernetesVirtualMachineInstanceType) SetDiskSize(v int64)`

SetDiskSize sets DiskSize field to given value.

### HasDiskSize

`func (o *KubernetesVirtualMachineInstanceType) HasDiskSize() bool`

HasDiskSize returns a boolean if a field has been set.

### GetMemory

`func (o *KubernetesVirtualMachineInstanceType) GetMemory() int64`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *KubernetesVirtualMachineInstanceType) GetMemoryOk() (*int64, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *KubernetesVirtualMachineInstanceType) SetMemory(v int64)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *KubernetesVirtualMachineInstanceType) HasMemory() bool`

HasMemory returns a boolean if a field has been set.

### GetOrganization

`func (o *KubernetesVirtualMachineInstanceType) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *KubernetesVirtualMachineInstanceType) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *KubernetesVirtualMachineInstanceType) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *KubernetesVirtualMachineInstanceType) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetProfiles

`func (o *KubernetesVirtualMachineInstanceType) GetProfiles() []KubernetesVirtualMachineInfrastructureProviderRelationship`

GetProfiles returns the Profiles field if non-nil, zero value otherwise.

### GetProfilesOk

`func (o *KubernetesVirtualMachineInstanceType) GetProfilesOk() (*[]KubernetesVirtualMachineInfrastructureProviderRelationship, bool)`

GetProfilesOk returns a tuple with the Profiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfiles

`func (o *KubernetesVirtualMachineInstanceType) SetProfiles(v []KubernetesVirtualMachineInfrastructureProviderRelationship)`

SetProfiles sets Profiles field to given value.

### HasProfiles

`func (o *KubernetesVirtualMachineInstanceType) HasProfiles() bool`

HasProfiles returns a boolean if a field has been set.

### SetProfilesNil

`func (o *KubernetesVirtualMachineInstanceType) SetProfilesNil(b bool)`

 SetProfilesNil sets the value for Profiles to be an explicit nil

### UnsetProfiles
`func (o *KubernetesVirtualMachineInstanceType) UnsetProfiles()`

UnsetProfiles ensures that no value is present for Profiles, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


