# Pagurus

## Function

Function is a CRD that indicates a function created by user. The controller will create a crab service for each given function.

## Crab

TODO: Crab is not needed, we can create a router as fission does. For function creation, we only add the code to share storage.

A crab is a application run as deployment in kubernetes for a given function. When called by triggers(often with input data), it will find a shell to execute and return the result.

## Shell

Shell is a long lived executor with suitable environments for some functions, when a function completes, it does not exit and waits for the next one.

fission function create --name hello --env python --code hello.py --executortype poolmgr
fission function create --name hello --env python --code hello.py --executortype poolmgr --type async --callback ...

Service userspace, when the pod tag itself as busy, the service will not send request to it.

To achive a event mesh, each Pod can decide whether or not to accept traffic. 
