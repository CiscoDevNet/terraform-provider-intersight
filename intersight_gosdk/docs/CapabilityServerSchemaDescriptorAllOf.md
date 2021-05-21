# CapabilityServerSchemaDescriptorAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "capability.ServerSchemaDescriptor"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "capability.ServerSchemaDescriptor"]
**LocatorLedName** | Pointer to **string** | Redfish property name for the server. | [optional] 
**RedfishSchema** | Pointer to **string** | Redfish version information for the server. | [optional] 

## Methods

### NewCapabilityServerSchemaDescriptorAllOf

`func NewCapabilityServerSchemaDescriptorAllOf(classId string, objectType string, ) *CapabilityServerSchemaDescriptorAllOf`

NewCapabilityServerSchemaDescriptorAllOf instantiates a new CapabilityServerSchemaDescriptorAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCapabilityServerSchemaDescriptorAllOfWithDefaults

`func NewCapabilityServerSchemaDescriptorAllOfWithDefaults() *CapabilityServerSchemaDescriptorAllOf`

NewCapabilityServerSchemaDescriptorAllOfWithDefaults instantiates a new CapabilityServerSchemaDescriptorAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *CapabilityServerSchemaDescriptorAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *CapabilityServerSchemaDescriptorAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *CapabilityServerSchemaDescriptorAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *CapabilityServerSchemaDescriptorAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CapabilityServerSchemaDescriptorAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CapabilityServerSchemaDescriptorAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetLocatorLedName

`func (o *CapabilityServerSchemaDescriptorAllOf) GetLocatorLedName() string`

GetLocatorLedName returns the LocatorLedName field if non-nil, zero value otherwise.

### GetLocatorLedNameOk

`func (o *CapabilityServerSchemaDescriptorAllOf) GetLocatorLedNameOk() (*string, bool)`

GetLocatorLedNameOk returns a tuple with the LocatorLedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocatorLedName

`func (o *CapabilityServerSchemaDescriptorAllOf) SetLocatorLedName(v string)`

SetLocatorLedName sets LocatorLedName field to given value.

### HasLocatorLedName

`func (o *CapabilityServerSchemaDescriptorAllOf) HasLocatorLedName() bool`

HasLocatorLedName returns a boolean if a field has been set.

### GetRedfishSchema

`func (o *CapabilityServerSchemaDescriptorAllOf) GetRedfishSchema() string`

GetRedfishSchema returns the RedfishSchema field if non-nil, zero value otherwise.

### GetRedfishSchemaOk

`func (o *CapabilityServerSchemaDescriptorAllOf) GetRedfishSchemaOk() (*string, bool)`

GetRedfishSchemaOk returns a tuple with the RedfishSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedfishSchema

`func (o *CapabilityServerSchemaDescriptorAllOf) SetRedfishSchema(v string)`

SetRedfishSchema sets RedfishSchema field to given value.

### HasRedfishSchema

`func (o *CapabilityServerSchemaDescriptorAllOf) HasRedfishSchema() bool`

HasRedfishSchema returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


