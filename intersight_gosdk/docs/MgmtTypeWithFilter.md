# MgmtTypeWithFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "mgmt.TypeWithFilter"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "mgmt.TypeWithFilter"]
**ChildTypes** | Pointer to [**[]MgmtChildTypeInfo**](MgmtChildTypeInfo.md) |  | [optional] 
**Filter** | Pointer to **string** | Filter used to obtain a list of objects of the given object type. | [optional] [readonly] 
**Name** | Pointer to **string** | The type of the managed object. | [optional] [readonly] 
**RestPath** | Pointer to **string** | The rest path for the MO type. | [optional] [readonly] 

## Methods

### NewMgmtTypeWithFilter

`func NewMgmtTypeWithFilter(classId string, objectType string, ) *MgmtTypeWithFilter`

NewMgmtTypeWithFilter instantiates a new MgmtTypeWithFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMgmtTypeWithFilterWithDefaults

`func NewMgmtTypeWithFilterWithDefaults() *MgmtTypeWithFilter`

NewMgmtTypeWithFilterWithDefaults instantiates a new MgmtTypeWithFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *MgmtTypeWithFilter) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *MgmtTypeWithFilter) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *MgmtTypeWithFilter) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *MgmtTypeWithFilter) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *MgmtTypeWithFilter) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *MgmtTypeWithFilter) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetChildTypes

`func (o *MgmtTypeWithFilter) GetChildTypes() []MgmtChildTypeInfo`

GetChildTypes returns the ChildTypes field if non-nil, zero value otherwise.

### GetChildTypesOk

`func (o *MgmtTypeWithFilter) GetChildTypesOk() (*[]MgmtChildTypeInfo, bool)`

GetChildTypesOk returns a tuple with the ChildTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildTypes

`func (o *MgmtTypeWithFilter) SetChildTypes(v []MgmtChildTypeInfo)`

SetChildTypes sets ChildTypes field to given value.

### HasChildTypes

`func (o *MgmtTypeWithFilter) HasChildTypes() bool`

HasChildTypes returns a boolean if a field has been set.

### SetChildTypesNil

`func (o *MgmtTypeWithFilter) SetChildTypesNil(b bool)`

 SetChildTypesNil sets the value for ChildTypes to be an explicit nil

### UnsetChildTypes
`func (o *MgmtTypeWithFilter) UnsetChildTypes()`

UnsetChildTypes ensures that no value is present for ChildTypes, not even an explicit nil
### GetFilter

`func (o *MgmtTypeWithFilter) GetFilter() string`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *MgmtTypeWithFilter) GetFilterOk() (*string, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *MgmtTypeWithFilter) SetFilter(v string)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *MgmtTypeWithFilter) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetName

`func (o *MgmtTypeWithFilter) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MgmtTypeWithFilter) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MgmtTypeWithFilter) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MgmtTypeWithFilter) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRestPath

`func (o *MgmtTypeWithFilter) GetRestPath() string`

GetRestPath returns the RestPath field if non-nil, zero value otherwise.

### GetRestPathOk

`func (o *MgmtTypeWithFilter) GetRestPathOk() (*string, bool)`

GetRestPathOk returns a tuple with the RestPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestPath

`func (o *MgmtTypeWithFilter) SetRestPath(v string)`

SetRestPath sets RestPath field to given value.

### HasRestPath

`func (o *MgmtTypeWithFilter) HasRestPath() bool`

HasRestPath returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


