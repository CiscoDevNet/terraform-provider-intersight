# CapabilityCatalog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | A unique name for the catalog. | [optional] 
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](organization.Organization.Relationship.md) |  | [optional] 
**Sections** | Pointer to [**[]CapabilitySectionRelationship**](capability.Section.Relationship.md) | An array of relationships to capabilitySection resources. | [optional] 

## Methods

### NewCapabilityCatalog

`func NewCapabilityCatalog() *CapabilityCatalog`

NewCapabilityCatalog instantiates a new CapabilityCatalog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCapabilityCatalogWithDefaults

`func NewCapabilityCatalogWithDefaults() *CapabilityCatalog`

NewCapabilityCatalogWithDefaults instantiates a new CapabilityCatalog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CapabilityCatalog) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CapabilityCatalog) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CapabilityCatalog) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CapabilityCatalog) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrganization

`func (o *CapabilityCatalog) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *CapabilityCatalog) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *CapabilityCatalog) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *CapabilityCatalog) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetSections

`func (o *CapabilityCatalog) GetSections() []CapabilitySectionRelationship`

GetSections returns the Sections field if non-nil, zero value otherwise.

### GetSectionsOk

`func (o *CapabilityCatalog) GetSectionsOk() (*[]CapabilitySectionRelationship, bool)`

GetSectionsOk returns a tuple with the Sections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSections

`func (o *CapabilityCatalog) SetSections(v []CapabilitySectionRelationship)`

SetSections sets Sections field to given value.

### HasSections

`func (o *CapabilityCatalog) HasSections() bool`

HasSections returns a boolean if a field has been set.

### SetSectionsNil

`func (o *CapabilityCatalog) SetSectionsNil(b bool)`

 SetSectionsNil sets the value for Sections to be an explicit nil

### UnsetSections
`func (o *CapabilityCatalog) UnsetSections()`

UnsetSections ensures that no value is present for Sections, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


