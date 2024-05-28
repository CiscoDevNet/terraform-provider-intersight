# KubernetesBaremetalNodeProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "kubernetes.BaremetalNodeProfile"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "kubernetes.BaremetalNodeProfile"]
**KubernetesNic** | Pointer to **string** | Network interface from NetworkInfo (by name) to use for kubernetes VIP. | [optional] 
**NetworkInfo** | Pointer to [**NullableKubernetesBaremetalNetworkInfo**](KubernetesBaremetalNetworkInfo.md) |  | [optional] 
**Server** | Pointer to [**NullableComputeRackUnitRelationship**](ComputeRackUnitRelationship.md) |  | [optional] 

## Methods

### NewKubernetesBaremetalNodeProfile

`func NewKubernetesBaremetalNodeProfile(classId string, objectType string, ) *KubernetesBaremetalNodeProfile`

NewKubernetesBaremetalNodeProfile instantiates a new KubernetesBaremetalNodeProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesBaremetalNodeProfileWithDefaults

`func NewKubernetesBaremetalNodeProfileWithDefaults() *KubernetesBaremetalNodeProfile`

NewKubernetesBaremetalNodeProfileWithDefaults instantiates a new KubernetesBaremetalNodeProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *KubernetesBaremetalNodeProfile) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *KubernetesBaremetalNodeProfile) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *KubernetesBaremetalNodeProfile) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *KubernetesBaremetalNodeProfile) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *KubernetesBaremetalNodeProfile) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *KubernetesBaremetalNodeProfile) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetKubernetesNic

`func (o *KubernetesBaremetalNodeProfile) GetKubernetesNic() string`

GetKubernetesNic returns the KubernetesNic field if non-nil, zero value otherwise.

### GetKubernetesNicOk

`func (o *KubernetesBaremetalNodeProfile) GetKubernetesNicOk() (*string, bool)`

GetKubernetesNicOk returns a tuple with the KubernetesNic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetesNic

`func (o *KubernetesBaremetalNodeProfile) SetKubernetesNic(v string)`

SetKubernetesNic sets KubernetesNic field to given value.

### HasKubernetesNic

`func (o *KubernetesBaremetalNodeProfile) HasKubernetesNic() bool`

HasKubernetesNic returns a boolean if a field has been set.

### GetNetworkInfo

`func (o *KubernetesBaremetalNodeProfile) GetNetworkInfo() KubernetesBaremetalNetworkInfo`

GetNetworkInfo returns the NetworkInfo field if non-nil, zero value otherwise.

### GetNetworkInfoOk

`func (o *KubernetesBaremetalNodeProfile) GetNetworkInfoOk() (*KubernetesBaremetalNetworkInfo, bool)`

GetNetworkInfoOk returns a tuple with the NetworkInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkInfo

`func (o *KubernetesBaremetalNodeProfile) SetNetworkInfo(v KubernetesBaremetalNetworkInfo)`

SetNetworkInfo sets NetworkInfo field to given value.

### HasNetworkInfo

`func (o *KubernetesBaremetalNodeProfile) HasNetworkInfo() bool`

HasNetworkInfo returns a boolean if a field has been set.

### SetNetworkInfoNil

`func (o *KubernetesBaremetalNodeProfile) SetNetworkInfoNil(b bool)`

 SetNetworkInfoNil sets the value for NetworkInfo to be an explicit nil

### UnsetNetworkInfo
`func (o *KubernetesBaremetalNodeProfile) UnsetNetworkInfo()`

UnsetNetworkInfo ensures that no value is present for NetworkInfo, not even an explicit nil
### GetServer

`func (o *KubernetesBaremetalNodeProfile) GetServer() ComputeRackUnitRelationship`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *KubernetesBaremetalNodeProfile) GetServerOk() (*ComputeRackUnitRelationship, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *KubernetesBaremetalNodeProfile) SetServer(v ComputeRackUnitRelationship)`

SetServer sets Server field to given value.

### HasServer

`func (o *KubernetesBaremetalNodeProfile) HasServer() bool`

HasServer returns a boolean if a field has been set.

### SetServerNil

`func (o *KubernetesBaremetalNodeProfile) SetServerNil(b bool)`

 SetServerNil sets the value for Server to be an explicit nil

### UnsetServer
`func (o *KubernetesBaremetalNodeProfile) UnsetServer()`

UnsetServer ensures that no value is present for Server, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


