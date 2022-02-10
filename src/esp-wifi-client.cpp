#include <ESP8266HTTPClient.h>
#include <ESP8266WiFi.h>

#define SSID "eduroam"
#define PASS "VABbad2019"

// create a static IP address for our esp

IPAddress local_IP(137, 195, 208, 7);
IPAddress gateway(137, 195, 208, 1);
IPAddress subnet(255, 255, 252, 0);

void setup()
{

  pinMode(2, OUTPUT);
  if (!WiFi.config(local_IP, gateway, subnet))
  {
    Serial.println("STA Failed to configure");

    Serial.begin(9600);     //Serial connection
    WiFi.begin(SSID, PASS); //WiFi connection

    while (WiFi.status() != WL_CONNECTED)
    { //Wait for the WiFI connection completion

      delay(500);
      Serial.println("Waiting for connection");
    }
  }
}

  void loop()
  {
    Serial.println("at start of loop");
    digitalWrite(2, HIGH);
    if (WiFi.status() == WL_CONNECTED)
    { //Check WiFi connection status

      HTTPClient http; //Declare object of class HTTPClient
/*
      http.begin("localhost:8080/website/login-success"); //Specify request destination
      http.addHeader("Content-Type", "text/plain");       //Specify content-type header

      int httpCode = http.POST("Hello from ESP8266"); //Send the request
      String payload = http.getString();              //Get the response payload

      Serial.println(httpCode); //Print HTTP return code
      Serial.println(payload);  //Print request response payload
      */
      Serial.println(WiFi.localIP());
      http.end(); //Close connection
    }
    else
    {

      Serial.println("Error in WiFi connection");
    }

    delay(1000); //Send a request every 30 seconds
  }