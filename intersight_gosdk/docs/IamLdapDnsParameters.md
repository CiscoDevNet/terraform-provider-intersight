# IamLdapDnsParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "iam.LdapDnsParameters"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "iam.LdapDnsParameters"]
**SearchDomain** | Pointer to **string** | Domain name that acts as a source for a DNS query. | [optional] 
**SearchForest** | Pointer to **string** | Forest name that acts as a source for a DNS query. | [optional] 
**Source** | Pointer to **string** | Source of the domain name used for the DNS SRV request. * &#x60;Extracted&#x60; - The domain name extracted-domain from the login ID. * &#x60;Configured&#x60; - The configured-search domain. * &#x60;ConfiguredExtracted&#x60; - The domain name extracted from the login ID than the configured-search domain. | [optional] [default to "Extracted"]

## Methods

### NewIamLdapDnsParameters

`func NewIamLdapDnsParameters(classId string, objectType string, ) *IamLdapDnsParameters`

NewIamLdapDnsParameters instantiates a new IamLdapDnsParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamLdapDnsParametersWithDefaults

`func NewIamLdapDnsParametersWithDefaults() *IamLdapDnsParameters`

NewIamLdapDnsParametersWithDefaults instantiates a new IamLdapDnsParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *IamLdapDnsParameters) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *IamLdapDnsParameters) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *IamLdapDnsParameters) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *IamLdapDnsParameters) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *IamLdapDnsParameters) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *IamLdapDnsParameters) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetSearchDomain

`func (o *IamLdapDnsParameters) GetSearchDomain() string`

GetSearchDomain returns the SearchDomain field if non-nil, zero value otherwise.

### GetSearchDomainOk

`func (o *IamLdapDnsParameters) GetSearchDomainOk() (*string, bool)`

GetSearchDomainOk returns a tuple with the SearchDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchDomain

`func (o *IamLdapDnsParameters) SetSearchDomain(v string)`

SetSearchDomain sets SearchDomain field to given value.

### HasSearchDomain

`func (o *IamLdapDnsParameters) HasSearchDomain() bool`

HasSearchDomain returns a boolean if a field has been set.

### GetSearchForest

`func (o *IamLdapDnsParameters) GetSearchForest() string`

GetSearchForest returns the SearchForest field if non-nil, zero value otherwise.

### GetSearchForestOk

`func (o *IamLdapDnsParameters) GetSearchForestOk() (*string, bool)`

GetSearchForestOk returns a tuple with the SearchForest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchForest

`func (o *IamLdapDnsParameters) SetSearchForest(v string)`

SetSearchForest sets SearchForest field to given value.

### HasSearchForest

`func (o *IamLdapDnsParameters) HasSearchForest() bool`

HasSearchForest returns a boolean if a field has been set.

### GetSource

`func (o *IamLdapDnsParameters) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *IamLdapDnsParameters) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *IamLdapDnsParameters) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *IamLdapDnsParameters) HasSource() bool`

HasSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


