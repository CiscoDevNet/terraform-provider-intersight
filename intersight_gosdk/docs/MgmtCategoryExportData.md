# MgmtCategoryExportData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "mgmt.CategoryExportData"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "mgmt.CategoryExportData"]
**ExportJobs** | Pointer to [**[]MoMoRef**](MoMoRef.md) |  | [optional] 
**Name** | Pointer to **string** | Name of the backup category. | [optional] [readonly] 
**ProcessedTypes** | Pointer to [**[]MgmtTypeWithFilter**](MgmtTypeWithFilter.md) |  | [optional] 

## Methods

### NewMgmtCategoryExportData

`func NewMgmtCategoryExportData(classId string, objectType string, ) *MgmtCategoryExportData`

NewMgmtCategoryExportData instantiates a new MgmtCategoryExportData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMgmtCategoryExportDataWithDefaults

`func NewMgmtCategoryExportDataWithDefaults() *MgmtCategoryExportData`

NewMgmtCategoryExportDataWithDefaults instantiates a new MgmtCategoryExportData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *MgmtCategoryExportData) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *MgmtCategoryExportData) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *MgmtCategoryExportData) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *MgmtCategoryExportData) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *MgmtCategoryExportData) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *MgmtCategoryExportData) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetExportJobs

`func (o *MgmtCategoryExportData) GetExportJobs() []MoMoRef`

GetExportJobs returns the ExportJobs field if non-nil, zero value otherwise.

### GetExportJobsOk

`func (o *MgmtCategoryExportData) GetExportJobsOk() (*[]MoMoRef, bool)`

GetExportJobsOk returns a tuple with the ExportJobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportJobs

`func (o *MgmtCategoryExportData) SetExportJobs(v []MoMoRef)`

SetExportJobs sets ExportJobs field to given value.

### HasExportJobs

`func (o *MgmtCategoryExportData) HasExportJobs() bool`

HasExportJobs returns a boolean if a field has been set.

### SetExportJobsNil

`func (o *MgmtCategoryExportData) SetExportJobsNil(b bool)`

 SetExportJobsNil sets the value for ExportJobs to be an explicit nil

### UnsetExportJobs
`func (o *MgmtCategoryExportData) UnsetExportJobs()`

UnsetExportJobs ensures that no value is present for ExportJobs, not even an explicit nil
### GetName

`func (o *MgmtCategoryExportData) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MgmtCategoryExportData) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MgmtCategoryExportData) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MgmtCategoryExportData) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProcessedTypes

`func (o *MgmtCategoryExportData) GetProcessedTypes() []MgmtTypeWithFilter`

GetProcessedTypes returns the ProcessedTypes field if non-nil, zero value otherwise.

### GetProcessedTypesOk

`func (o *MgmtCategoryExportData) GetProcessedTypesOk() (*[]MgmtTypeWithFilter, bool)`

GetProcessedTypesOk returns a tuple with the ProcessedTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessedTypes

`func (o *MgmtCategoryExportData) SetProcessedTypes(v []MgmtTypeWithFilter)`

SetProcessedTypes sets ProcessedTypes field to given value.

### HasProcessedTypes

`func (o *MgmtCategoryExportData) HasProcessedTypes() bool`

HasProcessedTypes returns a boolean if a field has been set.

### SetProcessedTypesNil

`func (o *MgmtCategoryExportData) SetProcessedTypesNil(b bool)`

 SetProcessedTypesNil sets the value for ProcessedTypes to be an explicit nil

### UnsetProcessedTypes
`func (o *MgmtCategoryExportData) UnsetProcessedTypes()`

UnsetProcessedTypes ensures that no value is present for ProcessedTypes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


