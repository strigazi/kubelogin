package cmd

import "github.com/spf13/pflag"

type tlsOptions struct {
	CACertFilename string
	CACertData     string
	SkipTLSVerify  bool
}

func (o *tlsOptions) addFlags(f *pflag.FlagSet) {
	f.StringVar(&o.CACertFilename, "certificate-authority", "", "Path to a cert file for the certificate authority")
	f.StringVar(&o.CACertData, "certificate-authority-data", "", "Base64 encoded cert for the certificate authority")
	f.BoolVar(&o.SkipTLSVerify, "insecure-skip-tls-verify", false, "If set, the server's certificate will not be checked for validity. This will make your HTTPS connections insecure")
}
