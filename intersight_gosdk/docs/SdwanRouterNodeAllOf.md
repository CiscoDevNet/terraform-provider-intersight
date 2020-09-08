# SdwanRouterNodeAllOf

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

### NewSdwanRouterNodeAllOf

`func NewSdwanRouterNodeAllOf() *SdwanRouterNodeAllOf`

NewSdwanRouterNodeAllOf instantiates a new SdwanRouterNodeAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSdwanRouterNodeAllOfWithDefaults

`func NewSdwanRouterNodeAllOfWithDefaults() *SdwanRouterNodeAllOf`

NewSdwanRouterNodeAllOfWithDefaults instantiates a new SdwanRouterNodeAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceTemplate

`func (o *SdwanRouterNodeAllOf) GetDeviceTemplate() string`

GetDeviceTemplate returns the DeviceTemplate field if non-nil, zero value otherwise.

### GetDeviceTemplateOk

`func (o *SdwanRouterNodeAllOf) GetDeviceTemplateOk() (*string, bool)`

GetDeviceTemplateOk returns a tuple with the DeviceTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceTemplate

`func (o *SdwanRouterNodeAllOf) SetDeviceTemplate(v string)`

SetDeviceTemplate sets DeviceTemplate field to given value.

### HasDeviceTemplate

`func (o *SdwanRouterNodeAllOf) HasDeviceTemplate() bool`

HasDeviceTemplate returns a boolean if a field has been set.

### GetName

`func (o *SdwanRouterNodeAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SdwanRouterNodeAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SdwanRouterNodeAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SdwanRouterNodeAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNetworkConfiguration

`func (o *SdwanRouterNodeAllOf) GetNetworkConfiguration() []SdwanNetworkConfigurationType`

GetNetworkConfiguration returns the NetworkConfiguration field if non-nil, zero value otherwise.

### GetNetworkConfigurationOk

`func (o *SdwanRouterNodeAllOf) GetNetworkConfigurationOk() (*[]SdwanNetworkConfigurationType, bool)`

GetNetworkConfigurationOk returns a tuple with the NetworkConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkConfiguration

`func (o *SdwanRouterNodeAllOf) SetNetworkConfiguration(v []SdwanNetworkConfigurationType)`

SetNetworkConfiguration sets NetworkConfiguration field to given value.

### HasNetworkConfiguration

`func (o *SdwanRouterNodeAllOf) HasNetworkConfiguration() bool`

HasNetworkConfiguration returns a boolean if a field has been set.

### GetTemplateInputs

`func (o *SdwanRouterNodeAllOf) GetTemplateInputs() []SdwanTemplateInputsType`

GetTemplateInputs returns the TemplateInputs field if non-nil, zero value otherwise.

### GetTemplateInputsOk

`func (o *SdwanRouterNodeAllOf) GetTemplateInputsOk() (*[]SdwanTemplateInputsType, bool)`

GetTemplateInputsOk returns a tuple with the TemplateInputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateInputs

`func (o *SdwanRouterNodeAllOf) SetTemplateInputs(v []SdwanTemplateInputsType)`

SetTemplateInputs sets TemplateInputs field to given value.

### HasTemplateInputs

`func (o *SdwanRouterNodeAllOf) HasTemplateInputs() bool`

HasTemplateInputs returns a boolean if a field has been set.

### GetUuid

`func (o *SdwanRouterNodeAllOf) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *SdwanRouterNodeAllOf) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *SdwanRouterNodeAllOf) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *SdwanRouterNodeAllOf) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetOrganization

`func (o *SdwanRouterNodeAllOf) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *SdwanRouterNodeAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *SdwanRouterNodeAllOf) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *SdwanRouterNodeAllOf) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetProfile

`func (o *SdwanRouterNodeAllOf) GetProfile() SdwanProfileRelationship`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *SdwanRouterNodeAllOf) GetProfileOk() (*SdwanProfileRelationship, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *SdwanRouterNodeAllOf) SetProfile(v SdwanProfileRelationship)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *SdwanRouterNodeAllOf) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### GetServerNode

`func (o *SdwanRouterNodeAllOf) GetServerNode() AssetDeviceRegistrationRelationship`

GetServerNode returns the ServerNode field if non-nil, zero value otherwise.

### GetServerNodeOk

`func (o *SdwanRouterNodeAllOf) GetServerNodeOk() (*AssetDeviceRegistrationRelationship, bool)`

GetServerNodeOk returns a tuple with the ServerNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerNode

`func (o *SdwanRouterNodeAllOf) SetServerNode(v AssetDeviceRegistrationRelationship)`

SetServerNode sets ServerNode field to given value.

### HasServerNode

`func (o *SdwanRouterNodeAllOf) HasServerNode() bool`

HasServerNode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


