module github.com/jonhadfield/sn-cli

go 1.17

require (
	github.com/asdine/storm/v3 v3.2.1
	github.com/briandowns/spinner v1.12.0
	github.com/divan/num2words v0.0.0-20170904212200-57dba452f942
	github.com/fatih/color v1.14.1
	github.com/fsnotify/fsnotify v1.6.0 // indirect
	github.com/godbus/dbus/v5 v5.1.0 // indirect
	github.com/ryanuber/columnize v2.1.2+incompatible
	github.com/spf13/viper v1.15.0
	github.com/stretchr/testify v1.8.1
	github.com/urfave/cli v1.22.5
	golang.org/x/crypto v0.7.0 // indirect
	gopkg.in/yaml.v2 v2.4.0
)

require (
	github.com/cpuguy83/go-md2man/v2 v2.0.0 // indirect
	github.com/danieljoos/wincred v1.1.2 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/hashicorp/hcl v1.0.0 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/magiconair/properties v1.8.7 // indirect
	github.com/matryer/try v0.0.0-20161228173917-9ac251b645a2 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.17 // indirect
	github.com/mitchellh/go-homedir v1.1.0 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/pelletier/go-toml v1.9.5 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/smartystreets/goconvey v1.6.4 // indirect
	github.com/spf13/afero v1.9.5 // indirect
	github.com/spf13/cast v1.5.0 // indirect
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/subosito/gotenv v1.4.2 // indirect
	github.com/zalando/go-keyring v0.2.2 // indirect
	go.etcd.io/bbolt v1.3.7 // indirect
	golang.org/x/sys v0.6.0 // indirect
	golang.org/x/term v0.6.0
	golang.org/x/text v0.8.0 // indirect
	gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c // indirect
	gopkg.in/ini.v1 v1.67.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

require github.com/meramsey/gosn-v2 v0.0.0-20230307161136-1fc0d4bbe4f4

require (
	github.com/alecthomas/gometalinter v3.0.0+incompatible // indirect
	github.com/alecthomas/units v0.0.0-20211218093645-b94a6e3cc137 // indirect
	github.com/alessio/shellescape v1.4.1 // indirect
	github.com/google/shlex v0.0.0-20191202100458-e7afc7fbc510 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/nicksnyder/go-i18n v1.10.1 // indirect
	github.com/pelletier/go-toml/v2 v2.0.7 // indirect
	github.com/russross/blackfriday/v2 v2.0.1 // indirect
	github.com/shurcooL/sanitized_anchor_name v1.0.0 // indirect
	golang.org/x/tools v0.6.0 // indirect
	golang.org/x/tools/cmd/cover v0.1.0-deprecated // indirect
	gopkg.in/alecthomas/kingpin.v3-unstable v3.0.0-20191105091915-95d230a53780 // indirect
)

//replace github.com/jonhadfield/gosn-v2 => ../gosn-v2
replace github.com/meramsey/gosn-v2 => /home/mike/Documents/dev/gosn-v2