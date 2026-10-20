# NiatelemetrySmartLicense

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "niatelemetry.SmartLicense"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "niatelemetry.SmartLicense"]
**ActiveMode** | Pointer to **string** | Indicate the mode smart license is curerntly running. | [optional] 
**AuthStatus** | Pointer to **string** | Authorization status of the smart license. | [optional] 
**LicenseUdi** | Pointer to **string** | License Udi of the smart license. | [optional] 
**SmartAccount** | Pointer to **string** | Smart licensing account name in CSSM and is retrieved from CSSM after regsitration. | [optional] 

## Methods

### NewNiatelemetrySmartLicense

`func NewNiatelemetrySmartLicense(classId string, objectType string, ) *NiatelemetrySmartLicense`

NewNiatelemetrySmartLicense instantiates a new NiatelemetrySmartLicense object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiatelemetrySmartLicenseWithDefaults

`func NewNiatelemetrySmartLicenseWithDefaults() *NiatelemetrySmartLicense`

NewNiatelemetrySmartLicenseWithDefaults instantiates a new NiatelemetrySmartLicense object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NiatelemetrySmartLicense) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NiatelemetrySmartLicense) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NiatelemetrySmartLicense) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NiatelemetrySmartLicense) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NiatelemetrySmartLicense) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NiatelemetrySmartLicense) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetActiveMode

`func (o *NiatelemetrySmartLicense) GetActiveMode() string`

GetActiveMode returns the ActiveMode field if non-nil, zero value otherwise.

### GetActiveModeOk

`func (o *NiatelemetrySmartLicense) GetActiveModeOk() (*string, bool)`

GetActiveModeOk returns a tuple with the ActiveMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveMode

`func (o *NiatelemetrySmartLicense) SetActiveMode(v string)`

SetActiveMode sets ActiveMode field to given value.

### HasActiveMode

`func (o *NiatelemetrySmartLicense) HasActiveMode() bool`

HasActiveMode returns a boolean if a field has been set.

### GetAuthStatus

`func (o *NiatelemetrySmartLicense) GetAuthStatus() string`

GetAuthStatus returns the AuthStatus field if non-nil, zero value otherwise.

### GetAuthStatusOk

`func (o *NiatelemetrySmartLicense) GetAuthStatusOk() (*string, bool)`

GetAuthStatusOk returns a tuple with the AuthStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthStatus

`func (o *NiatelemetrySmartLicense) SetAuthStatus(v string)`

SetAuthStatus sets AuthStatus field to given value.

### HasAuthStatus

`func (o *NiatelemetrySmartLicense) HasAuthStatus() bool`

HasAuthStatus returns a boolean if a field has been set.

### GetLicenseUdi

`func (o *NiatelemetrySmartLicense) GetLicenseUdi() string`

GetLicenseUdi returns the LicenseUdi field if non-nil, zero value otherwise.

### GetLicenseUdiOk

`func (o *NiatelemetrySmartLicense) GetLicenseUdiOk() (*string, bool)`

GetLicenseUdiOk returns a tuple with the LicenseUdi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseUdi

`func (o *NiatelemetrySmartLicense) SetLicenseUdi(v string)`

SetLicenseUdi sets LicenseUdi field to given value.

### HasLicenseUdi

`func (o *NiatelemetrySmartLicense) HasLicenseUdi() bool`

HasLicenseUdi returns a boolean if a field has been set.

### GetSmartAccount

`func (o *NiatelemetrySmartLicense) GetSmartAccount() string`

GetSmartAccount returns the SmartAccount field if non-nil, zero value otherwise.

### GetSmartAccountOk

`func (o *NiatelemetrySmartLicense) GetSmartAccountOk() (*string, bool)`

GetSmartAccountOk returns a tuple with the SmartAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmartAccount

`func (o *NiatelemetrySmartLicense) SetSmartAccount(v string)`

SetSmartAccount sets SmartAccount field to given value.

### HasSmartAccount

`func (o *NiatelemetrySmartLicense) HasSmartAccount() bool`

HasSmartAccount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


