# KubernetesHttpProxyPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "kubernetes.HttpProxyPolicy"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "kubernetes.HttpProxyPolicy"]
**HttpProxy** | Pointer to [**NullableKubernetesProxyConfig**](KubernetesProxyConfig.md) |  | [optional] 
**HttpsProxy** | Pointer to [**NullableKubernetesProxyConfig**](KubernetesProxyConfig.md) |  | [optional] 
**NoProxy** | Pointer to **[]string** |  | [optional] 
**Organization** | Pointer to [**NullableOrganizationOrganizationRelationship**](OrganizationOrganizationRelationship.md) |  | [optional] 

## Methods

### NewKubernetesHttpProxyPolicy

`func NewKubernetesHttpProxyPolicy(classId string, objectType string, ) *KubernetesHttpProxyPolicy`

NewKubernetesHttpProxyPolicy instantiates a new KubernetesHttpProxyPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesHttpProxyPolicyWithDefaults

`func NewKubernetesHttpProxyPolicyWithDefaults() *KubernetesHttpProxyPolicy`

NewKubernetesHttpProxyPolicyWithDefaults instantiates a new KubernetesHttpProxyPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *KubernetesHttpProxyPolicy) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *KubernetesHttpProxyPolicy) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *KubernetesHttpProxyPolicy) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *KubernetesHttpProxyPolicy) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *KubernetesHttpProxyPolicy) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *KubernetesHttpProxyPolicy) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetHttpProxy

`func (o *KubernetesHttpProxyPolicy) GetHttpProxy() KubernetesProxyConfig`

GetHttpProxy returns the HttpProxy field if non-nil, zero value otherwise.

### GetHttpProxyOk

`func (o *KubernetesHttpProxyPolicy) GetHttpProxyOk() (*KubernetesProxyConfig, bool)`

GetHttpProxyOk returns a tuple with the HttpProxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpProxy

`func (o *KubernetesHttpProxyPolicy) SetHttpProxy(v KubernetesProxyConfig)`

SetHttpProxy sets HttpProxy field to given value.

### HasHttpProxy

`func (o *KubernetesHttpProxyPolicy) HasHttpProxy() bool`

HasHttpProxy returns a boolean if a field has been set.

### SetHttpProxyNil

`func (o *KubernetesHttpProxyPolicy) SetHttpProxyNil(b bool)`

 SetHttpProxyNil sets the value for HttpProxy to be an explicit nil

### UnsetHttpProxy
`func (o *KubernetesHttpProxyPolicy) UnsetHttpProxy()`

UnsetHttpProxy ensures that no value is present for HttpProxy, not even an explicit nil
### GetHttpsProxy

`func (o *KubernetesHttpProxyPolicy) GetHttpsProxy() KubernetesProxyConfig`

GetHttpsProxy returns the HttpsProxy field if non-nil, zero value otherwise.

### GetHttpsProxyOk

`func (o *KubernetesHttpProxyPolicy) GetHttpsProxyOk() (*KubernetesProxyConfig, bool)`

GetHttpsProxyOk returns a tuple with the HttpsProxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpsProxy

`func (o *KubernetesHttpProxyPolicy) SetHttpsProxy(v KubernetesProxyConfig)`

SetHttpsProxy sets HttpsProxy field to given value.

### HasHttpsProxy

`func (o *KubernetesHttpProxyPolicy) HasHttpsProxy() bool`

HasHttpsProxy returns a boolean if a field has been set.

### SetHttpsProxyNil

`func (o *KubernetesHttpProxyPolicy) SetHttpsProxyNil(b bool)`

 SetHttpsProxyNil sets the value for HttpsProxy to be an explicit nil

### UnsetHttpsProxy
`func (o *KubernetesHttpProxyPolicy) UnsetHttpsProxy()`

UnsetHttpsProxy ensures that no value is present for HttpsProxy, not even an explicit nil
### GetNoProxy

`func (o *KubernetesHttpProxyPolicy) GetNoProxy() []string`

GetNoProxy returns the NoProxy field if non-nil, zero value otherwise.

### GetNoProxyOk

`func (o *KubernetesHttpProxyPolicy) GetNoProxyOk() (*[]string, bool)`

GetNoProxyOk returns a tuple with the NoProxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoProxy

`func (o *KubernetesHttpProxyPolicy) SetNoProxy(v []string)`

SetNoProxy sets NoProxy field to given value.

### HasNoProxy

`func (o *KubernetesHttpProxyPolicy) HasNoProxy() bool`

HasNoProxy returns a boolean if a field has been set.

### SetNoProxyNil

`func (o *KubernetesHttpProxyPolicy) SetNoProxyNil(b bool)`

 SetNoProxyNil sets the value for NoProxy to be an explicit nil

### UnsetNoProxy
`func (o *KubernetesHttpProxyPolicy) UnsetNoProxy()`

UnsetNoProxy ensures that no value is present for NoProxy, not even an explicit nil
### GetOrganization

`func (o *KubernetesHttpProxyPolicy) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *KubernetesHttpProxyPolicy) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *KubernetesHttpProxyPolicy) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *KubernetesHttpProxyPolicy) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### SetOrganizationNil

`func (o *KubernetesHttpProxyPolicy) SetOrganizationNil(b bool)`

 SetOrganizationNil sets the value for Organization to be an explicit nil

### UnsetOrganization
`func (o *KubernetesHttpProxyPolicy) UnsetOrganization()`

UnsetOrganization ensures that no value is present for Organization, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


