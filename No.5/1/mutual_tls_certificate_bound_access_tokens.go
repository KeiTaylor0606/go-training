package main

import (
	"crypto/sha256"
	"crypto/tls"
	"encoding/base64"
)

/*
フィンガープリントの作成
*/
func getThumbprintFromTLSState(state *tls.ConnectionState) (string, error) {
	var ErrMutualTLSConnection error
	if state == nil {
		return "", ErrMutualTLSConnection
	}
	PeerCertificates := state.PeerCertificates
	if PeerCertificates == nil {
		return "", ErrMutualTLSConnection
	}

	if len(PeerCertificates) <= 0 {
		return "", ErrMutualTLSConnection
	}

	cert := PeerCertificates[0]
	sum := sha256.Sum256(cert.Raw)
	return base64.RawURLEncoding.EncodeToString(sum[:]), nil
}

/*
フィンガープリントの作成（gPRCの場合）
*/
// func getCSFromContext(ctx context.Context) (*tls.ConnectionState, error) {
// 	peer, ok := peer.FromContext(ctx)
// 	if !ok {
// 		return nil, errors.New("failed to get peer")
// 	}

// 	if peer.AuthInfo == nil {
// 		return nil, errors.New("connection should be used TLS")
// 	}

// 	if peer.AuthInfo.AuthType() != "tls" {
// 		return nil, errors.New("connection should be used TLS")
// 	}

// 	tlsInfo, ok := peer.AuthInfo.(credentials.TLSInfo)
// 	if !ok {
// 		return nil, errors.New("connection should be used TLS")
// 	}
// 	return &tlsInfo.State, nil
// }

func MutualTlsCertificateBoundAccessTokens() {
	/*
		GoにおけるMutual TLS Certificate-Bound Access Tokensの実装
	*/
}
