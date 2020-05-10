# golangspell-core
Core Plugin with the main Golang Spell commands

## Golang Spell
The Core project contains the core commands (and the respective templates) of the platform [Golang Spell](https://github.com/danilovalente/golangspell).

## Environment

## Install
The golangspell-core spell is installed automatically by Golang Spell in the first execution.
If you need to install manually the golangspell-core spell for any reason, use the command

```bash
golangspell addspell github.com/danilovalente/golangspell-core golangspell-core
```

## Update
To update the golangspell-core version use the command

```bash
golangspell updatespell github.com/danilovalente/golangspell-core golangspell-core
```

![Spell Gopher](http://derobgfa8qo3s.cloudfront.net/images/gopher_spell.png)

## Version history
* 0.4.9 - Updated documentation
* 0.4.8 - Adjusted string template rendering
* 0.4.7 - Adjusted string templating
* 0.4.6 - Adjusted string templating feature
* 0.4.5 - Adjusted initspell command feedback text, added lint and adjusted code
* 0.4.4 - addspellcommand adjusted to add the new command specification to the Spell's buildconfig command
* 0.4.3 - Included golangci-lint and fixed lint issues
* 0.4.2 - Removed unused imports from template
* 0.4.1 - Fixed directory and naming issues
* 0.4.0 - Updated echo to the v4, added Prometheus to init template, created feature addspellcommand
