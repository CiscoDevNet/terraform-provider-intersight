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

func getApplianceUpgradeSchema() map[string]*schema.Schema {
	var schemaMap = make(map[string]*schema.Schema)
	schemaMap = map[string]*schema.Schema{"account": {
		Description: "A reference to a iamAccount resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
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
		"account_moid": {
			Description: "The Account ID for this managed object.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"active": {
			Description: "Indicates if the software upgrade is active or not.",
			Type:        schema.TypeBool,
			Optional:    true,
		},
		"additional_properties": {
			Type:             schema.TypeString,
			Optional:         true,
			DiffSuppressFunc: SuppressDiffAdditionProps,
		},
		"all_nodes_pingable": {
			Description: "True if all nodes in cluster are pingable, otherwise false.",
			Type:        schema.TypeBool,
			Optional:    true,
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
		"auto_created": {
			Description: "Indicates that the request was automatically created by the system.",
			Type:        schema.TypeBool,
			Optional:    true,
		},
		"class_id": {
			Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"completed_phases": {
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
					"current_node": {
						Description: "Id of the node the upgrade phase is running on.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"current_node_hostname": {
						Description: "Hostname of the node the upgrade phase is running on.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"elapsed_time": {
						Description: "Elapsed time in seconds to complete the current upgrade phase.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"end_time": {
						Description: "End date of the software upgrade phase.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"failed": {
						Description: "Indicates if the upgrade phase has failed or not.",
						Type:        schema.TypeBool,
						Optional:    true,
					},
					"message": {
						Description: "Status message set during the upgrade phase.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"name": {
						Description: "Name of the upgrade phase.\n* `init` - Upgrade service initialization phase.\n* `CheckCluster` - For a multinode system, check that all nodes in the cluster are connected and running.\n* `monitor` - Monitor a required upgrade.\n* `UpdateAnsibleHosts` - Update the ansible hosts file with latest template.\n* `SyncImages` - For a multinode system, sync image files between nodes.\n* `Prepare` - Upgrade service prepares folders and templated files.\n* `ServiceLoad` - Upgrade service loads the service images into the local docker cache.\n* `UiLoad` - Upgrade service loads the UI packages into the local cache.\n* `GenerateConfig` - Upgrade service generates the Kubernetes configuration files.\n* `DeployService` - Upgrade service deploys the Kubernetes services.\n* `UpgradeOS` - Run /opt/cisco/bin/onprem-upgrade-start.sh for each node.\n* `UpgradeServices` - Run /opt/cisco/bin/onprem-upgrade-start.sh per node.\n* `VerifyPlaybookSuccess` - Verify the upgrade playbook for UpgradeOS or UpgradeServices completed successfully.\n* `FinishUpgrade` - Run /opt/cisco/bin/onprem-upgrade-finish.sh for each node.\n* `Success` - Upgrade completed successfully.\n* `Fail` - Indicates that the upgrade process has failed.\n* `Cancel` - Indicates that the upgrade was canceled by the Intersight Appliance.\n* `Telemetry` - Upgrade service sends basic telemetry data to the Intersight.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"retry_count": {
						Description: "Retry count of the upgrade phase.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"start_time": {
						Description: "Start date of the software upgrade phase.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"status": {
						Description: "Status of the upgrade phase.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				},
			},
		},
		"create_time": {
			Description: "The time when this managed object was created.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"current_phase": {
			Description: "Current phase of the Intersight Appliance's software upgrade.",
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
					"current_node": {
						Description: "Id of the node the upgrade phase is running on.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"current_node_hostname": {
						Description: "Hostname of the node the upgrade phase is running on.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"elapsed_time": {
						Description: "Elapsed time in seconds to complete the current upgrade phase.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"end_time": {
						Description: "End date of the software upgrade phase.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"failed": {
						Description: "Indicates if the upgrade phase has failed or not.",
						Type:        schema.TypeBool,
						Optional:    true,
					},
					"message": {
						Description: "Status message set during the upgrade phase.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"name": {
						Description: "Name of the upgrade phase.\n* `init` - Upgrade service initialization phase.\n* `CheckCluster` - For a multinode system, check that all nodes in the cluster are connected and running.\n* `monitor` - Monitor a required upgrade.\n* `UpdateAnsibleHosts` - Update the ansible hosts file with latest template.\n* `SyncImages` - For a multinode system, sync image files between nodes.\n* `Prepare` - Upgrade service prepares folders and templated files.\n* `ServiceLoad` - Upgrade service loads the service images into the local docker cache.\n* `UiLoad` - Upgrade service loads the UI packages into the local cache.\n* `GenerateConfig` - Upgrade service generates the Kubernetes configuration files.\n* `DeployService` - Upgrade service deploys the Kubernetes services.\n* `UpgradeOS` - Run /opt/cisco/bin/onprem-upgrade-start.sh for each node.\n* `UpgradeServices` - Run /opt/cisco/bin/onprem-upgrade-start.sh per node.\n* `VerifyPlaybookSuccess` - Verify the upgrade playbook for UpgradeOS or UpgradeServices completed successfully.\n* `FinishUpgrade` - Run /opt/cisco/bin/onprem-upgrade-finish.sh for each node.\n* `Success` - Upgrade completed successfully.\n* `Fail` - Indicates that the upgrade process has failed.\n* `Cancel` - Indicates that the upgrade was canceled by the Intersight Appliance.\n* `Telemetry` - Upgrade service sends basic telemetry data to the Intersight.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"retry_count": {
						Description: "Retry count of the upgrade phase.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"start_time": {
						Description: "Start date of the software upgrade phase.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"status": {
						Description: "Status of the upgrade phase.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				},
			},
		},
		"description": {
			Description: "Description of the software upgrade.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"domain_group_moid": {
			Description: "The DomainGroup ID for this managed object.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"elapsed_time": {
			Description: "Elapsed time in seconds during the software upgrade.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"end_time": {
			Description: "End date of the software upgrade.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"error_code": {
			Description: "Error code for Intersight Appliance's software upgrade. In case of failure - this code will help decide if software upgrade can be retried.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"fingerprint": {
			Description: "Software upgrade manifest's fingerprint.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"image_bundle": {
			Description: "A reference to a applianceImageBundle resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
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
		"is_rolling_back": {
			Description: "Track if software upgrade is upgrading or rolling back.",
			Type:        schema.TypeBool,
			Optional:    true,
		},
		"is_user_triggered": {
			Description: "Indicates if the upgrade is triggered by user or due to schedule.",
			Type:        schema.TypeBool,
			Optional:    true,
		},
		"messages": {
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Schema{
				Type: schema.TypeString}},
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
		"node_info": {
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
					"gateway": {
						Description: "Gateway ip address of the cluster node.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"hostname": {
						Description: "Hostname of the cluster node.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"ip_address": {
						Description: "Ip address of the cluster node.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"netmask": {
						Description: "Netmask of the cluster node.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"node_id": {
						Description: "Id number of the cluster node.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"node_moid": {
						Description: "Moid of the corresponding appliance.ClusterInfo or appliance.NodeInfo mo.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"node_type": {
						Description: "The node type of Intersight Virtual Appliance.\n* `standalone` - Single Node Intersight Virtual Appliance.\n* `management` - Management node type when Intersight Virtual Appliance is running as management-worker deployment.\n* `hamanagement` - Management node type when Intersight Virtual Appliance is running as multi node HA deployment.\n* `metrics` - Metrics node when Intersight Virtual Appliance is running management-metrics node.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"status": {
						Description: "Status of the cluster node.\n* `Unknown` - The status of the appliance node is unknown.\n* `Operational` - The appliance node is operational.\n* `Impaired` - The appliance node is impaired.\n* `AttentionNeeded` - The appliance node needs attention.\n* `ReadyToJoin` - The node is ready to be added to a standalone Intersight Appliance to form a cluster.\n* `OutOfService` - The user has taken this node (part of a cluster) to out of service.\n* `ReadyForReplacement` - The cluster node is ready to be replaced.\n* `ReplacementInProgress` - The cluster node replacement is in progress.\n* `ReplacementFailed` - There was a failure during the cluster node replacement.\n* `WorkerNodeInstInProgress` - The worker node installation is in progress.\n* `WorkerNodeInstSuccess` - The worker node installation succeeded.\n* `WorkerNodeInstFailed` - The worker node installation failed.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				},
			},
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
		"previous_install_attempts": {
			Description: "The number of previous failed install attempts of the same upgrade version.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"requires": {
			Description: "A reference to a applianceUpgrade resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
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
		"rollback_needed": {
			Description: "Track if rollback is needed.",
			Type:        schema.TypeBool,
			Optional:    true,
		},
		"rollback_phases": {
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
					"current_node": {
						Description: "Id of the node the upgrade phase is running on.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"current_node_hostname": {
						Description: "Hostname of the node the upgrade phase is running on.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"elapsed_time": {
						Description: "Elapsed time in seconds to complete the current upgrade phase.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"end_time": {
						Description: "End date of the software upgrade phase.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"failed": {
						Description: "Indicates if the upgrade phase has failed or not.",
						Type:        schema.TypeBool,
						Optional:    true,
					},
					"message": {
						Description: "Status message set during the upgrade phase.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"name": {
						Description: "Name of the upgrade phase.\n* `init` - Upgrade service initialization phase.\n* `CheckCluster` - For a multinode system, check that all nodes in the cluster are connected and running.\n* `monitor` - Monitor a required upgrade.\n* `UpdateAnsibleHosts` - Update the ansible hosts file with latest template.\n* `SyncImages` - For a multinode system, sync image files between nodes.\n* `Prepare` - Upgrade service prepares folders and templated files.\n* `ServiceLoad` - Upgrade service loads the service images into the local docker cache.\n* `UiLoad` - Upgrade service loads the UI packages into the local cache.\n* `GenerateConfig` - Upgrade service generates the Kubernetes configuration files.\n* `DeployService` - Upgrade service deploys the Kubernetes services.\n* `UpgradeOS` - Run /opt/cisco/bin/onprem-upgrade-start.sh for each node.\n* `UpgradeServices` - Run /opt/cisco/bin/onprem-upgrade-start.sh per node.\n* `VerifyPlaybookSuccess` - Verify the upgrade playbook for UpgradeOS or UpgradeServices completed successfully.\n* `FinishUpgrade` - Run /opt/cisco/bin/onprem-upgrade-finish.sh for each node.\n* `Success` - Upgrade completed successfully.\n* `Fail` - Indicates that the upgrade process has failed.\n* `Cancel` - Indicates that the upgrade was canceled by the Intersight Appliance.\n* `Telemetry` - Upgrade service sends basic telemetry data to the Intersight.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"retry_count": {
						Description: "Retry count of the upgrade phase.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"start_time": {
						Description: "Start date of the software upgrade phase.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"status": {
						Description: "Status of the upgrade phase.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				},
			},
		},
		"rollback_status": {
			Description: "Status of the Intersight Appliance's software rollback status.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"services": {
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Schema{
				Type: schema.TypeString}},
		"shared_scope": {
			Description: "Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.\nObjects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"start_time": {
			Description: "Start date of the software upgrade. UI can modify startTime to re-schedule an upgrade.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"status": {
			Description: "Status of the Intersight Appliance's software upgrade.",
			Type:        schema.TypeString,
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
		"total_nodes": {
			Description: "Total number of nodes this upgrade will run on.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"total_phases": {
			Description: "TotalPhase represents the total number of the upgradePhases for one upgrade.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"ui_packages": {
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Schema{
				Type: schema.TypeString}},
		"nr_version": {
			Description: "Software upgrade manifest's version.",
			Type:        schema.TypeString,
			Optional:    true,
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
	}
	return schemaMap
}

func dataSourceApplianceUpgrade() *schema.Resource {
	var subSchema = getApplianceUpgradeSchema()
	var model = getApplianceUpgradeSchema()
	model["results"] = &schema.Schema{
		Type:     schema.TypeList,
		Elem:     &schema.Resource{Schema: subSchema},
		Computed: true,
	}
	return &schema.Resource{
		ReadContext: dataSourceApplianceUpgradeRead,
		Schema:      model}
}

func dataSourceApplianceUpgradeRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.ApplianceUpgrade{}
	if v, ok := d.GetOk("account"); ok {
		p := make([]models.IamAccountRelationship, 0, 1)
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
			p = append(p, models.MoMoRefAsIamAccountRelationship(o))
		}
		if len(p) > 0 {
			x := p[0]
			o.SetAccount(x)
		}
	}

	if v, ok := d.GetOk("account_moid"); ok {
		x := (v.(string))
		o.SetAccountMoid(x)
	}

	if v, ok := d.GetOkExists("active"); ok {
		x := (v.(bool))
		o.SetActive(x)
	}

	if v, ok := d.GetOk("additional_properties"); ok {
		x := []byte(v.(string))
		var x1 interface{}
		err := json.Unmarshal(x, &x1)
		if err == nil && x1 != nil {
			o.AdditionalProperties = x1.(map[string]interface{})
		}
	}

	if v, ok := d.GetOkExists("all_nodes_pingable"); ok {
		x := (v.(bool))
		o.SetAllNodesPingable(x)
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

	if v, ok := d.GetOkExists("auto_created"); ok {
		x := (v.(bool))
		o.SetAutoCreated(x)
	}

	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}

	if v, ok := d.GetOk("completed_phases"); ok {
		x := make([]models.OnpremUpgradePhase, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := &models.OnpremUpgradePhase{}
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
			o.SetClassId("onprem.UpgradePhase")
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			x = append(x, *o)
		}
		o.SetCompletedPhases(x)
	}

	if v, ok := d.GetOk("create_time"); ok {
		x, _ := time.Parse(time.RFC1123, v.(string))
		o.SetCreateTime(x)
	}

	if v, ok := d.GetOk("current_phase"); ok {
		p := make([]models.OnpremUpgradePhase, 0, 1)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			l := s[i].(map[string]interface{})
			o := &models.OnpremUpgradePhase{}
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
			o.SetClassId("onprem.UpgradePhase")
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
			o.SetCurrentPhase(x)
		}
	}

	if v, ok := d.GetOk("description"); ok {
		x := (v.(string))
		o.SetDescription(x)
	}

	if v, ok := d.GetOk("domain_group_moid"); ok {
		x := (v.(string))
		o.SetDomainGroupMoid(x)
	}

	if v, ok := d.GetOkExists("elapsed_time"); ok {
		x := int64(v.(int))
		o.SetElapsedTime(x)
	}

	if v, ok := d.GetOk("end_time"); ok {
		x, _ := time.Parse(time.RFC1123, v.(string))
		o.SetEndTime(x)
	}

	if v, ok := d.GetOkExists("error_code"); ok {
		x := int64(v.(int))
		o.SetErrorCode(x)
	}

	if v, ok := d.GetOk("fingerprint"); ok {
		x := (v.(string))
		o.SetFingerprint(x)
	}

	if v, ok := d.GetOk("image_bundle"); ok {
		p := make([]models.ApplianceImageBundleRelationship, 0, 1)
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
			p = append(p, models.MoMoRefAsApplianceImageBundleRelationship(o))
		}
		if len(p) > 0 {
			x := p[0]
			o.SetImageBundle(x)
		}
	}

	if v, ok := d.GetOkExists("is_rolling_back"); ok {
		x := (v.(bool))
		o.SetIsRollingBack(x)
	}

	if v, ok := d.GetOkExists("is_user_triggered"); ok {
		x := (v.(bool))
		o.SetIsUserTriggered(x)
	}

	if v, ok := d.GetOk("messages"); ok {
		x := make([]string, 0)
		y := reflect.ValueOf(v)
		for i := 0; i < y.Len(); i++ {
			if y.Index(i).Interface() != nil {
				x = append(x, y.Index(i).Interface().(string))
			}
		}
		o.SetMessages(x)
	}

	if v, ok := d.GetOk("mod_time"); ok {
		x, _ := time.Parse(time.RFC1123, v.(string))
		o.SetModTime(x)
	}

	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.SetMoid(x)
	}

	if v, ok := d.GetOk("node_info"); ok {
		x := make([]models.ApplianceNodeIpInfo, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := &models.ApplianceNodeIpInfo{}
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
			o.SetClassId("appliance.NodeIpInfo")
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			x = append(x, *o)
		}
		o.SetNodeInfo(x)
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

	if v, ok := d.GetOkExists("previous_install_attempts"); ok {
		x := int64(v.(int))
		o.SetPreviousInstallAttempts(x)
	}

	if v, ok := d.GetOk("requires"); ok {
		p := make([]models.ApplianceUpgradeRelationship, 0, 1)
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
			p = append(p, models.MoMoRefAsApplianceUpgradeRelationship(o))
		}
		if len(p) > 0 {
			x := p[0]
			o.SetRequires(x)
		}
	}

	if v, ok := d.GetOkExists("rollback_needed"); ok {
		x := (v.(bool))
		o.SetRollbackNeeded(x)
	}

	if v, ok := d.GetOk("rollback_phases"); ok {
		x := make([]models.OnpremUpgradePhase, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := &models.OnpremUpgradePhase{}
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
			o.SetClassId("onprem.UpgradePhase")
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			x = append(x, *o)
		}
		o.SetRollbackPhases(x)
	}

	if v, ok := d.GetOk("rollback_status"); ok {
		x := (v.(string))
		o.SetRollbackStatus(x)
	}

	if v, ok := d.GetOk("services"); ok {
		x := make([]string, 0)
		y := reflect.ValueOf(v)
		for i := 0; i < y.Len(); i++ {
			if y.Index(i).Interface() != nil {
				x = append(x, y.Index(i).Interface().(string))
			}
		}
		o.SetServices(x)
	}

	if v, ok := d.GetOk("shared_scope"); ok {
		x := (v.(string))
		o.SetSharedScope(x)
	}

	if v, ok := d.GetOk("start_time"); ok {
		x, _ := time.Parse(time.RFC1123, v.(string))
		o.SetStartTime(x)
	}

	if v, ok := d.GetOk("status"); ok {
		x := (v.(string))
		o.SetStatus(x)
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

	if v, ok := d.GetOkExists("total_nodes"); ok {
		x := int64(v.(int))
		o.SetTotalNodes(x)
	}

	if v, ok := d.GetOkExists("total_phases"); ok {
		x := int64(v.(int))
		o.SetTotalPhases(x)
	}

	if v, ok := d.GetOk("ui_packages"); ok {
		x := make([]string, 0)
		y := reflect.ValueOf(v)
		for i := 0; i < y.Len(); i++ {
			if y.Index(i).Interface() != nil {
				x = append(x, y.Index(i).Interface().(string))
			}
		}
		o.SetUiPackages(x)
	}

	if v, ok := d.GetOk("nr_version"); ok {
		x := (v.(string))
		o.SetVersion(x)
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

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of ApplianceUpgrade object failed with error : %s", err.Error())
	}
	countResponse, _, responseErr := conn.ApiClient.ApplianceApi.GetApplianceUpgradeList(conn.ctx).Filter(getRequestParams(data)).Count(true).Execute()
	if responseErr != nil {
		errorType := fmt.Sprintf("%T", responseErr)
		if strings.Contains(errorType, "GenericOpenAPIError") {
			responseErr := responseErr.(*models.GenericOpenAPIError)
			return diag.Errorf("error occurred while fetching count of ApplianceUpgrade: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		return diag.Errorf("error occurred while fetching count of ApplianceUpgrade: %s", responseErr.Error())
	}
	count := countResponse.MoDocumentCount.GetCount()
	if count == 0 {
		return diag.Errorf("your query for ApplianceUpgrade data source did not return any results. Please change your search criteria and try again")
	}
	var i int32
	var applianceUpgradeResults = make([]map[string]interface{}, 0, 0)
	for i = 0; i < count; i += 100 {
		resMo, _, responseErr := conn.ApiClient.ApplianceApi.GetApplianceUpgradeList(conn.ctx).Filter(getRequestParams(data)).Top(100).Skip(i).Execute()
		if responseErr != nil {
			errorType := fmt.Sprintf("%T", responseErr)
			if strings.Contains(errorType, "GenericOpenAPIError") {
				responseErr := responseErr.(*models.GenericOpenAPIError)
				return diag.Errorf("error occurred while fetching ApplianceUpgrade: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
			}
			return diag.Errorf("error occurred while fetching ApplianceUpgrade: %s", responseErr.Error())
		}
		results := resMo.ApplianceUpgradeList.GetResults()
		switch reflect.TypeOf(results).Kind() {
		case reflect.Slice:
			for k := 0; k < len(results); k++ {
				var s = results[k]
				var temp = make(map[string]interface{})

				temp["account"] = flattenMapIamAccountRelationship(s.GetAccount(), d)
				temp["account_moid"] = (s.GetAccountMoid())
				temp["active"] = (s.GetActive())
				temp["additional_properties"] = flattenAdditionalProperties(s.AdditionalProperties)
				temp["all_nodes_pingable"] = (s.GetAllNodesPingable())

				temp["ancestors"] = flattenListMoBaseMoRelationship(s.GetAncestors(), d)
				temp["auto_created"] = (s.GetAutoCreated())
				temp["class_id"] = (s.GetClassId())

				temp["completed_phases"] = flattenListOnpremUpgradePhase(s.GetCompletedPhases(), d)

				temp["create_time"] = (s.GetCreateTime()).String()

				temp["current_phase"] = flattenMapOnpremUpgradePhase(s.GetCurrentPhase(), d)
				temp["description"] = (s.GetDescription())
				temp["domain_group_moid"] = (s.GetDomainGroupMoid())
				temp["elapsed_time"] = (s.GetElapsedTime())

				temp["end_time"] = (s.GetEndTime()).String()
				temp["error_code"] = (s.GetErrorCode())
				temp["fingerprint"] = (s.GetFingerprint())

				temp["image_bundle"] = flattenMapApplianceImageBundleRelationship(s.GetImageBundle(), d)
				temp["is_rolling_back"] = (s.GetIsRollingBack())
				temp["is_user_triggered"] = (s.GetIsUserTriggered())
				temp["messages"] = (s.GetMessages())

				temp["mod_time"] = (s.GetModTime()).String()
				temp["moid"] = (s.GetMoid())

				temp["node_info"] = flattenListApplianceNodeIpInfo(s.GetNodeInfo(), d)
				temp["object_type"] = (s.GetObjectType())
				temp["owners"] = (s.GetOwners())

				temp["parent"] = flattenMapMoBaseMoRelationship(s.GetParent(), d)

				temp["permission_resources"] = flattenListMoBaseMoRelationship(s.GetPermissionResources(), d)
				temp["previous_install_attempts"] = (s.GetPreviousInstallAttempts())

				temp["requires"] = flattenMapApplianceUpgradeRelationship(s.GetRequires(), d)
				temp["rollback_needed"] = (s.GetRollbackNeeded())

				temp["rollback_phases"] = flattenListOnpremUpgradePhase(s.GetRollbackPhases(), d)
				temp["rollback_status"] = (s.GetRollbackStatus())
				temp["services"] = (s.GetServices())
				temp["shared_scope"] = (s.GetSharedScope())

				temp["start_time"] = (s.GetStartTime()).String()
				temp["status"] = (s.GetStatus())

				temp["tags"] = flattenListMoTag(s.GetTags(), d)
				temp["total_nodes"] = (s.GetTotalNodes())
				temp["total_phases"] = (s.GetTotalPhases())
				temp["ui_packages"] = (s.GetUiPackages())
				temp["nr_version"] = (s.GetVersion())

				temp["version_context"] = flattenMapMoVersionContext(s.GetVersionContext(), d)
				applianceUpgradeResults = append(applianceUpgradeResults, temp)
			}
		}
	}
	log.Println("length of results: ", len(applianceUpgradeResults))
	if err := d.Set("results", applianceUpgradeResults); err != nil {
		return diag.Errorf("error occurred while setting results: %s", err.Error())
	}
	d.SetId(applianceUpgradeResults[0]["moid"].(string))
	return de
}
