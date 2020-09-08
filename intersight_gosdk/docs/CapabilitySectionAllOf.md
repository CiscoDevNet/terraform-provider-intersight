# CapabilitySectionAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** | Administrative action to initialize/load the catalog section from a particular source. * &#x60;None&#x60; - Nil value to indicate that no action is required. * &#x60;LoadLocal&#x60; - Load the catalog file packaged with the service. * &#x60;LoadIntersightRepository&#x60; - Load a catalog file hosted in the Intersight Repository. | [optional] [readonly] [default to "None"]
**CatalogName** | Pointer to **string** | The catalog name reference. | [optional] 
**Name** | Pointer to **string** | A unique name for the section inside a catalog. | [optional] 
**Source** | Pointer to **string** | The configured source for this section of the catalog. * &#x60;Local&#x60; - The catalog file is packaged with the service. * &#x60;IntersightRepository&#x60; - The catalog file is hosted in the Intersight Repository. | [optional] [readonly] [default to "Local"]
**Version** | Pointer to **string** | Version of the section inside a catalog. | [optional] 
**Catalog** | Pointer to [**CapabilityCatalogRelationship**](capability.Catalog.Relationship.md) |  | [optional] 

## Methods

### NewCapabilitySectionAllOf

`func NewCapabilitySectionAllOf() *CapabilitySectionAllOf`

NewCapabilitySectionAllOf instantiates a new CapabilitySectionAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCapabilitySectionAllOfWithDefaults

`func NewCapabilitySectionAllOfWithDefaults() *CapabilitySectionAllOf`

NewCapabilitySectionAllOfWithDefaults instantiates a new CapabilitySectionAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *CapabilitySectionAllOf) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *CapabilitySectionAllOf) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *CapabilitySectionAllOf) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *CapabilitySectionAllOf) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetCatalogName

`func (o *CapabilitySectionAllOf) GetCatalogName() string`

GetCatalogName returns the CatalogName field if non-nil, zero value otherwise.

### GetCatalogNameOk

`func (o *CapabilitySectionAllOf) GetCatalogNameOk() (*string, bool)`

GetCatalogNameOk returns a tuple with the CatalogName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalogName

`func (o *CapabilitySectionAllOf) SetCatalogName(v string)`

SetCatalogName sets CatalogName field to given value.

### HasCatalogName

`func (o *CapabilitySectionAllOf) HasCatalogName() bool`

HasCatalogName returns a boolean if a field has been set.

### GetName

`func (o *CapabilitySectionAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CapabilitySectionAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CapabilitySectionAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CapabilitySectionAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSource

`func (o *CapabilitySectionAllOf) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *CapabilitySectionAllOf) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *CapabilitySectionAllOf) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *CapabilitySectionAllOf) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetVersion

`func (o *CapabilitySectionAllOf) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *CapabilitySectionAllOf) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *CapabilitySectionAllOf) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *CapabilitySectionAllOf) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetCatalog

`func (o *CapabilitySectionAllOf) GetCatalog() CapabilityCatalogRelationship`

GetCatalog returns the Catalog field if non-nil, zero value otherwise.

### GetCatalogOk

`func (o *CapabilitySectionAllOf) GetCatalogOk() (*CapabilityCatalogRelationship, bool)`

GetCatalogOk returns a tuple with the Catalog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalog

`func (o *CapabilitySectionAllOf) SetCatalog(v CapabilityCatalogRelationship)`

SetCatalog sets Catalog field to given value.

### HasCatalog

`func (o *CapabilitySectionAllOf) HasCatalog() bool`

HasCatalog returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


