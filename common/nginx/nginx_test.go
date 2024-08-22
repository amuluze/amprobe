package nginx

import "testing"

func TestPopulate(t *testing.T) {

}

func TestWriteConfig(t *testing.T) {
	siteGroup := SiteGroup{
		Sites: []*Site{
			{
				Name:        "test",
				ServerNames: []string{"test.com"},
				UrlPaths: []*UrlPath{
					{
						Op:      "pro",
						UrlPath: "/",
					},
					{
						Op:      "pro",
						UrlPath: "/test",
					},
				},
				ProxyConfig: &ProxyConfig{
					Upstreams: []*Upstream{
						{
							Address: "127.0.0.1",
							Port:    8080,
							Weight:  1,
						},
					},
					Keepalive:        65,
					KeepaliveTimeout: 65,
					Schema:           "http",
				},
				ListenConfigs: []*ListenConfig{
					{
						Address: "",
						Port:    8000,
					},
				},
			},
		},
	}
	err := WriteConfig("/Users/amu/Desktop/github/nginx", siteGroup)
	if err != nil {
		t.Errorf("write config error = %v", err)
		return
	}
}
