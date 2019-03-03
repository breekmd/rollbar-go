# rollbar-go fork details
This is a fork of the original go package

It adds functionality to allow providing your own "poster"
(for making POST requests)

This is useful when trying to add Rollbar logging into an
Appengine golang project and you need to make the POST
requests using url.Fetch and provide a context.

However, while this will work for Appengine, architecturally 
is not the best option, as request context has to be provided
to the appengine client on every post, whereas rollbar-go 
is pretty much a singleton pattern. Use this in case no other
options available.


only difference from the original library is that you can create a custom "poster"
that should handle a POST (i.e. using urlfetch.Client if used in appengine)
and this custom poster can be passed to rollbar.
Rollbar will make all API posts using the custom handler you provide
if no poster will be provided a default one will be used
```
poster := RollbarPoster{Context:context}
rollbar.SetPoster(poster)
```

# rollbar-go original readme

[Rollbar](https://rollbar.com) is a real-time exception reporting service for Go
and other languages. The Rollbar service will alert you of problems with your code
and help you understand them in a ways never possible before. We love it and we hope
you will too.

rollbar-go is a Golang Rollbar client that makes it easy to report errors to
Rollbar with full stacktraces. Errors are sent to Rollbar asynchronously in a
background goroutine.

Because Go's `error` type doesn't include stack information from when it was set
or allocated, we use the stack information from where the error was reported.

# Setup Instructions and Usage

1. [Sign up for a Rollbar account](https://rollbar.com/signup)
2. Follow the [Usage](https://docs.rollbar.com/docs/go#usage) example in our [Go SDK docs](https://docs.rollbar.com/docs/go) 
to get started for your platform.

# Documentation

[API docs on godoc.org](http://godoc.org/github.com/rollbar/rollbar-go)
