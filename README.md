# k8s-go-scope
HTTP (golang) service that exports as much as possible from inside k8s

### Deployment

    make package # place container image where k8s can get it.
    kubectl create -f k8s-go-scope.yaml
