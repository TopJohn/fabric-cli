/*
Copyright State Street Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package templates

// Config is a template string for SDK config yaml
const Config = `{{- $profile := .}}
client:
    organization: {{.Context.Organization}}

    cryptoconfig:
        path: {{.CryptoConfig}}

    credentialStore:
        path: "{{.CredentialStore}}/service-kvs"
        cryptoStore:
            path: "{{.CredentialStore}}/service-msp"

channels:
{{- range .Channels}}
    {{.ID}}:
    {{- if .Peers}}
        peers:
        {{- range .Peers}}
            {{.}}:
            {{- with (index $profile.Peers .)}}
                {{- range $k, $v := .ChannelOptions}}
                {{$k}}: {{$v}}
                {{- end}}
            {{- end}}
        {{- end}}
    {{- end}}
{{- end}}

organizations:
{{- range .Organizations}}
    {{.ID}}:
        mspid: {{.MSP.ID}}
        cryptoPath: {{.MSP.Store}}
        peers:
    {{- range .Peers}}
            - {{.}}
    {{- end}}
{{- end}}

orderers:
{{- range .Orderers}}
    {{.ID}}:
        url: {{.URL}}
        tlsCACerts:
            path: {{.TLS}}
{{- end}}

peers:
{{- range .Peers}}
    {{.ID}}:
        url: {{.URL}}
        eventUrl: {{.EventURL}}
    {{- if .GRPCOptions}}
        grpcOptions:
            {{- range $k, $v := .GRPCOptions}}
            {{$k}}: {{$v}}
            {{- end}}
    {{- end}}
        tlsCACerts:
            path: {{.TLS}}
{{- end}}
`