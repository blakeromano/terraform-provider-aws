// Code generated by internal/generate/servicepackage/main.go; DO NOT EDIT.

package apprunner

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/apprunner"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/types"
	"github.com/hashicorp/terraform-provider-aws/names"
)

type servicePackage struct{}

func (p *servicePackage) FrameworkDataSources(ctx context.Context) []*types.ServicePackageFrameworkDataSource {
	return []*types.ServicePackageFrameworkDataSource{
		{
			Factory:  newHostedZoneIDDataSource,
			TypeName: "aws_apprunner_hosted_zone_id",
			Name:     "Hosted Zone ID",
		},
	}
}

func (p *servicePackage) FrameworkResources(ctx context.Context) []*types.ServicePackageFrameworkResource {
	return []*types.ServicePackageFrameworkResource{
		{
			Factory:  newResourceDefaultAutoScalingConfigurationVersion,
			TypeName: "aws_apprunner_default_auto_scaling_configuration_version",
			Name:     "Default AutoScaling Configuration Version",
		},
		{
			Factory:  newDeploymentResource,
			TypeName: "aws_apprunner_deployment",
			Name:     "Deployment",
		},
	}
}

func (p *servicePackage) SDKDataSources(ctx context.Context) []*types.ServicePackageSDKDataSource {
	return []*types.ServicePackageSDKDataSource{}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  resourceAutoScalingConfigurationVersion,
			TypeName: "aws_apprunner_auto_scaling_configuration_version",
			Name:     "AutoScaling Configuration Version",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceConnection,
			TypeName: "aws_apprunner_connection",
			Name:     "Connection",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceCustomDomainAssociation,
			TypeName: "aws_apprunner_custom_domain_association",
			Name:     "Custom Domain Association",
		},
		{
			Factory:  resourceObservabilityConfiguration,
			TypeName: "aws_apprunner_observability_configuration",
			Name:     "Observability Configuration",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceService,
			TypeName: "aws_apprunner_service",
			Name:     "Service",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceVPCConnector,
			TypeName: "aws_apprunner_vpc_connector",
			Name:     "VPC Connector",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceVPCIngressConnection,
			TypeName: "aws_apprunner_vpc_ingress_connection",
			Name:     "VPC Ingress Connection",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.AppRunner
}

// NewClient returns a new AWS SDK for Go v2 client for this service package's AWS API.
func (p *servicePackage) NewClient(ctx context.Context, config map[string]any) (*apprunner.Client, error) {
	cfg := *(config["aws_sdkv2_config"].(*aws.Config))
	optFns := []func(*apprunner.Options){
		apprunner.WithEndpointResolverV2(newEndpointResolverV2()),
		withBaseEndpoint(config[names.AttrEndpoint].(string)),
		withExtraOptions(ctx, p, config),
	}

	return apprunner.NewFromConfig(cfg, optFns...), nil
}

// withExtraOptions returns a functional option that allows this service package to specify extra API client options.
// This option is always called after any generated options.
func withExtraOptions(ctx context.Context, sp conns.ServicePackage, config map[string]any) func(*apprunner.Options) {
	if v, ok := sp.(interface {
		withExtraOptions(context.Context, map[string]any) []func(*apprunner.Options)
	}); ok {
		optFns := v.withExtraOptions(ctx, config)

		return func(o *apprunner.Options) {
			for _, optFn := range optFns {
				optFn(o)
			}
		}
	}

	return func(*apprunner.Options) {}
}

func ServicePackage(ctx context.Context) conns.ServicePackage {
	return &servicePackage{}
}
