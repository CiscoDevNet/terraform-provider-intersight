# ConnectorFileChecksum

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hash** | Pointer to **string** | The calculated hash of the contents using the algorithm. | [optional] 
**HashAlgorithm** | Pointer to **string** | The hash algorithm used to calculate the checksum. * &#x60;crc&#x60; - A CRC hash as definded by RFC 3385. Generated with the IEEE polynomial. * &#x60;sha256&#x60; - A SHA256 hash as defined by RFC 4634. | [optional] [default to "crc"]

## Methods

### NewConnectorFileChecksum

`func NewConnectorFileChecksum() *ConnectorFileChecksum`

NewConnectorFileChecksum instantiates a new ConnectorFileChecksum object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectorFileChecksumWithDefaults

`func NewConnectorFileChecksumWithDefaults() *ConnectorFileChecksum`

NewConnectorFileChecksumWithDefaults instantiates a new ConnectorFileChecksum object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHash

`func (o *ConnectorFileChecksum) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *ConnectorFileChecksum) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *ConnectorFileChecksum) SetHash(v string)`

SetHash sets Hash field to given value.

### HasHash

`func (o *ConnectorFileChecksum) HasHash() bool`

HasHash returns a boolean if a field has been set.

### GetHashAlgorithm

`func (o *ConnectorFileChecksum) GetHashAlgorithm() string`

GetHashAlgorithm returns the HashAlgorithm field if non-nil, zero value otherwise.

### GetHashAlgorithmOk

`func (o *ConnectorFileChecksum) GetHashAlgorithmOk() (*string, bool)`

GetHashAlgorithmOk returns a tuple with the HashAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashAlgorithm

`func (o *ConnectorFileChecksum) SetHashAlgorithm(v string)`

SetHashAlgorithm sets HashAlgorithm field to given value.

### HasHashAlgorithm

`func (o *ConnectorFileChecksum) HasHashAlgorithm() bool`

HasHashAlgorithm returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


