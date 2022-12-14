// Code generated by codegen; DO NOT EDIT.

package compute

import (
	"context"
	"google.golang.org/api/iterator"

	pb "google.golang.org/genproto/googleapis/cloud/compute/v1"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugins/source/gcp/client"

	"cloud.google.com/go/compute/apiv1"
)

func SslPolicies() *schema.Table {
	return &schema.Table{
		Name:      "gcp_compute_ssl_policies",
		Resolver:  fetchSslPolicies,
		Multiplex: client.ProjectMultiplex,
		Columns: []schema.Column{
			{
				Name:     "project_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveProject,
			},
			{
				Name:     "creation_timestamp",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("CreationTimestamp"),
			},
			{
				Name:     "custom_features",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("CustomFeatures"),
			},
			{
				Name:     "description",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Description"),
			},
			{
				Name:     "enabled_features",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("EnabledFeatures"),
			},
			{
				Name:     "fingerprint",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Fingerprint"),
			},
			{
				Name:     "id",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("Id"),
			},
			{
				Name:     "kind",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Kind"),
			},
			{
				Name:     "min_tls_version",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("MinTlsVersion"),
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "profile",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Profile"),
			},
			{
				Name:     "region",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Region"),
			},
			{
				Name:     "self_link",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SelfLink"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "warnings",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Warnings"),
			},
		},
	}
}

func fetchSslPolicies(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	req := &pb.ListSslPoliciesRequest{
		Project: c.ProjectId,
	}
	gcpClient, err := compute.NewSslPoliciesRESTClient(ctx, c.ClientOptions...)
	if err != nil {
		return err
	}
	it := gcpClient.List(ctx, req, c.CallOptions...)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return err
		}

		res <- resp

	}
	return nil
}
