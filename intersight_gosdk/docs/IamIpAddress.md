# IamIpAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "iam.IpAddress"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "iam.IpAddress"]
**Address** | Pointer to **string** | The Trusted IP range&#39;s address. IP address, CIDR range, and IP address range formats are supported. For example &#39;12.13.14.15&#39;, &#39;12.13.14.0/24&#39;, and &#39;12.13.14.15-12.13.14.200&#39;. Reserved IP ranges &#39;127.0.0.1&#39;, &#39;10.0.0.0/8&#39;, &#39;172.16.0.0/12&#39;, and &#39;192.168.0.0/16&#39; are not allowed. | [optional] 
**Description** | Pointer to **string** | Description of Trusted IP address range. | [optional] 
**IpAccessManagement** | Pointer to [**IamIpAccessManagementRelationship**](IamIpAccessManagementRelationship.md) |  | [optional] 

## Methods

### NewIamIpAddress

`func NewIamIpAddress(classId string, objectType string, ) *IamIpAddress`

NewIamIpAddress instantiates a new IamIpAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamIpAddressWithDefaults

`func NewIamIpAddressWithDefaults() *IamIpAddress`

NewIamIpAddressWithDefaults instantiates a new IamIpAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *IamIpAddress) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *IamIpAddress) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *IamIpAddress) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *IamIpAddress) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *IamIpAddress) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *IamIpAddress) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAddress

`func (o *IamIpAddress) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *IamIpAddress) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *IamIpAddress) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *IamIpAddress) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetDescription

`func (o *IamIpAddress) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IamIpAddress) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IamIpAddress) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *IamIpAddress) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIpAccessManagement

`func (o *IamIpAddress) GetIpAccessManagement() IamIpAccessManagementRelationship`

GetIpAccessManagement returns the IpAccessManagement field if non-nil, zero value otherwise.

### GetIpAccessManagementOk

`func (o *IamIpAddress) GetIpAccessManagementOk() (*IamIpAccessManagementRelationship, bool)`

GetIpAccessManagementOk returns a tuple with the IpAccessManagement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAccessManagement

`func (o *IamIpAddress) SetIpAccessManagement(v IamIpAccessManagementRelationship)`

SetIpAccessManagement sets IpAccessManagement field to given value.

### HasIpAccessManagement

`func (o *IamIpAddress) HasIpAccessManagement() bool`

HasIpAccessManagement returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


