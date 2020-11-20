# PolicyConfigChange

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "policy.ConfigChange"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "policy.ConfigChange"]
**Changes** | Pointer to **[]string** |  | [optional] 
**Disruptions** | Pointer to **[]string** |  | [optional] 

## Methods

### NewPolicyConfigChange

`func NewPolicyConfigChange(classId string, objectType string, ) *PolicyConfigChange`

NewPolicyConfigChange instantiates a new PolicyConfigChange object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyConfigChangeWithDefaults

`func NewPolicyConfigChangeWithDefaults() *PolicyConfigChange`

NewPolicyConfigChangeWithDefaults instantiates a new PolicyConfigChange object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *PolicyConfigChange) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *PolicyConfigChange) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *PolicyConfigChange) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *PolicyConfigChange) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *PolicyConfigChange) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *PolicyConfigChange) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetChanges

`func (o *PolicyConfigChange) GetChanges() []string`

GetChanges returns the Changes field if non-nil, zero value otherwise.

### GetChangesOk

`func (o *PolicyConfigChange) GetChangesOk() (*[]string, bool)`

GetChangesOk returns a tuple with the Changes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChanges

`func (o *PolicyConfigChange) SetChanges(v []string)`

SetChanges sets Changes field to given value.

### HasChanges

`func (o *PolicyConfigChange) HasChanges() bool`

HasChanges returns a boolean if a field has been set.

### SetChangesNil

`func (o *PolicyConfigChange) SetChangesNil(b bool)`

 SetChangesNil sets the value for Changes to be an explicit nil

### UnsetChanges
`func (o *PolicyConfigChange) UnsetChanges()`

UnsetChanges ensures that no value is present for Changes, not even an explicit nil
### GetDisruptions

`func (o *PolicyConfigChange) GetDisruptions() []string`

GetDisruptions returns the Disruptions field if non-nil, zero value otherwise.

### GetDisruptionsOk

`func (o *PolicyConfigChange) GetDisruptionsOk() (*[]string, bool)`

GetDisruptionsOk returns a tuple with the Disruptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisruptions

`func (o *PolicyConfigChange) SetDisruptions(v []string)`

SetDisruptions sets Disruptions field to given value.

### HasDisruptions

`func (o *PolicyConfigChange) HasDisruptions() bool`

HasDisruptions returns a boolean if a field has been set.

### SetDisruptionsNil

`func (o *PolicyConfigChange) SetDisruptionsNil(b bool)`

 SetDisruptionsNil sets the value for Disruptions to be an explicit nil

### UnsetDisruptions
`func (o *PolicyConfigChange) UnsetDisruptions()`

UnsetDisruptions ensures that no value is present for Disruptions, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


