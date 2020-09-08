# ResourceLicenseResourceCountAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LicenseType** | Pointer to **string** | Type of licensing defined for this resource group. Used for licensing group. * &#x60;Base&#x60; - Base as a License type. It is default license type. * &#x60;Essential&#x60; - Essential as a License type. * &#x60;Standard&#x60; - Standard as a License type. * &#x60;Advantage&#x60; - Advantage as a License type. * &#x60;Premier&#x60; - Premier as a License type. | [optional] [readonly] [default to "Base"]
**ResourceCount** | Pointer to **int64** | The number of resource belongs to this licensing tier. | [optional] [readonly] 
**Account** | Pointer to [**IamAccountRelationship**](iam.Account.Relationship.md) |  | [optional] 
**LicenseGroups** | Pointer to [**[]ResourceGroupRelationship**](resource.Group.Relationship.md) | An array of relationships to resourceGroup resources. | [optional] [readonly] 

## Methods

### NewResourceLicenseResourceCountAllOf

`func NewResourceLicenseResourceCountAllOf() *ResourceLicenseResourceCountAllOf`

NewResourceLicenseResourceCountAllOf instantiates a new ResourceLicenseResourceCountAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceLicenseResourceCountAllOfWithDefaults

`func NewResourceLicenseResourceCountAllOfWithDefaults() *ResourceLicenseResourceCountAllOf`

NewResourceLicenseResourceCountAllOfWithDefaults instantiates a new ResourceLicenseResourceCountAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLicenseType

`func (o *ResourceLicenseResourceCountAllOf) GetLicenseType() string`

GetLicenseType returns the LicenseType field if non-nil, zero value otherwise.

### GetLicenseTypeOk

`func (o *ResourceLicenseResourceCountAllOf) GetLicenseTypeOk() (*string, bool)`

GetLicenseTypeOk returns a tuple with the LicenseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseType

`func (o *ResourceLicenseResourceCountAllOf) SetLicenseType(v string)`

SetLicenseType sets LicenseType field to given value.

### HasLicenseType

`func (o *ResourceLicenseResourceCountAllOf) HasLicenseType() bool`

HasLicenseType returns a boolean if a field has been set.

### GetResourceCount

`func (o *ResourceLicenseResourceCountAllOf) GetResourceCount() int64`

GetResourceCount returns the ResourceCount field if non-nil, zero value otherwise.

### GetResourceCountOk

`func (o *ResourceLicenseResourceCountAllOf) GetResourceCountOk() (*int64, bool)`

GetResourceCountOk returns a tuple with the ResourceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceCount

`func (o *ResourceLicenseResourceCountAllOf) SetResourceCount(v int64)`

SetResourceCount sets ResourceCount field to given value.

### HasResourceCount

`func (o *ResourceLicenseResourceCountAllOf) HasResourceCount() bool`

HasResourceCount returns a boolean if a field has been set.

### GetAccount

`func (o *ResourceLicenseResourceCountAllOf) GetAccount() IamAccountRelationship`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *ResourceLicenseResourceCountAllOf) GetAccountOk() (*IamAccountRelationship, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *ResourceLicenseResourceCountAllOf) SetAccount(v IamAccountRelationship)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *ResourceLicenseResourceCountAllOf) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetLicenseGroups

`func (o *ResourceLicenseResourceCountAllOf) GetLicenseGroups() []ResourceGroupRelationship`

GetLicenseGroups returns the LicenseGroups field if non-nil, zero value otherwise.

### GetLicenseGroupsOk

`func (o *ResourceLicenseResourceCountAllOf) GetLicenseGroupsOk() (*[]ResourceGroupRelationship, bool)`

GetLicenseGroupsOk returns a tuple with the LicenseGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseGroups

`func (o *ResourceLicenseResourceCountAllOf) SetLicenseGroups(v []ResourceGroupRelationship)`

SetLicenseGroups sets LicenseGroups field to given value.

### HasLicenseGroups

`func (o *ResourceLicenseResourceCountAllOf) HasLicenseGroups() bool`

HasLicenseGroups returns a boolean if a field has been set.

### SetLicenseGroupsNil

`func (o *ResourceLicenseResourceCountAllOf) SetLicenseGroupsNil(b bool)`

 SetLicenseGroupsNil sets the value for LicenseGroups to be an explicit nil

### UnsetLicenseGroups
`func (o *ResourceLicenseResourceCountAllOf) UnsetLicenseGroups()`

UnsetLicenseGroups ensures that no value is present for LicenseGroups, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


