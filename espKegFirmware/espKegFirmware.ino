/*********
  Rui Santos
  Complete project details at http://randomnerdtutorials.com  
*********/

// Including the ESP8266 WiFi library
#include <ESP8266WiFi.h>
#include <OneWire.h>
#include <DallasTemperature.h>
#include <WiFiClient.h>
#include <ESP8266WebServer.h>
#include <ESP8266mDNS.h>

MDNSResponder mdns;
ESP8266WebServer server(80);

#define heater1Pin 16
#define heater2Pin 5
#define ONE_WIRE_BUS 4

// Replace with your network details
const char* ssid = "WIFI_NETWORK_NAME";
const char* password = "WIFI_PASSWORD";

// Setup a oneWire instance to communicate with any OneWire devices (not just Maxim/Dallas temperature ICs)
OneWire oneWire(ONE_WIRE_BUS);

// Pass our oneWire reference to Dallas Temperature. 
DallasTemperature DS18B20(&oneWire);
char temperatureCString[6];

void heaterStatus(int pin) {
    int status = digitalRead(pin);
    server.send(200, "text/html", String(status));  
}

void toggleHeater(int pin) {
    int status = digitalRead(pin);
    if ( status == LOW ) {
      digitalWrite(pin, HIGH);
    } else {
      digitalWrite(pin, LOW);
    }
    server.send(200, "text/html", String(status));
}

// only runs once on boot
void setup() {
  // Initializing serial port for debugging purposes
  Serial.begin(115200);
  delay(10);

  // preparing GPIOs
  pinMode(heater1Pin, OUTPUT);
  digitalWrite(heater1Pin, LOW);
  pinMode(heater2Pin, OUTPUT);
  digitalWrite(heater2Pin, LOW);
  
  DS18B20.begin(); // IC Default 9 bit. If you have troubles consider upping it 12. Ups the delay giving the IC more time to process the temperature measurement
  
  // Connecting to WiFi network
  Serial.println();
  Serial.print("Connecting to ");
  Serial.println(ssid);
  
  WiFi.begin(ssid, password);
  
  while (WiFi.status() != WL_CONNECTED) {
    delay(500);
    Serial.print(".");
  }
  
  Serial.println("");
  Serial.println("WiFi connected");

  server.on("/", [](){
    server.send(200, "text/html", "OK");
  });
  
  server.on("/heater1", HTTP_GET, [](){
    heaterStatus(heater1Pin);
  });
  
  server.on("/heater2", HTTP_GET, [](){
    heaterStatus(heater2Pin);
  });

  server.on("/heater1", HTTP_POST, [](){
    toggleHeater(heater1Pin);
  });
  
  server.on("/heater2", HTTP_POST, [](){
    toggleHeater(heater2Pin);
  });
  
  server.on("/temp", [](){
    getTemperature();
    server.send(200, "text/html", temperatureCString);
  });

  server.begin();
  
  // Printing the ESP IP address
  Serial.println(WiFi.localIP());

  if (mdns.begin("esp8266", WiFi.localIP())) {
    Serial.println("MDNS responder started");
    while(1) { 
      delay(600);
    }
  }
}

void getTemperature() {
  float tempC;
  DS18B20.requestTemperatures(); 
  tempC = DS18B20.getTempCByIndex(0);
  dtostrf(tempC, 2, 2, temperatureCString);
  delay(100);
}

// runs over and over again
void loop() {
  server.handleClient();
}   
