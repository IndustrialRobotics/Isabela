package main

import (
	"crypto/md5"
	"encoding/hex"
	"crypto/aes"
	"crypto/cipher"
	"io"
	"crypto/rand"
	"fmt"
	"os"
	"io/ioutil"
)


/*Encrypt And Decrypt Data In A Golang Application With The Crypto Packages - YouTube -
turn OFF Windows Defender virus protection before running

output:
[137 217 113 202 98 70 71 173 130 158 242 138 166 60 169 227 207 226 203 50 193 133 159 24 206 18 55 249 166 219 26 96 235 194 188 175 139 170 70 255 233 25 161 169 117 222 164 211 218 129 210 20 17 202 27 58 224 175 37 172 125 10 214 53 65 147 136 59 200 173 9 189 102 211 108 64 157 146 113 211 245 173 195 183 127 253 199 8 11 187 206 54 236 236 1 110 167 194 218 8 37 177 38 178 236 189 105 112 125 159 103 212 92 6 15 185 64 120 89 128 55 56 148 105 178 95 219 243 207 255 52 74 0 96 173 167 111 146 202 90 138 28 155 17 229 229 50 141 180 50 70 232 104 57 239 53 44 248 7 22 219 217 51 99 172 166 145 24 23 5 41 61 54 250 138 121 166 86 58 15 33 158 148 54 79 45 84 205 56 22 123 63 211 179 18 202 83 140 14 28 3 233 246 249 75 61 15 54 139 185 177 15 23 130 8 114 172 150 5 19 59 248 25 217 119 226 159 15 173 6 182 136 194 5 3 92 228 181 109 149 67 183 69 213 113 133 245 122 192 241 138 41 244 179 223 36 3 228 187 152 103 38 247 177 117 25 251 167 75 115 127 129 255 15 210 245 99 89 232 229 65 28 111 103 116 56 143 87 221 37 184 90 164 230 114 222 194 106 160 189 57 133 243 71 237 24 176 223 161 0 188 101 170 106 160 217 224 113 178 246 181 242 135 208 125 119 192 110 61 134 91 171 23 5 65 88 250 107 7 37 111 125 251 176 138 93]
��q�bFG����<�����2����7���`�¼���F����uޤ�ځ��:�%�}
�5A��;ȭ	�f�l@��q���÷���6��n��%�&��ip}�g�\�@xY�78�i�_����4J `��o��Z����2��2F�h9�5,���3c���)=6��y�V:!��6O-T�8{?ӳ�S����K=6���r��;��w�����\�m�C�E�q��z��)���$付g&��u��Ks����cY��Aogt8�W�%�Z��r��j��9��G��ߡ �e�j���q�����}w�n=�[�AX�k%o}���]*/

func createHash(key string) string {
	hasher := md5.New()
	hasher.Write([]byte(key))
	return hex.EncodeToString(hasher.Sum(nil))
}

func encrypt(data []byte, passphrase string) []byte {
	block, _ := aes.NewCipher([]byte(createHash(passphrase)))
	gcm, _ := cipher.NewGCM(block)

	nonce := make([]byte, gcm.NonceSize())
	io.ReadFull(rand.Reader, nonce)
	ciphertext := gcm.Seal(nonce, nonce, data, nil)

	return ciphertext

}

func decrypt(data []byte, passphrase string) []byte {
	key := []byte(createHash(passphrase))
	block, _ := aes.NewCipher(key)
	gcm, _ := cipher.NewGCM(block)
	nonceSize := gcm.NonceSize()

	nonce, ciphertext := data[:nonceSize], data[nonceSize:]
	plaintext, _ := gcm.Open(nil, nonce, ciphertext, nil)
	return plaintext
}

func encryptFile(filename string, data []byte, passphrase string) {
	f, _ := os.Create(filename)
	defer f.Close()
	f.Write(encrypt(data, passphrase))
}

func decryptFile(filename string, passphrase string) []byte {
	data, _ := ioutil.ReadFile(filename)
	return decrypt(data, passphrase)
}

func main() {

	ciphertext := encrypt([]byte("Modbus communication protocol was developed in 1979, and almost 40 years later, it is still used on many industrial sites. It is pretty ... Modbus can also be used on an Ethernet network called “Modbus TCP”. ... For instance, you may have a chiller, whose manual says that the cooling setpoint is register “2”."), "password")

	fmt.Println(ciphertext)
	fmt.Println(string(ciphertext))

	plaintext := decrypt(ciphertext, "password")
	fmt.Println(string(plaintext))

	encryptFile("andycrypt.txt", []byte("apartment #244 on green island has been foreclosure"), "te42ta")

	plainfile := decryptFile("andycrypt.txt", "te42ta")
	fmt.Println(plainfile)

}
