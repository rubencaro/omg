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

It can run arbitrary scripts on your machine, just like `npm run <thing>` does. Even simpler: you would do `omg <thing>`. See `custom` section on `.omg.toml` file.

## What will it be able to do

It will be able to perform all things [bottler](https://github.com/rubencaro/bottler) can. Except for the _release_ part, which needs an inside man on the Erlang VM. That will be sorted out supporting _distillery_.

It will be able to do some of the things [goreleaser](https://goreleaser.com/) does. Like _builds_ for several platforms and _releasing_ to github.

## TODOs

* ...

## Changelog

### master

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