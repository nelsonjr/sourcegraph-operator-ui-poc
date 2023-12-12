# Operator Maintenance UI

## Components

This project contains the following components:

### Maintenance UI

A React + Material UI application that communicates with the Operator and gathers data and display status.

Features:

- Installation
- Health & Actions
- Upgrade

### Mock Operator API

In the [mock-api](./mock-api/) folder, a Go Server application that implements the Operator API companion to the Maintenance UI.

#### Mock Operator Debug Bar API

We also implement some test APIs to enable controlling the Mock Operator from the Maitenance UI.

## Running

1. Run the go application in the `mock-api` folder:

   ```
   cd mock-api
   go run ./cmd
   ```

2. Run the Maitenance UI:

   ```
   pnpm run dev
   ```
