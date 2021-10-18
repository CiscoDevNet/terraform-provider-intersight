# KubernetesBaremetalNetworkInfoAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "kubernetes.BaremetalNetworkInfo"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "kubernetes.BaremetalNetworkInfo"]
**Ethernets** | Pointer to [**[]KubernetesEthernet**](KubernetesEthernet.md) |  | [optional] 
**Ovsbonds** | Pointer to [**[]KubernetesOvsBond**](KubernetesOvsBond.md) |  | [optional] 

## Methods

### NewKubernetesBaremetalNetworkInfoAllOf

`func NewKubernetesBaremetalNetworkInfoAllOf(classId string, objectType string, ) *KubernetesBaremetalNetworkInfoAllOf`

NewKubernetesBaremetalNetworkInfoAllOf instantiates a new KubernetesBaremetalNetworkInfoAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesBaremetalNetworkInfoAllOfWithDefaults

`func NewKubernetesBaremetalNetworkInfoAllOfWithDefaults() *KubernetesBaremetalNetworkInfoAllOf`

NewKubernetesBaremetalNetworkInfoAllOfWithDefaults instantiates a new KubernetesBaremetalNetworkInfoAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *KubernetesBaremetalNetworkInfoAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *KubernetesBaremetalNetworkInfoAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *KubernetesBaremetalNetworkInfoAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *KubernetesBaremetalNetworkInfoAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *KubernetesBaremetalNetworkInfoAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *KubernetesBaremetalNetworkInfoAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetEthernets

`func (o *KubernetesBaremetalNetworkInfoAllOf) GetEthernets() []KubernetesEthernet`

GetEthernets returns the Ethernets field if non-nil, zero value otherwise.

### GetEthernetsOk

`func (o *KubernetesBaremetalNetworkInfoAllOf) GetEthernetsOk() (*[]KubernetesEthernet, bool)`

GetEthernetsOk returns a tuple with the Ethernets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEthernets

`func (o *KubernetesBaremetalNetworkInfoAllOf) SetEthernets(v []KubernetesEthernet)`

SetEthernets sets Ethernets field to given value.

### HasEthernets

`func (o *KubernetesBaremetalNetworkInfoAllOf) HasEthernets() bool`

HasEthernets returns a boolean if a field has been set.

### SetEthernetsNil

`func (o *KubernetesBaremetalNetworkInfoAllOf) SetEthernetsNil(b bool)`

 SetEthernetsNil sets the value for Ethernets to be an explicit nil

### UnsetEthernets
`func (o *KubernetesBaremetalNetworkInfoAllOf) UnsetEthernets()`

UnsetEthernets ensures that no value is present for Ethernets, not even an explicit nil
### GetOvsbonds

`func (o *KubernetesBaremetalNetworkInfoAllOf) GetOvsbonds() []KubernetesOvsBond`

GetOvsbonds returns the Ovsbonds field if non-nil, zero value otherwise.

### GetOvsbondsOk

`func (o *KubernetesBaremetalNetworkInfoAllOf) GetOvsbondsOk() (*[]KubernetesOvsBond, bool)`

GetOvsbondsOk returns a tuple with the Ovsbonds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOvsbonds

`func (o *KubernetesBaremetalNetworkInfoAllOf) SetOvsbonds(v []KubernetesOvsBond)`

SetOvsbonds sets Ovsbonds field to given value.

### HasOvsbonds

`func (o *KubernetesBaremetalNetworkInfoAllOf) HasOvsbonds() bool`

HasOvsbonds returns a boolean if a field has been set.

### SetOvsbondsNil

`func (o *KubernetesBaremetalNetworkInfoAllOf) SetOvsbondsNil(b bool)`

 SetOvsbondsNil sets the value for Ovsbonds to be an explicit nil

### UnsetOvsbonds
`func (o *KubernetesBaremetalNetworkInfoAllOf) UnsetOvsbonds()`

UnsetOvsbonds ensures that no value is present for Ovsbonds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


