# ResourceSourceToPermissionResourcesHolderAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "resource.SourceToPermissionResourcesHolder"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "resource.SourceToPermissionResourcesHolder"]
**SourceToPermissionResources** | Pointer to [**[]ResourceSourceToPermissionResources**](ResourceSourceToPermissionResources.md) |  | [optional] 

## Methods

### NewResourceSourceToPermissionResourcesHolderAllOf

`func NewResourceSourceToPermissionResourcesHolderAllOf(classId string, objectType string, ) *ResourceSourceToPermissionResourcesHolderAllOf`

NewResourceSourceToPermissionResourcesHolderAllOf instantiates a new ResourceSourceToPermissionResourcesHolderAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceSourceToPermissionResourcesHolderAllOfWithDefaults

`func NewResourceSourceToPermissionResourcesHolderAllOfWithDefaults() *ResourceSourceToPermissionResourcesHolderAllOf`

NewResourceSourceToPermissionResourcesHolderAllOfWithDefaults instantiates a new ResourceSourceToPermissionResourcesHolderAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ResourceSourceToPermissionResourcesHolderAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ResourceSourceToPermissionResourcesHolderAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ResourceSourceToPermissionResourcesHolderAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ResourceSourceToPermissionResourcesHolderAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ResourceSourceToPermissionResourcesHolderAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ResourceSourceToPermissionResourcesHolderAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetSourceToPermissionResources

`func (o *ResourceSourceToPermissionResourcesHolderAllOf) GetSourceToPermissionResources() []ResourceSourceToPermissionResources`

GetSourceToPermissionResources returns the SourceToPermissionResources field if non-nil, zero value otherwise.

### GetSourceToPermissionResourcesOk

`func (o *ResourceSourceToPermissionResourcesHolderAllOf) GetSourceToPermissionResourcesOk() (*[]ResourceSourceToPermissionResources, bool)`

GetSourceToPermissionResourcesOk returns a tuple with the SourceToPermissionResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceToPermissionResources

`func (o *ResourceSourceToPermissionResourcesHolderAllOf) SetSourceToPermissionResources(v []ResourceSourceToPermissionResources)`

SetSourceToPermissionResources sets SourceToPermissionResources field to given value.

### HasSourceToPermissionResources

`func (o *ResourceSourceToPermissionResourcesHolderAllOf) HasSourceToPermissionResources() bool`

HasSourceToPermissionResources returns a boolean if a field has been set.

### SetSourceToPermissionResourcesNil

`func (o *ResourceSourceToPermissionResourcesHolderAllOf) SetSourceToPermissionResourcesNil(b bool)`

 SetSourceToPermissionResourcesNil sets the value for SourceToPermissionResources to be an explicit nil

### UnsetSourceToPermissionResources
`func (o *ResourceSourceToPermissionResourcesHolderAllOf) UnsetSourceToPermissionResources()`

UnsetSourceToPermissionResources ensures that no value is present for SourceToPermissionResources, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


