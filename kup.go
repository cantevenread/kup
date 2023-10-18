package kup

import "fmt"

func Name() string {
	return "kup"
}

type Region string

type URL string

const (
	NA1  URL = "https://na1.api.riotgames.com/"
	BR1  URL = "https://br1.api.riotgames.com/"
	EUN1 URL = "https://eun1.api.riotgames.com/"
	EUW1 URL = "https://euw1.api.riotgames.com/"
	JP1  URL = "https://jp1.api.riotgames.com/"
	KR   URL = "https://kr.api.riotgames.com/"
	LA1  URL = "https://la1.api.riotgames.com/"
	LA2  URL = "https://la2.api.riotgames.com/"
	OC1  URL = "https://oc1.api.riotgames.com/"
	TR1  URL = "https://tr1.api.riotgames.com/"
	RU   URL = "https://ru.api.riotgames.com/"
	PH2  URL = "https://ph2.api.riotgames.com/"
	SG2  URL = "https://sg2.api.riotgames.com/"
	TH2  URL = "https://th2.api.riotgames.com/"
	TW2  URL = "https://tw2.api.riotgames.com/"
	VN2  URL = "https://vn2.api.riotgames.com/"
)

type KupClient struct {
	RGAPI     string
	Region    Region
	RegionURL URL
}

func NewKupClient(rgapi string, region Region) (KupClient, error) {
	regionURL := func() URL {
		switch region {
		case "NA1":
			return NA1
		case "BR1":
			return BR1
		case "EUN1":
			return EUN1
		case "EUW1":
			return EUW1
		case "JP1":
			return JP1
		case "KR":
			return KR
		case "LA1":
			return LA1
		case "LA2":
			return LA2
		case "OC1":
			return OC1
		case "TR1":
			return TR1
		case "RU":
			return RU
		case "PH2":
			return PH2
		case "SG2":
			return SG2
		case "TH2":
			return TH2
		case "TW2":
			return TW2
		case "VN2":
			return VN2
		default:
			return NA1

		}
	}
	switch region {
	case "NA1", "BR1", "EUN1", "EUW1", "JP1", "KR", "LA1", "LA2", "OC1", "TR1", "RU", "PH2", "SG2", "TH2", "TW2", "VN2":
		return KupClient{
			RGAPI:     rgapi,
			Region:    region,
			RegionURL: regionURL(),
		}, nil
	default:
		return KupClient{}, fmt.Errorf("invalid region: %s", region)
	}
}
