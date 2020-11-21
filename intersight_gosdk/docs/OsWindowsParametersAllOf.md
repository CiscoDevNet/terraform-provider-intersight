# OsWindowsParametersAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "os.WindowsParameters"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "os.WindowsParameters"]
**Edition** | Pointer to **string** | Lists all the editions supported for Windows Server installation. * &#x60;Standard&#x60; - Windows Standard Edition ideal for advanced features with limited virtualization. * &#x60;StandardCore&#x60; - Windows Standard Core Edition is a minimal installation option while installing Standard Core that is ideal for advanced features with limited virtualization. * &#x60;Datacenter&#x60; - Windows Standard Core Edition ideal for high requirements on IT workloads with largenumber fo virtual systems. * &#x60;DatacenterCore&#x60; - Windows Datacenter Core Edition is a minimal installation option while installing Datacenter Core that isideal for high requirements on IT workloads with largenumber for virtual systems. * &#x60;Core&#x60; - Microsoft Hyper-V is a native hypervisor to create and run virtual machines. | [optional] [default to "Standard"]

## Methods

### NewOsWindowsParametersAllOf

`func NewOsWindowsParametersAllOf(classId string, objectType string, ) *OsWindowsParametersAllOf`

NewOsWindowsParametersAllOf instantiates a new OsWindowsParametersAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOsWindowsParametersAllOfWithDefaults

`func NewOsWindowsParametersAllOfWithDefaults() *OsWindowsParametersAllOf`

NewOsWindowsParametersAllOfWithDefaults instantiates a new OsWindowsParametersAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *OsWindowsParametersAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *OsWindowsParametersAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *OsWindowsParametersAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *OsWindowsParametersAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *OsWindowsParametersAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *OsWindowsParametersAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetEdition

`func (o *OsWindowsParametersAllOf) GetEdition() string`

GetEdition returns the Edition field if non-nil, zero value otherwise.

### GetEditionOk

`func (o *OsWindowsParametersAllOf) GetEditionOk() (*string, bool)`

GetEditionOk returns a tuple with the Edition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdition

`func (o *OsWindowsParametersAllOf) SetEdition(v string)`

SetEdition sets Edition field to given value.

### HasEdition

`func (o *OsWindowsParametersAllOf) HasEdition() bool`

HasEdition returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


