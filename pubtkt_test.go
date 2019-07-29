package pubtkt_test

import (
	. "github.com/orange-cloudfoundry/go-auth-pubtkt"

	"crypto/tls"
	"net/http"
	"net/url"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Pubtkt", func() {

	Context("RawToTicket", func() {
		It("Should give correct ticket when it's not encrypted", func() {
			ticketRaw := "uid=myuser;validuntil=1;tokens=token1,token2;sig=mysignature"
			auth, err := NewAuthPubTkt(AuthPubTktOptions{TKTAuthPublicKey: "fake", TKTAuthCookieName: "fake", TKTAuthHeader: []string{"fake"}})
			Expect(err).ToNot(HaveOccurred())

			ticket, err := auth.RawToTicket(ticketRaw)
			Expect(err).ToNot(HaveOccurred())
			Expect(*ticket).Should(Equal(Ticket{
				Uid:        "myuser",
				Validuntil: time.Unix(1, 0),
				Tokens:     []string{"token1", "token2"},
				RawData:    "uid=myuser;validuntil=1;tokens=token1,token2",
				Sig:        "mysignature",
			}))
		})
		It("Should give correct ticket when it's encrypted", func() {
			ticketRaw := "NgJVDZTchnQ3CpQWRhLHExefvSPkFyLIaCyvnNy+XB/BHu+ah1ojR2ZBrALb0fIqKKdIpnVQ9OBuJl8MXa/NZw=="
			passPhrase := "mysuperpassphrase"
			auth, _ := NewAuthPubTkt(AuthPubTktOptions{
				TKTAuthPublicKey:           "fake",
				TKTAuthCookieName:          "fake",
				TKTAuthHeader:              []string{"fake"},
				TKTCypherTicketsWithPasswd: passPhrase,
				TKTCypherTicketsMethod:     "cbc",
			})
			ticket, err := auth.RawToTicket(ticketRaw)
			Expect(err).ToNot(HaveOccurred())
			Expect(*ticket).Should(Equal(Ticket{
				Uid:        "myuser",
				Validuntil: time.Unix(1, 0),
				Tokens:     []string{"token1", "token2"},
				Sig:        "mysignature",
				RawData:    "uid=myuser;validuntil=1;tokens=token1,token2",
			}))
		})
	})
	Context("RequestToTicket", func() {
		It("Should give correct ticket from cookie when it's set", func() {
			ticketRaw := "uid=myuser;validuntil=1;tokens=token1,token2;sig=mysignature"
			auth, err := NewAuthPubTkt(AuthPubTktOptions{
				TKTAuthPublicKey:  "fake",
				TKTAuthHeader:     []string{"cookie"},
				TKTAuthCookieName: "pubtkt",
			})
			Expect(err).ToNot(HaveOccurred())

			req, _ := http.NewRequest("GET", "http://local.com", nil)
			req.AddCookie(&http.Cookie{Name: "pubtkt", Value: url.QueryEscape(ticketRaw)})

			ticket, err := auth.RequestToTicket(req)
			Expect(err).ToNot(HaveOccurred())
			Expect(*ticket).Should(Equal(Ticket{
				Uid:        "myuser",
				Validuntil: time.Unix(1, 0),
				Tokens:     []string{"token1", "token2"},
				Sig:        "mysignature",
				RawData:    "uid=myuser;validuntil=1;tokens=token1,token2",
			}))
		})
		It("Should give correct ticket from header if it's set when it's set", func() {
			ticketRaw := "uid=myuser;validuntil=1;tokens=token1,token2;sig=mysignature"
			auth, err := NewAuthPubTkt(AuthPubTktOptions{
				TKTAuthPublicKey:  "fake",
				TKTAuthCookieName: "fake",
				TKTAuthHeader:     []string{"x-authpubtkt"},
			})
			Expect(err).ToNot(HaveOccurred())

			req, _ := http.NewRequest("GET", "http://local.com", nil)
			req.Header.Set("x-authpubtkt", ticketRaw)

			ticket, err := auth.RequestToTicket(req)
			Expect(err).ToNot(HaveOccurred())
			Expect(*ticket).Should(Equal(Ticket{
				Uid:        "myuser",
				Validuntil: time.Unix(1, 0),
				Tokens:     []string{"token1", "token2"},
				Sig:        "mysignature",
				RawData:    "uid=myuser;validuntil=1;tokens=token1,token2",
			}))
		})
		It("Should give correct ticket from cookie by cascading if no header is set", func() {
			ticketRaw := "uid=myuser;validuntil=1;tokens=token1,token2;sig=mysignature"
			auth, err := NewAuthPubTkt(AuthPubTktOptions{
				TKTAuthPublicKey:  "fake",
				TKTAuthHeader:     []string{"x-authpubtkt", "cookie"},
				TKTAuthCookieName: "pubtkt",
			})
			Expect(err).ToNot(HaveOccurred())

			req, _ := http.NewRequest("GET", "http://local.com", nil)
			req.AddCookie(&http.Cookie{Name: "pubtkt", Value: url.QueryEscape(ticketRaw)})

			ticket, err := auth.RequestToTicket(req)
			Expect(err).ToNot(HaveOccurred())
			Expect(*ticket).Should(Equal(Ticket{
				Uid:        "myuser",
				Validuntil: time.Unix(1, 0),
				Tokens:     []string{"token1", "token2"},
				Sig:        "mysignature",
				RawData:    "uid=myuser;validuntil=1;tokens=token1,token2",
			}))
		})
		It("Should give an error if no header or cookie are set", func() {
			auth, err := NewAuthPubTkt(AuthPubTktOptions{
				TKTAuthPublicKey:  "fake",
				TKTAuthHeader:     []string{"x-authpubtkt", "cookie"},
				TKTAuthCookieName: "pubtkt",
			})
			Expect(err).ToNot(HaveOccurred())

			req, _ := http.NewRequest("GET", "http://local.com", nil)
			_, err = auth.RequestToTicket(req)
			Expect(err).Should(HaveOccurred())
			Expect(err).Should(MatchError(NewErrNoTicket().Error()))
		})
	})
	Context("SignedTicket", func() {
		var defaultTicket *Ticket
		pubKeyRsa := `-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAx5JJ32izx2rZF4L7cnfv
e4aMew22Lu5GwJ6YgOj1hXKwYjPk0l+qyvCVAPVSKEOEf7ehtL3h+/XEDV+DDrdC
ZSjSrzT+RRV5tnQ+x7nbibSwT/VewAU0yz+C5cVuX5QWWDQV8sY7sAvvnJ3HJkpc
HqQ0Jvk0+w212h+CnZpuakO3M7yfq3yv8u93mEyUwcmix9dXx/9Cuoe18KDjULrj
UVMRcaQeXlAFau9nzd14LYruU81ShWmHNzvgMWhT5jYiEBlfF6jHso5e3d1nlX0n
tU03Z0V1stilqjL9L9DFQZUnpyQJSGu3HS2pf+G0NFDQnETEryKuD0vPIa17C0yE
zQIDAQAB
-----END PUBLIC KEY-----`
		privKeyRsa := `-----BEGIN RSA PRIVATE KEY-----
MIIEogIBAAKCAQEAx5JJ32izx2rZF4L7cnfve4aMew22Lu5GwJ6YgOj1hXKwYjPk
0l+qyvCVAPVSKEOEf7ehtL3h+/XEDV+DDrdCZSjSrzT+RRV5tnQ+x7nbibSwT/Ve
wAU0yz+C5cVuX5QWWDQV8sY7sAvvnJ3HJkpcHqQ0Jvk0+w212h+CnZpuakO3M7yf
q3yv8u93mEyUwcmix9dXx/9Cuoe18KDjULrjUVMRcaQeXlAFau9nzd14LYruU81S
hWmHNzvgMWhT5jYiEBlfF6jHso5e3d1nlX0ntU03Z0V1stilqjL9L9DFQZUnpyQJ
SGu3HS2pf+G0NFDQnETEryKuD0vPIa17C0yEzQIDAQABAoIBADS+DJfma9y/+C+m
hh1yZPAYVvgOX593iWtOfq9S4dAqx0KsxER01AZeHoLqUpQhg9rdBPdnV1nnUbDX
FSPGet2RNFzpfGl1i30Uq6LNE6AJCK+ZATluJs2wMz/WNZ083crhuQs4KH4WKXUS
nZcY/895Re8m3UGcFftXaIWOUq/yTfeIKnpo4paEEjtlBv7OBbqN5R0oWzV8KcI8
rUYshmwdbRjZCgLgKsh9NmJu4XMvrPRDw1RpnilP34zk4dHKT3YtopD1w/OHNusX
LdE4vzl2Af5xuzpq5BKcJHlyFTwVHNh2ROAxgttxGnEKUst7UC4l8Aih//frGGVY
iSNsaGECgYEA4lbEFBgVIV2c9S8Jqy+eAdN2DRwLfl2mut3IV6s4A/548ItD+o/B
4Bbs7U3DkWHVdqOVxlMioKTlUnHRIEWbKruyU2qR0EHm13Ka+5A/LSrXFhhS6W1I
B7VmSNkuMR2i+Bl2imduybv/R0TB8oL1H2KsLEddPGlTnfDDRyCF4w8CgYEA4bmG
aMYU8r7YloVL3RlABKBfZ5i3OGSS80sJCHiNNcBKNOimVcQ8WuIj1xdgb7FS2Wsm
w0u+pQVNFNTta9Aya1IUnjPEAR30Whje9vPDWisULB8A1LHJZZclhRN+TAnHfFSE
+hzX/WNARRxtn41wkUb8TXWTauFIeclv14Jm6mMCgYBUBdPKhdoqRb9Dwc98DjRB
B0g2u6eEpK2Fh1mdOgGUcf5hhcGZnY5iawBPY+Pq73+4CChMtIYRTsWW4ou4yD4/
Eei0UMaKojxY1MG2C6l4UgQAvszgvOHVgsDS4FTmhNL/+SvpW1Zka3br9RlutrDa
kt3JnkB9cbhJ1JwZpngtfwKBgB2rKvV+g6ZcW099ebk35mg1RBTzx7FL1tzPkgJh
Bf54pchKhAJl5qDzdE6DLhPANOmK4e9td0NtbPI+Am7XUKk2gqtAlpSnAUUOW9AN
vjGN0/rwoEVUkOofcbkKFOGdr3CyE0BkpicR8pKi5+2+w3r035i2yYeU0NfhGCg1
ZZ2NAoGAGnc90uOcUt7O2kPNk8dHw+RAo8XyaZaOHHskCDWLRKpzU6IbY1DmTKj8
LM7yfzgsimLlqoMZLm66j82/lNqKSiMlnTPN8sP2LoKZ0vKNc3iulT8g8BuFcHT/
xFn9Fdsp/a/Y4I0JEEA76FAarnRZxpUhLC8pKpF6CjUetnUwp7g=
-----END RSA PRIVATE KEY-----`

		pubKeyDsa := `-----BEGIN PUBLIC KEY-----
MIIDODCCAisGByqGSM44BAEwggIeAoIBAQDsgZumckiMIe4UNkxfHkg9OGiVJaDX
76g8oodG5dCwqxb5JN9QNeIj8mpyOPsVg0uq6v16MaCGkAwVIlIvCjBULMCzLdUP
N3Rcwf5rUWyrYu+nfs4XXEwFHMqOPl+GgkK3b6Z3jrf/FwjCuDJDqm6wvBVo1o5g
GIkfYZAn+pX8tpT8xeHXM5tCjeKEHz0Vwu6JahT03LmetqFCSoNJJi07Fa9NWyqJ
byJHi3lcIIJPimeCZsC2NC+rXzoJ292tfa2OgsON2btQLD/6daa7jdukpzL84FYA
qhqMJfaEEJgdoHlUrwxCEu6cCLUqWD2YaoRG7m69zQG4L5/vzCAH/AkZAhUAmjb+
JTswjkZ6og46tbkCaQ7JWFMCgf9iJkAD3D5gannM4q2kHaoPxmxPcMe93CstJ8Na
I4A69rTzYxJmpWz83HxCIh0JVWFtOajFguuUV3mTAmqIw1O3MFXLRiCWLSQLEFS4
oQ7jrFfuhKP1XC9gcTvyKLFhSyt2iU6j38XkZME1sH8McMEXVO+KfufsauZqYJni
1mNZ7uilWwkSEXIFfj91zdF1ELxwprsrwFFbClhTsLSgopx9/na0fOmi9pDyoJZV
ymo/dnF3/PF7guhF3Owj9JShzgKBGpEe3BB0bbhfvKHYE5QQOHF+qSGWDeXcEul1
lWqa8lAJyMnp5FCyIFYH1qNw5HBoayf7HlGaUd8oCk9q+boDggEFAAKCAQB8chaK
eH8/AMOUiOMUZulHGqF4MGLQDxDP854BDZSPcb10lUuXfRJvXdN1gM4T+E2oC3Jv
AeNepm7PqTSr23Hy+GQ7ey+LTO1Z2aSHjOf2rMOEJrqwj43zDxuRj5oMYstD3mR6
elGshc8N0HA7qIMoRdCkZfslh6vbd9SGQqIT3qkIh7jOcp8V3gi2fAkz5ZXmkODV
QOZkGj1O0vPk3+gVrVXKnQj5WF7wtT+iUEBs6L65hmXYMown4wzolJq+XT3lliAq
pzY+pD9CXCD6qkeXWXSxLfhlVs+42bPzebiS5lZr6yZ5knR0QM10oT+0ApODWmbd
StOB7bD9meH5/rOy
-----END PUBLIC KEY-----`
		privKeyDsa := `-----BEGIN DSA PRIVATE KEY-----
MIIDOwIBAAKCAQEA7IGbpnJIjCHuFDZMXx5IPTholSWg1++oPKKHRuXQsKsW+STf
UDXiI/Jqcjj7FYNLqur9ejGghpAMFSJSLwowVCzAsy3VDzd0XMH+a1Fsq2Lvp37O
F1xMBRzKjj5fhoJCt2+md463/xcIwrgyQ6pusLwVaNaOYBiJH2GQJ/qV/LaU/MXh
1zObQo3ihB89FcLuiWoU9Ny5nrahQkqDSSYtOxWvTVsqiW8iR4t5XCCCT4pngmbA
tjQvq186CdvdrX2tjoLDjdm7UCw/+nWmu43bpKcy/OBWAKoajCX2hBCYHaB5VK8M
QhLunAi1Klg9mGqERu5uvc0BuC+f78wgB/wJGQIVAJo2/iU7MI5GeqIOOrW5AmkO
yVhTAoH/YiZAA9w+YGp5zOKtpB2qD8ZsT3DHvdwrLSfDWiOAOva082MSZqVs/Nx8
QiIdCVVhbTmoxYLrlFd5kwJqiMNTtzBVy0Ygli0kCxBUuKEO46xX7oSj9VwvYHE7
8iixYUsrdolOo9/F5GTBNbB/DHDBF1Tvin7n7GrmamCZ4tZjWe7opVsJEhFyBX4/
dc3RdRC8cKa7K8BRWwpYU7C0oKKcff52tHzpovaQ8qCWVcpqP3Zxd/zxe4LoRdzs
I/SUoc4CgRqRHtwQdG24X7yh2BOUEDhxfqkhlg3l3BLpdZVqmvJQCcjJ6eRQsiBW
B9ajcORwaGsn+x5RmlHfKApPavm6AoIBAHxyFop4fz8Aw5SI4xRm6UcaoXgwYtAP
EM/zngENlI9xvXSVS5d9Em9d03WAzhP4TagLcm8B416mbs+pNKvbcfL4ZDt7L4tM
7VnZpIeM5/asw4QmurCPjfMPG5GPmgxiy0PeZHp6UayFzw3QcDuogyhF0KRl+yWH
q9t31IZCohPeqQiHuM5ynxXeCLZ8CTPlleaQ4NVA5mQaPU7S8+Tf6BWtVcqdCPlY
XvC1P6JQQGzovrmGZdgyjCfjDOiUmr5dPeWWICqnNj6kP0JcIPqqR5dZdLEt+GVW
z7jZs/N5uJLmVmvrJnmSdHRAzXShP7QCk4NaZt1K04HtsP2Z4fn+s7ICFAqAGLT+
aUu2Z8JldUWSq2sphCGj
-----END DSA PRIVATE KEY-----`
		BeforeEach(func() {
			TimeNowFunc = func() time.Time {
				return time.Unix(0, 0)
			}
			defaultTicket = &Ticket{
				Uid:        "myuser",
				Cip:        "127.0.0.1",
				Validuntil: time.Unix(1, 0),
				Tokens:     []string{"token1", "token2"},
			}
		})
		Context("SignTicket", func() {
			It("should sign ticket with private key when using rsa", func() {
				auth, err := NewAuthPubTkt(AuthPubTktOptions{
					TKTAuthPublicKey:  pubKeyRsa,
					TKTAuthPrivateKey: privKeyRsa,
					TKTAuthCookieName: "fake",
					TKTAuthHeader:     []string{"fake"},
				})
				Expect(err).ToNot(HaveOccurred())

				err = auth.SignTicket(defaultTicket)
				Expect(err).ToNot(HaveOccurred())
				Expect(defaultTicket.Sig).NotTo(BeEmpty())

				err = auth.VerifyTicket(defaultTicket, "127.0.0.1")
				Expect(err).ShouldNot(HaveOccurred())
			})
			It("should sign ticket with private key when using dsa", func() {
				auth, err := NewAuthPubTkt(AuthPubTktOptions{
					TKTAuthPublicKey:  pubKeyDsa,
					TKTAuthPrivateKey: privKeyDsa,
					TKTAuthCookieName: "fake",
					TKTAuthDigest:     "DSS1",
					TKTAuthHeader:     []string{"fake"},
				})
				Expect(err).ToNot(HaveOccurred())

				err = auth.SignTicket(defaultTicket)
				Expect(err).ToNot(HaveOccurred())
				Expect(defaultTicket.Sig).NotTo(BeEmpty())

				err = auth.VerifyTicket(defaultTicket, "127.0.0.1")
				Expect(err).ShouldNot(HaveOccurred())
			})
		})
		Context("TicketToRaw", func() {
			It("should create plain ticket string with no cipher", func() {
				auth, err := NewAuthPubTkt(AuthPubTktOptions{
					TKTAuthPublicKey:  pubKeyRsa,
					TKTAuthPrivateKey: privKeyRsa,
					TKTAuthCookieName: "fake",
					TKTAuthHeader:     []string{"fake"},
				})
				Expect(err).ToNot(HaveOccurred())

				raw, err := auth.TicketToRaw(defaultTicket)
				Expect(err).ToNot(HaveOccurred())

				tkt, err := auth.RawToTicket(raw)
				Expect(err).ToNot(HaveOccurred())
				Expect(tkt.Sig).NotTo(BeEmpty())
				Expect(defaultTicket.DataString()).To(Equal(tkt.DataString()))

			})
			It("should create plain ticket string with cipher", func() {
				auth, err := NewAuthPubTkt(AuthPubTktOptions{
					TKTAuthPublicKey:           pubKeyRsa,
					TKTAuthPrivateKey:          privKeyRsa,
					TKTAuthCookieName:          "fake",
					TKTCypherTicketsWithPasswd: "mypassphrase",
					TKTAuthHeader:              []string{"fake"},
				})
				Expect(err).ToNot(HaveOccurred())

				raw, err := auth.TicketToRaw(defaultTicket)
				Expect(err).ToNot(HaveOccurred())

				Expect(raw).ShouldNot(ContainSubstring(defaultTicket.DataString()))

				tkt, err := auth.RawToTicket(raw)
				Expect(err).ToNot(HaveOccurred())
				Expect(tkt.Sig).NotTo(BeEmpty())
				Expect(defaultTicket.DataString()).To(Equal(tkt.DataString()))
			})
		})
		Context("TicketInRequest", func() {
			It("should put ticket inside cookie when cookie required", func() {
				auth, err := NewAuthPubTkt(AuthPubTktOptions{
					TKTAuthPublicKey:  pubKeyRsa,
					TKTAuthPrivateKey: privKeyRsa,
					TKTAuthCookieName: "fake",
					TKTAuthHeader:     []string{"cookie"},
				})
				Expect(err).ToNot(HaveOccurred())

				req, _ := http.NewRequest("GET", "http://local.com", nil)
				req.TLS = &tls.ConnectionState{}
				req.RemoteAddr = "127.0.0.1:52332"

				err = auth.TicketInRequest(req, defaultTicket)
				Expect(err).ToNot(HaveOccurred())

				cookie, err := req.Cookie("fake")
				Expect(err).ToNot(HaveOccurred())

				tktRaw, _ := url.QueryUnescape(cookie.Value)

				tkt, err := auth.RawToTicket(tktRaw)
				Expect(err).ToNot(HaveOccurred())
				Expect(defaultTicket.DataString()).To(Equal(tkt.DataString()))
			})
			It("should put ticket inside header when header required", func() {
				auth, err := NewAuthPubTkt(AuthPubTktOptions{
					TKTAuthPublicKey:  pubKeyRsa,
					TKTAuthPrivateKey: privKeyRsa,
					TKTAuthHeader:     []string{"X-Pub-Tkt"},
				})
				Expect(err).ToNot(HaveOccurred())

				req, _ := http.NewRequest("GET", "http://local.com", nil)
				req.TLS = &tls.ConnectionState{}
				req.RemoteAddr = "127.0.0.1:52332"

				err = auth.TicketInRequest(req, defaultTicket)
				Expect(err).ToNot(HaveOccurred())

				tktRaw, _ := url.QueryUnescape(req.Header.Get("X-Pub-Tkt"))
				Expect(err).ToNot(HaveOccurred())

				tkt, err := auth.RawToTicket(tktRaw)
				Expect(err).ToNot(HaveOccurred())
				Expect(defaultTicket.DataString()).To(Equal(tkt.DataString()))
			})
		})
	})
	Context("Verify", func() {
		var defaultTicket *Ticket
		// with ticket data uid=myuser;cip=127.0.0.1;validuntil=1;tokens=token1,token2
		pubKeyRsa := `-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAx5JJ32izx2rZF4L7cnfv
e4aMew22Lu5GwJ6YgOj1hXKwYjPk0l+qyvCVAPVSKEOEf7ehtL3h+/XEDV+DDrdC
ZSjSrzT+RRV5tnQ+x7nbibSwT/VewAU0yz+C5cVuX5QWWDQV8sY7sAvvnJ3HJkpc
HqQ0Jvk0+w212h+CnZpuakO3M7yfq3yv8u93mEyUwcmix9dXx/9Cuoe18KDjULrj
UVMRcaQeXlAFau9nzd14LYruU81ShWmHNzvgMWhT5jYiEBlfF6jHso5e3d1nlX0n
tU03Z0V1stilqjL9L9DFQZUnpyQJSGu3HS2pf+G0NFDQnETEryKuD0vPIa17C0yE
zQIDAQAB
-----END PUBLIC KEY-----`

		sha1Sig := "CLB5SmRpGGiYwUM76MXfVS+h9cp9nq3G6xQ13/XrvTOXon2lR903Wuixz/zEt2ljZm9gSosfZmpa12k3csEOKqwGvZCDHJCfb/EibY/xDXJjgGv89XMtIwYSmDjJ1GJOuPG0YERZALIyfHmMLJZOXq6QalzQ/PRRNeZn93k+8KeetsO33W785vnSqDMkwL9JIJHHcxSd4pJLPsSUCQVPXJN5mWZWI56J0KHZht08klKc2EFx39jd4QImjWEu188HvQ5/NO4L6COjS/J29JrAGWN3IRvu7gq7Krzcm8wdkL1Hf4r2vsS1unpT6E0MfaIqLZOa9FPsvIp3EP4M2ugwLg=="
		dss1Sig := "MC0CFQCUijdPIW3yHoPtru/my77rdVrJVAIUfdtGuRsTd1B2V2LCET/LILiX/II="
		pubKeyDsa := `-----BEGIN PUBLIC KEY-----
MIIDODCCAisGByqGSM44BAEwggIeAoIBAQDsgZumckiMIe4UNkxfHkg9OGiVJaDX
76g8oodG5dCwqxb5JN9QNeIj8mpyOPsVg0uq6v16MaCGkAwVIlIvCjBULMCzLdUP
N3Rcwf5rUWyrYu+nfs4XXEwFHMqOPl+GgkK3b6Z3jrf/FwjCuDJDqm6wvBVo1o5g
GIkfYZAn+pX8tpT8xeHXM5tCjeKEHz0Vwu6JahT03LmetqFCSoNJJi07Fa9NWyqJ
byJHi3lcIIJPimeCZsC2NC+rXzoJ292tfa2OgsON2btQLD/6daa7jdukpzL84FYA
qhqMJfaEEJgdoHlUrwxCEu6cCLUqWD2YaoRG7m69zQG4L5/vzCAH/AkZAhUAmjb+
JTswjkZ6og46tbkCaQ7JWFMCgf9iJkAD3D5gannM4q2kHaoPxmxPcMe93CstJ8Na
I4A69rTzYxJmpWz83HxCIh0JVWFtOajFguuUV3mTAmqIw1O3MFXLRiCWLSQLEFS4
oQ7jrFfuhKP1XC9gcTvyKLFhSyt2iU6j38XkZME1sH8McMEXVO+KfufsauZqYJni
1mNZ7uilWwkSEXIFfj91zdF1ELxwprsrwFFbClhTsLSgopx9/na0fOmi9pDyoJZV
ymo/dnF3/PF7guhF3Owj9JShzgKBGpEe3BB0bbhfvKHYE5QQOHF+qSGWDeXcEul1
lWqa8lAJyMnp5FCyIFYH1qNw5HBoayf7HlGaUd8oCk9q+boDggEFAAKCAQB8chaK
eH8/AMOUiOMUZulHGqF4MGLQDxDP854BDZSPcb10lUuXfRJvXdN1gM4T+E2oC3Jv
AeNepm7PqTSr23Hy+GQ7ey+LTO1Z2aSHjOf2rMOEJrqwj43zDxuRj5oMYstD3mR6
elGshc8N0HA7qIMoRdCkZfslh6vbd9SGQqIT3qkIh7jOcp8V3gi2fAkz5ZXmkODV
QOZkGj1O0vPk3+gVrVXKnQj5WF7wtT+iUEBs6L65hmXYMown4wzolJq+XT3lliAq
pzY+pD9CXCD6qkeXWXSxLfhlVs+42bPzebiS5lZr6yZ5knR0QM10oT+0ApODWmbd
StOB7bD9meH5/rOy
-----END PUBLIC KEY-----`
		BeforeEach(func() {
			TimeNowFunc = func() time.Time {
				return time.Unix(0, 0)
			}
			defaultTicket = &Ticket{
				Uid:        "myuser",
				Cip:        "127.0.0.1",
				Validuntil: time.Unix(1, 0),
				Tokens:     []string{"token1", "token2"},
			}
		})
		Context("VerifyTicket", func() {
			It("Should complain about signature when signature isn't valid", func() {
				auth, err := NewAuthPubTkt(AuthPubTktOptions{
					TKTAuthPublicKey:  pubKeyRsa,
					TKTAuthCookieName: "fake",
					TKTAuthHeader:     []string{"fake"},
				})
				Expect(err).ToNot(HaveOccurred())
				defaultTicket.Sig = "mysignature"

				err = auth.VerifyTicket(defaultTicket, "")
				Expect(err).Should(HaveOccurred())
				_, isErrSigNotValid := err.(ErrSigNotValid)
				Expect(isErrSigNotValid).Should(BeTrue())
			})
			It("Should use either rsa or dsa verify when TKTAuthDigest is not set", func() {
				authRsa, err := NewAuthPubTkt(AuthPubTktOptions{
					TKTAuthPublicKey:  pubKeyRsa,
					TKTCheckIpEnabled: false,
					TKTAuthToken:      []string{"token1"},
					TKTAuthCookieName: "fake",
					TKTAuthHeader:     []string{"fake"},
				})
				Expect(err).ToNot(HaveOccurred())
				authDsa, err := NewAuthPubTkt(AuthPubTktOptions{
					TKTAuthPublicKey:  pubKeyDsa,
					TKTCheckIpEnabled: false,
					TKTAuthToken:      []string{"token1"},
					TKTAuthCookieName: "fake",
					TKTAuthHeader:     []string{"fake"},
				})

				Expect(err).ToNot(HaveOccurred())
				defaultTicket.Sig = sha1Sig
				err = authRsa.VerifyTicket(defaultTicket, "")
				Expect(err).ShouldNot(HaveOccurred())

				defaultTicket.Sig = dss1Sig
				err = authDsa.VerifyTicket(defaultTicket, "")
				Expect(err).ShouldNot(HaveOccurred())
			})
			It("should complain about token not found when user doesn't have the requested token", func() {
				auth, err := NewAuthPubTkt(AuthPubTktOptions{
					TKTAuthPublicKey:  pubKeyRsa,
					TKTCheckIpEnabled: false,
					TKTAuthToken:      []string{"requiredToken"},
					TKTAuthCookieName: "fake",
					TKTAuthHeader:     []string{"fake"},
				})
				Expect(err).ToNot(HaveOccurred())
				defaultTicket.Sig = sha1Sig

				err = auth.VerifyTicket(defaultTicket, "")
				Expect(err).Should(HaveOccurred())
				_, isType := err.(ErrNoValidToken)
				Expect(isType).Should(BeTrue())
			})
			It("should complain about ip if user doesn't have the ip inside the ticket", func() {
				auth, err := NewAuthPubTkt(AuthPubTktOptions{
					TKTAuthPublicKey:  pubKeyRsa,
					TKTCheckIpEnabled: true,
					TKTAuthToken:      []string{"token1"},
					TKTAuthCookieName: "fake",
					TKTAuthHeader:     []string{"fake"},
				})
				Expect(err).ToNot(HaveOccurred())
				defaultTicket.Sig = sha1Sig

				err = auth.VerifyTicket(defaultTicket, "fakeIP")
				Expect(err).Should(HaveOccurred())
				_, isType := err.(ErrWrongIp)
				Expect(isType).Should(BeTrue())
			})
			It("should complain about expiration if current time is higher than expiration time", func() {
				auth, err := NewAuthPubTkt(AuthPubTktOptions{
					TKTAuthPublicKey:  pubKeyRsa,
					TKTCheckIpEnabled: false,
					TKTAuthToken:      []string{"token1"},
					TKTAuthCookieName: "fake",
					TKTAuthHeader:     []string{"fake"},
				})
				Expect(err).ToNot(HaveOccurred())
				defaultTicket.Sig = sha1Sig
				TimeNowFunc = func() time.Time {
					return time.Unix(2, 0)
				}
				err = auth.VerifyTicket(defaultTicket, "")
				Expect(err).Should(HaveOccurred())
				_, isType := err.(ErrValidationExpired)
				Expect(isType).Should(BeTrue())
			})
			It("should return no error when all is valid", func() {
				auth, err := NewAuthPubTkt(AuthPubTktOptions{
					TKTAuthPublicKey:  pubKeyRsa,
					TKTCheckIpEnabled: true,
					TKTAuthToken:      []string{"token1"},
					TKTAuthCookieName: "fake",
					TKTAuthHeader:     []string{"fake"},
				})
				Expect(err).ToNot(HaveOccurred())

				defaultTicket.Sig = sha1Sig
				err = auth.VerifyTicket(defaultTicket, "127.0.0.1")
				Expect(err).ShouldNot(HaveOccurred())
			})
		})
		Context("VerifyTicket", func() {
			It("should return no error and a ticket when all is valid", func() {
				auth, err := NewAuthPubTkt(AuthPubTktOptions{
					TKTAuthPublicKey:  pubKeyRsa,
					TKTCheckIpEnabled: true,
					TKTAuthRequireSSL: true,
					TKTAuthCookieName: "pubtkt",
					TKTAuthToken:      []string{"token1"},
					TKTAuthHeader:     []string{"cookie"},
				})
				Expect(err).ToNot(HaveOccurred())

				defaultTicket.Sig = sha1Sig
				req, _ := http.NewRequest("GET", "http://local.com", nil)
				req.TLS = &tls.ConnectionState{}
				req.RemoteAddr = "127.0.0.1:52332"
				req.AddCookie(&http.Cookie{Name: "pubtkt", Value: url.QueryEscape(defaultTicket.String())})

				ticket, err := auth.VerifyFromRequest(req)
				Expect(err).ShouldNot(HaveOccurred())

				defaultTicket.RawData = "uid=myuser;cip=127.0.0.1;validuntil=1;tokens=token1,token2"
				Expect(ticket).Should(Equal(defaultTicket))

			})
			It("should return no error and a ticket when all is valid and ip in x-forwarded-for is correct", func() {
				auth, err := NewAuthPubTkt(AuthPubTktOptions{
					TKTAuthPublicKey:     pubKeyRsa,
					TKTCheckIpEnabled:    true,
					TKTCheckXForwardedIp: true,
					TKTAuthRequireSSL:    true,
					TKTAuthCookieName:    "pubtkt",
					TKTAuthToken:         []string{"token1"},
					TKTAuthHeader:        []string{"cookie"},
				})
				Expect(err).ToNot(HaveOccurred())
				defaultTicket.Sig = sha1Sig

				req, _ := http.NewRequest("GET", "http://local.com", nil)
				req.TLS = &tls.ConnectionState{}
				req.Header.Set("X-Forwarded-For", "127.0.0.1:6060")
				req.AddCookie(&http.Cookie{Name: "pubtkt", Value: url.QueryEscape(defaultTicket.String())})

				ticket, err := auth.VerifyFromRequest(req)
				Expect(err).ShouldNot(HaveOccurred())

				defaultTicket.RawData = "uid=myuser;cip=127.0.0.1;validuntil=1;tokens=token1,token2"
				Expect(ticket).Should(Equal(defaultTicket))

			})
			It("should complain if ssl is required and request is not tls", func() {
				auth, _ := NewAuthPubTkt(AuthPubTktOptions{
					TKTAuthPublicKey:  pubKeyRsa,
					TKTCheckIpEnabled: false,
					TKTAuthRequireSSL: true,
					TKTAuthToken:      []string{"token1"},
					TKTAuthCookieName: "fake",
					TKTAuthHeader:     []string{"fake"},
				})
				defaultTicket.Sig = sha1Sig
				req, _ := http.NewRequest("GET", "http://local.com", nil)
				_, err := auth.VerifyFromRequest(req)
				Expect(err).Should(HaveOccurred())
				_, isType := err.(ErrNoSSl)
				Expect(isType).Should(BeTrue())
			})
			It("should complain about ip if request remote address doesn't have the ip inside the ticket ", func() {
				auth, err := NewAuthPubTkt(AuthPubTktOptions{
					TKTAuthPublicKey:  pubKeyRsa,
					TKTCheckIpEnabled: true,
					TKTAuthRequireSSL: false,
					TKTAuthToken:      []string{"token1"},
					TKTAuthHeader:     []string{"cookie"},
					TKTAuthCookieName: "pubtkt",
				})
				Expect(err).ToNot(HaveOccurred())
				defaultTicket.Sig = sha1Sig
				req, _ := http.NewRequest("GET", "http://local.com", nil)
				req.RemoteAddr = "fakeip:52332"
				req.AddCookie(&http.Cookie{Name: "pubtkt", Value: url.QueryEscape(defaultTicket.String())})

				_, err = auth.VerifyFromRequest(req)
				Expect(err).Should(HaveOccurred())
				_, isType := err.(ErrWrongIp)
				Expect(isType).Should(BeTrue())
			})
			It("should complain about ip if request header x-forwarded-for doesn't have the ip inside the ticket ", func() {
				auth, err := NewAuthPubTkt(AuthPubTktOptions{
					TKTAuthPublicKey:     pubKeyRsa,
					TKTCheckIpEnabled:    true,
					TKTCheckXForwardedIp: true,
					TKTAuthRequireSSL:    false,
					TKTAuthToken:         []string{"token1"},
					TKTAuthHeader:        []string{"cookie"},
					TKTAuthCookieName:    "pubtkt",
				})
				Expect(err).ToNot(HaveOccurred())
				defaultTicket.Sig = sha1Sig
				req, _ := http.NewRequest("GET", "http://local.com", nil)
				req.Header.Set("X-Forwarded-For", "fakeip:52332")
				req.AddCookie(&http.Cookie{Name: "pubtkt", Value: url.QueryEscape(defaultTicket.String())})

				_, err = auth.VerifyFromRequest(req)
				Expect(err).Should(HaveOccurred())
				_, isType := err.(ErrWrongIp)
				Expect(isType).Should(BeTrue())
			})
		})
	})
})
