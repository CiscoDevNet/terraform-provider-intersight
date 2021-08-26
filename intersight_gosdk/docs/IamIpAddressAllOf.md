# IamIpAddressAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "iam.IpAddress"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "iam.IpAddress"]
**Address** | Pointer to **string** | The Trusted IP range&#39;s address. IP address, CIDR range, and IP address range formats are supported. For example &#39;12.13.14.15&#39;, &#39;12.13.14.0/24&#39;, and &#39;12.13.14.15-12.13.14.200&#39;. Reserved IP ranges &#39;127.0.0.1&#39;, &#39;10.0.0.0/8&#39;, &#39;172.16.0.0/12&#39;, and &#39;192.168.0.0/16&#39; are not allowed. | [optional] 
**Description** | Pointer to **string** | Description of Trusted IP address range. | [optional] 
**IpAccessManagement** | Pointer to [**IamIpAccessManagementRelationship**](IamIpAccessManagementRelationship.md) |  | [optional] 

## Methods

### NewIamIpAddressAllOf

`func NewIamIpAddressAllOf(classId string, objectType string, ) *IamIpAddressAllOf`

NewIamIpAddressAllOf instantiates a new IamIpAddressAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamIpAddressAllOfWithDefaults

`func NewIamIpAddressAllOfWithDefaults() *IamIpAddressAllOf`

NewIamIpAddressAllOfWithDefaults instantiates a new IamIpAddressAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *IamIpAddressAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *IamIpAddressAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *IamIpAddressAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *IamIpAddressAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *IamIpAddressAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *IamIpAddressAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAddress

`func (o *IamIpAddressAllOf) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *IamIpAddressAllOf) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *IamIpAddressAllOf) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *IamIpAddressAllOf) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetDescription

`func (o *IamIpAddressAllOf) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IamIpAddressAllOf) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IamIpAddressAllOf) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *IamIpAddressAllOf) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIpAccessManagement

`func (o *IamIpAddressAllOf) GetIpAccessManagement() IamIpAccessManagementRelationship`

GetIpAccessManagement returns the IpAccessManagement field if non-nil, zero value otherwise.

### GetIpAccessManagementOk

`func (o *IamIpAddressAllOf) GetIpAccessManagementOk() (*IamIpAccessManagementRelationship, bool)`

GetIpAccessManagementOk returns a tuple with the IpAccessManagement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAccessManagement

`func (o *IamIpAddressAllOf) SetIpAccessManagement(v IamIpAccessManagementRelationship)`

SetIpAccessManagement sets IpAccessManagement field to given value.

### HasIpAccessManagement

`func (o *IamIpAddressAllOf) HasIpAccessManagement() bool`

HasIpAccessManagement returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


