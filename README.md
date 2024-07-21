# PlantUML LSP

An implementation of the language server protocol (LSP) for PlantUML.

;
## Installation

TODO: (requires cloning stdlib repo)

---

## Capabilities

**Disclaimer: This project is very early in devlopment, so many features will be missing.**

Contributions are welcome, please open an issue or PR if you would like to add something.

#### Completion:
- [x] Core (WIP)
    - [x] Types, keywords, directives
    - [ ] Colors, skinparams (Figuring out if these should be included)
- [ ] stdlib/C4 (WIP)
    - [x] Procedures
    - [ ] Functions, globals, defines, constants, variables (Todo)
    - [ ] Snippets (Todo)
- [ ] Other stdlib (Todo)
- [ ] User Defined (Todo)

#### Hover
- [ ] Core (Backlog)
- [ ] stdlib/C4 (WIP)
- [ ] other stdlib (Backlog)
- [ ] User Defined (Backlog)

#### Definition
- [ ] Core (Backlog)
- [ ] stdlib/C4 (Backlog)
- [ ] other stdlib (Backlog)
- [ ] User Defined (Backlog)

#### Diagnostics
- [ ] Core (Backlog)
- [ ] stdlib/C4 (Backlog)
- [ ] other stdlib (Backlog)
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

