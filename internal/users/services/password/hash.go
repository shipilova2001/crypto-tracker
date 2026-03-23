package services
import bscrypt "golang.org/x/crypto/bcrypt"

const (
	MinCostHash     int = 4
	MaxCostHash     int = 31 // максимально допустимая стоимость, передаваемая в GenerateFromPassword
	DefaultCostHash int = 10 // стоимость, которая будет фактически установлена, если в GenerateFromPassword будет передана стоимость ниже MinCost
)

func GetHash(password string) []byte {
	hash, err := bscrypt.GenerateFromPassword([]byte(password), DefaultCostHash)
	if err != nil {
		return nil
	}
	return hash
}