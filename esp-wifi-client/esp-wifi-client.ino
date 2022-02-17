#include <ESP8266HTTPClient.h>
#include <ESP8266WiFi.h>

#define SSID "DandG-BHSS-2.4G"
#define PASS "Cookie23again"

WiFiClient wifiClient;

void setup()
{
  Serial.begin(9600);
  WiFi.begin(SSID, PASS);
  pinMode(2, OUTPUT);
}

  void loop()
  {
    HTTPClient http;
    String postData = "Message=Hello+from+ESP!";

    http.begin(wifiClient, "https://gonk-systems.herokuapp.com/ESP-requests");
    http.addHeader("Content-Type", "application/x-www-form-urlencoded");

    int httpCode = http.POST(postData);
    String payload = http.getString();

if (httpCode > 0) {
      // HTTP header has been send and Server response header has been handled
      Serial.printf("[HTTP] POST... code: %d\n", httpCode);

      // file found at server
      if (httpCode == HTTP_CODE_OK) {
        const String& payload = http.getString();
        Serial.println("received payload:\n<<");
        Serial.println(payload);
        Serial.println(">>");
      }
    } else {
      Serial.printf("[HTTP] POST... failed, error: %s\n", http.errorToString(httpCode).c_str());
    }
    Serial.println("here1");
    Serial.println(httpCode);
    Serial.println("here2");
    Serial.println(payload);

    http.end();

    delay(5000);
  }