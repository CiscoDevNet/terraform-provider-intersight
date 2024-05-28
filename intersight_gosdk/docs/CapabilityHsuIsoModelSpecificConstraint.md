# CapabilityHsuIsoModelSpecificConstraint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "capability.HsuIsoModelSpecificConstraint"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "capability.HsuIsoModelSpecificConstraint"]
**FileName** | Pointer to **string** | Name of the symbolic link to the actual iso file for this constraint. | [optional] [readonly] 
**MinVersion** | Pointer to **string** | The version below which the capability is not supported. | [optional] [readonly] 
**Model** | Pointer to **string** | The server model this constraint is to be enforced upon. | [optional] [readonly] 

## Methods

### NewCapabilityHsuIsoModelSpecificConstraint

`func NewCapabilityHsuIsoModelSpecificConstraint(classId string, objectType string, ) *CapabilityHsuIsoModelSpecificConstraint`

NewCapabilityHsuIsoModelSpecificConstraint instantiates a new CapabilityHsuIsoModelSpecificConstraint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCapabilityHsuIsoModelSpecificConstraintWithDefaults

`func NewCapabilityHsuIsoModelSpecificConstraintWithDefaults() *CapabilityHsuIsoModelSpecificConstraint`

NewCapabilityHsuIsoModelSpecificConstraintWithDefaults instantiates a new CapabilityHsuIsoModelSpecificConstraint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *CapabilityHsuIsoModelSpecificConstraint) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *CapabilityHsuIsoModelSpecificConstraint) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *CapabilityHsuIsoModelSpecificConstraint) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *CapabilityHsuIsoModelSpecificConstraint) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CapabilityHsuIsoModelSpecificConstraint) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CapabilityHsuIsoModelSpecificConstraint) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetFileName

`func (o *CapabilityHsuIsoModelSpecificConstraint) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *CapabilityHsuIsoModelSpecificConstraint) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *CapabilityHsuIsoModelSpecificConstraint) SetFileName(v string)`

SetFileName sets FileName field to given value.

### HasFileName

`func (o *CapabilityHsuIsoModelSpecificConstraint) HasFileName() bool`

HasFileName returns a boolean if a field has been set.

### GetMinVersion

`func (o *CapabilityHsuIsoModelSpecificConstraint) GetMinVersion() string`

GetMinVersion returns the MinVersion field if non-nil, zero value otherwise.

### GetMinVersionOk

`func (o *CapabilityHsuIsoModelSpecificConstraint) GetMinVersionOk() (*string, bool)`

GetMinVersionOk returns a tuple with the MinVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinVersion

`func (o *CapabilityHsuIsoModelSpecificConstraint) SetMinVersion(v string)`

SetMinVersion sets MinVersion field to given value.

### HasMinVersion

`func (o *CapabilityHsuIsoModelSpecificConstraint) HasMinVersion() bool`

HasMinVersion returns a boolean if a field has been set.

### GetModel

`func (o *CapabilityHsuIsoModelSpecificConstraint) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *CapabilityHsuIsoModelSpecificConstraint) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *CapabilityHsuIsoModelSpecificConstraint) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *CapabilityHsuIsoModelSpecificConstraint) HasModel() bool`

HasModel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


