/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package gxnet

import (
	"net"
	"strings"
)

import (
	perrors "github.com/pkg/errors"
)

var (
	privateBlocks []*net.IPNet
)

func init() {
	for _, b := range []string{"10.0.0.0/8", "172.16.0.0/12", "192.168.0.0/16"} {
		if _, block, err := net.ParseCIDR(b); err == nil {
			privateBlocks = append(privateBlocks, block)
		}
	}
}

// GetLocalIP get local ip
func GetLocalIP() (string, error) {
	faces, err := net.Interfaces()
	if err != nil {
		return "", perrors.WithStack(err)
	}

	var addr net.IP
	for _, face := range faces {
		if !isValidNetworkInterface(face) {
			continue
		}

		addrs, err := face.Addrs()
		if err != nil {
			return "", perrors.WithStack(err)
		}

		if ipv4, ok := getValidIPv4(addrs); ok {
			addr = ipv4
			if isPrivateIP(ipv4) {
				return ipv4.String(), nil
			}
		}
	}

	if addr == nil {
		return "", perrors.Errorf("can not get local IP")
	}

	return addr.String(), nil
}

func isPrivateIP(ip net.IP) bool {
	for _, priv := range privateBlocks {
		if priv.Contains(ip) {
			return true
		}
	}
	return false
}

func getValidIPv4(addrs []net.Addr) (net.IP, bool) {
	for _, addr := range addrs {
		var ip net.IP

		switch v := addr.(type) {
		case *net.IPNet:
			ip = v.IP
		case *net.IPAddr:
			ip = v.IP
		}

		if ip == nil || ip.IsLoopback() {
			continue
		}

		ip = ip.To4()
		if ip == nil {
			// not an valid ipv4 address
			continue
		}

		return ip, true
	}
	return nil, false
}

func isValidNetworkInterface(face net.Interface) bool {
	if face.Flags&net.FlagUp == 0 {
		// interface down
		return false
	}

	if face.Flags&net.FlagLoopback != 0 {
		// loopback interface
		return false
	}

	if strings.Contains(strings.ToLower(face.Name), "docker") {
		return false
	}

	return true
}

// refer from https://github.com/facebookarchive/grace/blob/master/gracenet/net.go#L180
func IsSameAddr(addr1, addr2 net.Addr) bool {
	if addr1.Network() != addr2.Network() {
		return false
	}

	addr1s := addr1.String()
	addr2s := addr2.String()
	if addr1s == addr2s {
		return true
	}

	// This allows for ipv6 vs ipv4 local addresses to compare as equal. This
	// scenario is common when listening on localhost.
	const ipv6prefix = "[::]"
	addr1s = strings.TrimPrefix(addr1s, ipv6prefix)
	addr2s = strings.TrimPrefix(addr2s, ipv6prefix)
	const ipv4prefix = "0.0.0.0"
	addr1s = strings.TrimPrefix(addr1s, ipv4prefix)
	addr2s = strings.TrimPrefix(addr2s, ipv4prefix)
	return addr1s == addr2s
}

// ListenOnTCPRandomPort a tcp server listening on a random port by tcp protocol
func ListenOnTCPRandomPort(ip string) (*net.TCPListener, error) {
	localAddr := net.TCPAddr{
		IP:   net.IPv4zero,
		Port: 0,
	}
	if len(ip) > 0 {
		localAddr.IP = net.ParseIP(ip)
	}

	// on some containers, u can not bind an random port by the following clause.
	// listener, err := net.Listen("tcp", ":0")

	return net.ListenTCP("tcp4", &localAddr)
}

// ListenOnUDPRandomPort a udp endpoint listening on a random port
func ListenOnUDPRandomPort(ip string) (*net.UDPConn, error) {
	localAddr := net.UDPAddr{
		IP:   net.IPv4zero,
		Port: 0,
	}
	if len(ip) > 0 {
		localAddr.IP = net.ParseIP(ip)
	}

	return net.ListenUDP("udp4", &localAddr)
}
