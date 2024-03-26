# ApplianceGroupOpStatusAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "appliance.GroupOpStatus"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "appliance.GroupOpStatus"]
**Description** | Pointer to **string** | Description of the group. | [optional] [readonly] 
**GroupName** | Pointer to **string** | The name of group, which includes Identity Management, Device Connector Service, Core Service, Analytics, Infra and Appliance. | [optional] [readonly] 
**OverallStatus** | Pointer to **string** | The overall API status from this group&#39;s applications. | [optional] [readonly] 
**Apps** | Pointer to [**[]ApplianceAppOpStatusRelationship**](ApplianceAppOpStatusRelationship.md) | An array of relationships to applianceAppOpStatus resources. | [optional] [readonly] 
**SystemOpStatus** | Pointer to [**ApplianceSystemOpStatusRelationship**](ApplianceSystemOpStatusRelationship.md) |  | [optional] 

## Methods

### NewApplianceGroupOpStatusAllOf

`func NewApplianceGroupOpStatusAllOf(classId string, objectType string, ) *ApplianceGroupOpStatusAllOf`

NewApplianceGroupOpStatusAllOf instantiates a new ApplianceGroupOpStatusAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplianceGroupOpStatusAllOfWithDefaults

`func NewApplianceGroupOpStatusAllOfWithDefaults() *ApplianceGroupOpStatusAllOf`

NewApplianceGroupOpStatusAllOfWithDefaults instantiates a new ApplianceGroupOpStatusAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ApplianceGroupOpStatusAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ApplianceGroupOpStatusAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ApplianceGroupOpStatusAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ApplianceGroupOpStatusAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ApplianceGroupOpStatusAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ApplianceGroupOpStatusAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDescription

`func (o *ApplianceGroupOpStatusAllOf) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ApplianceGroupOpStatusAllOf) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ApplianceGroupOpStatusAllOf) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ApplianceGroupOpStatusAllOf) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetGroupName

`func (o *ApplianceGroupOpStatusAllOf) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *ApplianceGroupOpStatusAllOf) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *ApplianceGroupOpStatusAllOf) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.

### HasGroupName

`func (o *ApplianceGroupOpStatusAllOf) HasGroupName() bool`

HasGroupName returns a boolean if a field has been set.

### GetOverallStatus

`func (o *ApplianceGroupOpStatusAllOf) GetOverallStatus() string`

GetOverallStatus returns the OverallStatus field if non-nil, zero value otherwise.

### GetOverallStatusOk

`func (o *ApplianceGroupOpStatusAllOf) GetOverallStatusOk() (*string, bool)`

GetOverallStatusOk returns a tuple with the OverallStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverallStatus

`func (o *ApplianceGroupOpStatusAllOf) SetOverallStatus(v string)`

SetOverallStatus sets OverallStatus field to given value.

### HasOverallStatus

`func (o *ApplianceGroupOpStatusAllOf) HasOverallStatus() bool`

HasOverallStatus returns a boolean if a field has been set.

### GetApps

`func (o *ApplianceGroupOpStatusAllOf) GetApps() []ApplianceAppOpStatusRelationship`

GetApps returns the Apps field if non-nil, zero value otherwise.

### GetAppsOk

`func (o *ApplianceGroupOpStatusAllOf) GetAppsOk() (*[]ApplianceAppOpStatusRelationship, bool)`

GetAppsOk returns a tuple with the Apps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApps

`func (o *ApplianceGroupOpStatusAllOf) SetApps(v []ApplianceAppOpStatusRelationship)`

SetApps sets Apps field to given value.

### HasApps

`func (o *ApplianceGroupOpStatusAllOf) HasApps() bool`

HasApps returns a boolean if a field has been set.

### SetAppsNil

`func (o *ApplianceGroupOpStatusAllOf) SetAppsNil(b bool)`

 SetAppsNil sets the value for Apps to be an explicit nil

### UnsetApps
`func (o *ApplianceGroupOpStatusAllOf) UnsetApps()`

UnsetApps ensures that no value is present for Apps, not even an explicit nil
### GetSystemOpStatus

`func (o *ApplianceGroupOpStatusAllOf) GetSystemOpStatus() ApplianceSystemOpStatusRelationship`

GetSystemOpStatus returns the SystemOpStatus field if non-nil, zero value otherwise.

### GetSystemOpStatusOk

`func (o *ApplianceGroupOpStatusAllOf) GetSystemOpStatusOk() (*ApplianceSystemOpStatusRelationship, bool)`

GetSystemOpStatusOk returns a tuple with the SystemOpStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemOpStatus

`func (o *ApplianceGroupOpStatusAllOf) SetSystemOpStatus(v ApplianceSystemOpStatusRelationship)`

SetSystemOpStatus sets SystemOpStatus field to given value.

### HasSystemOpStatus

`func (o *ApplianceGroupOpStatusAllOf) HasSystemOpStatus() bool`

HasSystemOpStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


