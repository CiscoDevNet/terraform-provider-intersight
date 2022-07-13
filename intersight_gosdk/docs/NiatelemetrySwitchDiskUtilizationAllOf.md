# NiatelemetrySwitchDiskUtilizationAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "niatelemetry.SwitchDiskUtilization"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "niatelemetry.SwitchDiskUtilization"]
**Avail** | Pointer to **string** | Available space in the given path. | [optional] 
**Name** | Pointer to **string** | Name of the path in the switch. | [optional] 
**Path** | Pointer to **string** | Path details for the name in the switch. | [optional] 
**Used** | Pointer to **string** | Used space in the given path. | [optional] 

## Methods

### NewNiatelemetrySwitchDiskUtilizationAllOf

`func NewNiatelemetrySwitchDiskUtilizationAllOf(classId string, objectType string, ) *NiatelemetrySwitchDiskUtilizationAllOf`

NewNiatelemetrySwitchDiskUtilizationAllOf instantiates a new NiatelemetrySwitchDiskUtilizationAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiatelemetrySwitchDiskUtilizationAllOfWithDefaults

`func NewNiatelemetrySwitchDiskUtilizationAllOfWithDefaults() *NiatelemetrySwitchDiskUtilizationAllOf`

NewNiatelemetrySwitchDiskUtilizationAllOfWithDefaults instantiates a new NiatelemetrySwitchDiskUtilizationAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NiatelemetrySwitchDiskUtilizationAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NiatelemetrySwitchDiskUtilizationAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NiatelemetrySwitchDiskUtilizationAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NiatelemetrySwitchDiskUtilizationAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NiatelemetrySwitchDiskUtilizationAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NiatelemetrySwitchDiskUtilizationAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAvail

`func (o *NiatelemetrySwitchDiskUtilizationAllOf) GetAvail() string`

GetAvail returns the Avail field if non-nil, zero value otherwise.

### GetAvailOk

`func (o *NiatelemetrySwitchDiskUtilizationAllOf) GetAvailOk() (*string, bool)`

GetAvailOk returns a tuple with the Avail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvail

`func (o *NiatelemetrySwitchDiskUtilizationAllOf) SetAvail(v string)`

SetAvail sets Avail field to given value.

### HasAvail

`func (o *NiatelemetrySwitchDiskUtilizationAllOf) HasAvail() bool`

HasAvail returns a boolean if a field has been set.

### GetName

`func (o *NiatelemetrySwitchDiskUtilizationAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NiatelemetrySwitchDiskUtilizationAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NiatelemetrySwitchDiskUtilizationAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NiatelemetrySwitchDiskUtilizationAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPath

`func (o *NiatelemetrySwitchDiskUtilizationAllOf) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *NiatelemetrySwitchDiskUtilizationAllOf) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *NiatelemetrySwitchDiskUtilizationAllOf) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *NiatelemetrySwitchDiskUtilizationAllOf) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetUsed

`func (o *NiatelemetrySwitchDiskUtilizationAllOf) GetUsed() string`

GetUsed returns the Used field if non-nil, zero value otherwise.

### GetUsedOk

`func (o *NiatelemetrySwitchDiskUtilizationAllOf) GetUsedOk() (*string, bool)`

GetUsedOk returns a tuple with the Used field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsed

`func (o *NiatelemetrySwitchDiskUtilizationAllOf) SetUsed(v string)`

SetUsed sets Used field to given value.

### HasUsed

`func (o *NiatelemetrySwitchDiskUtilizationAllOf) HasUsed() bool`

HasUsed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


