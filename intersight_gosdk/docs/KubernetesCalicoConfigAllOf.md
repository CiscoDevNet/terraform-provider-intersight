# KubernetesCalicoConfigAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "kubernetes.CalicoConfig"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "kubernetes.CalicoConfig"]
**IpVersion** | Pointer to **string** | IP version that can take on values v4 or v6. * &#x60;v4&#x60; - This refers to the IPv4 address. * &#x60;v6&#x60; - This refers to the IPv6 address. | [optional] [default to "v4"]
**Mtu** | Pointer to **int64** | Workload interface maximum transmission unit (MTU). | [optional] 

## Methods

### NewKubernetesCalicoConfigAllOf

`func NewKubernetesCalicoConfigAllOf(classId string, objectType string, ) *KubernetesCalicoConfigAllOf`

NewKubernetesCalicoConfigAllOf instantiates a new KubernetesCalicoConfigAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesCalicoConfigAllOfWithDefaults

`func NewKubernetesCalicoConfigAllOfWithDefaults() *KubernetesCalicoConfigAllOf`

NewKubernetesCalicoConfigAllOfWithDefaults instantiates a new KubernetesCalicoConfigAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *KubernetesCalicoConfigAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *KubernetesCalicoConfigAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *KubernetesCalicoConfigAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *KubernetesCalicoConfigAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *KubernetesCalicoConfigAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *KubernetesCalicoConfigAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetIpVersion

`func (o *KubernetesCalicoConfigAllOf) GetIpVersion() string`

GetIpVersion returns the IpVersion field if non-nil, zero value otherwise.

### GetIpVersionOk

`func (o *KubernetesCalicoConfigAllOf) GetIpVersionOk() (*string, bool)`

GetIpVersionOk returns a tuple with the IpVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpVersion

`func (o *KubernetesCalicoConfigAllOf) SetIpVersion(v string)`

SetIpVersion sets IpVersion field to given value.

### HasIpVersion

`func (o *KubernetesCalicoConfigAllOf) HasIpVersion() bool`

HasIpVersion returns a boolean if a field has been set.

### GetMtu

`func (o *KubernetesCalicoConfigAllOf) GetMtu() int64`

GetMtu returns the Mtu field if non-nil, zero value otherwise.

### GetMtuOk

`func (o *KubernetesCalicoConfigAllOf) GetMtuOk() (*int64, bool)`

GetMtuOk returns a tuple with the Mtu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtu

`func (o *KubernetesCalicoConfigAllOf) SetMtu(v int64)`

SetMtu sets Mtu field to given value.

### HasMtu

`func (o *KubernetesCalicoConfigAllOf) HasMtu() bool`

HasMtu returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


