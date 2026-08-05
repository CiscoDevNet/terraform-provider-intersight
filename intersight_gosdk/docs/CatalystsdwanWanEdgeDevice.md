# CatalystsdwanWanEdgeDevice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "catalystsdwan.WanEdgeDevice"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "catalystsdwan.WanEdgeDevice"]
**BfdSessions** | Pointer to **string** | Total number of BFD sessions from the WAN Edge Device to Cisco Catalyst SD-WAN Manager. | [optional] [readonly] 
**BfdSessionsUp** | Pointer to **string** | Number of BFD sessions in up state from the WAN Edge Device to Cisco Catalyst SD-WAN Manager. | [optional] [readonly] 
**ConfigDeviceStatus** | Pointer to **string** | Synchronization status of the WAN Edge Device configuration with Cisco Catalyst SD-WAN Manager. | [optional] [readonly] 
**ConfigLocked** | Pointer to **string** | Indicates if the device configuration is locked. | [optional] [readonly] 
**ControlConnections** | Pointer to **string** | Total number of control connections from the WAN Edge Device to Cisco Catalyst SD-WAN Manager. | [optional] [readonly] 
**ControlConnectionsUp** | Pointer to **string** | Number of control connections in up state from the WAN Edge Device to Cisco Catalyst SD-WAN Manager. | [optional] [readonly] 
**DeviceHealth** | Pointer to **string** | Health status of the WAN Edge Device. | [optional] [readonly] 
**DeviceId** | Pointer to **string** | Device ID of the WAN Edge Device. | [optional] [readonly] 
**DeviceState** | Pointer to **string** | Device state of the WAN Edge Device. | [optional] [readonly] 
**HostName** | Pointer to **string** | Host name of the WAN Edge Device. | [optional] [readonly] 
**ManagementPortHwAddr** | Pointer to **string** | MAC address of the WAN Edge Device management port. | [optional] [readonly] 
**ManagementPortIp** | Pointer to **string** | IP address of the WAN Edge Device management port. | [optional] [readonly] 
**ManagementPortSubnet** | Pointer to **string** | Subnet of the WAN Edge Device management port. | [optional] [readonly] 
**OrganizationName** | Pointer to **string** | The WAN Edge Device organization name. | [optional] [readonly] 
**Personality** | Pointer to **string** | Personality of the WAN Edge Device as seen in the Cisco Catalyst SD-WAN Manager (e.g., vEdge, cEdge). | [optional] [readonly] 
**Reachability** | Pointer to **string** | Reachability of the WAN Edge Device. * &#x60;Unknown&#x60; - Reachability status of the Catalyst SDWAN device from SDWAN Manager is unknown. * &#x60;Reachable&#x60; - The Catalyst SDWAN device is reachable from SDWAN Manager. * &#x60;Unreachable&#x60; - The Catalyst SDWAN device is unreachable from SDWAN Manager. * &#x60;Authentication Failed&#x60; - Authentication failed when SDWAN Manager attempted to connect to the Catalyst SDWAN device. | [optional] [readonly] [default to "Unknown"]
**SdwanUserTags** | Pointer to **string** | User defined tags associated with the WAN Edge Device from Cisco Catalyst SD-WAN Manager. | [optional] [readonly] 
**SiteId** | Pointer to **int64** | Site ID of the WAN Edge Device. | [optional] [readonly] 
**SiteName** | Pointer to **string** | Site name of the WAN Edge Device. | [optional] [readonly] 
**SpOrganizationName** | Pointer to **string** | The WAN Edge Device sp organization name. | [optional] [readonly] 
**SystemIp** | Pointer to **string** | System IP of the WAN Edge Device. | [optional] [readonly] 
**Tlocs** | Pointer to **string** | Total number of TLOCs from the WAN Edge Device to Cisco Catalyst SD-WAN Manager. | [optional] [readonly] 
**TlocsUp** | Pointer to **string** | Number of TLOCs in up state from the WAN Edge Device to Cisco Catalyst SD-WAN Manager. | [optional] [readonly] 
**Uptime** | Pointer to **string** | Uptime of the WAN Edge Device. | [optional] [readonly] 
**Validity** | Pointer to **string** | Validity of the WAN Edge Device. * &#x60;Unknown&#x60; - Validity of the Catalyst SDWAN device configuration is unknown. * &#x60;Valid&#x60; - Configuration on the Catalyst SDWAN device is valid. * &#x60;Invalid&#x60; - Configuration on the Catalyst SDWAN device is invalid. * &#x60;Prestaging&#x60; - Configuration on the Catalyst SDWAN device is in pre-staging state. * &#x60;Staging&#x60; - Configuration on the Catalyst SDWAN device is in staging state. * &#x60;Config Init&#x60; - Configuration on the Catalyst SDWAN device is in initialization state. | [optional] [readonly] [default to "Unknown"]
**RegisteredDevice** | Pointer to [**NullableAssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewCatalystsdwanWanEdgeDevice

`func NewCatalystsdwanWanEdgeDevice(classId string, objectType string, ) *CatalystsdwanWanEdgeDevice`

NewCatalystsdwanWanEdgeDevice instantiates a new CatalystsdwanWanEdgeDevice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalystsdwanWanEdgeDeviceWithDefaults

`func NewCatalystsdwanWanEdgeDeviceWithDefaults() *CatalystsdwanWanEdgeDevice`

NewCatalystsdwanWanEdgeDeviceWithDefaults instantiates a new CatalystsdwanWanEdgeDevice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *CatalystsdwanWanEdgeDevice) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *CatalystsdwanWanEdgeDevice) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *CatalystsdwanWanEdgeDevice) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *CatalystsdwanWanEdgeDevice) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CatalystsdwanWanEdgeDevice) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CatalystsdwanWanEdgeDevice) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetBfdSessions

`func (o *CatalystsdwanWanEdgeDevice) GetBfdSessions() string`

GetBfdSessions returns the BfdSessions field if non-nil, zero value otherwise.

### GetBfdSessionsOk

`func (o *CatalystsdwanWanEdgeDevice) GetBfdSessionsOk() (*string, bool)`

GetBfdSessionsOk returns a tuple with the BfdSessions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBfdSessions

`func (o *CatalystsdwanWanEdgeDevice) SetBfdSessions(v string)`

SetBfdSessions sets BfdSessions field to given value.

### HasBfdSessions

`func (o *CatalystsdwanWanEdgeDevice) HasBfdSessions() bool`

HasBfdSessions returns a boolean if a field has been set.

### GetBfdSessionsUp

`func (o *CatalystsdwanWanEdgeDevice) GetBfdSessionsUp() string`

GetBfdSessionsUp returns the BfdSessionsUp field if non-nil, zero value otherwise.

### GetBfdSessionsUpOk

`func (o *CatalystsdwanWanEdgeDevice) GetBfdSessionsUpOk() (*string, bool)`

GetBfdSessionsUpOk returns a tuple with the BfdSessionsUp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBfdSessionsUp

`func (o *CatalystsdwanWanEdgeDevice) SetBfdSessionsUp(v string)`

SetBfdSessionsUp sets BfdSessionsUp field to given value.

### HasBfdSessionsUp

`func (o *CatalystsdwanWanEdgeDevice) HasBfdSessionsUp() bool`

HasBfdSessionsUp returns a boolean if a field has been set.

### GetConfigDeviceStatus

`func (o *CatalystsdwanWanEdgeDevice) GetConfigDeviceStatus() string`

GetConfigDeviceStatus returns the ConfigDeviceStatus field if non-nil, zero value otherwise.

### GetConfigDeviceStatusOk

`func (o *CatalystsdwanWanEdgeDevice) GetConfigDeviceStatusOk() (*string, bool)`

GetConfigDeviceStatusOk returns a tuple with the ConfigDeviceStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigDeviceStatus

`func (o *CatalystsdwanWanEdgeDevice) SetConfigDeviceStatus(v string)`

SetConfigDeviceStatus sets ConfigDeviceStatus field to given value.

### HasConfigDeviceStatus

`func (o *CatalystsdwanWanEdgeDevice) HasConfigDeviceStatus() bool`

HasConfigDeviceStatus returns a boolean if a field has been set.

### GetConfigLocked

`func (o *CatalystsdwanWanEdgeDevice) GetConfigLocked() string`

GetConfigLocked returns the ConfigLocked field if non-nil, zero value otherwise.

### GetConfigLockedOk

`func (o *CatalystsdwanWanEdgeDevice) GetConfigLockedOk() (*string, bool)`

GetConfigLockedOk returns a tuple with the ConfigLocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigLocked

`func (o *CatalystsdwanWanEdgeDevice) SetConfigLocked(v string)`

SetConfigLocked sets ConfigLocked field to given value.

### HasConfigLocked

`func (o *CatalystsdwanWanEdgeDevice) HasConfigLocked() bool`

HasConfigLocked returns a boolean if a field has been set.

### GetControlConnections

`func (o *CatalystsdwanWanEdgeDevice) GetControlConnections() string`

GetControlConnections returns the ControlConnections field if non-nil, zero value otherwise.

### GetControlConnectionsOk

`func (o *CatalystsdwanWanEdgeDevice) GetControlConnectionsOk() (*string, bool)`

GetControlConnectionsOk returns a tuple with the ControlConnections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlConnections

`func (o *CatalystsdwanWanEdgeDevice) SetControlConnections(v string)`

SetControlConnections sets ControlConnections field to given value.

### HasControlConnections

`func (o *CatalystsdwanWanEdgeDevice) HasControlConnections() bool`

HasControlConnections returns a boolean if a field has been set.

### GetControlConnectionsUp

`func (o *CatalystsdwanWanEdgeDevice) GetControlConnectionsUp() string`

GetControlConnectionsUp returns the ControlConnectionsUp field if non-nil, zero value otherwise.

### GetControlConnectionsUpOk

`func (o *CatalystsdwanWanEdgeDevice) GetControlConnectionsUpOk() (*string, bool)`

GetControlConnectionsUpOk returns a tuple with the ControlConnectionsUp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlConnectionsUp

`func (o *CatalystsdwanWanEdgeDevice) SetControlConnectionsUp(v string)`

SetControlConnectionsUp sets ControlConnectionsUp field to given value.

### HasControlConnectionsUp

`func (o *CatalystsdwanWanEdgeDevice) HasControlConnectionsUp() bool`

HasControlConnectionsUp returns a boolean if a field has been set.

### GetDeviceHealth

`func (o *CatalystsdwanWanEdgeDevice) GetDeviceHealth() string`

GetDeviceHealth returns the DeviceHealth field if non-nil, zero value otherwise.

### GetDeviceHealthOk

`func (o *CatalystsdwanWanEdgeDevice) GetDeviceHealthOk() (*string, bool)`

GetDeviceHealthOk returns a tuple with the DeviceHealth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceHealth

`func (o *CatalystsdwanWanEdgeDevice) SetDeviceHealth(v string)`

SetDeviceHealth sets DeviceHealth field to given value.

### HasDeviceHealth

`func (o *CatalystsdwanWanEdgeDevice) HasDeviceHealth() bool`

HasDeviceHealth returns a boolean if a field has been set.

### GetDeviceId

`func (o *CatalystsdwanWanEdgeDevice) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *CatalystsdwanWanEdgeDevice) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *CatalystsdwanWanEdgeDevice) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *CatalystsdwanWanEdgeDevice) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### GetDeviceState

`func (o *CatalystsdwanWanEdgeDevice) GetDeviceState() string`

GetDeviceState returns the DeviceState field if non-nil, zero value otherwise.

### GetDeviceStateOk

`func (o *CatalystsdwanWanEdgeDevice) GetDeviceStateOk() (*string, bool)`

GetDeviceStateOk returns a tuple with the DeviceState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceState

`func (o *CatalystsdwanWanEdgeDevice) SetDeviceState(v string)`

SetDeviceState sets DeviceState field to given value.

### HasDeviceState

`func (o *CatalystsdwanWanEdgeDevice) HasDeviceState() bool`

HasDeviceState returns a boolean if a field has been set.

### GetHostName

`func (o *CatalystsdwanWanEdgeDevice) GetHostName() string`

GetHostName returns the HostName field if non-nil, zero value otherwise.

### GetHostNameOk

`func (o *CatalystsdwanWanEdgeDevice) GetHostNameOk() (*string, bool)`

GetHostNameOk returns a tuple with the HostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostName

`func (o *CatalystsdwanWanEdgeDevice) SetHostName(v string)`

SetHostName sets HostName field to given value.

### HasHostName

`func (o *CatalystsdwanWanEdgeDevice) HasHostName() bool`

HasHostName returns a boolean if a field has been set.

### GetManagementPortHwAddr

`func (o *CatalystsdwanWanEdgeDevice) GetManagementPortHwAddr() string`

GetManagementPortHwAddr returns the ManagementPortHwAddr field if non-nil, zero value otherwise.

### GetManagementPortHwAddrOk

`func (o *CatalystsdwanWanEdgeDevice) GetManagementPortHwAddrOk() (*string, bool)`

GetManagementPortHwAddrOk returns a tuple with the ManagementPortHwAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementPortHwAddr

`func (o *CatalystsdwanWanEdgeDevice) SetManagementPortHwAddr(v string)`

SetManagementPortHwAddr sets ManagementPortHwAddr field to given value.

### HasManagementPortHwAddr

`func (o *CatalystsdwanWanEdgeDevice) HasManagementPortHwAddr() bool`

HasManagementPortHwAddr returns a boolean if a field has been set.

### GetManagementPortIp

`func (o *CatalystsdwanWanEdgeDevice) GetManagementPortIp() string`

GetManagementPortIp returns the ManagementPortIp field if non-nil, zero value otherwise.

### GetManagementPortIpOk

`func (o *CatalystsdwanWanEdgeDevice) GetManagementPortIpOk() (*string, bool)`

GetManagementPortIpOk returns a tuple with the ManagementPortIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementPortIp

`func (o *CatalystsdwanWanEdgeDevice) SetManagementPortIp(v string)`

SetManagementPortIp sets ManagementPortIp field to given value.

### HasManagementPortIp

`func (o *CatalystsdwanWanEdgeDevice) HasManagementPortIp() bool`

HasManagementPortIp returns a boolean if a field has been set.

### GetManagementPortSubnet

`func (o *CatalystsdwanWanEdgeDevice) GetManagementPortSubnet() string`

GetManagementPortSubnet returns the ManagementPortSubnet field if non-nil, zero value otherwise.

### GetManagementPortSubnetOk

`func (o *CatalystsdwanWanEdgeDevice) GetManagementPortSubnetOk() (*string, bool)`

GetManagementPortSubnetOk returns a tuple with the ManagementPortSubnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementPortSubnet

`func (o *CatalystsdwanWanEdgeDevice) SetManagementPortSubnet(v string)`

SetManagementPortSubnet sets ManagementPortSubnet field to given value.

### HasManagementPortSubnet

`func (o *CatalystsdwanWanEdgeDevice) HasManagementPortSubnet() bool`

HasManagementPortSubnet returns a boolean if a field has been set.

### GetOrganizationName

`func (o *CatalystsdwanWanEdgeDevice) GetOrganizationName() string`

GetOrganizationName returns the OrganizationName field if non-nil, zero value otherwise.

### GetOrganizationNameOk

`func (o *CatalystsdwanWanEdgeDevice) GetOrganizationNameOk() (*string, bool)`

GetOrganizationNameOk returns a tuple with the OrganizationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationName

`func (o *CatalystsdwanWanEdgeDevice) SetOrganizationName(v string)`

SetOrganizationName sets OrganizationName field to given value.

### HasOrganizationName

`func (o *CatalystsdwanWanEdgeDevice) HasOrganizationName() bool`

HasOrganizationName returns a boolean if a field has been set.

### GetPersonality

`func (o *CatalystsdwanWanEdgeDevice) GetPersonality() string`

GetPersonality returns the Personality field if non-nil, zero value otherwise.

### GetPersonalityOk

`func (o *CatalystsdwanWanEdgeDevice) GetPersonalityOk() (*string, bool)`

GetPersonalityOk returns a tuple with the Personality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonality

`func (o *CatalystsdwanWanEdgeDevice) SetPersonality(v string)`

SetPersonality sets Personality field to given value.

### HasPersonality

`func (o *CatalystsdwanWanEdgeDevice) HasPersonality() bool`

HasPersonality returns a boolean if a field has been set.

### GetReachability

`func (o *CatalystsdwanWanEdgeDevice) GetReachability() string`

GetReachability returns the Reachability field if non-nil, zero value otherwise.

### GetReachabilityOk

`func (o *CatalystsdwanWanEdgeDevice) GetReachabilityOk() (*string, bool)`

GetReachabilityOk returns a tuple with the Reachability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReachability

`func (o *CatalystsdwanWanEdgeDevice) SetReachability(v string)`

SetReachability sets Reachability field to given value.

### HasReachability

`func (o *CatalystsdwanWanEdgeDevice) HasReachability() bool`

HasReachability returns a boolean if a field has been set.

### GetSdwanUserTags

`func (o *CatalystsdwanWanEdgeDevice) GetSdwanUserTags() string`

GetSdwanUserTags returns the SdwanUserTags field if non-nil, zero value otherwise.

### GetSdwanUserTagsOk

`func (o *CatalystsdwanWanEdgeDevice) GetSdwanUserTagsOk() (*string, bool)`

GetSdwanUserTagsOk returns a tuple with the SdwanUserTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSdwanUserTags

`func (o *CatalystsdwanWanEdgeDevice) SetSdwanUserTags(v string)`

SetSdwanUserTags sets SdwanUserTags field to given value.

### HasSdwanUserTags

`func (o *CatalystsdwanWanEdgeDevice) HasSdwanUserTags() bool`

HasSdwanUserTags returns a boolean if a field has been set.

### GetSiteId

`func (o *CatalystsdwanWanEdgeDevice) GetSiteId() int64`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *CatalystsdwanWanEdgeDevice) GetSiteIdOk() (*int64, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *CatalystsdwanWanEdgeDevice) SetSiteId(v int64)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *CatalystsdwanWanEdgeDevice) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetSiteName

`func (o *CatalystsdwanWanEdgeDevice) GetSiteName() string`

GetSiteName returns the SiteName field if non-nil, zero value otherwise.

### GetSiteNameOk

`func (o *CatalystsdwanWanEdgeDevice) GetSiteNameOk() (*string, bool)`

GetSiteNameOk returns a tuple with the SiteName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteName

`func (o *CatalystsdwanWanEdgeDevice) SetSiteName(v string)`

SetSiteName sets SiteName field to given value.

### HasSiteName

`func (o *CatalystsdwanWanEdgeDevice) HasSiteName() bool`

HasSiteName returns a boolean if a field has been set.

### GetSpOrganizationName

`func (o *CatalystsdwanWanEdgeDevice) GetSpOrganizationName() string`

GetSpOrganizationName returns the SpOrganizationName field if non-nil, zero value otherwise.

### GetSpOrganizationNameOk

`func (o *CatalystsdwanWanEdgeDevice) GetSpOrganizationNameOk() (*string, bool)`

GetSpOrganizationNameOk returns a tuple with the SpOrganizationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpOrganizationName

`func (o *CatalystsdwanWanEdgeDevice) SetSpOrganizationName(v string)`

SetSpOrganizationName sets SpOrganizationName field to given value.

### HasSpOrganizationName

`func (o *CatalystsdwanWanEdgeDevice) HasSpOrganizationName() bool`

HasSpOrganizationName returns a boolean if a field has been set.

### GetSystemIp

`func (o *CatalystsdwanWanEdgeDevice) GetSystemIp() string`

GetSystemIp returns the SystemIp field if non-nil, zero value otherwise.

### GetSystemIpOk

`func (o *CatalystsdwanWanEdgeDevice) GetSystemIpOk() (*string, bool)`

GetSystemIpOk returns a tuple with the SystemIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemIp

`func (o *CatalystsdwanWanEdgeDevice) SetSystemIp(v string)`

SetSystemIp sets SystemIp field to given value.

### HasSystemIp

`func (o *CatalystsdwanWanEdgeDevice) HasSystemIp() bool`

HasSystemIp returns a boolean if a field has been set.

### GetTlocs

`func (o *CatalystsdwanWanEdgeDevice) GetTlocs() string`

GetTlocs returns the Tlocs field if non-nil, zero value otherwise.

### GetTlocsOk

`func (o *CatalystsdwanWanEdgeDevice) GetTlocsOk() (*string, bool)`

GetTlocsOk returns a tuple with the Tlocs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlocs

`func (o *CatalystsdwanWanEdgeDevice) SetTlocs(v string)`

SetTlocs sets Tlocs field to given value.

### HasTlocs

`func (o *CatalystsdwanWanEdgeDevice) HasTlocs() bool`

HasTlocs returns a boolean if a field has been set.

### GetTlocsUp

`func (o *CatalystsdwanWanEdgeDevice) GetTlocsUp() string`

GetTlocsUp returns the TlocsUp field if non-nil, zero value otherwise.

### GetTlocsUpOk

`func (o *CatalystsdwanWanEdgeDevice) GetTlocsUpOk() (*string, bool)`

GetTlocsUpOk returns a tuple with the TlocsUp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlocsUp

`func (o *CatalystsdwanWanEdgeDevice) SetTlocsUp(v string)`

SetTlocsUp sets TlocsUp field to given value.

### HasTlocsUp

`func (o *CatalystsdwanWanEdgeDevice) HasTlocsUp() bool`

HasTlocsUp returns a boolean if a field has been set.

### GetUptime

`func (o *CatalystsdwanWanEdgeDevice) GetUptime() string`

GetUptime returns the Uptime field if non-nil, zero value otherwise.

### GetUptimeOk

`func (o *CatalystsdwanWanEdgeDevice) GetUptimeOk() (*string, bool)`

GetUptimeOk returns a tuple with the Uptime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUptime

`func (o *CatalystsdwanWanEdgeDevice) SetUptime(v string)`

SetUptime sets Uptime field to given value.

### HasUptime

`func (o *CatalystsdwanWanEdgeDevice) HasUptime() bool`

HasUptime returns a boolean if a field has been set.

### GetValidity

`func (o *CatalystsdwanWanEdgeDevice) GetValidity() string`

GetValidity returns the Validity field if non-nil, zero value otherwise.

### GetValidityOk

`func (o *CatalystsdwanWanEdgeDevice) GetValidityOk() (*string, bool)`

GetValidityOk returns a tuple with the Validity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidity

`func (o *CatalystsdwanWanEdgeDevice) SetValidity(v string)`

SetValidity sets Validity field to given value.

### HasValidity

`func (o *CatalystsdwanWanEdgeDevice) HasValidity() bool`

HasValidity returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *CatalystsdwanWanEdgeDevice) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *CatalystsdwanWanEdgeDevice) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *CatalystsdwanWanEdgeDevice) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *CatalystsdwanWanEdgeDevice) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### SetRegisteredDeviceNil

`func (o *CatalystsdwanWanEdgeDevice) SetRegisteredDeviceNil(b bool)`

 SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil

### UnsetRegisteredDevice
`func (o *CatalystsdwanWanEdgeDevice) UnsetRegisteredDevice()`

UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


