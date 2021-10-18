# KubernetesHyperFlexApVirtualMachineInfraConfigAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "kubernetes.HyperFlexApVirtualMachineInfraConfig"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "kubernetes.HyperFlexApVirtualMachineInfraConfig"]
**DiskMode** | Pointer to **string** | Disk mode to use for volumes. * &#x60;Block&#x60; - It is a Block virtual disk. * &#x60;Filesystem&#x60; - It is a File system virtual disk. * &#x60;&#x60; - Disk mode is either unknown or not supported. | [optional] [default to "Block"]

## Methods

### NewKubernetesHyperFlexApVirtualMachineInfraConfigAllOf

`func NewKubernetesHyperFlexApVirtualMachineInfraConfigAllOf(classId string, objectType string, ) *KubernetesHyperFlexApVirtualMachineInfraConfigAllOf`

NewKubernetesHyperFlexApVirtualMachineInfraConfigAllOf instantiates a new KubernetesHyperFlexApVirtualMachineInfraConfigAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesHyperFlexApVirtualMachineInfraConfigAllOfWithDefaults

`func NewKubernetesHyperFlexApVirtualMachineInfraConfigAllOfWithDefaults() *KubernetesHyperFlexApVirtualMachineInfraConfigAllOf`

NewKubernetesHyperFlexApVirtualMachineInfraConfigAllOfWithDefaults instantiates a new KubernetesHyperFlexApVirtualMachineInfraConfigAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *KubernetesHyperFlexApVirtualMachineInfraConfigAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *KubernetesHyperFlexApVirtualMachineInfraConfigAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *KubernetesHyperFlexApVirtualMachineInfraConfigAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *KubernetesHyperFlexApVirtualMachineInfraConfigAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *KubernetesHyperFlexApVirtualMachineInfraConfigAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *KubernetesHyperFlexApVirtualMachineInfraConfigAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDiskMode

`func (o *KubernetesHyperFlexApVirtualMachineInfraConfigAllOf) GetDiskMode() string`

GetDiskMode returns the DiskMode field if non-nil, zero value otherwise.

### GetDiskModeOk

`func (o *KubernetesHyperFlexApVirtualMachineInfraConfigAllOf) GetDiskModeOk() (*string, bool)`

GetDiskModeOk returns a tuple with the DiskMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskMode

`func (o *KubernetesHyperFlexApVirtualMachineInfraConfigAllOf) SetDiskMode(v string)`

SetDiskMode sets DiskMode field to given value.

### HasDiskMode

`func (o *KubernetesHyperFlexApVirtualMachineInfraConfigAllOf) HasDiskMode() bool`

HasDiskMode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


