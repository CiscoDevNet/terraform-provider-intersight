# KubernetesRoute

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "kubernetes.Route"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "kubernetes.Route"]
**To** | Pointer to **string** | The destination subnet, if set to 0.0.0.0/0 then the Route is considered a default route. | [optional] 
**Via** | Pointer to **string** | Via is the gateway for traffic destined for the subnet in the To property. | [optional] 

## Methods

### NewKubernetesRoute

`func NewKubernetesRoute(classId string, objectType string, ) *KubernetesRoute`

NewKubernetesRoute instantiates a new KubernetesRoute object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesRouteWithDefaults

`func NewKubernetesRouteWithDefaults() *KubernetesRoute`

NewKubernetesRouteWithDefaults instantiates a new KubernetesRoute object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *KubernetesRoute) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *KubernetesRoute) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *KubernetesRoute) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *KubernetesRoute) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *KubernetesRoute) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *KubernetesRoute) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetTo

`func (o *KubernetesRoute) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *KubernetesRoute) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *KubernetesRoute) SetTo(v string)`

SetTo sets To field to given value.

### HasTo

`func (o *KubernetesRoute) HasTo() bool`

HasTo returns a boolean if a field has been set.

### GetVia

`func (o *KubernetesRoute) GetVia() string`

GetVia returns the Via field if non-nil, zero value otherwise.

### GetViaOk

`func (o *KubernetesRoute) GetViaOk() (*string, bool)`

GetViaOk returns a tuple with the Via field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVia

`func (o *KubernetesRoute) SetVia(v string)`

SetVia sets Via field to given value.

### HasVia

`func (o *KubernetesRoute) HasVia() bool`

HasVia returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


