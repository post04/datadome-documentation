package main

type payload struct {
	Print         *fingerprint
	EventCounters *events
	JsType        string
	Cookie        string
	SiteKey       string
	Host          string
	Path          string
	ResponsePage  string
	Version       string
}

type fingerprint struct {
	TimeSinceLastEvent float64 `json:"ttst"`
	Ifov               bool    `json:"ifov"`
	TimeToDoWebGL      float64 `json:"tagpu"`
	Glvd               string  `json:"glvd"`
	Glrd               string  `json:"glrd"`
	Hc                 int     `json:"hc"`
	BrOh               int     `json:"br_oh"`
	BrOw               int     `json:"br_ow"`
	Ua                 string  `json:"ua"`
	Wbd                bool    `json:"wbd"`
	Wdif               bool    `json:"wdif"`
	Wdifrm             bool    `json:"wdifrm"`
	Npmtm              bool    `json:"npmtm"`
	BrH                int     `json:"br_h"`
	BrW                int     `json:"br_w"`
	CookieLength       int     `json:"nddc"`
	RsH                int     `json:"rs_h"`
	RsW                int     `json:"rs_w"`
	RsCd               int     `json:"rs_cd"`
	Phe                bool    `json:"phe"`
	Nm                 bool    `json:"nm"`
	Jsf                bool    `json:"jsf"`
	Lg                 string  `json:"lg"`
	Pr                 int     `json:"pr"`
	ArsH               int     `json:"ars_h"`
	ArsW               int     `json:"ars_w"`
	TimezoneOffset     int     `json:"tz"`
	StrSs              bool    `json:"str_ss"`
	StrLs              bool    `json:"str_ls"`
	StrIdb             bool    `json:"str_idb"`
	StrOdb             bool    `json:"str_odb"`
	Plgod              bool    `json:"plgod"`
	NavPluginLength    int     `json:"plg"`
	Plgne              bool    `json:"plgne"`
	Plgre              bool    `json:"plgre"`
	Plgof              bool    `json:"plgof"`
	Plggt              bool    `json:"plggt"`
	Pltod              bool    `json:"pltod"`
	Hcovdr             bool    `json:"hcovdr"`
	Hcovdr2            bool    `json:"hcovdr2"`
	Plovdr             bool    `json:"plovdr"`
	Plovdr2            bool    `json:"plovdr2"`
	Ftsovdr            bool    `json:"ftsovdr"`
	Ftsovdr2           bool    `json:"ftsovdr2"`
	Lb                 bool    `json:"lb"`
	EvalLength         int     `json:"eva"`
	Lo                 bool    `json:"lo"`
	TsMtp              int     `json:"ts_mtp"`
	TsTec              bool    `json:"ts_tec"`
	TsTsa              bool    `json:"ts_tsa"`
	Vnd                string  `json:"vnd"`
	Bid                string  `json:"bid"`
	Mmt                string  `json:"mmt"`
	Plu                string  `json:"plu"`
	Hdn                bool    `json:"hdn"`
	Awe                bool    `json:"awe"`
	Geb                bool    `json:"geb"`
	Dat                bool    `json:"dat"`
	Med                string  `json:"med"`
	Aco                string  `json:"aco"`
	Acmp               string  `json:"acmp"`
	Acw                string  `json:"acw"`
	Acma               string  `json:"acma"`
	Acaa               string  `json:"acaa"`
	Ac3                string  `json:"ac3"`
	Acf                string  `json:"acf"`
	Acmp4              string  `json:"acmp4"`
	Acmp3              string  `json:"acmp3"`
	Acwm               string  `json:"acwm"`
	Acots              string  `json:"acots"`
	Acmpts             string  `json:"acmpts"`
	Acwts              string  `json:"acwts"`
	Acmats             string  `json:"acmats"`
	Acaats             string  `json:"acaats"`
	Ac3Ts              string  `json:"ac3ts"`
	Acfts              string  `json:"acfts"`
	Acmp4Ts            string  `json:"acmp4ts"`
	Acmp3Ts            string  `json:"acmp3ts"`
	Acwmts             string  `json:"acwmts"`
	Vco                string  `json:"vco"`
	Vcots              bool    `json:"vcots"`
	Vch                string  `json:"vch"`
	Vchts              bool    `json:"vchts"`
	Vcw                string  `json:"vcw"`
	Vcwts              bool    `json:"vcwts"`
	Vc3                string  `json:"vc3"`
	Vc3Ts              bool    `json:"vc3ts"`
	Vcmp               string  `json:"vcmp"`
	Vcmpts             bool    `json:"vcmpts"`
	Vcq                string  `json:"vcq"`
	Vcqts              bool    `json:"vcqts"`
	Vc1                string  `json:"vc1"`
	Vc1Ts              bool    `json:"vc1ts"`
	Dvm                int     `json:"dvm"`
	Sqt                bool    `json:"sqt"`
	So                 string  `json:"so"`
	Wdw                bool    `json:"wdw"`
	Cokys              string  `json:"cokys"`
	Ecpc               bool    `json:"ecpc"`
	Lgs                bool    `json:"lgs"`
	Lgsod              bool    `json:"lgsod"`
	Psn                bool    `json:"psn"`
	Edp                bool    `json:"edp"`
	Addt               bool    `json:"addt"`
	Wsdc               bool    `json:"wsdc"`
	Ccsr               bool    `json:"ccsr"`
	Nuad               bool    `json:"nuad"`
	Bcda               bool    `json:"bcda"`
	Idn                bool    `json:"idn"`
	Capi               bool    `json:"capi"`
	Svde               bool    `json:"svde"`
	Vpbq               bool    `json:"vpbq"`
	Ucdv               bool    `json:"ucdv"`
	Spwn               bool    `json:"spwn"`
	Emt                bool    `json:"emt"`
	Bfr                bool    `json:"bfr"`
	Dbov               bool    `json:"dbov"`
	Prm                bool    `json:"prm"`
	Tzp                string  `json:"tzp"`
	Cvs                bool    `json:"cvs"`
	Usb                string  `json:"usb"`
	TimeSeconds        int64   `json:"jset"`
	// new values I saw on footlocker on le type
	NormalError     string  `json:"cfpfe,omitempty"`
	ErrorWithLineNo string  `json:"stcfp,omitempty"`
	Dcok            string  `json:"dcok,omitempty"`
	MpCx            int     `json:"mp_cx,omitempty"`
	MpCy            int     `json:"mp_cy,omitempty"`
	MpTr            bool    `json:"mp_tr,omitempty"`
	MpMx            *int    `json:"mp_mx,omitempty"`
	MpMy            *int    `json:"mp_my,omitempty"`
	MpSx            int     `json:"mp_sx,omitempty"`
	MpSy            int     `json:"mp_sy,omitempty"`
	MmMd            int     `json:"mm_md,omitempty"`
	EsSigmdn        float64 `json:"es_sigmdn,omitempty"`
	EsMumdn         float64 `json:"es_mumdn,omitempty"`
	EsDistmdn       float64 `json:"es_distmdn,omitempty"`
	EsAngsmdn       float64 `json:"es_angsmdn,omitempty"`
	EsAngemdn       float64 `json:"es_angemdn,omitempty"`
}

type events struct {
	// just do 2800-3200 cause it seems that's generally what I get normally before second request
	MouseMove     int        `json:"mousemove"`
	Click         int        `json:"click"`
	Scroll        int        `json:"scroll"`
	Touchstart    int        `json:"touchstart"`
	Touchend      int        `json:"touchend"`
	Touchmove     int        `json:"touchmove"`
	Keydown       int        `json:"keydown"`
	Keyup         int        `json:"keyup"`
	PageTimestamp int64      `json:"-"`
	Width         int        `json:"-"`
	Height        int        `json:"-"`
	Sigmdn        float64    `json:"-"`
	Mumdn         float64    `json:"-"`
	Dismdn        float64    `json:"-"`
	Angsmdn       float64    `json:"-"`
	Angemdn       float64    `json:"-"`
	LastPosition  *mouseMove `json:"-"`
	WeirdMD       *mouseMove `json:"-"`
	WeirdMDValue  int        `json:"-"`
}

type mouseMove struct {
	X         float64
	Y         float64
	MovementX int
	MovementY int
	ScreenX   int
	ScreenY   int
	Timestamp float64
}
