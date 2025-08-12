# FabricPortOperation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "fabric.PortOperation"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "fabric.PortOperation"]
**AdminAction** | Pointer to **string** | An operation that has to be perfomed on the switch or IOM port. Default value is None which means there will be no implicit port operation triggered. * &#x60;None&#x60; - No admin triggered action. * &#x60;ResetServerPortConfiguration&#x60; - Admin triggered operation to reset the server port to its original configuration. * &#x60;SetUserLabel&#x60; - Admin triggered operation to set the user label on the port. | [optional] [default to "None"]
**AdminState** | Pointer to **string** | Admin configured state to disable the port. * &#x60;Enabled&#x60; - Admin configured Enabled State. * &#x60;Disabled&#x60; - Admin configured Disabled State. | [optional] [default to "Enabled"]
**ConfigState** | Pointer to **string** | The configured state of these settings in the target chassis. The value is any one of Applied, Applying, Failed. Applied - This state denotes that the admin state changes are applied successfully in the target FI domain. Applying - This state denotes that the admin state changes are being applied in the target FI domain. Failed - This state denotes that the admin state changes could not be applied in the target FI domain. * &#x60;None&#x60; - Nil value when no action has been triggered by the user. * &#x60;Applied&#x60; - User configured settings are in applied state. * &#x60;Applying&#x60; - User settings are being applied on the target server. * &#x60;Failed&#x60; - User configured settings could not be applied. | [optional] [default to "None"]
**FexId** | Pointer to **int64** | FEX/IOM identifier to denote its Host ports in the format - FexId/SlotId/PortId. | [optional] 
**SwitchName** | Pointer to **string** | Name of the switch on which the port operation is performed. | [optional] [readonly] 
**UserLabel** | Pointer to **string** | The user defined label assigned to the a Port. | [optional] 
**NetworkElement** | Pointer to [**NullableNetworkElementRelationship**](NetworkElementRelationship.md) |  | [optional] 

## Methods

### NewFabricPortOperation

`func NewFabricPortOperation(classId string, objectType string, ) *FabricPortOperation`

NewFabricPortOperation instantiates a new FabricPortOperation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFabricPortOperationWithDefaults

`func NewFabricPortOperationWithDefaults() *FabricPortOperation`

NewFabricPortOperationWithDefaults instantiates a new FabricPortOperation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *FabricPortOperation) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *FabricPortOperation) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *FabricPortOperation) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *FabricPortOperation) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *FabricPortOperation) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *FabricPortOperation) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAdminAction

`func (o *FabricPortOperation) GetAdminAction() string`

GetAdminAction returns the AdminAction field if non-nil, zero value otherwise.

### GetAdminActionOk

`func (o *FabricPortOperation) GetAdminActionOk() (*string, bool)`

GetAdminActionOk returns a tuple with the AdminAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminAction

`func (o *FabricPortOperation) SetAdminAction(v string)`

SetAdminAction sets AdminAction field to given value.

### HasAdminAction

`func (o *FabricPortOperation) HasAdminAction() bool`

HasAdminAction returns a boolean if a field has been set.

### GetAdminState

`func (o *FabricPortOperation) GetAdminState() string`

GetAdminState returns the AdminState field if non-nil, zero value otherwise.

### GetAdminStateOk

`func (o *FabricPortOperation) GetAdminStateOk() (*string, bool)`

GetAdminStateOk returns a tuple with the AdminState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminState

`func (o *FabricPortOperation) SetAdminState(v string)`

SetAdminState sets AdminState field to given value.

### HasAdminState

`func (o *FabricPortOperation) HasAdminState() bool`

HasAdminState returns a boolean if a field has been set.

### GetConfigState

`func (o *FabricPortOperation) GetConfigState() string`

GetConfigState returns the ConfigState field if non-nil, zero value otherwise.

### GetConfigStateOk

`func (o *FabricPortOperation) GetConfigStateOk() (*string, bool)`

GetConfigStateOk returns a tuple with the ConfigState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigState

`func (o *FabricPortOperation) SetConfigState(v string)`

SetConfigState sets ConfigState field to given value.

### HasConfigState

`func (o *FabricPortOperation) HasConfigState() bool`

HasConfigState returns a boolean if a field has been set.

### GetFexId

`func (o *FabricPortOperation) GetFexId() int64`

GetFexId returns the FexId field if non-nil, zero value otherwise.

### GetFexIdOk

`func (o *FabricPortOperation) GetFexIdOk() (*int64, bool)`

GetFexIdOk returns a tuple with the FexId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFexId

`func (o *FabricPortOperation) SetFexId(v int64)`

SetFexId sets FexId field to given value.

### HasFexId

`func (o *FabricPortOperation) HasFexId() bool`

HasFexId returns a boolean if a field has been set.

### GetSwitchName

`func (o *FabricPortOperation) GetSwitchName() string`

GetSwitchName returns the SwitchName field if non-nil, zero value otherwise.

### GetSwitchNameOk

`func (o *FabricPortOperation) GetSwitchNameOk() (*string, bool)`

GetSwitchNameOk returns a tuple with the SwitchName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchName

`func (o *FabricPortOperation) SetSwitchName(v string)`

SetSwitchName sets SwitchName field to given value.

### HasSwitchName

`func (o *FabricPortOperation) HasSwitchName() bool`

HasSwitchName returns a boolean if a field has been set.

### GetUserLabel

`func (o *FabricPortOperation) GetUserLabel() string`

GetUserLabel returns the UserLabel field if non-nil, zero value otherwise.

### GetUserLabelOk

`func (o *FabricPortOperation) GetUserLabelOk() (*string, bool)`

GetUserLabelOk returns a tuple with the UserLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserLabel

`func (o *FabricPortOperation) SetUserLabel(v string)`

SetUserLabel sets UserLabel field to given value.

### HasUserLabel

`func (o *FabricPortOperation) HasUserLabel() bool`

HasUserLabel returns a boolean if a field has been set.

### GetNetworkElement

`func (o *FabricPortOperation) GetNetworkElement() NetworkElementRelationship`

GetNetworkElement returns the NetworkElement field if non-nil, zero value otherwise.

### GetNetworkElementOk

`func (o *FabricPortOperation) GetNetworkElementOk() (*NetworkElementRelationship, bool)`

GetNetworkElementOk returns a tuple with the NetworkElement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkElement

`func (o *FabricPortOperation) SetNetworkElement(v NetworkElementRelationship)`

SetNetworkElement sets NetworkElement field to given value.

### HasNetworkElement

`func (o *FabricPortOperation) HasNetworkElement() bool`

HasNetworkElement returns a boolean if a field has been set.

### SetNetworkElementNil

`func (o *FabricPortOperation) SetNetworkElementNil(b bool)`

 SetNetworkElementNil sets the value for NetworkElement to be an explicit nil

### UnsetNetworkElement
`func (o *FabricPortOperation) UnsetNetworkElement()`

UnsetNetworkElement ensures that no value is present for NetworkElement, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


