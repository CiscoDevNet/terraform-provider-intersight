# ResourceSourceToPermissionResourcesHolder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "resource.SourceToPermissionResourcesHolder"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "resource.SourceToPermissionResourcesHolder"]
**SourceToPermissionResources** | Pointer to [**[]ResourceSourceToPermissionResources**](ResourceSourceToPermissionResources.md) |  | [optional] 

## Methods

### NewResourceSourceToPermissionResourcesHolder

`func NewResourceSourceToPermissionResourcesHolder(classId string, objectType string, ) *ResourceSourceToPermissionResourcesHolder`

NewResourceSourceToPermissionResourcesHolder instantiates a new ResourceSourceToPermissionResourcesHolder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceSourceToPermissionResourcesHolderWithDefaults

`func NewResourceSourceToPermissionResourcesHolderWithDefaults() *ResourceSourceToPermissionResourcesHolder`

NewResourceSourceToPermissionResourcesHolderWithDefaults instantiates a new ResourceSourceToPermissionResourcesHolder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ResourceSourceToPermissionResourcesHolder) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ResourceSourceToPermissionResourcesHolder) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ResourceSourceToPermissionResourcesHolder) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ResourceSourceToPermissionResourcesHolder) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ResourceSourceToPermissionResourcesHolder) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ResourceSourceToPermissionResourcesHolder) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetSourceToPermissionResources

`func (o *ResourceSourceToPermissionResourcesHolder) GetSourceToPermissionResources() []ResourceSourceToPermissionResources`

GetSourceToPermissionResources returns the SourceToPermissionResources field if non-nil, zero value otherwise.

### GetSourceToPermissionResourcesOk

`func (o *ResourceSourceToPermissionResourcesHolder) GetSourceToPermissionResourcesOk() (*[]ResourceSourceToPermissionResources, bool)`

GetSourceToPermissionResourcesOk returns a tuple with the SourceToPermissionResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceToPermissionResources

`func (o *ResourceSourceToPermissionResourcesHolder) SetSourceToPermissionResources(v []ResourceSourceToPermissionResources)`

SetSourceToPermissionResources sets SourceToPermissionResources field to given value.

### HasSourceToPermissionResources

`func (o *ResourceSourceToPermissionResourcesHolder) HasSourceToPermissionResources() bool`

HasSourceToPermissionResources returns a boolean if a field has been set.

### SetSourceToPermissionResourcesNil

`func (o *ResourceSourceToPermissionResourcesHolder) SetSourceToPermissionResourcesNil(b bool)`

 SetSourceToPermissionResourcesNil sets the value for SourceToPermissionResources to be an explicit nil

### UnsetSourceToPermissionResources
`func (o *ResourceSourceToPermissionResourcesHolder) UnsetSourceToPermissionResources()`

UnsetSourceToPermissionResources ensures that no value is present for SourceToPermissionResources, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


