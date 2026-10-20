# CertificatemanagementKmipClientConfigurationBundle

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "certificatemanagement.KmipClientConfigurationBundle"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "certificatemanagement.KmipClientConfigurationBundle"]
**BundlePassword** | Pointer to **string** | Password for the certificate bundle/Private Key. It&#39;s required when the bundle type is PKCS#8/PKCS#12. | [optional] 
**BundleType** | Pointer to **string** | Type of the KMIP client certificate bundle. * &#x60;PKCS1&#x60; - The private key is in PKCS#1 format and does not contain certificate chain. * &#x60;PKCS8&#x60; - The private key here is in PKCS#8 format and may be password protected. * &#x60;PKCS12&#x60; - The bundle is in PKCS#12 format that contains client certificate, private key. This Bundle is password protected. | [optional] [default to "PKCS1"]
**CertificateBundle** | Pointer to **string** | Certificate Bundle that contains client certificate and Private Key. This Bundle is password protected. | [optional] 
**IsBundlePasswordSet** | Pointer to **bool** | Indicates whether the value of the &#39;bundlePassword&#39; property has been set. | [optional] [readonly] [default to false]
**IsPrivatekeySet** | Pointer to **bool** | Indicates whether the value of the &#39;privatekey&#39; property has been set. | [optional] [readonly] [default to false]
**Privatekey** | Pointer to **string** | Private Key which is used to validate the certificate. | [optional] 

## Methods

### NewCertificatemanagementKmipClientConfigurationBundle

`func NewCertificatemanagementKmipClientConfigurationBundle(classId string, objectType string, ) *CertificatemanagementKmipClientConfigurationBundle`

NewCertificatemanagementKmipClientConfigurationBundle instantiates a new CertificatemanagementKmipClientConfigurationBundle object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificatemanagementKmipClientConfigurationBundleWithDefaults

`func NewCertificatemanagementKmipClientConfigurationBundleWithDefaults() *CertificatemanagementKmipClientConfigurationBundle`

NewCertificatemanagementKmipClientConfigurationBundleWithDefaults instantiates a new CertificatemanagementKmipClientConfigurationBundle object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *CertificatemanagementKmipClientConfigurationBundle) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *CertificatemanagementKmipClientConfigurationBundle) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *CertificatemanagementKmipClientConfigurationBundle) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *CertificatemanagementKmipClientConfigurationBundle) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CertificatemanagementKmipClientConfigurationBundle) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CertificatemanagementKmipClientConfigurationBundle) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetBundlePassword

`func (o *CertificatemanagementKmipClientConfigurationBundle) GetBundlePassword() string`

GetBundlePassword returns the BundlePassword field if non-nil, zero value otherwise.

### GetBundlePasswordOk

`func (o *CertificatemanagementKmipClientConfigurationBundle) GetBundlePasswordOk() (*string, bool)`

GetBundlePasswordOk returns a tuple with the BundlePassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundlePassword

`func (o *CertificatemanagementKmipClientConfigurationBundle) SetBundlePassword(v string)`

SetBundlePassword sets BundlePassword field to given value.

### HasBundlePassword

`func (o *CertificatemanagementKmipClientConfigurationBundle) HasBundlePassword() bool`

HasBundlePassword returns a boolean if a field has been set.

### GetBundleType

`func (o *CertificatemanagementKmipClientConfigurationBundle) GetBundleType() string`

GetBundleType returns the BundleType field if non-nil, zero value otherwise.

### GetBundleTypeOk

`func (o *CertificatemanagementKmipClientConfigurationBundle) GetBundleTypeOk() (*string, bool)`

GetBundleTypeOk returns a tuple with the BundleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundleType

`func (o *CertificatemanagementKmipClientConfigurationBundle) SetBundleType(v string)`

SetBundleType sets BundleType field to given value.

### HasBundleType

`func (o *CertificatemanagementKmipClientConfigurationBundle) HasBundleType() bool`

HasBundleType returns a boolean if a field has been set.

### GetCertificateBundle

`func (o *CertificatemanagementKmipClientConfigurationBundle) GetCertificateBundle() string`

GetCertificateBundle returns the CertificateBundle field if non-nil, zero value otherwise.

### GetCertificateBundleOk

`func (o *CertificatemanagementKmipClientConfigurationBundle) GetCertificateBundleOk() (*string, bool)`

GetCertificateBundleOk returns a tuple with the CertificateBundle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateBundle

`func (o *CertificatemanagementKmipClientConfigurationBundle) SetCertificateBundle(v string)`

SetCertificateBundle sets CertificateBundle field to given value.

### HasCertificateBundle

`func (o *CertificatemanagementKmipClientConfigurationBundle) HasCertificateBundle() bool`

HasCertificateBundle returns a boolean if a field has been set.

### GetIsBundlePasswordSet

`func (o *CertificatemanagementKmipClientConfigurationBundle) GetIsBundlePasswordSet() bool`

GetIsBundlePasswordSet returns the IsBundlePasswordSet field if non-nil, zero value otherwise.

### GetIsBundlePasswordSetOk

`func (o *CertificatemanagementKmipClientConfigurationBundle) GetIsBundlePasswordSetOk() (*bool, bool)`

GetIsBundlePasswordSetOk returns a tuple with the IsBundlePasswordSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBundlePasswordSet

`func (o *CertificatemanagementKmipClientConfigurationBundle) SetIsBundlePasswordSet(v bool)`

SetIsBundlePasswordSet sets IsBundlePasswordSet field to given value.

### HasIsBundlePasswordSet

`func (o *CertificatemanagementKmipClientConfigurationBundle) HasIsBundlePasswordSet() bool`

HasIsBundlePasswordSet returns a boolean if a field has been set.

### GetIsPrivatekeySet

`func (o *CertificatemanagementKmipClientConfigurationBundle) GetIsPrivatekeySet() bool`

GetIsPrivatekeySet returns the IsPrivatekeySet field if non-nil, zero value otherwise.

### GetIsPrivatekeySetOk

`func (o *CertificatemanagementKmipClientConfigurationBundle) GetIsPrivatekeySetOk() (*bool, bool)`

GetIsPrivatekeySetOk returns a tuple with the IsPrivatekeySet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPrivatekeySet

`func (o *CertificatemanagementKmipClientConfigurationBundle) SetIsPrivatekeySet(v bool)`

SetIsPrivatekeySet sets IsPrivatekeySet field to given value.

### HasIsPrivatekeySet

`func (o *CertificatemanagementKmipClientConfigurationBundle) HasIsPrivatekeySet() bool`

HasIsPrivatekeySet returns a boolean if a field has been set.

### GetPrivatekey

`func (o *CertificatemanagementKmipClientConfigurationBundle) GetPrivatekey() string`

GetPrivatekey returns the Privatekey field if non-nil, zero value otherwise.

### GetPrivatekeyOk

`func (o *CertificatemanagementKmipClientConfigurationBundle) GetPrivatekeyOk() (*string, bool)`

GetPrivatekeyOk returns a tuple with the Privatekey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivatekey

`func (o *CertificatemanagementKmipClientConfigurationBundle) SetPrivatekey(v string)`

SetPrivatekey sets Privatekey field to given value.

### HasPrivatekey

`func (o *CertificatemanagementKmipClientConfigurationBundle) HasPrivatekey() bool`

HasPrivatekey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


