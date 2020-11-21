# SdcardVirtualDriveAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**Enable** | Pointer to **bool** | Enable the respective virtual drive to be available to the host. | [optional] 

## Methods

### NewSdcardVirtualDriveAllOf

`func NewSdcardVirtualDriveAllOf(classId string, objectType string, ) *SdcardVirtualDriveAllOf`

NewSdcardVirtualDriveAllOf instantiates a new SdcardVirtualDriveAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSdcardVirtualDriveAllOfWithDefaults

`func NewSdcardVirtualDriveAllOfWithDefaults() *SdcardVirtualDriveAllOf`

NewSdcardVirtualDriveAllOfWithDefaults instantiates a new SdcardVirtualDriveAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *SdcardVirtualDriveAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *SdcardVirtualDriveAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *SdcardVirtualDriveAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *SdcardVirtualDriveAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *SdcardVirtualDriveAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *SdcardVirtualDriveAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetEnable

`func (o *SdcardVirtualDriveAllOf) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *SdcardVirtualDriveAllOf) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *SdcardVirtualDriveAllOf) SetEnable(v bool)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *SdcardVirtualDriveAllOf) HasEnable() bool`

HasEnable returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


