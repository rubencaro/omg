# omg

Because all those helper scripts around your code should be amazing too.

## What

A tool to efficiently manage all those little things/scripts/files around _my_ code that make up for different stages of development. Such as compiling, releasing, deploying, packaging or publishing, but also linting, formatting, testing, benchmarking, etc.

I mean it to be stack independent, so I can use it with all my Javascript/Elixir/Go/Whatever projects. Also compatible with all the shapes and colors of helper scripts and tools I use with them.

-----

Everything below is just a __plan__...

-----

## Usage

You just download the compiled binary of the last release for your platform into the root of your project, and then `omg init`. That should create a commented `.omg.toml` file. You fill it in and you are good to go on asking `omg` to do things. `omg help` should help. Things like `omg deploy` or `omg release` or `omg publish` should start to work.

Both the `omg` binary and the `.omg.toml` file are meant to be included on  version control. Both are locked to your code. That's a big fat feature.

All documentation should be accessible from the `omg help` command and comments on `.omg.toml` itself. You shouldn't need to come back here to look for usage instructions for your particular copy of `omg`. That's another big fat feature.

## What can it do

It can run arbitrary scripts on your machine, just like `npm run <thing>` does.

It can perform all things [bottler](https://github.com/rubencaro/bottler) can. Except for the _release_ part, which needs an inside man on the Erlang VM. That will be sorted out supporting _distillery_.

It can do some of the things [goreleaser](https://goreleaser.com/) does. Like _builds_ for several platforms and _releasing_ to github.

## TODOS

* Everything, even fill the TODOs list in.
