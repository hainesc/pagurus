# Function & Pipeline

Function is a CRD that indicates a function created by user. The controller will create a crab service for each given function.

Pipeline is a CRD that indicates a pipeline created by user. The controller will create a crab service for each give pipeline.

Metadata: Name, Namespace, Revision
Spec: Code(ConfigMap), Runtime, Layer, Handler

Function
Pipeline
Revision
Runtime python 2.7...
Layer Dependency
Handler
ConfigMap
