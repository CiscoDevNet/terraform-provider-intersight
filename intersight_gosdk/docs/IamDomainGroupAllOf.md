# IamDomainGroupAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the domain-group. | [optional] [readonly] 
**Partition1** | Pointer to **int64** | The partition number domain group related messages are produced for &#39;Partition1&#39; category service topics. | [optional] [readonly] 
**Partition2** | Pointer to **int64** | In a cloud environment this parameter will indicate to which partition number domain group related messages are produced for &#39;Partition2&#39; category service topics. | [optional] [readonly] 
**Partition3** | Pointer to **int64** | In a cloud environment this parameter will indicate to which partition number domain group related messages are produced for &#39;Partition3&#39; category service topics. | [optional] [readonly] 
**PartitionKey** | Pointer to **string** | Partition key used for producing messages to Kafka partitions. By default Domain-group id will be used as parition key. For Domain-groups belonging to Early adopters domain-group id will be prefixed with &#39;H&#39; and used as partition key, such partition key will be treated differently and messages will always be produced to partition 0. | [optional] [readonly] 
**Usage** | Pointer to **int64** | The number of devices in the domain-group. Device registration notifications are processed to update the usage of the domain-group. The on-boarding account will have multiple domain-groups, and during the device registration least used domain-group will be selected for the device. | [optional] [readonly] 
**Account** | Pointer to [**IamAccountRelationship**](iam.Account.Relationship.md) |  | [optional] 

## Methods

### NewIamDomainGroupAllOf

`func NewIamDomainGroupAllOf() *IamDomainGroupAllOf`

NewIamDomainGroupAllOf instantiates a new IamDomainGroupAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamDomainGroupAllOfWithDefaults

`func NewIamDomainGroupAllOfWithDefaults() *IamDomainGroupAllOf`

NewIamDomainGroupAllOfWithDefaults instantiates a new IamDomainGroupAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *IamDomainGroupAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IamDomainGroupAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IamDomainGroupAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IamDomainGroupAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPartition1

`func (o *IamDomainGroupAllOf) GetPartition1() int64`

GetPartition1 returns the Partition1 field if non-nil, zero value otherwise.

### GetPartition1Ok

`func (o *IamDomainGroupAllOf) GetPartition1Ok() (*int64, bool)`

GetPartition1Ok returns a tuple with the Partition1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartition1

`func (o *IamDomainGroupAllOf) SetPartition1(v int64)`

SetPartition1 sets Partition1 field to given value.

### HasPartition1

`func (o *IamDomainGroupAllOf) HasPartition1() bool`

HasPartition1 returns a boolean if a field has been set.

### GetPartition2

`func (o *IamDomainGroupAllOf) GetPartition2() int64`

GetPartition2 returns the Partition2 field if non-nil, zero value otherwise.

### GetPartition2Ok

`func (o *IamDomainGroupAllOf) GetPartition2Ok() (*int64, bool)`

GetPartition2Ok returns a tuple with the Partition2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartition2

`func (o *IamDomainGroupAllOf) SetPartition2(v int64)`

SetPartition2 sets Partition2 field to given value.

### HasPartition2

`func (o *IamDomainGroupAllOf) HasPartition2() bool`

HasPartition2 returns a boolean if a field has been set.

### GetPartition3

`func (o *IamDomainGroupAllOf) GetPartition3() int64`

GetPartition3 returns the Partition3 field if non-nil, zero value otherwise.

### GetPartition3Ok

`func (o *IamDomainGroupAllOf) GetPartition3Ok() (*int64, bool)`

GetPartition3Ok returns a tuple with the Partition3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartition3

`func (o *IamDomainGroupAllOf) SetPartition3(v int64)`

SetPartition3 sets Partition3 field to given value.

### HasPartition3

`func (o *IamDomainGroupAllOf) HasPartition3() bool`

HasPartition3 returns a boolean if a field has been set.

### GetPartitionKey

`func (o *IamDomainGroupAllOf) GetPartitionKey() string`

GetPartitionKey returns the PartitionKey field if non-nil, zero value otherwise.

### GetPartitionKeyOk

`func (o *IamDomainGroupAllOf) GetPartitionKeyOk() (*string, bool)`

GetPartitionKeyOk returns a tuple with the PartitionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitionKey

`func (o *IamDomainGroupAllOf) SetPartitionKey(v string)`

SetPartitionKey sets PartitionKey field to given value.

### HasPartitionKey

`func (o *IamDomainGroupAllOf) HasPartitionKey() bool`

HasPartitionKey returns a boolean if a field has been set.

### GetUsage

`func (o *IamDomainGroupAllOf) GetUsage() int64`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *IamDomainGroupAllOf) GetUsageOk() (*int64, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *IamDomainGroupAllOf) SetUsage(v int64)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *IamDomainGroupAllOf) HasUsage() bool`

HasUsage returns a boolean if a field has been set.

### GetAccount

`func (o *IamDomainGroupAllOf) GetAccount() IamAccountRelationship`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *IamDomainGroupAllOf) GetAccountOk() (*IamAccountRelationship, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *IamDomainGroupAllOf) SetAccount(v IamAccountRelationship)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *IamDomainGroupAllOf) HasAccount() bool`

HasAccount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


