# Pagurus

## Function

Function is a CRD that indicates a function created by user. The controller will create a crab service for each given function.

## Crab

A crab is a application run as deployment in kubernetes for a given function. When called by triggers(often with input data), it will find a shell to execute and return the result.

## Shell

Shell is a long lived executor with suitable environments for some functions, when a function completes, it does not exit and waits for the next one.
