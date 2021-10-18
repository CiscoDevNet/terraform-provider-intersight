# TerraformRunstateAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "terraform.Runstate"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "terraform.Runstate"]
**RunId** | Pointer to **string** | Run identifier for every terraform execution. | [optional] 
**StateFile** | Pointer to **string** | StateFile identifier of terraform execution. | [optional] 

## Methods

### NewTerraformRunstateAllOf

`func NewTerraformRunstateAllOf(classId string, objectType string, ) *TerraformRunstateAllOf`

NewTerraformRunstateAllOf instantiates a new TerraformRunstateAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTerraformRunstateAllOfWithDefaults

`func NewTerraformRunstateAllOfWithDefaults() *TerraformRunstateAllOf`

NewTerraformRunstateAllOfWithDefaults instantiates a new TerraformRunstateAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *TerraformRunstateAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *TerraformRunstateAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *TerraformRunstateAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *TerraformRunstateAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *TerraformRunstateAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *TerraformRunstateAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetRunId

`func (o *TerraformRunstateAllOf) GetRunId() string`

GetRunId returns the RunId field if non-nil, zero value otherwise.

### GetRunIdOk

`func (o *TerraformRunstateAllOf) GetRunIdOk() (*string, bool)`

GetRunIdOk returns a tuple with the RunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunId

`func (o *TerraformRunstateAllOf) SetRunId(v string)`

SetRunId sets RunId field to given value.

### HasRunId

`func (o *TerraformRunstateAllOf) HasRunId() bool`

HasRunId returns a boolean if a field has been set.

### GetStateFile

`func (o *TerraformRunstateAllOf) GetStateFile() string`

GetStateFile returns the StateFile field if non-nil, zero value otherwise.

### GetStateFileOk

`func (o *TerraformRunstateAllOf) GetStateFileOk() (*string, bool)`

GetStateFileOk returns a tuple with the StateFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateFile

`func (o *TerraformRunstateAllOf) SetStateFile(v string)`

SetStateFile sets StateFile field to given value.

### HasStateFile

`func (o *TerraformRunstateAllOf) HasStateFile() bool`

HasStateFile returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


