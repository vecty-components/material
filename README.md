# Material Components For Vecty

Vecty Material is Vecty bindings for the [material-components-web][] JavaScript
library ([MDC](#Terminology)).

## Project Status

Material was under active development, but development was suspended in
2018. Development is currently ongoing to move the project to WASM, but
as of now the project is not fully functional.

Your help is requested and appreciated. With your support, it will be
possible to easily build a web application in Go using material
components. Help make this dream a reality by contributing to the project!

Monetary contributions are not accepted at this time; only pull requests.

## Getting Started

To simplify demos and development, this package makes use of gaepher. To
install it, run:

    go install bitbucket.org/xoviat/gaepher@6134584

Once gaepher is installed, change to the demos folder

    cd demos

Then run it with

    gaepher serve

Under the hood, gaepher will compile the demo using the go compiler,
then start an http server on port 4000 and serve a bootstrap html
file, the `wasm_exec.js` from your local go installation, and the wasm
file. This allows all code in the repository to be written in go without
html, css, or js. 

Gaepher is meant to make development of this project efficient and
painless. If you have questions/comments about running the demos,
or think the process could be more efficient, commment on the
[relevant issue](https://github.com/vecty-components/material/issues/45).

## Contributing

Any and all ideas, changes, bug reports, etc. are encouraged.

Get involved by:
- Opening issues and submitting merge requests at our [Gitlab project][Material]
  .
- Starting a discussion in the [Gopher Slack][] #GopherJS channel.

## Project Goals

- Expose an API that supports idiomatic Go programming while maintaining
  the general terminology of MDC resources.
- Minimal dependencies to keep generated JavaScript sizes as small as possible.
- Modular package organization so you can pick-and-choose components just like
  the MDC library.
- Limited scope. Only focus on wrapping the JavaScript functionality of the MDC
  library.

## Project Non-Goals

- HTML/CSS. These are areas of the MDC library that are best handled by projects
  that consume Material's packages. Check out
  [vecty-material][] as an example.

## Terminology

To help avoid ambiguity in code and documentation, here's a quick list of
definitions we are trying to enforce in Material.

| Term                        | Definition                                           |
| :---                        | :---------                                           |
| [material-components-web][] | The official upstream JavaScript library.            |
| MDC or MCW                  | Synonyms/abbreviations of `material-components-web`. |
| [Material][]                | This project.                                        |
| [Material component][]      | Our `github.com/vecty-components/material/*` Go packages.             |
| [Material menu][]           | An example of a specific type of Material component. |

[material-components-web]: https://github.com/material-components/material-components-web
[Material]: https://gitlab.com/agamigo/material
[Material component]: https://godoc.org/github.com/vecty-components/material
[Material menu]: https://godoc.org/github.com/vecty-components/material/menu
[vecty-material]: https://gitlab.com/agamigo/vecty-material
[Gopher Slack]: https://gophers.slack.com
[gl-issue-1]: https://gitlab.com/agamigo/material/issues/1
