# base-go

Base template for go projects
## Development Requirements

You will need to have installed the following:

- [Docker](https://www.docker.com/products/docker-desktop) (>= 18.09.2).

Note: The project was created with this specific version but older versions might also work.

## Git hooks

You need to run the commands to install the pre-commit:

```sh
make pre-commit install
```

It is recommended to install the git hooks present in `/githooks` folder:

git 2.9 or higher:

```sh
git config core.hooksPath githooks
```

git 2.8 or lower:

```sh
find .git/hooks -type l -exec rm {} \; && find githooks -type f -exec ln -sf ../../{} .git/hooks/ \;
```

**NOTE: If you change the `.pre-commit-config.yaml` you need remove the githook path configuration**

## Getting started

## Golang code format and style guide

[`golangci-lint`](https://github.com/golangci/golangci-lint) is used to enforce formatting and code style.

Additionally there should be an effort to follow these code style guidelines and recommendations:

- https://talks.golang.org/2014/names.slide
- https://golang.org/doc/effective_go.html#names
- https://blog.golang.org/package-names
- https://github.com/golang/go/wiki/CodeReviewComments
- [Style guideline for Go packages](https://rakyll.org/style-packages) (rakyll/JBD)

## Folder structure

This repository follows the [golang-standards/project-layout](https://github.com/golang-standards/project-layout/tree/5b325fed762cac9d0c63fcd412a33de6bfcac389).
In the future if the need arises to create a new top level folder there should be an effort to continue to follow the [golang-standards/project-layout](https://github.com/golang-standards/project-layout/tree/5b325fed762cac9d0c63fcd412a33de6bfcac389).

### Go Directories

#### `/cmd`

Main applications.

The directory name for each application should match the name of the executable (e.g., `/cmd/myapp`).

Very few code should be present in the application directory. If the code can be imported and used in other projects, then it should live in the `/pkg` directory. If the code is not reusable or shareable, then it should live in the `/internal` directory.

It's common to have a small `main` function that imports and invokes the code from the `/internal` and `/pkg` directories and nothing else.

#### `/pkg`

Library code that's ok to use by external applications (e.g., `/pkg/mypubliclib`). Other projects will import these libraries expecting them to work, so think twice before putting something here :-) Note that the `internal` directory is a better way to ensure the private packages are not importable because it's enforced by Go. The `/pkg` directory is still a good way to explicitly communicate that the code in that directory is safe for use by others. The [`I'll take pkg over internal`](https://travisjeffery.com/b/2019/11/i-ll-take-pkg-over-internal/) blog post by Travis Jeffery provides a good overview of the `pkg` and `internal` directories and when it might make sense to use them.

It's also a way to group Go code in one place when your root directory contains lots of non-Go components and directories making it easier to run various Go tools (as mentioned in these talks: [`Best Practices for Industrial Programming`](https://www.youtube.com/watch?v=PTE4VJIdHPg) from GopherCon EU 2018, [GopherCon 2018: Kat Zien - How Do You Structure Your Go Apps](https://www.youtube.com/watch?v=oL6JBUk6tj0) and [GoLab 2018 - Massimiliano Pippi - Project layout patterns in Go](https://www.youtube.com/watch?v=3gQa1LWwuzk)).

This is a common layout pattern, but it's not universally accepted and some in the Go community don't recommend it.

Ok not to use it if the app project is really small and where an extra level of nesting doesn't add much value.

#### `/build`

Packaging and Continuous Integration.

Cloud (AMI), container (Docker), OS (deb, rpm, pkg) package configurations and scripts should be in the `/build/package` directory.

CI (travis, circle, drone) configurations and scripts should be in the `/build/ci` directory. Note that some of the CI tools (e.g., Travis CI) are very picky about the location of their config files. Try putting the config files in the `/build/ci` directory linking them to the location where the CI tools expect them (when possible).

### Other Directories

#### `/docs`

Design and user documents (in addition to your godoc generated documentation).

#### `/githooks`

Git hooks.

## Contributing

The `develop` branch is where the source code of HEAD always reflects a state with the latest delivered development changes for the next release.

A new feature should be developed in a branch whose name follows the pattern: `feature/<ISSUE-CODE>-<SHORT-DESCRIPTION>`, example: `feature/middleware-1057-add-user-groups`.

A bug fix should be developed in a branch whose name follows the pattern: `fix/<ISSUE-CODE>-<SHORT-DESCRIPTION>`, example: `fix/middleware-1057-fix-user-roles`.

The branch name should be lowercase, having uppercase characters is known to cause problems with some tools.
If no issue exist, the `<ISSUE-CODE>-` should be omitted.

When a new feature or bug fix is completed, a new PR (Pull Request) into the `develop` branch should be created so the changes can be reviewed.

The PR title should start with the issue number, for example `MIDDLEWARE-1057: Add user groups` if a issue exists.

All PRs should have one of the following labels:

- feature
- bugfix

All PRs should have the requester as assignee.

Also if a PR is still a WIP (Work In Progress) it should be marked with the `WIP` label and it should be ignored by the reviewers. This means the PR still has pending changes.

The label `RFC` (Request For Comments) should be added to any PR that it is ready to be reviewed.

The PR reviewer(s) should delegate the merge (when the PR is approved) to the original commiter.
