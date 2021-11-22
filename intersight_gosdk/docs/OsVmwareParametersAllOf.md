# OsVmwareParametersAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "os.VmwareParameters"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "os.VmwareParameters"]
**Vlanid** | Pointer to **int64** | Specify the VLAN ID in which the ESXi host is turned on. Valid values ranges between 1 – 4095. | [optional] 

## Methods

### NewOsVmwareParametersAllOf

`func NewOsVmwareParametersAllOf(classId string, objectType string, ) *OsVmwareParametersAllOf`

NewOsVmwareParametersAllOf instantiates a new OsVmwareParametersAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOsVmwareParametersAllOfWithDefaults

`func NewOsVmwareParametersAllOfWithDefaults() *OsVmwareParametersAllOf`

NewOsVmwareParametersAllOfWithDefaults instantiates a new OsVmwareParametersAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *OsVmwareParametersAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *OsVmwareParametersAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *OsVmwareParametersAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *OsVmwareParametersAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *OsVmwareParametersAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *OsVmwareParametersAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetVlanid

`func (o *OsVmwareParametersAllOf) GetVlanid() int64`

GetVlanid returns the Vlanid field if non-nil, zero value otherwise.

### GetVlanidOk

`func (o *OsVmwareParametersAllOf) GetVlanidOk() (*int64, bool)`

GetVlanidOk returns a tuple with the Vlanid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanid

`func (o *OsVmwareParametersAllOf) SetVlanid(v int64)`

SetVlanid sets Vlanid field to given value.

### HasVlanid

`func (o *OsVmwareParametersAllOf) HasVlanid() bool`

HasVlanid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


