# CEX CLI

A command-line interface for interacting with the CEX cryptocurrency exchange.

## Features

- Market data retrieval
- Account balance information
- Trading operations (place/cancel/amend order, set leverage, set position mode)
- Account management
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

> **注意：所有需要登录的操作（如下单、改单、查余额等）都必须正确配置 API Key/Secret/Passphrase，否则会报 `50103` 错误。
> 
> 示例报错：
> `{"msg":"Request header OK-ACCESS-KEY can not be empty.","code":"50103"}`

## CLI 命令结构

```bash
cexcli [command] [subcommand] [flags]
```

**主命令：**
- `market`    市场行情相关
- `account`   账户相关操作（余额、持仓、配置、杠杆、持仓模式等）
- `trade`     交易相关（下单、改单、撤单、查单、挂单列表）

## 全局参数

| 参数            | 说明                        |
|-----------------|----------------------------|
| --api-key       | CEX API key                |
| --api-secret    | CEX API secret             |
| --passphrase    | CEX API passphrase         |
| --config        | 指定配置文件路径            |
| --simple        | 简洁输出模式                |
| --simulated     | 使用模拟盘                  |

## 主要命令与示例

### 市场行情

```bash
# 获取单个产品行情
cexcli market ticker BTC-USDT

# 获取所有产品行情
cexcli market tickers
```

### 账户相关

```bash
# 查询账户余额
cexcli account balance

# 查询账户持仓
cexcli account positions

# 查询账户配置
cexcli account config --simple

# 设置杠杆
cexcli account leverage --instId BTC-USDT-SWAP --lever 10 --mgnMode cross

# 设置持仓模式
cexcli account position-mode --posMode long_short_mode
```

### 交易相关

```bash
# 下单
cexcli trade order --instId BTC-USDT-SWAP --tdMode cross --side buy --ordType limit --sz 1 --px 30000

# 撤单
cexcli trade cancel --instId BTC-USDT-SWAP --ordId 123456

# 查询订单
cexcli trade get-order --instId BTC-USDT-SWAP --ordId 123456

# 查询挂单
cexcli trade pending --instId BTC-USDT-SWAP

# 修改订单
cexcli trade amend --instId BTC-USDT-SWAP --ordId 123456 --newPx 31000
```

## 获取帮助

```bash
cexcli --help
cexcli trade --help
cexcli trade order --help
```

**如需更多命令参数说明，请用 `--help` 查看详细用法。**

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request. 