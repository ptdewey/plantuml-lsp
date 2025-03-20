# PlantUML LSP

A Visual Studio Code extension that integrates the PlantUML LSP (Language Server Protocol), allowing for enhanced language features like autocompletion, hover, diagnostics, and more.

Disclaimer: This project is in early development, and features may be missing. Contributions are welcome—please open an issue or PR if you'd like to help!

## Setup

1. Install PlantUML stdlib (optional but recommended):
    ```sh
    git clone https://github.com/plantuml/plantuml-stdlib.git
    # Alternatively, extract stdlib from the plantuml executable:
    plantuml -extractstdlib
    ```
2. For diagnostics, provide path to plantuml jar file or executable (`plantuml` is fine if it is on your PATH)

### Configuration

If you want to configure the extension (e.g., to point to custom plantuml.jar or executable files), you can modify your settings.json or configure them directly in your workspace.

The following options can be set in your user settings
```json
{
    "plantuml-lsp.execPath": "plantuml",
    "plantuml-lsp.jarPath": "/path/to/plantuml.jar",
    "plantuml-lsp.stdLibPath": "/path/to/plantuml-stdlib"
}
```

* NOTE: This assumes plantuml is set up as a filetype (ensure you have the plantuml extension) already
* NOTE: cmd flags `--exec-path` and `--jar-path` are used to derive diagnostics. These flags can optionally be omitted
    * This argument allows for use of plantuml via a Jar or a system visible binary.
    * Only one of these flags should be specified at a given time.


## Capabilities

#### Completion:
- [x] Core (WIP)
    - [x] Types, keywords, directives
    - [ ] Colors, skinparams (Figuring out if these should be included)
- [ ] stdlib/C4 (WIP)
    - [x] Procedures
    - [ ] Functions, globals, defines, constants, variables (Todo)
- [ ] Other stdlib (Todo)
- [ ] User Defined (Todo)

- [ ] Snippets (WIP)
    - [ ] Core (Todo)
    - [x] stdlib/C4

#### Hover
- [x] Core
    - [ ] Example usage (Backlog)
- [x] stdlib/C4
    - [ ] Example usage (Backlog)
- [ ] Other stdlib (Backlog)
- [ ] User Defined (Backlog)

#### Definition
- [ ] Core (Might be impossible with plantuml project structure)
- [ ] stdlib/C4 (Todo)
- [ ] Other stdlib (Backlog)
- [ ] User Defined (Backlog)

#### Diagnostics
- Diagnostics currently depend on `plantuml -syntax` output
- [x] Core
- [x] stdlib/C4
- [x] Other stdlib
- [x] User Defined


## Troubleshooting

Issues can be filed at [https://github.com/ptdewey/plantuml-lsp/issues](https://github.com/ptdewey/plantuml-lsp/issues)
