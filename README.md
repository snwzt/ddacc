# Download Accelerator CLI

A stupid simple download accelerator that I wrote to learn Go concepts better.

Building the binary:

```sh
git clone https://github.com/snwzt/ddacc.git
cd ddacc
go mod tidy
make test # for tests
make build # this will create the ddacc binary in /bin directory 
```

| Usage                  | Command                                                      |
| ---------------------- | ------------------------------------------------------------ |
| Download a single file | `ddacc url <file-url>`                                     |
| Batch downloads        | `ddacc readfile <path-to-file>`                            |
| Help                   | `ddacc help`                                               |
| Help for subcommands   | `ddacc <subcommand> help` or `ddacc <subcommand> --help` |
| Shell completion       | `ddacc completion`                                         |

Number of connections and download path can be configured with `--connections` or `-c` and `--directory` or `-d` respectively.