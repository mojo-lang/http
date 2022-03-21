package http

import (
    "github.com/mojo-lang/core/go/pkg/logs"
    "github.com/mojo-lang/lang/go/pkg/mojo/lang"
)

const (
    RouterTemplateValuesAttributeName     = "router_template_values"
    RouterTemplateValuesAttributeFullName = "http.router_template_values"
)

func GetRouterTemplateValues(attributes ...*lang.Attribute) map[string]string {
    values := make(map[string]string)
    for _, attribute := range attributes {
        if attribute.IsSameName(RouterTemplateValuesAttributeFullName) {
            for _, argument := range attribute.Arguments {
                if len(argument.Label) > 0 {
                    if value, err := argument.GetString(); err != nil {
                        logs.Warnw("failed to parse the string argument",
                            "attribute", RouterTemplateValuesAttributeFullName, "label", argument.Label)
                    } else {
                        values[argument.Label] = value
                    }
                }
            }
        }
    }

    return values
}
