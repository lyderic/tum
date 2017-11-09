package main

import(
  "testing"
)

func TestTunnel(t *testing.T) {

  tunnel := Tunnel {
    Id: 42,
    Description: "Dummy Tunnel for Testing",
    Host: "somehost",
    LocalPort: 8000,
    RemotePort: 8000,
  }

  if tunnel.getId() != 42 {
    t.Error("A tunnel could not be created")
  }

  if tunnel.toString() != "42 somehost 8000 8000 'Dummy Tunnel for Testing'" {
    t.Error("Incorrect tunnel string representation")
  }

}
