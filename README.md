# oci-token-cache

Cache oci login token. This command cache oci login token into `~/.oci/token-cache.json` and re-use for kubectl.

## Usage

Currently, your `~/.kube/config` is like below.

```
---
apiVersion: v1
kind: ""
clusters:
- name: cluster-zzzzzzzz
  cluster:
    server: https://xxx.xxx.xxx.xxx:6443
    certificate-authority-data: XXXXXXXXXXXXXXXXXXXXXXXXXX
users:
- name: user-zzzzzzzz
  user:
    exec:
      apiVersion: client.authentication.k8s.io/v1beta1
      command: oci
      args:
      - ce
      - cluster
      - generate-token
      - --cluster-id
      - ocid1.cluster.oc1.yyyyyyy.YYYYYYYYYYYYYYYYYYYY
      - --region
      - yyyyyyy
      env: []
contexts:
- name: context-zzzzzzzz
  context:
    cluster: cluster-zzzzzzzz
    user: user-zzzzzzzz
current-context: context-zzzzzzzz
```

Just replace `command` to `oci-token-cache`, and shift `args` like below.

```
users:
- name: user-zzzzzzzz
  user:
    exec:
      apiVersion: client.authentication.k8s.io/v1beta1
      command: oci-token-cache
      args:
      - oci
      - ce
      - cluster
      - generate-token
      - --cluster-id
      - ocid1.cluster.oc1.yyyyyyy.YYYYYYYYYYYYYYYYYYYY
      - --region
      - yyyyyyy
      env: []
```

## Installation

```
go install github.com/mattn/oci-token-cache@latest
```

## License

MIT

## Author

Yasuhiro Matsumoto (a.k.a. mattn)
