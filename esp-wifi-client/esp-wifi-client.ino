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
    String postData = "hello from ESP";

    http.begin(wifiClient, "https://hidden-beyond-76584.herokuapp.com/website/");
    http.addHeader("Content-Type", "text/plain");

    int httpCode = http.POST(postData);
    String payload = http.getString();

    Serial.println(httpCode);
    Serial.println(payload);

    http.end();

    delay(5000);
  }