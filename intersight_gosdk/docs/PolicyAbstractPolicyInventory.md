# PolicyAbstractPolicyInventory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**Description** | Pointer to **string** | Description of the policy. | [optional] [readonly] 
**Name** | Pointer to **string** | Name of the inventoried policy object. | [optional] [readonly] 

## Methods

### NewPolicyAbstractPolicyInventory

`func NewPolicyAbstractPolicyInventory(classId string, objectType string, ) *PolicyAbstractPolicyInventory`

NewPolicyAbstractPolicyInventory instantiates a new PolicyAbstractPolicyInventory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyAbstractPolicyInventoryWithDefaults

`func NewPolicyAbstractPolicyInventoryWithDefaults() *PolicyAbstractPolicyInventory`

NewPolicyAbstractPolicyInventoryWithDefaults instantiates a new PolicyAbstractPolicyInventory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *PolicyAbstractPolicyInventory) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *PolicyAbstractPolicyInventory) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *PolicyAbstractPolicyInventory) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *PolicyAbstractPolicyInventory) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *PolicyAbstractPolicyInventory) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *PolicyAbstractPolicyInventory) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDescription

`func (o *PolicyAbstractPolicyInventory) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PolicyAbstractPolicyInventory) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PolicyAbstractPolicyInventory) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PolicyAbstractPolicyInventory) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *PolicyAbstractPolicyInventory) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PolicyAbstractPolicyInventory) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PolicyAbstractPolicyInventory) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PolicyAbstractPolicyInventory) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


