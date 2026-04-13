# VnicServiceEthIf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "vnic.ServiceEthIf"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "vnic.ServiceEthIf"]
**AdapterId** | Pointer to **string** | Adapter Information for server. | [optional] 
**Name** | Pointer to **string** | Name of the service virtual ethernet interface. | [optional] 
**OverriddenList** | Pointer to **[]string** |  | [optional] 
**Placement** | Pointer to [**NullableVnicPlacementSettings**](VnicPlacementSettings.md) |  | [optional] 
**TemplateActions** | Pointer to [**[]MotemplateActionEntry**](MotemplateActionEntry.md) |  | [optional] 
**TemplateSyncErrors** | Pointer to [**[]MotemplateSyncError**](MotemplateSyncError.md) |  | [optional] 
**TemplateSyncStatus** | Pointer to **string** | The sync status of the current MO wrt the attached Template MO. * &#x60;None&#x60; - The Enum value represents that the object is not attached to any template. * &#x60;OK&#x60; - The Enum value represents that the object values are in sync with attached template. * &#x60;Scheduled&#x60; - The Enum value represents that the object sync from attached template is scheduled from template. * &#x60;InProgress&#x60; - The Enum value represents that the object sync with the attached template is in progress. * &#x60;OutOfSync&#x60; - The Enum value represents that the object values are not in sync with attached template. | [optional] [readonly] [default to "None"]
**VifId** | Pointer to **int64** | The Vif Id should be same as the channel number of the vethernet created on switch in order to set up the data path. The property is applicable only for Fabric Interconnect attached servers where a vethernet is created on the switch for every vNIC. | [optional] [readonly] 
**Profile** | Pointer to [**NullablePolicyAbstractConfigProfileRelationship**](PolicyAbstractConfigProfileRelationship.md) |  | [optional] 

## Methods

### NewVnicServiceEthIf

`func NewVnicServiceEthIf(classId string, objectType string, ) *VnicServiceEthIf`

NewVnicServiceEthIf instantiates a new VnicServiceEthIf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVnicServiceEthIfWithDefaults

`func NewVnicServiceEthIfWithDefaults() *VnicServiceEthIf`

NewVnicServiceEthIfWithDefaults instantiates a new VnicServiceEthIf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *VnicServiceEthIf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *VnicServiceEthIf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *VnicServiceEthIf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *VnicServiceEthIf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VnicServiceEthIf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VnicServiceEthIf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAdapterId

`func (o *VnicServiceEthIf) GetAdapterId() string`

GetAdapterId returns the AdapterId field if non-nil, zero value otherwise.

### GetAdapterIdOk

`func (o *VnicServiceEthIf) GetAdapterIdOk() (*string, bool)`

GetAdapterIdOk returns a tuple with the AdapterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdapterId

`func (o *VnicServiceEthIf) SetAdapterId(v string)`

SetAdapterId sets AdapterId field to given value.

### HasAdapterId

`func (o *VnicServiceEthIf) HasAdapterId() bool`

HasAdapterId returns a boolean if a field has been set.

### GetName

`func (o *VnicServiceEthIf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VnicServiceEthIf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VnicServiceEthIf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VnicServiceEthIf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOverriddenList

`func (o *VnicServiceEthIf) GetOverriddenList() []string`

GetOverriddenList returns the OverriddenList field if non-nil, zero value otherwise.

### GetOverriddenListOk

`func (o *VnicServiceEthIf) GetOverriddenListOk() (*[]string, bool)`

GetOverriddenListOk returns a tuple with the OverriddenList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverriddenList

`func (o *VnicServiceEthIf) SetOverriddenList(v []string)`

SetOverriddenList sets OverriddenList field to given value.

### HasOverriddenList

`func (o *VnicServiceEthIf) HasOverriddenList() bool`

HasOverriddenList returns a boolean if a field has been set.

### SetOverriddenListNil

`func (o *VnicServiceEthIf) SetOverriddenListNil(b bool)`

 SetOverriddenListNil sets the value for OverriddenList to be an explicit nil

### UnsetOverriddenList
`func (o *VnicServiceEthIf) UnsetOverriddenList()`

UnsetOverriddenList ensures that no value is present for OverriddenList, not even an explicit nil
### GetPlacement

`func (o *VnicServiceEthIf) GetPlacement() VnicPlacementSettings`

GetPlacement returns the Placement field if non-nil, zero value otherwise.

### GetPlacementOk

`func (o *VnicServiceEthIf) GetPlacementOk() (*VnicPlacementSettings, bool)`

GetPlacementOk returns a tuple with the Placement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlacement

`func (o *VnicServiceEthIf) SetPlacement(v VnicPlacementSettings)`

SetPlacement sets Placement field to given value.

### HasPlacement

`func (o *VnicServiceEthIf) HasPlacement() bool`

HasPlacement returns a boolean if a field has been set.

### SetPlacementNil

`func (o *VnicServiceEthIf) SetPlacementNil(b bool)`

 SetPlacementNil sets the value for Placement to be an explicit nil

### UnsetPlacement
`func (o *VnicServiceEthIf) UnsetPlacement()`

UnsetPlacement ensures that no value is present for Placement, not even an explicit nil
### GetTemplateActions

`func (o *VnicServiceEthIf) GetTemplateActions() []MotemplateActionEntry`

GetTemplateActions returns the TemplateActions field if non-nil, zero value otherwise.

### GetTemplateActionsOk

`func (o *VnicServiceEthIf) GetTemplateActionsOk() (*[]MotemplateActionEntry, bool)`

GetTemplateActionsOk returns a tuple with the TemplateActions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateActions

`func (o *VnicServiceEthIf) SetTemplateActions(v []MotemplateActionEntry)`

SetTemplateActions sets TemplateActions field to given value.

### HasTemplateActions

`func (o *VnicServiceEthIf) HasTemplateActions() bool`

HasTemplateActions returns a boolean if a field has been set.

### SetTemplateActionsNil

`func (o *VnicServiceEthIf) SetTemplateActionsNil(b bool)`

 SetTemplateActionsNil sets the value for TemplateActions to be an explicit nil

### UnsetTemplateActions
`func (o *VnicServiceEthIf) UnsetTemplateActions()`

UnsetTemplateActions ensures that no value is present for TemplateActions, not even an explicit nil
### GetTemplateSyncErrors

`func (o *VnicServiceEthIf) GetTemplateSyncErrors() []MotemplateSyncError`

GetTemplateSyncErrors returns the TemplateSyncErrors field if non-nil, zero value otherwise.

### GetTemplateSyncErrorsOk

`func (o *VnicServiceEthIf) GetTemplateSyncErrorsOk() (*[]MotemplateSyncError, bool)`

GetTemplateSyncErrorsOk returns a tuple with the TemplateSyncErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateSyncErrors

`func (o *VnicServiceEthIf) SetTemplateSyncErrors(v []MotemplateSyncError)`

SetTemplateSyncErrors sets TemplateSyncErrors field to given value.

### HasTemplateSyncErrors

`func (o *VnicServiceEthIf) HasTemplateSyncErrors() bool`

HasTemplateSyncErrors returns a boolean if a field has been set.

### SetTemplateSyncErrorsNil

`func (o *VnicServiceEthIf) SetTemplateSyncErrorsNil(b bool)`

 SetTemplateSyncErrorsNil sets the value for TemplateSyncErrors to be an explicit nil

### UnsetTemplateSyncErrors
`func (o *VnicServiceEthIf) UnsetTemplateSyncErrors()`

UnsetTemplateSyncErrors ensures that no value is present for TemplateSyncErrors, not even an explicit nil
### GetTemplateSyncStatus

`func (o *VnicServiceEthIf) GetTemplateSyncStatus() string`

GetTemplateSyncStatus returns the TemplateSyncStatus field if non-nil, zero value otherwise.

### GetTemplateSyncStatusOk

`func (o *VnicServiceEthIf) GetTemplateSyncStatusOk() (*string, bool)`

GetTemplateSyncStatusOk returns a tuple with the TemplateSyncStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateSyncStatus

`func (o *VnicServiceEthIf) SetTemplateSyncStatus(v string)`

SetTemplateSyncStatus sets TemplateSyncStatus field to given value.

### HasTemplateSyncStatus

`func (o *VnicServiceEthIf) HasTemplateSyncStatus() bool`

HasTemplateSyncStatus returns a boolean if a field has been set.

### GetVifId

`func (o *VnicServiceEthIf) GetVifId() int64`

GetVifId returns the VifId field if non-nil, zero value otherwise.

### GetVifIdOk

`func (o *VnicServiceEthIf) GetVifIdOk() (*int64, bool)`

GetVifIdOk returns a tuple with the VifId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVifId

`func (o *VnicServiceEthIf) SetVifId(v int64)`

SetVifId sets VifId field to given value.

### HasVifId

`func (o *VnicServiceEthIf) HasVifId() bool`

HasVifId returns a boolean if a field has been set.

### GetProfile

`func (o *VnicServiceEthIf) GetProfile() PolicyAbstractConfigProfileRelationship`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *VnicServiceEthIf) GetProfileOk() (*PolicyAbstractConfigProfileRelationship, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *VnicServiceEthIf) SetProfile(v PolicyAbstractConfigProfileRelationship)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *VnicServiceEthIf) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### SetProfileNil

`func (o *VnicServiceEthIf) SetProfileNil(b bool)`

 SetProfileNil sets the value for Profile to be an explicit nil

### UnsetProfile
`func (o *VnicServiceEthIf) UnsetProfile()`

UnsetProfile ensures that no value is present for Profile, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


