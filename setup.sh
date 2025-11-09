./create_kind_cluster.sh

kubectl create namespace default
bazel run //app:deploy.apply
