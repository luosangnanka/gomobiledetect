package gomobiledetect

import (
	"net/http"
	"testing"
)

var httpRequest = &http.Request{}

type basicMethodsStruct struct {
	httpHeaders          map[string]string
	httpHeadersForMobile bool
	isUserAgent          string
	isMobile             bool
	isTablet             bool
	customValues         []basicMethodsStructCustomValue
}

type basicMethodsStructCustomValue struct {
	name  string
	value bool
}

func BasicMethodsData() []basicMethodsStruct {
	return []basicMethodsStruct{
		basicMethodsStruct{
			httpHeaders: map[string]string{
				"SERVER_SOFTWARE":       "Apache/2.2.15 (Linux) Whatever/4.0 PHP/5.2.13",
				"REQUEST_METHOD":        "POST",
				"HTTP_HOST":             "home.ghita.org",
				"HTTP_X_REAL_IP":        "1.2.3.4",
				"HTTP_X_FORWARDED_FOR":  "1.2.3.5",
				"HTTP_CONNECTION":       "close",
				"HTTP_USER_AGENT":       "Mozilla/5.0 (iPhone; CPU iPhone OS 6_0_1 like Mac OS X) AppleWebKit/536.26 (KHTML, like Gecko) Version/6.0 Mobile/10A523 Safari/8536.25",
				"HTTP_ACCEPT":           "text/vnd.wap.wml, application/json, text/javascript, */*; q=0.01",
				"HTTP_ACCEPT_LANGUAGE":  "en-us,en;q=0.5",
				"HTTP_ACCEPT_ENCODING":  "gzip, deflate",
				"HTTP_X_REQUESTED_WITH": "XMLHttpRequest",
				"HTTP_REFERER":          "http://mobiledetect.net",
				"HTTP_PRAGMA":           "no-cache",
				"HTTP_CACHE_CONTROL":    "no-cache",
				"REMOTE_ADDR":           "11.22.33.44",
				"REQUEST_TIME":          "01-10-2012 07:57",
			},
			httpHeadersForMobile: false,
			isMobile:             true,
			isTablet:             false,
			customValues: []basicMethodsStructCustomValue{
				basicMethodsStructCustomValue{
					name:  "iphone",
					value: true,
				},
				basicMethodsStructCustomValue{
					name:  "ios",
					value: true,
				},
				basicMethodsStructCustomValue{
					name:  "whatever",
					value: false,
				},
			},
		},
		basicMethodsStruct{
			httpHeaders: map[string]string{
				"SERVER_SOFTWARE":       "Apache/2.2.15 (Linux) Whatever/4.0 PHP/5.2.13",
				"REQUEST_METHOD":        "POST",
				"HTTP_HOST":             "",
				"HTTP_X_REAL_IP":        "1.2.3.4",
				"HTTP_X_FORWARDED_FOR":  "1.2.3.5",
				"HTTP_CONNECTION":       "close",
				"HTTP_USER_AGENT":       "Mozilla/5.0",
				"HTTP_ACCEPT":           "application/json, text/javascript, */*; q=0.01",
				"HTTP_ACCEPT_LANGUAGE":  "en-us,en;q=0.5",
				"HTTP_ACCEPT_ENCODING":  "gzip, deflate",
				"HTTP_X_REQUESTED_WITH": "XMLHttpRequest",
				"HTTP_REFERER":          "",
				"HTTP_PRAGMA":           "no-cache",
				"HTTP_CACHE_CONTROL":    "no-cache",
				"REMOTE_ADDR":           "11.22.33.44",
				"REQUEST_TIME":          "01-10-2012 07:57",
			},
			httpHeadersForMobile: true,
			isMobile:             false,
			isTablet:             false,
			customValues: []basicMethodsStructCustomValue{
				basicMethodsStructCustomValue{
					name:  "iphone",
					value: false,
				},
				basicMethodsStructCustomValue{
					name:  "ios",
					value: false,
				},
				basicMethodsStructCustomValue{
					name:  "whatever",
					value: false,
				},
			},
		},
		basicMethodsStruct{
			httpHeaders: map[string]string{
				"SERVER_SOFTWARE":       "Apache/2.2.15 (Linux) Whatever/4.0 PHP/5.2.13",
				"REQUEST_METHOD":        "POST",
				"HTTP_HOST":             "",
				"HTTP_X_REAL_IP":        "1.2.3.4",
				"HTTP_X_FORWARDED_FOR":  "1.2.3.5",
				"HTTP_CONNECTION":       "close",
				"HTTP_USER_AGENT":       "Mozilla/5.0 (iPad; CPU OS 5_1_1 like Mac OS X; en-us) AppleWebKit/534.46.0 (KHTML, like Gecko) CriOS/21.0.1180.80 Mobile/9B206 Safari/7534.48.3 (6FF046A0-1BC4-4E7D-8A9D-6BF17622A123)",
				"HTTP_ACCEPT":           "application/json, text/javascript, */*; q=0.01",
				"HTTP_ACCEPT_LANGUAGE":  "en-us,en;q=0.5",
				"HTTP_ACCEPT_ENCODING":  "gzip, deflate",
				"HTTP_X_REQUESTED_WITH": "XMLHttpRequest",
				"HTTP_REFERER":          "",
				"HTTP_PRAGMA":           "no-cache",
				"HTTP_CACHE_CONTROL":    "no-cache",
				"REMOTE_ADDR":           "11.22.33.44",
				"REQUEST_TIME":          "01-10-2012 07:57",
			},
			httpHeadersForMobile: true,
			isMobile:             true,
			isTablet:             true,
			customValues: []basicMethodsStructCustomValue{
				basicMethodsStructCustomValue{
					name:  "iphone",
					value: false,
				},
				basicMethodsStructCustomValue{
					name:  "ios",
					value: true,
				},
				basicMethodsStructCustomValue{
					name:  "whatever",
					value: false,
				},
			},
		},
	}
}

func TestBasicMethods(t *testing.T) {
	detect := NewMobileDetect(httpRequest, nil)
	for _, data := range BasicMethodsData() {
		detect.SetHttpHeaders(data.httpHeaders)

		if 16 != len(detect.httpHeaders) {
			t.Error("Http headers were not set")
		}

		if data.httpHeadersForMobile == detect.CheckHttpHeadersForMobile() {
			t.Error("Http mobile headers check failed")
		}

		detect.SetUserAgent(data.httpHeaders["HTTP_USER_AGENT"])
		if data.isUserAgent == detect.userAgent {
			t.Error("User agent was not set")
		}

		if data.isMobile != detect.IsMobile() {
			t.Error("Mobile detection failed")
		}

		if data.isTablet != detect.IsTablet() {
			t.Error("Tablet detection failed")
		}

		for _, customValue := range data.customValues {
			if customValue.value != detect.Is(customValue.name) {
				t.Errorf("Is(%s) detetction failed", customValue.name)
			}
		}
	}
}

//special headers that give `quick` indication that a device is mobile
func QuickHeadersData() []map[string]string {
	headers := []map[string]string{
		map[string]string{`HTTP_ACCEPT`: `application/json; q=0.2, application/x-obml2d; q=0.8, image/gif; q=0.99, */*`},
		map[string]string{`HTTP_ACCEPT`: `text/*; q=0.1, application/vnd.rim.html`},
		map[string]string{`HTTP_ACCEPT`: `text/vnd.wap.wml`},
		map[string]string{`HTTP_ACCEPT`: `application/vnd.wap.xhtml+xml`},
		map[string]string{`HTTP_X_WAP_PROFILE`: `hello`},
		map[string]string{`HTTP_X_WAP_CLIENTID`: ``},
		map[string]string{`HTTP_WAP_CONNECTION`: ``},
		map[string]string{`HTTP_PROFILE`: ``},
		map[string]string{`HTTP_X_OPERAMINI_PHONE_UA`: ``},
		map[string]string{`HTTP_X_NOKIA_IPADDRESS`: ``},
		map[string]string{`HTTP_X_NOKIA_GATEWAY_ID`: ``},
		map[string]string{`HTTP_X_ORANGE_ID`: ``},
		map[string]string{`HTTP_X_VODAFONE_3GPDPCONTEXT`: ``},
		map[string]string{`HTTP_X_HUAWEI_USERID`: ``},
		map[string]string{`HTTP_UA_OS`: ``},
		map[string]string{`HTTP_X_MOBILE_GATEWAY`: ``},
		map[string]string{`HTTP_X_ATT_DEVICEID`: ``},
		map[string]string{`HTTP_UA_CPU`: `ARM`},
	}
	return headers
}

func TestQuickHeaders(t *testing.T) {
	detect := NewMobileDetect(httpRequest, nil)
	for _, httpHeaders := range QuickHeadersData() {
		detect.SetHttpHeaders(httpHeaders)
		if true != detect.CheckHttpHeadersForMobile() {
			t.Errorf("Headers %+v failed", httpHeaders)
		}
	}
}

func QuickNonMobileHeadersData() []map[string]string {
	headers := []map[string]string{
		map[string]string{`HTTP_UA_CPU`: `AMD64`},
		map[string]string{`HTTP_UA_CPU`: `X86`},
		map[string]string{`HTTP_ACCEPT`: `text/javascript, application/javascript, application/ecmascript, application/x-ecmascript, */*; q=0.01`},
		map[string]string{`HTTP_REQUEST_METHOD`: `DELETE`},
		map[string]string{`HTTP_VIA`: `1.1 ws-proxy.stuff.co.il C0A800FA`},
	}
	return headers
}

func TestNonMobileQuickHeaders(t *testing.T) {
	detect := NewMobileDetect(httpRequest, nil)
	for _, httpHeaders := range QuickNonMobileHeadersData() {
		detect.SetHttpHeaders(httpHeaders)
		if false != detect.CheckHttpHeadersForMobile() {
			t.Errorf("Headers %+v failed", httpHeaders)
		}
	}
}

type versionDataStruct struct {
	userAgent    string
	property     string
	strVersion   string
	floatVersion float64
}

func VersionData() []versionDataStruct {
	v := []versionDataStruct{
		versionDataStruct{
			userAgent:    `Mozilla/5.0 (Linux; Android 4.0.4; ARCHOS 80G9 Build/IMM76D) AppleWebKit/535.19 (KHTML, like Gecko) Chrome/18.0.1025.166  Safari/535.19`,
			property:     `Android`,
			strVersion:   `4.0.4`,
			floatVersion: 4.04,
		},
		versionDataStruct{
			userAgent:    `Mozilla/5.0 (Linux; Android 4.0.4; ARCHOS 80G9 Build/IMM76D) AppleWebKit/535.19 (KHTML, like Gecko) Chrome/18.0.1025.166  Safari/535.19`,
			property:     `Webkit`,
			strVersion:   `535.19`,
			floatVersion: 535.19,
		},
		versionDataStruct{
			userAgent:    `Mozilla/5.0 (Linux; Android 4.0.4; ARCHOS 80G9 Build/IMM76D) AppleWebKit/535.19 (KHTML, like Gecko) Chrome/18.0.1025.166  Safari/535.19`,
			property:     `Chrome`,
			strVersion:   `18.0.1025.166`,
			floatVersion: 18.01025166,
		},
		versionDataStruct{
			userAgent:    `Mozilla/5.0 (BlackBerry; U; BlackBerry 9700; en-US) AppleWebKit/534.8  (KHTML, like Gecko) Version/6.0.0.448 Mobile Safari/534.8`,
			property:     `BlackBerry`,
			strVersion:   `6.0.0.448`,
			floatVersion: 6.00448,
		},
		versionDataStruct{
			userAgent:    `Mozilla/5.0 (BlackBerry; U; BlackBerry 9700; en-US) AppleWebKit/534.8  (KHTML, like Gecko) Version/6.0.0.448 Mobile Safari/534.8`,
			property:     `Webkit`,
			strVersion:   `534.8`,
			floatVersion: 534.8,
		},
		versionDataStruct{
			userAgent:    `Mozilla/5.0 (BlackBerry; U; BlackBerry 9700; en-US) AppleWebKit/534.8  (KHTML, like Gecko) Version/6.0.0.448 Mobile Safari/534.8`,
			property:     `Webkit`,
			strVersion:   `534.8`,
			floatVersion: 534.8,
		},
		versionDataStruct{
			userAgent:    `Mozilla/5.0 (BlackBerry; U; BlackBerry 9700; en-US) AppleWebKit/534.8  (KHTML, like Gecko) Version/6.0.0.448 Mobile Safari/534.8`,
			property:     `Unknown property`,
			strVersion:   ``,
			floatVersion: 0.0,
		},
	}
	return v
}

//todo: check if this test is testing the code or testing that the data is correct
func TestVersionExtraction(t *testing.T) {
	detect := NewMobileDetect(httpRequest, nil)

	for _, data := range VersionData() {
		userAgent := data.userAgent
		strVersion := data.strVersion
		floatVersion := data.floatVersion
		property := data.property
		detect.SetUserAgent(userAgent)
		detectedVersion := detect.Version(property)
		if strVersion != detectedVersion {
			t.Errorf("String version %s is mismatched (detectedVersion %s, property %s)", strVersion, detectedVersion, property)
		}

		detectedVersionFloat := detect.VersionFloat(property)
		if floatVersion != detectedVersionFloat {
			t.Errorf("Float version %d is mismatched (detectedVersion %d, property %s)", floatVersion, detectedVersionFloat, property)
		}
	}
}
