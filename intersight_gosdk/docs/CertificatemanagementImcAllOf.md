# CertificatemanagementImcAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "certificatemanagement.Imc"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "certificatemanagement.Imc"]
**CertType** | Pointer to **string** | Certificate Type for the certificate management. * &#x60;None&#x60; - Set certificate on the selected end point . * &#x60;KMIPClient&#x60; - Set KMIP certificate on the selected end point. | [optional] [default to "None"]
**IsPrivatekeySet** | Pointer to **bool** | Indicates whether the value of the &#39;privatekey&#39; property has been set. | [optional] [readonly] [default to false]
**Privatekey** | Pointer to **string** | Private Key which is used to validate the certificate. | [optional] 

## Methods

### NewCertificatemanagementImcAllOf

`func NewCertificatemanagementImcAllOf(classId string, objectType string, ) *CertificatemanagementImcAllOf`

NewCertificatemanagementImcAllOf instantiates a new CertificatemanagementImcAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificatemanagementImcAllOfWithDefaults

`func NewCertificatemanagementImcAllOfWithDefaults() *CertificatemanagementImcAllOf`

NewCertificatemanagementImcAllOfWithDefaults instantiates a new CertificatemanagementImcAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *CertificatemanagementImcAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *CertificatemanagementImcAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *CertificatemanagementImcAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *CertificatemanagementImcAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CertificatemanagementImcAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CertificatemanagementImcAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCertType

`func (o *CertificatemanagementImcAllOf) GetCertType() string`

GetCertType returns the CertType field if non-nil, zero value otherwise.

### GetCertTypeOk

`func (o *CertificatemanagementImcAllOf) GetCertTypeOk() (*string, bool)`

GetCertTypeOk returns a tuple with the CertType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertType

`func (o *CertificatemanagementImcAllOf) SetCertType(v string)`

SetCertType sets CertType field to given value.

### HasCertType

`func (o *CertificatemanagementImcAllOf) HasCertType() bool`

HasCertType returns a boolean if a field has been set.

### GetIsPrivatekeySet

`func (o *CertificatemanagementImcAllOf) GetIsPrivatekeySet() bool`

GetIsPrivatekeySet returns the IsPrivatekeySet field if non-nil, zero value otherwise.

### GetIsPrivatekeySetOk

`func (o *CertificatemanagementImcAllOf) GetIsPrivatekeySetOk() (*bool, bool)`

GetIsPrivatekeySetOk returns a tuple with the IsPrivatekeySet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPrivatekeySet

`func (o *CertificatemanagementImcAllOf) SetIsPrivatekeySet(v bool)`

SetIsPrivatekeySet sets IsPrivatekeySet field to given value.

### HasIsPrivatekeySet

`func (o *CertificatemanagementImcAllOf) HasIsPrivatekeySet() bool`

HasIsPrivatekeySet returns a boolean if a field has been set.

### GetPrivatekey

`func (o *CertificatemanagementImcAllOf) GetPrivatekey() string`

GetPrivatekey returns the Privatekey field if non-nil, zero value otherwise.

### GetPrivatekeyOk

`func (o *CertificatemanagementImcAllOf) GetPrivatekeyOk() (*string, bool)`

GetPrivatekeyOk returns a tuple with the Privatekey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivatekey

`func (o *CertificatemanagementImcAllOf) SetPrivatekey(v string)`

SetPrivatekey sets Privatekey field to given value.

### HasPrivatekey

`func (o *CertificatemanagementImcAllOf) HasPrivatekey() bool`

HasPrivatekey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


