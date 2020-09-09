# IamLdapDnsParametersAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SearchDomain** | Pointer to **string** | Domain name that acts as a source for a DNS query. | [optional] 
**SearchForest** | Pointer to **string** | Forest name that acts as a source for a DNS query. | [optional] 
**Source** | Pointer to **string** | Source of the domain name used for the DNS SRV request. * &#x60;Extracted&#x60; - The domain name extracted-domain from the login ID. * &#x60;Configured&#x60; - The configured-search domain. * &#x60;ConfiguredExtracted&#x60; - The domain name extracted from the login ID than the configured-search domain. | [optional] [default to "Extracted"]

## Methods

### NewIamLdapDnsParametersAllOf

`func NewIamLdapDnsParametersAllOf() *IamLdapDnsParametersAllOf`

NewIamLdapDnsParametersAllOf instantiates a new IamLdapDnsParametersAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamLdapDnsParametersAllOfWithDefaults

`func NewIamLdapDnsParametersAllOfWithDefaults() *IamLdapDnsParametersAllOf`

NewIamLdapDnsParametersAllOfWithDefaults instantiates a new IamLdapDnsParametersAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSearchDomain

`func (o *IamLdapDnsParametersAllOf) GetSearchDomain() string`

GetSearchDomain returns the SearchDomain field if non-nil, zero value otherwise.

### GetSearchDomainOk

`func (o *IamLdapDnsParametersAllOf) GetSearchDomainOk() (*string, bool)`

GetSearchDomainOk returns a tuple with the SearchDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchDomain

`func (o *IamLdapDnsParametersAllOf) SetSearchDomain(v string)`

SetSearchDomain sets SearchDomain field to given value.

### HasSearchDomain

`func (o *IamLdapDnsParametersAllOf) HasSearchDomain() bool`

HasSearchDomain returns a boolean if a field has been set.

### GetSearchForest

`func (o *IamLdapDnsParametersAllOf) GetSearchForest() string`

GetSearchForest returns the SearchForest field if non-nil, zero value otherwise.

### GetSearchForestOk

`func (o *IamLdapDnsParametersAllOf) GetSearchForestOk() (*string, bool)`

GetSearchForestOk returns a tuple with the SearchForest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchForest

`func (o *IamLdapDnsParametersAllOf) SetSearchForest(v string)`

SetSearchForest sets SearchForest field to given value.

### HasSearchForest

`func (o *IamLdapDnsParametersAllOf) HasSearchForest() bool`

HasSearchForest returns a boolean if a field has been set.

### GetSource

`func (o *IamLdapDnsParametersAllOf) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *IamLdapDnsParametersAllOf) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *IamLdapDnsParametersAllOf) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *IamLdapDnsParametersAllOf) HasSource() bool`

HasSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


