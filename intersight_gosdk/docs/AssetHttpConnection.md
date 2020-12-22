# AssetHttpConnection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "asset.HttpConnection"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "asset.HttpConnection"]
**CertificateAuthority** | Pointer to **string** | The certificate authority of the target. If set and connection is secure the connection will be validate using the servers identity with this certificate. If not set no validation will be done of the identity. | [optional] 
**IsSecure** | Pointer to **bool** | Indicates whether a connection to the target should be established using HTTPS. | [optional] [default to true]
**ManagementAddress** | Pointer to **string** | The DNS hostname or IP Address, either IPv4 or IPv6, to be used to connect to the managed target. | [optional] 
**Port** | Pointer to **int64** | The port number to be used to connect to the managed target. Values 1-65535 indicate a port number to be used. A value of 0 is not a valid port number and instead indicates that the default management port, as defined by the documentation of the managed target, should be used to establish a connection. | [optional] 

## Methods

### NewAssetHttpConnection

`func NewAssetHttpConnection(classId string, objectType string, ) *AssetHttpConnection`

NewAssetHttpConnection instantiates a new AssetHttpConnection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetHttpConnectionWithDefaults

`func NewAssetHttpConnectionWithDefaults() *AssetHttpConnection`

NewAssetHttpConnectionWithDefaults instantiates a new AssetHttpConnection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *AssetHttpConnection) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *AssetHttpConnection) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *AssetHttpConnection) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *AssetHttpConnection) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *AssetHttpConnection) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *AssetHttpConnection) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCertificateAuthority

`func (o *AssetHttpConnection) GetCertificateAuthority() string`

GetCertificateAuthority returns the CertificateAuthority field if non-nil, zero value otherwise.

### GetCertificateAuthorityOk

`func (o *AssetHttpConnection) GetCertificateAuthorityOk() (*string, bool)`

GetCertificateAuthorityOk returns a tuple with the CertificateAuthority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateAuthority

`func (o *AssetHttpConnection) SetCertificateAuthority(v string)`

SetCertificateAuthority sets CertificateAuthority field to given value.

### HasCertificateAuthority

`func (o *AssetHttpConnection) HasCertificateAuthority() bool`

HasCertificateAuthority returns a boolean if a field has been set.

### GetIsSecure

`func (o *AssetHttpConnection) GetIsSecure() bool`

GetIsSecure returns the IsSecure field if non-nil, zero value otherwise.

### GetIsSecureOk

`func (o *AssetHttpConnection) GetIsSecureOk() (*bool, bool)`

GetIsSecureOk returns a tuple with the IsSecure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSecure

`func (o *AssetHttpConnection) SetIsSecure(v bool)`

SetIsSecure sets IsSecure field to given value.

### HasIsSecure

`func (o *AssetHttpConnection) HasIsSecure() bool`

HasIsSecure returns a boolean if a field has been set.

### GetManagementAddress

`func (o *AssetHttpConnection) GetManagementAddress() string`

GetManagementAddress returns the ManagementAddress field if non-nil, zero value otherwise.

### GetManagementAddressOk

`func (o *AssetHttpConnection) GetManagementAddressOk() (*string, bool)`

GetManagementAddressOk returns a tuple with the ManagementAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementAddress

`func (o *AssetHttpConnection) SetManagementAddress(v string)`

SetManagementAddress sets ManagementAddress field to given value.

### HasManagementAddress

`func (o *AssetHttpConnection) HasManagementAddress() bool`

HasManagementAddress returns a boolean if a field has been set.

### GetPort

`func (o *AssetHttpConnection) GetPort() int64`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *AssetHttpConnection) GetPortOk() (*int64, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *AssetHttpConnection) SetPort(v int64)`

SetPort sets Port field to given value.

### HasPort

`func (o *AssetHttpConnection) HasPort() bool`

HasPort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


