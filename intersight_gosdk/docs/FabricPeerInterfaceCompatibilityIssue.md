# FabricPeerInterfaceCompatibilityIssue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "fabric.PeerInterfaceCompatibilityIssue"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "fabric.PeerInterfaceCompatibilityIssue"]
**DesiredValue** | Pointer to **string** | Value configured in Intersight intent. | [optional] [readonly] 
**IssueType** | Pointer to **string** | Type of compatibility issue detected. * &#x60;None&#x60; - No peer port compatibility issue was detected. * &#x60;Mode Mismatch&#x60; - The port mode on the Secure Router backplane in Intersight differs from the mode set in Catalyst SD-WAN Manager. * &#x60;Allowed VLANs Mismatch&#x60; - The allowed VLANs on the Secure Router backplane in Intersight differ from those in Catalyst SD-WAN Manager. VLAN 4047 is automatically added on the Secure Router and is included in the comparison. * &#x60;Native VLAN Mismatch&#x60; - The native VLAN on the Secure Router backplane in Intersight differs from the native VLAN in Catalyst SD-WAN Manager. * &#x60;Access VLAN Mismatch&#x60; - The access VLAN on the Secure Router backplane in Intersight differs from the access VLAN in Catalyst SD-WAN Manager. * &#x60;MTU Mismatch&#x60; - The MTU on the switch in Intersight differs from the MTU on the Secure Router in Catalyst SD-WAN Manager. | [optional] [readonly] [default to "None"]
**LocalInterface** | Pointer to [**MoMoRef**](MoMoRef.md) |  | [optional] 
**PeerDevice** | Pointer to [**MoMoRef**](MoMoRef.md) |  | [optional] 
**PeerInterface** | Pointer to [**MoMoRef**](MoMoRef.md) |  | [optional] 
**PeerInterfaceName** | Pointer to **string** | User-friendly name of the peer interface corresponding to the mismatch. | [optional] [readonly] 
**PeerValue** | Pointer to **string** | Value reported by Catalyst SD-WAN Manager. | [optional] [readonly] 
**PeerValueModTime** | Pointer to **string** | Time when the peer value used for this compatibility check was captured. | [optional] [readonly] 
**Source** | Pointer to **string** | Source of the local endpoint where the mismatch was detected, including the slot identifier when applicable. | [optional] [readonly] 
**SourcePolicy** | Pointer to [**MoMoRef**](MoMoRef.md) |  | [optional] 

## Methods

### NewFabricPeerInterfaceCompatibilityIssue

`func NewFabricPeerInterfaceCompatibilityIssue(classId string, objectType string, ) *FabricPeerInterfaceCompatibilityIssue`

NewFabricPeerInterfaceCompatibilityIssue instantiates a new FabricPeerInterfaceCompatibilityIssue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFabricPeerInterfaceCompatibilityIssueWithDefaults

`func NewFabricPeerInterfaceCompatibilityIssueWithDefaults() *FabricPeerInterfaceCompatibilityIssue`

NewFabricPeerInterfaceCompatibilityIssueWithDefaults instantiates a new FabricPeerInterfaceCompatibilityIssue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *FabricPeerInterfaceCompatibilityIssue) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *FabricPeerInterfaceCompatibilityIssue) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *FabricPeerInterfaceCompatibilityIssue) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *FabricPeerInterfaceCompatibilityIssue) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *FabricPeerInterfaceCompatibilityIssue) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *FabricPeerInterfaceCompatibilityIssue) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDesiredValue

`func (o *FabricPeerInterfaceCompatibilityIssue) GetDesiredValue() string`

GetDesiredValue returns the DesiredValue field if non-nil, zero value otherwise.

### GetDesiredValueOk

`func (o *FabricPeerInterfaceCompatibilityIssue) GetDesiredValueOk() (*string, bool)`

GetDesiredValueOk returns a tuple with the DesiredValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesiredValue

`func (o *FabricPeerInterfaceCompatibilityIssue) SetDesiredValue(v string)`

SetDesiredValue sets DesiredValue field to given value.

### HasDesiredValue

`func (o *FabricPeerInterfaceCompatibilityIssue) HasDesiredValue() bool`

HasDesiredValue returns a boolean if a field has been set.

### GetIssueType

`func (o *FabricPeerInterfaceCompatibilityIssue) GetIssueType() string`

GetIssueType returns the IssueType field if non-nil, zero value otherwise.

### GetIssueTypeOk

`func (o *FabricPeerInterfaceCompatibilityIssue) GetIssueTypeOk() (*string, bool)`

GetIssueTypeOk returns a tuple with the IssueType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueType

`func (o *FabricPeerInterfaceCompatibilityIssue) SetIssueType(v string)`

SetIssueType sets IssueType field to given value.

### HasIssueType

`func (o *FabricPeerInterfaceCompatibilityIssue) HasIssueType() bool`

HasIssueType returns a boolean if a field has been set.

### GetLocalInterface

`func (o *FabricPeerInterfaceCompatibilityIssue) GetLocalInterface() MoMoRef`

GetLocalInterface returns the LocalInterface field if non-nil, zero value otherwise.

### GetLocalInterfaceOk

`func (o *FabricPeerInterfaceCompatibilityIssue) GetLocalInterfaceOk() (*MoMoRef, bool)`

GetLocalInterfaceOk returns a tuple with the LocalInterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalInterface

`func (o *FabricPeerInterfaceCompatibilityIssue) SetLocalInterface(v MoMoRef)`

SetLocalInterface sets LocalInterface field to given value.

### HasLocalInterface

`func (o *FabricPeerInterfaceCompatibilityIssue) HasLocalInterface() bool`

HasLocalInterface returns a boolean if a field has been set.

### GetPeerDevice

`func (o *FabricPeerInterfaceCompatibilityIssue) GetPeerDevice() MoMoRef`

GetPeerDevice returns the PeerDevice field if non-nil, zero value otherwise.

### GetPeerDeviceOk

`func (o *FabricPeerInterfaceCompatibilityIssue) GetPeerDeviceOk() (*MoMoRef, bool)`

GetPeerDeviceOk returns a tuple with the PeerDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerDevice

`func (o *FabricPeerInterfaceCompatibilityIssue) SetPeerDevice(v MoMoRef)`

SetPeerDevice sets PeerDevice field to given value.

### HasPeerDevice

`func (o *FabricPeerInterfaceCompatibilityIssue) HasPeerDevice() bool`

HasPeerDevice returns a boolean if a field has been set.

### GetPeerInterface

`func (o *FabricPeerInterfaceCompatibilityIssue) GetPeerInterface() MoMoRef`

GetPeerInterface returns the PeerInterface field if non-nil, zero value otherwise.

### GetPeerInterfaceOk

`func (o *FabricPeerInterfaceCompatibilityIssue) GetPeerInterfaceOk() (*MoMoRef, bool)`

GetPeerInterfaceOk returns a tuple with the PeerInterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerInterface

`func (o *FabricPeerInterfaceCompatibilityIssue) SetPeerInterface(v MoMoRef)`

SetPeerInterface sets PeerInterface field to given value.

### HasPeerInterface

`func (o *FabricPeerInterfaceCompatibilityIssue) HasPeerInterface() bool`

HasPeerInterface returns a boolean if a field has been set.

### GetPeerInterfaceName

`func (o *FabricPeerInterfaceCompatibilityIssue) GetPeerInterfaceName() string`

GetPeerInterfaceName returns the PeerInterfaceName field if non-nil, zero value otherwise.

### GetPeerInterfaceNameOk

`func (o *FabricPeerInterfaceCompatibilityIssue) GetPeerInterfaceNameOk() (*string, bool)`

GetPeerInterfaceNameOk returns a tuple with the PeerInterfaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerInterfaceName

`func (o *FabricPeerInterfaceCompatibilityIssue) SetPeerInterfaceName(v string)`

SetPeerInterfaceName sets PeerInterfaceName field to given value.

### HasPeerInterfaceName

`func (o *FabricPeerInterfaceCompatibilityIssue) HasPeerInterfaceName() bool`

HasPeerInterfaceName returns a boolean if a field has been set.

### GetPeerValue

`func (o *FabricPeerInterfaceCompatibilityIssue) GetPeerValue() string`

GetPeerValue returns the PeerValue field if non-nil, zero value otherwise.

### GetPeerValueOk

`func (o *FabricPeerInterfaceCompatibilityIssue) GetPeerValueOk() (*string, bool)`

GetPeerValueOk returns a tuple with the PeerValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerValue

`func (o *FabricPeerInterfaceCompatibilityIssue) SetPeerValue(v string)`

SetPeerValue sets PeerValue field to given value.

### HasPeerValue

`func (o *FabricPeerInterfaceCompatibilityIssue) HasPeerValue() bool`

HasPeerValue returns a boolean if a field has been set.

### GetPeerValueModTime

`func (o *FabricPeerInterfaceCompatibilityIssue) GetPeerValueModTime() string`

GetPeerValueModTime returns the PeerValueModTime field if non-nil, zero value otherwise.

### GetPeerValueModTimeOk

`func (o *FabricPeerInterfaceCompatibilityIssue) GetPeerValueModTimeOk() (*string, bool)`

GetPeerValueModTimeOk returns a tuple with the PeerValueModTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerValueModTime

`func (o *FabricPeerInterfaceCompatibilityIssue) SetPeerValueModTime(v string)`

SetPeerValueModTime sets PeerValueModTime field to given value.

### HasPeerValueModTime

`func (o *FabricPeerInterfaceCompatibilityIssue) HasPeerValueModTime() bool`

HasPeerValueModTime returns a boolean if a field has been set.

### GetSource

`func (o *FabricPeerInterfaceCompatibilityIssue) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *FabricPeerInterfaceCompatibilityIssue) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *FabricPeerInterfaceCompatibilityIssue) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *FabricPeerInterfaceCompatibilityIssue) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetSourcePolicy

`func (o *FabricPeerInterfaceCompatibilityIssue) GetSourcePolicy() MoMoRef`

GetSourcePolicy returns the SourcePolicy field if non-nil, zero value otherwise.

### GetSourcePolicyOk

`func (o *FabricPeerInterfaceCompatibilityIssue) GetSourcePolicyOk() (*MoMoRef, bool)`

GetSourcePolicyOk returns a tuple with the SourcePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourcePolicy

`func (o *FabricPeerInterfaceCompatibilityIssue) SetSourcePolicy(v MoMoRef)`

SetSourcePolicy sets SourcePolicy field to given value.

### HasSourcePolicy

`func (o *FabricPeerInterfaceCompatibilityIssue) HasSourcePolicy() bool`

HasSourcePolicy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


