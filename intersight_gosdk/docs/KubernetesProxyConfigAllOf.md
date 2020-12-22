# KubernetesProxyConfigAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "kubernetes.ProxyConfig"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "kubernetes.ProxyConfig"]
**Hostname** | Pointer to **string** | HTTP/HTTPS Proxy server FQDN or IP. | [optional] 
**IsPasswordSet** | Pointer to **bool** | Indicates whether the value of the &#39;password&#39; property has been set. | [optional] [readonly] [default to false]
**Password** | Pointer to **string** | The password for the HTTP/HTTPS Proxy. | [optional] 
**Port** | Pointer to **int64** | The HTTP Proxy port number. The port number of the HTTP/HTTPS proxy must be between 1 and 65535, inclusive. | [optional] 
**Protocol** | Pointer to **string** | Protocol to use for the HTTP/HTTPS Proxy. | [optional] 
**Username** | Pointer to **string** | The username for the HTTP/HTTPS Proxy. | [optional] 

## Methods

### NewKubernetesProxyConfigAllOf

`func NewKubernetesProxyConfigAllOf(classId string, objectType string, ) *KubernetesProxyConfigAllOf`

NewKubernetesProxyConfigAllOf instantiates a new KubernetesProxyConfigAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesProxyConfigAllOfWithDefaults

`func NewKubernetesProxyConfigAllOfWithDefaults() *KubernetesProxyConfigAllOf`

NewKubernetesProxyConfigAllOfWithDefaults instantiates a new KubernetesProxyConfigAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *KubernetesProxyConfigAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *KubernetesProxyConfigAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *KubernetesProxyConfigAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *KubernetesProxyConfigAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *KubernetesProxyConfigAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *KubernetesProxyConfigAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetHostname

`func (o *KubernetesProxyConfigAllOf) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *KubernetesProxyConfigAllOf) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *KubernetesProxyConfigAllOf) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *KubernetesProxyConfigAllOf) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetIsPasswordSet

`func (o *KubernetesProxyConfigAllOf) GetIsPasswordSet() bool`

GetIsPasswordSet returns the IsPasswordSet field if non-nil, zero value otherwise.

### GetIsPasswordSetOk

`func (o *KubernetesProxyConfigAllOf) GetIsPasswordSetOk() (*bool, bool)`

GetIsPasswordSetOk returns a tuple with the IsPasswordSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPasswordSet

`func (o *KubernetesProxyConfigAllOf) SetIsPasswordSet(v bool)`

SetIsPasswordSet sets IsPasswordSet field to given value.

### HasIsPasswordSet

`func (o *KubernetesProxyConfigAllOf) HasIsPasswordSet() bool`

HasIsPasswordSet returns a boolean if a field has been set.

### GetPassword

`func (o *KubernetesProxyConfigAllOf) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *KubernetesProxyConfigAllOf) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *KubernetesProxyConfigAllOf) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *KubernetesProxyConfigAllOf) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetPort

`func (o *KubernetesProxyConfigAllOf) GetPort() int64`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *KubernetesProxyConfigAllOf) GetPortOk() (*int64, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *KubernetesProxyConfigAllOf) SetPort(v int64)`

SetPort sets Port field to given value.

### HasPort

`func (o *KubernetesProxyConfigAllOf) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetProtocol

`func (o *KubernetesProxyConfigAllOf) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *KubernetesProxyConfigAllOf) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *KubernetesProxyConfigAllOf) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *KubernetesProxyConfigAllOf) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetUsername

`func (o *KubernetesProxyConfigAllOf) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *KubernetesProxyConfigAllOf) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *KubernetesProxyConfigAllOf) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *KubernetesProxyConfigAllOf) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


