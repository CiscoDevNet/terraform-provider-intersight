# FabricFcUplinkPcRole

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "fabric.FcUplinkPcRole"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "fabric.FcUplinkPcRole"]
**AdminSpeed** | Pointer to **string** | Admin configured speed for the port. * &#x60;16Gbps&#x60; - Admin configurable speed 16Gbps. * &#x60;8Gbps&#x60; - Admin configurable speed 8Gbps. * &#x60;32Gbps&#x60; - Admin configurable speed 32Gbps. * &#x60;Auto&#x60; - Admin configurable speed AUTO ( default ). | [optional] [default to "16Gbps"]
**FillPattern** | Pointer to **string** | Fill pattern to differentiate the configs in NPIV. * &#x60;Idle&#x60; - Fc Fill Pattern type Idle. * &#x60;Arbff&#x60; - Fc Fill Pattern type Arbff. | [optional] [default to "Idle"]
**VsanId** | Pointer to **int64** | Virtual San Identifier associated to the FC port. | [optional] 

## Methods

### NewFabricFcUplinkPcRole

`func NewFabricFcUplinkPcRole(classId string, objectType string, ) *FabricFcUplinkPcRole`

NewFabricFcUplinkPcRole instantiates a new FabricFcUplinkPcRole object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFabricFcUplinkPcRoleWithDefaults

`func NewFabricFcUplinkPcRoleWithDefaults() *FabricFcUplinkPcRole`

NewFabricFcUplinkPcRoleWithDefaults instantiates a new FabricFcUplinkPcRole object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *FabricFcUplinkPcRole) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *FabricFcUplinkPcRole) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *FabricFcUplinkPcRole) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *FabricFcUplinkPcRole) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *FabricFcUplinkPcRole) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *FabricFcUplinkPcRole) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAdminSpeed

`func (o *FabricFcUplinkPcRole) GetAdminSpeed() string`

GetAdminSpeed returns the AdminSpeed field if non-nil, zero value otherwise.

### GetAdminSpeedOk

`func (o *FabricFcUplinkPcRole) GetAdminSpeedOk() (*string, bool)`

GetAdminSpeedOk returns a tuple with the AdminSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminSpeed

`func (o *FabricFcUplinkPcRole) SetAdminSpeed(v string)`

SetAdminSpeed sets AdminSpeed field to given value.

### HasAdminSpeed

`func (o *FabricFcUplinkPcRole) HasAdminSpeed() bool`

HasAdminSpeed returns a boolean if a field has been set.

### GetFillPattern

`func (o *FabricFcUplinkPcRole) GetFillPattern() string`

GetFillPattern returns the FillPattern field if non-nil, zero value otherwise.

### GetFillPatternOk

`func (o *FabricFcUplinkPcRole) GetFillPatternOk() (*string, bool)`

GetFillPatternOk returns a tuple with the FillPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFillPattern

`func (o *FabricFcUplinkPcRole) SetFillPattern(v string)`

SetFillPattern sets FillPattern field to given value.

### HasFillPattern

`func (o *FabricFcUplinkPcRole) HasFillPattern() bool`

HasFillPattern returns a boolean if a field has been set.

### GetVsanId

`func (o *FabricFcUplinkPcRole) GetVsanId() int64`

GetVsanId returns the VsanId field if non-nil, zero value otherwise.

### GetVsanIdOk

`func (o *FabricFcUplinkPcRole) GetVsanIdOk() (*int64, bool)`

GetVsanIdOk returns a tuple with the VsanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVsanId

`func (o *FabricFcUplinkPcRole) SetVsanId(v int64)`

SetVsanId sets VsanId field to given value.

### HasVsanId

`func (o *FabricFcUplinkPcRole) HasVsanId() bool`

HasVsanId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


