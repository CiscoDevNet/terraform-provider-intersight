package intersight

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"strings"
	"time"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func getHciNodeSchema() map[string]*schema.Schema {
	var schemaMap = make(map[string]*schema.Schema)
	schemaMap = map[string]*schema.Schema{"account_moid": {
		Description: "The Account ID for this managed object.",
		Type:        schema.TypeString,
		Optional:    true,
	},
		"additional_properties": {
			Type:             schema.TypeString,
			Optional:         true,
			DiffSuppressFunc: SuppressDiffAdditionProps,
		},
		"ancestors": {
			Description: "An array of relationships to moBaseMo resources.",
			Type:        schema.TypeList,
			Optional:    true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"additional_properties": {
						Type:             schema.TypeString,
						Optional:         true,
						DiffSuppressFunc: SuppressDiffAdditionProps,
					},
					"class_id": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"moid": {
						Description: "The Moid of the referenced REST resource.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the remote type referred by this relationship.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"selector": {
						Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				},
			},
		},
		"block_model": {
			Description: "The rackable unit model of the node.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"block_serial": {
			Description: "The rackable unit serial number of the node.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"boot_time_usecs": {
			Description: "The boot time in microseconds of the node.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"class_id": {
			Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"cluster": {
			Description: "A reference to a hciCluster resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
			Type:        schema.TypeList,
			MaxItems:    1,
			Optional:    true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"additional_properties": {
						Type:             schema.TypeString,
						Optional:         true,
						DiffSuppressFunc: SuppressDiffAdditionProps,
					},
					"class_id": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"moid": {
						Description: "The Moid of the referenced REST resource.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the remote type referred by this relationship.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"selector": {
						Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				},
			},
		},
		"cluster_ext_id": {
			Description: "The unique identifier of the cluster.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"cluster_name": {
			Description: "The name of the cluster this node belongs to.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"controller_vm_backplane_address": {
			Description: "The backplane IP address of the controller VM.",
			Type:        schema.TypeList,
			MaxItems:    1,
			Optional:    true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"additional_properties": {
						Type:             schema.TypeString,
						Optional:         true,
						DiffSuppressFunc: SuppressDiffAdditionProps,
					},
					"class_id": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"ipv4_address": {
						Description: "An IPv4 address in this IpAddress.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"ipv4_prefix_length": {
						Description: "The prefix length of the IPv4 address.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"ipv6_address": {
						Description: "An IPv6 address in this IpAddress.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"ipv6_prefix_length": {
						Description: "The prefix length of the IPv6 address.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				},
			},
		},
		"controller_vm_external_address": {
			Description: "The external IP address of the controller VM.",
			Type:        schema.TypeList,
			MaxItems:    1,
			Optional:    true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"additional_properties": {
						Type:             schema.TypeString,
						Optional:         true,
						DiffSuppressFunc: SuppressDiffAdditionProps,
					},
					"class_id": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"ipv4_address": {
						Description: "An IPv4 address in this IpAddress.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"ipv4_prefix_length": {
						Description: "The prefix length of the IPv4 address.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"ipv6_address": {
						Description: "An IPv6 address in this IpAddress.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"ipv6_prefix_length": {
						Description: "The prefix length of the IPv6 address.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				},
			},
		},
		"controller_vm_id": {
			Description: "The identifier number of the controller VM.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"controller_vm_maintanence_mode": {
			Description: "The maintenance mode status of the controller VM.",
			Type:        schema.TypeBool,
			Optional:    true,
		},
		"controller_vm_nat_ip": {
			Description: "The NAT IP address of the controller VM.",
			Type:        schema.TypeList,
			MaxItems:    1,
			Optional:    true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"additional_properties": {
						Type:             schema.TypeString,
						Optional:         true,
						DiffSuppressFunc: SuppressDiffAdditionProps,
					},
					"class_id": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"ipv4_address": {
						Description: "An IPv4 address in this IpAddress.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"ipv4_prefix_length": {
						Description: "The prefix length of the IPv4 address.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"ipv6_address": {
						Description: "An IPv6 address in this IpAddress.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"ipv6_prefix_length": {
						Description: "The prefix length of the IPv6 address.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				},
			},
		},
		"controller_vm_nat_port": {
			Description: "The NAT port of the controller VM.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"controller_vm_rdma_backplane_address": {
			Description: "The RDMA backplane IP address of the controller VM.",
			Type:        schema.TypeList,
			MaxItems:    1,
			Optional:    true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"additional_properties": {
						Type:             schema.TypeString,
						Optional:         true,
						DiffSuppressFunc: SuppressDiffAdditionProps,
					},
					"class_id": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"ipv4_address": {
						Description: "An IPv4 address in this IpAddress.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"ipv4_prefix_length": {
						Description: "The prefix length of the IPv4 address.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"ipv6_address": {
						Description: "An IPv6 address in this IpAddress.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"ipv6_prefix_length": {
						Description: "The prefix length of the IPv6 address.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				},
			},
		},
		"controller_vm_server_uuid": {
			Description: "The Rackable unit UUID of the server.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"cpu_capacity_hz": {
			Description: "The CPU capacity in Hz of the node.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"cpu_frequency_hz": {
			Description: "The CPU frequency in Hz on the node.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"cpu_model": {
			Description: "The CPU model of the node.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"cpu_usage_hz": {
			Description: "The CPU usage in Hz of the node.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"create_time": {
			Description: "The time when this managed object was created.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"default_vhd_container_uuid": {
			Description: "The default VHD container UUID of the node.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"default_vhd_location": {
			Description: "The default VHD location of the node.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"default_vm_container_uuid": {
			Description: "The default VM container UUID of the node.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"default_vm_location": {
			Description: "The default VM location of the node.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"disk_count": {
			Description: "The number of disks on the node.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"disks": {
			Description: "An array of relationships to hciDisk resources.",
			Type:        schema.TypeList,
			Optional:    true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"additional_properties": {
						Type:             schema.TypeString,
						Optional:         true,
						DiffSuppressFunc: SuppressDiffAdditionProps,
					},
					"class_id": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"moid": {
						Description: "The Moid of the referenced REST resource.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the remote type referred by this relationship.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"selector": {
						Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				},
			},
		},
		"domain_group_moid": {
			Description: "The DomainGroup ID for this managed object.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"failover_cluster_fqdn": {
			Description: "The failover cluster FQDN of the node.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"failover_cluster_node_status": {
			Description: "The failover cluster node status of the node.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"gpu_count": {
			Description: "The number of GPUs on the node.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"gpu_driver_version": {
			Description: "The GPU driver version of the node.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"gpus": {
			Description: "An array of relationships to hciGpu resources.",
			Type:        schema.TypeList,
			Optional:    true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"additional_properties": {
						Type:             schema.TypeString,
						Optional:         true,
						DiffSuppressFunc: SuppressDiffAdditionProps,
					},
					"class_id": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"moid": {
						Description: "The Moid of the referenced REST resource.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the remote type referred by this relationship.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"selector": {
						Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				},
			},
		},
		"has_csr": {
			Description: "Certificate signing request status of the node.",
			Type:        schema.TypeBool,
			Optional:    true,
		},
		"host_name": {
			Description: "The name of the host the node is running on.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"host_type": {
			Description: "The type of the host, e.g. HYPER_CONVERGED, COMPUTE_ONLY, STORAGE_ONLY.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"hypervisor_acropolis_connection_state": {
			Description: "The connection state of the hypervisor, e.g. CONNECTED, DISCONNECTED, NOT_AVAILABLE.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"hypervisor_external_address": {
			Description: "The external IP address of the hypervisor on this node.",
			Type:        schema.TypeList,
			MaxItems:    1,
			Optional:    true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"additional_properties": {
						Type:             schema.TypeString,
						Optional:         true,
						DiffSuppressFunc: SuppressDiffAdditionProps,
					},
					"class_id": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"ipv4_address": {
						Description: "An IPv4 address in this IpAddress.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"ipv4_prefix_length": {
						Description: "The prefix length of the IPv4 address.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"ipv6_address": {
						Description: "An IPv6 address in this IpAddress.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"ipv6_prefix_length": {
						Description: "The prefix length of the IPv6 address.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				},
			},
		},
		"hypervisor_number_of_vms": {
			Description: "The number of VMs managed on this node.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"hypervisor_state": {
			Description: "The hypervisor state e.g. ACROPOLIS_NORMAL, ENTERING_MAINTENANCE_MODE,\nENTERED_MAINTENANCE_MODE, RESERVED_FOR_HA_FAILOVER, ENTERING_MAINTENANCE_MODE_FROM_HA_FAILOVER,\nRESERVING_FOR_HA_FAILOVER, HA_FAILOVER_SOURCE, HA_FAILOVER_TARGET, HA_HEALING_SOURCE,\nHA_HEALING_TARGET.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"hypervisor_type": {
			Description: "The hypervisor type, e.g. AHV, ESX, HYPERV, XEN, NATIVEHOST etc.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"hypervisor_user_name": {
			Description: "The user name of the hypervisor on this node.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"hypervisor_version": {
			Description: "The version of the hypervisor on this node.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"ipmi_ip": {
			Description: "The IPMI IP address of the controller.",
			Type:        schema.TypeList,
			MaxItems:    1,
			Optional:    true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"additional_properties": {
						Type:             schema.TypeString,
						Optional:         true,
						DiffSuppressFunc: SuppressDiffAdditionProps,
					},
					"class_id": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"ipv4_address": {
						Description: "An IPv4 address in this IpAddress.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"ipv4_prefix_length": {
						Description: "The prefix length of the IPv4 address.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"ipv6_address": {
						Description: "An IPv6 address in this IpAddress.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"ipv6_prefix_length": {
						Description: "The prefix length of the IPv6 address.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				},
			},
		},
		"ipmi_username": {
			Description: "The IPMI user name of the controller.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"is_degraded": {
			Description: "The degraded status of the node.",
			Type:        schema.TypeBool,
			Optional:    true,
		},
		"is_hardware_virtualized": {
			Description: "The hardware virtualization status of the node.",
			Type:        schema.TypeBool,
			Optional:    true,
		},
		"is_secure_booted": {
			Description: "The secure boot status of the node.",
			Type:        schema.TypeBool,
			Optional:    true,
		},
		"key_management_device_to_cert_status": {
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"additional_properties": {
						Type:             schema.TypeString,
						Optional:         true,
						DiffSuppressFunc: SuppressDiffAdditionProps,
					},
					"class_id": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"key_management_server_name": {
						Description: "The name of the key management server.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"status": {
						Description: "The status of the certificate.",
						Type:        schema.TypeBool,
						Optional:    true,
					},
				},
			},
		},
		"maintenance_state": {
			Description: "The maintenance state of the node.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"memory_capacity_bytes": {
			Description: "The memory capacity in bytes of the node.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"memory_size_bytes": {
			Description: "The memory size in bytes of the node.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"memory_usage_bytes": {
			Description: "The memory usage in bytes of the node.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"mod_time": {
			Description: "The time when this managed object was last modified.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"moid": {
			Description: "The unique identifier of this Managed Object instance.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"node_ext_id": {
			Description: "The unique identifier of the node.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"node_serial": {
			Description: "The serial number of this node.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"node_status": {
			Description: "The status of the node such as NORMAL, TO_BE_REMOVED, OK_TO_BE_REMOVED,\nNEW_NODE, TO_BE_PREPROTECTED, PREPROTECTED.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"number_of_cpu_cores": {
			Description: "The number of CPU cores on the node.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"number_of_cpu_sockets": {
			Description: "The number of sockets on the node.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"number_of_cpu_threads": {
			Description: "The number of threads on the node.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"object_type": {
			Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"owners": {
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Schema{
				Type: schema.TypeString}},
		"parent": {
			Description: "A reference to a moBaseMo resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
			Type:        schema.TypeList,
			MaxItems:    1,
			Optional:    true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"additional_properties": {
						Type:             schema.TypeString,
						Optional:         true,
						DiffSuppressFunc: SuppressDiffAdditionProps,
					},
					"class_id": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"moid": {
						Description: "The Moid of the referenced REST resource.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the remote type referred by this relationship.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"selector": {
						Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				},
			},
		},
		"permission_resources": {
			Description: "An array of relationships to moBaseMo resources.",
			Type:        schema.TypeList,
			Optional:    true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"additional_properties": {
						Type:             schema.TypeString,
						Optional:         true,
						DiffSuppressFunc: SuppressDiffAdditionProps,
					},
					"class_id": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"moid": {
						Description: "The Moid of the referenced REST resource.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the remote type referred by this relationship.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"selector": {
						Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				},
			},
		},
		"physical_server": {
			Description: "A reference to a computePhysical resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
			Type:        schema.TypeList,
			MaxItems:    1,
			Optional:    true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"additional_properties": {
						Type:             schema.TypeString,
						Optional:         true,
						DiffSuppressFunc: SuppressDiffAdditionProps,
					},
					"class_id": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"moid": {
						Description: "The Moid of the referenced REST resource.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the remote type referred by this relationship.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"selector": {
						Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				},
			},
		},
		"reboot_pending": {
			Description: "The reboot pending status of the node.",
			Type:        schema.TypeBool,
			Optional:    true,
		},
		"registered_device": {
			Description: "A reference to a assetDeviceRegistration resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
			Type:        schema.TypeList,
			MaxItems:    1,
			Optional:    true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"additional_properties": {
						Type:             schema.TypeString,
						Optional:         true,
						DiffSuppressFunc: SuppressDiffAdditionProps,
					},
					"class_id": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"moid": {
						Description: "The Moid of the referenced REST resource.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the remote type referred by this relationship.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"selector": {
						Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				},
			},
		},
		"shared_scope": {
			Description: "Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.\nObjects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"storage_capacity_bytes": {
			Description: "The storage capacity in bytes of the node.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"storage_usage_bytes": {
			Description: "The storage usage in bytes of the node.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"tags": {
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"additional_properties": {
						Type:             schema.TypeString,
						Optional:         true,
						DiffSuppressFunc: SuppressDiffAdditionProps,
					},
					"key": {
						Description: "The string representation of a tag key.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"value": {
						Description: "The string representation of a tag value.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				},
			},
		},
		"version_context": {
			Description: "The versioning info for this managed object.",
			Type:        schema.TypeList,
			MaxItems:    1,
			Optional:    true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"additional_properties": {
						Type:             schema.TypeString,
						Optional:         true,
						DiffSuppressFunc: SuppressDiffAdditionProps,
					},
					"class_id": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"interested_mos": {
						Type:     schema.TypeList,
						Optional: true,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"additional_properties": {
									Type:             schema.TypeString,
									Optional:         true,
									DiffSuppressFunc: SuppressDiffAdditionProps,
								},
								"class_id": {
									Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"moid": {
									Description: "The Moid of the referenced REST resource.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"object_type": {
									Description: "The fully-qualified name of the remote type referred by this relationship.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"selector": {
									Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
									Type:        schema.TypeString,
									Optional:    true,
								},
							},
						},
					},
					"marked_for_deletion": {
						Description: "The flag to indicate if snapshot is marked for deletion or not. If flag is set then snapshot will be removed after the successful deployment of the policy.",
						Type:        schema.TypeBool,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"ref_mo": {
						Description: "A reference to the original Managed Object.",
						Type:        schema.TypeList,
						MaxItems:    1,
						Optional:    true,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"additional_properties": {
									Type:             schema.TypeString,
									Optional:         true,
									DiffSuppressFunc: SuppressDiffAdditionProps,
								},
								"class_id": {
									Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"moid": {
									Description: "The Moid of the referenced REST resource.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"object_type": {
									Description: "The fully-qualified name of the remote type referred by this relationship.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"selector": {
									Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
									Type:        schema.TypeString,
									Optional:    true,
								},
							},
						},
					},
					"timestamp": {
						Description: "The time this versioned Managed Object was created.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"nr_version": {
						Description: "The version of the Managed Object, e.g. an incrementing number or a hash id.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"version_type": {
						Description: "Specifies type of version. Currently the only supported value is \"Configured\"\nthat is used to keep track of snapshots of policies and profiles that are intended\nto be configured to target endpoints.\n* `Modified` - Version created every time an object is modified.\n* `Configured` - Version created every time an object is configured to the service profile.\n* `Deployed` - Version created for objects related to a service profile when it is deployed.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				},
			},
		},
		"vms": {
			Description: "An array of relationships to hciBaseVm resources.",
			Type:        schema.TypeList,
			Optional:    true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"additional_properties": {
						Type:             schema.TypeString,
						Optional:         true,
						DiffSuppressFunc: SuppressDiffAdditionProps,
					},
					"class_id": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"moid": {
						Description: "The Moid of the referenced REST resource.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the remote type referred by this relationship.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"selector": {
						Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				},
			},
		},
	}
	return schemaMap
}

func dataSourceHciNode() *schema.Resource {
	var subSchema = getHciNodeSchema()
	var model = getHciNodeSchema()
	model["results"] = &schema.Schema{
		Type:     schema.TypeList,
		Elem:     &schema.Resource{Schema: subSchema},
		Computed: true,
	}
	return &schema.Resource{
		ReadContext: dataSourceHciNodeRead,
		Schema:      model}
}

func dataSourceHciNodeRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.HciNode{}
	if v, ok := d.GetOk("account_moid"); ok {
		x := (v.(string))
		o.SetAccountMoid(x)
	}

	if v, ok := d.GetOk("additional_properties"); ok {
		x := []byte(v.(string))
		var x1 interface{}
		err := json.Unmarshal(x, &x1)
		if err == nil && x1 != nil {
			o.AdditionalProperties = x1.(map[string]interface{})
		}
	}

	if v, ok := d.GetOk("ancestors"); ok {
		x := make([]models.MoBaseMoRelationship, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := &models.MoMoRef{}
			l := s[i].(map[string]interface{})
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.AdditionalProperties = x1.(map[string]interface{})
					}
				}
			}
			o.SetClassId("mo.MoRef")
			if v, ok := l["moid"]; ok {
				{
					x := (v.(string))
					o.SetMoid(x)
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			if v, ok := l["selector"]; ok {
				{
					x := (v.(string))
					o.SetSelector(x)
				}
			}
			x = append(x, models.MoMoRefAsMoBaseMoRelationship(o))
		}
		o.SetAncestors(x)
	}

	if v, ok := d.GetOk("block_model"); ok {
		x := (v.(string))
		o.SetBlockModel(x)
	}

	if v, ok := d.GetOk("block_serial"); ok {
		x := (v.(string))
		o.SetBlockSerial(x)
	}

	if v, ok := d.GetOkExists("boot_time_usecs"); ok {
		x := int64(v.(int))
		o.SetBootTimeUsecs(x)
	}

	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}

	if v, ok := d.GetOk("cluster"); ok {
		p := make([]models.HciClusterRelationship, 0, 1)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			l := s[i].(map[string]interface{})
			o := &models.MoMoRef{}
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.AdditionalProperties = x1.(map[string]interface{})
					}
				}
			}
			o.SetClassId("mo.MoRef")
			if v, ok := l["moid"]; ok {
				{
					x := (v.(string))
					o.SetMoid(x)
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			if v, ok := l["selector"]; ok {
				{
					x := (v.(string))
					o.SetSelector(x)
				}
			}
			p = append(p, models.MoMoRefAsHciClusterRelationship(o))
		}
		if len(p) > 0 {
			x := p[0]
			o.SetCluster(x)
		}
	}

	if v, ok := d.GetOk("cluster_ext_id"); ok {
		x := (v.(string))
		o.SetClusterExtId(x)
	}

	if v, ok := d.GetOk("cluster_name"); ok {
		x := (v.(string))
		o.SetClusterName(x)
	}

	if v, ok := d.GetOk("controller_vm_backplane_address"); ok {
		p := make([]models.HciIpAddress, 0, 1)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			l := s[i].(map[string]interface{})
			o := &models.HciIpAddress{}
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.AdditionalProperties = x1.(map[string]interface{})
					}
				}
			}
			o.SetClassId("hci.IpAddress")
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			p = append(p, *o)
		}
		if len(p) > 0 {
			x := p[0]
			o.SetControllerVmBackplaneAddress(x)
		}
	}

	if v, ok := d.GetOk("controller_vm_external_address"); ok {
		p := make([]models.HciIpAddress, 0, 1)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			l := s[i].(map[string]interface{})
			o := &models.HciIpAddress{}
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.AdditionalProperties = x1.(map[string]interface{})
					}
				}
			}
			o.SetClassId("hci.IpAddress")
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			p = append(p, *o)
		}
		if len(p) > 0 {
			x := p[0]
			o.SetControllerVmExternalAddress(x)
		}
	}

	if v, ok := d.GetOkExists("controller_vm_id"); ok {
		x := int32(v.(int))
		o.SetControllerVmId(x)
	}

	if v, ok := d.GetOkExists("controller_vm_maintanence_mode"); ok {
		x := (v.(bool))
		o.SetControllerVmMaintanenceMode(x)
	}

	if v, ok := d.GetOk("controller_vm_nat_ip"); ok {
		p := make([]models.HciIpAddress, 0, 1)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			l := s[i].(map[string]interface{})
			o := &models.HciIpAddress{}
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.AdditionalProperties = x1.(map[string]interface{})
					}
				}
			}
			o.SetClassId("hci.IpAddress")
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			p = append(p, *o)
		}
		if len(p) > 0 {
			x := p[0]
			o.SetControllerVmNatIp(x)
		}
	}

	if v, ok := d.GetOkExists("controller_vm_nat_port"); ok {
		x := int32(v.(int))
		o.SetControllerVmNatPort(x)
	}

	if v, ok := d.GetOk("controller_vm_rdma_backplane_address"); ok {
		p := make([]models.HciIpAddress, 0, 1)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			l := s[i].(map[string]interface{})
			o := &models.HciIpAddress{}
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.AdditionalProperties = x1.(map[string]interface{})
					}
				}
			}
			o.SetClassId("hci.IpAddress")
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			p = append(p, *o)
		}
		if len(p) > 0 {
			x := p[0]
			o.SetControllerVmRdmaBackplaneAddress(x)
		}
	}

	if v, ok := d.GetOk("controller_vm_server_uuid"); ok {
		x := (v.(string))
		o.SetControllerVmServerUuid(x)
	}

	if v, ok := d.GetOkExists("cpu_capacity_hz"); ok {
		x := int64(v.(int))
		o.SetCpuCapacityHz(x)
	}

	if v, ok := d.GetOkExists("cpu_frequency_hz"); ok {
		x := int64(v.(int))
		o.SetCpuFrequencyHz(x)
	}

	if v, ok := d.GetOk("cpu_model"); ok {
		x := (v.(string))
		o.SetCpuModel(x)
	}

	if v, ok := d.GetOkExists("cpu_usage_hz"); ok {
		x := int64(v.(int))
		o.SetCpuUsageHz(x)
	}

	if v, ok := d.GetOk("create_time"); ok {
		x, _ := time.Parse(time.RFC1123, v.(string))
		o.SetCreateTime(x)
	}

	if v, ok := d.GetOk("default_vhd_container_uuid"); ok {
		x := (v.(string))
		o.SetDefaultVhdContainerUuid(x)
	}

	if v, ok := d.GetOk("default_vhd_location"); ok {
		x := (v.(string))
		o.SetDefaultVhdLocation(x)
	}

	if v, ok := d.GetOk("default_vm_container_uuid"); ok {
		x := (v.(string))
		o.SetDefaultVmContainerUuid(x)
	}

	if v, ok := d.GetOk("default_vm_location"); ok {
		x := (v.(string))
		o.SetDefaultVmLocation(x)
	}

	if v, ok := d.GetOkExists("disk_count"); ok {
		x := int64(v.(int))
		o.SetDiskCount(x)
	}

	if v, ok := d.GetOk("disks"); ok {
		x := make([]models.HciDiskRelationship, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := &models.MoMoRef{}
			l := s[i].(map[string]interface{})
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.AdditionalProperties = x1.(map[string]interface{})
					}
				}
			}
			o.SetClassId("mo.MoRef")
			if v, ok := l["moid"]; ok {
				{
					x := (v.(string))
					o.SetMoid(x)
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			if v, ok := l["selector"]; ok {
				{
					x := (v.(string))
					o.SetSelector(x)
				}
			}
			x = append(x, models.MoMoRefAsHciDiskRelationship(o))
		}
		o.SetDisks(x)
	}

	if v, ok := d.GetOk("domain_group_moid"); ok {
		x := (v.(string))
		o.SetDomainGroupMoid(x)
	}

	if v, ok := d.GetOk("failover_cluster_fqdn"); ok {
		x := (v.(string))
		o.SetFailoverClusterFqdn(x)
	}

	if v, ok := d.GetOk("failover_cluster_node_status"); ok {
		x := (v.(string))
		o.SetFailoverClusterNodeStatus(x)
	}

	if v, ok := d.GetOkExists("gpu_count"); ok {
		x := int64(v.(int))
		o.SetGpuCount(x)
	}

	if v, ok := d.GetOk("gpu_driver_version"); ok {
		x := (v.(string))
		o.SetGpuDriverVersion(x)
	}

	if v, ok := d.GetOk("gpus"); ok {
		x := make([]models.HciGpuRelationship, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := &models.MoMoRef{}
			l := s[i].(map[string]interface{})
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.AdditionalProperties = x1.(map[string]interface{})
					}
				}
			}
			o.SetClassId("mo.MoRef")
			if v, ok := l["moid"]; ok {
				{
					x := (v.(string))
					o.SetMoid(x)
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			if v, ok := l["selector"]; ok {
				{
					x := (v.(string))
					o.SetSelector(x)
				}
			}
			x = append(x, models.MoMoRefAsHciGpuRelationship(o))
		}
		o.SetGpus(x)
	}

	if v, ok := d.GetOkExists("has_csr"); ok {
		x := (v.(bool))
		o.SetHasCsr(x)
	}

	if v, ok := d.GetOk("host_name"); ok {
		x := (v.(string))
		o.SetHostName(x)
	}

	if v, ok := d.GetOk("host_type"); ok {
		x := (v.(string))
		o.SetHostType(x)
	}

	if v, ok := d.GetOk("hypervisor_acropolis_connection_state"); ok {
		x := (v.(string))
		o.SetHypervisorAcropolisConnectionState(x)
	}

	if v, ok := d.GetOk("hypervisor_external_address"); ok {
		p := make([]models.HciIpAddress, 0, 1)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			l := s[i].(map[string]interface{})
			o := &models.HciIpAddress{}
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.AdditionalProperties = x1.(map[string]interface{})
					}
				}
			}
			o.SetClassId("hci.IpAddress")
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			p = append(p, *o)
		}
		if len(p) > 0 {
			x := p[0]
			o.SetHypervisorExternalAddress(x)
		}
	}

	if v, ok := d.GetOkExists("hypervisor_number_of_vms"); ok {
		x := int64(v.(int))
		o.SetHypervisorNumberOfVms(x)
	}

	if v, ok := d.GetOk("hypervisor_state"); ok {
		x := (v.(string))
		o.SetHypervisorState(x)
	}

	if v, ok := d.GetOk("hypervisor_type"); ok {
		x := (v.(string))
		o.SetHypervisorType(x)
	}

	if v, ok := d.GetOk("hypervisor_user_name"); ok {
		x := (v.(string))
		o.SetHypervisorUserName(x)
	}

	if v, ok := d.GetOk("hypervisor_version"); ok {
		x := (v.(string))
		o.SetHypervisorVersion(x)
	}

	if v, ok := d.GetOk("ipmi_ip"); ok {
		p := make([]models.HciIpAddress, 0, 1)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			l := s[i].(map[string]interface{})
			o := &models.HciIpAddress{}
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.AdditionalProperties = x1.(map[string]interface{})
					}
				}
			}
			o.SetClassId("hci.IpAddress")
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			p = append(p, *o)
		}
		if len(p) > 0 {
			x := p[0]
			o.SetIpmiIp(x)
		}
	}

	if v, ok := d.GetOk("ipmi_username"); ok {
		x := (v.(string))
		o.SetIpmiUsername(x)
	}

	if v, ok := d.GetOkExists("is_degraded"); ok {
		x := (v.(bool))
		o.SetIsDegraded(x)
	}

	if v, ok := d.GetOkExists("is_hardware_virtualized"); ok {
		x := (v.(bool))
		o.SetIsHardwareVirtualized(x)
	}

	if v, ok := d.GetOkExists("is_secure_booted"); ok {
		x := (v.(bool))
		o.SetIsSecureBooted(x)
	}

	if v, ok := d.GetOk("key_management_device_to_cert_status"); ok {
		x := make([]models.HciKeyManagementDeviceToCertStatusInfo, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := &models.HciKeyManagementDeviceToCertStatusInfo{}
			l := s[i].(map[string]interface{})
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.AdditionalProperties = x1.(map[string]interface{})
					}
				}
			}
			o.SetClassId("hci.KeyManagementDeviceToCertStatusInfo")
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			x = append(x, *o)
		}
		o.SetKeyManagementDeviceToCertStatus(x)
	}

	if v, ok := d.GetOk("maintenance_state"); ok {
		x := (v.(string))
		o.SetMaintenanceState(x)
	}

	if v, ok := d.GetOkExists("memory_capacity_bytes"); ok {
		x := int64(v.(int))
		o.SetMemoryCapacityBytes(x)
	}

	if v, ok := d.GetOkExists("memory_size_bytes"); ok {
		x := int64(v.(int))
		o.SetMemorySizeBytes(x)
	}

	if v, ok := d.GetOkExists("memory_usage_bytes"); ok {
		x := int64(v.(int))
		o.SetMemoryUsageBytes(x)
	}

	if v, ok := d.GetOk("mod_time"); ok {
		x, _ := time.Parse(time.RFC1123, v.(string))
		o.SetModTime(x)
	}

	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.SetMoid(x)
	}

	if v, ok := d.GetOk("node_ext_id"); ok {
		x := (v.(string))
		o.SetNodeExtId(x)
	}

	if v, ok := d.GetOk("node_serial"); ok {
		x := (v.(string))
		o.SetNodeSerial(x)
	}

	if v, ok := d.GetOk("node_status"); ok {
		x := (v.(string))
		o.SetNodeStatus(x)
	}

	if v, ok := d.GetOkExists("number_of_cpu_cores"); ok {
		x := int64(v.(int))
		o.SetNumberOfCpuCores(x)
	}

	if v, ok := d.GetOkExists("number_of_cpu_sockets"); ok {
		x := int64(v.(int))
		o.SetNumberOfCpuSockets(x)
	}

	if v, ok := d.GetOkExists("number_of_cpu_threads"); ok {
		x := int64(v.(int))
		o.SetNumberOfCpuThreads(x)
	}

	if v, ok := d.GetOk("object_type"); ok {
		x := (v.(string))
		o.SetObjectType(x)
	}

	if v, ok := d.GetOk("owners"); ok {
		x := make([]string, 0)
		y := reflect.ValueOf(v)
		for i := 0; i < y.Len(); i++ {
			if y.Index(i).Interface() != nil {
				x = append(x, y.Index(i).Interface().(string))
			}
		}
		o.SetOwners(x)
	}

	if v, ok := d.GetOk("parent"); ok {
		p := make([]models.MoBaseMoRelationship, 0, 1)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			l := s[i].(map[string]interface{})
			o := &models.MoMoRef{}
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.AdditionalProperties = x1.(map[string]interface{})
					}
				}
			}
			o.SetClassId("mo.MoRef")
			if v, ok := l["moid"]; ok {
				{
					x := (v.(string))
					o.SetMoid(x)
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			if v, ok := l["selector"]; ok {
				{
					x := (v.(string))
					o.SetSelector(x)
				}
			}
			p = append(p, models.MoMoRefAsMoBaseMoRelationship(o))
		}
		if len(p) > 0 {
			x := p[0]
			o.SetParent(x)
		}
	}

	if v, ok := d.GetOk("permission_resources"); ok {
		x := make([]models.MoBaseMoRelationship, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := &models.MoMoRef{}
			l := s[i].(map[string]interface{})
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.AdditionalProperties = x1.(map[string]interface{})
					}
				}
			}
			o.SetClassId("mo.MoRef")
			if v, ok := l["moid"]; ok {
				{
					x := (v.(string))
					o.SetMoid(x)
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			if v, ok := l["selector"]; ok {
				{
					x := (v.(string))
					o.SetSelector(x)
				}
			}
			x = append(x, models.MoMoRefAsMoBaseMoRelationship(o))
		}
		o.SetPermissionResources(x)
	}

	if v, ok := d.GetOk("physical_server"); ok {
		p := make([]models.ComputePhysicalRelationship, 0, 1)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			l := s[i].(map[string]interface{})
			o := &models.MoMoRef{}
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.AdditionalProperties = x1.(map[string]interface{})
					}
				}
			}
			o.SetClassId("mo.MoRef")
			if v, ok := l["moid"]; ok {
				{
					x := (v.(string))
					o.SetMoid(x)
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			if v, ok := l["selector"]; ok {
				{
					x := (v.(string))
					o.SetSelector(x)
				}
			}
			p = append(p, models.MoMoRefAsComputePhysicalRelationship(o))
		}
		if len(p) > 0 {
			x := p[0]
			o.SetPhysicalServer(x)
		}
	}

	if v, ok := d.GetOkExists("reboot_pending"); ok {
		x := (v.(bool))
		o.SetRebootPending(x)
	}

	if v, ok := d.GetOk("registered_device"); ok {
		p := make([]models.AssetDeviceRegistrationRelationship, 0, 1)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			l := s[i].(map[string]interface{})
			o := &models.MoMoRef{}
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.AdditionalProperties = x1.(map[string]interface{})
					}
				}
			}
			o.SetClassId("mo.MoRef")
			if v, ok := l["moid"]; ok {
				{
					x := (v.(string))
					o.SetMoid(x)
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			if v, ok := l["selector"]; ok {
				{
					x := (v.(string))
					o.SetSelector(x)
				}
			}
			p = append(p, models.MoMoRefAsAssetDeviceRegistrationRelationship(o))
		}
		if len(p) > 0 {
			x := p[0]
			o.SetRegisteredDevice(x)
		}
	}

	if v, ok := d.GetOk("shared_scope"); ok {
		x := (v.(string))
		o.SetSharedScope(x)
	}

	if v, ok := d.GetOkExists("storage_capacity_bytes"); ok {
		x := int64(v.(int))
		o.SetStorageCapacityBytes(x)
	}

	if v, ok := d.GetOkExists("storage_usage_bytes"); ok {
		x := int64(v.(int))
		o.SetStorageUsageBytes(x)
	}

	if v, ok := d.GetOk("tags"); ok {
		x := make([]models.MoTag, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := &models.MoTag{}
			l := s[i].(map[string]interface{})
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.AdditionalProperties = x1.(map[string]interface{})
					}
				}
			}
			if v, ok := l["key"]; ok {
				{
					x := (v.(string))
					o.SetKey(x)
				}
			}
			if v, ok := l["value"]; ok {
				{
					x := (v.(string))
					o.SetValue(x)
				}
			}
			x = append(x, *o)
		}
		o.SetTags(x)
	}

	if v, ok := d.GetOk("version_context"); ok {
		p := make([]models.MoVersionContext, 0, 1)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			l := s[i].(map[string]interface{})
			o := &models.MoVersionContext{}
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.AdditionalProperties = x1.(map[string]interface{})
					}
				}
			}
			o.SetClassId("mo.VersionContext")
			if v, ok := l["interested_mos"]; ok {
				{
					x := make([]models.MoMoRef, 0)
					s := v.([]interface{})
					for i := 0; i < len(s); i++ {
						o := models.NewMoMoRefWithDefaults()
						l := s[i].(map[string]interface{})
						if v, ok := l["additional_properties"]; ok {
							{
								x := []byte(v.(string))
								var x1 interface{}
								err := json.Unmarshal(x, &x1)
								if err == nil && x1 != nil {
									o.AdditionalProperties = x1.(map[string]interface{})
								}
							}
						}
						o.SetClassId("mo.MoRef")
						if v, ok := l["moid"]; ok {
							{
								x := (v.(string))
								o.SetMoid(x)
							}
						}
						if v, ok := l["object_type"]; ok {
							{
								x := (v.(string))
								o.SetObjectType(x)
							}
						}
						if v, ok := l["selector"]; ok {
							{
								x := (v.(string))
								o.SetSelector(x)
							}
						}
						x = append(x, *o)
					}
					if len(x) > 0 {
						o.SetInterestedMos(x)
					}
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			p = append(p, *o)
		}
		if len(p) > 0 {
			x := p[0]
			o.SetVersionContext(x)
		}
	}

	if v, ok := d.GetOk("vms"); ok {
		x := make([]models.HciBaseVmRelationship, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := &models.MoMoRef{}
			l := s[i].(map[string]interface{})
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.AdditionalProperties = x1.(map[string]interface{})
					}
				}
			}
			o.SetClassId("mo.MoRef")
			if v, ok := l["moid"]; ok {
				{
					x := (v.(string))
					o.SetMoid(x)
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			if v, ok := l["selector"]; ok {
				{
					x := (v.(string))
					o.SetSelector(x)
				}
			}
			x = append(x, models.MoMoRefAsHciBaseVmRelationship(o))
		}
		o.SetVms(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of HciNode object failed with error : %s", err.Error())
	}
	countResponse, _, responseErr := conn.ApiClient.HciApi.GetHciNodeList(conn.ctx).Filter(getRequestParams(data)).Count(true).Execute()
	if responseErr != nil {
		errorType := fmt.Sprintf("%T", responseErr)
		if strings.Contains(errorType, "GenericOpenAPIError") {
			responseErr := responseErr.(*models.GenericOpenAPIError)
			return diag.Errorf("error occurred while fetching count of HciNode: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		return diag.Errorf("error occurred while fetching count of HciNode: %s", responseErr.Error())
	}
	count := countResponse.MoDocumentCount.GetCount()
	if count == 0 {
		return diag.Errorf("your query for HciNode data source did not return any results. Please change your search criteria and try again")
	}
	var i int32
	var hciNodeResults = make([]map[string]interface{}, 0, 0)
	for i = 0; i < count; i += 100 {
		resMo, _, responseErr := conn.ApiClient.HciApi.GetHciNodeList(conn.ctx).Filter(getRequestParams(data)).Top(100).Skip(i).Execute()
		if responseErr != nil {
			errorType := fmt.Sprintf("%T", responseErr)
			if strings.Contains(errorType, "GenericOpenAPIError") {
				responseErr := responseErr.(*models.GenericOpenAPIError)
				return diag.Errorf("error occurred while fetching HciNode: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
			}
			return diag.Errorf("error occurred while fetching HciNode: %s", responseErr.Error())
		}
		results := resMo.HciNodeList.GetResults()
		switch reflect.TypeOf(results).Kind() {
		case reflect.Slice:
			for k := 0; k < len(results); k++ {
				var s = results[k]
				var temp = make(map[string]interface{})
				temp["account_moid"] = (s.GetAccountMoid())
				temp["additional_properties"] = flattenAdditionalProperties(s.AdditionalProperties)

				temp["ancestors"] = flattenListMoBaseMoRelationship(s.GetAncestors(), d)
				temp["block_model"] = (s.GetBlockModel())
				temp["block_serial"] = (s.GetBlockSerial())
				temp["boot_time_usecs"] = (s.GetBootTimeUsecs())
				temp["class_id"] = (s.GetClassId())

				temp["cluster"] = flattenMapHciClusterRelationship(s.GetCluster(), d)
				temp["cluster_ext_id"] = (s.GetClusterExtId())
				temp["cluster_name"] = (s.GetClusterName())

				temp["controller_vm_backplane_address"] = flattenMapHciIpAddress(s.GetControllerVmBackplaneAddress(), d)

				temp["controller_vm_external_address"] = flattenMapHciIpAddress(s.GetControllerVmExternalAddress(), d)
				temp["controller_vm_id"] = (s.GetControllerVmId())
				temp["controller_vm_maintanence_mode"] = (s.GetControllerVmMaintanenceMode())

				temp["controller_vm_nat_ip"] = flattenMapHciIpAddress(s.GetControllerVmNatIp(), d)
				temp["controller_vm_nat_port"] = (s.GetControllerVmNatPort())

				temp["controller_vm_rdma_backplane_address"] = flattenMapHciIpAddress(s.GetControllerVmRdmaBackplaneAddress(), d)
				temp["controller_vm_server_uuid"] = (s.GetControllerVmServerUuid())
				temp["cpu_capacity_hz"] = (s.GetCpuCapacityHz())
				temp["cpu_frequency_hz"] = (s.GetCpuFrequencyHz())
				temp["cpu_model"] = (s.GetCpuModel())
				temp["cpu_usage_hz"] = (s.GetCpuUsageHz())

				temp["create_time"] = (s.GetCreateTime()).String()
				temp["default_vhd_container_uuid"] = (s.GetDefaultVhdContainerUuid())
				temp["default_vhd_location"] = (s.GetDefaultVhdLocation())
				temp["default_vm_container_uuid"] = (s.GetDefaultVmContainerUuid())
				temp["default_vm_location"] = (s.GetDefaultVmLocation())
				temp["disk_count"] = (s.GetDiskCount())

				temp["disks"] = flattenListHciDiskRelationship(s.GetDisks(), d)
				temp["domain_group_moid"] = (s.GetDomainGroupMoid())
				temp["failover_cluster_fqdn"] = (s.GetFailoverClusterFqdn())
				temp["failover_cluster_node_status"] = (s.GetFailoverClusterNodeStatus())
				temp["gpu_count"] = (s.GetGpuCount())
				temp["gpu_driver_version"] = (s.GetGpuDriverVersion())

				temp["gpus"] = flattenListHciGpuRelationship(s.GetGpus(), d)
				temp["has_csr"] = (s.GetHasCsr())
				temp["host_name"] = (s.GetHostName())
				temp["host_type"] = (s.GetHostType())
				temp["hypervisor_acropolis_connection_state"] = (s.GetHypervisorAcropolisConnectionState())

				temp["hypervisor_external_address"] = flattenMapHciIpAddress(s.GetHypervisorExternalAddress(), d)
				temp["hypervisor_number_of_vms"] = (s.GetHypervisorNumberOfVms())
				temp["hypervisor_state"] = (s.GetHypervisorState())
				temp["hypervisor_type"] = (s.GetHypervisorType())
				temp["hypervisor_user_name"] = (s.GetHypervisorUserName())
				temp["hypervisor_version"] = (s.GetHypervisorVersion())

				temp["ipmi_ip"] = flattenMapHciIpAddress(s.GetIpmiIp(), d)
				temp["ipmi_username"] = (s.GetIpmiUsername())
				temp["is_degraded"] = (s.GetIsDegraded())
				temp["is_hardware_virtualized"] = (s.GetIsHardwareVirtualized())
				temp["is_secure_booted"] = (s.GetIsSecureBooted())

				temp["key_management_device_to_cert_status"] = flattenListHciKeyManagementDeviceToCertStatusInfo(s.GetKeyManagementDeviceToCertStatus(), d)
				temp["maintenance_state"] = (s.GetMaintenanceState())
				temp["memory_capacity_bytes"] = (s.GetMemoryCapacityBytes())
				temp["memory_size_bytes"] = (s.GetMemorySizeBytes())
				temp["memory_usage_bytes"] = (s.GetMemoryUsageBytes())

				temp["mod_time"] = (s.GetModTime()).String()
				temp["moid"] = (s.GetMoid())
				temp["node_ext_id"] = (s.GetNodeExtId())
				temp["node_serial"] = (s.GetNodeSerial())
				temp["node_status"] = (s.GetNodeStatus())
				temp["number_of_cpu_cores"] = (s.GetNumberOfCpuCores())
				temp["number_of_cpu_sockets"] = (s.GetNumberOfCpuSockets())
				temp["number_of_cpu_threads"] = (s.GetNumberOfCpuThreads())
				temp["object_type"] = (s.GetObjectType())
				temp["owners"] = (s.GetOwners())

				temp["parent"] = flattenMapMoBaseMoRelationship(s.GetParent(), d)

				temp["permission_resources"] = flattenListMoBaseMoRelationship(s.GetPermissionResources(), d)

				temp["physical_server"] = flattenMapComputePhysicalRelationship(s.GetPhysicalServer(), d)
				temp["reboot_pending"] = (s.GetRebootPending())

				temp["registered_device"] = flattenMapAssetDeviceRegistrationRelationship(s.GetRegisteredDevice(), d)
				temp["shared_scope"] = (s.GetSharedScope())
				temp["storage_capacity_bytes"] = (s.GetStorageCapacityBytes())
				temp["storage_usage_bytes"] = (s.GetStorageUsageBytes())

				temp["tags"] = flattenListMoTag(s.GetTags(), d)

				temp["version_context"] = flattenMapMoVersionContext(s.GetVersionContext(), d)

				temp["vms"] = flattenListHciBaseVmRelationship(s.GetVms(), d)
				hciNodeResults = append(hciNodeResults, temp)
			}
		}
	}
	log.Println("length of results: ", len(hciNodeResults))
	if err := d.Set("results", hciNodeResults); err != nil {
		return diag.Errorf("error occurred while setting results: %s", err.Error())
	}
	d.SetId(hciNodeResults[0]["moid"].(string))
	return de
}
