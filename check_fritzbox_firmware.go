package main

import (
    "flag"
    "fmt"
    "os"
    "net/http"
    "strings"
    "io/ioutil"
)

const (
    EXIT_OK int = 0
    EXIT_WARNING int = 1
    EXIT_CRITICAL int = 2
    EXIT_UNKNOWN int = 3
)

func main() {
    ip := flag.String("ip", "192.168.0.1", "The IP of the Fritz!Box")
    flag.Parse()

    if IsUpgradeAvailable(GetInfo(*ip)) {
        ExitCritical("Firmware upgrade available")
    } else {
        ExitOK("No firmware update available")
    }
}

func GetInfo(ip string) string {
    stringReader := strings.NewReader(`
    <?xml version="1.0" encoding="utf-8"?>
       <s:Envelope s:encodingStyle="http://schemas.xmlsoap.org/soap/encoding/" xmlns:s="http://schemas.xmlsoap.org/soap/envelope/">
       <s:Body>
           <u:GetInfo xmlns:u="urn:dslforum-org:service:UserInterface:1">
           </u:GetInfo>
       </s:Body>
    </s:Envelope>
    `)

    client := &http.Client{}
    request, _ := http.NewRequest("POST", fmt.Sprintf("http://%s:49000/upnp/control/userif", ip), stringReader)
    request.Header.Set("Content-Type", "text/xml; charset=\"utf-8\"")
    request.Header.Set("SoapAction", "urn:dslforum-org:service:UserInterface:1#GetInfo")

    response, err := client.Do(request)

    if err != nil {
        ExitCritical(fmt.Sprintf("Was not able to connect to the Fritz!Box: %s", err.Error()))
    } else {
        defer response.Body.Close()

        switch response.StatusCode {
        case http.StatusOK:
            // Parse resulting XML
            getInfoContent, _ := ioutil.ReadAll(response.Body)
            return string(getInfoContent)
        case http.StatusNotFound:
            ExitWarning("Was able to contact the Fritz!Box but status info collection is disabled")
        default:
            ExitUnknown(fmt.Sprintf("Received an unexpected status code %d", response.StatusCode))
        }
    }

    return ""
}

func IsUpgradeAvailable(info string) bool {
    updateAvailable := false

    if ! strings.Contains(info, "<NewUpgradeAvailable>0</NewUpgradeAvailable>") {
        updateAvailable = true
    }

    return updateAvailable
}

func ExitUnknown(message string) {
    fmt.Printf("UNKNOWN - %s\n", message)
    os.Exit(EXIT_UNKNOWN)
}

func ExitCritical(message string) {
    fmt.Printf("CRITICAL - %s\n", message)
    os.Exit(EXIT_CRITICAL)
}

func ExitWarning(message string) {
    fmt.Printf("WARNING - %s\n", message)
    os.Exit(EXIT_WARNING)
}

func ExitOK(message string) {
    fmt.Printf("OK - %s\n", message)
    os.Exit(EXIT_OK)
}
