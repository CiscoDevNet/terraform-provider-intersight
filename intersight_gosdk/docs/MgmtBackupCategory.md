# MgmtBackupCategory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "mgmt.BackupCategory"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "mgmt.BackupCategory"]
**Name** | Pointer to **string** | The name of the category. | [optional] [readonly] 
**Order** | Pointer to **int64** | The precedence or order for this category during backup operations is captured in the order property. The order will be applied within the context of the containment scope. | [optional] [readonly] 
**ReportingServices** | Pointer to **[]string** |  | [optional] 
**Scope** | Pointer to **string** | The containment scope to be used for this category of objects. * &#x60;Account&#x60; - The objects are present at the account level. * &#x60;Organization&#x60; - The objects are present within the context of an organization. | [optional] [readonly] [default to "Account"]

## Methods

### NewMgmtBackupCategory

`func NewMgmtBackupCategory(classId string, objectType string, ) *MgmtBackupCategory`

NewMgmtBackupCategory instantiates a new MgmtBackupCategory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMgmtBackupCategoryWithDefaults

`func NewMgmtBackupCategoryWithDefaults() *MgmtBackupCategory`

NewMgmtBackupCategoryWithDefaults instantiates a new MgmtBackupCategory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *MgmtBackupCategory) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *MgmtBackupCategory) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *MgmtBackupCategory) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *MgmtBackupCategory) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *MgmtBackupCategory) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *MgmtBackupCategory) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetName

`func (o *MgmtBackupCategory) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MgmtBackupCategory) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MgmtBackupCategory) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MgmtBackupCategory) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrder

`func (o *MgmtBackupCategory) GetOrder() int64`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *MgmtBackupCategory) GetOrderOk() (*int64, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *MgmtBackupCategory) SetOrder(v int64)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *MgmtBackupCategory) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetReportingServices

`func (o *MgmtBackupCategory) GetReportingServices() []string`

GetReportingServices returns the ReportingServices field if non-nil, zero value otherwise.

### GetReportingServicesOk

`func (o *MgmtBackupCategory) GetReportingServicesOk() (*[]string, bool)`

GetReportingServicesOk returns a tuple with the ReportingServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportingServices

`func (o *MgmtBackupCategory) SetReportingServices(v []string)`

SetReportingServices sets ReportingServices field to given value.

### HasReportingServices

`func (o *MgmtBackupCategory) HasReportingServices() bool`

HasReportingServices returns a boolean if a field has been set.

### SetReportingServicesNil

`func (o *MgmtBackupCategory) SetReportingServicesNil(b bool)`

 SetReportingServicesNil sets the value for ReportingServices to be an explicit nil

### UnsetReportingServices
`func (o *MgmtBackupCategory) UnsetReportingServices()`

UnsetReportingServices ensures that no value is present for ReportingServices, not even an explicit nil
### GetScope

`func (o *MgmtBackupCategory) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *MgmtBackupCategory) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *MgmtBackupCategory) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *MgmtBackupCategory) HasScope() bool`

HasScope returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


