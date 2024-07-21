# PlantUML LSP

An implementation of the language server protocol (LSP) for PlantUML.

**Disclaimer: This project is very early in devlopment, so many features will be missing.**
Contributions are welcome, please open an issue or PR if you would like to add something.


## Installation

TODO:
(for the best experience, this lsp requires cloning stdlib repo)

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

- [ ] Snippets (Todo)
    - [ ] Core (Todo)
    - [ ] stdlib/C4 (Todo)

#### Hover
- [x] Core
    - [ ] Examples usage (Backlog)
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

