# Enum property for integer
resource "intersight_sol_policy" "enum_int" {
    baud_rate = 10000
}


# Enum property for string
resource "intersight_workflow_workflow_info" "enum_string" {
    action = "retryfailed"
}

# Maximum integer
resource "intersight_hyperflex_cluster_replication_network_policy" "max_int"{
	replication_bandwidth_mbps = 100001
}


# Minimum integer
resource "intersight_kubernetes_node_group_profile" "min_int"{
    maxsize = 0
}

# Range Minimum Integer
resource "intersight_hyperflex_cluster_backup_policy" "min_range_int"{
    snapshot_retention_count = 0
}


#Range Maximum Integer
resource "intersight_hyperflex_cluster_backup_policy" "max_range_int"{
    snapshot_retention_count = 31
}


#Maximum Float
resource "intersight_tam_security_advisory" "max_float"{
    base_score = 11
}

#Minimum length string
resource "intersight_virtualization_virtual_disk" "min_length_string"{
     capacity = ""
}

#Maximum length string
resource "intersight_snmp_policy" "max_length_string"{
     access_community_string = "testinggggggggggggggggggggggggggggggg"
}

#String length range minimum
resource "intersight_virtualization_virtual_machine" "min_length_range_string"{
       name = ""
}

#String length range maximum
resource "intersight_virtualization_virtual_machine" "max_length_range_string"{
    name = "testingg_stringggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggg"
}

#Match pattern
resource "intersight_hyperflex_sys_config_policy" "string_match_pattern"{
     dns_servers = ["testing", "soemthing"]
}

#Match pattern, Max length string
resource "intersight_hyperflex_sys_config_policy" "match_pattern_max_length"{
     dns_domain_name = "~`tetsingngggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggg"
}

#Match pattern, Max range
resource "intersight_iam_end_point_user_role" "match_pattern_max_range"{
     password = "~`testingg_stringggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggg"
}

#Match pattern, Min range
resource "intersight_iam_end_point_user_role" "match_pattern_min_range"{
     password = "~`test"
}

#Min items
resource "intersight_softwarerepository_release" "min_items" {
  supported_models = []
}

#Max items
resource "intersight_ntp_policy" "max_items" {
  ntp_servers = ["10.10.10.10", "10.10.10.11", "10.10.10.12", "10.10.10.13", "10.10.10.14"]
}

# Please add new resources above this line. Also, add the new resource added into this file to "resource_list.txt" file. The file will be read for testing each resource separately in Jenkins job

