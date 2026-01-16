# ConnectorEncryptionTaskMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PublicEncryptionKey** | Pointer to **string** | The public encryption key in PEM format to use to encrypt the data. Data should be encrypted using RSA-OAEP with SHA-256 hashing function. For example, a device can use this key to encrypt data securely at rest and to transport it securely to Intersight. Intersight has the capability to decrypt this data. The encrypted data should be base64 encoded into a string format for persistence/transit. | [optional] 

## Methods

### NewConnectorEncryptionTaskMessage

`func NewConnectorEncryptionTaskMessage() *ConnectorEncryptionTaskMessage`

NewConnectorEncryptionTaskMessage instantiates a new ConnectorEncryptionTaskMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectorEncryptionTaskMessageWithDefaults

`func NewConnectorEncryptionTaskMessageWithDefaults() *ConnectorEncryptionTaskMessage`

NewConnectorEncryptionTaskMessageWithDefaults instantiates a new ConnectorEncryptionTaskMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPublicEncryptionKey

`func (o *ConnectorEncryptionTaskMessage) GetPublicEncryptionKey() string`

GetPublicEncryptionKey returns the PublicEncryptionKey field if non-nil, zero value otherwise.

### GetPublicEncryptionKeyOk

`func (o *ConnectorEncryptionTaskMessage) GetPublicEncryptionKeyOk() (*string, bool)`

GetPublicEncryptionKeyOk returns a tuple with the PublicEncryptionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicEncryptionKey

`func (o *ConnectorEncryptionTaskMessage) SetPublicEncryptionKey(v string)`

SetPublicEncryptionKey sets PublicEncryptionKey field to given value.

### HasPublicEncryptionKey

`func (o *ConnectorEncryptionTaskMessage) HasPublicEncryptionKey() bool`

HasPublicEncryptionKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


