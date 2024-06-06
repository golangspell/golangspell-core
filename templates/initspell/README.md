# {{.NewSpellName}}
{{.NewSpellName}} Spell with additional Golang Spell commands for TODO: Add here the description of your new Spell

## Golang Spell
The Core project contains the core commands (and the respective templates) of the platform [Golang Spell](https://github.com/golangspell/golangspell).

![alt text](https://golangspell.com/golangspell/blob/master/img/gopher_spell.png?raw=true)

## Test and coverage

Run the tests

```sh 
TESTRUN=true go test ./... -coverprofile=cover.out

go tool cover -html=cover.out
```

Install [golangci-lint](https://github.com/golangci/golangci-lint#install) and run lint:

```sh
golangci-lint run
```

## Install
To install the {{.NewSpellName}} spell use the command
golangspell addspell {{.ModuleName}} {{.NewSpellName}}

## Update
To update the {{.NewSpellName}} version use the command
golangspell updatespell {{.ModuleName}} {{.NewSpellName}}