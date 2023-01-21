<div align="center">

# uwkendo discord

</div>

utility kendo related discord bot

## Setting up for development

We are using golang version 1.19.5, you can install for your system or choice, or using [gvm](https://github.com/moovweb/gvm):
```sh
gvm install go1.19.5 -B
gvm use go1.19.5
```

Next, install developer git hooks
```sh
make devsetup
```

Create the local copy of .env
```sh
cp env.example .env
```

