# SoftwarerepositoryCatalogAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the catalog. The names are populated and predefined during MO creation. | [optional] [readonly] 
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](organization.Organization.Relationship.md) |  | [optional] 
**System** | Pointer to [**IamSystemRelationship**](iam.System.Relationship.md) |  | [optional] 

## Methods

### NewSoftwarerepositoryCatalogAllOf

`func NewSoftwarerepositoryCatalogAllOf() *SoftwarerepositoryCatalogAllOf`

NewSoftwarerepositoryCatalogAllOf instantiates a new SoftwarerepositoryCatalogAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSoftwarerepositoryCatalogAllOfWithDefaults

`func NewSoftwarerepositoryCatalogAllOfWithDefaults() *SoftwarerepositoryCatalogAllOf`

NewSoftwarerepositoryCatalogAllOfWithDefaults instantiates a new SoftwarerepositoryCatalogAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SoftwarerepositoryCatalogAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SoftwarerepositoryCatalogAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SoftwarerepositoryCatalogAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SoftwarerepositoryCatalogAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrganization

`func (o *SoftwarerepositoryCatalogAllOf) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *SoftwarerepositoryCatalogAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *SoftwarerepositoryCatalogAllOf) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *SoftwarerepositoryCatalogAllOf) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetSystem

`func (o *SoftwarerepositoryCatalogAllOf) GetSystem() IamSystemRelationship`

GetSystem returns the System field if non-nil, zero value otherwise.

### GetSystemOk

`func (o *SoftwarerepositoryCatalogAllOf) GetSystemOk() (*IamSystemRelationship, bool)`

GetSystemOk returns a tuple with the System field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystem

`func (o *SoftwarerepositoryCatalogAllOf) SetSystem(v IamSystemRelationship)`

SetSystem sets System field to given value.

### HasSystem

`func (o *SoftwarerepositoryCatalogAllOf) HasSystem() bool`

HasSystem returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


