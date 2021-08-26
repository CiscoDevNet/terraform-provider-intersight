# HyperflexLicense

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hyperflex.License"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hyperflex.License"]
**ComplianceState** | Pointer to **string** | Is the cluster complaint with the license entitlements? | [optional] [readonly] 
**GetOutOfComplianceStartAt** | Pointer to **string** | Out of compliance date of the cluster | [optional] [readonly] 
**InEvaluation** | Pointer to **bool** | Is the cluster in evalution period? | [optional] [readonly] 
**LicenseAuthorization** | Pointer to [**NullableHyperflexHxLicenseAuthorizationDetailsDt**](HyperflexHxLicenseAuthorizationDetailsDt.md) |  | [optional] 
**LicenseRegistration** | Pointer to [**NullableHyperflexHxRegistrationDetailsDt**](HyperflexHxRegistrationDetailsDt.md) |  | [optional] 
**LicenseType** | Pointer to **string** | The type of license applied on the cluster | [optional] [readonly] 
**PlrEnabled** | Pointer to **bool** | Is reservation enabled for the cluster? | [optional] [readonly] 
**SmartLicensingEnabled** | Pointer to **bool** | Is Smart Licensing Enabled for this cluster? | [optional] [readonly] 
**Cluster** | Pointer to [**HyperflexClusterRelationship**](HyperflexClusterRelationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewHyperflexLicense

`func NewHyperflexLicense(classId string, objectType string, ) *HyperflexLicense`

NewHyperflexLicense instantiates a new HyperflexLicense object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexLicenseWithDefaults

`func NewHyperflexLicenseWithDefaults() *HyperflexLicense`

NewHyperflexLicenseWithDefaults instantiates a new HyperflexLicense object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HyperflexLicense) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HyperflexLicense) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HyperflexLicense) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HyperflexLicense) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HyperflexLicense) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HyperflexLicense) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetComplianceState

`func (o *HyperflexLicense) GetComplianceState() string`

GetComplianceState returns the ComplianceState field if non-nil, zero value otherwise.

### GetComplianceStateOk

`func (o *HyperflexLicense) GetComplianceStateOk() (*string, bool)`

GetComplianceStateOk returns a tuple with the ComplianceState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplianceState

`func (o *HyperflexLicense) SetComplianceState(v string)`

SetComplianceState sets ComplianceState field to given value.

### HasComplianceState

`func (o *HyperflexLicense) HasComplianceState() bool`

HasComplianceState returns a boolean if a field has been set.

### GetGetOutOfComplianceStartAt

`func (o *HyperflexLicense) GetGetOutOfComplianceStartAt() string`

GetGetOutOfComplianceStartAt returns the GetOutOfComplianceStartAt field if non-nil, zero value otherwise.

### GetGetOutOfComplianceStartAtOk

`func (o *HyperflexLicense) GetGetOutOfComplianceStartAtOk() (*string, bool)`

GetGetOutOfComplianceStartAtOk returns a tuple with the GetOutOfComplianceStartAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGetOutOfComplianceStartAt

`func (o *HyperflexLicense) SetGetOutOfComplianceStartAt(v string)`

SetGetOutOfComplianceStartAt sets GetOutOfComplianceStartAt field to given value.

### HasGetOutOfComplianceStartAt

`func (o *HyperflexLicense) HasGetOutOfComplianceStartAt() bool`

HasGetOutOfComplianceStartAt returns a boolean if a field has been set.

### GetInEvaluation

`func (o *HyperflexLicense) GetInEvaluation() bool`

GetInEvaluation returns the InEvaluation field if non-nil, zero value otherwise.

### GetInEvaluationOk

`func (o *HyperflexLicense) GetInEvaluationOk() (*bool, bool)`

GetInEvaluationOk returns a tuple with the InEvaluation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInEvaluation

`func (o *HyperflexLicense) SetInEvaluation(v bool)`

SetInEvaluation sets InEvaluation field to given value.

### HasInEvaluation

`func (o *HyperflexLicense) HasInEvaluation() bool`

HasInEvaluation returns a boolean if a field has been set.

### GetLicenseAuthorization

`func (o *HyperflexLicense) GetLicenseAuthorization() HyperflexHxLicenseAuthorizationDetailsDt`

GetLicenseAuthorization returns the LicenseAuthorization field if non-nil, zero value otherwise.

### GetLicenseAuthorizationOk

`func (o *HyperflexLicense) GetLicenseAuthorizationOk() (*HyperflexHxLicenseAuthorizationDetailsDt, bool)`

GetLicenseAuthorizationOk returns a tuple with the LicenseAuthorization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseAuthorization

`func (o *HyperflexLicense) SetLicenseAuthorization(v HyperflexHxLicenseAuthorizationDetailsDt)`

SetLicenseAuthorization sets LicenseAuthorization field to given value.

### HasLicenseAuthorization

`func (o *HyperflexLicense) HasLicenseAuthorization() bool`

HasLicenseAuthorization returns a boolean if a field has been set.

### SetLicenseAuthorizationNil

`func (o *HyperflexLicense) SetLicenseAuthorizationNil(b bool)`

 SetLicenseAuthorizationNil sets the value for LicenseAuthorization to be an explicit nil

### UnsetLicenseAuthorization
`func (o *HyperflexLicense) UnsetLicenseAuthorization()`

UnsetLicenseAuthorization ensures that no value is present for LicenseAuthorization, not even an explicit nil
### GetLicenseRegistration

`func (o *HyperflexLicense) GetLicenseRegistration() HyperflexHxRegistrationDetailsDt`

GetLicenseRegistration returns the LicenseRegistration field if non-nil, zero value otherwise.

### GetLicenseRegistrationOk

`func (o *HyperflexLicense) GetLicenseRegistrationOk() (*HyperflexHxRegistrationDetailsDt, bool)`

GetLicenseRegistrationOk returns a tuple with the LicenseRegistration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseRegistration

`func (o *HyperflexLicense) SetLicenseRegistration(v HyperflexHxRegistrationDetailsDt)`

SetLicenseRegistration sets LicenseRegistration field to given value.

### HasLicenseRegistration

`func (o *HyperflexLicense) HasLicenseRegistration() bool`

HasLicenseRegistration returns a boolean if a field has been set.

### SetLicenseRegistrationNil

`func (o *HyperflexLicense) SetLicenseRegistrationNil(b bool)`

 SetLicenseRegistrationNil sets the value for LicenseRegistration to be an explicit nil

### UnsetLicenseRegistration
`func (o *HyperflexLicense) UnsetLicenseRegistration()`

UnsetLicenseRegistration ensures that no value is present for LicenseRegistration, not even an explicit nil
### GetLicenseType

`func (o *HyperflexLicense) GetLicenseType() string`

GetLicenseType returns the LicenseType field if non-nil, zero value otherwise.

### GetLicenseTypeOk

`func (o *HyperflexLicense) GetLicenseTypeOk() (*string, bool)`

GetLicenseTypeOk returns a tuple with the LicenseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseType

`func (o *HyperflexLicense) SetLicenseType(v string)`

SetLicenseType sets LicenseType field to given value.

### HasLicenseType

`func (o *HyperflexLicense) HasLicenseType() bool`

HasLicenseType returns a boolean if a field has been set.

### GetPlrEnabled

`func (o *HyperflexLicense) GetPlrEnabled() bool`

GetPlrEnabled returns the PlrEnabled field if non-nil, zero value otherwise.

### GetPlrEnabledOk

`func (o *HyperflexLicense) GetPlrEnabledOk() (*bool, bool)`

GetPlrEnabledOk returns a tuple with the PlrEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlrEnabled

`func (o *HyperflexLicense) SetPlrEnabled(v bool)`

SetPlrEnabled sets PlrEnabled field to given value.

### HasPlrEnabled

`func (o *HyperflexLicense) HasPlrEnabled() bool`

HasPlrEnabled returns a boolean if a field has been set.

### GetSmartLicensingEnabled

`func (o *HyperflexLicense) GetSmartLicensingEnabled() bool`

GetSmartLicensingEnabled returns the SmartLicensingEnabled field if non-nil, zero value otherwise.

### GetSmartLicensingEnabledOk

`func (o *HyperflexLicense) GetSmartLicensingEnabledOk() (*bool, bool)`

GetSmartLicensingEnabledOk returns a tuple with the SmartLicensingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmartLicensingEnabled

`func (o *HyperflexLicense) SetSmartLicensingEnabled(v bool)`

SetSmartLicensingEnabled sets SmartLicensingEnabled field to given value.

### HasSmartLicensingEnabled

`func (o *HyperflexLicense) HasSmartLicensingEnabled() bool`

HasSmartLicensingEnabled returns a boolean if a field has been set.

### GetCluster

`func (o *HyperflexLicense) GetCluster() HyperflexClusterRelationship`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *HyperflexLicense) GetClusterOk() (*HyperflexClusterRelationship, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *HyperflexLicense) SetCluster(v HyperflexClusterRelationship)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *HyperflexLicense) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *HyperflexLicense) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *HyperflexLicense) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *HyperflexLicense) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *HyperflexLicense) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


