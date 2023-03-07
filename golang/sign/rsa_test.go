package sign

import (
	"fmt"
	"testing"
)

func TestGenerateRsaKey(t *testing.T) {
	pri, pub := GenerateRsaKey(1024)

	fmt.Println(`private_key: `)
	fmt.Println(string(pri))
	fmt.Println(`public_key: `)
	fmt.Println(string(pub))

}
