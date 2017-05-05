## Clean OS memory cache on Kubernetes node

DeamonSet which pulls memory statistics of kubernetes node and flush it if limit is reached.

Details: https://linux-mm.org/Drop_Caches

### Usage

```
Usage of ./flusher:
  -buffers_limit string
      Maximum cache buffer size (default "10 MB")
  -cached_limit string
      Maximum cached memory size (default "900 MB")
  -drop_caches_file_path string
      Mounted host file path (default "/var/host_sys_vm/drop_caches")
  -interval int
      Timeout in seconds (default 360)
  -reset_cache_option int
      Option to reset (default 1)
```

### Build

`make all`

### Install

`helm install --name k8s-memory-flusher chart/`

`helm upgrade k8s-memory-flusher chart/`

`helm delete k8s-memory-flusher`
