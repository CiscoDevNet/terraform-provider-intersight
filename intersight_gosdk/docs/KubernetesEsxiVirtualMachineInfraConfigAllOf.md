# KubernetesEsxiVirtualMachineInfraConfigAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "kubernetes.EsxiVirtualMachineInfraConfig"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "kubernetes.EsxiVirtualMachineInfraConfig"]
**Cluster** | Pointer to **string** | Name of the vSphere cluster on which the virtual machines are created. | [optional] 
**Datastore** | Pointer to **string** | Name of the datasore on which the virtual machine disks are created. | [optional] 
**IsPassphraseSet** | Pointer to **bool** | Indicates whether the value of the &#39;passphrase&#39; property has been set. | [optional] [readonly] [default to false]
**Passphrase** | Pointer to **string** | Passphrase for the vcenter user. | [optional] 
**ResourcePool** | Pointer to **string** | Name of the vSphere resource pool on which the virtual machines are created. | [optional] 

## Methods

### NewKubernetesEsxiVirtualMachineInfraConfigAllOf

`func NewKubernetesEsxiVirtualMachineInfraConfigAllOf(classId string, objectType string, ) *KubernetesEsxiVirtualMachineInfraConfigAllOf`

NewKubernetesEsxiVirtualMachineInfraConfigAllOf instantiates a new KubernetesEsxiVirtualMachineInfraConfigAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesEsxiVirtualMachineInfraConfigAllOfWithDefaults

`func NewKubernetesEsxiVirtualMachineInfraConfigAllOfWithDefaults() *KubernetesEsxiVirtualMachineInfraConfigAllOf`

NewKubernetesEsxiVirtualMachineInfraConfigAllOfWithDefaults instantiates a new KubernetesEsxiVirtualMachineInfraConfigAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *KubernetesEsxiVirtualMachineInfraConfigAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *KubernetesEsxiVirtualMachineInfraConfigAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *KubernetesEsxiVirtualMachineInfraConfigAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *KubernetesEsxiVirtualMachineInfraConfigAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *KubernetesEsxiVirtualMachineInfraConfigAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *KubernetesEsxiVirtualMachineInfraConfigAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCluster

`func (o *KubernetesEsxiVirtualMachineInfraConfigAllOf) GetCluster() string`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *KubernetesEsxiVirtualMachineInfraConfigAllOf) GetClusterOk() (*string, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *KubernetesEsxiVirtualMachineInfraConfigAllOf) SetCluster(v string)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *KubernetesEsxiVirtualMachineInfraConfigAllOf) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetDatastore

`func (o *KubernetesEsxiVirtualMachineInfraConfigAllOf) GetDatastore() string`

GetDatastore returns the Datastore field if non-nil, zero value otherwise.

### GetDatastoreOk

`func (o *KubernetesEsxiVirtualMachineInfraConfigAllOf) GetDatastoreOk() (*string, bool)`

GetDatastoreOk returns a tuple with the Datastore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastore

`func (o *KubernetesEsxiVirtualMachineInfraConfigAllOf) SetDatastore(v string)`

SetDatastore sets Datastore field to given value.

### HasDatastore

`func (o *KubernetesEsxiVirtualMachineInfraConfigAllOf) HasDatastore() bool`

HasDatastore returns a boolean if a field has been set.

### GetIsPassphraseSet

`func (o *KubernetesEsxiVirtualMachineInfraConfigAllOf) GetIsPassphraseSet() bool`

GetIsPassphraseSet returns the IsPassphraseSet field if non-nil, zero value otherwise.

### GetIsPassphraseSetOk

`func (o *KubernetesEsxiVirtualMachineInfraConfigAllOf) GetIsPassphraseSetOk() (*bool, bool)`

GetIsPassphraseSetOk returns a tuple with the IsPassphraseSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPassphraseSet

`func (o *KubernetesEsxiVirtualMachineInfraConfigAllOf) SetIsPassphraseSet(v bool)`

SetIsPassphraseSet sets IsPassphraseSet field to given value.

### HasIsPassphraseSet

`func (o *KubernetesEsxiVirtualMachineInfraConfigAllOf) HasIsPassphraseSet() bool`

HasIsPassphraseSet returns a boolean if a field has been set.

### GetPassphrase

`func (o *KubernetesEsxiVirtualMachineInfraConfigAllOf) GetPassphrase() string`

GetPassphrase returns the Passphrase field if non-nil, zero value otherwise.

### GetPassphraseOk

`func (o *KubernetesEsxiVirtualMachineInfraConfigAllOf) GetPassphraseOk() (*string, bool)`

GetPassphraseOk returns a tuple with the Passphrase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassphrase

`func (o *KubernetesEsxiVirtualMachineInfraConfigAllOf) SetPassphrase(v string)`

SetPassphrase sets Passphrase field to given value.

### HasPassphrase

`func (o *KubernetesEsxiVirtualMachineInfraConfigAllOf) HasPassphrase() bool`

HasPassphrase returns a boolean if a field has been set.

### GetResourcePool

`func (o *KubernetesEsxiVirtualMachineInfraConfigAllOf) GetResourcePool() string`

GetResourcePool returns the ResourcePool field if non-nil, zero value otherwise.

### GetResourcePoolOk

`func (o *KubernetesEsxiVirtualMachineInfraConfigAllOf) GetResourcePoolOk() (*string, bool)`

GetResourcePoolOk returns a tuple with the ResourcePool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePool

`func (o *KubernetesEsxiVirtualMachineInfraConfigAllOf) SetResourcePool(v string)`

SetResourcePool sets ResourcePool field to given value.

### HasResourcePool

`func (o *KubernetesEsxiVirtualMachineInfraConfigAllOf) HasResourcePool() bool`

HasResourcePool returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


