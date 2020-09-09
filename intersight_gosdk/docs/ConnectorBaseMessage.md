# ConnectorBaseMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EncryptedAesKey** | Pointer to **string** | The secure properties that have large text content as value can be encrypted using AES key. In these cases, the AES key needs to be encrypted using the device connector public key and passed as the value for this property. The secure properties that are encrypted using the AES key are mapped against the property name with prefix &#39;AES&#39; in SecureProperties dictionary. | [optional] 
**EncryptionKey** | Pointer to **string** | The public key that was used to encrypt the values present in SecureProperties dictionary. If the given public key is not same as device connector&#39;s public key, an error reponse with appropriate error message is thrown back. | [optional] 
**SecureProperties** | Pointer to **interface{}** | A dictionary of encrypted secure values mapped against the secure property name. The values that are encrypted using AES key must be mapped against the secure property name with a &#39;AES&#39; prefix Device connector expects the message body to be a golang template and the template can use the secure property names as placeholders. | [optional] 

## Methods

### NewConnectorBaseMessage

`func NewConnectorBaseMessage() *ConnectorBaseMessage`

NewConnectorBaseMessage instantiates a new ConnectorBaseMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectorBaseMessageWithDefaults

`func NewConnectorBaseMessageWithDefaults() *ConnectorBaseMessage`

NewConnectorBaseMessageWithDefaults instantiates a new ConnectorBaseMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEncryptedAesKey

`func (o *ConnectorBaseMessage) GetEncryptedAesKey() string`

GetEncryptedAesKey returns the EncryptedAesKey field if non-nil, zero value otherwise.

### GetEncryptedAesKeyOk

`func (o *ConnectorBaseMessage) GetEncryptedAesKeyOk() (*string, bool)`

GetEncryptedAesKeyOk returns a tuple with the EncryptedAesKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptedAesKey

`func (o *ConnectorBaseMessage) SetEncryptedAesKey(v string)`

SetEncryptedAesKey sets EncryptedAesKey field to given value.

### HasEncryptedAesKey

`func (o *ConnectorBaseMessage) HasEncryptedAesKey() bool`

HasEncryptedAesKey returns a boolean if a field has been set.

### GetEncryptionKey

`func (o *ConnectorBaseMessage) GetEncryptionKey() string`

GetEncryptionKey returns the EncryptionKey field if non-nil, zero value otherwise.

### GetEncryptionKeyOk

`func (o *ConnectorBaseMessage) GetEncryptionKeyOk() (*string, bool)`

GetEncryptionKeyOk returns a tuple with the EncryptionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionKey

`func (o *ConnectorBaseMessage) SetEncryptionKey(v string)`

SetEncryptionKey sets EncryptionKey field to given value.

### HasEncryptionKey

`func (o *ConnectorBaseMessage) HasEncryptionKey() bool`

HasEncryptionKey returns a boolean if a field has been set.

### GetSecureProperties

`func (o *ConnectorBaseMessage) GetSecureProperties() interface{}`

GetSecureProperties returns the SecureProperties field if non-nil, zero value otherwise.

### GetSecurePropertiesOk

`func (o *ConnectorBaseMessage) GetSecurePropertiesOk() (*interface{}, bool)`

GetSecurePropertiesOk returns a tuple with the SecureProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureProperties

`func (o *ConnectorBaseMessage) SetSecureProperties(v interface{})`

SetSecureProperties sets SecureProperties field to given value.

### HasSecureProperties

`func (o *ConnectorBaseMessage) HasSecureProperties() bool`

HasSecureProperties returns a boolean if a field has been set.

### SetSecurePropertiesNil

`func (o *ConnectorBaseMessage) SetSecurePropertiesNil(b bool)`

 SetSecurePropertiesNil sets the value for SecureProperties to be an explicit nil

### UnsetSecureProperties
`func (o *ConnectorBaseMessage) UnsetSecureProperties()`

UnsetSecureProperties ensures that no value is present for SecureProperties, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


