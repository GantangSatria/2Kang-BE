package middleware

import (
    "2Kang/pkg/utils"
    "fmt"
    "github.com/gofiber/fiber/v3"
    "strconv"
)

// func JWTProtected(c fiber.Ctx) error {
//     // 1. Validasi Token
//     _, claims, err := utils.ValidateJWT(c)
//     if err != nil {
//         return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized: Token invalid"})
//     }

//     // 2. Ambil ID dari claim
//     // Pastikan saat login kamu set key-nya "id" (huruf kecil)
//     idFloat := claims["id"].(float64) 
//     idClaim := uint(idFloat)
    
//     if idClaim == 0 {
//         fmt.Println("DEBUG JWT: ID Claim is nil!")
//         return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Token tidak memiliki ID"})
//     }

//     //var finalUserID uint

//     // 3. LOGIKA ANTI-CRASH (Handle Float64 dari JSON JWT)
//     // switch v := idClaim.(type) {
//     // case float64:
//     //     finalUserID = uint(v) // Ini yang biasanya terjadi
//     // case string:
//     //     // Jaga-jaga kalau token simpan id sebagai string "3"
//     //     parsed, err := strconv.ParseUint(v, 10, 32)
//     //     if err != nil {
//     //         return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Format ID di token salah"})
//     //     }
//     //     finalUserID = uint(parsed)
//     // case int:
//     //     finalUserID = uint(v)
//     // default:
//     //     fmt.Printf("DEBUG JWT: Tipe data aneh: %T\n", v)
//     //     return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Tipe ID tidak dikenali"})
//     // }

//     // 4. Simpan di Locals sebagai UINT
//     c.Locals("user_id", idClaim)
    
//     // Debug print supaya kelihatan di terminal
//     fmt.Printf("DEBUG MIDDLEWARE: Berhasil dapat ID: %d\n", idClaim)
//     fmt.Printf("DEBUG CLAIM id tipe: %T, value: %v\n", idClaim, idClaim)

//     return c.Next()
// }

func JWTProtected(c fiber.Ctx) error {
    // 1. Validasi token
    _, claims, err := utils.ValidateJWT(c)
    if err != nil {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
    }

    idRaw, ok := claims["id"]
    if !ok {
        return c.Status(401).JSON(fiber.Map{"error": "ID not found in token"})
    }

    var id uint

    switch v := idRaw.(type) {
    case float64:
        id = uint(v)
    case int:
        id = uint(v)
    case uint:
        id = v
    case string:
        parsed, err := strconv.ParseUint(v, 10, 64)
        if err != nil {
            return c.Status(400).JSON(fiber.Map{"error": "Invalid ID format"})
        }
        id = uint(parsed)
    default:
        return c.Status(400).JSON(fiber.Map{"error": "Unknown ID type"})
    }

    // simpan
    c.Locals("user_id", id)

    fmt.Printf("DEBUG JWT ID → %d (%T)\n", id, idRaw)

    return c.Next()
}
