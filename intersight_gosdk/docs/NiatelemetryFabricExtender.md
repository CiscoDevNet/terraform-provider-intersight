# NiatelemetryFabricExtender

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "niatelemetry.FabricExtender"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "niatelemetry.FabricExtender"]
**Description** | Pointer to **string** | Descriptive text configured for the fabric extender. | [optional] [readonly] 
**Id** | Pointer to **int64** | Numeric identifier assigned to the fabric extender. | [optional] [readonly] 
**InterfaceNames** | Pointer to **[]string** |  | [optional] 
**Model** | Pointer to **string** | Hardware model reported for the fabric extender. | [optional] [readonly] 
**Name** | Pointer to **string** | Display name assigned to the fabric extender. | [optional] [readonly] 
**PinningMaxLinks** | Pointer to **int64** | Maximum number of pinning links supported by the fabric extender. | [optional] [readonly] 
**PortChannel** | Pointer to **string** | Port channel associated with the fabric extender. | [optional] [readonly] 
**SerialNumber** | Pointer to **string** | Serial number reported for the fabric extender. | [optional] [readonly] 
**SoftwareVersion** | Pointer to **string** | Software version running on the fabric extender. | [optional] [readonly] 
**State** | Pointer to **string** | Current operational state reported for the fabric extender. | [optional] [readonly] 
**VpcId** | Pointer to **int64** | Optional virtual port channel identifier for the fabric extender. | [optional] [readonly] 

## Methods

### NewNiatelemetryFabricExtender

`func NewNiatelemetryFabricExtender(classId string, objectType string, ) *NiatelemetryFabricExtender`

NewNiatelemetryFabricExtender instantiates a new NiatelemetryFabricExtender object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiatelemetryFabricExtenderWithDefaults

`func NewNiatelemetryFabricExtenderWithDefaults() *NiatelemetryFabricExtender`

NewNiatelemetryFabricExtenderWithDefaults instantiates a new NiatelemetryFabricExtender object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NiatelemetryFabricExtender) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NiatelemetryFabricExtender) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NiatelemetryFabricExtender) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NiatelemetryFabricExtender) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NiatelemetryFabricExtender) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NiatelemetryFabricExtender) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDescription

`func (o *NiatelemetryFabricExtender) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NiatelemetryFabricExtender) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *NiatelemetryFabricExtender) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *NiatelemetryFabricExtender) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *NiatelemetryFabricExtender) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NiatelemetryFabricExtender) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NiatelemetryFabricExtender) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *NiatelemetryFabricExtender) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInterfaceNames

`func (o *NiatelemetryFabricExtender) GetInterfaceNames() []string`

GetInterfaceNames returns the InterfaceNames field if non-nil, zero value otherwise.

### GetInterfaceNamesOk

`func (o *NiatelemetryFabricExtender) GetInterfaceNamesOk() (*[]string, bool)`

GetInterfaceNamesOk returns a tuple with the InterfaceNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceNames

`func (o *NiatelemetryFabricExtender) SetInterfaceNames(v []string)`

SetInterfaceNames sets InterfaceNames field to given value.

### HasInterfaceNames

`func (o *NiatelemetryFabricExtender) HasInterfaceNames() bool`

HasInterfaceNames returns a boolean if a field has been set.

### SetInterfaceNamesNil

`func (o *NiatelemetryFabricExtender) SetInterfaceNamesNil(b bool)`

 SetInterfaceNamesNil sets the value for InterfaceNames to be an explicit nil

### UnsetInterfaceNames
`func (o *NiatelemetryFabricExtender) UnsetInterfaceNames()`

UnsetInterfaceNames ensures that no value is present for InterfaceNames, not even an explicit nil
### GetModel

`func (o *NiatelemetryFabricExtender) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *NiatelemetryFabricExtender) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *NiatelemetryFabricExtender) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *NiatelemetryFabricExtender) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetName

`func (o *NiatelemetryFabricExtender) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NiatelemetryFabricExtender) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NiatelemetryFabricExtender) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NiatelemetryFabricExtender) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPinningMaxLinks

`func (o *NiatelemetryFabricExtender) GetPinningMaxLinks() int64`

GetPinningMaxLinks returns the PinningMaxLinks field if non-nil, zero value otherwise.

### GetPinningMaxLinksOk

`func (o *NiatelemetryFabricExtender) GetPinningMaxLinksOk() (*int64, bool)`

GetPinningMaxLinksOk returns a tuple with the PinningMaxLinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPinningMaxLinks

`func (o *NiatelemetryFabricExtender) SetPinningMaxLinks(v int64)`

SetPinningMaxLinks sets PinningMaxLinks field to given value.

### HasPinningMaxLinks

`func (o *NiatelemetryFabricExtender) HasPinningMaxLinks() bool`

HasPinningMaxLinks returns a boolean if a field has been set.

### GetPortChannel

`func (o *NiatelemetryFabricExtender) GetPortChannel() string`

GetPortChannel returns the PortChannel field if non-nil, zero value otherwise.

### GetPortChannelOk

`func (o *NiatelemetryFabricExtender) GetPortChannelOk() (*string, bool)`

GetPortChannelOk returns a tuple with the PortChannel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortChannel

`func (o *NiatelemetryFabricExtender) SetPortChannel(v string)`

SetPortChannel sets PortChannel field to given value.

### HasPortChannel

`func (o *NiatelemetryFabricExtender) HasPortChannel() bool`

HasPortChannel returns a boolean if a field has been set.

### GetSerialNumber

`func (o *NiatelemetryFabricExtender) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *NiatelemetryFabricExtender) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *NiatelemetryFabricExtender) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *NiatelemetryFabricExtender) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### GetSoftwareVersion

`func (o *NiatelemetryFabricExtender) GetSoftwareVersion() string`

GetSoftwareVersion returns the SoftwareVersion field if non-nil, zero value otherwise.

### GetSoftwareVersionOk

`func (o *NiatelemetryFabricExtender) GetSoftwareVersionOk() (*string, bool)`

GetSoftwareVersionOk returns a tuple with the SoftwareVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareVersion

`func (o *NiatelemetryFabricExtender) SetSoftwareVersion(v string)`

SetSoftwareVersion sets SoftwareVersion field to given value.

### HasSoftwareVersion

`func (o *NiatelemetryFabricExtender) HasSoftwareVersion() bool`

HasSoftwareVersion returns a boolean if a field has been set.

### GetState

`func (o *NiatelemetryFabricExtender) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *NiatelemetryFabricExtender) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *NiatelemetryFabricExtender) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *NiatelemetryFabricExtender) HasState() bool`

HasState returns a boolean if a field has been set.

### GetVpcId

`func (o *NiatelemetryFabricExtender) GetVpcId() int64`

GetVpcId returns the VpcId field if non-nil, zero value otherwise.

### GetVpcIdOk

`func (o *NiatelemetryFabricExtender) GetVpcIdOk() (*int64, bool)`

GetVpcIdOk returns a tuple with the VpcId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcId

`func (o *NiatelemetryFabricExtender) SetVpcId(v int64)`

SetVpcId sets VpcId field to given value.

### HasVpcId

`func (o *NiatelemetryFabricExtender) HasVpcId() bool`

HasVpcId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


