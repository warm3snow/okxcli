# CEX CLI

A command-line interface for interacting with the CEX cryptocurrency exchange.

## Features

- Market data retrieval
- Trading operations (coming soon)
- Account management (coming soon)
- Configuration management with config files and environment variables
- Support for simulated trading

## Installation

```bash
go install github.com/warm3snow/cexcli@latest
```

## Configuration

The CLI can be configured in several ways:

1. Configuration file (`configs/cex.yaml`):
   ```yaml
   cex:
     base_url: "https://www.okx.com"
     api:
       api_key: "YOUR_API_KEY"
       secret_key: "YOUR_SECRET_KEY"
       passphrase: "YOUR_PASSPHRASE"
       is_simulated: false
   ```

2. Command line flags:
   ```bash
   cexcli --api-key=YOUR_API_KEY --api-secret=YOUR_SECRET_KEY --passphrase=YOUR_PASSPHRASE --simulated
   ```

3. Environment variables:
   ```bash
   export CEX_OKEX_API_API_KEY=YOUR_API_KEY
   export CEX_OKEX_API_SECRET_KEY=YOUR_SECRET_KEY
   export CEX_OKEX_API_PASSPHRASE=YOUR_PASSPHRASE
   export CEX_OKEX_API_IS_SIMULATED=true
   ```

The CLI will look for the configuration file in the following locations (in order):
1. Path specified by `--config` flag
2. `./cex.yaml`
3. `./configs/cex.yaml`
4. `$HOME/.cexcli/cex.yaml`
5. `/etc/cexcli/cex.yaml`

## Usage

### Market Data

Get ticker information:
```bash
cexcli market ticker BTC-USDT
```

Use simulated trading mode:
```bash
cexcli --simulated market ticker BTC-USDT
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request. 