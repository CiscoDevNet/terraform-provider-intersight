# HclServerCatalog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hcl.ServerCatalog"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hcl.ServerCatalog"]
**ProcessorFamily** | Pointer to **[]string** |  | [optional] 
**ServerPid** | Pointer to **string** | Three part ID representing the server model as returned by UCSM/CIMC XML APIs. | [optional] 

## Methods

### NewHclServerCatalog

`func NewHclServerCatalog(classId string, objectType string, ) *HclServerCatalog`

NewHclServerCatalog instantiates a new HclServerCatalog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHclServerCatalogWithDefaults

`func NewHclServerCatalogWithDefaults() *HclServerCatalog`

NewHclServerCatalogWithDefaults instantiates a new HclServerCatalog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HclServerCatalog) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HclServerCatalog) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HclServerCatalog) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HclServerCatalog) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HclServerCatalog) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HclServerCatalog) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetProcessorFamily

`func (o *HclServerCatalog) GetProcessorFamily() []string`

GetProcessorFamily returns the ProcessorFamily field if non-nil, zero value otherwise.

### GetProcessorFamilyOk

`func (o *HclServerCatalog) GetProcessorFamilyOk() (*[]string, bool)`

GetProcessorFamilyOk returns a tuple with the ProcessorFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessorFamily

`func (o *HclServerCatalog) SetProcessorFamily(v []string)`

SetProcessorFamily sets ProcessorFamily field to given value.

### HasProcessorFamily

`func (o *HclServerCatalog) HasProcessorFamily() bool`

HasProcessorFamily returns a boolean if a field has been set.

### SetProcessorFamilyNil

`func (o *HclServerCatalog) SetProcessorFamilyNil(b bool)`

 SetProcessorFamilyNil sets the value for ProcessorFamily to be an explicit nil

### UnsetProcessorFamily
`func (o *HclServerCatalog) UnsetProcessorFamily()`

UnsetProcessorFamily ensures that no value is present for ProcessorFamily, not even an explicit nil
### GetServerPid

`func (o *HclServerCatalog) GetServerPid() string`

GetServerPid returns the ServerPid field if non-nil, zero value otherwise.

### GetServerPidOk

`func (o *HclServerCatalog) GetServerPidOk() (*string, bool)`

GetServerPidOk returns a tuple with the ServerPid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerPid

`func (o *HclServerCatalog) SetServerPid(v string)`

SetServerPid sets ServerPid field to given value.

### HasServerPid

`func (o *HclServerCatalog) HasServerPid() bool`

HasServerPid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


