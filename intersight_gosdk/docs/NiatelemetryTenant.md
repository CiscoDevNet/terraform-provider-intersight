# NiatelemetryTenant

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "niatelemetry.Tenant"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "niatelemetry.Tenant"]
**BfdIfPolCount** | Pointer to **int64** | Number of Bidirectional Forwarding Detection bfdIfPol Model Objects. | [optional] 
**BfdIfpCount** | Pointer to **int64** | Number of objects with Bidirectional Forwarding Detection Interface Policy. | [optional] 
**DhcpRsProvCount** | Pointer to **int64** | Number of tenants with Dynamic Host Configuration Protocol enabled. | [optional] 
**Dn** | Pointer to **string** | Dn for the tenants present. | [optional] 
**FhsBdPolCount** | Pointer to **int64** | Number of objects with First hop security. Checks for presence of IP source gaurd, dynamic arp inspection. | [optional] 
**FvApCount** | Pointer to **int64** | Number of application profiles per tenant. | [optional] 
**FvBdCount** | Pointer to **int64** | Number of bridge domains with hardware proxy enabled per tenant. | [optional] 
**FvBdSubnetCount** | Pointer to **int64** | Number of bridge domain subnets per tenant. | [optional] 
**FvBdnoArpCount** | Pointer to **int64** | Number of bridge domains with ARP flodding disabled per tenant. | [optional] 
**FvCepCount** | Pointer to **int64** | Count of number of endpoints per tenant. | [optional] 
**FvRsBdToFhsCount** | Pointer to **int64** | Number of objects with First hop security. Checks for presence of IP source gaurd, dynamic arp inspection. | [optional] 
**FvRsBdToOutCount** | Pointer to **int64** | Number of bridge domains connected to layer 3 out per tenant. | [optional] 
**FvSiteConnpCount** | Pointer to **int64** | Number of Multi-Site objects. | [optional] 
**FvSubnetCount** | Pointer to **int64** | Number of subnets per tenant. | [optional] 
**IpStaticRouteCount** | Pointer to **int64** | Number of IP static routes per tenant. | [optional] 
**L3MulticastCount** | Pointer to **int64** | Number of layer 3 multicasts. | [optional] 
**L3MulticastCtxCount** | Pointer to **int64** | Number of layer 3 multicast CtxP. | [optional] 
**L3MulticastIfCount** | Pointer to **int64** | Number of layer 3 multicast IfP. | [optional] 
**L3outCount** | Pointer to **int64** | Number of L3 out objects for the tenants present. | [optional] 
**QosCustomPolCount** | Pointer to **int64** | Number of Quality Of Service Custom Policy. | [optional] 
**RecordType** | Pointer to **string** | Type of record DCNM / APIC / SE. This determines the type of platform where inventory was collected. | [optional] 
**RecordVersion** | Pointer to **string** | Version of record being pushed. This determines what was the API version for data available from the device. | [optional] 
**SiteName** | Pointer to **string** | The Site name represents an APIC cluster. Service Engine can onboard multiple APIC clusters / sites. | [optional] 
**Ssm** | Pointer to **string** | SSM property feature usage. | [optional] 
**SsmCount** | Pointer to **int64** | Number of context-level ssm translate policies per tenant. | [optional] 
**TcamOptCount** | Pointer to **int64** | Number of TCAM optimization enabled per tenant. | [optional] 
**TraceRouteEpCount** | Pointer to **int64** | Number of ITrace route endpoint per tenant. | [optional] 
**TraceRouteEpExtCount** | Pointer to **int64** | Number of ITrace endpoint external routes per tenant. | [optional] 
**TraceRouteExtEpCount** | Pointer to **int64** | Number of ITrace external endpoint routes per tenant. | [optional] 
**TraceRouteExtExtCount** | Pointer to **int64** | Number of ITrace external routes per tenant. | [optional] 
**VnsAbsGraphCount** | Pointer to **int64** | Number of objects with L4 to L7 Services graph. | [optional] 
**VnsBackupPolCount** | Pointer to **int64** | Number of objects with Policy Based Routing standby Node. Checks the Policy Based Routing backup policy. | [optional] 
**VnsRedirectDestCount** | Pointer to **int64** | Number of objects with Policy Based Routing standby Node. Policy based redirect requires a destination to redirect traffic. | [optional] 
**VnsSvcRedirectPolCount** | Pointer to **int64** | Number of Policy Based Routing and Policy Based Service Insertion objects. Includes without L4-L7 package. | [optional] 
**VrfCount** | Pointer to **int64** | Number of Vrfs per tenant. | [optional] 
**VzBrCpCount** | Pointer to **int64** | Number of Zoning Policy per tenant. | [optional] 
**VzRtConsCount** | Pointer to **int64** | Number of Client Contract between End Point Groups per tenant. | [optional] 
**VzRtProvCount** | Pointer to **int64** | Number of Client Contract between End Point Groups per tenant. | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewNiatelemetryTenant

`func NewNiatelemetryTenant(classId string, objectType string, ) *NiatelemetryTenant`

NewNiatelemetryTenant instantiates a new NiatelemetryTenant object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiatelemetryTenantWithDefaults

`func NewNiatelemetryTenantWithDefaults() *NiatelemetryTenant`

NewNiatelemetryTenantWithDefaults instantiates a new NiatelemetryTenant object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NiatelemetryTenant) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NiatelemetryTenant) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NiatelemetryTenant) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NiatelemetryTenant) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NiatelemetryTenant) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NiatelemetryTenant) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetBfdIfPolCount

`func (o *NiatelemetryTenant) GetBfdIfPolCount() int64`

GetBfdIfPolCount returns the BfdIfPolCount field if non-nil, zero value otherwise.

### GetBfdIfPolCountOk

`func (o *NiatelemetryTenant) GetBfdIfPolCountOk() (*int64, bool)`

GetBfdIfPolCountOk returns a tuple with the BfdIfPolCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBfdIfPolCount

`func (o *NiatelemetryTenant) SetBfdIfPolCount(v int64)`

SetBfdIfPolCount sets BfdIfPolCount field to given value.

### HasBfdIfPolCount

`func (o *NiatelemetryTenant) HasBfdIfPolCount() bool`

HasBfdIfPolCount returns a boolean if a field has been set.

### GetBfdIfpCount

`func (o *NiatelemetryTenant) GetBfdIfpCount() int64`

GetBfdIfpCount returns the BfdIfpCount field if non-nil, zero value otherwise.

### GetBfdIfpCountOk

`func (o *NiatelemetryTenant) GetBfdIfpCountOk() (*int64, bool)`

GetBfdIfpCountOk returns a tuple with the BfdIfpCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBfdIfpCount

`func (o *NiatelemetryTenant) SetBfdIfpCount(v int64)`

SetBfdIfpCount sets BfdIfpCount field to given value.

### HasBfdIfpCount

`func (o *NiatelemetryTenant) HasBfdIfpCount() bool`

HasBfdIfpCount returns a boolean if a field has been set.

### GetDhcpRsProvCount

`func (o *NiatelemetryTenant) GetDhcpRsProvCount() int64`

GetDhcpRsProvCount returns the DhcpRsProvCount field if non-nil, zero value otherwise.

### GetDhcpRsProvCountOk

`func (o *NiatelemetryTenant) GetDhcpRsProvCountOk() (*int64, bool)`

GetDhcpRsProvCountOk returns a tuple with the DhcpRsProvCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpRsProvCount

`func (o *NiatelemetryTenant) SetDhcpRsProvCount(v int64)`

SetDhcpRsProvCount sets DhcpRsProvCount field to given value.

### HasDhcpRsProvCount

`func (o *NiatelemetryTenant) HasDhcpRsProvCount() bool`

HasDhcpRsProvCount returns a boolean if a field has been set.

### GetDn

`func (o *NiatelemetryTenant) GetDn() string`

GetDn returns the Dn field if non-nil, zero value otherwise.

### GetDnOk

`func (o *NiatelemetryTenant) GetDnOk() (*string, bool)`

GetDnOk returns a tuple with the Dn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDn

`func (o *NiatelemetryTenant) SetDn(v string)`

SetDn sets Dn field to given value.

### HasDn

`func (o *NiatelemetryTenant) HasDn() bool`

HasDn returns a boolean if a field has been set.

### GetFhsBdPolCount

`func (o *NiatelemetryTenant) GetFhsBdPolCount() int64`

GetFhsBdPolCount returns the FhsBdPolCount field if non-nil, zero value otherwise.

### GetFhsBdPolCountOk

`func (o *NiatelemetryTenant) GetFhsBdPolCountOk() (*int64, bool)`

GetFhsBdPolCountOk returns a tuple with the FhsBdPolCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFhsBdPolCount

`func (o *NiatelemetryTenant) SetFhsBdPolCount(v int64)`

SetFhsBdPolCount sets FhsBdPolCount field to given value.

### HasFhsBdPolCount

`func (o *NiatelemetryTenant) HasFhsBdPolCount() bool`

HasFhsBdPolCount returns a boolean if a field has been set.

### GetFvApCount

`func (o *NiatelemetryTenant) GetFvApCount() int64`

GetFvApCount returns the FvApCount field if non-nil, zero value otherwise.

### GetFvApCountOk

`func (o *NiatelemetryTenant) GetFvApCountOk() (*int64, bool)`

GetFvApCountOk returns a tuple with the FvApCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFvApCount

`func (o *NiatelemetryTenant) SetFvApCount(v int64)`

SetFvApCount sets FvApCount field to given value.

### HasFvApCount

`func (o *NiatelemetryTenant) HasFvApCount() bool`

HasFvApCount returns a boolean if a field has been set.

### GetFvBdCount

`func (o *NiatelemetryTenant) GetFvBdCount() int64`

GetFvBdCount returns the FvBdCount field if non-nil, zero value otherwise.

### GetFvBdCountOk

`func (o *NiatelemetryTenant) GetFvBdCountOk() (*int64, bool)`

GetFvBdCountOk returns a tuple with the FvBdCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFvBdCount

`func (o *NiatelemetryTenant) SetFvBdCount(v int64)`

SetFvBdCount sets FvBdCount field to given value.

### HasFvBdCount

`func (o *NiatelemetryTenant) HasFvBdCount() bool`

HasFvBdCount returns a boolean if a field has been set.

### GetFvBdSubnetCount

`func (o *NiatelemetryTenant) GetFvBdSubnetCount() int64`

GetFvBdSubnetCount returns the FvBdSubnetCount field if non-nil, zero value otherwise.

### GetFvBdSubnetCountOk

`func (o *NiatelemetryTenant) GetFvBdSubnetCountOk() (*int64, bool)`

GetFvBdSubnetCountOk returns a tuple with the FvBdSubnetCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFvBdSubnetCount

`func (o *NiatelemetryTenant) SetFvBdSubnetCount(v int64)`

SetFvBdSubnetCount sets FvBdSubnetCount field to given value.

### HasFvBdSubnetCount

`func (o *NiatelemetryTenant) HasFvBdSubnetCount() bool`

HasFvBdSubnetCount returns a boolean if a field has been set.

### GetFvBdnoArpCount

`func (o *NiatelemetryTenant) GetFvBdnoArpCount() int64`

GetFvBdnoArpCount returns the FvBdnoArpCount field if non-nil, zero value otherwise.

### GetFvBdnoArpCountOk

`func (o *NiatelemetryTenant) GetFvBdnoArpCountOk() (*int64, bool)`

GetFvBdnoArpCountOk returns a tuple with the FvBdnoArpCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFvBdnoArpCount

`func (o *NiatelemetryTenant) SetFvBdnoArpCount(v int64)`

SetFvBdnoArpCount sets FvBdnoArpCount field to given value.

### HasFvBdnoArpCount

`func (o *NiatelemetryTenant) HasFvBdnoArpCount() bool`

HasFvBdnoArpCount returns a boolean if a field has been set.

### GetFvCepCount

`func (o *NiatelemetryTenant) GetFvCepCount() int64`

GetFvCepCount returns the FvCepCount field if non-nil, zero value otherwise.

### GetFvCepCountOk

`func (o *NiatelemetryTenant) GetFvCepCountOk() (*int64, bool)`

GetFvCepCountOk returns a tuple with the FvCepCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFvCepCount

`func (o *NiatelemetryTenant) SetFvCepCount(v int64)`

SetFvCepCount sets FvCepCount field to given value.

### HasFvCepCount

`func (o *NiatelemetryTenant) HasFvCepCount() bool`

HasFvCepCount returns a boolean if a field has been set.

### GetFvRsBdToFhsCount

`func (o *NiatelemetryTenant) GetFvRsBdToFhsCount() int64`

GetFvRsBdToFhsCount returns the FvRsBdToFhsCount field if non-nil, zero value otherwise.

### GetFvRsBdToFhsCountOk

`func (o *NiatelemetryTenant) GetFvRsBdToFhsCountOk() (*int64, bool)`

GetFvRsBdToFhsCountOk returns a tuple with the FvRsBdToFhsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFvRsBdToFhsCount

`func (o *NiatelemetryTenant) SetFvRsBdToFhsCount(v int64)`

SetFvRsBdToFhsCount sets FvRsBdToFhsCount field to given value.

### HasFvRsBdToFhsCount

`func (o *NiatelemetryTenant) HasFvRsBdToFhsCount() bool`

HasFvRsBdToFhsCount returns a boolean if a field has been set.

### GetFvRsBdToOutCount

`func (o *NiatelemetryTenant) GetFvRsBdToOutCount() int64`

GetFvRsBdToOutCount returns the FvRsBdToOutCount field if non-nil, zero value otherwise.

### GetFvRsBdToOutCountOk

`func (o *NiatelemetryTenant) GetFvRsBdToOutCountOk() (*int64, bool)`

GetFvRsBdToOutCountOk returns a tuple with the FvRsBdToOutCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFvRsBdToOutCount

`func (o *NiatelemetryTenant) SetFvRsBdToOutCount(v int64)`

SetFvRsBdToOutCount sets FvRsBdToOutCount field to given value.

### HasFvRsBdToOutCount

`func (o *NiatelemetryTenant) HasFvRsBdToOutCount() bool`

HasFvRsBdToOutCount returns a boolean if a field has been set.

### GetFvSiteConnpCount

`func (o *NiatelemetryTenant) GetFvSiteConnpCount() int64`

GetFvSiteConnpCount returns the FvSiteConnpCount field if non-nil, zero value otherwise.

### GetFvSiteConnpCountOk

`func (o *NiatelemetryTenant) GetFvSiteConnpCountOk() (*int64, bool)`

GetFvSiteConnpCountOk returns a tuple with the FvSiteConnpCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFvSiteConnpCount

`func (o *NiatelemetryTenant) SetFvSiteConnpCount(v int64)`

SetFvSiteConnpCount sets FvSiteConnpCount field to given value.

### HasFvSiteConnpCount

`func (o *NiatelemetryTenant) HasFvSiteConnpCount() bool`

HasFvSiteConnpCount returns a boolean if a field has been set.

### GetFvSubnetCount

`func (o *NiatelemetryTenant) GetFvSubnetCount() int64`

GetFvSubnetCount returns the FvSubnetCount field if non-nil, zero value otherwise.

### GetFvSubnetCountOk

`func (o *NiatelemetryTenant) GetFvSubnetCountOk() (*int64, bool)`

GetFvSubnetCountOk returns a tuple with the FvSubnetCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFvSubnetCount

`func (o *NiatelemetryTenant) SetFvSubnetCount(v int64)`

SetFvSubnetCount sets FvSubnetCount field to given value.

### HasFvSubnetCount

`func (o *NiatelemetryTenant) HasFvSubnetCount() bool`

HasFvSubnetCount returns a boolean if a field has been set.

### GetIpStaticRouteCount

`func (o *NiatelemetryTenant) GetIpStaticRouteCount() int64`

GetIpStaticRouteCount returns the IpStaticRouteCount field if non-nil, zero value otherwise.

### GetIpStaticRouteCountOk

`func (o *NiatelemetryTenant) GetIpStaticRouteCountOk() (*int64, bool)`

GetIpStaticRouteCountOk returns a tuple with the IpStaticRouteCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpStaticRouteCount

`func (o *NiatelemetryTenant) SetIpStaticRouteCount(v int64)`

SetIpStaticRouteCount sets IpStaticRouteCount field to given value.

### HasIpStaticRouteCount

`func (o *NiatelemetryTenant) HasIpStaticRouteCount() bool`

HasIpStaticRouteCount returns a boolean if a field has been set.

### GetL3MulticastCount

`func (o *NiatelemetryTenant) GetL3MulticastCount() int64`

GetL3MulticastCount returns the L3MulticastCount field if non-nil, zero value otherwise.

### GetL3MulticastCountOk

`func (o *NiatelemetryTenant) GetL3MulticastCountOk() (*int64, bool)`

GetL3MulticastCountOk returns a tuple with the L3MulticastCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetL3MulticastCount

`func (o *NiatelemetryTenant) SetL3MulticastCount(v int64)`

SetL3MulticastCount sets L3MulticastCount field to given value.

### HasL3MulticastCount

`func (o *NiatelemetryTenant) HasL3MulticastCount() bool`

HasL3MulticastCount returns a boolean if a field has been set.

### GetL3MulticastCtxCount

`func (o *NiatelemetryTenant) GetL3MulticastCtxCount() int64`

GetL3MulticastCtxCount returns the L3MulticastCtxCount field if non-nil, zero value otherwise.

### GetL3MulticastCtxCountOk

`func (o *NiatelemetryTenant) GetL3MulticastCtxCountOk() (*int64, bool)`

GetL3MulticastCtxCountOk returns a tuple with the L3MulticastCtxCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetL3MulticastCtxCount

`func (o *NiatelemetryTenant) SetL3MulticastCtxCount(v int64)`

SetL3MulticastCtxCount sets L3MulticastCtxCount field to given value.

### HasL3MulticastCtxCount

`func (o *NiatelemetryTenant) HasL3MulticastCtxCount() bool`

HasL3MulticastCtxCount returns a boolean if a field has been set.

### GetL3MulticastIfCount

`func (o *NiatelemetryTenant) GetL3MulticastIfCount() int64`

GetL3MulticastIfCount returns the L3MulticastIfCount field if non-nil, zero value otherwise.

### GetL3MulticastIfCountOk

`func (o *NiatelemetryTenant) GetL3MulticastIfCountOk() (*int64, bool)`

GetL3MulticastIfCountOk returns a tuple with the L3MulticastIfCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetL3MulticastIfCount

`func (o *NiatelemetryTenant) SetL3MulticastIfCount(v int64)`

SetL3MulticastIfCount sets L3MulticastIfCount field to given value.

### HasL3MulticastIfCount

`func (o *NiatelemetryTenant) HasL3MulticastIfCount() bool`

HasL3MulticastIfCount returns a boolean if a field has been set.

### GetL3outCount

`func (o *NiatelemetryTenant) GetL3outCount() int64`

GetL3outCount returns the L3outCount field if non-nil, zero value otherwise.

### GetL3outCountOk

`func (o *NiatelemetryTenant) GetL3outCountOk() (*int64, bool)`

GetL3outCountOk returns a tuple with the L3outCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetL3outCount

`func (o *NiatelemetryTenant) SetL3outCount(v int64)`

SetL3outCount sets L3outCount field to given value.

### HasL3outCount

`func (o *NiatelemetryTenant) HasL3outCount() bool`

HasL3outCount returns a boolean if a field has been set.

### GetQosCustomPolCount

`func (o *NiatelemetryTenant) GetQosCustomPolCount() int64`

GetQosCustomPolCount returns the QosCustomPolCount field if non-nil, zero value otherwise.

### GetQosCustomPolCountOk

`func (o *NiatelemetryTenant) GetQosCustomPolCountOk() (*int64, bool)`

GetQosCustomPolCountOk returns a tuple with the QosCustomPolCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQosCustomPolCount

`func (o *NiatelemetryTenant) SetQosCustomPolCount(v int64)`

SetQosCustomPolCount sets QosCustomPolCount field to given value.

### HasQosCustomPolCount

`func (o *NiatelemetryTenant) HasQosCustomPolCount() bool`

HasQosCustomPolCount returns a boolean if a field has been set.

### GetRecordType

`func (o *NiatelemetryTenant) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *NiatelemetryTenant) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *NiatelemetryTenant) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *NiatelemetryTenant) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetRecordVersion

`func (o *NiatelemetryTenant) GetRecordVersion() string`

GetRecordVersion returns the RecordVersion field if non-nil, zero value otherwise.

### GetRecordVersionOk

`func (o *NiatelemetryTenant) GetRecordVersionOk() (*string, bool)`

GetRecordVersionOk returns a tuple with the RecordVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordVersion

`func (o *NiatelemetryTenant) SetRecordVersion(v string)`

SetRecordVersion sets RecordVersion field to given value.

### HasRecordVersion

`func (o *NiatelemetryTenant) HasRecordVersion() bool`

HasRecordVersion returns a boolean if a field has been set.

### GetSiteName

`func (o *NiatelemetryTenant) GetSiteName() string`

GetSiteName returns the SiteName field if non-nil, zero value otherwise.

### GetSiteNameOk

`func (o *NiatelemetryTenant) GetSiteNameOk() (*string, bool)`

GetSiteNameOk returns a tuple with the SiteName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteName

`func (o *NiatelemetryTenant) SetSiteName(v string)`

SetSiteName sets SiteName field to given value.

### HasSiteName

`func (o *NiatelemetryTenant) HasSiteName() bool`

HasSiteName returns a boolean if a field has been set.

### GetSsm

`func (o *NiatelemetryTenant) GetSsm() string`

GetSsm returns the Ssm field if non-nil, zero value otherwise.

### GetSsmOk

`func (o *NiatelemetryTenant) GetSsmOk() (*string, bool)`

GetSsmOk returns a tuple with the Ssm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsm

`func (o *NiatelemetryTenant) SetSsm(v string)`

SetSsm sets Ssm field to given value.

### HasSsm

`func (o *NiatelemetryTenant) HasSsm() bool`

HasSsm returns a boolean if a field has been set.

### GetSsmCount

`func (o *NiatelemetryTenant) GetSsmCount() int64`

GetSsmCount returns the SsmCount field if non-nil, zero value otherwise.

### GetSsmCountOk

`func (o *NiatelemetryTenant) GetSsmCountOk() (*int64, bool)`

GetSsmCountOk returns a tuple with the SsmCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsmCount

`func (o *NiatelemetryTenant) SetSsmCount(v int64)`

SetSsmCount sets SsmCount field to given value.

### HasSsmCount

`func (o *NiatelemetryTenant) HasSsmCount() bool`

HasSsmCount returns a boolean if a field has been set.

### GetTcamOptCount

`func (o *NiatelemetryTenant) GetTcamOptCount() int64`

GetTcamOptCount returns the TcamOptCount field if non-nil, zero value otherwise.

### GetTcamOptCountOk

`func (o *NiatelemetryTenant) GetTcamOptCountOk() (*int64, bool)`

GetTcamOptCountOk returns a tuple with the TcamOptCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcamOptCount

`func (o *NiatelemetryTenant) SetTcamOptCount(v int64)`

SetTcamOptCount sets TcamOptCount field to given value.

### HasTcamOptCount

`func (o *NiatelemetryTenant) HasTcamOptCount() bool`

HasTcamOptCount returns a boolean if a field has been set.

### GetTraceRouteEpCount

`func (o *NiatelemetryTenant) GetTraceRouteEpCount() int64`

GetTraceRouteEpCount returns the TraceRouteEpCount field if non-nil, zero value otherwise.

### GetTraceRouteEpCountOk

`func (o *NiatelemetryTenant) GetTraceRouteEpCountOk() (*int64, bool)`

GetTraceRouteEpCountOk returns a tuple with the TraceRouteEpCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceRouteEpCount

`func (o *NiatelemetryTenant) SetTraceRouteEpCount(v int64)`

SetTraceRouteEpCount sets TraceRouteEpCount field to given value.

### HasTraceRouteEpCount

`func (o *NiatelemetryTenant) HasTraceRouteEpCount() bool`

HasTraceRouteEpCount returns a boolean if a field has been set.

### GetTraceRouteEpExtCount

`func (o *NiatelemetryTenant) GetTraceRouteEpExtCount() int64`

GetTraceRouteEpExtCount returns the TraceRouteEpExtCount field if non-nil, zero value otherwise.

### GetTraceRouteEpExtCountOk

`func (o *NiatelemetryTenant) GetTraceRouteEpExtCountOk() (*int64, bool)`

GetTraceRouteEpExtCountOk returns a tuple with the TraceRouteEpExtCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceRouteEpExtCount

`func (o *NiatelemetryTenant) SetTraceRouteEpExtCount(v int64)`

SetTraceRouteEpExtCount sets TraceRouteEpExtCount field to given value.

### HasTraceRouteEpExtCount

`func (o *NiatelemetryTenant) HasTraceRouteEpExtCount() bool`

HasTraceRouteEpExtCount returns a boolean if a field has been set.

### GetTraceRouteExtEpCount

`func (o *NiatelemetryTenant) GetTraceRouteExtEpCount() int64`

GetTraceRouteExtEpCount returns the TraceRouteExtEpCount field if non-nil, zero value otherwise.

### GetTraceRouteExtEpCountOk

`func (o *NiatelemetryTenant) GetTraceRouteExtEpCountOk() (*int64, bool)`

GetTraceRouteExtEpCountOk returns a tuple with the TraceRouteExtEpCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceRouteExtEpCount

`func (o *NiatelemetryTenant) SetTraceRouteExtEpCount(v int64)`

SetTraceRouteExtEpCount sets TraceRouteExtEpCount field to given value.

### HasTraceRouteExtEpCount

`func (o *NiatelemetryTenant) HasTraceRouteExtEpCount() bool`

HasTraceRouteExtEpCount returns a boolean if a field has been set.

### GetTraceRouteExtExtCount

`func (o *NiatelemetryTenant) GetTraceRouteExtExtCount() int64`

GetTraceRouteExtExtCount returns the TraceRouteExtExtCount field if non-nil, zero value otherwise.

### GetTraceRouteExtExtCountOk

`func (o *NiatelemetryTenant) GetTraceRouteExtExtCountOk() (*int64, bool)`

GetTraceRouteExtExtCountOk returns a tuple with the TraceRouteExtExtCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceRouteExtExtCount

`func (o *NiatelemetryTenant) SetTraceRouteExtExtCount(v int64)`

SetTraceRouteExtExtCount sets TraceRouteExtExtCount field to given value.

### HasTraceRouteExtExtCount

`func (o *NiatelemetryTenant) HasTraceRouteExtExtCount() bool`

HasTraceRouteExtExtCount returns a boolean if a field has been set.

### GetVnsAbsGraphCount

`func (o *NiatelemetryTenant) GetVnsAbsGraphCount() int64`

GetVnsAbsGraphCount returns the VnsAbsGraphCount field if non-nil, zero value otherwise.

### GetVnsAbsGraphCountOk

`func (o *NiatelemetryTenant) GetVnsAbsGraphCountOk() (*int64, bool)`

GetVnsAbsGraphCountOk returns a tuple with the VnsAbsGraphCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVnsAbsGraphCount

`func (o *NiatelemetryTenant) SetVnsAbsGraphCount(v int64)`

SetVnsAbsGraphCount sets VnsAbsGraphCount field to given value.

### HasVnsAbsGraphCount

`func (o *NiatelemetryTenant) HasVnsAbsGraphCount() bool`

HasVnsAbsGraphCount returns a boolean if a field has been set.

### GetVnsBackupPolCount

`func (o *NiatelemetryTenant) GetVnsBackupPolCount() int64`

GetVnsBackupPolCount returns the VnsBackupPolCount field if non-nil, zero value otherwise.

### GetVnsBackupPolCountOk

`func (o *NiatelemetryTenant) GetVnsBackupPolCountOk() (*int64, bool)`

GetVnsBackupPolCountOk returns a tuple with the VnsBackupPolCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVnsBackupPolCount

`func (o *NiatelemetryTenant) SetVnsBackupPolCount(v int64)`

SetVnsBackupPolCount sets VnsBackupPolCount field to given value.

### HasVnsBackupPolCount

`func (o *NiatelemetryTenant) HasVnsBackupPolCount() bool`

HasVnsBackupPolCount returns a boolean if a field has been set.

### GetVnsRedirectDestCount

`func (o *NiatelemetryTenant) GetVnsRedirectDestCount() int64`

GetVnsRedirectDestCount returns the VnsRedirectDestCount field if non-nil, zero value otherwise.

### GetVnsRedirectDestCountOk

`func (o *NiatelemetryTenant) GetVnsRedirectDestCountOk() (*int64, bool)`

GetVnsRedirectDestCountOk returns a tuple with the VnsRedirectDestCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVnsRedirectDestCount

`func (o *NiatelemetryTenant) SetVnsRedirectDestCount(v int64)`

SetVnsRedirectDestCount sets VnsRedirectDestCount field to given value.

### HasVnsRedirectDestCount

`func (o *NiatelemetryTenant) HasVnsRedirectDestCount() bool`

HasVnsRedirectDestCount returns a boolean if a field has been set.

### GetVnsSvcRedirectPolCount

`func (o *NiatelemetryTenant) GetVnsSvcRedirectPolCount() int64`

GetVnsSvcRedirectPolCount returns the VnsSvcRedirectPolCount field if non-nil, zero value otherwise.

### GetVnsSvcRedirectPolCountOk

`func (o *NiatelemetryTenant) GetVnsSvcRedirectPolCountOk() (*int64, bool)`

GetVnsSvcRedirectPolCountOk returns a tuple with the VnsSvcRedirectPolCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVnsSvcRedirectPolCount

`func (o *NiatelemetryTenant) SetVnsSvcRedirectPolCount(v int64)`

SetVnsSvcRedirectPolCount sets VnsSvcRedirectPolCount field to given value.

### HasVnsSvcRedirectPolCount

`func (o *NiatelemetryTenant) HasVnsSvcRedirectPolCount() bool`

HasVnsSvcRedirectPolCount returns a boolean if a field has been set.

### GetVrfCount

`func (o *NiatelemetryTenant) GetVrfCount() int64`

GetVrfCount returns the VrfCount field if non-nil, zero value otherwise.

### GetVrfCountOk

`func (o *NiatelemetryTenant) GetVrfCountOk() (*int64, bool)`

GetVrfCountOk returns a tuple with the VrfCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrfCount

`func (o *NiatelemetryTenant) SetVrfCount(v int64)`

SetVrfCount sets VrfCount field to given value.

### HasVrfCount

`func (o *NiatelemetryTenant) HasVrfCount() bool`

HasVrfCount returns a boolean if a field has been set.

### GetVzBrCpCount

`func (o *NiatelemetryTenant) GetVzBrCpCount() int64`

GetVzBrCpCount returns the VzBrCpCount field if non-nil, zero value otherwise.

### GetVzBrCpCountOk

`func (o *NiatelemetryTenant) GetVzBrCpCountOk() (*int64, bool)`

GetVzBrCpCountOk returns a tuple with the VzBrCpCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVzBrCpCount

`func (o *NiatelemetryTenant) SetVzBrCpCount(v int64)`

SetVzBrCpCount sets VzBrCpCount field to given value.

### HasVzBrCpCount

`func (o *NiatelemetryTenant) HasVzBrCpCount() bool`

HasVzBrCpCount returns a boolean if a field has been set.

### GetVzRtConsCount

`func (o *NiatelemetryTenant) GetVzRtConsCount() int64`

GetVzRtConsCount returns the VzRtConsCount field if non-nil, zero value otherwise.

### GetVzRtConsCountOk

`func (o *NiatelemetryTenant) GetVzRtConsCountOk() (*int64, bool)`

GetVzRtConsCountOk returns a tuple with the VzRtConsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVzRtConsCount

`func (o *NiatelemetryTenant) SetVzRtConsCount(v int64)`

SetVzRtConsCount sets VzRtConsCount field to given value.

### HasVzRtConsCount

`func (o *NiatelemetryTenant) HasVzRtConsCount() bool`

HasVzRtConsCount returns a boolean if a field has been set.

### GetVzRtProvCount

`func (o *NiatelemetryTenant) GetVzRtProvCount() int64`

GetVzRtProvCount returns the VzRtProvCount field if non-nil, zero value otherwise.

### GetVzRtProvCountOk

`func (o *NiatelemetryTenant) GetVzRtProvCountOk() (*int64, bool)`

GetVzRtProvCountOk returns a tuple with the VzRtProvCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVzRtProvCount

`func (o *NiatelemetryTenant) SetVzRtProvCount(v int64)`

SetVzRtProvCount sets VzRtProvCount field to given value.

### HasVzRtProvCount

`func (o *NiatelemetryTenant) HasVzRtProvCount() bool`

HasVzRtProvCount returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *NiatelemetryTenant) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *NiatelemetryTenant) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *NiatelemetryTenant) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *NiatelemetryTenant) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


