load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")
load("@bazel_gazelle//:def.bzl", "gazelle")

# gazelle:prefix github.com/rockwotj/godav
gazelle(name = "gazelle")

go_library(
    name = "godav_lib",
    srcs = ["webdav.go"],
    importpath = "github.com/rockwotj/godav",
    visibility = ["//visibility:private"],
    deps = ["@org_golang_x_net//webdav:go_default_library"],
)

go_binary(
    name = "godav",
    embed = [":godav_lib"],
    visibility = ["//visibility:public"],
)
