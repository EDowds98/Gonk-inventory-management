//ESP dependancies
#include <ESP8266HTTPClient.h>
#include <ESP8266WiFi.h>
#include <ArduinoJson.h>

#define SSID "DandG-BHSS-2.4G"
#define PASS "Cookie23again"

//#define SSID "Starlight"
//#define PASS "Omicron2311"

//#define SSID "Aidan"
//#define PASS "Flopster"


//state variables
bool lookToRead = true; //will be false in future
bool start = true;
bool sendReady = false;

//time variables
const int sampleCount = 25;
const int pulseDuration = 100;
//read start must be more than sampleCount less than pulse duration
const int sampleStart = (pulseDuration-sampleCount)/2;
// this should equal pulseDuration - sampleStart - sampleCount 
const int sigEnd = pulseDuration-(sampleStart+sampleCount);

//variables.
int clearState8[8];
int readNum = 0;
int turn = 300;

//container for storing sample array
struct cont {
  bool complete = false;
  int curSamplePos = 0;
  unsigned long readTimer = 0;
  int samples[sampleCount];
};
//stores state of a single shelf module
struct module {
  cont sampleCont;
  //readNum of 8 means checked
  int readNum = 0;
  bool state[8];
  bool complete = false;
};
//stores 8 shelf modules making up a shelf
struct shelf {
  module curMod;
  bool shelf[8][8];
  int readNum = 0;
  bool complete = false;
};

WiFiClient wifiClient;
StaticJsonDocument<1536> doc;

shelf rShelf;

void setup() {
  
  Serial.begin(115200);     //Serial connection
  WiFi.begin(SSID, PASS);   //WiFi connection

  //Wait for the WiFI connection completion
  while (WiFi.status() != WL_CONNECTED) {  delay(500); Serial.println("Waiting for connection"); }
}
void loop() {
  //replace this with RGB sensor code
  if(digitalRead(2) == HIGH) { start = true;}
  
  if (start == true && lookToRead == true) {
    rShelf = readShelf(rShelf);
    if (rShelf.complete) {
      for(int i = 0; i < 8; i++) {
        Serial.println();
        Serial.print("Module ");
        Serial.print(i);
        Serial.print(": ");
        JsonArray mod = doc.createNestedArray(String(i));
        for(int j = 0; j < 8; j++) {Serial.print(rShelf.shelf[i][j]); mod.add(rShelf.shelf[i][j]); }}
        Serial.println(doc.as<String>()); 
        start = false ;lookToRead = false; sendReady = true;
    }
    if (sendReady == true) {
      if (StringESPSend(doc) == 200) {
        sendReady == false;
      }
      delay(2000);
    }
  } 
}
shelf readShelf(shelf unit) {
  if(unit.readNum >= 8) {unit.complete = true; return unit;}
  unit.curMod = readShelfModState(unit.curMod);
  if(unit.curMod.complete == true) {
    for (int j = 0; j < 8; j++)  {
      unit.shelf[unit.readNum][j] = unit.curMod.state[j];
    }
    unit.readNum = unit.readNum+1;
    unit.curMod = moduleReset(unit.curMod);
  }
  return unit;
}

module moduleReset(module mod) {
  mod.complete = false;
  mod.readNum = 0;
  return mod;
}
module readShelfModState(module mod) {
  if(mod.readNum >= 8) {mod.complete = true; return mod;}
  mod.sampleCont = sampling(mod.sampleCont);
  if((mod.sampleCont.complete == true)) {
    mod.state[mod.readNum] = procSample(mod.sampleCont);
    mod.readNum += 1;
    mod.sampleCont = contReset(mod.sampleCont);
  }
  return mod;
}

cont contReset(cont box) {
  box.complete = false;
  box.readTimer = 0;
  box.curSamplePos = 0;
  return box;
}

cont sampling(cont box) {
  if (box.readTimer == 0) {box.readTimer = (millis()+sampleStart);
  //Serial.print("readTimer: ");
  //Serial.println(box.readTimer);
  }
  if(( box.curSamplePos < sampleCount )&&(millis() > box.readTimer)) {
    box.samples[box.curSamplePos] = analogRead(A0);
    /*
    Serial.print("time: ");
    Serial.println(box.readTimer);
    Serial.print("current sample: ");
    Serial.print1ln(box.curSamplePos);
    Serial.print("current sample value: ");
    Serial.println(box.samples[box.curSamplePos]);
   */
    box.curSamplePos = box.curSamplePos + 1;
    box.readTimer = box.readTimer +1;
  }
  // old if(box.curSamplePos >=25
  if(millis() >= box.readTimer +sigEnd) {box.complete = true;}
  return box;
}

bool procSample(cont box) {
  int val = 0;
  for( int j = 0; j < sampleCount; j++) { val += box.samples[j];}
  val = val/sampleCount;
  if (val > turn) {return true;}
  else {return false;}
}
//Sends Json data as String to Server
int StringESPSend(StaticJsonDocument<1536> data) {
 
  if (WiFi.status() == WL_CONNECTED) { //Check WiFi connection status
    HTTPClient http;
    
    String dataStr = data.as<String>();
    String message = "Message="+ dataStr; 
   
    http.begin(wifiClient, "http://gonk-systems.herokuapp.com/ESP-requests");      //Specify request destination
    http.addHeader("Content-Type", "application/json");  //Specify content-type header
 
    //int httpCode = http.POST(message);   //Send the request
    int httpCode = http.POST(dataStr);   //Send the request
    String payload = http.getString();                 //Get the response payload
 
    Serial.println(httpCode);   //Print HTTP return code
    Serial.println(payload);    //Print request response payload
 
    http.end();  //Close connection
    return httpCode;
  } else { Serial.println("Error in WiFi connection"); return -1;}
}
