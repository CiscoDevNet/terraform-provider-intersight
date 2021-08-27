# HyperflexLicenseAllOf

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

### NewHyperflexLicenseAllOf

`func NewHyperflexLicenseAllOf(classId string, objectType string, ) *HyperflexLicenseAllOf`

NewHyperflexLicenseAllOf instantiates a new HyperflexLicenseAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexLicenseAllOfWithDefaults

`func NewHyperflexLicenseAllOfWithDefaults() *HyperflexLicenseAllOf`

NewHyperflexLicenseAllOfWithDefaults instantiates a new HyperflexLicenseAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HyperflexLicenseAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HyperflexLicenseAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HyperflexLicenseAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HyperflexLicenseAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HyperflexLicenseAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HyperflexLicenseAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetComplianceState

`func (o *HyperflexLicenseAllOf) GetComplianceState() string`

GetComplianceState returns the ComplianceState field if non-nil, zero value otherwise.

### GetComplianceStateOk

`func (o *HyperflexLicenseAllOf) GetComplianceStateOk() (*string, bool)`

GetComplianceStateOk returns a tuple with the ComplianceState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplianceState

`func (o *HyperflexLicenseAllOf) SetComplianceState(v string)`

SetComplianceState sets ComplianceState field to given value.

### HasComplianceState

`func (o *HyperflexLicenseAllOf) HasComplianceState() bool`

HasComplianceState returns a boolean if a field has been set.

### GetGetOutOfComplianceStartAt

`func (o *HyperflexLicenseAllOf) GetGetOutOfComplianceStartAt() string`

GetGetOutOfComplianceStartAt returns the GetOutOfComplianceStartAt field if non-nil, zero value otherwise.

### GetGetOutOfComplianceStartAtOk

`func (o *HyperflexLicenseAllOf) GetGetOutOfComplianceStartAtOk() (*string, bool)`

GetGetOutOfComplianceStartAtOk returns a tuple with the GetOutOfComplianceStartAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGetOutOfComplianceStartAt

`func (o *HyperflexLicenseAllOf) SetGetOutOfComplianceStartAt(v string)`

SetGetOutOfComplianceStartAt sets GetOutOfComplianceStartAt field to given value.

### HasGetOutOfComplianceStartAt

`func (o *HyperflexLicenseAllOf) HasGetOutOfComplianceStartAt() bool`

HasGetOutOfComplianceStartAt returns a boolean if a field has been set.

### GetInEvaluation

`func (o *HyperflexLicenseAllOf) GetInEvaluation() bool`

GetInEvaluation returns the InEvaluation field if non-nil, zero value otherwise.

### GetInEvaluationOk

`func (o *HyperflexLicenseAllOf) GetInEvaluationOk() (*bool, bool)`

GetInEvaluationOk returns a tuple with the InEvaluation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInEvaluation

`func (o *HyperflexLicenseAllOf) SetInEvaluation(v bool)`

SetInEvaluation sets InEvaluation field to given value.

### HasInEvaluation

`func (o *HyperflexLicenseAllOf) HasInEvaluation() bool`

HasInEvaluation returns a boolean if a field has been set.

### GetLicenseAuthorization

`func (o *HyperflexLicenseAllOf) GetLicenseAuthorization() HyperflexHxLicenseAuthorizationDetailsDt`

GetLicenseAuthorization returns the LicenseAuthorization field if non-nil, zero value otherwise.

### GetLicenseAuthorizationOk

`func (o *HyperflexLicenseAllOf) GetLicenseAuthorizationOk() (*HyperflexHxLicenseAuthorizationDetailsDt, bool)`

GetLicenseAuthorizationOk returns a tuple with the LicenseAuthorization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseAuthorization

`func (o *HyperflexLicenseAllOf) SetLicenseAuthorization(v HyperflexHxLicenseAuthorizationDetailsDt)`

SetLicenseAuthorization sets LicenseAuthorization field to given value.

### HasLicenseAuthorization

`func (o *HyperflexLicenseAllOf) HasLicenseAuthorization() bool`

HasLicenseAuthorization returns a boolean if a field has been set.

### SetLicenseAuthorizationNil

`func (o *HyperflexLicenseAllOf) SetLicenseAuthorizationNil(b bool)`

 SetLicenseAuthorizationNil sets the value for LicenseAuthorization to be an explicit nil

### UnsetLicenseAuthorization
`func (o *HyperflexLicenseAllOf) UnsetLicenseAuthorization()`

UnsetLicenseAuthorization ensures that no value is present for LicenseAuthorization, not even an explicit nil
### GetLicenseRegistration

`func (o *HyperflexLicenseAllOf) GetLicenseRegistration() HyperflexHxRegistrationDetailsDt`

GetLicenseRegistration returns the LicenseRegistration field if non-nil, zero value otherwise.

### GetLicenseRegistrationOk

`func (o *HyperflexLicenseAllOf) GetLicenseRegistrationOk() (*HyperflexHxRegistrationDetailsDt, bool)`

GetLicenseRegistrationOk returns a tuple with the LicenseRegistration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseRegistration

`func (o *HyperflexLicenseAllOf) SetLicenseRegistration(v HyperflexHxRegistrationDetailsDt)`

SetLicenseRegistration sets LicenseRegistration field to given value.

### HasLicenseRegistration

`func (o *HyperflexLicenseAllOf) HasLicenseRegistration() bool`

HasLicenseRegistration returns a boolean if a field has been set.

### SetLicenseRegistrationNil

`func (o *HyperflexLicenseAllOf) SetLicenseRegistrationNil(b bool)`

 SetLicenseRegistrationNil sets the value for LicenseRegistration to be an explicit nil

### UnsetLicenseRegistration
`func (o *HyperflexLicenseAllOf) UnsetLicenseRegistration()`

UnsetLicenseRegistration ensures that no value is present for LicenseRegistration, not even an explicit nil
### GetLicenseType

`func (o *HyperflexLicenseAllOf) GetLicenseType() string`

GetLicenseType returns the LicenseType field if non-nil, zero value otherwise.

### GetLicenseTypeOk

`func (o *HyperflexLicenseAllOf) GetLicenseTypeOk() (*string, bool)`

GetLicenseTypeOk returns a tuple with the LicenseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseType

`func (o *HyperflexLicenseAllOf) SetLicenseType(v string)`

SetLicenseType sets LicenseType field to given value.

### HasLicenseType

`func (o *HyperflexLicenseAllOf) HasLicenseType() bool`

HasLicenseType returns a boolean if a field has been set.

### GetPlrEnabled

`func (o *HyperflexLicenseAllOf) GetPlrEnabled() bool`

GetPlrEnabled returns the PlrEnabled field if non-nil, zero value otherwise.

### GetPlrEnabledOk

`func (o *HyperflexLicenseAllOf) GetPlrEnabledOk() (*bool, bool)`

GetPlrEnabledOk returns a tuple with the PlrEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlrEnabled

`func (o *HyperflexLicenseAllOf) SetPlrEnabled(v bool)`

SetPlrEnabled sets PlrEnabled field to given value.

### HasPlrEnabled

`func (o *HyperflexLicenseAllOf) HasPlrEnabled() bool`

HasPlrEnabled returns a boolean if a field has been set.

### GetSmartLicensingEnabled

`func (o *HyperflexLicenseAllOf) GetSmartLicensingEnabled() bool`

GetSmartLicensingEnabled returns the SmartLicensingEnabled field if non-nil, zero value otherwise.

### GetSmartLicensingEnabledOk

`func (o *HyperflexLicenseAllOf) GetSmartLicensingEnabledOk() (*bool, bool)`

GetSmartLicensingEnabledOk returns a tuple with the SmartLicensingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmartLicensingEnabled

`func (o *HyperflexLicenseAllOf) SetSmartLicensingEnabled(v bool)`

SetSmartLicensingEnabled sets SmartLicensingEnabled field to given value.

### HasSmartLicensingEnabled

`func (o *HyperflexLicenseAllOf) HasSmartLicensingEnabled() bool`

HasSmartLicensingEnabled returns a boolean if a field has been set.

### GetCluster

`func (o *HyperflexLicenseAllOf) GetCluster() HyperflexClusterRelationship`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *HyperflexLicenseAllOf) GetClusterOk() (*HyperflexClusterRelationship, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *HyperflexLicenseAllOf) SetCluster(v HyperflexClusterRelationship)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *HyperflexLicenseAllOf) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *HyperflexLicenseAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *HyperflexLicenseAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *HyperflexLicenseAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *HyperflexLicenseAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


