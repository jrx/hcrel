# hcrel

A small CLI tool to fetch the latest release version of HashiCorp products.

## Usage

```sh
go build -o hcrel .
./hcrel -product vault -license oss
```

### Flags

| Flag        | Default   | Description                                      |
|-------------|-----------|--------------------------------------------------|
| `-product`  | `vault`   | Product name: `vault`, `consul`, `boundary`, etc |
| `-license`  | `oss`     | License class: `oss`, `enterprise`, `hcp`        |

## Requirements

- Go 1.26+
