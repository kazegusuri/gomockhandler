# gomockhandler

**If you find any bugs or have feature requests, please feel free to create an issue.**

gomockhandler is handler of [golang/mock](https://github.com/golang/mock), as the name implies.

`gomockhandler` use one config file to generate all mocks.

With `gomockhandler`,

- You can generate mocks **faster** :rocket:.
- You can check if mock is **up-to-date** :sparkles:.
- You can manage your mocks in **one config file** :books:.

Here is some example of the mock being generated in half the time with `gomockhandler`. (I ran `mockgen` to generate same mocks in `go generate ./...`)

<img width="509" alt="Screen Shot 2021-03-25 at 11 55 54" src="https://user-images.githubusercontent.com/44139130/112411968-1050b180-8d61-11eb-8321-d9a890de292a.png">

<img width="529" alt="Screen Shot 2021-03-25 at 11 56 29" src="https://user-images.githubusercontent.com/44139130/112412003-2494ae80-8d61-11eb-8b0f-08098ee6992e.png">

This speedup is only due to the fact that gomockhandler is running mockgen in parallel, so depending on the power of the PC you are using, you may or may not be able to get a higher speedup than this.
Also, the gomockhandler has a function to skip mock generation if the mock source file has not been updated, but in this experiment, all mocks were generated. (None of the mock generation was skipped.)

## Background

Some of you may often manage your mocks with `go generate` like below.

```
//go:generate mockgen -source=$GOFILE -destination=mock_$GOFILE -package=$GOPACKAG
```

But, it will take long time to generate a log of mocks with `go generate ./...`, because `go generate` executes mockgen one by one. And we cannot easily check if mock is up-to-date.

`gomockhandler` is created to solve all of these problems.

And with this background, it is designed to make it easy to switch from managing mocks with `go generate` to managing mocks with gomockhandler.

## Install

You have to install `mockgen` first.

### Go version < 1.16
```
GO111MODULE=on go get github.com/golang/mock/mockgen
GO111MODULE=on go get github.com/sanposhiho/gomockhandler
```
### Go 1.16+
```
go install github.com/golang/mock/mockgen@latest
go install github.com/sanposhiho/gomockhandler@latest
```

## How to use

`gomockhandler` is designed to be **simple** and does only three things.

- generate/edit a config with CLI
- generate mocks from config
- check if mocks are up-to-date

These are the options for using the gomockhandler.

```
-config string
  The path to config file.
  The default value is "./gomockhandler.json"
  
-target_dir string
  Only mocks under the specified directory will be targeted for commands `check` and `mockgen`.
  By default, all files will be targeted.
  
-f bool
  If true, it will also generate mocks whose source has not been updated.
  The default value is false.
```

## generate mock

You can generate all mocks from config.

```
gomockhandler -config=/path/to/gomockhandler.json mockgen
```

## check if mock is up-to-date

You can check if the mock is generated based on the latest interface.

It is useful for ci.

```
gomockhandler -config=/path/to/gomockhandler.json check
```
If some mocks are not up to date, you can see the error and `gomockhandler` will exit with exit-code 1

```
2021/03/10 22:17:12 [WARN] mock is not up to date. source: ./interfaces/user.go, destination: ./interfaces/../mock/user.go
2021/03/10 22:17:12 mocks is not up-to-date
```

## configuring

You need a config for `gomockhandler`.

However, you don't need to generate/edit the config directly, it can be generated/edited from CLI.

### configuring a new mock

You can configure a new mock to be generated with CLI. It will also check if mockgen will run correctly with that option.

If a config file does not exist, a config file will be created.

`mockgen` has two modes of operation: source and reflect, and gomockhandler support both.

See [golang/mock#running-mockgen](https://github.com/golang/mock#running-mockgen) for more information about the two modes and mockgen options.

Source mode:
```
gomockhandler -config=/path/to/gomockhandler.json -source=foo.go -destination=./mock/ [other mockgen options]
```

Reflect mode:
```
gomockhandler -config=/path/to/gomockhandler.json -destination=./mock/ [other mockgen options] database/sql/driver Conn,Driver
```

---

You can use all options of mockgen to add a new mock.

For example, suppose you want to configure the mock generated by the following mockgen command to be generated by gomockhandler

```
mockgen -source=foo.go -destination=../mock/
```

The following command will add the information of the mock you want to generate to the configuration.
As you can see, you just need to think about the option `config`. (The default value is `./gomockhandler.json`)

```
gomockhandler -config=/path/to/gomockhandler.json -source=foo.go -destination=../mock/
```

---

gomockhandler is designed to make it easy to switch from managing mocks with `go generate` to managing mocks with gomockhandler.

If you use `go:generate` to execute mockgen now, you can generate the config file by rewriting `go:generate` comment a little bit.

Replace from `mockgen` to `gomockhandler -config=/path/to/gomockhandler.json` in all `go:generate` comments, and run `go generate ./...` in your project. And then,

```
- //go:generate mockgen -source=$GOFILE -destination=mock_$GOFILE -package=$GOPACKAG
+ //go:generate gomockhandler -config=/path/to/gomockhandler.json -source=$GOFILE -destination=mock_$GOFILE -package=$GOPACKAG
```

After generating the config, your `go:generate` comments are no longer needed. You've been released from a slow-mockgen with `go generate`!

Let's delete all `go:generate` comments for mockgen in your project.

### Recommendations

- name the config file `gomockhandler.json`, and place it in a location where the gomockhandler is likely to be run frequently.
- use **Source Mode**, if there is no great reason to use Reflect mode., which is much faster because the gomockhandler will skip processing if the source file has not been changed.

### delete mocks to be generated from config

You can remove the mocks to be generated from the config.

```
gomockhandler -config=/path/to/gomockhandler.json -destination=./mock/user.go deletemock 
```

## edit config manually

You can edit the config manually.

But, it is **RECOMMENDED** to use [CLI](https://github.com/sanposhiho/gomockhandler#generateedit-a-config), especially for adding/editing mocks. (This is because CLI will check if mockgen works correctly with that option, and then edit the config.)

The config json file has the following format.

```
{
	"mocks": {
		"mock/user.go": {
			"checksum": "qxZ/pLjLBtib7o+kDLOzOQ==",
			"source_checksum": "UUyR0gaRX4IbPPAttwOCXw==",
			"mode": "SOURCE_MODE",
			"source_mode_runner": {
				"source": "interfaces/user.go",
				"destination": "mock/user.go"
			}
		},
		"mock/user2.go": {
			"checksum": "qxZ/pLjLBtib7o+kDLOzOQ==",
			"source_checksum": "AAAAAAAAAAAAAAAAAAAAAA==",
			"mode": "REFLECT_MODE",
			"reflect_mode_runner": {
				"package_name": "playground/interfaces",
				"interfaces": "User2",
				"destination": "mock/user2.go"
			}
		}
	}
}
```

As mentioned above, there are two modes of mockgen, and the format of the config is slightly different depending on which mode you are using.
In the `***mode-runner` field, specify the option to be used when running mockgen.

In the `checksum` field, the checksum of the currently generated mock is stored. With this checksum, the gomockhandler checks if the mock is the same as the mock generated from the latest interface.

In the `source_checksum` field, the checksum of the currently generated mock's source file is stored. This field is only valid in source mode, and when reflect mode is used, the value will be `AAAAAAAAAAAAAAAAAAAAAA==`.
