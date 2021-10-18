# KubernetesBaremetalNetworkInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "kubernetes.BaremetalNetworkInfo"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "kubernetes.BaremetalNetworkInfo"]
**Ethernets** | Pointer to [**[]KubernetesEthernet**](KubernetesEthernet.md) |  | [optional] 
**Ovsbonds** | Pointer to [**[]KubernetesOvsBond**](KubernetesOvsBond.md) |  | [optional] 

## Methods

### NewKubernetesBaremetalNetworkInfo

`func NewKubernetesBaremetalNetworkInfo(classId string, objectType string, ) *KubernetesBaremetalNetworkInfo`

NewKubernetesBaremetalNetworkInfo instantiates a new KubernetesBaremetalNetworkInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesBaremetalNetworkInfoWithDefaults

`func NewKubernetesBaremetalNetworkInfoWithDefaults() *KubernetesBaremetalNetworkInfo`

NewKubernetesBaremetalNetworkInfoWithDefaults instantiates a new KubernetesBaremetalNetworkInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *KubernetesBaremetalNetworkInfo) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *KubernetesBaremetalNetworkInfo) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *KubernetesBaremetalNetworkInfo) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *KubernetesBaremetalNetworkInfo) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *KubernetesBaremetalNetworkInfo) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *KubernetesBaremetalNetworkInfo) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetEthernets

`func (o *KubernetesBaremetalNetworkInfo) GetEthernets() []KubernetesEthernet`

GetEthernets returns the Ethernets field if non-nil, zero value otherwise.

### GetEthernetsOk

`func (o *KubernetesBaremetalNetworkInfo) GetEthernetsOk() (*[]KubernetesEthernet, bool)`

GetEthernetsOk returns a tuple with the Ethernets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEthernets

`func (o *KubernetesBaremetalNetworkInfo) SetEthernets(v []KubernetesEthernet)`

SetEthernets sets Ethernets field to given value.

### HasEthernets

`func (o *KubernetesBaremetalNetworkInfo) HasEthernets() bool`

HasEthernets returns a boolean if a field has been set.

### SetEthernetsNil

`func (o *KubernetesBaremetalNetworkInfo) SetEthernetsNil(b bool)`

 SetEthernetsNil sets the value for Ethernets to be an explicit nil

### UnsetEthernets
`func (o *KubernetesBaremetalNetworkInfo) UnsetEthernets()`

UnsetEthernets ensures that no value is present for Ethernets, not even an explicit nil
### GetOvsbonds

`func (o *KubernetesBaremetalNetworkInfo) GetOvsbonds() []KubernetesOvsBond`

GetOvsbonds returns the Ovsbonds field if non-nil, zero value otherwise.

### GetOvsbondsOk

`func (o *KubernetesBaremetalNetworkInfo) GetOvsbondsOk() (*[]KubernetesOvsBond, bool)`

GetOvsbondsOk returns a tuple with the Ovsbonds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOvsbonds

`func (o *KubernetesBaremetalNetworkInfo) SetOvsbonds(v []KubernetesOvsBond)`

SetOvsbonds sets Ovsbonds field to given value.

### HasOvsbonds

`func (o *KubernetesBaremetalNetworkInfo) HasOvsbonds() bool`

HasOvsbonds returns a boolean if a field has been set.

### SetOvsbondsNil

`func (o *KubernetesBaremetalNetworkInfo) SetOvsbondsNil(b bool)`

 SetOvsbondsNil sets the value for Ovsbonds to be an explicit nil

### UnsetOvsbonds
`func (o *KubernetesBaremetalNetworkInfo) UnsetOvsbonds()`

UnsetOvsbonds ensures that no value is present for Ovsbonds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


