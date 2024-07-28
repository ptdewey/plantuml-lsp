# PlantUML LSP

An implementation of the language server protocol (LSP) for PlantUML.

**Disclaimer: This project is very early in devlopment, so many features will be missing.**
Contributions are welcome, please open an issue or PR if you would like to add something.


## Installation

```sh
# clone repository
git clone https://github.com/ptdewey/plantuml-lsp.git
cd plantuml-lsp

# build lsp binary
go build
cd ..

# (optional but recommended) clone stdlib repo
git clone https://github.com/plantuml/plantuml-stdlib.git
# or alternatively, extract stdlib from plantuml executable
plantuml -extractstdlib
```

---

### Setup

Neovim (with lspconfig):

```lua
config = function()
    local lspconfig = require("lspconfig")
    local configs = require("lspconfig.configs")
    if not configs.plantuml_lsp then
        configs.plantuml_lsp = {
            default_config = {
                cmd = { "/path/to/plantuml_lsp", "--stdlib-path=/path/to/plantuml-stdlib" },
                filetypes = { "plantuml" },
                root_dir = function(fname)
                    return lspconfig.util.find_git_ancestor(fname) or lspconfig.util.path.dirname(fname)
                end,
                settings = {},
            }
        }
    end
    lspconfig.plantuml_lsp.setup {}
end,
```

* NOTE: This assumes plantuml is set up as a filetype already

---

VS Code:

TODO (I don't use VS Code)

---

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
- [ ] Core (Backlog)
- [ ] stdlib/C4 (Backlog)
- [ ] Other stdlib (Backlog)
- [ ] User Defined (Backlog)

#### Other Language Server Features
Other language server features are not currently planned.
See the [language server protocol specification](https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#languageFeatures) for more info about other features.

---

## Developer Notes

#### LSP Specification
- [https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/](https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/)

#### PlantUML Language specification:
- [https://github.com/plantuml/plantuml.js/blob/main/plantuml-docs/docs_from_alphadoc/developers.adoc](https://github.com/plantuml/plantuml.js/blob/main/plantuml-docs/docs_from_alphadoc/developers.adoc)
    - `plantuml -language` is a very helpful command (outputs language keywords)
    - `cat text.puml | plantuml -syntax` is also useful (checks if thing is valid plantuml, could be good for static analysis)
        - might be too slow for static analysis

