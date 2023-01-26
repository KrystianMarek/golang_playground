.PHONY: create_clusters
create_clusters:
	kind create cluster --config kind-cluster-with-extramounts.yaml
	./kind-with-registry.sh

.PHONY: clean
clean:
	docker rm -f /kind-registry
	./kind-local/teardown-kind-with-registry.sh
	kind delete clusters --all