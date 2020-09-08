# ResourceGroupAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of this resource group. | [optional] 
**PerTypeCombinedSelector** | Pointer to [**[]ResourcePerTypeCombinedSelector**](resource.PerTypeCombinedSelector.md) |  | [optional] 
**Qualifier** | Pointer to **string** | Qualifier shall be used to specify if we want to organize resources using multiple resource group or single For an account, resource groups can be of only one of the above types. (Both the types are mutually exclusive for an account.). * &#x60;Allow-Selectors&#x60; - Resources will be added to resource groups based on ODATA filter. Multiple resource group can be created to organize resources. * &#x60;Allow-All&#x60; - All resources will become part of the Resource Group. Only one resource group can be created to organize resources. | [optional] [default to "Allow-Selectors"]
**Selectors** | Pointer to [**[]ResourceSelector**](resource.Selector.md) |  | [optional] 
**Account** | Pointer to [**IamAccountRelationship**](iam.Account.Relationship.md) |  | [optional] 
**Organizations** | Pointer to [**[]OrganizationOrganizationRelationship**](organization.Organization.Relationship.md) | An array of relationships to organizationOrganization resources. | [optional] 

## Methods

### NewResourceGroupAllOf

`func NewResourceGroupAllOf() *ResourceGroupAllOf`

NewResourceGroupAllOf instantiates a new ResourceGroupAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceGroupAllOfWithDefaults

`func NewResourceGroupAllOfWithDefaults() *ResourceGroupAllOf`

NewResourceGroupAllOfWithDefaults instantiates a new ResourceGroupAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ResourceGroupAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ResourceGroupAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ResourceGroupAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ResourceGroupAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPerTypeCombinedSelector

`func (o *ResourceGroupAllOf) GetPerTypeCombinedSelector() []ResourcePerTypeCombinedSelector`

GetPerTypeCombinedSelector returns the PerTypeCombinedSelector field if non-nil, zero value otherwise.

### GetPerTypeCombinedSelectorOk

`func (o *ResourceGroupAllOf) GetPerTypeCombinedSelectorOk() (*[]ResourcePerTypeCombinedSelector, bool)`

GetPerTypeCombinedSelectorOk returns a tuple with the PerTypeCombinedSelector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerTypeCombinedSelector

`func (o *ResourceGroupAllOf) SetPerTypeCombinedSelector(v []ResourcePerTypeCombinedSelector)`

SetPerTypeCombinedSelector sets PerTypeCombinedSelector field to given value.

### HasPerTypeCombinedSelector

`func (o *ResourceGroupAllOf) HasPerTypeCombinedSelector() bool`

HasPerTypeCombinedSelector returns a boolean if a field has been set.

### GetQualifier

`func (o *ResourceGroupAllOf) GetQualifier() string`

GetQualifier returns the Qualifier field if non-nil, zero value otherwise.

### GetQualifierOk

`func (o *ResourceGroupAllOf) GetQualifierOk() (*string, bool)`

GetQualifierOk returns a tuple with the Qualifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQualifier

`func (o *ResourceGroupAllOf) SetQualifier(v string)`

SetQualifier sets Qualifier field to given value.

### HasQualifier

`func (o *ResourceGroupAllOf) HasQualifier() bool`

HasQualifier returns a boolean if a field has been set.

### GetSelectors

`func (o *ResourceGroupAllOf) GetSelectors() []ResourceSelector`

GetSelectors returns the Selectors field if non-nil, zero value otherwise.

### GetSelectorsOk

`func (o *ResourceGroupAllOf) GetSelectorsOk() (*[]ResourceSelector, bool)`

GetSelectorsOk returns a tuple with the Selectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectors

`func (o *ResourceGroupAllOf) SetSelectors(v []ResourceSelector)`

SetSelectors sets Selectors field to given value.

### HasSelectors

`func (o *ResourceGroupAllOf) HasSelectors() bool`

HasSelectors returns a boolean if a field has been set.

### GetAccount

`func (o *ResourceGroupAllOf) GetAccount() IamAccountRelationship`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *ResourceGroupAllOf) GetAccountOk() (*IamAccountRelationship, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *ResourceGroupAllOf) SetAccount(v IamAccountRelationship)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *ResourceGroupAllOf) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetOrganizations

`func (o *ResourceGroupAllOf) GetOrganizations() []OrganizationOrganizationRelationship`

GetOrganizations returns the Organizations field if non-nil, zero value otherwise.

### GetOrganizationsOk

`func (o *ResourceGroupAllOf) GetOrganizationsOk() (*[]OrganizationOrganizationRelationship, bool)`

GetOrganizationsOk returns a tuple with the Organizations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizations

`func (o *ResourceGroupAllOf) SetOrganizations(v []OrganizationOrganizationRelationship)`

SetOrganizations sets Organizations field to given value.

### HasOrganizations

`func (o *ResourceGroupAllOf) HasOrganizations() bool`

HasOrganizations returns a boolean if a field has been set.

### SetOrganizationsNil

`func (o *ResourceGroupAllOf) SetOrganizationsNil(b bool)`

 SetOrganizationsNil sets the value for Organizations to be an explicit nil

### UnsetOrganizations
`func (o *ResourceGroupAllOf) UnsetOrganizations()`

UnsetOrganizations ensures that no value is present for Organizations, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


