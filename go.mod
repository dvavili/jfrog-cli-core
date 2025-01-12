module github.com/jfrog/jfrog-cli-core/v2

go 1.14

require (
	github.com/buger/jsonparser v1.1.1
	github.com/c-bata/go-prompt v0.2.5
	github.com/chzyer/readline v0.0.0-20180603132655-2972be24d48e
	github.com/danwakefield/fnmatch v0.0.0-20160403171240-cbb64ac3d964 // indirect
	github.com/denormal/go-gitignore v0.0.0-20180930084346-ae8ad1d07817
	github.com/google/uuid v1.3.0
	github.com/gookit/color v1.4.2
	github.com/jedib0t/go-pretty/v6 v6.2.2
	github.com/jfrog/build-info-go v1.1.0
	github.com/jfrog/gofrog v1.1.1
	github.com/jfrog/jfrog-client-go v1.10.0
	github.com/magiconair/properties v1.8.5
	github.com/manifoldco/promptui v0.8.0
	github.com/pkg/browser v0.0.0-20210911075715-681adbf594b8
	github.com/pkg/errors v0.9.1
	github.com/spf13/viper v1.8.1
	github.com/stretchr/testify v1.7.0
	github.com/urfave/cli v1.22.5
	golang.org/x/crypto v0.0.0-20211202192323-5770296d904e
	golang.org/x/mod v0.4.2
	gopkg.in/yaml.v2 v2.4.0
)

// Exclude vulnerable dependencies
exclude (
	github.com/miekg/dns v1.0.14
	github.com/pkg/sftp v1.10.1
)

// replace github.com/jfrog/jfrog-client-go => github.com/jfrog/jfrog-client-go v1.9.1-0.20220217145821-3e29721b7973

// replace github.com/jfrog/build-info-go => github.com/jfrog/build-info-go v1.0.2-0.20220222144839-297d97db5248

// replace github.com/jfrog/gofrog => github.com/jfrog/gofrog v1.0.7-0.20211128152632-e218c460d703
