/// the get router attribute for the method in interface
attribute get: Router

/// the post router attribute for the method in interface
attribute post: Router

/// the put router attribute for the method in interface
attribute put: Router

/// the patch router attribute for the method in interface
attribute patch: Router

/// the delete router attribute for the method in interface
attribute delete: Router

/// the head router attribute for the method in interface
attribute head: Router

/// the trace router attribute for the method in interface
attribute trace: Router

/// the options router attribute for the method in interface
attribute options: Router

/// the body attribute indicating the field in struct is mapping to http body for request or response
attribute body: Bool

attribute query: String

attribute header: TemplateHeader

attribute fragment: TemplateString

/// will replace the `{{key}}` block in the router, specially in the `path` field
attribute router_template_values: Object
