# CertificatemanagementRootCaCertificateAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "certificatemanagement.RootCaCertificate"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "certificatemanagement.RootCaCertificate"]
**CertificateName** | Pointer to **string** | A name that helps identify a certificate. It can be any string that adheres to the following constraints. It should start and end with an alphanumeric character. It can have underscores and hyphens. It cannot be more than 30 characters. | [optional] 

## Methods

### NewCertificatemanagementRootCaCertificateAllOf

`func NewCertificatemanagementRootCaCertificateAllOf(classId string, objectType string, ) *CertificatemanagementRootCaCertificateAllOf`

NewCertificatemanagementRootCaCertificateAllOf instantiates a new CertificatemanagementRootCaCertificateAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificatemanagementRootCaCertificateAllOfWithDefaults

`func NewCertificatemanagementRootCaCertificateAllOfWithDefaults() *CertificatemanagementRootCaCertificateAllOf`

NewCertificatemanagementRootCaCertificateAllOfWithDefaults instantiates a new CertificatemanagementRootCaCertificateAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *CertificatemanagementRootCaCertificateAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *CertificatemanagementRootCaCertificateAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *CertificatemanagementRootCaCertificateAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *CertificatemanagementRootCaCertificateAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CertificatemanagementRootCaCertificateAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CertificatemanagementRootCaCertificateAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCertificateName

`func (o *CertificatemanagementRootCaCertificateAllOf) GetCertificateName() string`

GetCertificateName returns the CertificateName field if non-nil, zero value otherwise.

### GetCertificateNameOk

`func (o *CertificatemanagementRootCaCertificateAllOf) GetCertificateNameOk() (*string, bool)`

GetCertificateNameOk returns a tuple with the CertificateName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateName

`func (o *CertificatemanagementRootCaCertificateAllOf) SetCertificateName(v string)`

SetCertificateName sets CertificateName field to given value.

### HasCertificateName

`func (o *CertificatemanagementRootCaCertificateAllOf) HasCertificateName() bool`

HasCertificateName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


