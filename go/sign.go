package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
	"errors"
	"fmt"
)

var (
	PublicKey  = "-----BEGIN PUBLIC KEY-----\nMIICojANBgkqhkiG9w0BAQEFAAOCAo8AMIICigKCAoEAqgShz8zOTz2AZlNXIzLp\noFhrbTH2eN2bVlQvPA1fTw+U7+KD6OdN1isgwaDJKPz2gV7lF5FXE+wgb8VuvaoF\nO+/AiKvFYQNOqFmDKjXavFIuv6bBphL/SfcIwcVfAbwkIJCxT7R2YHfUDQWvBlXu\nYtoaJshkGU15g+Wn2d5AyuwK6KjqDyC9Frp8u7kHKihi80w3G4e1c+06KN15FuTb\nxiC1+CuasQA7350PFjOHUrMk5Ap7rzUH5rIXwFKmaIpoN7Vu3/HIKMFELBjqf7jY\n7zId3QFCqc4tHH2uHV0C17bN3CxdZdLs3fprF2iuhfF3M6nM0kU1a5I5c0Vz56B/\nps5dOqpCdiqlATa18jTtYLy/xn61fX46a5iaeauJbinBOcEDwEE9oOkUfjJuQq+X\nwo/SIGcT00lFTEJdSt9eb9/m0ck1xcyiSWhVaFOSIaRPOvsl+5uoowgE1ofYBE3i\nF6VtyZQVpwUaeJgxF/i348n6YclYHhSB1RBqEGfWHuRhsE9rvuMoKpXRCsL5BGCv\nEe/SspH+imrAeIldOpA/6RKGYNdMuswhR5rPF+UUEAYWzqzagISuvHKjECy0kcTx\njK54caMzm5UzNQAfPS/pIWbQdzbfQbx/3tvrczK32SQB/7z1CnDKuJDWyOLgPa0W\nquJEbtsY5eH61FhzEBHt3IOOaNPuvLg2aHnSPMAdpJfUQsCykDoML2g16c4EvqrV\nbMDDl25HSg9AZm2gze91aDGgWqY8BgUDeO06qk2gRTolwZDSgAqEb0tAltBY+hnN\n60e3gFUz6wAOjbiqLSf5CavrLISYBrl6/Ut3cfin4j8pfy2xCMFzytGWV5QCL65v\nQwIDAQAB\n-----END PUBLIC KEY-----"
	PrivateKey = "-----BEGIN RSA PRIVATE KEY-----\nMIILagIBAAKCAoEAqgShz8zOTz2AZlNXIzLpoFhrbTH2eN2bVlQvPA1fTw+U7+KD\n6OdN1isgwaDJKPz2gV7lF5FXE+wgb8VuvaoFO+/AiKvFYQNOqFmDKjXavFIuv6bB\nphL/SfcIwcVfAbwkIJCxT7R2YHfUDQWvBlXuYtoaJshkGU15g+Wn2d5AyuwK6Kjq\nDyC9Frp8u7kHKihi80w3G4e1c+06KN15FuTbxiC1+CuasQA7350PFjOHUrMk5Ap7\nrzUH5rIXwFKmaIpoN7Vu3/HIKMFELBjqf7jY7zId3QFCqc4tHH2uHV0C17bN3Cxd\nZdLs3fprF2iuhfF3M6nM0kU1a5I5c0Vz56B/ps5dOqpCdiqlATa18jTtYLy/xn61\nfX46a5iaeauJbinBOcEDwEE9oOkUfjJuQq+Xwo/SIGcT00lFTEJdSt9eb9/m0ck1\nxcyiSWhVaFOSIaRPOvsl+5uoowgE1ofYBE3iF6VtyZQVpwUaeJgxF/i348n6YclY\nHhSB1RBqEGfWHuRhsE9rvuMoKpXRCsL5BGCvEe/SspH+imrAeIldOpA/6RKGYNdM\nuswhR5rPF+UUEAYWzqzagISuvHKjECy0kcTxjK54caMzm5UzNQAfPS/pIWbQdzbf\nQbx/3tvrczK32SQB/7z1CnDKuJDWyOLgPa0WquJEbtsY5eH61FhzEBHt3IOOaNPu\nvLg2aHnSPMAdpJfUQsCykDoML2g16c4EvqrVbMDDl25HSg9AZm2gze91aDGgWqY8\nBgUDeO06qk2gRTolwZDSgAqEb0tAltBY+hnN60e3gFUz6wAOjbiqLSf5CavrLISY\nBrl6/Ut3cfin4j8pfy2xCMFzytGWV5QCL65vQwIDAQABAoICgQCihqUMWVNLSpaI\nutiMrGnLP1sKWn2r0uRgpT82H+5hTJrqqlXA+gRdXIMGtiY2SYN37nf/jI96Wvoy\n/sA7DOHOBwW9YL7hW2EA6/jIDA1agxFvYgCyOmzSjxVO7tdAeFLs0oS9ldynoUUT\nYcCLxO+KjcX8Fwohaf9kh9BqI906XmSbPoXpCjnYFVORvHHN6ieVPuFbLuLVvAwh\n7n2H1iUFhFWPtFc/JuMVgzVd2spc0yXL3P2ZYo94B/1Oe3PacZzozEb+S0o1M3Nw\n2CyrBdBS8Ey97SuF2bMkDjy4kyZZgcXO2JY79WM7W9yt1wMmxcpWl9ZVl5VWjMTg\nnZ8L6FC+630RzJqWDXQI8f0CWXgmRqjpHk1udpFwM8RP1RweNEojZJCqY3N/c/CE\nJr9ifisL8aJEGWVJUDyvhq1aGqdTN16LTd9zZsJBet+UW7Mfd7YFKyZF5oQpaWCg\nQ2/BFaIudg5Xa0Fo7f07moMtaT6W2OY8cs3nnrpuvyWy9PzQglY+hVt9gqYCLkf0\n6Wmc8xFU0riJwYgRPA0/2S3BO+DVvISGMXo65Ue98/XeKKRbvLoEMlL6gHS6cacX\ngBp7MRrurSWeTqJ4imH05M5gEBUTGXdsjoH5WTWn+VMiI65gv6Lzs/JtAW7T760j\ntTlwnc3g0sQtmt4M26YzY5kIqCcPbRt+SEqSUnUmqrhVB7in9c/9jyxSXwb4bPev\n9j6ZS2MZh1ZlS3cE361OIfGIJkaxxKyVwF3XUub6jXM0weOE65RGPYQ68Myu6DwE\nSt6pe7LU5Ir72+FweQxnZEw5PquabjVq7PaLArmxPtUqStTFXnHqNAo1oWaxj30P\no3bRXRAZAoIBQQDW3Cnn5nbPGUgfIbk+j53LCBs2E0Uwzlbt6Tc3XnLwhjZ9f/vH\nZNStMfO2s4gQs1Auv5AKxAxIlBnwsB2lxlxNLgLdHYz2uE7Q0qWfEk5Lsp1SLUzt\n3U+10kbfkwlf1BJfcjbGfkk//HLb7kSyaTdn8rf/eiuw15GtW2Q8XCJbZU0iM3yK\nTLfkc3RGcWCHWm41I83NM9k9us5MmdYrR/E5DlEgA61YUF/GohBppTWozGnIzxDn\nLO0gVA0CbqB4NzedXED6RGSi3D0ZkYTSHOG+D6lkLNVb9nEVfcaNFswYd7CnRDuu\n90WXocwUNyDTg829A4YZvW9yCVuBCl9KAMTmHG0ZN05/4HgmtVTxdl66YM6q/Deq\n/BhQCfpacwAMOyXacJYcaREia4NwmGC4Pd59BB1B4wsai4P0QB83DU4wxwKCAUEA\nypJxaIi7/KVs6Szh/k1BILDO9xEFe3a79XdS8LY+X5tHv4UnX92QPJgLvzWW4UmT\nXe2xWdRStS8XWa/QUjsr0lMOLBUcsXtZABE1N4/1Pmu3Nu3gsfIHLzrkPBfA2+pm\nYH8Wqh7Qd6JdvwV+Y8ymXGDeacIty4Th/+WNQXPTV+BFVSIwDnuszDHi/DdjG5hR\n94xp/N4aLoaqvaQmqVdCXwN6kyf4DuOFYwYDFtGDhs0B5lGm9sNTFAw6u83zMPUY\ntE8zjLJtyat3BBPkin0ILxZjdEDd+azdOCScHXH/NLauKdCzD8WBr++oiLHLKOK2\nm9uQwLvKB8SWD8U3vMBl3PE22oxpMvVk5SACrATtuWZUD4J5JtQyCoJRingQtBS0\nuQ8NfMwAupc683sIjZmFAfI9d5Y3fUnJsH7DSvDeCaUCggFBALMrt3pdsjezUnQ7\nkBl+c8xKfz6RxUIU32LX9yaiCWFP2oK/RjuZdszS+fQiNFYQ4c2KC1ke8UtYzWTr\nyF1kU7Y8p0CykslvGrKC5oAHKOlxEkXj5uH4iq8JofZwzAFwfL3BJbJSv0yvfSyv\nvfO1T/HL8uoFfAelTCLvMmAvSKtEh5T8sF836KWGLbFMfH2CyKJhsA3trbwPWleA\ndZt0C9FTgg0bY+Ngau3eWadvkJapl5FQMz1ENp2d2Zdsnz6AdU0xEDJi9DTnScBX\nnBYG5ayzrMMPrABbOyP71aX0G9BPhy2E9exK4Wq4DtHQwvv/ToLPPHWchiHncbni\n52Fz0L4/eC7t8vUqP6ev0RzmY8oyTkVBvD5GqROdrnrcqW6JqGBHQLaCWzqDw4ES\nYJ1jYnSICMnXxnulWi1Fh4JGAy0jAoIBQBVTEzMw6bEv7OYi/mtZ0JhTZIYnuY8E\nscamWAnU9X31B7mGl92EonFuhBYN2yM+LpA9vGS7cDV4jpm3Y1O0vOk4Kw7C7jFU\nzcqvBQ4jTmrSxge3ROcWlmEHbw1QZtH/u+U+m7Eu6z0cz4eTWXnCCmVBIUi60z1Q\ncMYScStJuR7QEhxjK/HqXQdl3QI2Bm3aPA62LlYbP2BUdtbd1+ANuoPsBtRE6d0P\nYHRLJroZNiJTpHaNc+kYKL+8hdZNWON3ebFJgSC3i80mCrJYMb/OZ7UlQzTyuytU\nYjlNyU1kObxf2re2K5NmdD+EGBbDuKJGM3j2xSc03ZCxxiZE2Zg6HL8EeYTKMSzX\nHlbsZrubDTduKT00u/I1dksa6uCLwR5j655rkjzn8M+zqdYFxqG5XoIU8RmMq6y5\nvsygm4hhEM6tAoIBQQDStUMa8HHtqF69b4z95rMhdooHfU681AWnZW8A4ssLy6Qi\nOG8ZyRB3IXifYn7H2G5HqPmT0kXsOHxAEZjOhjC9JFi3m3XbjKF3D1obOU0PCNDJ\nHW3ml4VVllVLW7GyMTtPL3V8zoM3J6bvNXpiTObTMtDdA1fsiwb6XEjEQv5XGmrq\nEDkMVLuNA729SBdAACW9NsVoB263iTecdlW88j8x2DZlQLOv5bl6F3/ektzKO3Kb\nrY6em1FZiskpCYRzysx5sbopS/AfNRoiGW+Zu0dKRpUOzQcyFylRRt9FUW9JTpmg\nAPChv5pJ4hUfd27T6hg4aly5yY7B9EYbDJkfrrLp9NlJkTpJtia5lhMHEwclE4eS\n00QwgxVu4REfDpuh0WOBAtBHu1rn6lQQtRFDzsD7xT/yoB7qQYF3m/+o0dnv3Q==\n-----END RSA PRIVATE KEY-----"
)

func RsaVerySignWithSha256(data, signData []byte, publicKey string) bool {
	block, _ := pem.Decode([]byte(publicKey))
	if block == nil {
		panic(errors.New("public key error"))
	}
	pubKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		fmt.Println(err)
		return false
	}

	hashed := sha256.Sum256(data)

	//sig, _ := base64.StdEncoding.DecodeString(string(signData))  base64 decode
	sig, _ := hex.DecodeString(string(signData)) // hex decode

	err = rsa.VerifyPKCS1v15(pubKey.(*rsa.PublicKey), crypto.SHA256, hashed[:], sig)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

func RsaSignWithSha256(data []byte, keyBytes []byte) (string, error) {
	h := sha256.New()
	h.Write(data)
	hashed := h.Sum(nil)
	block, _ := pem.Decode(keyBytes)
	if block == nil {
		return "", errors.New("private key error")
	}
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return "", err
	}

	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hashed)
	if err != nil {
		return "", err

	}
	str := hex.EncodeToString(signature)
	return str, nil
}

func RsaEncrypt(data, keyBytes []byte) ([]byte, error) {
	block, _ := pem.Decode(keyBytes)
	if block == nil {
		return nil, errors.New("public key error")
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		fmt.Println("ParsePKIXPublicKey Err:", err)
		return nil, err
	}
	pub := pubInterface.(*rsa.PublicKey)
	ciphertext, err := rsa.EncryptPKCS1v15(rand.Reader, pub, data)
	if err != nil {
		fmt.Println("EncryptPKCS1v15 Err:", err)
		return nil, err
	}
	return ciphertext, nil
}

func RsaDecrypt(ciphertext, keyBytes []byte) ([]byte, error) {
	block, _ := pem.Decode(keyBytes)
	if block == nil {
		return nil, errors.New("private key error!")
	}
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	data, err := rsa.DecryptPKCS1v15(rand.Reader, priv, ciphertext)
	if err != nil {
		return nil, err
	}
	return data, nil
}

/**
todo php
function SHA256Hex($str){
    $re=hash('sha256', $str, true);
    return bin2hex($re);
}

todo nodejs:

 let hash = crypto.createHash('sha256')
	.update(str, 'utf8')
        .digest('hex');

todo java:
https://www.geeksforgeeks.org/sha-256-hash-in-java/

todo python:
import hashlib

a_string = 'GeeksForGeeks'

hashed_string = hashlib.sha256(a_string.encode('utf-8')).hexdigest()
print(hashed_string)
*/

// todo golang
func Hash256(byte2 []byte) string {
	hash := sha256.Sum256(byte2)
	code := hex.EncodeToString(hash[:]) //hex
	return code
}
