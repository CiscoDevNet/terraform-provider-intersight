# CapabilityCertificatePropertyConstraints

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "capability.CertificatePropertyConstraints"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "capability.CertificatePropertyConstraints"]
**NumOfRootCertificatesSupported** | Pointer to **int64** | Maximum number of root Certificates supported in FI. | [optional] [readonly] 

## Methods

### NewCapabilityCertificatePropertyConstraints

`func NewCapabilityCertificatePropertyConstraints(classId string, objectType string, ) *CapabilityCertificatePropertyConstraints`

NewCapabilityCertificatePropertyConstraints instantiates a new CapabilityCertificatePropertyConstraints object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCapabilityCertificatePropertyConstraintsWithDefaults

`func NewCapabilityCertificatePropertyConstraintsWithDefaults() *CapabilityCertificatePropertyConstraints`

NewCapabilityCertificatePropertyConstraintsWithDefaults instantiates a new CapabilityCertificatePropertyConstraints object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *CapabilityCertificatePropertyConstraints) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *CapabilityCertificatePropertyConstraints) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *CapabilityCertificatePropertyConstraints) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *CapabilityCertificatePropertyConstraints) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CapabilityCertificatePropertyConstraints) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CapabilityCertificatePropertyConstraints) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetNumOfRootCertificatesSupported

`func (o *CapabilityCertificatePropertyConstraints) GetNumOfRootCertificatesSupported() int64`

GetNumOfRootCertificatesSupported returns the NumOfRootCertificatesSupported field if non-nil, zero value otherwise.

### GetNumOfRootCertificatesSupportedOk

`func (o *CapabilityCertificatePropertyConstraints) GetNumOfRootCertificatesSupportedOk() (*int64, bool)`

GetNumOfRootCertificatesSupportedOk returns a tuple with the NumOfRootCertificatesSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumOfRootCertificatesSupported

`func (o *CapabilityCertificatePropertyConstraints) SetNumOfRootCertificatesSupported(v int64)`

SetNumOfRootCertificatesSupported sets NumOfRootCertificatesSupported field to given value.

### HasNumOfRootCertificatesSupported

`func (o *CapabilityCertificatePropertyConstraints) HasNumOfRootCertificatesSupported() bool`

HasNumOfRootCertificatesSupported returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


