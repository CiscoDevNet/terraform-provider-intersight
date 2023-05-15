# ResourceSelectionCriteriaAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "resource.SelectionCriteria"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "resource.SelectionCriteria"]
**Description** | Pointer to **string** | The description of the Resource Selection Criteria. | [optional] 
**Name** | Pointer to **string** | Name of the Resource Selection Criteria. | [optional] 
**PlaceHolders** | Pointer to **[]string** |  | [optional] 
**Selectors** | Pointer to [**[]ResourceSelector**](ResourceSelector.md) |  | [optional] 
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](OrganizationOrganizationRelationship.md) |  | [optional] 

## Methods

### NewResourceSelectionCriteriaAllOf

`func NewResourceSelectionCriteriaAllOf(classId string, objectType string, ) *ResourceSelectionCriteriaAllOf`

NewResourceSelectionCriteriaAllOf instantiates a new ResourceSelectionCriteriaAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceSelectionCriteriaAllOfWithDefaults

`func NewResourceSelectionCriteriaAllOfWithDefaults() *ResourceSelectionCriteriaAllOf`

NewResourceSelectionCriteriaAllOfWithDefaults instantiates a new ResourceSelectionCriteriaAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ResourceSelectionCriteriaAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ResourceSelectionCriteriaAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ResourceSelectionCriteriaAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ResourceSelectionCriteriaAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ResourceSelectionCriteriaAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ResourceSelectionCriteriaAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDescription

`func (o *ResourceSelectionCriteriaAllOf) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ResourceSelectionCriteriaAllOf) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ResourceSelectionCriteriaAllOf) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ResourceSelectionCriteriaAllOf) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *ResourceSelectionCriteriaAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ResourceSelectionCriteriaAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ResourceSelectionCriteriaAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ResourceSelectionCriteriaAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPlaceHolders

`func (o *ResourceSelectionCriteriaAllOf) GetPlaceHolders() []string`

GetPlaceHolders returns the PlaceHolders field if non-nil, zero value otherwise.

### GetPlaceHoldersOk

`func (o *ResourceSelectionCriteriaAllOf) GetPlaceHoldersOk() (*[]string, bool)`

GetPlaceHoldersOk returns a tuple with the PlaceHolders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaceHolders

`func (o *ResourceSelectionCriteriaAllOf) SetPlaceHolders(v []string)`

SetPlaceHolders sets PlaceHolders field to given value.

### HasPlaceHolders

`func (o *ResourceSelectionCriteriaAllOf) HasPlaceHolders() bool`

HasPlaceHolders returns a boolean if a field has been set.

### SetPlaceHoldersNil

`func (o *ResourceSelectionCriteriaAllOf) SetPlaceHoldersNil(b bool)`

 SetPlaceHoldersNil sets the value for PlaceHolders to be an explicit nil

### UnsetPlaceHolders
`func (o *ResourceSelectionCriteriaAllOf) UnsetPlaceHolders()`

UnsetPlaceHolders ensures that no value is present for PlaceHolders, not even an explicit nil
### GetSelectors

`func (o *ResourceSelectionCriteriaAllOf) GetSelectors() []ResourceSelector`

GetSelectors returns the Selectors field if non-nil, zero value otherwise.

### GetSelectorsOk

`func (o *ResourceSelectionCriteriaAllOf) GetSelectorsOk() (*[]ResourceSelector, bool)`

GetSelectorsOk returns a tuple with the Selectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectors

`func (o *ResourceSelectionCriteriaAllOf) SetSelectors(v []ResourceSelector)`

SetSelectors sets Selectors field to given value.

### HasSelectors

`func (o *ResourceSelectionCriteriaAllOf) HasSelectors() bool`

HasSelectors returns a boolean if a field has been set.

### SetSelectorsNil

`func (o *ResourceSelectionCriteriaAllOf) SetSelectorsNil(b bool)`

 SetSelectorsNil sets the value for Selectors to be an explicit nil

### UnsetSelectors
`func (o *ResourceSelectionCriteriaAllOf) UnsetSelectors()`

UnsetSelectors ensures that no value is present for Selectors, not even an explicit nil
### GetOrganization

`func (o *ResourceSelectionCriteriaAllOf) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *ResourceSelectionCriteriaAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *ResourceSelectionCriteriaAllOf) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *ResourceSelectionCriteriaAllOf) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


