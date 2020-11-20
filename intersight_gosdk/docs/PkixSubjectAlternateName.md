# PkixSubjectAlternateName

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "pkix.SubjectAlternateName"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "pkix.SubjectAlternateName"]
**DnsName** | Pointer to **[]string** |  | [optional] 
**EmailAddress** | Pointer to **[]string** |  | [optional] 
**IpAddress** | Pointer to **[]string** |  | [optional] 
**Uri** | Pointer to **[]string** |  | [optional] 

## Methods

### NewPkixSubjectAlternateName

`func NewPkixSubjectAlternateName(classId string, objectType string, ) *PkixSubjectAlternateName`

NewPkixSubjectAlternateName instantiates a new PkixSubjectAlternateName object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPkixSubjectAlternateNameWithDefaults

`func NewPkixSubjectAlternateNameWithDefaults() *PkixSubjectAlternateName`

NewPkixSubjectAlternateNameWithDefaults instantiates a new PkixSubjectAlternateName object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *PkixSubjectAlternateName) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *PkixSubjectAlternateName) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *PkixSubjectAlternateName) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *PkixSubjectAlternateName) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *PkixSubjectAlternateName) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *PkixSubjectAlternateName) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDnsName

`func (o *PkixSubjectAlternateName) GetDnsName() []string`

GetDnsName returns the DnsName field if non-nil, zero value otherwise.

### GetDnsNameOk

`func (o *PkixSubjectAlternateName) GetDnsNameOk() (*[]string, bool)`

GetDnsNameOk returns a tuple with the DnsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsName

`func (o *PkixSubjectAlternateName) SetDnsName(v []string)`

SetDnsName sets DnsName field to given value.

### HasDnsName

`func (o *PkixSubjectAlternateName) HasDnsName() bool`

HasDnsName returns a boolean if a field has been set.

### SetDnsNameNil

`func (o *PkixSubjectAlternateName) SetDnsNameNil(b bool)`

 SetDnsNameNil sets the value for DnsName to be an explicit nil

### UnsetDnsName
`func (o *PkixSubjectAlternateName) UnsetDnsName()`

UnsetDnsName ensures that no value is present for DnsName, not even an explicit nil
### GetEmailAddress

`func (o *PkixSubjectAlternateName) GetEmailAddress() []string`

GetEmailAddress returns the EmailAddress field if non-nil, zero value otherwise.

### GetEmailAddressOk

`func (o *PkixSubjectAlternateName) GetEmailAddressOk() (*[]string, bool)`

GetEmailAddressOk returns a tuple with the EmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddress

`func (o *PkixSubjectAlternateName) SetEmailAddress(v []string)`

SetEmailAddress sets EmailAddress field to given value.

### HasEmailAddress

`func (o *PkixSubjectAlternateName) HasEmailAddress() bool`

HasEmailAddress returns a boolean if a field has been set.

### SetEmailAddressNil

`func (o *PkixSubjectAlternateName) SetEmailAddressNil(b bool)`

 SetEmailAddressNil sets the value for EmailAddress to be an explicit nil

### UnsetEmailAddress
`func (o *PkixSubjectAlternateName) UnsetEmailAddress()`

UnsetEmailAddress ensures that no value is present for EmailAddress, not even an explicit nil
### GetIpAddress

`func (o *PkixSubjectAlternateName) GetIpAddress() []string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *PkixSubjectAlternateName) GetIpAddressOk() (*[]string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *PkixSubjectAlternateName) SetIpAddress(v []string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *PkixSubjectAlternateName) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### SetIpAddressNil

`func (o *PkixSubjectAlternateName) SetIpAddressNil(b bool)`

 SetIpAddressNil sets the value for IpAddress to be an explicit nil

### UnsetIpAddress
`func (o *PkixSubjectAlternateName) UnsetIpAddress()`

UnsetIpAddress ensures that no value is present for IpAddress, not even an explicit nil
### GetUri

`func (o *PkixSubjectAlternateName) GetUri() []string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *PkixSubjectAlternateName) GetUriOk() (*[]string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *PkixSubjectAlternateName) SetUri(v []string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *PkixSubjectAlternateName) HasUri() bool`

HasUri returns a boolean if a field has been set.

### SetUriNil

`func (o *PkixSubjectAlternateName) SetUriNil(b bool)`

 SetUriNil sets the value for Uri to be an explicit nil

### UnsetUri
`func (o *PkixSubjectAlternateName) UnsetUri()`

UnsetUri ensures that no value is present for Uri, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


