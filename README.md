## Clean OS memory cache on Kubernetes node

DeamonSet which pulls memory statistics of kubernetes node and flush it if limit is reached.

Details: https://linux-mm.org/Drop_Caches

### Usage

```
Usage of ./flusher:
  -drop_caches_file_path string
    	Mounted host file path (default "/var/host_sys_vm/drop_caches")
  -interval int
    	Timeout in seconds (default 360)
  -limit string
    	Maximum cache buffer size (default "10 MB")
  -reset_cache_option int
    	Option to reset (default 1)
```

### Build

`make all`

### Install

`helm install --name k8s-memory-flusher chart/`

`helm upgrade k8s-memory-flusher chart/`

`helm delete k8s-memory-flusher`
