# OsVmwareParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "os.VmwareParameters"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "os.VmwareParameters"]
**Vlanid** | Pointer to **int64** | Specify the VLAN ID in which the ESXi host is turned on. Valid values ranges between 1 – 4095. | [optional] 

## Methods

### NewOsVmwareParameters

`func NewOsVmwareParameters(classId string, objectType string, ) *OsVmwareParameters`

NewOsVmwareParameters instantiates a new OsVmwareParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOsVmwareParametersWithDefaults

`func NewOsVmwareParametersWithDefaults() *OsVmwareParameters`

NewOsVmwareParametersWithDefaults instantiates a new OsVmwareParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *OsVmwareParameters) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *OsVmwareParameters) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *OsVmwareParameters) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *OsVmwareParameters) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *OsVmwareParameters) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *OsVmwareParameters) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetVlanid

`func (o *OsVmwareParameters) GetVlanid() int64`

GetVlanid returns the Vlanid field if non-nil, zero value otherwise.

### GetVlanidOk

`func (o *OsVmwareParameters) GetVlanidOk() (*int64, bool)`

GetVlanidOk returns a tuple with the Vlanid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanid

`func (o *OsVmwareParameters) SetVlanid(v int64)`

SetVlanid sets Vlanid field to given value.

### HasVlanid

`func (o *OsVmwareParameters) HasVlanid() bool`

HasVlanid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


