# pomdok 🍏

Simple wrapper to [Symfony Go Binary](https://symfony.com/download) for multi-app.

This README does not aim to explain Symfony binary, if you want more details you can read [the article](https://jolicode.com/blog/my-local-server-with-the-symfony-binary) I made to introduce it or [official documentation](https://symfony.com/doc/current/setup/symfony_server.html).

# Installation 💾

pomdok is the only thing you have to install, the symfony binary will be automatically installed if not present on your computer.

## Mac

You can install this binary through [Homebrew](https://brew.sh/):

```bash
brew tap jolicode/pomdok git@github.com:jolicode/pomdok.git
brew install pomdok
```

## Linux

Download [last release](https://github.com/jolicode/pomdok/releases), extract it and you'll have the binary. I suggest you to put it in `/usr/local/bin/` to be easier to use but you can do whatever you want 🤷

# Getting started 🚀

First, you need a configuration file that we call `pomdok.yaml` in your project root as following:
```yaml
pomdok:
  tld: 'test'
  projects:
    - domain: 'api.project'
      path: '/apps/api'
    - domain: 'www.project'
      path: '/apps/front'
    - domain: 'admin.project'
      path: '/apps/back-office'
```

You'll need at least `tld` field and one project to have a valid configuration.

To init `pomdok` for your project run:
```bash
pomdok init
```
You can add `--config=configuration.yaml` option if your configuration file is not in current folder with default name `pomdok.yaml`.

Then to start your applications 🎉
```
pomdok start
```

And to stop them:
```
pomdok stop
```

## Symfony related

To make pomdok works, we're using symfony CLI. Some setup on this side is needed:
- You have to setup Symfony CLI proxy (you can find how on [this slide](https://speakerdeck.com/fabpot/symfony-local-web-server-dot-dot-dot-reloaded?slide=32) or on the [official documentation](https://symfony.com/doc/current/setup/symfony_server.html#setting-up-the-local-proxy))
- And to install Symfony CLI certificate authority through `symfony local:server:ca:install`

# Troubleshooting 🤕

## Some debug tips 🔧

Because this tool use symfony CLI to run your servers, here is some advices to debug when need:
- You can check running servers on `http://127.0.0.1:7080/`
- You used start command but server is still stopped in the list ? Go in the app folder then use: `symfony local:server:start`, you'll have full logs and order to see what's happening !

## Everything working but I have untrusted https ❌

This tool does not run `symfony ca:install` command since it needs sudo. This install local certificate authority. Just run it and you'll have trusted https for you apps 😉

## My website isn't working 😢

When you start pomdok, and you don't have the symfony proxy already launched, you're site won't be reachable. You have to close your web browser (really quit it, not reduce it like mac usually do).

# Commands 🛠

## init

```
pomdok init
```

Will sync your project with Symfony binary configuration.
This command can run anywhere inside your project tree, it will search in current directory and will goes into parent one if nothing and again and again until finding your project configuration.

## start

```
pomdok start
```

Will start symfony proxy if needed and all your apps

## stop

```
pomdok start
```

Will stop all your apps

## check

```
pomdok check
```

Will check your OS and needed binaries:
- OS: should be Linux or Darwin (MacOS)
- PHP: you need local php installation
- Symfony: and the symfony binary 😉

## install

```
sudo pomdok install
```

Will install all needed binaries :
- PHP: from `apt` or `brew` depending on OS
- Symfony: with `wget` command
⚠ This command obviously needs `sudo` or being logged as root.

# Sponsor

[![JoliCode](https://jolicode.com/images/logo.svg)](https://jolicode.com)

Open Source time sponsored by JoliCode
