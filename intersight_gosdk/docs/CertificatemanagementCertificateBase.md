# CertificatemanagementCertificateBase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | [default to "certificatemanagement.Imc"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | [default to "certificatemanagement.Imc"]
**Certificate** | Pointer to [**NullableX509Certificate**](X509Certificate.md) |  | [optional] 
**Enabled** | Pointer to **bool** | Enable/Disable the certificate in Certificate Management policy. | [optional] [default to true]
**IsPrivatekeySet** | Pointer to **bool** | Indicates whether the value of the &#39;privatekey&#39; property has been set. | [optional] [readonly] [default to false]
**Privatekey** | Pointer to **string** | Private Key which is used to validate the certificate. | [optional] 

## Methods

### NewCertificatemanagementCertificateBase

`func NewCertificatemanagementCertificateBase(classId string, objectType string, ) *CertificatemanagementCertificateBase`

NewCertificatemanagementCertificateBase instantiates a new CertificatemanagementCertificateBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificatemanagementCertificateBaseWithDefaults

`func NewCertificatemanagementCertificateBaseWithDefaults() *CertificatemanagementCertificateBase`

NewCertificatemanagementCertificateBaseWithDefaults instantiates a new CertificatemanagementCertificateBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *CertificatemanagementCertificateBase) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *CertificatemanagementCertificateBase) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *CertificatemanagementCertificateBase) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *CertificatemanagementCertificateBase) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CertificatemanagementCertificateBase) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CertificatemanagementCertificateBase) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCertificate

`func (o *CertificatemanagementCertificateBase) GetCertificate() X509Certificate`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *CertificatemanagementCertificateBase) GetCertificateOk() (*X509Certificate, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *CertificatemanagementCertificateBase) SetCertificate(v X509Certificate)`

SetCertificate sets Certificate field to given value.

### HasCertificate

`func (o *CertificatemanagementCertificateBase) HasCertificate() bool`

HasCertificate returns a boolean if a field has been set.

### SetCertificateNil

`func (o *CertificatemanagementCertificateBase) SetCertificateNil(b bool)`

 SetCertificateNil sets the value for Certificate to be an explicit nil

### UnsetCertificate
`func (o *CertificatemanagementCertificateBase) UnsetCertificate()`

UnsetCertificate ensures that no value is present for Certificate, not even an explicit nil
### GetEnabled

`func (o *CertificatemanagementCertificateBase) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CertificatemanagementCertificateBase) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CertificatemanagementCertificateBase) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *CertificatemanagementCertificateBase) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetIsPrivatekeySet

`func (o *CertificatemanagementCertificateBase) GetIsPrivatekeySet() bool`

GetIsPrivatekeySet returns the IsPrivatekeySet field if non-nil, zero value otherwise.

### GetIsPrivatekeySetOk

`func (o *CertificatemanagementCertificateBase) GetIsPrivatekeySetOk() (*bool, bool)`

GetIsPrivatekeySetOk returns a tuple with the IsPrivatekeySet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPrivatekeySet

`func (o *CertificatemanagementCertificateBase) SetIsPrivatekeySet(v bool)`

SetIsPrivatekeySet sets IsPrivatekeySet field to given value.

### HasIsPrivatekeySet

`func (o *CertificatemanagementCertificateBase) HasIsPrivatekeySet() bool`

HasIsPrivatekeySet returns a boolean if a field has been set.

### GetPrivatekey

`func (o *CertificatemanagementCertificateBase) GetPrivatekey() string`

GetPrivatekey returns the Privatekey field if non-nil, zero value otherwise.

### GetPrivatekeyOk

`func (o *CertificatemanagementCertificateBase) GetPrivatekeyOk() (*string, bool)`

GetPrivatekeyOk returns a tuple with the Privatekey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivatekey

`func (o *CertificatemanagementCertificateBase) SetPrivatekey(v string)`

SetPrivatekey sets Privatekey field to given value.

### HasPrivatekey

`func (o *CertificatemanagementCertificateBase) HasPrivatekey() bool`

HasPrivatekey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


