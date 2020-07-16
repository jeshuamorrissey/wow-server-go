load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "packet_spec",
    srcs = ["packet_spec.go"],
    importpath = "util/gen_pkt/packet_spec",
    visibility = ["//util/gen_pkt:__subpackages__"],
)

go_test(
    name = "packet_spec_test",
    srcs = ["packet_spec_test.go"],
    deps = [
        ":packet_spec",
        "@com_github_alecthomas_participle//:go_default_library",
        "@com_github_stretchr_testify//assert:go_default_library",
    ],
)
