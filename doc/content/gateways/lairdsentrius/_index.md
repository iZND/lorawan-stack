---
title: "Laird Sentrius™ RG1xx LoRaWAN Gateway"
description: ""
---

The **Laird Sentrius™ RG1xx** is a secure, scalable, robust LoRaWAN-Enabled Gateway that supports Wi-Fi, Ethernet, and LTE backends.

This page contains information about connecting the Laird Sentrius™ RG1xx to {{% tts %}}

<!--more-->

The technical specifications can be found in [Laird's official documentation](https://www.lairdconnect.com/wireless-modules/lorawan-solutions/sentrius-rg1xx-lorawan-gateway-wi-fi-ethernet-optional-lte-us-only#documentation). The Laird Sentrius™ RG1xx LoRaWAN Gateway supports two ways to connect with {{% tts %}}, using either the Semtech Packet Router or LoRa Basic Station.

{{< figure src="laird.png" alt="Laird" >}}

## Requirements

1. User account on {{% tts %}} with rights to create gateways.
2. Laird Sentrius™ RG1xx LoRaWAN Gateway connected to the internet via ethernet or cellular.
3. CA certificate for the Basic Station (if using Basic Station).

## Registration

Create a gateway by following the instructions for the [Console]({{< ref "/getting-started/console#create-gateway" >}}) or the [CLI]({{< ref "/getting-started/cli#create-gateway" >}}).

The **EUI** of the gateway can be found on the configuration web page of the gateway as **Gateway ID**, or on the bottom label as **DevEUI**.

{{< figure src="eui.png" alt="Gateway EUI" >}}

## Configuration via Browser

The Laird Sentrius™ RG1xx LoRaWAN Gateway features a Wi-Fi quick config that allows you to configure it without connecting via ethernet. To expose the Wi-Fi access point, do the following:

1. Depress and hold the user button for 7 seconds.
2. From a wirelessly enabled device perform a scan.
3. Connect to the access point rg1xx29378B, where “29378B” are the last six digits of the Ethernet MAC address found on the label on the bottom of the gateway

> The network is secured with WPA2 with a password that is the same as the SSID.  It is recommended that the default password is changed for security reasons.  The password can be changed on the Wi-Fi -> Advanced web page.

See [Laird's official documentation](https://www.lairdconnect.com/wireless-modules/lorawan-solutions/sentrius-rg1xx-lorawan-gateway-wi-fi-ethernet-optional-lte-us-only#documentation) for more information.

### Login

To log into the gateway web interface, follow these steps:

1. Determine the last three bytes of your gateway’s Ethernet MAC address. This can be found on the label on the bottom of the gateway; the last three bytes are highlighted
2. Enter the URL into the web browser to access the web interface. For example, for the gateway used in this guide, the URL is https://rg1xx29378B.local., where “29378B” are the last 6 digits of the Ethernet MAC address. In Wi-Fi quick config mode, the gateway can also be accessed via the IP address at https://192.168.1.1 
3. Accept the self-signed security certificate in the browser.
4. Log on using the following the default username **sentrius** and default password **RG1xx**

{{< figure src="login.png" alt="Login" >}}

## Configuring Basic Station

To set up LoRa packet forwarding on the gateway, follow these steps.

Click the **LoRa** tab in the main menu.

Click on the **Forwarder** tab.

{{< figure src="plus.png" alt="Create new server" >}}

Choose **Semtech Basic Station** as the **Mode**.

Enter the CUPS and LNS configuration information provided.

Click Update.

If your configuration was successful, your gateway will connect to {{% tts %}} after a couple of seconds.
