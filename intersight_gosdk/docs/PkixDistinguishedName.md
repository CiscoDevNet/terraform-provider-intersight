# PkixDistinguishedName

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CommonName** | Pointer to **string** | A required component that identifies a person or an object. | [optional] [readonly] 
**Country** | Pointer to **[]string** |  | [optional] 
**Locality** | Pointer to **[]string** |  | [optional] 
**Organization** | Pointer to **[]string** |  | [optional] 
**OrganizationalUnit** | Pointer to **[]string** |  | [optional] 
**State** | Pointer to **[]string** |  | [optional] 

## Methods

### NewPkixDistinguishedName

`func NewPkixDistinguishedName() *PkixDistinguishedName`

NewPkixDistinguishedName instantiates a new PkixDistinguishedName object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPkixDistinguishedNameWithDefaults

`func NewPkixDistinguishedNameWithDefaults() *PkixDistinguishedName`

NewPkixDistinguishedNameWithDefaults instantiates a new PkixDistinguishedName object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommonName

`func (o *PkixDistinguishedName) GetCommonName() string`

GetCommonName returns the CommonName field if non-nil, zero value otherwise.

### GetCommonNameOk

`func (o *PkixDistinguishedName) GetCommonNameOk() (*string, bool)`

GetCommonNameOk returns a tuple with the CommonName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommonName

`func (o *PkixDistinguishedName) SetCommonName(v string)`

SetCommonName sets CommonName field to given value.

### HasCommonName

`func (o *PkixDistinguishedName) HasCommonName() bool`

HasCommonName returns a boolean if a field has been set.

### GetCountry

`func (o *PkixDistinguishedName) GetCountry() []string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *PkixDistinguishedName) GetCountryOk() (*[]string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *PkixDistinguishedName) SetCountry(v []string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *PkixDistinguishedName) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetLocality

`func (o *PkixDistinguishedName) GetLocality() []string`

GetLocality returns the Locality field if non-nil, zero value otherwise.

### GetLocalityOk

`func (o *PkixDistinguishedName) GetLocalityOk() (*[]string, bool)`

GetLocalityOk returns a tuple with the Locality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocality

`func (o *PkixDistinguishedName) SetLocality(v []string)`

SetLocality sets Locality field to given value.

### HasLocality

`func (o *PkixDistinguishedName) HasLocality() bool`

HasLocality returns a boolean if a field has been set.

### GetOrganization

`func (o *PkixDistinguishedName) GetOrganization() []string`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *PkixDistinguishedName) GetOrganizationOk() (*[]string, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *PkixDistinguishedName) SetOrganization(v []string)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *PkixDistinguishedName) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetOrganizationalUnit

`func (o *PkixDistinguishedName) GetOrganizationalUnit() []string`

GetOrganizationalUnit returns the OrganizationalUnit field if non-nil, zero value otherwise.

### GetOrganizationalUnitOk

`func (o *PkixDistinguishedName) GetOrganizationalUnitOk() (*[]string, bool)`

GetOrganizationalUnitOk returns a tuple with the OrganizationalUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationalUnit

`func (o *PkixDistinguishedName) SetOrganizationalUnit(v []string)`

SetOrganizationalUnit sets OrganizationalUnit field to given value.

### HasOrganizationalUnit

`func (o *PkixDistinguishedName) HasOrganizationalUnit() bool`

HasOrganizationalUnit returns a boolean if a field has been set.

### GetState

`func (o *PkixDistinguishedName) GetState() []string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *PkixDistinguishedName) GetStateOk() (*[]string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *PkixDistinguishedName) SetState(v []string)`

SetState sets State field to given value.

### HasState

`func (o *PkixDistinguishedName) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


