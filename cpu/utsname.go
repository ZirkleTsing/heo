package cpu

import "github.com/mcai/acogo/cpu/mem"

type Utsname struct {
	Sysname    string
	Nodename   string
	Release    string
	Version    string
	Machine    string
	Domainname string
}

func NewUtsname() *Utsname {
	var utsname = &Utsname{
	}

	return utsname
}

func (utsname *Utsname) GetBytes(littleEndian bool) []byte {
	var sysname_buf = []byte(utsname.Sysname + "\x00")
	var nodename_buf = []byte(utsname.Nodename + "\x00")
	var release_buf = []byte(utsname.Release + "\x00")
	var version_buf = []byte(utsname.Version + "\x00")
	var machine_buf = []byte(utsname.Machine + "\x00")
	var domainname_buf = []byte(utsname.Domainname + "\x00")

	var _sysname_size = uint32(64 + 1)
	var size_of = _sysname_size * 6

	var buf = make([]byte, size_of)

	var memory = mem.NewSimpleMemory(littleEndian, buf)

	memory.WriteBlockAt(0, _sysname_size, sysname_buf)
	memory.WriteBlockAt(_sysname_size, _sysname_size, nodename_buf)
	memory.WriteBlockAt(_sysname_size * 2, _sysname_size, release_buf)
	memory.WriteBlockAt(_sysname_size * 3, _sysname_size, version_buf)
	memory.WriteBlockAt(_sysname_size * 4, _sysname_size, machine_buf)
	memory.WriteBlockAt(_sysname_size * 5, _sysname_size, domainname_buf)

	return buf
}
