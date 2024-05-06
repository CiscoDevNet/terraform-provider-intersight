# AssetSudiInfoAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "asset.SudiInfo"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "asset.SudiInfo"]
**Pid** | Pointer to **string** | The device model (PID) extracted from the X.509 SUDI leaf certificate. | [optional] [readonly] 
**SerialNumber** | Pointer to **string** | The device SerialNumber extracted from the X.509 SUDI leaf certificate. | [optional] [readonly] 
**Signature** | Pointer to **string** | The base64-encoding of a SUDI signature, signed with the provided Trust Anchor Module (TAM) private key and signatureScheme. | [optional] [readonly] 
**Status** | Pointer to **string** | The validation status of the device. * &#x60;DeviceStatusUnknown&#x60; - SUDI validation is done on the establishment of a connection. Before a device connects or after it disconnects, the SUDI validation status is set to this value. * &#x60;Verified&#x60; - The SUDI signature and certificate have been validated by the device. The device PID, Serial Number, Status and X.509 certificates were validated by a Trusted Anchor Module. The Intersight services have not cross-validated the SUDI signature. * &#x60;EndToEndVerified&#x60; - The SUDI signature and certificate have been validated by the device and by the Intersight services. The device PID, Serial Number, Status and X.509 certificates were validated by a Trusted Anchor Module. * &#x60;CertificateValidationFailed&#x60; - Validation of the certificate signing chain failed. * &#x60;UnsupportedSignatureScheme&#x60; - The SUDI signature scheme is not supported. * &#x60;SignatureValidationFailed&#x60; - Validation of the SUDI signature has failed. * &#x60;UnsupportedFirmware&#x60; - The firmware version of the Cisco IMC that is installed does not contain the SUDI APIs needed to perform validation. * &#x60;UnsupportedHardware&#x60; - The device is a model that does not contain a Trust Anchor Module (TAM) and thus cannot be validated. * &#x60;DeviceError&#x60; - The device returned an error when it received a SUDIrequrest request. * &#x60;DeviceNotResponding&#x60; - A request was sent to the device, but no response was received. * &#x60;InvalidSignatureEncoding&#x60; - The encoding of the SUDI signature is invalid. * &#x60;MissingSignature&#x60; - The SUDI signature is missing. * &#x60;MissingSignatureScheme&#x60; - The SUDI signature scheme is missing. * &#x60;MissingCertificate&#x60; - The SUDI certificate is missing. * &#x60;InvalidCertificateEncoding&#x60; - The encoding of the SUDI certificate is invalid. * &#x60;InvalidIdentity&#x60; - The model or serial number of the device is either missing or does not match the certificate attributes. * &#x60;Suspect&#x60; - The DC claims the SUDI has been verified, but there is no way to validate the claim. | [optional] [readonly] [default to "DeviceStatusUnknown"]
**StatusDetails** | Pointer to **string** | The detailed validation status of the device, such as an error message explaining why the validation failed. | [optional] [readonly] 
**SudiCertificate** | Pointer to [**NullableX509Certificate**](X509Certificate.md) |  | [optional] 

## Methods

### NewAssetSudiInfoAllOf

`func NewAssetSudiInfoAllOf(classId string, objectType string, ) *AssetSudiInfoAllOf`

NewAssetSudiInfoAllOf instantiates a new AssetSudiInfoAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetSudiInfoAllOfWithDefaults

`func NewAssetSudiInfoAllOfWithDefaults() *AssetSudiInfoAllOf`

NewAssetSudiInfoAllOfWithDefaults instantiates a new AssetSudiInfoAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *AssetSudiInfoAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *AssetSudiInfoAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *AssetSudiInfoAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *AssetSudiInfoAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *AssetSudiInfoAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *AssetSudiInfoAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


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

### GetStatusDetails

`func (o *AssetSudiInfoAllOf) GetStatusDetails() string`

GetStatusDetails returns the StatusDetails field if non-nil, zero value otherwise.

### GetStatusDetailsOk

`func (o *AssetSudiInfoAllOf) GetStatusDetailsOk() (*string, bool)`

GetStatusDetailsOk returns a tuple with the StatusDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusDetails

`func (o *AssetSudiInfoAllOf) SetStatusDetails(v string)`

SetStatusDetails sets StatusDetails field to given value.

### HasStatusDetails

`func (o *AssetSudiInfoAllOf) HasStatusDetails() bool`

HasStatusDetails returns a boolean if a field has been set.

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

### SetSudiCertificateNil

`func (o *AssetSudiInfoAllOf) SetSudiCertificateNil(b bool)`

 SetSudiCertificateNil sets the value for SudiCertificate to be an explicit nil

### UnsetSudiCertificate
`func (o *AssetSudiInfoAllOf) UnsetSudiCertificate()`

UnsetSudiCertificate ensures that no value is present for SudiCertificate, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


