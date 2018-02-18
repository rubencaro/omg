# omg

Because all those helper scripts around your code should be amazing too.

## What

A tool to efficiently manage all those little things/scripts/files around _my_ code that make up for different stages of development. Such as compiling, releasing, deploying, packaging or publishing, but also linting, formatting, testing, benchmarking, etc.

I mean it to be stack independent, so I can use it with all my Javascript/Elixir/Go/Whatever projects. Also compatible with all the shapes and colors of helper scripts and tools I use with them.

## Usage

You just download the compiled binary of the last release for your platform into the root of your project, and then `omg`. That should show you some guidance and create a documented `.omg.toml` file. You fill it in and you are good to go on asking `omg` to do things. `omg help` should help. Things like `omg deploy` or `omg release` or `omg publish` should start to work as you define them.

Both the `omg` binary and the `.omg.toml` file are meant to be included on version control. Both are locked to your code. That's a big fat feature. You may choose to have a global `omg` binary and keep only `.omg.toml` included into VCS though. It's up to you. It that case you just need to put the binary on your system `PATH` (ex. `~/.local/bin`) and you're done. When updating `omg`, or when working from different machines, you may take a look at the version numbers and the [Changelog](#changelog) to ensure everything in your `.omg.toml` will work as expected.

All documentation should be accessible from the `omg help` command and comments on `.omg.toml` itself. You shouldn't need to come back here to look for usage instructions for your particular copy of `omg`. That's another big fat feature.

## What can it do

You could always run `omg help` and see for yourself, and also then self generated `.omg.toml` file, but here is a brief description.

It can resolve a list of servers (name, IPs, etc.) from a fixed list on config, or from a configured GCE project, and match them with a given list of names or a given regex. That list of servers can be applied to commands. That would mean different things for different commands, but the point is you can abstract where server list definition comes from.

It can run arbitrary scripts on your machine, just like `npm` does. Even simpler: you would do `omg <thing>`. See `customs` section on `.omg.toml` file for details. It can do that once, but also for each of the servers on the resolved server list. Communication between OMG and the custom script is done via environment variables.

It can run arbitrary scripts on remote servers too. Just run `omg run <thing>` and it will run it on all configured servers. See `omg help run` for details.

It can open a terminal with an SSH session on any remote server from the resolved server list. You can configure the actual `terminal` command. Once configured, you just run `omg goto <server name>`.

## What will it be able to do

It will be able to perform all things [bottler](https://github.com/rubencaro/bottler) can. Except for the _release_ part, which needs an inside man on the Erlang VM. That will be sorted out supporting _distillery_.

It will be able to do some of the things [goreleaser](https://goreleaser.com/) does. Like _builds_ for several platforms and _releasing_ to github.

Maybe some things will be delegated to scripts to keep the OMG binary small and not to carry too much project specific logic.

## TODOs

* Support to ship to a remote server
* Force target server definition for dangerous commands
* Pass __all__ data to custom scripts as JSON

## Changelog

### master

* Support to run a remote script
* Support `each` for custom commands
* Pass some data to custom commands

### 0.3

* Add flags help to command help text
* support for `servers` flag
* support for `match` flag

### 0.2

* support GCE to get machine names for `goto`
* `goto` command
* Get rid of Viper
  * Parse TOML using [go-toml](https://github.com/pelletier/go-toml) directly.
  * Parse flags using stdlib.
  * Get environment from stdlib.

### 0.1

* Run `custom` scripts
* `help` command
* Stop using Cobra
* `version` command
* Autogenerate '.omg.toml'