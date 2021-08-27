# ApplianceGroupStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "appliance.GroupStatus"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "appliance.GroupStatus"]
**Description** | Pointer to **string** | Description of the group. | [optional] [readonly] 
**GroupName** | Pointer to **string** | The name of group, which includes Identity Management, Device Connector Service, Core Service, Analytics, Internal and Appliance. | [optional] [readonly] 
**OverallStatus** | Pointer to **string** | The overall API status from this group&#39;s applications. | [optional] [readonly] 
**Apps** | Pointer to [**[]ApplianceAppStatusRelationship**](ApplianceAppStatusRelationship.md) | An array of relationships to applianceAppStatus resources. | [optional] [readonly] 
**SystemStatus** | Pointer to [**ApplianceSystemStatusRelationship**](ApplianceSystemStatusRelationship.md) |  | [optional] 

## Methods

### NewApplianceGroupStatus

`func NewApplianceGroupStatus(classId string, objectType string, ) *ApplianceGroupStatus`

NewApplianceGroupStatus instantiates a new ApplianceGroupStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplianceGroupStatusWithDefaults

`func NewApplianceGroupStatusWithDefaults() *ApplianceGroupStatus`

NewApplianceGroupStatusWithDefaults instantiates a new ApplianceGroupStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ApplianceGroupStatus) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ApplianceGroupStatus) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ApplianceGroupStatus) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ApplianceGroupStatus) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ApplianceGroupStatus) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ApplianceGroupStatus) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDescription

`func (o *ApplianceGroupStatus) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ApplianceGroupStatus) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ApplianceGroupStatus) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ApplianceGroupStatus) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetGroupName

`func (o *ApplianceGroupStatus) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *ApplianceGroupStatus) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *ApplianceGroupStatus) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.

### HasGroupName

`func (o *ApplianceGroupStatus) HasGroupName() bool`

HasGroupName returns a boolean if a field has been set.

### GetOverallStatus

`func (o *ApplianceGroupStatus) GetOverallStatus() string`

GetOverallStatus returns the OverallStatus field if non-nil, zero value otherwise.

### GetOverallStatusOk

`func (o *ApplianceGroupStatus) GetOverallStatusOk() (*string, bool)`

GetOverallStatusOk returns a tuple with the OverallStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverallStatus

`func (o *ApplianceGroupStatus) SetOverallStatus(v string)`

SetOverallStatus sets OverallStatus field to given value.

### HasOverallStatus

`func (o *ApplianceGroupStatus) HasOverallStatus() bool`

HasOverallStatus returns a boolean if a field has been set.

### GetApps

`func (o *ApplianceGroupStatus) GetApps() []ApplianceAppStatusRelationship`

GetApps returns the Apps field if non-nil, zero value otherwise.

### GetAppsOk

`func (o *ApplianceGroupStatus) GetAppsOk() (*[]ApplianceAppStatusRelationship, bool)`

GetAppsOk returns a tuple with the Apps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApps

`func (o *ApplianceGroupStatus) SetApps(v []ApplianceAppStatusRelationship)`

SetApps sets Apps field to given value.

### HasApps

`func (o *ApplianceGroupStatus) HasApps() bool`

HasApps returns a boolean if a field has been set.

### SetAppsNil

`func (o *ApplianceGroupStatus) SetAppsNil(b bool)`

 SetAppsNil sets the value for Apps to be an explicit nil

### UnsetApps
`func (o *ApplianceGroupStatus) UnsetApps()`

UnsetApps ensures that no value is present for Apps, not even an explicit nil
### GetSystemStatus

`func (o *ApplianceGroupStatus) GetSystemStatus() ApplianceSystemStatusRelationship`

GetSystemStatus returns the SystemStatus field if non-nil, zero value otherwise.

### GetSystemStatusOk

`func (o *ApplianceGroupStatus) GetSystemStatusOk() (*ApplianceSystemStatusRelationship, bool)`

GetSystemStatusOk returns a tuple with the SystemStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemStatus

`func (o *ApplianceGroupStatus) SetSystemStatus(v ApplianceSystemStatusRelationship)`

SetSystemStatus sets SystemStatus field to given value.

### HasSystemStatus

`func (o *ApplianceGroupStatus) HasSystemStatus() bool`

HasSystemStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


