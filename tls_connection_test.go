package tls_example

import "testing"

func TestTlsHandshake(t *testing.T)  {
	conn, err := createTlsConnection()
	if err != nil {
		t.Fatalf("Error: %s", err)
	}

	if conn == nil {
		t.Fatalf("Connection is nil!")
	}

	err = conn.Close()
	if err != nil {
		t.Fatalf("Unable to close connection")
	}
}
