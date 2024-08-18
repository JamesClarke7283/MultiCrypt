# Multi Crypt
This is a program written in Go with Fyne, which enables you to use different symmetric encyption algorithms to encrypt/decrypt UTF-8 text.

## Project Structure
FYI the module name is `github.com/JamesClarke7283/MultiCrypt`
```
.
├── docs
│   ├── LICENSE.md
│   └── PLAN.md
├── go.mod
├── LICENSE
├── Makefile
├── .default_config.toml
├── README.md
└── src
    ├── shared - shared logic between frontend and backend like `logging.go`,etc.
    ├── backend - contains the core code files for doing encryption/decryption for different algorithms.
    ├── frontend - contains all the different frontend components in Fyne.
    └── main.go - Entrypoint/Main window for the app.
```

## Extra

### Logging

Please use either the builtin logging or the most popular library for it.
I need it to also in a platform indiependent way to put the log files in the home folder(please use a different package for platform indiependent file placement).

We need dotenv support also, and env support, with the `LOG_LEVEL` environment variable, which can be for example `DEBUG` but if none is specified the default is `INFO`.

We provide in the `src/shared` package, the `logging.go` file this exports functions that let us get the logger easily.

We also need coloured logs so we can see it easier.

### Themeing

```toml
[appearance.theme.light]
background = "FFFFFF"
foreground = "000000"
primary = "0000FF"
secondary = "00FF00"
tertiary = "FF0000"

[appearance.theme.dark]
foreground = "FFFFFF"
background = "000000"
primary = "0000FF"
secondary = "00FF00"
tertiary = "FF0000"

[appearance]
selected_theme = "system"
```
This is the example `./default_config.toml`, we configure all constants/data in this file, so we seperate code from data. If no file exists in the home folder under the appropriate config directory, on linux this would be (`~/.config/MultiCrypt`)(please use a platform indiependent package to support multiple OS's), we create one called `config.toml` in that config dir prior mentioned and write the default config to it, if it does exist we load from it.

For themeing we select `system` by default, but it can be set to `light`, `dark` or any other theme added to the config file.

## Phases of development

### Phase 1: Laying the foundations:

We implement the Logging, Themeing, make a basic main window with encrypt/decrypt functionality with 1 symmetric encryption algorithm like AES-256.

We need a Key, Message and Ciphertext box.

### Phase 2: Improvements

We add a settings button which opens the settings and lets us change the theme and font size of the application under a `Appearance` tab.

We add a `Copy To Clipboard`(for ciphertext) and Generate Random Key button, the generate random key button opens a popup where you can specify the characterset and length of the key.

We add more algorithms like Serpent and add a new dropdown(in addition to the `Cipher/Algorithm` one, called `Variant` where you can specify the `Mode`/`Bits` of the cipher as a dropdown). we also add the variants.