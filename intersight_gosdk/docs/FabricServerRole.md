# FabricServerRole

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "fabric.ServerRole"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "fabric.ServerRole"]
**AutoNegotiationDisabled** | Pointer to **bool** | Auto negotiation configuration for server port. This configuration is required only for FEX Model N9K-C93180YC-FX3 connected with 100G speed port on UCS-FI-6536 and should be set as True. | [optional] [default to false]
**Fec** | Pointer to **string** | Forward error correction configuration for server port. This configuration is required only for FEX Model N9K-C93180YC-FX3 connected with 25G speed ports on UCS-FI-6454/UCS-FI-64108/UCS-FI-6652 and should be set as Cl74. * &#x60;Auto&#x60; - Forward error correction option &#39;Auto&#39;. Supported speeds are Auto, 1Gbps, 10Gbps, 25Gbps, 40Gbps and 100 Gbps. * &#x60;Cl91&#x60; - Forward error correction option &#39;cl91&#39;. Supported speeds are 25Gbps and 100 Gbps. * &#x60;Cl74&#x60; - Forward error correction option &#39;cl74&#39;. Supported speeds are 25Gbps. * &#x60;rs-cons16&#x60; - Forward error correction option \&quot;rs-cons16\&quot;. Supported speeds are 25Gbps. * &#x60;rs-ieee&#x60; - Forward error correction option \&quot;rs-ieee\&quot;. Supported speeds are 25Gbps. * &#x60;Off&#x60; - Turn off forward error correction. Supported speeds are 25Gbps and 100 Gbps. | [optional] [default to "Auto"]
**PreferredDeviceId** | Pointer to **int64** | Preferred device ID to be configured by user for the connected device. This ID must be specified together with the &#39;PreferredDeviceType&#39; property. This ID will only takes effect if the actual connected device matches the &#39;PreferredDeviceType&#39;. If the preferred ID is not available, the ID is automatically allocated and assigned by the system. If different preferred IDs are specified for the ports connected to the same device, only the preferred ID (if specified) of the port that is discovered first will be considered. | [optional] 
**PreferredDeviceType** | Pointer to **string** | Device type for which preferred ID to be configured. If the actual connected device does not match the specified device type, the system ignores the &#39;PreferredDeviceId&#39; property. * &#x60;Auto&#x60; - Preferred Id will be ignored if specified with this type. * &#x60;RackServer&#x60; - Connected device type is Rack Unit Server. * &#x60;Chassis&#x60; - Connected device type is Chassis. | [optional] [default to "Auto"]

## Methods

### NewFabricServerRole

`func NewFabricServerRole(classId string, objectType string, ) *FabricServerRole`

NewFabricServerRole instantiates a new FabricServerRole object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFabricServerRoleWithDefaults

`func NewFabricServerRoleWithDefaults() *FabricServerRole`

NewFabricServerRoleWithDefaults instantiates a new FabricServerRole object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *FabricServerRole) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *FabricServerRole) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *FabricServerRole) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *FabricServerRole) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *FabricServerRole) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *FabricServerRole) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAutoNegotiationDisabled

`func (o *FabricServerRole) GetAutoNegotiationDisabled() bool`

GetAutoNegotiationDisabled returns the AutoNegotiationDisabled field if non-nil, zero value otherwise.

### GetAutoNegotiationDisabledOk

`func (o *FabricServerRole) GetAutoNegotiationDisabledOk() (*bool, bool)`

GetAutoNegotiationDisabledOk returns a tuple with the AutoNegotiationDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoNegotiationDisabled

`func (o *FabricServerRole) SetAutoNegotiationDisabled(v bool)`

SetAutoNegotiationDisabled sets AutoNegotiationDisabled field to given value.

### HasAutoNegotiationDisabled

`func (o *FabricServerRole) HasAutoNegotiationDisabled() bool`

HasAutoNegotiationDisabled returns a boolean if a field has been set.

### GetFec

`func (o *FabricServerRole) GetFec() string`

GetFec returns the Fec field if non-nil, zero value otherwise.

### GetFecOk

`func (o *FabricServerRole) GetFecOk() (*string, bool)`

GetFecOk returns a tuple with the Fec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFec

`func (o *FabricServerRole) SetFec(v string)`

SetFec sets Fec field to given value.

### HasFec

`func (o *FabricServerRole) HasFec() bool`

HasFec returns a boolean if a field has been set.

### GetPreferredDeviceId

`func (o *FabricServerRole) GetPreferredDeviceId() int64`

GetPreferredDeviceId returns the PreferredDeviceId field if non-nil, zero value otherwise.

### GetPreferredDeviceIdOk

`func (o *FabricServerRole) GetPreferredDeviceIdOk() (*int64, bool)`

GetPreferredDeviceIdOk returns a tuple with the PreferredDeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredDeviceId

`func (o *FabricServerRole) SetPreferredDeviceId(v int64)`

SetPreferredDeviceId sets PreferredDeviceId field to given value.

### HasPreferredDeviceId

`func (o *FabricServerRole) HasPreferredDeviceId() bool`

HasPreferredDeviceId returns a boolean if a field has been set.

### GetPreferredDeviceType

`func (o *FabricServerRole) GetPreferredDeviceType() string`

GetPreferredDeviceType returns the PreferredDeviceType field if non-nil, zero value otherwise.

### GetPreferredDeviceTypeOk

`func (o *FabricServerRole) GetPreferredDeviceTypeOk() (*string, bool)`

GetPreferredDeviceTypeOk returns a tuple with the PreferredDeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredDeviceType

`func (o *FabricServerRole) SetPreferredDeviceType(v string)`

SetPreferredDeviceType sets PreferredDeviceType field to given value.

### HasPreferredDeviceType

`func (o *FabricServerRole) HasPreferredDeviceType() bool`

HasPreferredDeviceType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


