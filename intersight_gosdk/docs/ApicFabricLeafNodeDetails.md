# ApicFabricLeafNodeDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "apic.FabricLeafNodeDetails"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "apic.FabricLeafNodeDetails"]
**Id** | Pointer to **string** | Id of an leaf node within the Cisco Application Policy Infrastructure Controller (APIC). | [optional] 
**Name** | Pointer to **string** | Name of an leaf node within the Cisco Application Policy Infrastructure Controller (APIC). | [optional] 

## Methods

### NewApicFabricLeafNodeDetails

`func NewApicFabricLeafNodeDetails(classId string, objectType string, ) *ApicFabricLeafNodeDetails`

NewApicFabricLeafNodeDetails instantiates a new ApicFabricLeafNodeDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApicFabricLeafNodeDetailsWithDefaults

`func NewApicFabricLeafNodeDetailsWithDefaults() *ApicFabricLeafNodeDetails`

NewApicFabricLeafNodeDetailsWithDefaults instantiates a new ApicFabricLeafNodeDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ApicFabricLeafNodeDetails) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ApicFabricLeafNodeDetails) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ApicFabricLeafNodeDetails) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ApicFabricLeafNodeDetails) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ApicFabricLeafNodeDetails) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ApicFabricLeafNodeDetails) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetId

`func (o *ApicFabricLeafNodeDetails) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApicFabricLeafNodeDetails) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApicFabricLeafNodeDetails) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ApicFabricLeafNodeDetails) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ApicFabricLeafNodeDetails) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApicFabricLeafNodeDetails) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApicFabricLeafNodeDetails) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApicFabricLeafNodeDetails) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


