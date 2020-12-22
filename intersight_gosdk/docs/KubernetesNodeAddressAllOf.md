# KubernetesNodeAddressAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "kubernetes.NodeAddress"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "kubernetes.NodeAddress"]
**Address** | Pointer to **string** | The address of type specified by the type field. | [optional] 
**Type** | Pointer to **string** | The address type of the Node. The usage of the IP address depending on the cloud provider or bare metal configuration. HostName - The hostname as reported by the node&#39;s kernel. ExternalIP - Typically the IP address of the node that is externally routable (available from outside the cluster) InternalIP - Typically the IP address of the node that is routable only within the cluster. | [optional] 

## Methods

### NewKubernetesNodeAddressAllOf

`func NewKubernetesNodeAddressAllOf(classId string, objectType string, ) *KubernetesNodeAddressAllOf`

NewKubernetesNodeAddressAllOf instantiates a new KubernetesNodeAddressAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesNodeAddressAllOfWithDefaults

`func NewKubernetesNodeAddressAllOfWithDefaults() *KubernetesNodeAddressAllOf`

NewKubernetesNodeAddressAllOfWithDefaults instantiates a new KubernetesNodeAddressAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *KubernetesNodeAddressAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *KubernetesNodeAddressAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *KubernetesNodeAddressAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *KubernetesNodeAddressAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *KubernetesNodeAddressAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *KubernetesNodeAddressAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAddress

`func (o *KubernetesNodeAddressAllOf) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *KubernetesNodeAddressAllOf) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *KubernetesNodeAddressAllOf) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *KubernetesNodeAddressAllOf) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetType

`func (o *KubernetesNodeAddressAllOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *KubernetesNodeAddressAllOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *KubernetesNodeAddressAllOf) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *KubernetesNodeAddressAllOf) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


