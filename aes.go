package main
import ( "crypto/aes"; "fmt" )

func main ()  {
	plaintext := []byte("hi three")
	key := []byte("12345678")
	cipher, _ := aes.NewCipher(key)

	ciphertext := make([]byte, 16)
	cipher.Encrypt(ciphertext, plaintext)
	fmt.Printf("Ciphertext: %x\n", ciphertext)
}