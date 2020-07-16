load("@io_bazel_rules_go//go:def.bzl", "go_binary")

go_binary(
    name = "gen_pkt",
    srcs = ["gen_pkt.go"],
    visibility = ["//visibility:public"],
    deps = [
        "//util/gen_pkt/packet_spec",
        "@com_github_alecthomas_participle//:go_default_library",
    ],
)
