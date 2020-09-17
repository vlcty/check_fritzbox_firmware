# check_fritzbox_firmware

Checkt eine Fritz!Box via TR-064, ob ein Firmwareupdate aussteht.

## Beispielaufruf:

Das Programm nimmt nur einen Parameter: Die IP zur Fritz!Box. Diese gibt man über den Parameter `--ip` mit an:

```
./check_fritzbox_firmware --ip 10.10.0.1
```

Wird die IP nicht mit angegeben, dann wird die IP `192.168.0.1` verwendet.

## Vorbereitungen, Security, etc

Gegebenenfalls muss [die Expertenansicht aktiviert werden](https://avm.de/service/fritzbox/fritzbox-7490/wissensdatenbank/publication/show/1652_Erweiterte-Ansicht-der-Benutzeroberflaeche-aktivieren/).

### Auslesen via TR-064 erlauben

Damit man die Info über TR-064 auslesen kann muss man dies in der Fritz!Box zuerst erlauben. Das findet man unter "Heimnetz" -> "Netzwerk" -> Tab "Netzwerkeinstellungen" -> Sektion "Heimnetzfreigaben" -> Haken setzen bei "Zugriff für Anwendungen zulassen"

Der Haken bei "Statusinformationen über UPnP übertragen" muss nicht gesetzt werden.

### Security und Authentifizierung

Die Anfragen erfolgen alle unverschlüsselt. Eine Authentifizierungung um diese Informationen auszulesen ist ebenfalls nicht nötig.

## Manuelle Update-Abfrage

Via cURL kann man die Info auch manuell oder zum debuggen auslesen:

```bash
curl -v -4 -k                                 \
     "http://192.168.0.1:49000/upnp/control/userif"                                     \
     -H 'Content-Type: text/xml; charset="utf-8"'                           \
     -H 'SoapAction: urn:dslforum-org:service:UserInterface:1#GetInfo' \
     -d '<?xml version="1.0" encoding="utf-8"?>
        <s:Envelope s:encodingStyle="http://schemas.xmlsoap.org/soap/encoding/" xmlns:s="http://schemas.xmlsoap.org/soap/envelope/">
        <s:Body>
            <u:GetInfo xmlns:u="urn:dslforum-org:service:UserInterface:1">
            </u:GetInfo>
        </s:Body>
     </s:Envelope>'

```

Gegebenenfalls die IP Adresse ändern und ausführen. Als Antwort erhält man dann XML.

Bei einem ausstehenden Update:

```xml
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
```

Bei keinem ausstehenden Update:

```xml
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
```

Das Programm wertet nur den Key `NewUpgradeAvailable` aus.

## Getestete Fritz!Box Modelle

Ich habe das Programm gegen zwei Modelle getestet:

* Fritz!Box 7490 (VDSL)
* Fritz!Box 7590 (VDSL)
* Fritz!Box 7582 (G.Fast)
* Fritz!Box 7583 (G.Fast)

## Referenzen

* [AVM Dokumentation zum UserInterface endpoint](https://avm.de/fileadmin/user_upload/Global/Service/Schnittstellen/userifSCPD.pdf)
* [Beispiel cURL Aufruf (ganz unten in der Sektion "Fritz!Box Dialer")](http://chris.cnie.de/netzwerk/fritzbox.html)
