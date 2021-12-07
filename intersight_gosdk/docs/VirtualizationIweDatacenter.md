# VirtualizationIweDatacenter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "virtualization.IweDatacenter"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "virtualization.IweDatacenter"]
**Account** | Pointer to [**IamAccountRelationship**](IamAccountRelationship.md) |  | [optional] 
**HypervisorManager** | Pointer to [**VirtualizationCiscoHypervisorManagerRelationship**](VirtualizationCiscoHypervisorManagerRelationship.md) |  | [optional] 

## Methods

### NewVirtualizationIweDatacenter

`func NewVirtualizationIweDatacenter(classId string, objectType string, ) *VirtualizationIweDatacenter`

NewVirtualizationIweDatacenter instantiates a new VirtualizationIweDatacenter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualizationIweDatacenterWithDefaults

`func NewVirtualizationIweDatacenterWithDefaults() *VirtualizationIweDatacenter`

NewVirtualizationIweDatacenterWithDefaults instantiates a new VirtualizationIweDatacenter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *VirtualizationIweDatacenter) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *VirtualizationIweDatacenter) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *VirtualizationIweDatacenter) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *VirtualizationIweDatacenter) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VirtualizationIweDatacenter) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VirtualizationIweDatacenter) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAccount

`func (o *VirtualizationIweDatacenter) GetAccount() IamAccountRelationship`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *VirtualizationIweDatacenter) GetAccountOk() (*IamAccountRelationship, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *VirtualizationIweDatacenter) SetAccount(v IamAccountRelationship)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *VirtualizationIweDatacenter) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetHypervisorManager

`func (o *VirtualizationIweDatacenter) GetHypervisorManager() VirtualizationCiscoHypervisorManagerRelationship`

GetHypervisorManager returns the HypervisorManager field if non-nil, zero value otherwise.

### GetHypervisorManagerOk

`func (o *VirtualizationIweDatacenter) GetHypervisorManagerOk() (*VirtualizationCiscoHypervisorManagerRelationship, bool)`

GetHypervisorManagerOk returns a tuple with the HypervisorManager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHypervisorManager

`func (o *VirtualizationIweDatacenter) SetHypervisorManager(v VirtualizationCiscoHypervisorManagerRelationship)`

SetHypervisorManager sets HypervisorManager field to given value.

### HasHypervisorManager

`func (o *VirtualizationIweDatacenter) HasHypervisorManager() bool`

HasHypervisorManager returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


