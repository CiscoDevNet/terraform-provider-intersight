# AssetHttpConnectionAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Credential** | Pointer to [**AssetCredential**](asset.Credential.md) |  | [optional] 
**IsSecure** | Pointer to **bool** | Indicates whether a connection to the target should be established using HTTPS. | [optional] 
**ManagementAddress** | Pointer to **string** | The DNS hostname or IP Address, either IPv4 or IPv6, to be used to connect to the managed target. | [optional] 
**Port** | Pointer to **int64** | The port number to be used to to connect to the managed target. Values 1-65535 indicate a port number to be used. A value of 0 is not a valid port number and instead indicates that the default management port, as defined by the documentation of the managed target, should be used to establish a connection. | [optional] 

## Methods

### NewAssetHttpConnectionAllOf

`func NewAssetHttpConnectionAllOf() *AssetHttpConnectionAllOf`

NewAssetHttpConnectionAllOf instantiates a new AssetHttpConnectionAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetHttpConnectionAllOfWithDefaults

`func NewAssetHttpConnectionAllOfWithDefaults() *AssetHttpConnectionAllOf`

NewAssetHttpConnectionAllOfWithDefaults instantiates a new AssetHttpConnectionAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCredential

`func (o *AssetHttpConnectionAllOf) GetCredential() AssetCredential`

GetCredential returns the Credential field if non-nil, zero value otherwise.

### GetCredentialOk

`func (o *AssetHttpConnectionAllOf) GetCredentialOk() (*AssetCredential, bool)`

GetCredentialOk returns a tuple with the Credential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredential

`func (o *AssetHttpConnectionAllOf) SetCredential(v AssetCredential)`

SetCredential sets Credential field to given value.

### HasCredential

`func (o *AssetHttpConnectionAllOf) HasCredential() bool`

HasCredential returns a boolean if a field has been set.

### GetIsSecure

`func (o *AssetHttpConnectionAllOf) GetIsSecure() bool`

GetIsSecure returns the IsSecure field if non-nil, zero value otherwise.

### GetIsSecureOk

`func (o *AssetHttpConnectionAllOf) GetIsSecureOk() (*bool, bool)`

GetIsSecureOk returns a tuple with the IsSecure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSecure

`func (o *AssetHttpConnectionAllOf) SetIsSecure(v bool)`

SetIsSecure sets IsSecure field to given value.

### HasIsSecure

`func (o *AssetHttpConnectionAllOf) HasIsSecure() bool`

HasIsSecure returns a boolean if a field has been set.

### GetManagementAddress

`func (o *AssetHttpConnectionAllOf) GetManagementAddress() string`

GetManagementAddress returns the ManagementAddress field if non-nil, zero value otherwise.

### GetManagementAddressOk

`func (o *AssetHttpConnectionAllOf) GetManagementAddressOk() (*string, bool)`

GetManagementAddressOk returns a tuple with the ManagementAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementAddress

`func (o *AssetHttpConnectionAllOf) SetManagementAddress(v string)`

SetManagementAddress sets ManagementAddress field to given value.

### HasManagementAddress

`func (o *AssetHttpConnectionAllOf) HasManagementAddress() bool`

HasManagementAddress returns a boolean if a field has been set.

### GetPort

`func (o *AssetHttpConnectionAllOf) GetPort() int64`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *AssetHttpConnectionAllOf) GetPortOk() (*int64, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *AssetHttpConnectionAllOf) SetPort(v int64)`

SetPort sets Port field to given value.

### HasPort

`func (o *AssetHttpConnectionAllOf) HasPort() bool`

HasPort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


