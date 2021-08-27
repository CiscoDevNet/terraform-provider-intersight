# OrganizationOrganizationAllOf

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

### NewOrganizationOrganizationAllOf

`func NewOrganizationOrganizationAllOf(classId string, objectType string, ) *OrganizationOrganizationAllOf`

NewOrganizationOrganizationAllOf instantiates a new OrganizationOrganizationAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationOrganizationAllOfWithDefaults

`func NewOrganizationOrganizationAllOfWithDefaults() *OrganizationOrganizationAllOf`

NewOrganizationOrganizationAllOfWithDefaults instantiates a new OrganizationOrganizationAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *OrganizationOrganizationAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *OrganizationOrganizationAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *OrganizationOrganizationAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *OrganizationOrganizationAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *OrganizationOrganizationAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *OrganizationOrganizationAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDescription

`func (o *OrganizationOrganizationAllOf) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OrganizationOrganizationAllOf) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OrganizationOrganizationAllOf) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *OrganizationOrganizationAllOf) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *OrganizationOrganizationAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OrganizationOrganizationAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OrganizationOrganizationAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OrganizationOrganizationAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAccount

`func (o *OrganizationOrganizationAllOf) GetAccount() IamAccountRelationship`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *OrganizationOrganizationAllOf) GetAccountOk() (*IamAccountRelationship, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *OrganizationOrganizationAllOf) SetAccount(v IamAccountRelationship)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *OrganizationOrganizationAllOf) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetResourceGroups

`func (o *OrganizationOrganizationAllOf) GetResourceGroups() []ResourceGroupRelationship`

GetResourceGroups returns the ResourceGroups field if non-nil, zero value otherwise.

### GetResourceGroupsOk

`func (o *OrganizationOrganizationAllOf) GetResourceGroupsOk() (*[]ResourceGroupRelationship, bool)`

GetResourceGroupsOk returns a tuple with the ResourceGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceGroups

`func (o *OrganizationOrganizationAllOf) SetResourceGroups(v []ResourceGroupRelationship)`

SetResourceGroups sets ResourceGroups field to given value.

### HasResourceGroups

`func (o *OrganizationOrganizationAllOf) HasResourceGroups() bool`

HasResourceGroups returns a boolean if a field has been set.

### SetResourceGroupsNil

`func (o *OrganizationOrganizationAllOf) SetResourceGroupsNil(b bool)`

 SetResourceGroupsNil sets the value for ResourceGroups to be an explicit nil

### UnsetResourceGroups
`func (o *OrganizationOrganizationAllOf) UnsetResourceGroups()`

UnsetResourceGroups ensures that no value is present for ResourceGroups, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


