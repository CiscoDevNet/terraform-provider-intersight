# TerraformRunstate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "terraform.Runstate"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "terraform.Runstate"]
**RunId** | Pointer to **string** | Run identifier for every terraform execution. | [optional] 
**StateFile** | Pointer to **string** | StateFile identifier of terraform execution. | [optional] 

## Methods

### NewTerraformRunstate

`func NewTerraformRunstate(classId string, objectType string, ) *TerraformRunstate`

NewTerraformRunstate instantiates a new TerraformRunstate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTerraformRunstateWithDefaults

`func NewTerraformRunstateWithDefaults() *TerraformRunstate`

NewTerraformRunstateWithDefaults instantiates a new TerraformRunstate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *TerraformRunstate) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *TerraformRunstate) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *TerraformRunstate) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *TerraformRunstate) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *TerraformRunstate) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *TerraformRunstate) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetRunId

`func (o *TerraformRunstate) GetRunId() string`

GetRunId returns the RunId field if non-nil, zero value otherwise.

### GetRunIdOk

`func (o *TerraformRunstate) GetRunIdOk() (*string, bool)`

GetRunIdOk returns a tuple with the RunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunId

`func (o *TerraformRunstate) SetRunId(v string)`

SetRunId sets RunId field to given value.

### HasRunId

`func (o *TerraformRunstate) HasRunId() bool`

HasRunId returns a boolean if a field has been set.

### GetStateFile

`func (o *TerraformRunstate) GetStateFile() string`

GetStateFile returns the StateFile field if non-nil, zero value otherwise.

### GetStateFileOk

`func (o *TerraformRunstate) GetStateFileOk() (*string, bool)`

GetStateFileOk returns a tuple with the StateFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateFile

`func (o *TerraformRunstate) SetStateFile(v string)`

SetStateFile sets StateFile field to given value.

### HasStateFile

`func (o *TerraformRunstate) HasStateFile() bool`

HasStateFile returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


