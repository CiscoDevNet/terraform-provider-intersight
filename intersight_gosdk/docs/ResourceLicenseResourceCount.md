# ResourceLicenseResourceCount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "resource.LicenseResourceCount"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "resource.LicenseResourceCount"]
**LicenseType** | Pointer to **string** | Type of licensing defined for this resource group. Used for licensing group. * &#x60;Base&#x60; - Base as a License type. It is default license type. * &#x60;Essential&#x60; - Essential as a License type. * &#x60;Standard&#x60; - Standard as a License type. * &#x60;Advantage&#x60; - Advantage as a License type. * &#x60;Premier&#x60; - Premier as a License type. * &#x60;IWO-Essential&#x60; - IWO-Essential as a License type. * &#x60;IWO-Advantage&#x60; - IWO-Advantage as a License type. * &#x60;IWO-Premier&#x60; - IWO-Premier as a License type. | [optional] [readonly] [default to "Base"]
**ResourceCount** | Pointer to **int64** | The number of resource belongs to this licensing tier. | [optional] [readonly] 
**Account** | Pointer to [**IamAccountRelationship**](IamAccountRelationship.md) |  | [optional] 
**LicenseGroups** | Pointer to [**[]ResourceGroupRelationship**](ResourceGroupRelationship.md) | An array of relationships to resourceGroup resources. | [optional] [readonly] 

## Methods

### NewResourceLicenseResourceCount

`func NewResourceLicenseResourceCount(classId string, objectType string, ) *ResourceLicenseResourceCount`

NewResourceLicenseResourceCount instantiates a new ResourceLicenseResourceCount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceLicenseResourceCountWithDefaults

`func NewResourceLicenseResourceCountWithDefaults() *ResourceLicenseResourceCount`

NewResourceLicenseResourceCountWithDefaults instantiates a new ResourceLicenseResourceCount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ResourceLicenseResourceCount) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ResourceLicenseResourceCount) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ResourceLicenseResourceCount) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ResourceLicenseResourceCount) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ResourceLicenseResourceCount) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ResourceLicenseResourceCount) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetLicenseType

`func (o *ResourceLicenseResourceCount) GetLicenseType() string`

GetLicenseType returns the LicenseType field if non-nil, zero value otherwise.

### GetLicenseTypeOk

`func (o *ResourceLicenseResourceCount) GetLicenseTypeOk() (*string, bool)`

GetLicenseTypeOk returns a tuple with the LicenseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseType

`func (o *ResourceLicenseResourceCount) SetLicenseType(v string)`

SetLicenseType sets LicenseType field to given value.

### HasLicenseType

`func (o *ResourceLicenseResourceCount) HasLicenseType() bool`

HasLicenseType returns a boolean if a field has been set.

### GetResourceCount

`func (o *ResourceLicenseResourceCount) GetResourceCount() int64`

GetResourceCount returns the ResourceCount field if non-nil, zero value otherwise.

### GetResourceCountOk

`func (o *ResourceLicenseResourceCount) GetResourceCountOk() (*int64, bool)`

GetResourceCountOk returns a tuple with the ResourceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceCount

`func (o *ResourceLicenseResourceCount) SetResourceCount(v int64)`

SetResourceCount sets ResourceCount field to given value.

### HasResourceCount

`func (o *ResourceLicenseResourceCount) HasResourceCount() bool`

HasResourceCount returns a boolean if a field has been set.

### GetAccount

`func (o *ResourceLicenseResourceCount) GetAccount() IamAccountRelationship`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *ResourceLicenseResourceCount) GetAccountOk() (*IamAccountRelationship, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *ResourceLicenseResourceCount) SetAccount(v IamAccountRelationship)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *ResourceLicenseResourceCount) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetLicenseGroups

`func (o *ResourceLicenseResourceCount) GetLicenseGroups() []ResourceGroupRelationship`

GetLicenseGroups returns the LicenseGroups field if non-nil, zero value otherwise.

### GetLicenseGroupsOk

`func (o *ResourceLicenseResourceCount) GetLicenseGroupsOk() (*[]ResourceGroupRelationship, bool)`

GetLicenseGroupsOk returns a tuple with the LicenseGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseGroups

`func (o *ResourceLicenseResourceCount) SetLicenseGroups(v []ResourceGroupRelationship)`

SetLicenseGroups sets LicenseGroups field to given value.

### HasLicenseGroups

`func (o *ResourceLicenseResourceCount) HasLicenseGroups() bool`

HasLicenseGroups returns a boolean if a field has been set.

### SetLicenseGroupsNil

`func (o *ResourceLicenseResourceCount) SetLicenseGroupsNil(b bool)`

 SetLicenseGroupsNil sets the value for LicenseGroups to be an explicit nil

### UnsetLicenseGroups
`func (o *ResourceLicenseResourceCount) UnsetLicenseGroups()`

UnsetLicenseGroups ensures that no value is present for LicenseGroups, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


