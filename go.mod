module github.com/jonhadfield/sn-cli

go 1.14

require (
	github.com/cpuguy83/go-md2man/v2 v2.0.0 // indirect
	github.com/divan/num2words v0.0.0-20170904212200-57dba452f942
	github.com/fatih/color v1.9.0
	github.com/fsnotify/fsnotify v1.4.9 // indirect
	github.com/jonhadfield/gosn v0.0.2
	github.com/jonhadfield/gosn-v2 v0.0.0-20200511210922-3703bc5fdde3
	github.com/mitchellh/mapstructure v1.3.0 // indirect
	github.com/pelletier/go-toml v1.7.0 // indirect
	github.com/ryanuber/columnize v2.1.0+incompatible
	github.com/spf13/afero v1.2.2 // indirect
	github.com/spf13/cast v1.3.1 // indirect
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/spf13/viper v1.7.0
	github.com/stretchr/testify v1.5.1
	github.com/urfave/cli v1.22.4
	github.com/zalando/go-keyring v0.0.0-20191216155332-f41212d4b175 // indirect
	golang.org/x/crypto v0.0.0-20200510223506-06a226fb4e37
	golang.org/x/sys v0.0.0-20200509044756-6aff5f38e54f // indirect
	gopkg.in/ini.v1 v1.56.0 // indirect
	gopkg.in/yaml.v2 v2.2.8
)


replace (
	github.com/jonhadfield/gosn-v2  => ../gosn-v2
)