# ResourceBlueprintResourceQualifier

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "resource.BlueprintResourceQualifier"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "resource.BlueprintResourceQualifier"]
**BlueprintRefName** | Pointer to **string** | The reference name of the blueprint this qualifier block applies to—the same identifier used as BlueprintRefName in composed workload definitions (dependent or nested blueprint instance). | [optional] 
**IsParentResourceQualifier** | Pointer to **bool** | When true, this qualifier block applies to the parent chassis resource (anchor blueprint). Set by the system when building qualification policy from workload deployment; not user-editable. | [optional] [readonly] 
**Qualifiers** | Pointer to [**[]ResourceResourceQualifier**](ResourceResourceQualifier.md) |  | [optional] 
**ResourceType** | Pointer to **string** | Identifies what infrastructure class the nested qualifiers apply to—for example &#x60;equipment.Chassis&#x60; for a chassis anchor, or &#x60;compute.Blade&#x60; / &#x60;compute.RackUnit&#x60; for servers under a chassis. The runtime defines allowed values and matching semantics. | [optional] 

## Methods

### NewResourceBlueprintResourceQualifier

`func NewResourceBlueprintResourceQualifier(classId string, objectType string, ) *ResourceBlueprintResourceQualifier`

NewResourceBlueprintResourceQualifier instantiates a new ResourceBlueprintResourceQualifier object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceBlueprintResourceQualifierWithDefaults

`func NewResourceBlueprintResourceQualifierWithDefaults() *ResourceBlueprintResourceQualifier`

NewResourceBlueprintResourceQualifierWithDefaults instantiates a new ResourceBlueprintResourceQualifier object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ResourceBlueprintResourceQualifier) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ResourceBlueprintResourceQualifier) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ResourceBlueprintResourceQualifier) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ResourceBlueprintResourceQualifier) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ResourceBlueprintResourceQualifier) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ResourceBlueprintResourceQualifier) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetBlueprintRefName

`func (o *ResourceBlueprintResourceQualifier) GetBlueprintRefName() string`

GetBlueprintRefName returns the BlueprintRefName field if non-nil, zero value otherwise.

### GetBlueprintRefNameOk

`func (o *ResourceBlueprintResourceQualifier) GetBlueprintRefNameOk() (*string, bool)`

GetBlueprintRefNameOk returns a tuple with the BlueprintRefName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlueprintRefName

`func (o *ResourceBlueprintResourceQualifier) SetBlueprintRefName(v string)`

SetBlueprintRefName sets BlueprintRefName field to given value.

### HasBlueprintRefName

`func (o *ResourceBlueprintResourceQualifier) HasBlueprintRefName() bool`

HasBlueprintRefName returns a boolean if a field has been set.

### GetIsParentResourceQualifier

`func (o *ResourceBlueprintResourceQualifier) GetIsParentResourceQualifier() bool`

GetIsParentResourceQualifier returns the IsParentResourceQualifier field if non-nil, zero value otherwise.

### GetIsParentResourceQualifierOk

`func (o *ResourceBlueprintResourceQualifier) GetIsParentResourceQualifierOk() (*bool, bool)`

GetIsParentResourceQualifierOk returns a tuple with the IsParentResourceQualifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsParentResourceQualifier

`func (o *ResourceBlueprintResourceQualifier) SetIsParentResourceQualifier(v bool)`

SetIsParentResourceQualifier sets IsParentResourceQualifier field to given value.

### HasIsParentResourceQualifier

`func (o *ResourceBlueprintResourceQualifier) HasIsParentResourceQualifier() bool`

HasIsParentResourceQualifier returns a boolean if a field has been set.

### GetQualifiers

`func (o *ResourceBlueprintResourceQualifier) GetQualifiers() []ResourceResourceQualifier`

GetQualifiers returns the Qualifiers field if non-nil, zero value otherwise.

### GetQualifiersOk

`func (o *ResourceBlueprintResourceQualifier) GetQualifiersOk() (*[]ResourceResourceQualifier, bool)`

GetQualifiersOk returns a tuple with the Qualifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQualifiers

`func (o *ResourceBlueprintResourceQualifier) SetQualifiers(v []ResourceResourceQualifier)`

SetQualifiers sets Qualifiers field to given value.

### HasQualifiers

`func (o *ResourceBlueprintResourceQualifier) HasQualifiers() bool`

HasQualifiers returns a boolean if a field has been set.

### SetQualifiersNil

`func (o *ResourceBlueprintResourceQualifier) SetQualifiersNil(b bool)`

 SetQualifiersNil sets the value for Qualifiers to be an explicit nil

### UnsetQualifiers
`func (o *ResourceBlueprintResourceQualifier) UnsetQualifiers()`

UnsetQualifiers ensures that no value is present for Qualifiers, not even an explicit nil
### GetResourceType

`func (o *ResourceBlueprintResourceQualifier) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *ResourceBlueprintResourceQualifier) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *ResourceBlueprintResourceQualifier) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *ResourceBlueprintResourceQualifier) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


