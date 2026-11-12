# BlueprintResourceTypeQualifier

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "blueprint.ResourceTypeQualifier"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "blueprint.ResourceTypeQualifier"]
**OdataResourceQualifier** | Pointer to [**BlueprintOdataResourceQualifier**](BlueprintOdataResourceQualifier.md) |  | [optional] 
**ResourceType** | Pointer to **string** | The Intersight managed object type (for example, compute.Blade or equipment.Chassis) that this blueprint may bind to. | [optional] 

## Methods

### NewBlueprintResourceTypeQualifier

`func NewBlueprintResourceTypeQualifier(classId string, objectType string, ) *BlueprintResourceTypeQualifier`

NewBlueprintResourceTypeQualifier instantiates a new BlueprintResourceTypeQualifier object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlueprintResourceTypeQualifierWithDefaults

`func NewBlueprintResourceTypeQualifierWithDefaults() *BlueprintResourceTypeQualifier`

NewBlueprintResourceTypeQualifierWithDefaults instantiates a new BlueprintResourceTypeQualifier object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *BlueprintResourceTypeQualifier) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *BlueprintResourceTypeQualifier) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *BlueprintResourceTypeQualifier) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *BlueprintResourceTypeQualifier) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *BlueprintResourceTypeQualifier) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *BlueprintResourceTypeQualifier) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetOdataResourceQualifier

`func (o *BlueprintResourceTypeQualifier) GetOdataResourceQualifier() BlueprintOdataResourceQualifier`

GetOdataResourceQualifier returns the OdataResourceQualifier field if non-nil, zero value otherwise.

### GetOdataResourceQualifierOk

`func (o *BlueprintResourceTypeQualifier) GetOdataResourceQualifierOk() (*BlueprintOdataResourceQualifier, bool)`

GetOdataResourceQualifierOk returns a tuple with the OdataResourceQualifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOdataResourceQualifier

`func (o *BlueprintResourceTypeQualifier) SetOdataResourceQualifier(v BlueprintOdataResourceQualifier)`

SetOdataResourceQualifier sets OdataResourceQualifier field to given value.

### HasOdataResourceQualifier

`func (o *BlueprintResourceTypeQualifier) HasOdataResourceQualifier() bool`

HasOdataResourceQualifier returns a boolean if a field has been set.

### GetResourceType

`func (o *BlueprintResourceTypeQualifier) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *BlueprintResourceTypeQualifier) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *BlueprintResourceTypeQualifier) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *BlueprintResourceTypeQualifier) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


