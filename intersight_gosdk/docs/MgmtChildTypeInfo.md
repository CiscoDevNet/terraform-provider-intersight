# MgmtChildTypeInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "mgmt.ChildTypeInfo"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "mgmt.ChildTypeInfo"]
**ChildTypes** | Pointer to [**[]MgmtChildTypeInfo**](MgmtChildTypeInfo.md) |  | [optional] 
**Name** | Pointer to **string** | The name of the child MO type. | [optional] [readonly] 
**ParentRelationName** | Pointer to **string** | The name of the relationship to the parent type. | [optional] [readonly] 
**RestPath** | Pointer to **string** | The rest path for the child MO type. | [optional] [readonly] 

## Methods

### NewMgmtChildTypeInfo

`func NewMgmtChildTypeInfo(classId string, objectType string, ) *MgmtChildTypeInfo`

NewMgmtChildTypeInfo instantiates a new MgmtChildTypeInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMgmtChildTypeInfoWithDefaults

`func NewMgmtChildTypeInfoWithDefaults() *MgmtChildTypeInfo`

NewMgmtChildTypeInfoWithDefaults instantiates a new MgmtChildTypeInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *MgmtChildTypeInfo) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *MgmtChildTypeInfo) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *MgmtChildTypeInfo) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *MgmtChildTypeInfo) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *MgmtChildTypeInfo) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *MgmtChildTypeInfo) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetChildTypes

`func (o *MgmtChildTypeInfo) GetChildTypes() []MgmtChildTypeInfo`

GetChildTypes returns the ChildTypes field if non-nil, zero value otherwise.

### GetChildTypesOk

`func (o *MgmtChildTypeInfo) GetChildTypesOk() (*[]MgmtChildTypeInfo, bool)`

GetChildTypesOk returns a tuple with the ChildTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildTypes

`func (o *MgmtChildTypeInfo) SetChildTypes(v []MgmtChildTypeInfo)`

SetChildTypes sets ChildTypes field to given value.

### HasChildTypes

`func (o *MgmtChildTypeInfo) HasChildTypes() bool`

HasChildTypes returns a boolean if a field has been set.

### SetChildTypesNil

`func (o *MgmtChildTypeInfo) SetChildTypesNil(b bool)`

 SetChildTypesNil sets the value for ChildTypes to be an explicit nil

### UnsetChildTypes
`func (o *MgmtChildTypeInfo) UnsetChildTypes()`

UnsetChildTypes ensures that no value is present for ChildTypes, not even an explicit nil
### GetName

`func (o *MgmtChildTypeInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MgmtChildTypeInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MgmtChildTypeInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MgmtChildTypeInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetParentRelationName

`func (o *MgmtChildTypeInfo) GetParentRelationName() string`

GetParentRelationName returns the ParentRelationName field if non-nil, zero value otherwise.

### GetParentRelationNameOk

`func (o *MgmtChildTypeInfo) GetParentRelationNameOk() (*string, bool)`

GetParentRelationNameOk returns a tuple with the ParentRelationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentRelationName

`func (o *MgmtChildTypeInfo) SetParentRelationName(v string)`

SetParentRelationName sets ParentRelationName field to given value.

### HasParentRelationName

`func (o *MgmtChildTypeInfo) HasParentRelationName() bool`

HasParentRelationName returns a boolean if a field has been set.

### GetRestPath

`func (o *MgmtChildTypeInfo) GetRestPath() string`

GetRestPath returns the RestPath field if non-nil, zero value otherwise.

### GetRestPathOk

`func (o *MgmtChildTypeInfo) GetRestPathOk() (*string, bool)`

GetRestPathOk returns a tuple with the RestPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestPath

`func (o *MgmtChildTypeInfo) SetRestPath(v string)`

SetRestPath sets RestPath field to given value.

### HasRestPath

`func (o *MgmtChildTypeInfo) HasRestPath() bool`

HasRestPath returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


