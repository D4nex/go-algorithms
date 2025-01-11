package main

import(
	"fmt")

func XORencrypt(message, key string) string {
	//Convert message string to bytes
	//Convert key string to bytess
	//XORED message with key and return encrypted message(m ^ k = e)

	message2bytes := []byte(message)
	key2bytes := []byte(key)

	encrypted := make([]byte, len(message2bytes))

	for i := range message2bytes {
		encrypted[i] = message2bytes[i] ^ key2bytes[i%len(key2bytes)]
	}

	return string(encrypted)
}

func XORdecrypt(encrypted, key string) string {
	//Reuse the XORencrypt func
	//XORED encrypted message with the key and return plaintext message(e ^ k = m)
	return XORencrypt(encrypted, key)
}

func main() {

	message := "Hello Everyone!"
	key := "My secret key"

	encrypted := XORencrypt(message, key)
	fmt.Printf("Encrypted message: %s\n", encrypted)

	decrypted := XORdecrypt(encrypted, key)
	fmt.Printf("Decrypted message: %s\n", decrypted)
}
