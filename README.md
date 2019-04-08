# pomdok 🍏

Simple wrapper to [Symfony Go Binary](https://symfony.com/download) for multi-app.

This README does not aim to explain Symfony binary, if you want more details you can read [the article](https://jolicode.com/blog/my-local-server-with-the-symfony-binary) I made to introduce it or [official documentation](https://symfony.com/doc/current/setup/symfony_server.html).

# Getting started 🚀

First you need a configuration file in your project root as following:
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

You'll need atleast `tld` field and one project to have a valid configuration.

To init `pomdok` for your project run:
```bash
pomdok init
```
You can add `--config=confiuration.yaml` option if your configuration file is not in current folder with default name `pomdok.yaml`.

Then to start your applications 🎉
```
pomdok start
```

And to stop them:
```
pomdok stop
```

# Troubleshooting 🤕

## Everything working but I have untrusted https

This tool does not run `symfony ca:install` command since it needs sudo. This install local certificate authority. Just run it and you'll have trusted https for you apps 😉

# Commands 🛠

## init

```
pomdok init
```

Will sync your project with Symfony binary configuration 

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
