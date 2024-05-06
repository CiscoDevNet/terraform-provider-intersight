# CapabilityTemplateCatalog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "capability.TemplateCatalog"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "capability.TemplateCatalog"]
**AllowedOverrideList** | Pointer to **[]string** |  | [optional] 

## Methods

### NewCapabilityTemplateCatalog

`func NewCapabilityTemplateCatalog(classId string, objectType string, ) *CapabilityTemplateCatalog`

NewCapabilityTemplateCatalog instantiates a new CapabilityTemplateCatalog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCapabilityTemplateCatalogWithDefaults

`func NewCapabilityTemplateCatalogWithDefaults() *CapabilityTemplateCatalog`

NewCapabilityTemplateCatalogWithDefaults instantiates a new CapabilityTemplateCatalog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *CapabilityTemplateCatalog) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *CapabilityTemplateCatalog) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *CapabilityTemplateCatalog) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *CapabilityTemplateCatalog) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CapabilityTemplateCatalog) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CapabilityTemplateCatalog) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAllowedOverrideList

`func (o *CapabilityTemplateCatalog) GetAllowedOverrideList() []string`

GetAllowedOverrideList returns the AllowedOverrideList field if non-nil, zero value otherwise.

### GetAllowedOverrideListOk

`func (o *CapabilityTemplateCatalog) GetAllowedOverrideListOk() (*[]string, bool)`

GetAllowedOverrideListOk returns a tuple with the AllowedOverrideList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedOverrideList

`func (o *CapabilityTemplateCatalog) SetAllowedOverrideList(v []string)`

SetAllowedOverrideList sets AllowedOverrideList field to given value.

### HasAllowedOverrideList

`func (o *CapabilityTemplateCatalog) HasAllowedOverrideList() bool`

HasAllowedOverrideList returns a boolean if a field has been set.

### SetAllowedOverrideListNil

`func (o *CapabilityTemplateCatalog) SetAllowedOverrideListNil(b bool)`

 SetAllowedOverrideListNil sets the value for AllowedOverrideList to be an explicit nil

### UnsetAllowedOverrideList
`func (o *CapabilityTemplateCatalog) UnsetAllowedOverrideList()`

UnsetAllowedOverrideList ensures that no value is present for AllowedOverrideList, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


