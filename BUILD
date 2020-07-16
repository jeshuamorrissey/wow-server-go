load("@io_bazel_rules_go//go:def.bzl", "go_path")
load("@bazel_gazelle//:def.bzl", "gazelle")

package_group(
    name = "all",
    packages = [
        "//...",
    ],
)

# gopath defines a directory that is structured in a way that is compatible
# with standard Go tools. Things like godoc, editors and refactor tools should
# work as expected.
#
# The files in this tree are symlinks to the true sources.
go_path(
    name = "gopath",
    mode = "link",
    deps = [
        "//util/gen_pkt",

        # Test dependencies.
        "@com_github_stretchr_testify//assert:go_default_library",
    ],
)

# Allows you to run gazelle to add new go repos.
# bazel run //:gazelle -- update-repos github.com/...
gazelle(name = "gazelle")
