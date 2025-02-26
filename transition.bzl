# Contents of transition.bzl

# implementation of the transition
# settings: dict where keys are inputs parameter
def _multiarch_transition(settings, attr):
    # attr.platforms = [Label("//:linux_arm64"), Label("//:linux_amd64")]
    # print("Requested platforms for transition:", attr.platforms)
    # print("Settings:", settings)
    # print("Attr:", attr)
    transitions = [{"//command_line_option:platforms": str(platform)} for platform in attr.platforms]
    
    print("Transitions being applied:", transitions)
    # transitions = [{"//command_line_option:platforms": "@@//:linux_arm64"}, {"//command_line_option:platforms": "@@//:linux_amd64"}]
    return transitions

# reads a set of input build settings and writes a set of output build settings
multiarch_transition = transition(
    implementation = _multiarch_transition,
    inputs = [],
    outputs = ["//command_line_option:platforms"],
)

def _impl(ctx):
    image_files = [f.path for f in ctx.files.image]
    print("Files included in the image:", image_files)
    return DefaultInfo(files = depset(ctx.files.image))

multi_arch = rule(
    implementation = _impl,
    attrs = {
        "image": attr.label(cfg = multiarch_transition),
        "platforms": attr.label_list(),
        "_allowlist_function_transition": attr.label(
            default = "@bazel_tools//tools/allowlists/function_transition_allowlist",
        ),
    },
)
