# FmcDeviceHaPair

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "fmc.DeviceHaPair"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "fmc.DeviceHaPair"]
**DeviceId** | Pointer to **string** | Unique Identifier of the Device. | [optional] [readonly] 
**DomainId** | Pointer to **string** | Unique Identifier of the Domain. | [optional] [readonly] 
**Name** | Pointer to **string** | A user provided name of the Device. | [optional] [readonly] 

## Methods

### NewFmcDeviceHaPair

`func NewFmcDeviceHaPair(classId string, objectType string, ) *FmcDeviceHaPair`

NewFmcDeviceHaPair instantiates a new FmcDeviceHaPair object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFmcDeviceHaPairWithDefaults

`func NewFmcDeviceHaPairWithDefaults() *FmcDeviceHaPair`

NewFmcDeviceHaPairWithDefaults instantiates a new FmcDeviceHaPair object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *FmcDeviceHaPair) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *FmcDeviceHaPair) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *FmcDeviceHaPair) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *FmcDeviceHaPair) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *FmcDeviceHaPair) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *FmcDeviceHaPair) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDeviceId

`func (o *FmcDeviceHaPair) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *FmcDeviceHaPair) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *FmcDeviceHaPair) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *FmcDeviceHaPair) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### GetDomainId

`func (o *FmcDeviceHaPair) GetDomainId() string`

GetDomainId returns the DomainId field if non-nil, zero value otherwise.

### GetDomainIdOk

`func (o *FmcDeviceHaPair) GetDomainIdOk() (*string, bool)`

GetDomainIdOk returns a tuple with the DomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainId

`func (o *FmcDeviceHaPair) SetDomainId(v string)`

SetDomainId sets DomainId field to given value.

### HasDomainId

`func (o *FmcDeviceHaPair) HasDomainId() bool`

HasDomainId returns a boolean if a field has been set.

### GetName

`func (o *FmcDeviceHaPair) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FmcDeviceHaPair) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FmcDeviceHaPair) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FmcDeviceHaPair) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


