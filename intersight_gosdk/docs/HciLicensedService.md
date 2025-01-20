# HciLicensedService

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hci.LicensedService"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hci.LicensedService"]
**EnforcementActions** | Pointer to **[]string** |  | [optional] 
**EnforcementLevel** | Pointer to **string** | The level of enforcement applied on any violotions. | [optional] [readonly] 
**IsCompliant** | Pointer to **bool** | Indicate if the service is in-compliance with the license type. | [optional] [readonly] 
**LicenseType** | Pointer to **string** | The type of license used by the service. | [optional] [readonly] 
**Name** | Pointer to **string** | The name of the licensed service. | [optional] [readonly] 
**Violations** | Pointer to [**[]HciServiceViolation**](HciServiceViolation.md) |  | [optional] 

## Methods

### NewHciLicensedService

`func NewHciLicensedService(classId string, objectType string, ) *HciLicensedService`

NewHciLicensedService instantiates a new HciLicensedService object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHciLicensedServiceWithDefaults

`func NewHciLicensedServiceWithDefaults() *HciLicensedService`

NewHciLicensedServiceWithDefaults instantiates a new HciLicensedService object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HciLicensedService) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HciLicensedService) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HciLicensedService) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HciLicensedService) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HciLicensedService) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HciLicensedService) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetEnforcementActions

`func (o *HciLicensedService) GetEnforcementActions() []string`

GetEnforcementActions returns the EnforcementActions field if non-nil, zero value otherwise.

### GetEnforcementActionsOk

`func (o *HciLicensedService) GetEnforcementActionsOk() (*[]string, bool)`

GetEnforcementActionsOk returns a tuple with the EnforcementActions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnforcementActions

`func (o *HciLicensedService) SetEnforcementActions(v []string)`

SetEnforcementActions sets EnforcementActions field to given value.

### HasEnforcementActions

`func (o *HciLicensedService) HasEnforcementActions() bool`

HasEnforcementActions returns a boolean if a field has been set.

### SetEnforcementActionsNil

`func (o *HciLicensedService) SetEnforcementActionsNil(b bool)`

 SetEnforcementActionsNil sets the value for EnforcementActions to be an explicit nil

### UnsetEnforcementActions
`func (o *HciLicensedService) UnsetEnforcementActions()`

UnsetEnforcementActions ensures that no value is present for EnforcementActions, not even an explicit nil
### GetEnforcementLevel

`func (o *HciLicensedService) GetEnforcementLevel() string`

GetEnforcementLevel returns the EnforcementLevel field if non-nil, zero value otherwise.

### GetEnforcementLevelOk

`func (o *HciLicensedService) GetEnforcementLevelOk() (*string, bool)`

GetEnforcementLevelOk returns a tuple with the EnforcementLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnforcementLevel

`func (o *HciLicensedService) SetEnforcementLevel(v string)`

SetEnforcementLevel sets EnforcementLevel field to given value.

### HasEnforcementLevel

`func (o *HciLicensedService) HasEnforcementLevel() bool`

HasEnforcementLevel returns a boolean if a field has been set.

### GetIsCompliant

`func (o *HciLicensedService) GetIsCompliant() bool`

GetIsCompliant returns the IsCompliant field if non-nil, zero value otherwise.

### GetIsCompliantOk

`func (o *HciLicensedService) GetIsCompliantOk() (*bool, bool)`

GetIsCompliantOk returns a tuple with the IsCompliant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCompliant

`func (o *HciLicensedService) SetIsCompliant(v bool)`

SetIsCompliant sets IsCompliant field to given value.

### HasIsCompliant

`func (o *HciLicensedService) HasIsCompliant() bool`

HasIsCompliant returns a boolean if a field has been set.

### GetLicenseType

`func (o *HciLicensedService) GetLicenseType() string`

GetLicenseType returns the LicenseType field if non-nil, zero value otherwise.

### GetLicenseTypeOk

`func (o *HciLicensedService) GetLicenseTypeOk() (*string, bool)`

GetLicenseTypeOk returns a tuple with the LicenseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseType

`func (o *HciLicensedService) SetLicenseType(v string)`

SetLicenseType sets LicenseType field to given value.

### HasLicenseType

`func (o *HciLicensedService) HasLicenseType() bool`

HasLicenseType returns a boolean if a field has been set.

### GetName

`func (o *HciLicensedService) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HciLicensedService) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HciLicensedService) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *HciLicensedService) HasName() bool`

HasName returns a boolean if a field has been set.

### GetViolations

`func (o *HciLicensedService) GetViolations() []HciServiceViolation`

GetViolations returns the Violations field if non-nil, zero value otherwise.

### GetViolationsOk

`func (o *HciLicensedService) GetViolationsOk() (*[]HciServiceViolation, bool)`

GetViolationsOk returns a tuple with the Violations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViolations

`func (o *HciLicensedService) SetViolations(v []HciServiceViolation)`

SetViolations sets Violations field to given value.

### HasViolations

`func (o *HciLicensedService) HasViolations() bool`

HasViolations returns a boolean if a field has been set.

### SetViolationsNil

`func (o *HciLicensedService) SetViolationsNil(b bool)`

 SetViolationsNil sets the value for Violations to be an explicit nil

### UnsetViolations
`func (o *HciLicensedService) UnsetViolations()`

UnsetViolations ensures that no value is present for Violations, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


