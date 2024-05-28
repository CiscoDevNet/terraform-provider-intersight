# ResourceSelectionCriteria

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "resource.SelectionCriteria"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "resource.SelectionCriteria"]
**Description** | Pointer to **string** | The description of the Resource Selection Criteria. | [optional] 
**Name** | Pointer to **string** | Name of the Resource Selection Criteria. | [optional] 
**PlaceHolders** | Pointer to **[]string** |  | [optional] 
**Selectors** | Pointer to [**[]ResourceSelector**](ResourceSelector.md) |  | [optional] 
**Organization** | Pointer to [**NullableOrganizationOrganizationRelationship**](OrganizationOrganizationRelationship.md) |  | [optional] 

## Methods

### NewResourceSelectionCriteria

`func NewResourceSelectionCriteria(classId string, objectType string, ) *ResourceSelectionCriteria`

NewResourceSelectionCriteria instantiates a new ResourceSelectionCriteria object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceSelectionCriteriaWithDefaults

`func NewResourceSelectionCriteriaWithDefaults() *ResourceSelectionCriteria`

NewResourceSelectionCriteriaWithDefaults instantiates a new ResourceSelectionCriteria object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ResourceSelectionCriteria) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ResourceSelectionCriteria) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ResourceSelectionCriteria) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ResourceSelectionCriteria) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ResourceSelectionCriteria) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ResourceSelectionCriteria) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDescription

`func (o *ResourceSelectionCriteria) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ResourceSelectionCriteria) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ResourceSelectionCriteria) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ResourceSelectionCriteria) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *ResourceSelectionCriteria) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ResourceSelectionCriteria) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ResourceSelectionCriteria) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ResourceSelectionCriteria) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPlaceHolders

`func (o *ResourceSelectionCriteria) GetPlaceHolders() []string`

GetPlaceHolders returns the PlaceHolders field if non-nil, zero value otherwise.

### GetPlaceHoldersOk

`func (o *ResourceSelectionCriteria) GetPlaceHoldersOk() (*[]string, bool)`

GetPlaceHoldersOk returns a tuple with the PlaceHolders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaceHolders

`func (o *ResourceSelectionCriteria) SetPlaceHolders(v []string)`

SetPlaceHolders sets PlaceHolders field to given value.

### HasPlaceHolders

`func (o *ResourceSelectionCriteria) HasPlaceHolders() bool`

HasPlaceHolders returns a boolean if a field has been set.

### SetPlaceHoldersNil

`func (o *ResourceSelectionCriteria) SetPlaceHoldersNil(b bool)`

 SetPlaceHoldersNil sets the value for PlaceHolders to be an explicit nil

### UnsetPlaceHolders
`func (o *ResourceSelectionCriteria) UnsetPlaceHolders()`

UnsetPlaceHolders ensures that no value is present for PlaceHolders, not even an explicit nil
### GetSelectors

`func (o *ResourceSelectionCriteria) GetSelectors() []ResourceSelector`

GetSelectors returns the Selectors field if non-nil, zero value otherwise.

### GetSelectorsOk

`func (o *ResourceSelectionCriteria) GetSelectorsOk() (*[]ResourceSelector, bool)`

GetSelectorsOk returns a tuple with the Selectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectors

`func (o *ResourceSelectionCriteria) SetSelectors(v []ResourceSelector)`

SetSelectors sets Selectors field to given value.

### HasSelectors

`func (o *ResourceSelectionCriteria) HasSelectors() bool`

HasSelectors returns a boolean if a field has been set.

### SetSelectorsNil

`func (o *ResourceSelectionCriteria) SetSelectorsNil(b bool)`

 SetSelectorsNil sets the value for Selectors to be an explicit nil

### UnsetSelectors
`func (o *ResourceSelectionCriteria) UnsetSelectors()`

UnsetSelectors ensures that no value is present for Selectors, not even an explicit nil
### GetOrganization

`func (o *ResourceSelectionCriteria) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *ResourceSelectionCriteria) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *ResourceSelectionCriteria) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *ResourceSelectionCriteria) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### SetOrganizationNil

`func (o *ResourceSelectionCriteria) SetOrganizationNil(b bool)`

 SetOrganizationNil sets the value for Organization to be an explicit nil

### UnsetOrganization
`func (o *ResourceSelectionCriteria) UnsetOrganization()`

UnsetOrganization ensures that no value is present for Organization, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


