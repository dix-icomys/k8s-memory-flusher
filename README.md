## Flush system memory cache on kubernetes worker

### Configuration

ENV variables:

- INTERVAL (Default 360) Timeout in seconds
- BUFFERS_LIMIT (Default 100 MB) Maximum cache buffer size
- RESET_CACHE_OPTION (Default 1) Option to reset, Details: https://linux-mm.org/Drop_Caches
- DROP_CACHES_FILE_PATH (/var/docker_host_drop_caches) Mounted host file path

### Build

make all

### Install to k8s

helm install --name k8s-memory-flusher chart/
