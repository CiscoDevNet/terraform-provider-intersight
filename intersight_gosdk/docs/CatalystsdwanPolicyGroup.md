# CatalystsdwanPolicyGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "catalystsdwan.PolicyGroup"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "catalystsdwan.PolicyGroup"]
**Devices** | Pointer to **[]string** |  | [optional] 
**Name** | Pointer to **string** | The Catalyst SDWAN policy group name. | [optional] 
**NumberOfDevices** | Pointer to **int64** | The Catalyst SDWAN policy group number of devices. | [optional] 
**NumberOfDevicesUpToDate** | Pointer to **int64** | The Catalyst SDWAN policy group number of devices up to date. | [optional] 
**PolicyGroupId** | Pointer to **string** | UUID for the Catalyst SDWAN policy group. | [optional] 
**Solution** | Pointer to **string** | The Catalyst SDWAN policy group solution. | [optional] 
**Version** | Pointer to **int64** | The Catalyst SDWAN policy group version. | [optional] 

## Methods

### NewCatalystsdwanPolicyGroup

`func NewCatalystsdwanPolicyGroup(classId string, objectType string, ) *CatalystsdwanPolicyGroup`

NewCatalystsdwanPolicyGroup instantiates a new CatalystsdwanPolicyGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalystsdwanPolicyGroupWithDefaults

`func NewCatalystsdwanPolicyGroupWithDefaults() *CatalystsdwanPolicyGroup`

NewCatalystsdwanPolicyGroupWithDefaults instantiates a new CatalystsdwanPolicyGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *CatalystsdwanPolicyGroup) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *CatalystsdwanPolicyGroup) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *CatalystsdwanPolicyGroup) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *CatalystsdwanPolicyGroup) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CatalystsdwanPolicyGroup) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CatalystsdwanPolicyGroup) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDevices

`func (o *CatalystsdwanPolicyGroup) GetDevices() []string`

GetDevices returns the Devices field if non-nil, zero value otherwise.

### GetDevicesOk

`func (o *CatalystsdwanPolicyGroup) GetDevicesOk() (*[]string, bool)`

GetDevicesOk returns a tuple with the Devices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevices

`func (o *CatalystsdwanPolicyGroup) SetDevices(v []string)`

SetDevices sets Devices field to given value.

### HasDevices

`func (o *CatalystsdwanPolicyGroup) HasDevices() bool`

HasDevices returns a boolean if a field has been set.

### SetDevicesNil

`func (o *CatalystsdwanPolicyGroup) SetDevicesNil(b bool)`

 SetDevicesNil sets the value for Devices to be an explicit nil

### UnsetDevices
`func (o *CatalystsdwanPolicyGroup) UnsetDevices()`

UnsetDevices ensures that no value is present for Devices, not even an explicit nil
### GetName

`func (o *CatalystsdwanPolicyGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CatalystsdwanPolicyGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CatalystsdwanPolicyGroup) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CatalystsdwanPolicyGroup) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNumberOfDevices

`func (o *CatalystsdwanPolicyGroup) GetNumberOfDevices() int64`

GetNumberOfDevices returns the NumberOfDevices field if non-nil, zero value otherwise.

### GetNumberOfDevicesOk

`func (o *CatalystsdwanPolicyGroup) GetNumberOfDevicesOk() (*int64, bool)`

GetNumberOfDevicesOk returns a tuple with the NumberOfDevices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfDevices

`func (o *CatalystsdwanPolicyGroup) SetNumberOfDevices(v int64)`

SetNumberOfDevices sets NumberOfDevices field to given value.

### HasNumberOfDevices

`func (o *CatalystsdwanPolicyGroup) HasNumberOfDevices() bool`

HasNumberOfDevices returns a boolean if a field has been set.

### GetNumberOfDevicesUpToDate

`func (o *CatalystsdwanPolicyGroup) GetNumberOfDevicesUpToDate() int64`

GetNumberOfDevicesUpToDate returns the NumberOfDevicesUpToDate field if non-nil, zero value otherwise.

### GetNumberOfDevicesUpToDateOk

`func (o *CatalystsdwanPolicyGroup) GetNumberOfDevicesUpToDateOk() (*int64, bool)`

GetNumberOfDevicesUpToDateOk returns a tuple with the NumberOfDevicesUpToDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfDevicesUpToDate

`func (o *CatalystsdwanPolicyGroup) SetNumberOfDevicesUpToDate(v int64)`

SetNumberOfDevicesUpToDate sets NumberOfDevicesUpToDate field to given value.

### HasNumberOfDevicesUpToDate

`func (o *CatalystsdwanPolicyGroup) HasNumberOfDevicesUpToDate() bool`

HasNumberOfDevicesUpToDate returns a boolean if a field has been set.

### GetPolicyGroupId

`func (o *CatalystsdwanPolicyGroup) GetPolicyGroupId() string`

GetPolicyGroupId returns the PolicyGroupId field if non-nil, zero value otherwise.

### GetPolicyGroupIdOk

`func (o *CatalystsdwanPolicyGroup) GetPolicyGroupIdOk() (*string, bool)`

GetPolicyGroupIdOk returns a tuple with the PolicyGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyGroupId

`func (o *CatalystsdwanPolicyGroup) SetPolicyGroupId(v string)`

SetPolicyGroupId sets PolicyGroupId field to given value.

### HasPolicyGroupId

`func (o *CatalystsdwanPolicyGroup) HasPolicyGroupId() bool`

HasPolicyGroupId returns a boolean if a field has been set.

### GetSolution

`func (o *CatalystsdwanPolicyGroup) GetSolution() string`

GetSolution returns the Solution field if non-nil, zero value otherwise.

### GetSolutionOk

`func (o *CatalystsdwanPolicyGroup) GetSolutionOk() (*string, bool)`

GetSolutionOk returns a tuple with the Solution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSolution

`func (o *CatalystsdwanPolicyGroup) SetSolution(v string)`

SetSolution sets Solution field to given value.

### HasSolution

`func (o *CatalystsdwanPolicyGroup) HasSolution() bool`

HasSolution returns a boolean if a field has been set.

### GetVersion

`func (o *CatalystsdwanPolicyGroup) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *CatalystsdwanPolicyGroup) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *CatalystsdwanPolicyGroup) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *CatalystsdwanPolicyGroup) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


