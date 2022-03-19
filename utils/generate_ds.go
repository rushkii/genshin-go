package utils

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"time"
)

// Creates a new ds for authentication.
func GenerateDs() string {
	t := time.Now().Unix() // current timestamp
	r := Base64(6)         // 6 random chars
	data := fmt.Sprintf("salt=6s25p5ox5y14umn1p61aqyyvbvvl3lrt&t=%d&r=%s", t, r)
	sum := md5.Sum([]byte(data))
	h := hex.EncodeToString(sum[:])
	return fmt.Sprintf("%d,%s,%s", t, r, h)
}

// func GenerateDsCn() string {
// 	t := time.Now().Unix()       // current timestamp
// 	r := RandInt(100001, 200000) // 6 random chars
// 	data := fmt.Sprintf("salt=6cqshh5dhw73bzxn20oexa9k516chk7s&t=%v&r=%s", t, r)
// 	sum := md5.Sum([]byte(data))
// 	h := string(sum[:])
// 	return fmt.Sprintf("%v,%s,%s", t, r, h)
// }

// def generate_cn_ds(salt: str, body: Any = None, query: Mapping[str, Any] = None) -> str:
//     """Creates a new chinese ds for authentication."""
//     t = int(time.time())
//     r = random.randint(100001, 200000)
//     b = json.dumps(body) if body else ""
//     q = "&".join(f"{k}={v}" for k, v in sorted(query.items())) if query else ""

//     h = hashlib.md5(f"salt={salt}&t={t}&r={r}&b={b}&q={q}".encode()).hexdigest()
//     return f"{t},{r},{h}"
