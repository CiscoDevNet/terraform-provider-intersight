# IamDomainGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "iam.DomainGroup"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "iam.DomainGroup"]
**Name** | Pointer to **string** | The name of the domain-group. | [optional] [readonly] 
**Partition1** | Pointer to **int64** | The partition number domain group related messages are produced for &#39;Partition1&#39; category service topics. | [optional] [readonly] 
**Partition2** | Pointer to **int64** | In a cloud environment this parameter will indicate to which partition number domain group related messages are produced for &#39;Partition2&#39; category service topics. | [optional] [readonly] 
**Partition3** | Pointer to **int64** | In a cloud environment this parameter will indicate to which partition number domain group related messages are produced for &#39;Partition3&#39; category service topics. | [optional] [readonly] 
**PartitionKey** | Pointer to **string** | Partition key used for producing messages to Kafka partitions. By default Domain-group id will be used as partition key. For Domain-groups belonging to Early adopters domain-group id will be prefixed with &#39;H&#39; and used as partition key, such partition key will be treated differently and messages will always be produced to partition 0. | [optional] [readonly] 
**Usage** | Pointer to **int64** | The number of devices in the domain-group. Device registration notifications are processed to update the usage of the domain-group. The on-boarding account will have multiple domain-groups, and during the device registration least used domain-group will be selected for the device. | [optional] [readonly] 
**Account** | Pointer to [**IamAccountRelationship**](IamAccountRelationship.md) |  | [optional] 

## Methods

### NewIamDomainGroup

`func NewIamDomainGroup(classId string, objectType string, ) *IamDomainGroup`

NewIamDomainGroup instantiates a new IamDomainGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamDomainGroupWithDefaults

`func NewIamDomainGroupWithDefaults() *IamDomainGroup`

NewIamDomainGroupWithDefaults instantiates a new IamDomainGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *IamDomainGroup) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *IamDomainGroup) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *IamDomainGroup) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *IamDomainGroup) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *IamDomainGroup) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *IamDomainGroup) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetName

`func (o *IamDomainGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IamDomainGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IamDomainGroup) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IamDomainGroup) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPartition1

`func (o *IamDomainGroup) GetPartition1() int64`

GetPartition1 returns the Partition1 field if non-nil, zero value otherwise.

### GetPartition1Ok

`func (o *IamDomainGroup) GetPartition1Ok() (*int64, bool)`

GetPartition1Ok returns a tuple with the Partition1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartition1

`func (o *IamDomainGroup) SetPartition1(v int64)`

SetPartition1 sets Partition1 field to given value.

### HasPartition1

`func (o *IamDomainGroup) HasPartition1() bool`

HasPartition1 returns a boolean if a field has been set.

### GetPartition2

`func (o *IamDomainGroup) GetPartition2() int64`

GetPartition2 returns the Partition2 field if non-nil, zero value otherwise.

### GetPartition2Ok

`func (o *IamDomainGroup) GetPartition2Ok() (*int64, bool)`

GetPartition2Ok returns a tuple with the Partition2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartition2

`func (o *IamDomainGroup) SetPartition2(v int64)`

SetPartition2 sets Partition2 field to given value.

### HasPartition2

`func (o *IamDomainGroup) HasPartition2() bool`

HasPartition2 returns a boolean if a field has been set.

### GetPartition3

`func (o *IamDomainGroup) GetPartition3() int64`

GetPartition3 returns the Partition3 field if non-nil, zero value otherwise.

### GetPartition3Ok

`func (o *IamDomainGroup) GetPartition3Ok() (*int64, bool)`

GetPartition3Ok returns a tuple with the Partition3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartition3

`func (o *IamDomainGroup) SetPartition3(v int64)`

SetPartition3 sets Partition3 field to given value.

### HasPartition3

`func (o *IamDomainGroup) HasPartition3() bool`

HasPartition3 returns a boolean if a field has been set.

### GetPartitionKey

`func (o *IamDomainGroup) GetPartitionKey() string`

GetPartitionKey returns the PartitionKey field if non-nil, zero value otherwise.

### GetPartitionKeyOk

`func (o *IamDomainGroup) GetPartitionKeyOk() (*string, bool)`

GetPartitionKeyOk returns a tuple with the PartitionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitionKey

`func (o *IamDomainGroup) SetPartitionKey(v string)`

SetPartitionKey sets PartitionKey field to given value.

### HasPartitionKey

`func (o *IamDomainGroup) HasPartitionKey() bool`

HasPartitionKey returns a boolean if a field has been set.

### GetUsage

`func (o *IamDomainGroup) GetUsage() int64`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *IamDomainGroup) GetUsageOk() (*int64, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *IamDomainGroup) SetUsage(v int64)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *IamDomainGroup) HasUsage() bool`

HasUsage returns a boolean if a field has been set.

### GetAccount

`func (o *IamDomainGroup) GetAccount() IamAccountRelationship`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *IamDomainGroup) GetAccountOk() (*IamAccountRelationship, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *IamDomainGroup) SetAccount(v IamAccountRelationship)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *IamDomainGroup) HasAccount() bool`

HasAccount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


