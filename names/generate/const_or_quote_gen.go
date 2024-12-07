// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
// Code generated by internal/generate/attrconsts/main.go; DO NOT EDIT.

package names

import (
	"fmt"
)

// ConstOrQuote returns the constant name for the given attribute if it exists.
// Otherwise, it returns the attribute quoted. This is intended for use in
// generated code and templates.
func ConstOrQuote(constant string) string {
	allConstants := map[string]string{
		"arn":                           "AttrARN",
		"arns":                          "AttrARNs",
		"aws_account_id":                "AttrAWSAccountID",
		"access_key":                    "AttrAccessKey",
		"account_id":                    "AttrAccountID",
		"action":                        "AttrAction",
		"actions":                       "AttrActions",
		"address":                       "AttrAddress",
		"alias":                         "AttrAlias",
		"allocated_storage":             "AttrAllocatedStorage",
		"allow_major_version_upgrade":   "AttrAllowMajorVersionUpgrade",
		"application_id":                "AttrApplicationID",
		"apply_immediately":             "AttrApplyImmediately",
		"association_id":                "AttrAssociationID",
		"attributes":                    "AttrAttributes",
		"auto_minor_version_upgrade":    "AttrAutoMinorVersionUpgrade",
		"availability_zone":             "AttrAvailabilityZone",
		"availability_zones":            "AttrAvailabilityZones",
		"bucket":                        "AttrBucket",
		"bucket_name":                   "AttrBucketName",
		"bucket_prefix":                 "AttrBucketPrefix",
		"cidr_block":                    "AttrCIDRBlock",
		"capacity_provider_strategy":    "AttrCapacityProviderStrategy",
		"catalog_id":                    "AttrCatalogID",
		"certificate":                   "AttrCertificate",
		"certificate_arn":               "AttrCertificateARN",
		"certificate_chain":             "AttrCertificateChain",
		"client_id":                     "AttrClientID",
		"client_secret":                 "AttrClientSecret",
		"cloudwatch_log_group_arn":      "AttrCloudWatchLogGroupARN",
		"cloudwatch_logs":               "AttrCloudWatchLogs",
		"cluster_identifier":            "AttrClusterIdentifier",
		"cluster_name":                  "AttrClusterName",
		"comment":                       "AttrComment",
		"condition":                     "AttrCondition",
		"configuration":                 "AttrConfiguration",
		"connection_id":                 "AttrConnectionID",
		"content":                       "AttrContent",
		"content_type":                  "AttrContentType",
		"create_time":                   "AttrCreateTime",
		"created_at":                    "AttrCreatedAt",
		"created_date":                  "AttrCreatedDate",
		"created_time":                  "AttrCreatedTime",
		"creation_date":                 "AttrCreationDate",
		"creation_time":                 "AttrCreationTime",
		"dns_name":                      "AttrDNSName",
		"database":                      "AttrDatabase",
		"database_name":                 "AttrDatabaseName",
		"default_action":                "AttrDefaultAction",
		"default_value":                 "AttrDefaultValue",
		"delete_on_termination":         "AttrDeleteOnTermination",
		"deletion_protection":           "AttrDeletionProtection",
		"description":                   "AttrDescription",
		"destination":                   "AttrDestination",
		"destination_arn":               "AttrDestinationARN",
		"device_name":                   "AttrDeviceName",
		"display_name":                  "AttrDisplayName",
		"domain":                        "AttrDomain",
		"domain_name":                   "AttrDomainName",
		"duration":                      "AttrDuration",
		"email":                         "AttrEmail",
		"enabled":                       "AttrEnabled",
		"encrypted":                     "AttrEncrypted",
		"encryption_configuration":      "AttrEncryptionConfiguration",
		"endpoint":                      "AttrEndpoint",
		"endpoint_type":                 "AttrEndpointType",
		"endpoints":                     "AttrEndpoints",
		"engine":                        "AttrEngine",
		"engine_version":                "AttrEngineVersion",
		"environment":                   "AttrEnvironment",
		"execution_role_arn":            "AttrExecutionRoleARN",
		"expected_bucket_owner":         "AttrExpectedBucketOwner",
		"expression":                    "AttrExpression",
		"external_id":                   "AttrExternalID",
		"family":                        "AttrFamily",
		"field":                         "AttrField",
		"file_system_id":                "AttrFileSystemID",
		"filter":                        "AttrFilter",
		"final_snapshot_identifier":     "AttrFinalSnapshotIdentifier",
		"force_delete":                  "AttrForceDelete",
		"force_destroy":                 "AttrForceDestroy",
		"format":                        "AttrFormat",
		"function_arn":                  "AttrFunctionARN",
		"group_name":                    "AttrGroupName",
		"header":                        "AttrHeader",
		"health_check":                  "AttrHealthCheck",
		"hosted_zone_id":                "AttrHostedZoneID",
		"iam_role_arn":                  "AttrIAMRoleARN",
		"id":                            "AttrID",
		"ids":                           "AttrIDs",
		"iops":                          "AttrIOPS",
		"ip_address":                    "AttrIPAddress",
		"ip_address_type":               "AttrIPAddressType",
		"ip_addresses":                  "AttrIPAddresses",
		"identifier":                    "AttrIdentifier",
		"instance_count":                "AttrInstanceCount",
		"instance_id":                   "AttrInstanceID",
		"instance_type":                 "AttrInstanceType",
		"interval":                      "AttrInterval",
		"issuer":                        "AttrIssuer",
		"json":                          "AttrJSON",
		"kms_key":                       "AttrKMSKey",
		"kms_key_arn":                   "AttrKMSKeyARN",
		"kms_key_id":                    "AttrKMSKeyID",
		"key":                           "AttrKey",
		"key_id":                        "AttrKeyID",
		"language_code":                 "AttrLanguageCode",
		"last_updated_date":             "AttrLastUpdatedDate",
		"last_updated_time":             "AttrLastUpdatedTime",
		"launch_template":               "AttrLaunchTemplate",
		"location":                      "AttrLocation",
		"log_group_name":                "AttrLogGroupName",
		"logging_configuration":         "AttrLoggingConfiguration",
		"max":                           "AttrMax",
		"max_capacity":                  "AttrMaxCapacity",
		"message":                       "AttrMessage",
		"metric_name":                   "AttrMetricName",
		"min":                           "AttrMin",
		"mode":                          "AttrMode",
		"most_recent":                   "AttrMostRecent",
		"name":                          "AttrName",
		"name_prefix":                   "AttrNamePrefix",
		"name_suffix":                   "AttrNameSuffix",
		"names":                         "AttrNames",
		"namespace":                     "AttrNamespace",
		"network_configuration":         "AttrNetworkConfiguration",
		"network_interface_id":          "AttrNetworkInterfaceID",
		"owner":                         "AttrOwner",
		"owner_account_id":              "AttrOwnerAccountID",
		"owner_id":                      "AttrOwnerID",
		"parameter":                     "AttrParameter",
		"parameter_group_name":          "AttrParameterGroupName",
		"parameters":                    "AttrParameters",
		"password":                      "AttrPassword",
		"path":                          "AttrPath",
		"permissions":                   "AttrPermissions",
		"policy":                        "AttrPolicy",
		"port":                          "AttrPort",
		"preferred_maintenance_window":  "AttrPreferredMaintenanceWindow",
		"prefix":                        "AttrPrefix",
		"principal":                     "AttrPrincipal",
		"priority":                      "AttrPriority",
		"private_key":                   "AttrPrivateKey",
		"profile":                       "AttrProfile",
		"propagate_tags":                "AttrPropagateTags",
		"properties":                    "AttrProperties",
		"protocol":                      "AttrProtocol",
		"provider_name":                 "AttrProviderName",
		"public_key":                    "AttrPublicKey",
		"publicly_accessible":           "AttrPubliclyAccessible",
		"region":                        "AttrRegion",
		"repository_name":               "AttrRepositoryName",
		"resource_arn":                  "AttrResourceARN",
		"resource_id":                   "AttrResourceID",
		"resource_owner":                "AttrResourceOwner",
		"resource_tags":                 "AttrResourceTags",
		"resource_type":                 "AttrResourceType",
		"resources":                     "AttrResources",
		"retention_period":              "AttrRetentionPeriod",
		"role":                          "AttrRole",
		"role_arn":                      "AttrRoleARN",
		"rule":                          "AttrRule",
		"s3_bucket":                     "AttrS3Bucket",
		"s3_bucket_name":                "AttrS3BucketName",
		"s3_key_prefix":                 "AttrS3KeyPrefix",
		"sns_topic_arn":                 "AttrSNSTopicARN",
		"schedule":                      "AttrSchedule",
		"schedule_expression":           "AttrScheduleExpression",
		"schema":                        "AttrSchema",
		"scope":                         "AttrScope",
		"secret_key":                    "AttrSecretKey",
		"security_group_ids":            "AttrSecurityGroupIDs",
		"security_groups":               "AttrSecurityGroups",
		"service_name":                  "AttrServiceName",
		"service_role":                  "AttrServiceRole",
		"service_role_arn":              "AttrServiceRoleARN",
		"session":                       "AttrSession",
		"shared_config_files":           "AttrSharedConfigFiles",
		"size":                          "AttrSize",
		"skip_credentials_validation":   "AttrSkipCredentialsValidation",
		"skip_destroy":                  "AttrSkipDestroy",
		"skip_requesting_account_id":    "AttrSkipRequestingAccountID",
		"snapshot_id":                   "AttrSnapshotID",
		"source":                        "AttrSource",
		"source_type":                   "AttrSourceType",
		"stage":                         "AttrStage",
		"start_time":                    "AttrStartTime",
		"state":                         "AttrState",
		"status":                        "AttrStatus",
		"status_code":                   "AttrStatusCode",
		"status_message":                "AttrStatusMessage",
		"status_reason":                 "AttrStatusReason",
		"storage_class":                 "AttrStorageClass",
		"storage_encrypted":             "AttrStorageEncrypted",
		"storage_type":                  "AttrStorageType",
		"stream_arn":                    "AttrStreamARN",
		"subnet_id":                     "AttrSubnetID",
		"subnet_ids":                    "AttrSubnetIDs",
		"subnets":                       "AttrSubnets",
		"suffix":                        "AttrSuffix",
		"table_name":                    "AttrTableName",
		"tags":                          "AttrTags",
		"tags_all":                      "AttrTagsAll",
		"target":                        "AttrTarget",
		"target_arn":                    "AttrTargetARN",
		"throughput":                    "AttrThroughput",
		"timeout":                       "AttrTimeout",
		"timeouts":                      "AttrTimeouts",
		"topic_arn":                     "AttrTopicARN",
		"transit_gateway_attachment_id": "AttrTransitGatewayAttachmentID",
		"transit_gateway_id":            "AttrTransitGatewayID",
		"triggers":                      "AttrTriggers",
		"type":                          "AttrType",
		"uri":                           "AttrURI",
		"url":                           "AttrURL",
		"unit":                          "AttrUnit",
		"user_name":                     "AttrUserName",
		"user_pool_id":                  "AttrUserPoolID",
		"username":                      "AttrUsername",
		"vpc_config":                    "AttrVPCConfig",
		"vpc_configuration":             "AttrVPCConfiguration",
		"vpc_endpoint_id":               "AttrVPCEndpointID",
		"vpc_id":                        "AttrVPCID",
		"vpc_security_group_ids":        "AttrVPCSecurityGroupIDs",
		"value":                         "AttrValue",
		"values":                        "AttrValues",
		"version":                       "AttrVersion",
		"virtual_name":                  "AttrVirtualName",
		"volume_size":                   "AttrVolumeSize",
		"volume_type":                   "AttrVolumeType",
		"weight":                        "AttrWeight",
	}

	if v, ok := allConstants[constant]; ok {
		return fmt.Sprintf("names.%s", v)
	}
	return fmt.Sprintf("%q", constant)
}
