# NiatelemetryLink

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "niatelemetry.Link"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "niatelemetry.Link"]
**DstFabricName** | Pointer to **string** | Destination fabric name for the link. | [optional] [readonly] 
**DstInterfaceName** | Pointer to **string** | Destination interface name for the link. | [optional] [readonly] 
**DstSwitchId** | Pointer to **string** | Destination switch identifier for the link. | [optional] [readonly] 
**DstSwitchModelName** | Pointer to **string** | Destination switch model for the link. | [optional] [readonly] 
**DstSwitchName** | Pointer to **string** | Destination switch name for the link. | [optional] [readonly] 
**DstSwitchRole** | Pointer to **string** | Destination switch role for the link. | [optional] [readonly] 
**LinkDiscovered** | Pointer to **bool** | Indicates whether the link is discovered. | [optional] [readonly] 
**LinkId** | Pointer to **string** | Unique identifier for the link. | [optional] [readonly] 
**LinkPlanned** | Pointer to **bool** | Indicates whether the link is planned. | [optional] [readonly] 
**LinkPresent** | Pointer to **bool** | Indicates whether the link is present. | [optional] [readonly] 
**LinkType** | Pointer to **string** | A description of the type of link. | [optional] [readonly] 
**PortChannel** | Pointer to **bool** | Indicates whether the link is a port channel. | [optional] [readonly] 
**SrcFabricName** | Pointer to **string** | Source fabric name for the link. | [optional] [readonly] 
**SrcInterfaceAdminStatus** | Pointer to **string** | Administrative status of the source interface. | [optional] [readonly] 
**SrcInterfaceName** | Pointer to **string** | Source interface name for the link. | [optional] [readonly] 
**SrcInterfaceOperStatus** | Pointer to **string** | Operational status of the source interface. | [optional] [readonly] 
**SrcSwitchId** | Pointer to **string** | Source switch identifier for the link. | [optional] [readonly] 
**SrcSwitchModelName** | Pointer to **string** | Source switch model for the link. | [optional] [readonly] 
**SrcSwitchName** | Pointer to **string** | Source switch name for the link. | [optional] [readonly] 
**SrcSwitchRole** | Pointer to **string** | Source switch role for the link. | [optional] [readonly] 
**Fabric** | Pointer to [**NullableNiatelemetryFabricRelationship**](NiatelemetryFabricRelationship.md) |  | [optional] 

## Methods

### NewNiatelemetryLink

`func NewNiatelemetryLink(classId string, objectType string, ) *NiatelemetryLink`

NewNiatelemetryLink instantiates a new NiatelemetryLink object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiatelemetryLinkWithDefaults

`func NewNiatelemetryLinkWithDefaults() *NiatelemetryLink`

NewNiatelemetryLinkWithDefaults instantiates a new NiatelemetryLink object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NiatelemetryLink) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NiatelemetryLink) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NiatelemetryLink) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NiatelemetryLink) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NiatelemetryLink) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NiatelemetryLink) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDstFabricName

`func (o *NiatelemetryLink) GetDstFabricName() string`

GetDstFabricName returns the DstFabricName field if non-nil, zero value otherwise.

### GetDstFabricNameOk

`func (o *NiatelemetryLink) GetDstFabricNameOk() (*string, bool)`

GetDstFabricNameOk returns a tuple with the DstFabricName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDstFabricName

`func (o *NiatelemetryLink) SetDstFabricName(v string)`

SetDstFabricName sets DstFabricName field to given value.

### HasDstFabricName

`func (o *NiatelemetryLink) HasDstFabricName() bool`

HasDstFabricName returns a boolean if a field has been set.

### GetDstInterfaceName

`func (o *NiatelemetryLink) GetDstInterfaceName() string`

GetDstInterfaceName returns the DstInterfaceName field if non-nil, zero value otherwise.

### GetDstInterfaceNameOk

`func (o *NiatelemetryLink) GetDstInterfaceNameOk() (*string, bool)`

GetDstInterfaceNameOk returns a tuple with the DstInterfaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDstInterfaceName

`func (o *NiatelemetryLink) SetDstInterfaceName(v string)`

SetDstInterfaceName sets DstInterfaceName field to given value.

### HasDstInterfaceName

`func (o *NiatelemetryLink) HasDstInterfaceName() bool`

HasDstInterfaceName returns a boolean if a field has been set.

### GetDstSwitchId

`func (o *NiatelemetryLink) GetDstSwitchId() string`

GetDstSwitchId returns the DstSwitchId field if non-nil, zero value otherwise.

### GetDstSwitchIdOk

`func (o *NiatelemetryLink) GetDstSwitchIdOk() (*string, bool)`

GetDstSwitchIdOk returns a tuple with the DstSwitchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDstSwitchId

`func (o *NiatelemetryLink) SetDstSwitchId(v string)`

SetDstSwitchId sets DstSwitchId field to given value.

### HasDstSwitchId

`func (o *NiatelemetryLink) HasDstSwitchId() bool`

HasDstSwitchId returns a boolean if a field has been set.

### GetDstSwitchModelName

`func (o *NiatelemetryLink) GetDstSwitchModelName() string`

GetDstSwitchModelName returns the DstSwitchModelName field if non-nil, zero value otherwise.

### GetDstSwitchModelNameOk

`func (o *NiatelemetryLink) GetDstSwitchModelNameOk() (*string, bool)`

GetDstSwitchModelNameOk returns a tuple with the DstSwitchModelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDstSwitchModelName

`func (o *NiatelemetryLink) SetDstSwitchModelName(v string)`

SetDstSwitchModelName sets DstSwitchModelName field to given value.

### HasDstSwitchModelName

`func (o *NiatelemetryLink) HasDstSwitchModelName() bool`

HasDstSwitchModelName returns a boolean if a field has been set.

### GetDstSwitchName

`func (o *NiatelemetryLink) GetDstSwitchName() string`

GetDstSwitchName returns the DstSwitchName field if non-nil, zero value otherwise.

### GetDstSwitchNameOk

`func (o *NiatelemetryLink) GetDstSwitchNameOk() (*string, bool)`

GetDstSwitchNameOk returns a tuple with the DstSwitchName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDstSwitchName

`func (o *NiatelemetryLink) SetDstSwitchName(v string)`

SetDstSwitchName sets DstSwitchName field to given value.

### HasDstSwitchName

`func (o *NiatelemetryLink) HasDstSwitchName() bool`

HasDstSwitchName returns a boolean if a field has been set.

### GetDstSwitchRole

`func (o *NiatelemetryLink) GetDstSwitchRole() string`

GetDstSwitchRole returns the DstSwitchRole field if non-nil, zero value otherwise.

### GetDstSwitchRoleOk

`func (o *NiatelemetryLink) GetDstSwitchRoleOk() (*string, bool)`

GetDstSwitchRoleOk returns a tuple with the DstSwitchRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDstSwitchRole

`func (o *NiatelemetryLink) SetDstSwitchRole(v string)`

SetDstSwitchRole sets DstSwitchRole field to given value.

### HasDstSwitchRole

`func (o *NiatelemetryLink) HasDstSwitchRole() bool`

HasDstSwitchRole returns a boolean if a field has been set.

### GetLinkDiscovered

`func (o *NiatelemetryLink) GetLinkDiscovered() bool`

GetLinkDiscovered returns the LinkDiscovered field if non-nil, zero value otherwise.

### GetLinkDiscoveredOk

`func (o *NiatelemetryLink) GetLinkDiscoveredOk() (*bool, bool)`

GetLinkDiscoveredOk returns a tuple with the LinkDiscovered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkDiscovered

`func (o *NiatelemetryLink) SetLinkDiscovered(v bool)`

SetLinkDiscovered sets LinkDiscovered field to given value.

### HasLinkDiscovered

`func (o *NiatelemetryLink) HasLinkDiscovered() bool`

HasLinkDiscovered returns a boolean if a field has been set.

### GetLinkId

`func (o *NiatelemetryLink) GetLinkId() string`

GetLinkId returns the LinkId field if non-nil, zero value otherwise.

### GetLinkIdOk

`func (o *NiatelemetryLink) GetLinkIdOk() (*string, bool)`

GetLinkIdOk returns a tuple with the LinkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkId

`func (o *NiatelemetryLink) SetLinkId(v string)`

SetLinkId sets LinkId field to given value.

### HasLinkId

`func (o *NiatelemetryLink) HasLinkId() bool`

HasLinkId returns a boolean if a field has been set.

### GetLinkPlanned

`func (o *NiatelemetryLink) GetLinkPlanned() bool`

GetLinkPlanned returns the LinkPlanned field if non-nil, zero value otherwise.

### GetLinkPlannedOk

`func (o *NiatelemetryLink) GetLinkPlannedOk() (*bool, bool)`

GetLinkPlannedOk returns a tuple with the LinkPlanned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkPlanned

`func (o *NiatelemetryLink) SetLinkPlanned(v bool)`

SetLinkPlanned sets LinkPlanned field to given value.

### HasLinkPlanned

`func (o *NiatelemetryLink) HasLinkPlanned() bool`

HasLinkPlanned returns a boolean if a field has been set.

### GetLinkPresent

`func (o *NiatelemetryLink) GetLinkPresent() bool`

GetLinkPresent returns the LinkPresent field if non-nil, zero value otherwise.

### GetLinkPresentOk

`func (o *NiatelemetryLink) GetLinkPresentOk() (*bool, bool)`

GetLinkPresentOk returns a tuple with the LinkPresent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkPresent

`func (o *NiatelemetryLink) SetLinkPresent(v bool)`

SetLinkPresent sets LinkPresent field to given value.

### HasLinkPresent

`func (o *NiatelemetryLink) HasLinkPresent() bool`

HasLinkPresent returns a boolean if a field has been set.

### GetLinkType

`func (o *NiatelemetryLink) GetLinkType() string`

GetLinkType returns the LinkType field if non-nil, zero value otherwise.

### GetLinkTypeOk

`func (o *NiatelemetryLink) GetLinkTypeOk() (*string, bool)`

GetLinkTypeOk returns a tuple with the LinkType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkType

`func (o *NiatelemetryLink) SetLinkType(v string)`

SetLinkType sets LinkType field to given value.

### HasLinkType

`func (o *NiatelemetryLink) HasLinkType() bool`

HasLinkType returns a boolean if a field has been set.

### GetPortChannel

`func (o *NiatelemetryLink) GetPortChannel() bool`

GetPortChannel returns the PortChannel field if non-nil, zero value otherwise.

### GetPortChannelOk

`func (o *NiatelemetryLink) GetPortChannelOk() (*bool, bool)`

GetPortChannelOk returns a tuple with the PortChannel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortChannel

`func (o *NiatelemetryLink) SetPortChannel(v bool)`

SetPortChannel sets PortChannel field to given value.

### HasPortChannel

`func (o *NiatelemetryLink) HasPortChannel() bool`

HasPortChannel returns a boolean if a field has been set.

### GetSrcFabricName

`func (o *NiatelemetryLink) GetSrcFabricName() string`

GetSrcFabricName returns the SrcFabricName field if non-nil, zero value otherwise.

### GetSrcFabricNameOk

`func (o *NiatelemetryLink) GetSrcFabricNameOk() (*string, bool)`

GetSrcFabricNameOk returns a tuple with the SrcFabricName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcFabricName

`func (o *NiatelemetryLink) SetSrcFabricName(v string)`

SetSrcFabricName sets SrcFabricName field to given value.

### HasSrcFabricName

`func (o *NiatelemetryLink) HasSrcFabricName() bool`

HasSrcFabricName returns a boolean if a field has been set.

### GetSrcInterfaceAdminStatus

`func (o *NiatelemetryLink) GetSrcInterfaceAdminStatus() string`

GetSrcInterfaceAdminStatus returns the SrcInterfaceAdminStatus field if non-nil, zero value otherwise.

### GetSrcInterfaceAdminStatusOk

`func (o *NiatelemetryLink) GetSrcInterfaceAdminStatusOk() (*string, bool)`

GetSrcInterfaceAdminStatusOk returns a tuple with the SrcInterfaceAdminStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcInterfaceAdminStatus

`func (o *NiatelemetryLink) SetSrcInterfaceAdminStatus(v string)`

SetSrcInterfaceAdminStatus sets SrcInterfaceAdminStatus field to given value.

### HasSrcInterfaceAdminStatus

`func (o *NiatelemetryLink) HasSrcInterfaceAdminStatus() bool`

HasSrcInterfaceAdminStatus returns a boolean if a field has been set.

### GetSrcInterfaceName

`func (o *NiatelemetryLink) GetSrcInterfaceName() string`

GetSrcInterfaceName returns the SrcInterfaceName field if non-nil, zero value otherwise.

### GetSrcInterfaceNameOk

`func (o *NiatelemetryLink) GetSrcInterfaceNameOk() (*string, bool)`

GetSrcInterfaceNameOk returns a tuple with the SrcInterfaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcInterfaceName

`func (o *NiatelemetryLink) SetSrcInterfaceName(v string)`

SetSrcInterfaceName sets SrcInterfaceName field to given value.

### HasSrcInterfaceName

`func (o *NiatelemetryLink) HasSrcInterfaceName() bool`

HasSrcInterfaceName returns a boolean if a field has been set.

### GetSrcInterfaceOperStatus

`func (o *NiatelemetryLink) GetSrcInterfaceOperStatus() string`

GetSrcInterfaceOperStatus returns the SrcInterfaceOperStatus field if non-nil, zero value otherwise.

### GetSrcInterfaceOperStatusOk

`func (o *NiatelemetryLink) GetSrcInterfaceOperStatusOk() (*string, bool)`

GetSrcInterfaceOperStatusOk returns a tuple with the SrcInterfaceOperStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcInterfaceOperStatus

`func (o *NiatelemetryLink) SetSrcInterfaceOperStatus(v string)`

SetSrcInterfaceOperStatus sets SrcInterfaceOperStatus field to given value.

### HasSrcInterfaceOperStatus

`func (o *NiatelemetryLink) HasSrcInterfaceOperStatus() bool`

HasSrcInterfaceOperStatus returns a boolean if a field has been set.

### GetSrcSwitchId

`func (o *NiatelemetryLink) GetSrcSwitchId() string`

GetSrcSwitchId returns the SrcSwitchId field if non-nil, zero value otherwise.

### GetSrcSwitchIdOk

`func (o *NiatelemetryLink) GetSrcSwitchIdOk() (*string, bool)`

GetSrcSwitchIdOk returns a tuple with the SrcSwitchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcSwitchId

`func (o *NiatelemetryLink) SetSrcSwitchId(v string)`

SetSrcSwitchId sets SrcSwitchId field to given value.

### HasSrcSwitchId

`func (o *NiatelemetryLink) HasSrcSwitchId() bool`

HasSrcSwitchId returns a boolean if a field has been set.

### GetSrcSwitchModelName

`func (o *NiatelemetryLink) GetSrcSwitchModelName() string`

GetSrcSwitchModelName returns the SrcSwitchModelName field if non-nil, zero value otherwise.

### GetSrcSwitchModelNameOk

`func (o *NiatelemetryLink) GetSrcSwitchModelNameOk() (*string, bool)`

GetSrcSwitchModelNameOk returns a tuple with the SrcSwitchModelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcSwitchModelName

`func (o *NiatelemetryLink) SetSrcSwitchModelName(v string)`

SetSrcSwitchModelName sets SrcSwitchModelName field to given value.

### HasSrcSwitchModelName

`func (o *NiatelemetryLink) HasSrcSwitchModelName() bool`

HasSrcSwitchModelName returns a boolean if a field has been set.

### GetSrcSwitchName

`func (o *NiatelemetryLink) GetSrcSwitchName() string`

GetSrcSwitchName returns the SrcSwitchName field if non-nil, zero value otherwise.

### GetSrcSwitchNameOk

`func (o *NiatelemetryLink) GetSrcSwitchNameOk() (*string, bool)`

GetSrcSwitchNameOk returns a tuple with the SrcSwitchName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcSwitchName

`func (o *NiatelemetryLink) SetSrcSwitchName(v string)`

SetSrcSwitchName sets SrcSwitchName field to given value.

### HasSrcSwitchName

`func (o *NiatelemetryLink) HasSrcSwitchName() bool`

HasSrcSwitchName returns a boolean if a field has been set.

### GetSrcSwitchRole

`func (o *NiatelemetryLink) GetSrcSwitchRole() string`

GetSrcSwitchRole returns the SrcSwitchRole field if non-nil, zero value otherwise.

### GetSrcSwitchRoleOk

`func (o *NiatelemetryLink) GetSrcSwitchRoleOk() (*string, bool)`

GetSrcSwitchRoleOk returns a tuple with the SrcSwitchRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcSwitchRole

`func (o *NiatelemetryLink) SetSrcSwitchRole(v string)`

SetSrcSwitchRole sets SrcSwitchRole field to given value.

### HasSrcSwitchRole

`func (o *NiatelemetryLink) HasSrcSwitchRole() bool`

HasSrcSwitchRole returns a boolean if a field has been set.

### GetFabric

`func (o *NiatelemetryLink) GetFabric() NiatelemetryFabricRelationship`

GetFabric returns the Fabric field if non-nil, zero value otherwise.

### GetFabricOk

`func (o *NiatelemetryLink) GetFabricOk() (*NiatelemetryFabricRelationship, bool)`

GetFabricOk returns a tuple with the Fabric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFabric

`func (o *NiatelemetryLink) SetFabric(v NiatelemetryFabricRelationship)`

SetFabric sets Fabric field to given value.

### HasFabric

`func (o *NiatelemetryLink) HasFabric() bool`

HasFabric returns a boolean if a field has been set.

### SetFabricNil

`func (o *NiatelemetryLink) SetFabricNil(b bool)`

 SetFabricNil sets the value for Fabric to be an explicit nil

### UnsetFabric
`func (o *NiatelemetryLink) UnsetFabric()`

UnsetFabric ensures that no value is present for Fabric, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


