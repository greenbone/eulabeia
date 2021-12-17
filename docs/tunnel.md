# Tunneling concept

In eulabeia there can be sensors that are not in the same network and cannot reach the same mqtt broker as the director.

Since there are the requirements for each machine to have a SSH Server the decision was made to mitigate that limitation by using SSH and requiring the sensor to have a local MQTT broker running as well.

This allows us to use the same sensor setup on those machines without having to switch mechanics but just use another network.

## Director

To not having to rely on stable connections the decision was made to establish connection from the director to the host in a timely fashion to pull messages from that broker.

Since the connection will be controlled by the director and not having to cache that much data there was a decision to immediately send data to those machines.

To be capable of reusing the same client id for fetching and sending it checks if there is already a connection open and reuses it accordingly.

## Sensor
