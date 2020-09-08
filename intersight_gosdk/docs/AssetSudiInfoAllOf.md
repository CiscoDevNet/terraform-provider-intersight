# AssetSudiInfoAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pid** | Pointer to **string** | The device model (PID) extracted from the X.509 SUDI Leaf Certificate. | [optional] 
**SerialNumber** | Pointer to **string** | The device SerialNumber extracted from the X.509 SUDI Leaf Certiicate. | [optional] 
**Signature** | Pointer to **string** | The signature is obtained by taking the base64 encoding of the Serial Number + PID + Status, taking the SHA256 hash and then signing with the SUDI X.509 Leaf Certifiate. | [optional] 
**Status** | Pointer to **string** | The validation status of the device. * &#x60;DeviceStatusUnknown&#x60; - SUDI validation is done on the establishment of a connection. Before a device connects or after it disconnects, the SUDI validation status is set to this value. * &#x60;Verified&#x60; - The device returned a valid PID, Serial Number, Status and X.509 Leaf Certificate. The certificate signing chain was validated. * &#x60;CertificateValidationFailed&#x60; - Validation of the certificate signing chain failed. * &#x60;UnsupportedFirmware&#x60; - The firmware version of the Cisco IMC that is installed does not contain the SUDI APIs needed to perform validation. * &#x60;UnsupportedHardware&#x60; - The device is a model that does not contain a Trust Anchor Module (TAM) and thus cannot be validated. * &#x60;DeviceNotResponding&#x60; - An request was sent to the device, but no response was received. | [optional] [default to "DeviceStatusUnknown"]
**SudiCertificate** | Pointer to [**X509Certificate**](x509.Certificate.md) |  | [optional] 

## Methods

### NewAssetSudiInfoAllOf

`func NewAssetSudiInfoAllOf() *AssetSudiInfoAllOf`

NewAssetSudiInfoAllOf instantiates a new AssetSudiInfoAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetSudiInfoAllOfWithDefaults

`func NewAssetSudiInfoAllOfWithDefaults() *AssetSudiInfoAllOf`

NewAssetSudiInfoAllOfWithDefaults instantiates a new AssetSudiInfoAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPid

`func (o *AssetSudiInfoAllOf) GetPid() string`

GetPid returns the Pid field if non-nil, zero value otherwise.

### GetPidOk

`func (o *AssetSudiInfoAllOf) GetPidOk() (*string, bool)`

GetPidOk returns a tuple with the Pid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPid

`func (o *AssetSudiInfoAllOf) SetPid(v string)`

SetPid sets Pid field to given value.

### HasPid

`func (o *AssetSudiInfoAllOf) HasPid() bool`

HasPid returns a boolean if a field has been set.

### GetSerialNumber

`func (o *AssetSudiInfoAllOf) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *AssetSudiInfoAllOf) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *AssetSudiInfoAllOf) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *AssetSudiInfoAllOf) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### GetSignature

`func (o *AssetSudiInfoAllOf) GetSignature() string`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *AssetSudiInfoAllOf) GetSignatureOk() (*string, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *AssetSudiInfoAllOf) SetSignature(v string)`

SetSignature sets Signature field to given value.

### HasSignature

`func (o *AssetSudiInfoAllOf) HasSignature() bool`

HasSignature returns a boolean if a field has been set.

### GetStatus

`func (o *AssetSudiInfoAllOf) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AssetSudiInfoAllOf) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AssetSudiInfoAllOf) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AssetSudiInfoAllOf) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSudiCertificate

`func (o *AssetSudiInfoAllOf) GetSudiCertificate() X509Certificate`

GetSudiCertificate returns the SudiCertificate field if non-nil, zero value otherwise.

### GetSudiCertificateOk

`func (o *AssetSudiInfoAllOf) GetSudiCertificateOk() (*X509Certificate, bool)`

GetSudiCertificateOk returns a tuple with the SudiCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSudiCertificate

`func (o *AssetSudiInfoAllOf) SetSudiCertificate(v X509Certificate)`

SetSudiCertificate sets SudiCertificate field to given value.

### HasSudiCertificate

`func (o *AssetSudiInfoAllOf) HasSudiCertificate() bool`

HasSudiCertificate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


