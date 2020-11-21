# PkixDistinguishedName

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "pkix.DistinguishedName"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "pkix.DistinguishedName"]
**CommonName** | Pointer to **string** | A required component that identifies a person or an object. | [optional] [readonly] 
**Country** | Pointer to **[]string** |  | [optional] 
**Locality** | Pointer to **[]string** |  | [optional] 
**Organization** | Pointer to **[]string** |  | [optional] 
**OrganizationalUnit** | Pointer to **[]string** |  | [optional] 
**State** | Pointer to **[]string** |  | [optional] 

## Methods

### NewPkixDistinguishedName

`func NewPkixDistinguishedName(classId string, objectType string, ) *PkixDistinguishedName`

NewPkixDistinguishedName instantiates a new PkixDistinguishedName object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPkixDistinguishedNameWithDefaults

`func NewPkixDistinguishedNameWithDefaults() *PkixDistinguishedName`

NewPkixDistinguishedNameWithDefaults instantiates a new PkixDistinguishedName object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *PkixDistinguishedName) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *PkixDistinguishedName) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *PkixDistinguishedName) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *PkixDistinguishedName) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *PkixDistinguishedName) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *PkixDistinguishedName) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


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

### SetCountryNil

`func (o *PkixDistinguishedName) SetCountryNil(b bool)`

 SetCountryNil sets the value for Country to be an explicit nil

### UnsetCountry
`func (o *PkixDistinguishedName) UnsetCountry()`

UnsetCountry ensures that no value is present for Country, not even an explicit nil
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

### SetLocalityNil

`func (o *PkixDistinguishedName) SetLocalityNil(b bool)`

 SetLocalityNil sets the value for Locality to be an explicit nil

### UnsetLocality
`func (o *PkixDistinguishedName) UnsetLocality()`

UnsetLocality ensures that no value is present for Locality, not even an explicit nil
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

### SetOrganizationNil

`func (o *PkixDistinguishedName) SetOrganizationNil(b bool)`

 SetOrganizationNil sets the value for Organization to be an explicit nil

### UnsetOrganization
`func (o *PkixDistinguishedName) UnsetOrganization()`

UnsetOrganization ensures that no value is present for Organization, not even an explicit nil
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

### SetOrganizationalUnitNil

`func (o *PkixDistinguishedName) SetOrganizationalUnitNil(b bool)`

 SetOrganizationalUnitNil sets the value for OrganizationalUnit to be an explicit nil

### UnsetOrganizationalUnit
`func (o *PkixDistinguishedName) UnsetOrganizationalUnit()`

UnsetOrganizationalUnit ensures that no value is present for OrganizationalUnit, not even an explicit nil
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

### SetStateNil

`func (o *PkixDistinguishedName) SetStateNil(b bool)`

 SetStateNil sets the value for State to be an explicit nil

### UnsetState
`func (o *PkixDistinguishedName) UnsetState()`

UnsetState ensures that no value is present for State, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


