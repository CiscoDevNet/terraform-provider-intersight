# ApplianceGroupOpStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "appliance.GroupOpStatus"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "appliance.GroupOpStatus"]
**Description** | Pointer to **string** | Description of the group. | [optional] [readonly] 
**GroupName** | Pointer to **string** | The name of group, which includes Identity Management, Device Connector Service, Core Service, Analytics, Infra and Appliance. | [optional] [readonly] 
**OverallStatus** | Pointer to **string** | The overall API status from this group&#39;s applications. | [optional] [readonly] 
**Apps** | Pointer to [**[]ApplianceAppOpStatusRelationship**](ApplianceAppOpStatusRelationship.md) | An array of relationships to applianceAppOpStatus resources. | [optional] [readonly] 
**SystemOpStatus** | Pointer to [**NullableApplianceSystemOpStatusRelationship**](ApplianceSystemOpStatusRelationship.md) |  | [optional] 

## Methods

### NewApplianceGroupOpStatus

`func NewApplianceGroupOpStatus(classId string, objectType string, ) *ApplianceGroupOpStatus`

NewApplianceGroupOpStatus instantiates a new ApplianceGroupOpStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplianceGroupOpStatusWithDefaults

`func NewApplianceGroupOpStatusWithDefaults() *ApplianceGroupOpStatus`

NewApplianceGroupOpStatusWithDefaults instantiates a new ApplianceGroupOpStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ApplianceGroupOpStatus) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ApplianceGroupOpStatus) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ApplianceGroupOpStatus) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ApplianceGroupOpStatus) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ApplianceGroupOpStatus) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ApplianceGroupOpStatus) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDescription

`func (o *ApplianceGroupOpStatus) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ApplianceGroupOpStatus) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ApplianceGroupOpStatus) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ApplianceGroupOpStatus) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetGroupName

`func (o *ApplianceGroupOpStatus) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *ApplianceGroupOpStatus) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *ApplianceGroupOpStatus) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.

### HasGroupName

`func (o *ApplianceGroupOpStatus) HasGroupName() bool`

HasGroupName returns a boolean if a field has been set.

### GetOverallStatus

`func (o *ApplianceGroupOpStatus) GetOverallStatus() string`

GetOverallStatus returns the OverallStatus field if non-nil, zero value otherwise.

### GetOverallStatusOk

`func (o *ApplianceGroupOpStatus) GetOverallStatusOk() (*string, bool)`

GetOverallStatusOk returns a tuple with the OverallStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverallStatus

`func (o *ApplianceGroupOpStatus) SetOverallStatus(v string)`

SetOverallStatus sets OverallStatus field to given value.

### HasOverallStatus

`func (o *ApplianceGroupOpStatus) HasOverallStatus() bool`

HasOverallStatus returns a boolean if a field has been set.

### GetApps

`func (o *ApplianceGroupOpStatus) GetApps() []ApplianceAppOpStatusRelationship`

GetApps returns the Apps field if non-nil, zero value otherwise.

### GetAppsOk

`func (o *ApplianceGroupOpStatus) GetAppsOk() (*[]ApplianceAppOpStatusRelationship, bool)`

GetAppsOk returns a tuple with the Apps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApps

`func (o *ApplianceGroupOpStatus) SetApps(v []ApplianceAppOpStatusRelationship)`

SetApps sets Apps field to given value.

### HasApps

`func (o *ApplianceGroupOpStatus) HasApps() bool`

HasApps returns a boolean if a field has been set.

### SetAppsNil

`func (o *ApplianceGroupOpStatus) SetAppsNil(b bool)`

 SetAppsNil sets the value for Apps to be an explicit nil

### UnsetApps
`func (o *ApplianceGroupOpStatus) UnsetApps()`

UnsetApps ensures that no value is present for Apps, not even an explicit nil
### GetSystemOpStatus

`func (o *ApplianceGroupOpStatus) GetSystemOpStatus() ApplianceSystemOpStatusRelationship`

GetSystemOpStatus returns the SystemOpStatus field if non-nil, zero value otherwise.

### GetSystemOpStatusOk

`func (o *ApplianceGroupOpStatus) GetSystemOpStatusOk() (*ApplianceSystemOpStatusRelationship, bool)`

GetSystemOpStatusOk returns a tuple with the SystemOpStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemOpStatus

`func (o *ApplianceGroupOpStatus) SetSystemOpStatus(v ApplianceSystemOpStatusRelationship)`

SetSystemOpStatus sets SystemOpStatus field to given value.

### HasSystemOpStatus

`func (o *ApplianceGroupOpStatus) HasSystemOpStatus() bool`

HasSystemOpStatus returns a boolean if a field has been set.

### SetSystemOpStatusNil

`func (o *ApplianceGroupOpStatus) SetSystemOpStatusNil(b bool)`

 SetSystemOpStatusNil sets the value for SystemOpStatus to be an explicit nil

### UnsetSystemOpStatus
`func (o *ApplianceGroupOpStatus) UnsetSystemOpStatus()`

UnsetSystemOpStatus ensures that no value is present for SystemOpStatus, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


