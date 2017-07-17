# ntwrk
CLI and server for testing network performance. Supports macOS and 64-bit Linux.

## Installation
### Client
1. Download the latest release for your platform from the [releases](https://github.com/waits/ntwrk/releases) page.
2. `# install ntwrk-darwin-amd64 /usr/local/bin/ntwrk` (adjust for your platform)

### Server (for Linux with systemd)
1. Download the latest `ntwrk-linux-amd64` binary from the [releases](https://github.com/waits/ntwrk/releases) page.
2. `# install ntwrk-linux-amd64 /usr/local/bin/ntwrk`
3. `# curl https://github.com/waits/ntwrk/blob/master/etc/ntwrk.service > /etc/systemd/system/ntwrk.service`
4. `# systemctl enable ntwrk && systemctl start ntwrk`

## Usage
```
usage: ntwrk <command> [arguments]

commands:
    help	Show this help message
    ip		Print external IP address
    server	Start a test server
    test	Run performance tests
    update	Check for and download an updated binary
    version	Print version number
```

The `ip` and `test` commands take a `-host` flag to test against a custom server (default is `ntwrk.waits.io`). The client and server currently only run on port 1600.
