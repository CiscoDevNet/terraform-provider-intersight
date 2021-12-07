# VirtualizationCiscoHypervisorManagerAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "virtualization.CiscoHypervisorManager"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "virtualization.CiscoHypervisorManager"]
**Account** | Pointer to [**IamAccountRelationship**](IamAccountRelationship.md) |  | [optional] 

## Methods

### NewVirtualizationCiscoHypervisorManagerAllOf

`func NewVirtualizationCiscoHypervisorManagerAllOf(classId string, objectType string, ) *VirtualizationCiscoHypervisorManagerAllOf`

NewVirtualizationCiscoHypervisorManagerAllOf instantiates a new VirtualizationCiscoHypervisorManagerAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualizationCiscoHypervisorManagerAllOfWithDefaults

`func NewVirtualizationCiscoHypervisorManagerAllOfWithDefaults() *VirtualizationCiscoHypervisorManagerAllOf`

NewVirtualizationCiscoHypervisorManagerAllOfWithDefaults instantiates a new VirtualizationCiscoHypervisorManagerAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *VirtualizationCiscoHypervisorManagerAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *VirtualizationCiscoHypervisorManagerAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *VirtualizationCiscoHypervisorManagerAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *VirtualizationCiscoHypervisorManagerAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VirtualizationCiscoHypervisorManagerAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VirtualizationCiscoHypervisorManagerAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAccount

`func (o *VirtualizationCiscoHypervisorManagerAllOf) GetAccount() IamAccountRelationship`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *VirtualizationCiscoHypervisorManagerAllOf) GetAccountOk() (*IamAccountRelationship, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *VirtualizationCiscoHypervisorManagerAllOf) SetAccount(v IamAccountRelationship)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *VirtualizationCiscoHypervisorManagerAllOf) HasAccount() bool`

HasAccount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


