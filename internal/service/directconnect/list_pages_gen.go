// Code generated by "internal/generate/listpages/main.go -ListOps=DescribeDirectConnectGateways,DescribeDirectConnectGatewayAssociations,DescribeDirectConnectGatewayAssociationProposals"; DO NOT EDIT.

package directconnect

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/directconnect"
)

func describeDirectConnectGatewayAssociationProposalsPages(conn *directconnect.DirectConnect, input *directconnect.DescribeDirectConnectGatewayAssociationProposalsInput, fn func(*directconnect.DescribeDirectConnectGatewayAssociationProposalsOutput, bool) bool) error {
	return describeDirectConnectGatewayAssociationProposalsPagesWithContext(context.Background(), conn, input, fn)
}

func describeDirectConnectGatewayAssociationProposalsPagesWithContext(ctx context.Context, conn *directconnect.DirectConnect, input *directconnect.DescribeDirectConnectGatewayAssociationProposalsInput, fn func(*directconnect.DescribeDirectConnectGatewayAssociationProposalsOutput, bool) bool) error {
	for {
		output, err := conn.DescribeDirectConnectGatewayAssociationProposalsWithContext(ctx, input)
		if err != nil {
			return err
		}

		lastPage := aws.StringValue(output.NextToken) == ""
		if !fn(output, lastPage) || lastPage {
			break
		}

		input.NextToken = output.NextToken
	}
	return nil
}

func describeDirectConnectGatewayAssociationsPages(conn *directconnect.DirectConnect, input *directconnect.DescribeDirectConnectGatewayAssociationsInput, fn func(*directconnect.DescribeDirectConnectGatewayAssociationsOutput, bool) bool) error {
	return describeDirectConnectGatewayAssociationsPagesWithContext(context.Background(), conn, input, fn)
}

func describeDirectConnectGatewayAssociationsPagesWithContext(ctx context.Context, conn *directconnect.DirectConnect, input *directconnect.DescribeDirectConnectGatewayAssociationsInput, fn func(*directconnect.DescribeDirectConnectGatewayAssociationsOutput, bool) bool) error {
	for {
		output, err := conn.DescribeDirectConnectGatewayAssociationsWithContext(ctx, input)
		if err != nil {
			return err
		}

		lastPage := aws.StringValue(output.NextToken) == ""
		if !fn(output, lastPage) || lastPage {
			break
		}

		input.NextToken = output.NextToken
	}
	return nil
}

func describeDirectConnectGatewaysPages(conn *directconnect.DirectConnect, input *directconnect.DescribeDirectConnectGatewaysInput, fn func(*directconnect.DescribeDirectConnectGatewaysOutput, bool) bool) error {
	return describeDirectConnectGatewaysPagesWithContext(context.Background(), conn, input, fn)
}

func describeDirectConnectGatewaysPagesWithContext(ctx context.Context, conn *directconnect.DirectConnect, input *directconnect.DescribeDirectConnectGatewaysInput, fn func(*directconnect.DescribeDirectConnectGatewaysOutput, bool) bool) error {
	for {
		output, err := conn.DescribeDirectConnectGatewaysWithContext(ctx, input)
		if err != nil {
			return err
		}

		lastPage := aws.StringValue(output.NextToken) == ""
		if !fn(output, lastPage) || lastPage {
			break
		}

		input.NextToken = output.NextToken
	}
	return nil
}
