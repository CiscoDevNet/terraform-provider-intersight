# KubernetesLoadBalancer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "kubernetes.LoadBalancer"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "kubernetes.LoadBalancer"]
**IpAddresses** | Pointer to **[]string** |  | [optional] 

## Methods

### NewKubernetesLoadBalancer

`func NewKubernetesLoadBalancer(classId string, objectType string, ) *KubernetesLoadBalancer`

NewKubernetesLoadBalancer instantiates a new KubernetesLoadBalancer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesLoadBalancerWithDefaults

`func NewKubernetesLoadBalancerWithDefaults() *KubernetesLoadBalancer`

NewKubernetesLoadBalancerWithDefaults instantiates a new KubernetesLoadBalancer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *KubernetesLoadBalancer) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *KubernetesLoadBalancer) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *KubernetesLoadBalancer) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *KubernetesLoadBalancer) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *KubernetesLoadBalancer) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *KubernetesLoadBalancer) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetIpAddresses

`func (o *KubernetesLoadBalancer) GetIpAddresses() []string`

GetIpAddresses returns the IpAddresses field if non-nil, zero value otherwise.

### GetIpAddressesOk

`func (o *KubernetesLoadBalancer) GetIpAddressesOk() (*[]string, bool)`

GetIpAddressesOk returns a tuple with the IpAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddresses

`func (o *KubernetesLoadBalancer) SetIpAddresses(v []string)`

SetIpAddresses sets IpAddresses field to given value.

### HasIpAddresses

`func (o *KubernetesLoadBalancer) HasIpAddresses() bool`

HasIpAddresses returns a boolean if a field has been set.

### SetIpAddressesNil

`func (o *KubernetesLoadBalancer) SetIpAddressesNil(b bool)`

 SetIpAddressesNil sets the value for IpAddresses to be an explicit nil

### UnsetIpAddresses
`func (o *KubernetesLoadBalancer) UnsetIpAddresses()`

UnsetIpAddresses ensures that no value is present for IpAddresses, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


