# adams
Go linter that checks for calls to panic(). For example in request handling code paths of a microservice, it is undesirable to panic.

This is a work in progress. It works but doesnt check for reassigns of panic in the surrounding scope.
