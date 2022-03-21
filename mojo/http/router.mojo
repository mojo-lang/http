/// the router information for the method in the interface
type Router {
    path: TemplateString @1
    
    headers: [TemplateHeader] @2

    method: Method @10
}
