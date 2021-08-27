# ResourceGroupAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "resource.Group"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "resource.Group"]
**Name** | Pointer to **string** | The name of this resource group. | [optional] 
**PerTypeCombinedSelector** | Pointer to [**[]ResourcePerTypeCombinedSelector**](ResourcePerTypeCombinedSelector.md) |  | [optional] 
**Qualifier** | Pointer to **string** | Qualifier shall be used to specify if we want to organize resources using multiple resource group or single For an account, resource groups can be of only one of the above types. (Both the types are mutually exclusive for an account.). * &#x60;Allow-Selectors&#x60; - Resources will be added to resource groups based on ODATA filter. Multiple resource group can be created to organize resources. * &#x60;Allow-All&#x60; - All resources will become part of the Resource Group. Only one resource group can be created to organize resources. | [optional] [default to "Allow-Selectors"]
**Selectors** | Pointer to [**[]ResourceSelector**](ResourceSelector.md) |  | [optional] 
**Account** | Pointer to [**IamAccountRelationship**](IamAccountRelationship.md) |  | [optional] 
**Organizations** | Pointer to [**[]OrganizationOrganizationRelationship**](OrganizationOrganizationRelationship.md) | An array of relationships to organizationOrganization resources. | [optional] 

## Methods

### NewResourceGroupAllOf

`func NewResourceGroupAllOf(classId string, objectType string, ) *ResourceGroupAllOf`

NewResourceGroupAllOf instantiates a new ResourceGroupAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceGroupAllOfWithDefaults

`func NewResourceGroupAllOfWithDefaults() *ResourceGroupAllOf`

NewResourceGroupAllOfWithDefaults instantiates a new ResourceGroupAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ResourceGroupAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ResourceGroupAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ResourceGroupAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ResourceGroupAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ResourceGroupAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ResourceGroupAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


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

### SetPerTypeCombinedSelectorNil

`func (o *ResourceGroupAllOf) SetPerTypeCombinedSelectorNil(b bool)`

 SetPerTypeCombinedSelectorNil sets the value for PerTypeCombinedSelector to be an explicit nil

### UnsetPerTypeCombinedSelector
`func (o *ResourceGroupAllOf) UnsetPerTypeCombinedSelector()`

UnsetPerTypeCombinedSelector ensures that no value is present for PerTypeCombinedSelector, not even an explicit nil
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

### SetSelectorsNil

`func (o *ResourceGroupAllOf) SetSelectorsNil(b bool)`

 SetSelectorsNil sets the value for Selectors to be an explicit nil

### UnsetSelectors
`func (o *ResourceGroupAllOf) UnsetSelectors()`

UnsetSelectors ensures that no value is present for Selectors, not even an explicit nil
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


