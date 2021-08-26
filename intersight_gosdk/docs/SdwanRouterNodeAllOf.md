# SdwanRouterNodeAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "sdwan.RouterNode"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "sdwan.RouterNode"]
**DeviceTemplate** | Pointer to **string** | Name of the Cisco vManage device template that the current device should be attached to. A device template consists of many feature templates that contain SD-WAN vEdge router configuration. | [optional] 
**Name** | Pointer to **string** | Name of the router node object. | [optional] 
**NetworkConfiguration** | Pointer to [**[]SdwanNetworkConfigurationType**](SdwanNetworkConfigurationType.md) |  | [optional] 
**TemplateInputs** | Pointer to [**[]SdwanTemplateInputsType**](SdwanTemplateInputsType.md) |  | [optional] 
**Uuid** | Pointer to **string** | Uniquely identifies the router by its chassis number. | [optional] 
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](OrganizationOrganizationRelationship.md) |  | [optional] 
**Profile** | Pointer to [**SdwanProfileRelationship**](SdwanProfileRelationship.md) |  | [optional] 
**ServerNode** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewSdwanRouterNodeAllOf

`func NewSdwanRouterNodeAllOf(classId string, objectType string, ) *SdwanRouterNodeAllOf`

NewSdwanRouterNodeAllOf instantiates a new SdwanRouterNodeAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSdwanRouterNodeAllOfWithDefaults

`func NewSdwanRouterNodeAllOfWithDefaults() *SdwanRouterNodeAllOf`

NewSdwanRouterNodeAllOfWithDefaults instantiates a new SdwanRouterNodeAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *SdwanRouterNodeAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *SdwanRouterNodeAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *SdwanRouterNodeAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *SdwanRouterNodeAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *SdwanRouterNodeAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *SdwanRouterNodeAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


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

### SetNetworkConfigurationNil

`func (o *SdwanRouterNodeAllOf) SetNetworkConfigurationNil(b bool)`

 SetNetworkConfigurationNil sets the value for NetworkConfiguration to be an explicit nil

### UnsetNetworkConfiguration
`func (o *SdwanRouterNodeAllOf) UnsetNetworkConfiguration()`

UnsetNetworkConfiguration ensures that no value is present for NetworkConfiguration, not even an explicit nil
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

### SetTemplateInputsNil

`func (o *SdwanRouterNodeAllOf) SetTemplateInputsNil(b bool)`

 SetTemplateInputsNil sets the value for TemplateInputs to be an explicit nil

### UnsetTemplateInputs
`func (o *SdwanRouterNodeAllOf) UnsetTemplateInputs()`

UnsetTemplateInputs ensures that no value is present for TemplateInputs, not even an explicit nil
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


