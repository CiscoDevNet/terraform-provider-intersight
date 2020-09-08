# PolicyConfigChange

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Changes** | Pointer to **[]string** |  | [optional] 
**Disruptions** | Pointer to **[]string** |  | [optional] 

## Methods

### NewPolicyConfigChange

`func NewPolicyConfigChange() *PolicyConfigChange`

NewPolicyConfigChange instantiates a new PolicyConfigChange object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyConfigChangeWithDefaults

`func NewPolicyConfigChangeWithDefaults() *PolicyConfigChange`

NewPolicyConfigChangeWithDefaults instantiates a new PolicyConfigChange object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


