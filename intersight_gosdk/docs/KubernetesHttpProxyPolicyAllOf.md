# KubernetesHttpProxyPolicyAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "kubernetes.HttpProxyPolicy"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "kubernetes.HttpProxyPolicy"]
**HttpProxy** | Pointer to [**NullableKubernetesProxyConfig**](KubernetesProxyConfig.md) |  | [optional] 
**HttpsProxy** | Pointer to [**NullableKubernetesProxyConfig**](KubernetesProxyConfig.md) |  | [optional] 
**NoProxy** | Pointer to **[]string** |  | [optional] 
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](OrganizationOrganizationRelationship.md) |  | [optional] 

## Methods

### NewKubernetesHttpProxyPolicyAllOf

`func NewKubernetesHttpProxyPolicyAllOf(classId string, objectType string, ) *KubernetesHttpProxyPolicyAllOf`

NewKubernetesHttpProxyPolicyAllOf instantiates a new KubernetesHttpProxyPolicyAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesHttpProxyPolicyAllOfWithDefaults

`func NewKubernetesHttpProxyPolicyAllOfWithDefaults() *KubernetesHttpProxyPolicyAllOf`

NewKubernetesHttpProxyPolicyAllOfWithDefaults instantiates a new KubernetesHttpProxyPolicyAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *KubernetesHttpProxyPolicyAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *KubernetesHttpProxyPolicyAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *KubernetesHttpProxyPolicyAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *KubernetesHttpProxyPolicyAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *KubernetesHttpProxyPolicyAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *KubernetesHttpProxyPolicyAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetHttpProxy

`func (o *KubernetesHttpProxyPolicyAllOf) GetHttpProxy() KubernetesProxyConfig`

GetHttpProxy returns the HttpProxy field if non-nil, zero value otherwise.

### GetHttpProxyOk

`func (o *KubernetesHttpProxyPolicyAllOf) GetHttpProxyOk() (*KubernetesProxyConfig, bool)`

GetHttpProxyOk returns a tuple with the HttpProxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpProxy

`func (o *KubernetesHttpProxyPolicyAllOf) SetHttpProxy(v KubernetesProxyConfig)`

SetHttpProxy sets HttpProxy field to given value.

### HasHttpProxy

`func (o *KubernetesHttpProxyPolicyAllOf) HasHttpProxy() bool`

HasHttpProxy returns a boolean if a field has been set.

### SetHttpProxyNil

`func (o *KubernetesHttpProxyPolicyAllOf) SetHttpProxyNil(b bool)`

 SetHttpProxyNil sets the value for HttpProxy to be an explicit nil

### UnsetHttpProxy
`func (o *KubernetesHttpProxyPolicyAllOf) UnsetHttpProxy()`

UnsetHttpProxy ensures that no value is present for HttpProxy, not even an explicit nil
### GetHttpsProxy

`func (o *KubernetesHttpProxyPolicyAllOf) GetHttpsProxy() KubernetesProxyConfig`

GetHttpsProxy returns the HttpsProxy field if non-nil, zero value otherwise.

### GetHttpsProxyOk

`func (o *KubernetesHttpProxyPolicyAllOf) GetHttpsProxyOk() (*KubernetesProxyConfig, bool)`

GetHttpsProxyOk returns a tuple with the HttpsProxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpsProxy

`func (o *KubernetesHttpProxyPolicyAllOf) SetHttpsProxy(v KubernetesProxyConfig)`

SetHttpsProxy sets HttpsProxy field to given value.

### HasHttpsProxy

`func (o *KubernetesHttpProxyPolicyAllOf) HasHttpsProxy() bool`

HasHttpsProxy returns a boolean if a field has been set.

### SetHttpsProxyNil

`func (o *KubernetesHttpProxyPolicyAllOf) SetHttpsProxyNil(b bool)`

 SetHttpsProxyNil sets the value for HttpsProxy to be an explicit nil

### UnsetHttpsProxy
`func (o *KubernetesHttpProxyPolicyAllOf) UnsetHttpsProxy()`

UnsetHttpsProxy ensures that no value is present for HttpsProxy, not even an explicit nil
### GetNoProxy

`func (o *KubernetesHttpProxyPolicyAllOf) GetNoProxy() []string`

GetNoProxy returns the NoProxy field if non-nil, zero value otherwise.

### GetNoProxyOk

`func (o *KubernetesHttpProxyPolicyAllOf) GetNoProxyOk() (*[]string, bool)`

GetNoProxyOk returns a tuple with the NoProxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoProxy

`func (o *KubernetesHttpProxyPolicyAllOf) SetNoProxy(v []string)`

SetNoProxy sets NoProxy field to given value.

### HasNoProxy

`func (o *KubernetesHttpProxyPolicyAllOf) HasNoProxy() bool`

HasNoProxy returns a boolean if a field has been set.

### SetNoProxyNil

`func (o *KubernetesHttpProxyPolicyAllOf) SetNoProxyNil(b bool)`

 SetNoProxyNil sets the value for NoProxy to be an explicit nil

### UnsetNoProxy
`func (o *KubernetesHttpProxyPolicyAllOf) UnsetNoProxy()`

UnsetNoProxy ensures that no value is present for NoProxy, not even an explicit nil
### GetOrganization

`func (o *KubernetesHttpProxyPolicyAllOf) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *KubernetesHttpProxyPolicyAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *KubernetesHttpProxyPolicyAllOf) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *KubernetesHttpProxyPolicyAllOf) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


