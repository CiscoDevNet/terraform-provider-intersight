# ApplianceSetupInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BuildType** | Pointer to **string** | Build type of the Intersight Appliance setup (e.g. release or debug). | [optional] [readonly] 
**Capabilities** | Pointer to [**[]ApplianceKeyValuePair**](appliance.KeyValuePair.md) |  | [optional] 
**CloudUrl** | Pointer to **string** | URL of the Intersight to which this Intersight Appliance is connected to. | [optional] [readonly] 
**DeploymentMode** | Pointer to **string** | Indicates where Intersight Appliance is installed in air-gapped or connected mode. In connected mode, Intersight Appliance is claimed by Intesight SaaS. In air-gapped mode, Intersight Appliance does not connect to any Cisco services. * &#x60;Connected&#x60; - In connected mode, Intersight Appliance connects to Intersight SaaS and other cisco.com services. * &#x60;Private&#x60; - In private mode, Intersight Appliance does not connect to Intersight SaaS or any cisco.com services. | [optional] [readonly] [default to "Connected"]
**EndTime** | Pointer to [**time.Time**](time.Time.md) | End date of the Intersight Appliance&#39;s initial setup. | [optional] [readonly] 
**SetupStates** | Pointer to **[]string** |  | [optional] 
**StartTime** | Pointer to [**time.Time**](time.Time.md) | Start date of the Intersight Appliance&#39;s initial setup. | [optional] [readonly] 
**Account** | Pointer to [**IamAccountRelationship**](iam.Account.Relationship.md) |  | [optional] 

## Methods

### NewApplianceSetupInfo

`func NewApplianceSetupInfo() *ApplianceSetupInfo`

NewApplianceSetupInfo instantiates a new ApplianceSetupInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplianceSetupInfoWithDefaults

`func NewApplianceSetupInfoWithDefaults() *ApplianceSetupInfo`

NewApplianceSetupInfoWithDefaults instantiates a new ApplianceSetupInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBuildType

`func (o *ApplianceSetupInfo) GetBuildType() string`

GetBuildType returns the BuildType field if non-nil, zero value otherwise.

### GetBuildTypeOk

`func (o *ApplianceSetupInfo) GetBuildTypeOk() (*string, bool)`

GetBuildTypeOk returns a tuple with the BuildType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildType

`func (o *ApplianceSetupInfo) SetBuildType(v string)`

SetBuildType sets BuildType field to given value.

### HasBuildType

`func (o *ApplianceSetupInfo) HasBuildType() bool`

HasBuildType returns a boolean if a field has been set.

### GetCapabilities

`func (o *ApplianceSetupInfo) GetCapabilities() []ApplianceKeyValuePair`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *ApplianceSetupInfo) GetCapabilitiesOk() (*[]ApplianceKeyValuePair, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *ApplianceSetupInfo) SetCapabilities(v []ApplianceKeyValuePair)`

SetCapabilities sets Capabilities field to given value.

### HasCapabilities

`func (o *ApplianceSetupInfo) HasCapabilities() bool`

HasCapabilities returns a boolean if a field has been set.

### GetCloudUrl

`func (o *ApplianceSetupInfo) GetCloudUrl() string`

GetCloudUrl returns the CloudUrl field if non-nil, zero value otherwise.

### GetCloudUrlOk

`func (o *ApplianceSetupInfo) GetCloudUrlOk() (*string, bool)`

GetCloudUrlOk returns a tuple with the CloudUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudUrl

`func (o *ApplianceSetupInfo) SetCloudUrl(v string)`

SetCloudUrl sets CloudUrl field to given value.

### HasCloudUrl

`func (o *ApplianceSetupInfo) HasCloudUrl() bool`

HasCloudUrl returns a boolean if a field has been set.

### GetDeploymentMode

`func (o *ApplianceSetupInfo) GetDeploymentMode() string`

GetDeploymentMode returns the DeploymentMode field if non-nil, zero value otherwise.

### GetDeploymentModeOk

`func (o *ApplianceSetupInfo) GetDeploymentModeOk() (*string, bool)`

GetDeploymentModeOk returns a tuple with the DeploymentMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentMode

`func (o *ApplianceSetupInfo) SetDeploymentMode(v string)`

SetDeploymentMode sets DeploymentMode field to given value.

### HasDeploymentMode

`func (o *ApplianceSetupInfo) HasDeploymentMode() bool`

HasDeploymentMode returns a boolean if a field has been set.

### GetEndTime

`func (o *ApplianceSetupInfo) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *ApplianceSetupInfo) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *ApplianceSetupInfo) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *ApplianceSetupInfo) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetSetupStates

`func (o *ApplianceSetupInfo) GetSetupStates() []string`

GetSetupStates returns the SetupStates field if non-nil, zero value otherwise.

### GetSetupStatesOk

`func (o *ApplianceSetupInfo) GetSetupStatesOk() (*[]string, bool)`

GetSetupStatesOk returns a tuple with the SetupStates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetupStates

`func (o *ApplianceSetupInfo) SetSetupStates(v []string)`

SetSetupStates sets SetupStates field to given value.

### HasSetupStates

`func (o *ApplianceSetupInfo) HasSetupStates() bool`

HasSetupStates returns a boolean if a field has been set.

### GetStartTime

`func (o *ApplianceSetupInfo) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *ApplianceSetupInfo) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *ApplianceSetupInfo) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *ApplianceSetupInfo) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetAccount

`func (o *ApplianceSetupInfo) GetAccount() IamAccountRelationship`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *ApplianceSetupInfo) GetAccountOk() (*IamAccountRelationship, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *ApplianceSetupInfo) SetAccount(v IamAccountRelationship)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *ApplianceSetupInfo) HasAccount() bool`

HasAccount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


