# EtherPeerInterfaceCompatibilityIssues

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "ether.PeerInterfaceCompatibilityIssues"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "ether.PeerInterfaceCompatibilityIssues"]
**ActiveIssueCount** | Pointer to **int64** | The count of active compatibility issues with this peer interface. | [optional] [readonly] 
**Issues** | Pointer to [**[]EtherPeerPortCompatibilityDetail**](EtherPeerPortCompatibilityDetail.md) |  | [optional] 
**PeerDevice** | Pointer to [**MoMoRef**](MoMoRef.md) |  | [optional] 
**PeerInterface** | Pointer to [**MoMoRef**](MoMoRef.md) |  | [optional] 
**PeerInterfaceName** | Pointer to **string** | User-friendly name of the peer interface used for compatibility evaluation. | [optional] [readonly] 

## Methods

### NewEtherPeerInterfaceCompatibilityIssues

`func NewEtherPeerInterfaceCompatibilityIssues(classId string, objectType string, ) *EtherPeerInterfaceCompatibilityIssues`

NewEtherPeerInterfaceCompatibilityIssues instantiates a new EtherPeerInterfaceCompatibilityIssues object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEtherPeerInterfaceCompatibilityIssuesWithDefaults

`func NewEtherPeerInterfaceCompatibilityIssuesWithDefaults() *EtherPeerInterfaceCompatibilityIssues`

NewEtherPeerInterfaceCompatibilityIssuesWithDefaults instantiates a new EtherPeerInterfaceCompatibilityIssues object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *EtherPeerInterfaceCompatibilityIssues) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *EtherPeerInterfaceCompatibilityIssues) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *EtherPeerInterfaceCompatibilityIssues) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *EtherPeerInterfaceCompatibilityIssues) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *EtherPeerInterfaceCompatibilityIssues) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *EtherPeerInterfaceCompatibilityIssues) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetActiveIssueCount

`func (o *EtherPeerInterfaceCompatibilityIssues) GetActiveIssueCount() int64`

GetActiveIssueCount returns the ActiveIssueCount field if non-nil, zero value otherwise.

### GetActiveIssueCountOk

`func (o *EtherPeerInterfaceCompatibilityIssues) GetActiveIssueCountOk() (*int64, bool)`

GetActiveIssueCountOk returns a tuple with the ActiveIssueCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveIssueCount

`func (o *EtherPeerInterfaceCompatibilityIssues) SetActiveIssueCount(v int64)`

SetActiveIssueCount sets ActiveIssueCount field to given value.

### HasActiveIssueCount

`func (o *EtherPeerInterfaceCompatibilityIssues) HasActiveIssueCount() bool`

HasActiveIssueCount returns a boolean if a field has been set.

### GetIssues

`func (o *EtherPeerInterfaceCompatibilityIssues) GetIssues() []EtherPeerPortCompatibilityDetail`

GetIssues returns the Issues field if non-nil, zero value otherwise.

### GetIssuesOk

`func (o *EtherPeerInterfaceCompatibilityIssues) GetIssuesOk() (*[]EtherPeerPortCompatibilityDetail, bool)`

GetIssuesOk returns a tuple with the Issues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssues

`func (o *EtherPeerInterfaceCompatibilityIssues) SetIssues(v []EtherPeerPortCompatibilityDetail)`

SetIssues sets Issues field to given value.

### HasIssues

`func (o *EtherPeerInterfaceCompatibilityIssues) HasIssues() bool`

HasIssues returns a boolean if a field has been set.

### SetIssuesNil

`func (o *EtherPeerInterfaceCompatibilityIssues) SetIssuesNil(b bool)`

 SetIssuesNil sets the value for Issues to be an explicit nil

### UnsetIssues
`func (o *EtherPeerInterfaceCompatibilityIssues) UnsetIssues()`

UnsetIssues ensures that no value is present for Issues, not even an explicit nil
### GetPeerDevice

`func (o *EtherPeerInterfaceCompatibilityIssues) GetPeerDevice() MoMoRef`

GetPeerDevice returns the PeerDevice field if non-nil, zero value otherwise.

### GetPeerDeviceOk

`func (o *EtherPeerInterfaceCompatibilityIssues) GetPeerDeviceOk() (*MoMoRef, bool)`

GetPeerDeviceOk returns a tuple with the PeerDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerDevice

`func (o *EtherPeerInterfaceCompatibilityIssues) SetPeerDevice(v MoMoRef)`

SetPeerDevice sets PeerDevice field to given value.

### HasPeerDevice

`func (o *EtherPeerInterfaceCompatibilityIssues) HasPeerDevice() bool`

HasPeerDevice returns a boolean if a field has been set.

### GetPeerInterface

`func (o *EtherPeerInterfaceCompatibilityIssues) GetPeerInterface() MoMoRef`

GetPeerInterface returns the PeerInterface field if non-nil, zero value otherwise.

### GetPeerInterfaceOk

`func (o *EtherPeerInterfaceCompatibilityIssues) GetPeerInterfaceOk() (*MoMoRef, bool)`

GetPeerInterfaceOk returns a tuple with the PeerInterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerInterface

`func (o *EtherPeerInterfaceCompatibilityIssues) SetPeerInterface(v MoMoRef)`

SetPeerInterface sets PeerInterface field to given value.

### HasPeerInterface

`func (o *EtherPeerInterfaceCompatibilityIssues) HasPeerInterface() bool`

HasPeerInterface returns a boolean if a field has been set.

### GetPeerInterfaceName

`func (o *EtherPeerInterfaceCompatibilityIssues) GetPeerInterfaceName() string`

GetPeerInterfaceName returns the PeerInterfaceName field if non-nil, zero value otherwise.

### GetPeerInterfaceNameOk

`func (o *EtherPeerInterfaceCompatibilityIssues) GetPeerInterfaceNameOk() (*string, bool)`

GetPeerInterfaceNameOk returns a tuple with the PeerInterfaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerInterfaceName

`func (o *EtherPeerInterfaceCompatibilityIssues) SetPeerInterfaceName(v string)`

SetPeerInterfaceName sets PeerInterfaceName field to given value.

### HasPeerInterfaceName

`func (o *EtherPeerInterfaceCompatibilityIssues) HasPeerInterfaceName() bool`

HasPeerInterfaceName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


