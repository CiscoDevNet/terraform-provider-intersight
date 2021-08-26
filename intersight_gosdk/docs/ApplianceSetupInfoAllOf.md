# ApplianceSetupInfoAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "appliance.SetupInfo"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "appliance.SetupInfo"]
**BuildType** | Pointer to **string** | Build type of the Intersight Appliance setup (e.g. release or debug). | [optional] [readonly] 
**Capabilities** | Pointer to [**[]ApplianceKeyValuePair**](ApplianceKeyValuePair.md) |  | [optional] 
**CloudUrl** | Pointer to **string** | URL of the Intersight to which this Intersight Appliance is connected to. | [optional] [readonly] 
**DeploymentMode** | Pointer to **string** | Indicates where Intersight Appliance is installed in air-gapped or connected mode. In connected mode, Intersight Appliance is claimed by Intesight SaaS. In air-gapped mode, Intersight Appliance does not connect to any Cisco services. * &#x60;Connected&#x60; - In connected mode, Intersight Appliance connects to Intersight SaaS and other cisco.com services. * &#x60;Private&#x60; - In private mode, Intersight Appliance does not connect to Intersight SaaS or any cisco.com services. | [optional] [readonly] [default to "Connected"]
**EndTime** | Pointer to **time.Time** | End date of the Intersight Appliance&#39;s initial setup. | [optional] [readonly] 
**SetupStates** | Pointer to **[]string** |  | [optional] 
**StartTime** | Pointer to **time.Time** | Start date of the Intersight Appliance&#39;s initial setup. | [optional] [readonly] 
**Account** | Pointer to [**IamAccountRelationship**](IamAccountRelationship.md) |  | [optional] 

## Methods

### NewApplianceSetupInfoAllOf

`func NewApplianceSetupInfoAllOf(classId string, objectType string, ) *ApplianceSetupInfoAllOf`

NewApplianceSetupInfoAllOf instantiates a new ApplianceSetupInfoAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplianceSetupInfoAllOfWithDefaults

`func NewApplianceSetupInfoAllOfWithDefaults() *ApplianceSetupInfoAllOf`

NewApplianceSetupInfoAllOfWithDefaults instantiates a new ApplianceSetupInfoAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ApplianceSetupInfoAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ApplianceSetupInfoAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ApplianceSetupInfoAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ApplianceSetupInfoAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ApplianceSetupInfoAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ApplianceSetupInfoAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetBuildType

`func (o *ApplianceSetupInfoAllOf) GetBuildType() string`

GetBuildType returns the BuildType field if non-nil, zero value otherwise.

### GetBuildTypeOk

`func (o *ApplianceSetupInfoAllOf) GetBuildTypeOk() (*string, bool)`

GetBuildTypeOk returns a tuple with the BuildType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildType

`func (o *ApplianceSetupInfoAllOf) SetBuildType(v string)`

SetBuildType sets BuildType field to given value.

### HasBuildType

`func (o *ApplianceSetupInfoAllOf) HasBuildType() bool`

HasBuildType returns a boolean if a field has been set.

### GetCapabilities

`func (o *ApplianceSetupInfoAllOf) GetCapabilities() []ApplianceKeyValuePair`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *ApplianceSetupInfoAllOf) GetCapabilitiesOk() (*[]ApplianceKeyValuePair, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *ApplianceSetupInfoAllOf) SetCapabilities(v []ApplianceKeyValuePair)`

SetCapabilities sets Capabilities field to given value.

### HasCapabilities

`func (o *ApplianceSetupInfoAllOf) HasCapabilities() bool`

HasCapabilities returns a boolean if a field has been set.

### SetCapabilitiesNil

`func (o *ApplianceSetupInfoAllOf) SetCapabilitiesNil(b bool)`

 SetCapabilitiesNil sets the value for Capabilities to be an explicit nil

### UnsetCapabilities
`func (o *ApplianceSetupInfoAllOf) UnsetCapabilities()`

UnsetCapabilities ensures that no value is present for Capabilities, not even an explicit nil
### GetCloudUrl

`func (o *ApplianceSetupInfoAllOf) GetCloudUrl() string`

GetCloudUrl returns the CloudUrl field if non-nil, zero value otherwise.

### GetCloudUrlOk

`func (o *ApplianceSetupInfoAllOf) GetCloudUrlOk() (*string, bool)`

GetCloudUrlOk returns a tuple with the CloudUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudUrl

`func (o *ApplianceSetupInfoAllOf) SetCloudUrl(v string)`

SetCloudUrl sets CloudUrl field to given value.

### HasCloudUrl

`func (o *ApplianceSetupInfoAllOf) HasCloudUrl() bool`

HasCloudUrl returns a boolean if a field has been set.

### GetDeploymentMode

`func (o *ApplianceSetupInfoAllOf) GetDeploymentMode() string`

GetDeploymentMode returns the DeploymentMode field if non-nil, zero value otherwise.

### GetDeploymentModeOk

`func (o *ApplianceSetupInfoAllOf) GetDeploymentModeOk() (*string, bool)`

GetDeploymentModeOk returns a tuple with the DeploymentMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentMode

`func (o *ApplianceSetupInfoAllOf) SetDeploymentMode(v string)`

SetDeploymentMode sets DeploymentMode field to given value.

### HasDeploymentMode

`func (o *ApplianceSetupInfoAllOf) HasDeploymentMode() bool`

HasDeploymentMode returns a boolean if a field has been set.

### GetEndTime

`func (o *ApplianceSetupInfoAllOf) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *ApplianceSetupInfoAllOf) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *ApplianceSetupInfoAllOf) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *ApplianceSetupInfoAllOf) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetSetupStates

`func (o *ApplianceSetupInfoAllOf) GetSetupStates() []string`

GetSetupStates returns the SetupStates field if non-nil, zero value otherwise.

### GetSetupStatesOk

`func (o *ApplianceSetupInfoAllOf) GetSetupStatesOk() (*[]string, bool)`

GetSetupStatesOk returns a tuple with the SetupStates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetupStates

`func (o *ApplianceSetupInfoAllOf) SetSetupStates(v []string)`

SetSetupStates sets SetupStates field to given value.

### HasSetupStates

`func (o *ApplianceSetupInfoAllOf) HasSetupStates() bool`

HasSetupStates returns a boolean if a field has been set.

### SetSetupStatesNil

`func (o *ApplianceSetupInfoAllOf) SetSetupStatesNil(b bool)`

 SetSetupStatesNil sets the value for SetupStates to be an explicit nil

### UnsetSetupStates
`func (o *ApplianceSetupInfoAllOf) UnsetSetupStates()`

UnsetSetupStates ensures that no value is present for SetupStates, not even an explicit nil
### GetStartTime

`func (o *ApplianceSetupInfoAllOf) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *ApplianceSetupInfoAllOf) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *ApplianceSetupInfoAllOf) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *ApplianceSetupInfoAllOf) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetAccount

`func (o *ApplianceSetupInfoAllOf) GetAccount() IamAccountRelationship`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *ApplianceSetupInfoAllOf) GetAccountOk() (*IamAccountRelationship, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *ApplianceSetupInfoAllOf) SetAccount(v IamAccountRelationship)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *ApplianceSetupInfoAllOf) HasAccount() bool`

HasAccount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


