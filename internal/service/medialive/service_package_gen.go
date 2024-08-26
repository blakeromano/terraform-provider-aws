// Code generated by internal/generate/servicepackage/main.go; DO NOT EDIT.

package medialive

import (
	"context"

	aws_sdkv2 "github.com/aws/aws-sdk-go-v2/aws"
	medialive_sdkv2 "github.com/aws/aws-sdk-go-v2/service/medialive"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/types"
	"github.com/hashicorp/terraform-provider-aws/names"
)

type servicePackage struct{}

func (p *servicePackage) FrameworkDataSources(ctx context.Context) []*types.ServicePackageFrameworkDataSource {
	return []*types.ServicePackageFrameworkDataSource{
		{
			Factory: newDataSourceInput,
			Name:    "Input",
		},
	}
}

func (p *servicePackage) FrameworkResources(ctx context.Context) []*types.ServicePackageFrameworkResource {
	return []*types.ServicePackageFrameworkResource{
		{
			Factory: newResourceMultiplexProgram,
		},
	}
}

func (p *servicePackage) SDKDataSources(ctx context.Context) []*types.ServicePackageSDKDataSource {
	return []*types.ServicePackageSDKDataSource{}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  ResourceChannel,
			TypeName: "aws_medialive_channel",
			Name:     "Channel",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  ResourceInput,
			TypeName: "aws_medialive_input",
			Name:     "Input",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  ResourceInputSecurityGroup,
			TypeName: "aws_medialive_input_security_group",
			Name:     "Input Security Group",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  ResourceMultiplex,
			TypeName: "aws_medialive_multiplex",
			Name:     "Multiplex",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.MediaLive
}

// NewClient returns a new AWS SDK for Go v2 client for this service package's AWS API.
func (p *servicePackage) NewClient(ctx context.Context, config map[string]any) (*medialive_sdkv2.Client, error) {
	cfg := *(config["aws_sdkv2_config"].(*aws_sdkv2.Config))

	return medialive_sdkv2.NewFromConfig(cfg,
		medialive_sdkv2.WithEndpointResolverV2(newEndpointResolverSDKv2()),
		withBaseEndpoint(config[names.AttrEndpoint].(string)),
	), nil
}

func ServicePackage(ctx context.Context) conns.ServicePackage {
	return &servicePackage{}
}
