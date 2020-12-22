# KubernetesNodeSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "kubernetes.NodeSpec"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "kubernetes.NodeSpec"]
**PodCidr** | Pointer to **string** | Node Pod CIDR. In Kubernetes, the workload (Pod) is allocated to an IP address by Kubernetes. The IP address is from a Pod CIDR of the cluster. Each node will (mostly) evenly be distributed the Pod CIDR. | [optional] 
**ProviderId** | Pointer to **string** | Kubernetes can be running on a specific cloud provider such as Openstack, Amazon Web Services, vCenter etc. Each cloud provider will have a specific cloud provider ID. This field is to represent that ID for the corresponding Kubernetes cluster. | [optional] 

## Methods

### NewKubernetesNodeSpec

`func NewKubernetesNodeSpec(classId string, objectType string, ) *KubernetesNodeSpec`

NewKubernetesNodeSpec instantiates a new KubernetesNodeSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesNodeSpecWithDefaults

`func NewKubernetesNodeSpecWithDefaults() *KubernetesNodeSpec`

NewKubernetesNodeSpecWithDefaults instantiates a new KubernetesNodeSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *KubernetesNodeSpec) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *KubernetesNodeSpec) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *KubernetesNodeSpec) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *KubernetesNodeSpec) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *KubernetesNodeSpec) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *KubernetesNodeSpec) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetPodCidr

`func (o *KubernetesNodeSpec) GetPodCidr() string`

GetPodCidr returns the PodCidr field if non-nil, zero value otherwise.

### GetPodCidrOk

`func (o *KubernetesNodeSpec) GetPodCidrOk() (*string, bool)`

GetPodCidrOk returns a tuple with the PodCidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodCidr

`func (o *KubernetesNodeSpec) SetPodCidr(v string)`

SetPodCidr sets PodCidr field to given value.

### HasPodCidr

`func (o *KubernetesNodeSpec) HasPodCidr() bool`

HasPodCidr returns a boolean if a field has been set.

### GetProviderId

`func (o *KubernetesNodeSpec) GetProviderId() string`

GetProviderId returns the ProviderId field if non-nil, zero value otherwise.

### GetProviderIdOk

`func (o *KubernetesNodeSpec) GetProviderIdOk() (*string, bool)`

GetProviderIdOk returns a tuple with the ProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderId

`func (o *KubernetesNodeSpec) SetProviderId(v string)`

SetProviderId sets ProviderId field to given value.

### HasProviderId

`func (o *KubernetesNodeSpec) HasProviderId() bool`

HasProviderId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


