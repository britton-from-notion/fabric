package microsoft

import (
	"context"
	"fmt"

	"github.com/blackstork-io/fabric/internal/microsoft/client"
	"github.com/blackstork-io/fabric/plugin"
	"github.com/blackstork-io/fabric/plugin/dataspec"
)

type ClientLoadFn func() client.Client

var DefaultClientLoader ClientLoadFn = client.New

func Plugin(version string, loader ClientLoadFn) *plugin.Schema {
	if loader == nil {
		loader = DefaultClientLoader
	}
	return &plugin.Schema{
		Doc:     "The `microsoft` plugin for Microsoft services.",
		Name:    "blackstork/microsoft",
		Version: version,
		DataSources: plugin.DataSources{
			"microsoft_sentinel_incidents": makeMicrosoftSentinelIncidentsDataSource(loader),
		},
	}
}

func makeClient(ctx context.Context, loader ClientLoadFn, cfg *dataspec.Block) (client.Client, error) {
	if cfg == nil {
		return nil, fmt.Errorf("configuration is required")
	}
	cli := loader()
	res, err := cli.GetClientCredentialsToken(ctx, &client.GetClientCredentialsTokenReq{
		TenantID:     cfg.GetAttrVal("tenant_id").AsString(),
		ClientID:     cfg.GetAttrVal("client_id").AsString(),
		ClientSecret: cfg.GetAttrVal("client_secret").AsString(),
	})
	if err != nil {
		return nil, err
	}
	cli.UseAuth(res.AccessToken)
	return cli, nil
}
