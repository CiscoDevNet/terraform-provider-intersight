# OrganizationOrganization

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "organization.Organization"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "organization.Organization"]
**Description** | Pointer to **string** | The informative description about the usage of this organization. | [optional] 
**Name** | Pointer to **string** | The name of the organization. There can be multiple organizations under an account. | [optional] 
**Account** | Pointer to [**IamAccountRelationship**](IamAccountRelationship.md) |  | [optional] 
**ResourceGroups** | Pointer to [**[]ResourceGroupRelationship**](ResourceGroupRelationship.md) | An array of relationships to resourceGroup resources. | [optional] 

## Methods

### NewOrganizationOrganization

`func NewOrganizationOrganization(classId string, objectType string, ) *OrganizationOrganization`

NewOrganizationOrganization instantiates a new OrganizationOrganization object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationOrganizationWithDefaults

`func NewOrganizationOrganizationWithDefaults() *OrganizationOrganization`

NewOrganizationOrganizationWithDefaults instantiates a new OrganizationOrganization object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *OrganizationOrganization) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *OrganizationOrganization) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *OrganizationOrganization) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *OrganizationOrganization) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *OrganizationOrganization) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *OrganizationOrganization) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDescription

`func (o *OrganizationOrganization) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OrganizationOrganization) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OrganizationOrganization) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *OrganizationOrganization) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *OrganizationOrganization) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OrganizationOrganization) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OrganizationOrganization) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OrganizationOrganization) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAccount

`func (o *OrganizationOrganization) GetAccount() IamAccountRelationship`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *OrganizationOrganization) GetAccountOk() (*IamAccountRelationship, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *OrganizationOrganization) SetAccount(v IamAccountRelationship)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *OrganizationOrganization) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetResourceGroups

`func (o *OrganizationOrganization) GetResourceGroups() []ResourceGroupRelationship`

GetResourceGroups returns the ResourceGroups field if non-nil, zero value otherwise.

### GetResourceGroupsOk

`func (o *OrganizationOrganization) GetResourceGroupsOk() (*[]ResourceGroupRelationship, bool)`

GetResourceGroupsOk returns a tuple with the ResourceGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceGroups

`func (o *OrganizationOrganization) SetResourceGroups(v []ResourceGroupRelationship)`

SetResourceGroups sets ResourceGroups field to given value.

### HasResourceGroups

`func (o *OrganizationOrganization) HasResourceGroups() bool`

HasResourceGroups returns a boolean if a field has been set.

### SetResourceGroupsNil

`func (o *OrganizationOrganization) SetResourceGroupsNil(b bool)`

 SetResourceGroupsNil sets the value for ResourceGroups to be an explicit nil

### UnsetResourceGroups
`func (o *OrganizationOrganization) UnsetResourceGroups()`

UnsetResourceGroups ensures that no value is present for ResourceGroups, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


