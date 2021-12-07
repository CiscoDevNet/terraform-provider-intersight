# KubernetesIpV4Config

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "kubernetes.IpV4Config"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "kubernetes.IpV4Config"]
**Ip** | Pointer to **string** | IPv4 Address in CIDR format. | [optional] 
**Lease** | Pointer to [**MoMoRef**](MoMoRef.md) |  | [optional] 

## Methods

### NewKubernetesIpV4Config

`func NewKubernetesIpV4Config(classId string, objectType string, ) *KubernetesIpV4Config`

NewKubernetesIpV4Config instantiates a new KubernetesIpV4Config object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesIpV4ConfigWithDefaults

`func NewKubernetesIpV4ConfigWithDefaults() *KubernetesIpV4Config`

NewKubernetesIpV4ConfigWithDefaults instantiates a new KubernetesIpV4Config object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *KubernetesIpV4Config) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *KubernetesIpV4Config) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *KubernetesIpV4Config) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *KubernetesIpV4Config) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *KubernetesIpV4Config) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *KubernetesIpV4Config) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetIp

`func (o *KubernetesIpV4Config) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *KubernetesIpV4Config) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *KubernetesIpV4Config) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *KubernetesIpV4Config) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetLease

`func (o *KubernetesIpV4Config) GetLease() MoMoRef`

GetLease returns the Lease field if non-nil, zero value otherwise.

### GetLeaseOk

`func (o *KubernetesIpV4Config) GetLeaseOk() (*MoMoRef, bool)`

GetLeaseOk returns a tuple with the Lease field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLease

`func (o *KubernetesIpV4Config) SetLease(v MoMoRef)`

SetLease sets Lease field to given value.

### HasLease

`func (o *KubernetesIpV4Config) HasLease() bool`

HasLease returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


