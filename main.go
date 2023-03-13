// This program demonstrates attaching an eBPF program to a network interface to block source IPs/CIDRs
// This example depends on bpf_link, available in Linux kernel version 5.7 or newer.
package main

//go:generate go run github.com/cilium/ebpf/cmd/bpf2go -cc $BPF_CLANG -cflags $BPF_CFLAGS bpf xdp_synproxy_kern.c -- -I./headers

func main() {
}
