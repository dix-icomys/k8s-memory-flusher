package main

import (
  "fmt"
  "io/ioutil"
  "github.com/c2h5oh/datasize"
  "github.com/shirou/gopsutil/mem"
  "flag"
  "strconv"
  "time"
)

var interval = flag.Int("interval", 360, "Timeout in seconds")
var reset_cache_option = flag.Int("reset_cache_option", 1, "Option to reset")
var buffers_limit_string = flag.String("buffers_limit", "10 MB", "Maximum cache buffer size")
var cached_limit_string = flag.String("cached_limit", "900 MB", "Maximum cached memory size")
var drop_caches_file_path = flag.String("drop_caches_file_path", "/var/host_sys_vm/drop_caches", "Mounted host file path")

func main() {
  flag.Parse()

  var buffers_limit datasize.ByteSize
  var cached_limit datasize.ByteSize

  buffers_limit.UnmarshalText([]byte(*buffers_limit_string))
  cached_limit.UnmarshalText([]byte(*cached_limit_string))

  for {
    v, _ := mem.VirtualMemory()

    fmt.Printf("Buffers: %v\n", (datasize.ByteSize(v.Buffers) * datasize.B).String())
    fmt.Printf("Cached: %v\n", (datasize.ByteSize(v.Cached) * datasize.B).String())

    if (datasize.ByteSize(v.Buffers) > buffers_limit || datasize.ByteSize(v.Cached) > cached_limit) {
      fmt.Println("Cleaning memory...")

      ioutil.WriteFile(*drop_caches_file_path, []byte(strconv.Itoa(*reset_cache_option)), 0644)
    }

    // Sleep until interval
    fmt.Println("Sleeping for", *interval, "seconds")
    time.Sleep(time.Duration(*interval) * time.Second)
  }
}
