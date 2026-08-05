# MgmtCategoryImportData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "mgmt.CategoryImportData"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "mgmt.CategoryImportData"]
**FileName** | Pointer to **string** | Name of the file that contains configurations for this category. | [optional] [readonly] 
**Name** | Pointer to **string** | Name of the backup category. | [optional] [readonly] 
**ProcessedTypes** | Pointer to [**[]MgmtTypeWithFilter**](MgmtTypeWithFilter.md) |  | [optional] 
**Status** | Pointer to **string** | Import status for this category of objects. * &#x60;Scheduled&#x60; - Backup or restore process has been scheduled. * &#x60;NotInitiated&#x60; - Backup or restore process has not been initiated. * &#x60;InProgress&#x60; - Backup or restore process is in progress. * &#x60;Failed&#x60; - Backup or restore process has failed. * &#x60;Completed&#x60; - Backup or restore process has completed. * &#x60;NotComplete&#x60; - Backup or restore process has partially imported configuration. | [optional] [readonly] [default to "Scheduled"]

## Methods

### NewMgmtCategoryImportData

`func NewMgmtCategoryImportData(classId string, objectType string, ) *MgmtCategoryImportData`

NewMgmtCategoryImportData instantiates a new MgmtCategoryImportData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMgmtCategoryImportDataWithDefaults

`func NewMgmtCategoryImportDataWithDefaults() *MgmtCategoryImportData`

NewMgmtCategoryImportDataWithDefaults instantiates a new MgmtCategoryImportData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *MgmtCategoryImportData) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *MgmtCategoryImportData) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *MgmtCategoryImportData) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *MgmtCategoryImportData) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *MgmtCategoryImportData) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *MgmtCategoryImportData) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetFileName

`func (o *MgmtCategoryImportData) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *MgmtCategoryImportData) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *MgmtCategoryImportData) SetFileName(v string)`

SetFileName sets FileName field to given value.

### HasFileName

`func (o *MgmtCategoryImportData) HasFileName() bool`

HasFileName returns a boolean if a field has been set.

### GetName

`func (o *MgmtCategoryImportData) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MgmtCategoryImportData) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MgmtCategoryImportData) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MgmtCategoryImportData) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProcessedTypes

`func (o *MgmtCategoryImportData) GetProcessedTypes() []MgmtTypeWithFilter`

GetProcessedTypes returns the ProcessedTypes field if non-nil, zero value otherwise.

### GetProcessedTypesOk

`func (o *MgmtCategoryImportData) GetProcessedTypesOk() (*[]MgmtTypeWithFilter, bool)`

GetProcessedTypesOk returns a tuple with the ProcessedTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessedTypes

`func (o *MgmtCategoryImportData) SetProcessedTypes(v []MgmtTypeWithFilter)`

SetProcessedTypes sets ProcessedTypes field to given value.

### HasProcessedTypes

`func (o *MgmtCategoryImportData) HasProcessedTypes() bool`

HasProcessedTypes returns a boolean if a field has been set.

### SetProcessedTypesNil

`func (o *MgmtCategoryImportData) SetProcessedTypesNil(b bool)`

 SetProcessedTypesNil sets the value for ProcessedTypes to be an explicit nil

### UnsetProcessedTypes
`func (o *MgmtCategoryImportData) UnsetProcessedTypes()`

UnsetProcessedTypes ensures that no value is present for ProcessedTypes, not even an explicit nil
### GetStatus

`func (o *MgmtCategoryImportData) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MgmtCategoryImportData) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MgmtCategoryImportData) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MgmtCategoryImportData) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


