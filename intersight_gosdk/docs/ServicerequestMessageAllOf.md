# ServicerequestMessageAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "servicerequest.Message"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "servicerequest.Message"]
**ServiceItemName** | Pointer to **string** | The service item in which operation is perfomed. | [optional] [readonly] 

## Methods

### NewServicerequestMessageAllOf

`func NewServicerequestMessageAllOf(classId string, objectType string, ) *ServicerequestMessageAllOf`

NewServicerequestMessageAllOf instantiates a new ServicerequestMessageAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServicerequestMessageAllOfWithDefaults

`func NewServicerequestMessageAllOfWithDefaults() *ServicerequestMessageAllOf`

NewServicerequestMessageAllOfWithDefaults instantiates a new ServicerequestMessageAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ServicerequestMessageAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ServicerequestMessageAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ServicerequestMessageAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ServicerequestMessageAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ServicerequestMessageAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ServicerequestMessageAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetServiceItemName

`func (o *ServicerequestMessageAllOf) GetServiceItemName() string`

GetServiceItemName returns the ServiceItemName field if non-nil, zero value otherwise.

### GetServiceItemNameOk

`func (o *ServicerequestMessageAllOf) GetServiceItemNameOk() (*string, bool)`

GetServiceItemNameOk returns a tuple with the ServiceItemName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceItemName

`func (o *ServicerequestMessageAllOf) SetServiceItemName(v string)`

SetServiceItemName sets ServiceItemName field to given value.

### HasServiceItemName

`func (o *ServicerequestMessageAllOf) HasServiceItemName() bool`

HasServiceItemName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


