package main

import (
    "testing"
)

const (
    WITH_OUTSTANDING_UPGRADE bool = true
    WITHOUT_OUTSTANDING_UPGRADE bool = false
)

func GenerateGetInfoString(isUpgradeOutstanding bool) string {
    if isUpgradeOutstanding {
        return `
        <?xml version="1.0"?>
        <s:Envelope xmlns:s="http://schemas.xmlsoap.org/soap/envelope/" s:encodingStyle="http://schemas.xmlsoap.org/soap/encoding/">
            <s:Body>
                <u:GetInfoResponse xmlns:u="urn:dslforum-org:service:UserInterface:1">
                    <NewUpgradeAvailable>1</NewUpgradeAvailable>
                    <NewPasswordRequired>0</NewPasswordRequired>
                    <NewPasswordUserSelectable>1</NewPasswordUserSelectable>
                    <NewWarrantyDate>0001-01-01T00:00:00</NewWarrantyDate>
                    <NewX_AVM-DE_Version>113.07.01</NewX_AVM-DE_Version>
                    <NewX_AVM-DE_DownloadURL>http://download.avm.de/fritzbox/fritzbox-7490/deutschland/fritz.os/FRITZ.Box_7490.113.07.01.image</NewX_AVM-DE_DownloadURL>
                    <NewX_AVM-DE_InfoURL>http://download.avm.de/fritzbox/fritzbox-7490/deutschland/fritz.os/info_de.txt</NewX_AVM-DE_InfoURL>
                    <NewX_AVM-DE_UpdateState>Error</NewX_AVM-DE_UpdateState>
                    <NewX_AVM-DE_LaborVersion></NewX_AVM-DE_LaborVersion>
                </u:GetInfoResponse>
            </s:Body>
        </s:Envelope>
        `
    } else {
        return `
        <?xml version="1.0"?>
        <s:Envelope xmlns:s="http://schemas.xmlsoap.org/soap/envelope/" s:encodingStyle="http://schemas.xmlsoap.org/soap/encoding/">
            <s:Body>
                <u:GetInfoResponse xmlns:u="urn:dslforum-org:service:UserInterface:1">
                    <NewUpgradeAvailable>0</NewUpgradeAvailable>
                    <NewPasswordRequired>0</NewPasswordRequired>
                    <NewPasswordUserSelectable>1</NewPasswordUserSelectable>
                    <NewWarrantyDate>0001-01-01T00:00:00</NewWarrantyDate>
                    <NewX_AVM-DE_Version></NewX_AVM-DE_Version>
                    <NewX_AVM-DE_DownloadURL></NewX_AVM-DE_DownloadURL>
                    <NewX_AVM-DE_InfoURL></NewX_AVM-DE_InfoURL>
                    <NewX_AVM-DE_UpdateState>Stopped</NewX_AVM-DE_UpdateState>
                    <NewX_AVM-DE_LaborVersion></NewX_AVM-DE_LaborVersion>
                </u:GetInfoResponse>
            </s:Body>
        </s:Envelope>
        `
    }
}

func TestNoUpgradeAvailable(t *testing.T) {
    if IsUpgradeAvailable(GenerateGetInfoString(WITHOUT_OUTSTANDING_UPGRADE)) {
        t.Error("Detected outstanding upgrade when no upgrade is outstanding")
    }
}

func TestUpgradeAvailable(t *testing.T) {
    if IsUpgradeAvailable(GenerateGetInfoString(WITH_OUTSTANDING_UPGRADE)) == false {
        t.Error("Did not detect outstanding upgrade")
    }
}
