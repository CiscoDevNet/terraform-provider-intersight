# ResourceSourceToPermissionResourcesAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "resource.SourceToPermissionResources"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "resource.SourceToPermissionResources"]
**PermissionResources** | Pointer to [**[]MoMoRef**](MoMoRef.md) |  | [optional] 
**SourceObject** | Pointer to [**MoMoRef**](MoMoRef.md) |  | [optional] 

## Methods

### NewResourceSourceToPermissionResourcesAllOf

`func NewResourceSourceToPermissionResourcesAllOf(classId string, objectType string, ) *ResourceSourceToPermissionResourcesAllOf`

NewResourceSourceToPermissionResourcesAllOf instantiates a new ResourceSourceToPermissionResourcesAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceSourceToPermissionResourcesAllOfWithDefaults

`func NewResourceSourceToPermissionResourcesAllOfWithDefaults() *ResourceSourceToPermissionResourcesAllOf`

NewResourceSourceToPermissionResourcesAllOfWithDefaults instantiates a new ResourceSourceToPermissionResourcesAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ResourceSourceToPermissionResourcesAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ResourceSourceToPermissionResourcesAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ResourceSourceToPermissionResourcesAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ResourceSourceToPermissionResourcesAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ResourceSourceToPermissionResourcesAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ResourceSourceToPermissionResourcesAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetPermissionResources

`func (o *ResourceSourceToPermissionResourcesAllOf) GetPermissionResources() []MoMoRef`

GetPermissionResources returns the PermissionResources field if non-nil, zero value otherwise.

### GetPermissionResourcesOk

`func (o *ResourceSourceToPermissionResourcesAllOf) GetPermissionResourcesOk() (*[]MoMoRef, bool)`

GetPermissionResourcesOk returns a tuple with the PermissionResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionResources

`func (o *ResourceSourceToPermissionResourcesAllOf) SetPermissionResources(v []MoMoRef)`

SetPermissionResources sets PermissionResources field to given value.

### HasPermissionResources

`func (o *ResourceSourceToPermissionResourcesAllOf) HasPermissionResources() bool`

HasPermissionResources returns a boolean if a field has been set.

### SetPermissionResourcesNil

`func (o *ResourceSourceToPermissionResourcesAllOf) SetPermissionResourcesNil(b bool)`

 SetPermissionResourcesNil sets the value for PermissionResources to be an explicit nil

### UnsetPermissionResources
`func (o *ResourceSourceToPermissionResourcesAllOf) UnsetPermissionResources()`

UnsetPermissionResources ensures that no value is present for PermissionResources, not even an explicit nil
### GetSourceObject

`func (o *ResourceSourceToPermissionResourcesAllOf) GetSourceObject() MoMoRef`

GetSourceObject returns the SourceObject field if non-nil, zero value otherwise.

### GetSourceObjectOk

`func (o *ResourceSourceToPermissionResourcesAllOf) GetSourceObjectOk() (*MoMoRef, bool)`

GetSourceObjectOk returns a tuple with the SourceObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceObject

`func (o *ResourceSourceToPermissionResourcesAllOf) SetSourceObject(v MoMoRef)`

SetSourceObject sets SourceObject field to given value.

### HasSourceObject

`func (o *ResourceSourceToPermissionResourcesAllOf) HasSourceObject() bool`

HasSourceObject returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


