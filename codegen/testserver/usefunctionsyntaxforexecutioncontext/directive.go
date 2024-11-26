package usefunctionsyntaxforexecutioncontext

import (
    "context"
    "log"

    "github.com/99designs/gqlgen/graphql"
)

// LogDirective implementation
func LogDirective(ctx context.Context, obj interface{}, next graphql.Resolver, message *string) (res interface{}, err error) {
    log.Printf("Log Directive: %s",*message)

    // Proceed with the next resolver in the chain
    return next(ctx)
}
