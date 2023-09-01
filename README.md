# adams
Go linter that checks for calls to panic(). For example in request handling code paths of a microservice, it is undesirable to panic.

This is a work in progress. It works but I dont think its performant to render the node in every check so beware of that.
