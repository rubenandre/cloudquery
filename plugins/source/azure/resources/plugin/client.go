package plugin

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/v4/message"
	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/cloudquery/plugin-sdk/v4/scheduler"
	"github.com/cloudquery/plugin-sdk/v4/schema"

	"github.com/rs/zerolog"
)

type Client struct {
	plugin.UnimplementedDestination
	schduler   *scheduler.Scheduler
	syncClient *client.Client
	options    plugin.NewClientOptions
}

func NewClient(ctx context.Context, logger zerolog.Logger, specBytes []byte, options plugin.NewClientOptions) (plugin.Client, error) {
	c := &Client{
		options: options,
	}
	if options.NoConnection {
		return c, nil
	}
	spec := &client.Spec{}
	if err := json.Unmarshal(specBytes, spec); err != nil {
		return nil, err
	}
	spec.SetDefaults()
	syncClient, err := client.New(ctx, logger, spec)
	if err != nil {
		return nil, err
	}
	c.syncClient = syncClient.(*client.Client)
	c.schduler = scheduler.NewScheduler(scheduler.WithLogger(logger), scheduler.WithConcurrency(spec.Concurrency))
	return c, nil
}

func (*Client) Close(_ context.Context) error {
	return nil
}

func (*Client) Tables(_ context.Context, options plugin.TableOptions) (schema.Tables, error) {
	tables := tables()
	tables, err := tables.FilterDfs(options.Tables, options.SkipTables, options.SkipDependentTables)
	if err != nil {
		return nil, err
	}
	return tables, nil
}

func (c *Client) Sync(ctx context.Context, options plugin.SyncOptions, res chan<- message.SyncMessage) error {
	if c.options.NoConnection {
		return fmt.Errorf("no connection")
	}
	tables := tables()
	tables, err := tables.FilterDfs(options.Tables, options.SkipTables, options.SkipDependentTables)
	if err != nil {
		return err
	}
	return c.schduler.Sync(ctx, c.syncClient.Duplicate(), tables, res)
}
