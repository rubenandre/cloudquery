---
name: Typeform
stage: Preview
title: Typeform Source Plugin
description: CloudQuery Typeform source plugin documentation
---
# Typeform Source Plugin

:badge

The CloudQuery Typeform plugin pulls data from [Typeform](https://www.typeform.com/) and loads it into any supported CloudQuery destination (e.g. PostgreSQL, BigQuery, Snowflake, and [more](https://hub.cloudquery.io/plugins/destination)).

See [tables](/docs/plugins/sources/typeform/tables) for a list of resources supported.

## Example Configuration

This example syncs from Typeform to a Postgres destination. The (top level) source spec section is described in the [Source Spec Reference](/docs/reference/source-spec).

:configuration

## Authentication

:authentication

## Configuration Reference

This is the (nested) spec used by the Typeform source plugin:

- `access_token` (string, required):

  Your [personal access token](https://www.typeform.com/developers/get-started/personal-access-token/) from the Typeform Dashboard.

- `base_url` (string, optional. Default: `https://api.typeform.com`):

  The base URL to fetch data from. Use https://api.eu.typeform.com if your account is stored in the EU.

- `concurrency` (integer, optional. Default: 100):

  Maximum number of requests to perform concurrently.

- `queue_size` (integer, optional. Default: 10000):

  Maximum number of items to have in the queue before waiting for an unfinished request to complete.
