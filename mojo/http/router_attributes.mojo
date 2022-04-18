
/// the base path attribute for the method in interface
@target(DeclType.function, DeclType.interface)
attribute base_path: String

/// the get router attribute for the method in interface
@target(DeclType.function)
attribute get: Router

/// the post router attribute for the method in interface
@target(DeclType.function)
attribute post: Router

/// the put router attribute for the method in interface
@target(DeclType.function)
attribute put: Router

/// the patch router attribute for the method in interface
@target(DeclType.function)
attribute patch: Router

/// the delete router attribute for the method in interface
@target(DeclType.function)
attribute delete: Router

/// the head router attribute for the method in interface
@target(DeclType.function)
attribute head: Router

/// the trace router attribute for the method in interface
@target(DeclType.function)
attribute trace: Router

/// the options router attribute for the method in interface
@target(DeclType.function)
attribute options: Router

/// will replace the `{{key}}` block in the router, specially in the `path` field
@target(DeclType.function, DeclType.interface)
attribute router_template_values: Object
