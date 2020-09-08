# ResourceGroup

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

### NewResourceGroup

`func NewResourceGroup() *ResourceGroup`

NewResourceGroup instantiates a new ResourceGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceGroupWithDefaults

`func NewResourceGroupWithDefaults() *ResourceGroup`

NewResourceGroupWithDefaults instantiates a new ResourceGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ResourceGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ResourceGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ResourceGroup) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ResourceGroup) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPerTypeCombinedSelector

`func (o *ResourceGroup) GetPerTypeCombinedSelector() []ResourcePerTypeCombinedSelector`

GetPerTypeCombinedSelector returns the PerTypeCombinedSelector field if non-nil, zero value otherwise.

### GetPerTypeCombinedSelectorOk

`func (o *ResourceGroup) GetPerTypeCombinedSelectorOk() (*[]ResourcePerTypeCombinedSelector, bool)`

GetPerTypeCombinedSelectorOk returns a tuple with the PerTypeCombinedSelector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerTypeCombinedSelector

`func (o *ResourceGroup) SetPerTypeCombinedSelector(v []ResourcePerTypeCombinedSelector)`

SetPerTypeCombinedSelector sets PerTypeCombinedSelector field to given value.

### HasPerTypeCombinedSelector

`func (o *ResourceGroup) HasPerTypeCombinedSelector() bool`

HasPerTypeCombinedSelector returns a boolean if a field has been set.

### GetQualifier

`func (o *ResourceGroup) GetQualifier() string`

GetQualifier returns the Qualifier field if non-nil, zero value otherwise.

### GetQualifierOk

`func (o *ResourceGroup) GetQualifierOk() (*string, bool)`

GetQualifierOk returns a tuple with the Qualifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQualifier

`func (o *ResourceGroup) SetQualifier(v string)`

SetQualifier sets Qualifier field to given value.

### HasQualifier

`func (o *ResourceGroup) HasQualifier() bool`

HasQualifier returns a boolean if a field has been set.

### GetSelectors

`func (o *ResourceGroup) GetSelectors() []ResourceSelector`

GetSelectors returns the Selectors field if non-nil, zero value otherwise.

### GetSelectorsOk

`func (o *ResourceGroup) GetSelectorsOk() (*[]ResourceSelector, bool)`

GetSelectorsOk returns a tuple with the Selectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectors

`func (o *ResourceGroup) SetSelectors(v []ResourceSelector)`

SetSelectors sets Selectors field to given value.

### HasSelectors

`func (o *ResourceGroup) HasSelectors() bool`

HasSelectors returns a boolean if a field has been set.

### GetAccount

`func (o *ResourceGroup) GetAccount() IamAccountRelationship`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *ResourceGroup) GetAccountOk() (*IamAccountRelationship, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *ResourceGroup) SetAccount(v IamAccountRelationship)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *ResourceGroup) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetOrganizations

`func (o *ResourceGroup) GetOrganizations() []OrganizationOrganizationRelationship`

GetOrganizations returns the Organizations field if non-nil, zero value otherwise.

### GetOrganizationsOk

`func (o *ResourceGroup) GetOrganizationsOk() (*[]OrganizationOrganizationRelationship, bool)`

GetOrganizationsOk returns a tuple with the Organizations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizations

`func (o *ResourceGroup) SetOrganizations(v []OrganizationOrganizationRelationship)`

SetOrganizations sets Organizations field to given value.

### HasOrganizations

`func (o *ResourceGroup) HasOrganizations() bool`

HasOrganizations returns a boolean if a field has been set.

### SetOrganizationsNil

`func (o *ResourceGroup) SetOrganizationsNil(b bool)`

 SetOrganizationsNil sets the value for Organizations to be an explicit nil

### UnsetOrganizations
`func (o *ResourceGroup) UnsetOrganizations()`

UnsetOrganizations ensures that no value is present for Organizations, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


