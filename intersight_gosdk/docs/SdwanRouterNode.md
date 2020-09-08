# SdwanRouterNode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceTemplate** | Pointer to **string** | Name of the Cisco vManage device template that the current device should be attached to. A device template consists of many feature templates that contain SD-WAN vEdge router configuration. | [optional] 
**Name** | Pointer to **string** | Name of the router node object. | [optional] 
**NetworkConfiguration** | Pointer to [**[]SdwanNetworkConfigurationType**](sdwan.NetworkConfigurationType.md) |  | [optional] 
**TemplateInputs** | Pointer to [**[]SdwanTemplateInputsType**](sdwan.TemplateInputsType.md) |  | [optional] 
**Uuid** | Pointer to **string** | Uniquely identifies the router by its chassis number. | [optional] 
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](organization.Organization.Relationship.md) |  | [optional] 
**Profile** | Pointer to [**SdwanProfileRelationship**](sdwan.Profile.Relationship.md) |  | [optional] 
**ServerNode** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 

## Methods

### NewSdwanRouterNode

`func NewSdwanRouterNode() *SdwanRouterNode`

NewSdwanRouterNode instantiates a new SdwanRouterNode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSdwanRouterNodeWithDefaults

`func NewSdwanRouterNodeWithDefaults() *SdwanRouterNode`

NewSdwanRouterNodeWithDefaults instantiates a new SdwanRouterNode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceTemplate

`func (o *SdwanRouterNode) GetDeviceTemplate() string`

GetDeviceTemplate returns the DeviceTemplate field if non-nil, zero value otherwise.

### GetDeviceTemplateOk

`func (o *SdwanRouterNode) GetDeviceTemplateOk() (*string, bool)`

GetDeviceTemplateOk returns a tuple with the DeviceTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceTemplate

`func (o *SdwanRouterNode) SetDeviceTemplate(v string)`

SetDeviceTemplate sets DeviceTemplate field to given value.

### HasDeviceTemplate

`func (o *SdwanRouterNode) HasDeviceTemplate() bool`

HasDeviceTemplate returns a boolean if a field has been set.

### GetName

`func (o *SdwanRouterNode) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SdwanRouterNode) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SdwanRouterNode) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SdwanRouterNode) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNetworkConfiguration

`func (o *SdwanRouterNode) GetNetworkConfiguration() []SdwanNetworkConfigurationType`

GetNetworkConfiguration returns the NetworkConfiguration field if non-nil, zero value otherwise.

### GetNetworkConfigurationOk

`func (o *SdwanRouterNode) GetNetworkConfigurationOk() (*[]SdwanNetworkConfigurationType, bool)`

GetNetworkConfigurationOk returns a tuple with the NetworkConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkConfiguration

`func (o *SdwanRouterNode) SetNetworkConfiguration(v []SdwanNetworkConfigurationType)`

SetNetworkConfiguration sets NetworkConfiguration field to given value.

### HasNetworkConfiguration

`func (o *SdwanRouterNode) HasNetworkConfiguration() bool`

HasNetworkConfiguration returns a boolean if a field has been set.

### GetTemplateInputs

`func (o *SdwanRouterNode) GetTemplateInputs() []SdwanTemplateInputsType`

GetTemplateInputs returns the TemplateInputs field if non-nil, zero value otherwise.

### GetTemplateInputsOk

`func (o *SdwanRouterNode) GetTemplateInputsOk() (*[]SdwanTemplateInputsType, bool)`

GetTemplateInputsOk returns a tuple with the TemplateInputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateInputs

`func (o *SdwanRouterNode) SetTemplateInputs(v []SdwanTemplateInputsType)`

SetTemplateInputs sets TemplateInputs field to given value.

### HasTemplateInputs

`func (o *SdwanRouterNode) HasTemplateInputs() bool`

HasTemplateInputs returns a boolean if a field has been set.

### GetUuid

`func (o *SdwanRouterNode) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *SdwanRouterNode) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *SdwanRouterNode) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *SdwanRouterNode) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetOrganization

`func (o *SdwanRouterNode) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *SdwanRouterNode) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *SdwanRouterNode) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *SdwanRouterNode) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetProfile

`func (o *SdwanRouterNode) GetProfile() SdwanProfileRelationship`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *SdwanRouterNode) GetProfileOk() (*SdwanProfileRelationship, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *SdwanRouterNode) SetProfile(v SdwanProfileRelationship)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *SdwanRouterNode) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### GetServerNode

`func (o *SdwanRouterNode) GetServerNode() AssetDeviceRegistrationRelationship`

GetServerNode returns the ServerNode field if non-nil, zero value otherwise.

### GetServerNodeOk

`func (o *SdwanRouterNode) GetServerNodeOk() (*AssetDeviceRegistrationRelationship, bool)`

GetServerNodeOk returns a tuple with the ServerNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerNode

`func (o *SdwanRouterNode) SetServerNode(v AssetDeviceRegistrationRelationship)`

SetServerNode sets ServerNode field to given value.

### HasServerNode

`func (o *SdwanRouterNode) HasServerNode() bool`

HasServerNode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


