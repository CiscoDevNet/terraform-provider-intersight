# ConnectorDownloadStatusAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Checksum** | Pointer to [**ConnectorFileChecksum**](connector.FileChecksum.md) |  | [optional] 
**DownloadError** | Pointer to **string** | Any error encountered. Set to empty when download is in progress or completed. | [optional] 
**DownloadProgress** | Pointer to **int64** | The download progress of the file represented as a percentage between 0% and 100%. If progress reporting is not possible a value of -1 is sent. | [optional] 
**DownloadRetries** | Pointer to **int64** | The number of retries the plugin attempted before succeeding or failing the download. | [optional] 
**Sha256checksum** | Pointer to **string** | The sha256checksum of the downloaded file as calculated by the download plugin after successfully downloading a file. | [optional] 

## Methods

### NewConnectorDownloadStatusAllOf

`func NewConnectorDownloadStatusAllOf() *ConnectorDownloadStatusAllOf`

NewConnectorDownloadStatusAllOf instantiates a new ConnectorDownloadStatusAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectorDownloadStatusAllOfWithDefaults

`func NewConnectorDownloadStatusAllOfWithDefaults() *ConnectorDownloadStatusAllOf`

NewConnectorDownloadStatusAllOfWithDefaults instantiates a new ConnectorDownloadStatusAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChecksum

`func (o *ConnectorDownloadStatusAllOf) GetChecksum() ConnectorFileChecksum`

GetChecksum returns the Checksum field if non-nil, zero value otherwise.

### GetChecksumOk

`func (o *ConnectorDownloadStatusAllOf) GetChecksumOk() (*ConnectorFileChecksum, bool)`

GetChecksumOk returns a tuple with the Checksum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecksum

`func (o *ConnectorDownloadStatusAllOf) SetChecksum(v ConnectorFileChecksum)`

SetChecksum sets Checksum field to given value.

### HasChecksum

`func (o *ConnectorDownloadStatusAllOf) HasChecksum() bool`

HasChecksum returns a boolean if a field has been set.

### GetDownloadError

`func (o *ConnectorDownloadStatusAllOf) GetDownloadError() string`

GetDownloadError returns the DownloadError field if non-nil, zero value otherwise.

### GetDownloadErrorOk

`func (o *ConnectorDownloadStatusAllOf) GetDownloadErrorOk() (*string, bool)`

GetDownloadErrorOk returns a tuple with the DownloadError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadError

`func (o *ConnectorDownloadStatusAllOf) SetDownloadError(v string)`

SetDownloadError sets DownloadError field to given value.

### HasDownloadError

`func (o *ConnectorDownloadStatusAllOf) HasDownloadError() bool`

HasDownloadError returns a boolean if a field has been set.

### GetDownloadProgress

`func (o *ConnectorDownloadStatusAllOf) GetDownloadProgress() int64`

GetDownloadProgress returns the DownloadProgress field if non-nil, zero value otherwise.

### GetDownloadProgressOk

`func (o *ConnectorDownloadStatusAllOf) GetDownloadProgressOk() (*int64, bool)`

GetDownloadProgressOk returns a tuple with the DownloadProgress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadProgress

`func (o *ConnectorDownloadStatusAllOf) SetDownloadProgress(v int64)`

SetDownloadProgress sets DownloadProgress field to given value.

### HasDownloadProgress

`func (o *ConnectorDownloadStatusAllOf) HasDownloadProgress() bool`

HasDownloadProgress returns a boolean if a field has been set.

### GetDownloadRetries

`func (o *ConnectorDownloadStatusAllOf) GetDownloadRetries() int64`

GetDownloadRetries returns the DownloadRetries field if non-nil, zero value otherwise.

### GetDownloadRetriesOk

`func (o *ConnectorDownloadStatusAllOf) GetDownloadRetriesOk() (*int64, bool)`

GetDownloadRetriesOk returns a tuple with the DownloadRetries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadRetries

`func (o *ConnectorDownloadStatusAllOf) SetDownloadRetries(v int64)`

SetDownloadRetries sets DownloadRetries field to given value.

### HasDownloadRetries

`func (o *ConnectorDownloadStatusAllOf) HasDownloadRetries() bool`

HasDownloadRetries returns a boolean if a field has been set.

### GetSha256checksum

`func (o *ConnectorDownloadStatusAllOf) GetSha256checksum() string`

GetSha256checksum returns the Sha256checksum field if non-nil, zero value otherwise.

### GetSha256checksumOk

`func (o *ConnectorDownloadStatusAllOf) GetSha256checksumOk() (*string, bool)`

GetSha256checksumOk returns a tuple with the Sha256checksum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha256checksum

`func (o *ConnectorDownloadStatusAllOf) SetSha256checksum(v string)`

SetSha256checksum sets Sha256checksum field to given value.

### HasSha256checksum

`func (o *ConnectorDownloadStatusAllOf) HasSha256checksum() bool`

HasSha256checksum returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


