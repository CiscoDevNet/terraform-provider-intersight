# KubernetesIpV4ConfigAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "kubernetes.IpV4Config"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "kubernetes.IpV4Config"]
**Ip** | Pointer to **string** | IPv4 Address in CIDR format. | [optional] 
**Lease** | Pointer to [**MoMoRef**](MoMoRef.md) |  | [optional] 

## Methods

### NewKubernetesIpV4ConfigAllOf

`func NewKubernetesIpV4ConfigAllOf(classId string, objectType string, ) *KubernetesIpV4ConfigAllOf`

NewKubernetesIpV4ConfigAllOf instantiates a new KubernetesIpV4ConfigAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesIpV4ConfigAllOfWithDefaults

`func NewKubernetesIpV4ConfigAllOfWithDefaults() *KubernetesIpV4ConfigAllOf`

NewKubernetesIpV4ConfigAllOfWithDefaults instantiates a new KubernetesIpV4ConfigAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *KubernetesIpV4ConfigAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *KubernetesIpV4ConfigAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *KubernetesIpV4ConfigAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *KubernetesIpV4ConfigAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *KubernetesIpV4ConfigAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *KubernetesIpV4ConfigAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetIp

`func (o *KubernetesIpV4ConfigAllOf) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *KubernetesIpV4ConfigAllOf) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *KubernetesIpV4ConfigAllOf) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *KubernetesIpV4ConfigAllOf) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetLease

`func (o *KubernetesIpV4ConfigAllOf) GetLease() MoMoRef`

GetLease returns the Lease field if non-nil, zero value otherwise.

### GetLeaseOk

`func (o *KubernetesIpV4ConfigAllOf) GetLeaseOk() (*MoMoRef, bool)`

GetLeaseOk returns a tuple with the Lease field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLease

`func (o *KubernetesIpV4ConfigAllOf) SetLease(v MoMoRef)`

SetLease sets Lease field to given value.

### HasLease

`func (o *KubernetesIpV4ConfigAllOf) HasLease() bool`

HasLease returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


