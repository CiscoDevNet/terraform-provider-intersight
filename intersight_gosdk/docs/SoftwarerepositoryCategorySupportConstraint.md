# SoftwarerepositoryCategorySupportConstraint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConstraintId** | Pointer to **string** | Identifier for this managed object. | [optional] 
**FilteredModels** | Pointer to [**[]SoftwarerepositoryConstraintModels**](softwarerepository.ConstraintModels.md) |  | [optional] 
**MdfId** | Pointer to **string** | Cisco software repository image category identifier. | [optional] 
**MinVersion** | Pointer to **string** | Minimum image version from where the models can be supported. | [optional] 
**ParseFromImageName** | Pointer to **bool** | Fields which tells if the constraint is based on image name parsing. | [optional] 
**SupportedModels** | Pointer to **[]string** |  | [optional] 

## Methods

### NewSoftwarerepositoryCategorySupportConstraint

`func NewSoftwarerepositoryCategorySupportConstraint() *SoftwarerepositoryCategorySupportConstraint`

NewSoftwarerepositoryCategorySupportConstraint instantiates a new SoftwarerepositoryCategorySupportConstraint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSoftwarerepositoryCategorySupportConstraintWithDefaults

`func NewSoftwarerepositoryCategorySupportConstraintWithDefaults() *SoftwarerepositoryCategorySupportConstraint`

NewSoftwarerepositoryCategorySupportConstraintWithDefaults instantiates a new SoftwarerepositoryCategorySupportConstraint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConstraintId

`func (o *SoftwarerepositoryCategorySupportConstraint) GetConstraintId() string`

GetConstraintId returns the ConstraintId field if non-nil, zero value otherwise.

### GetConstraintIdOk

`func (o *SoftwarerepositoryCategorySupportConstraint) GetConstraintIdOk() (*string, bool)`

GetConstraintIdOk returns a tuple with the ConstraintId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConstraintId

`func (o *SoftwarerepositoryCategorySupportConstraint) SetConstraintId(v string)`

SetConstraintId sets ConstraintId field to given value.

### HasConstraintId

`func (o *SoftwarerepositoryCategorySupportConstraint) HasConstraintId() bool`

HasConstraintId returns a boolean if a field has been set.

### GetFilteredModels

`func (o *SoftwarerepositoryCategorySupportConstraint) GetFilteredModels() []SoftwarerepositoryConstraintModels`

GetFilteredModels returns the FilteredModels field if non-nil, zero value otherwise.

### GetFilteredModelsOk

`func (o *SoftwarerepositoryCategorySupportConstraint) GetFilteredModelsOk() (*[]SoftwarerepositoryConstraintModels, bool)`

GetFilteredModelsOk returns a tuple with the FilteredModels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilteredModels

`func (o *SoftwarerepositoryCategorySupportConstraint) SetFilteredModels(v []SoftwarerepositoryConstraintModels)`

SetFilteredModels sets FilteredModels field to given value.

### HasFilteredModels

`func (o *SoftwarerepositoryCategorySupportConstraint) HasFilteredModels() bool`

HasFilteredModels returns a boolean if a field has been set.

### GetMdfId

`func (o *SoftwarerepositoryCategorySupportConstraint) GetMdfId() string`

GetMdfId returns the MdfId field if non-nil, zero value otherwise.

### GetMdfIdOk

`func (o *SoftwarerepositoryCategorySupportConstraint) GetMdfIdOk() (*string, bool)`

GetMdfIdOk returns a tuple with the MdfId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMdfId

`func (o *SoftwarerepositoryCategorySupportConstraint) SetMdfId(v string)`

SetMdfId sets MdfId field to given value.

### HasMdfId

`func (o *SoftwarerepositoryCategorySupportConstraint) HasMdfId() bool`

HasMdfId returns a boolean if a field has been set.

### GetMinVersion

`func (o *SoftwarerepositoryCategorySupportConstraint) GetMinVersion() string`

GetMinVersion returns the MinVersion field if non-nil, zero value otherwise.

### GetMinVersionOk

`func (o *SoftwarerepositoryCategorySupportConstraint) GetMinVersionOk() (*string, bool)`

GetMinVersionOk returns a tuple with the MinVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinVersion

`func (o *SoftwarerepositoryCategorySupportConstraint) SetMinVersion(v string)`

SetMinVersion sets MinVersion field to given value.

### HasMinVersion

`func (o *SoftwarerepositoryCategorySupportConstraint) HasMinVersion() bool`

HasMinVersion returns a boolean if a field has been set.

### GetParseFromImageName

`func (o *SoftwarerepositoryCategorySupportConstraint) GetParseFromImageName() bool`

GetParseFromImageName returns the ParseFromImageName field if non-nil, zero value otherwise.

### GetParseFromImageNameOk

`func (o *SoftwarerepositoryCategorySupportConstraint) GetParseFromImageNameOk() (*bool, bool)`

GetParseFromImageNameOk returns a tuple with the ParseFromImageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParseFromImageName

`func (o *SoftwarerepositoryCategorySupportConstraint) SetParseFromImageName(v bool)`

SetParseFromImageName sets ParseFromImageName field to given value.

### HasParseFromImageName

`func (o *SoftwarerepositoryCategorySupportConstraint) HasParseFromImageName() bool`

HasParseFromImageName returns a boolean if a field has been set.

### GetSupportedModels

`func (o *SoftwarerepositoryCategorySupportConstraint) GetSupportedModels() []string`

GetSupportedModels returns the SupportedModels field if non-nil, zero value otherwise.

### GetSupportedModelsOk

`func (o *SoftwarerepositoryCategorySupportConstraint) GetSupportedModelsOk() (*[]string, bool)`

GetSupportedModelsOk returns a tuple with the SupportedModels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedModels

`func (o *SoftwarerepositoryCategorySupportConstraint) SetSupportedModels(v []string)`

SetSupportedModels sets SupportedModels field to given value.

### HasSupportedModels

`func (o *SoftwarerepositoryCategorySupportConstraint) HasSupportedModels() bool`

HasSupportedModels returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


