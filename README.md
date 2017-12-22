# CLF - CLI

> Cloudflare CLI, you can use this cli for manage your resource on cloudflare CDN

### Usage

```
➜  ~ clf --help
This CLI will help you manage Cloudflare resource, such as DNS record, Zone..
	Example:
		clf dns create record dns --type A --address --zone example.com --enable-proxy true

Usage:
  clf [command]

Available Commands:
  config      Manage clf config on your local machine
  dns         A Cloudflare DNS manager CLI
  help        Help about any command
  zone        Action with zones on cloudflare

Flags:
  -h, --help   help for clf

Use "clf [command] --help" for more information about a command.
```
### Requirements

- OS installed

### Contributing

- `CLF CLI` use [Cloudflare SDK](https://github.com/cloudflare/cloudflare-go)

- If you want to modify or custom some features, please send me a pull request. Thanks.

### License
[Apache License](LICENSE)