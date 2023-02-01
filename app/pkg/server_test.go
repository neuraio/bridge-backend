package pkg

import "testing"

func TestServer_Start(t *testing.T) {
	NewServer([]string{"http://103.23.44.17:18575"}).Start()
}
