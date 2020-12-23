# git-config-switcher

A tool to easily switch ``git config``

## Demo

[![asciicast](https://asciinema.org/a/ogZDrtqdil8l3NOrJGWD9PN4y.svg)](https://asciinema.org/a/ogZDrtqdil8l3NOrJGWD9PN4y)   


## Install

```bash
$ go get github.com/DuGlaser/git-config-switcher
```

## Usage
1. Create ``git-config-switch.toml`` in ``$HOME/.config``
2. Write in A as shown below
```toml
[[config]]
name="test"

  [config.user]
  name="name1"
  email="email1"

[[config]]
name="test2"

  [config.user]
  name="name2"
  email="email2"
```

3. Run ``git-config-switch``  


![SS](https://user-images.githubusercontent.com/50506482/102792288-a506e480-43eb-11eb-9403-2b1ead54dfae.png)
