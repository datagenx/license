package generate

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/datagenx/license-generator/internal/initializers"
	"github.com/datagenx/license-generator/models"
)

type Slic struct {
	*models.SignedLicense
}

type Rlic struct {
	*models.RawLicense
}

func (rl *Rlic) Generate() (Slic, string, error) {

	var sl = Slic{}

	licstring, err := rl.string()
	if err != nil {
		return sl, "", err
	}

	st := initializers.Setup()
	signature := st.Sign(licstring)

	err = sl.marshal(licstring)
	if err != nil {
		fmt.Println("marshaling sl", err)
		return sl, "", err
	}
	sl.Signature = signature

	lic, err := sl.string()
	if err != nil {
		return sl, "", err
	}
	log.Println(string(lic))
	return sl, string(lic), nil
}

func (rl *Rlic) string() ([]byte, error) {
	licenseAsString, err := json.Marshal(rl)
	if err != nil {
		return nil, err
	}

	return licenseAsString, nil
}

func (sl *Slic) string() ([]byte, error) {
	license, err := json.Marshal(sl)
	if err != nil {
		return nil, err
	}

	return license, nil
}

func (sl *Slic) marshal(licstring []byte) error {
	if err := json.Unmarshal(licstring, sl); err != nil {
		return err
	}
	return nil
}
