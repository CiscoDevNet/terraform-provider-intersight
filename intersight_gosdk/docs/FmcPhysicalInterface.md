# FmcPhysicalInterface

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "fmc.PhysicalInterface"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "fmc.PhysicalInterface"]
**DeviceId** | Pointer to **string** | Physical interface device Id. | [optional] [readonly] 
**DomainId** | Pointer to **string** | Physical interface domain Id. | [optional] [readonly] 
**Name** | Pointer to **string** | The name for the physical interface. | [optional] [readonly] 
**PhysicalInterfaceId** | Pointer to **string** | The id for the physical interface. | [optional] [readonly] 

## Methods

### NewFmcPhysicalInterface

`func NewFmcPhysicalInterface(classId string, objectType string, ) *FmcPhysicalInterface`

NewFmcPhysicalInterface instantiates a new FmcPhysicalInterface object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFmcPhysicalInterfaceWithDefaults

`func NewFmcPhysicalInterfaceWithDefaults() *FmcPhysicalInterface`

NewFmcPhysicalInterfaceWithDefaults instantiates a new FmcPhysicalInterface object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *FmcPhysicalInterface) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *FmcPhysicalInterface) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *FmcPhysicalInterface) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *FmcPhysicalInterface) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *FmcPhysicalInterface) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *FmcPhysicalInterface) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDeviceId

`func (o *FmcPhysicalInterface) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *FmcPhysicalInterface) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *FmcPhysicalInterface) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *FmcPhysicalInterface) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### GetDomainId

`func (o *FmcPhysicalInterface) GetDomainId() string`

GetDomainId returns the DomainId field if non-nil, zero value otherwise.

### GetDomainIdOk

`func (o *FmcPhysicalInterface) GetDomainIdOk() (*string, bool)`

GetDomainIdOk returns a tuple with the DomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainId

`func (o *FmcPhysicalInterface) SetDomainId(v string)`

SetDomainId sets DomainId field to given value.

### HasDomainId

`func (o *FmcPhysicalInterface) HasDomainId() bool`

HasDomainId returns a boolean if a field has been set.

### GetName

`func (o *FmcPhysicalInterface) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FmcPhysicalInterface) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FmcPhysicalInterface) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FmcPhysicalInterface) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPhysicalInterfaceId

`func (o *FmcPhysicalInterface) GetPhysicalInterfaceId() string`

GetPhysicalInterfaceId returns the PhysicalInterfaceId field if non-nil, zero value otherwise.

### GetPhysicalInterfaceIdOk

`func (o *FmcPhysicalInterface) GetPhysicalInterfaceIdOk() (*string, bool)`

GetPhysicalInterfaceIdOk returns a tuple with the PhysicalInterfaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalInterfaceId

`func (o *FmcPhysicalInterface) SetPhysicalInterfaceId(v string)`

SetPhysicalInterfaceId sets PhysicalInterfaceId field to given value.

### HasPhysicalInterfaceId

`func (o *FmcPhysicalInterface) HasPhysicalInterfaceId() bool`

HasPhysicalInterfaceId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


