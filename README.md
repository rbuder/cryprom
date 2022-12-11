## Synopsis

This is a prometheus exporter which tracks arbitrary coins on CoinCap's API

## Usage

```
export COINS="bitcoin,ethereum"
./cryprom
```

Get all supported coins from here:
```
curl --location --request GET 'api.coincap.io/v2/assets' | jq -r '.data | .[].id
```

## Sample
```
# HELP current_rate_coin_usd Current exchange rate
# TYPE current_rate_coin_usd gauge
current_rate_coin_usd{coin="bitcoin"} 17182.408525906096
current_rate_coin_usd{coin="ethereum"} 1274.0311497166313
```

## Thanks

<https://github.com/metonymic-smokey/prom-instrumentation>
